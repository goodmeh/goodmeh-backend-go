package summarizer

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"goodmeh/app/utils/arr"
	"net/http"
	"os"
	"time"

	"github.com/invopop/jsonschema"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/tiktoken-go/tokenizer"
)

const (
	TOKEN_LIMIT            = 110000
	REVIEW_COUNT_THRESHOLD = 10
	TEXT_LENGTH_THRESHOLD  = 1000
)

type SummariesSchema struct {
	Summary         string `json:"summary"`
	BusinessSummary string `json:"business_summary"`
}

type IndividualSummaryInput struct {
	ID   string
	Text string
}

type OpenAiBatchRequestInput struct {
	CustomId string                        `json:"custom_id"`
	Method   string                        `json:"method"`
	Body     any                           `json:"body"`
	Url      openai.BatchNewParamsEndpoint `json:"url"`
}

type OpenAiBatchRequestOutput struct {
	CustomId string `json:"custom_id"`
	Error    *openai.ErrorObject
	ID       string `json:"id"`
	Response *struct {
		RequestID  string `json:"request_id"`
		StatusCode int    `json:"status_code"`
		Body       any    `json:"body"`
	}
}

func generateSchema[T any]() interface{} {
	// Structured Outputs uses a subset of JSON schema
	// These flags are necessary to comply with the subset
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)
	return schema
}

var SUMMARY_RESPONSE_FORMAT = openai.ChatCompletionNewParamsResponseFormatUnion{
	OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{
		JSONSchema: openai.ResponseFormatJSONSchemaJSONSchemaParam{
			Name:   "SummariesSchema",
			Schema: generateSchema[SummariesSchema](),
		},
	},
}

type OpenAiSummarizer struct {
	Client openai.Client
}

func NewSummarizer() OpenAiSummarizer {
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
	)
	return OpenAiSummarizer{
		Client: client,
	}
}

func (s OpenAiSummarizer) preprocessReviews(text []string) (string, error) {
	encoding, err := tokenizer.ForModel(tokenizer.GPT4o)
	if err != nil {
		return "", err
	}

	tokenCount := 0
	for i, review := range text {
		encodedReview, _, err := encoding.Encode(review)
		if err != nil {
			return "", err
		}
		tokenCount += len(encodedReview)
		if tokenCount > TOKEN_LIMIT {
			text = text[:i]
			break
		}
	}

	result, err := json.Marshal(text)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (s OpenAiSummarizer) summarize(combinedReviews string) (string, error) {
	chatCompletion, err := s.Client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT4oMini,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage("The data you will be given to work with are a collection of individual reviews of a place on Google maps as submitted by users. The goal is to generate two summaries with a different purpose for a different audience, based on the same set of review data. The first summary is intended for potential customers. The second summary is intended for business owners or managers. Use the following steps in order to generate the two summaries with the highest quality and accuracy. Step 1 - Read in all of the data provided to establish a context and understanding. If there are many reviews, identify key details and important  information that should be reflected in the summary. If there are few reviews, do not generate information or detail where there is none, if there is insufficient data to provide a good summary, you can respond with a statement to state that. Step 2 - Adopt the persona of a connoisseur and professional critic. This summary should let customers understand the essence of the place with the main aim of aiding them in deciding if they would enjoy it or not. Step 3 - Adopt the persona of a consultant and advisor. This summary should help business owners or managers understand what they are doing well by highlighting what customers are delighted by. Also help them improve their operations by surfacing any blindspots by including customer feedback if any. Keep the summaries to a maximum of 6 sentences each. You can use markdown bold and italic formatting to highlight key points."),
			openai.UserMessage(combinedReviews),
		},
		ResponseFormat: SUMMARY_RESPONSE_FORMAT,
	})
	if err != nil {
		return "", err
	}
	return chatCompletion.Choices[0].Message.Content, nil
}

func (s OpenAiSummarizer) postProcessSummary(output string) SummariesSchema {
	var summaries SummariesSchema
	err := json.Unmarshal([]byte(output), &summaries)
	if err != nil {
		return SummariesSchema{}
	}
	return summaries
}

