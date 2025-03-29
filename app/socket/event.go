package socket

import (
	"slices"

	"github.com/olahol/melody"
)

const (
	JOIN  = "join"
	LEAVE = "leave"
)

type Event struct {
	Name   string
	Data   any
	Socket *melody.Session
}

type EventPayload struct {
	Name string `json:"name"`
	Data any    `json:"data"`
}

type EventHandler func(Event)

func (server *Server) AddToRoom(event Event) {
	server.rmu.Lock()
	defer server.rmu.Unlock()

	roomID := event.Data.(string)

	if _, ok := server.rooms[roomID]; !ok {
		server.rooms[roomID] = []*melody.Session{}
	}

	server.rooms[roomID] = append(server.rooms[roomID], event.Socket)
}

func (server *Server) RemoveFromRoom(event Event) {
	server.rmu.Lock()
	defer server.rmu.Unlock()

	roomID := event.Data.(string)

	if sessions, ok := server.rooms[roomID]; ok {
		for i, s := range sessions {
			if s == event.Socket {
				server.rooms[roomID] = slices.Delete(server.rooms[roomID], i, i+1)
				return
			}
		}
	}
}
