package socket

import (
	"encoding/json"
	"slices"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

type Server struct {
	melody   *melody.Melody
	rooms    map[string][]*melody.Session
	rmu      sync.RWMutex
	handlers map[string]EventHandler
}

func NewServer() Server {
	m := melody.New()
	return Server{
		melody:   m,
		rooms:    make(map[string][]*melody.Session),
		handlers: make(map[string]EventHandler),
	}
}

func (server *Server) InitListeners(router *gin.Engine) {
	router.GET("/ws", func(ctx *gin.Context) {
		server.melody.HandleRequest(ctx.Writer, ctx.Request)
	})

	server.On(JOIN, server.AddToRoom)
	server.On(LEAVE, server.RemoveFromRoom)

	server.melody.HandleMessage(func(s *melody.Session, msg []byte) {
		var payload EventPayload
		err := json.Unmarshal(msg, &payload)
		if err != nil {
			return
		}
		if handler, ok := server.handlers[payload.Name]; ok {
			handler(Event{
				Name:   payload.Name,
				Data:   payload.Data,
				Socket: s,
			})
		}
	})

	server.melody.HandleDisconnect(func(s *melody.Session) {
		server.rmu.Lock()
		defer server.rmu.Unlock()
		for roomID, sessions := range server.rooms {
			for i, session := range sessions {
				if session == s {
					server.rooms[roomID] = slices.Delete(server.rooms[roomID], i, i+1)
					break
				}
			}
		}
	})
}

func (server *Server) On(event string, handler EventHandler) {
	server.handlers[event] = handler
}

func (server *Server) To(roomID string, data any) error {
	server.rmu.RLock()
	defer server.rmu.RUnlock()

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if sessions, ok := server.rooms[roomID]; ok {
		return server.melody.BroadcastMultiple(dataBytes, sessions)
	}
	return nil
}
