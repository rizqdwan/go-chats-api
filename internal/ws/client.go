package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type Message struct {
	Content string `json:"content"`
	RoomID string `json:"roomId"`
	Username string `json:"username"`
}

type Client struct {
	Conn *websocket.Conn
	Message chan *Message
	ID string `json:"id"`
	RoomID string `json:"roomId"`
	Username string `json:"username"`
}

func (cl *Client) writeMessage() {
	defer func ()  {
		cl.Conn.Close()
	}()

	for {
		msg, ok := <-cl.Message
		if !ok {
			return
		}

		cl.Conn.WriteJSON(msg)
	}
}

func (cl *Client) readMessage(hub *Hub) {
	defer func ()  {
		hub.Unregister <- cl
		cl.Conn.Close()	
	}()

	for {
		_, m, err := cl.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}

			break
		}

		msg := &Message{
			Content: string(m),
			RoomID: cl.RoomID,
			Username: cl.Username,
		}

		hub.Broadcast <- msg
	}
}