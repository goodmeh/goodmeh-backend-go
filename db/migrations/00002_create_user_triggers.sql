-- +goose Up
-- +goose StatementBegin
CREATE FUNCTION on_review_change() RETURNS TRIGGER AS $$ BEGIN IF TG_OP = 'INSERT'
AND LENGTH(NEW.text) > 200 THEN
UPDATE "user"
SET long_review_count = long_review_count + 1
WHERE id = NEW.user_id;
ELSEIF TG_OP = 'UPDATE'
AND (LENGTH(OLD.text) > 200) != (LENGTH(NEW.text) > 200) THEN
UPDATE "user"
SET long_review_count = long_review_count + (
        CASE
            WHEN LENGTH(NEW.text) > 200 THEN 1
            ELSE -1
        END
    )
WHERE id = NEW.user_id;
END IF;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER review_change_trigger
AFTER
INSERT
    OR
UPDATE ON review FOR EACH ROW EXECUTE FUNCTION on_review_change();
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER review_change_trigger ON review;
DROP FUNCTION on_review_change();
-- +goose StatementEnd