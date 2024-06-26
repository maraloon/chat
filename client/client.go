package client

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	nickname string
	hub      *Hub

	// The websocket connection.
	websocket *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.websocket.Close()
	}()
	c.websocket.SetReadLimit(maxMessageSize)
	c.websocket.SetReadDeadline(time.Now().Add(pongWait))
	c.websocket.SetPongHandler(func(string) error { c.websocket.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.websocket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		// nicknameBytes := []byte(c.nickname + ": ")
		// newMessage := make([]byte, 0, len(nicknameBytes)+len(message))
		// newMessage = append(newMessage, nicknameBytes...)
		// newMessage = append(newMessage, message...)

		// parts := []string{c.nickname, ": ", string(message)}

		var builder strings.Builder
		for _, part := range []string{c.nickname, ": ", string(message)} {
			builder.WriteString(part)
		}
		messageWithNickname := []byte(builder.String())

		message = bytes.TrimSpace(bytes.Replace(messageWithNickname, newline, space, -1))
		c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.websocket.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.websocket.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.websocket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.websocket.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.websocket.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.websocket.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{
		nickname:  generateNickname(),
		hub:       hub,
		websocket: websocket,
		send:      make(chan []byte, 256),
	}
	hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

func generateNickname() string {
	length := 6
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define the alphabet
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	// Create a buffer to store the generated string
	result := make([]byte, length)

	// Generate random characters from the alphabet
	for i := 0; i < length; i++ {
		result[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(result)
}