func (s OpenAiSummarizer) SummarizeReviews(text []string) (SummariesSchema, error) {
	if len(text) < TEXT_LENGTH_THRESHOLD && arr.Reduce(text, func(count int, reviewText string) int {
		return count + len(reviewText)
	}, 0) < TEXT_LENGTH_THRESHOLD {
		return SummariesSchema{}, nil
	}

	combinedReviews, err := s.preprocessReviews(text)
	if err != nil {
		return SummariesSchema{}, err
	}
	ouptut, err := s.summarize(combinedReviews)
	if err != nil {
		return SummariesSchema{}, err
	}
	return s.postProcessSummary(ouptut), nil
}

func (s OpenAiSummarizer) SummarizeIndividualReviews(reviews []IndividualSummaryInput) (string, error) {
	if len(reviews) == 0 {
		return "", nil
	}
	fileName := fmt.Sprintf("%d_reviews.jsonl", time.Now().Unix())
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			return
		}
		err = os.Remove(fileName)
		if err != nil {
			return
		}
	}()
	for _, review := range reviews {
		body := openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage("The data provided is an individual review of a place on Google maps as submitted by a user. The objective is to summarize the review. This is done twice with a different persona and objective each time. One version is intended for potential customers. The other version is intended for business managers. Use the following steps in order to generate the two summaries. Keep the summaries to a maximum of 3 sentences each. Step 1 - Process the review data provided to establish a context and understanding. Identify key details and important information that should be reflected in the summary. Step 2 - Adopt the persona and voice of the user. The summary should look and sound like it was from the original user. Adopt their language and writing style, but focus on distilling and improving the clarity and quality of the text. Step 3 - Adopt the persona of a business analyst looking at customer feedback to extract insights. Pay attention to what the user was pleased or pained by. Adopt an objective and neutral writing style and tone. Focus on extracting valuable insight from the review data."),
				openai.UserMessage(review.Text),
			},
			Model:          openai.ChatModelGPT4oMini,
			ResponseFormat: SUMMARY_RESPONSE_FORMAT,
		}
		json, err := json.Marshal(OpenAiBatchRequestInput{
			CustomId: review.ID,
			Method:   http.MethodPost,
			Body:     body,
			Url:      openai.BatchNewParamsEndpointV1ChatCompletions,
		})
		if err != nil {
			return "", err
		}
		if _, err = f.Write(append(json, '\n')); err != nil {
			return "", err
		}
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return "", err
	}
	file, err := s.Client.Files.New(context.TODO(), openai.FileNewParams{
		File:    f,
		Purpose: openai.FilePurposeBatch,
	})
	if err != nil {
		return "", err
	}
	batchJob, err := s.Client.Batches.New(context.TODO(), openai.BatchNewParams{
		CompletionWindow: openai.BatchNewParamsCompletionWindow24h,
		Endpoint:         openai.BatchNewParamsEndpointV1ChatCompletions,
		InputFileID:      file.ID,
	})
	if err != nil {
		return "", err
	}
	return batchJob.ID, nil
}

func (s OpenAiSummarizer) GetBatchJobResult(jobId string) (map[string]SummariesSchema, error) {
	results := make(map[string]SummariesSchema)
	batchJob, err := s.Client.Batches.Get(context.TODO(), jobId)
	if err != nil {
		return nil, err
	}
	if batchJob.Status == openai.BatchStatusFailed ||
		batchJob.Status == openai.BatchStatusCancelling ||
		batchJob.Status == openai.BatchStatusCancelled ||
		batchJob.Status == openai.BatchStatusExpired {
		return nil, fmt.Errorf("batch job failed")
	}
	if batchJob.Status != openai.BatchStatusCompleted {
		return nil, nil
	}
	response, err := s.Client.Files.Content(context.TODO(), batchJob.OutputFileID)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	reader := bufio.NewReader(response.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		var (
			result   OpenAiBatchRequestOutput
			response openai.ChatCompletion
		)
		err = json.Unmarshal(line, &result)
		if err != nil {
			return nil, err
		}
		response, ok := result.Response.Body.(openai.ChatCompletion)
		if !ok {
			return nil, fmt.Errorf("invalid response type")
		}
		results[result.CustomId] = s.postProcessSummary(response.Choices[0].Message.Content)
	}
	return results, nil
}
