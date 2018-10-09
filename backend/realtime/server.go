package realtime

import (
	"log"
	"net/http"
	"strings"

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

func (server *Server) updateOperation(message string) {
	str := strings.Split(message, ":")
	if strings.Contains(str[0], "Loads") {
		return
	}

	operation := strings.ToLower(str[0])
	replacedOperation := strings.Replace(operation, " ", "_", -1)
	vmid := strings.TrimSpace(str[1])
	// TODO: vendors id set
	req, err := http.NewRequest("PUT", "http://0.0.0.0:8000/api/cluster/vendors/0/virtual_machines/"+vmid+"/"+replacedOperation, nil)
	if err != nil {
		log.Println(err)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
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
			server.updateOperation(message)

		case err := <-server.errorCh:
			log.Println("Error:", err.Error())

		case <-server.doneCh:
			return
		}
	}
}
