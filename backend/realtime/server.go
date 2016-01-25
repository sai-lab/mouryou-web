package realtime

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	pattern     string
	clients     map[int]*Client
	addCh       chan *Client
	removeCh    chan *Client
	broadcastCh chan string
	doneCh      chan bool
	errorCh     chan error
}

func NewServer(pattern string) *Server {
	return &Server{
		pattern:     pattern,
		clients:     make(map[int]*Client),
		addCh:       make(chan *Client),
		removeCh:    make(chan *Client),
		broadcastCh: make(chan string),
		doneCh:      make(chan bool),
		errorCh:     make(chan error),
	}
}

func (server *Server) Add(client *Client) {
	server.addCh <- client
}

func (server *Server) Remove(client *Client) {
	server.removeCh <- client
}

func (server *Server) Broadcast(message string) {
	server.broadcastCh <- message
}

func (server *Server) Done() {
	server.doneCh <- true
}

func (server *Server) Error(err error) {
	server.errorCh <- err
}

func (server *Server) broadcast(message string) {
	for _, client := range server.clients {
		client.Write(message)
	}
}

func (server *Server) Listen() {
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				server.errorCh <- err
			}
		}()

		client := NewClient(ws, server)
		server.Add(client)
		client.Listen()
	}

	http.Handle(server.pattern, websocket.Handler(onConnected))

	for {
		select {

		case client := <-server.addCh:
			log.Println("Add:", client.id)
			server.clients[client.id] = client

		case client := <-server.removeCh:
			log.Println("Delete:", client.id)
			delete(server.clients, client.id)

		case message := <-server.broadcastCh:
			log.Println("Broadcast:", message)
			server.broadcast(message)

		case err := <-server.errorCh:
			log.Println("Error:", err.Error())

		case <-server.doneCh:
			return
		}
	}
}
