package realtime

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/websocket"
)

var totalID int = 0

type Client struct {
	id     int
	ws     *websocket.Conn
	server *Server
	ch     chan string
	doneCh chan bool
}

func NewClient(ws *websocket.Conn, server *Server) *Client {
	totalID++

	return &Client{
		id:     totalID,
		ws:     ws,
		server: server,
		ch:     make(chan string, 10),
		doneCh: make(chan bool),
	}
}

func (client *Client) Conn() *websocket.Conn {
	return client.ws
}

func (client *Client) Write(message string) {
	select {
	case client.ch <- message:
	default:
		client.server.Remove(client)
		client.server.Error(fmt.Errorf("Disconnect:", client.id))
	}
}

func (client *Client) Done() {
	client.doneCh <- true
}

func (client *Client) Listen() {
	go client.listenWrite()
	client.listenRead()
}

func (client *Client) listenWrite() {
	for {
		select {

		case message := <-client.ch:
			log.Println("Send:", message)
			err := websocket.Message.Send(client.ws, message)
			if err != nil {
				log.Println("Error:", err)
			}

		case <-client.doneCh:
			client.server.Remove(client)
			client.doneCh <- true
			return
		}
	}
}

func (client *Client) listenRead() {
	for {
		select {

		case <-client.doneCh:
			client.server.Remove(client)
			client.doneCh <- true
			return

		default:
			var message string
			err := websocket.Message.Receive(client.ws, &message)
			if err == io.EOF {
				client.doneCh <- true
			} else if err != nil {
				client.server.Error(err)
			} else {
				client.server.Broadcast(message)
			}
		}
	}
}
