package __obf_eb68a9852b2589b3

import (
	"encoding/json"
	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/gorilla/websocket"
	"strings"
	"sync"
	"time"
)

// constants for action type
const (
	// from client
	__obf_d2b06147fc286e86 = "publish"
	__obf_ac323d22c93b0b4e = "subscribe"
	__obf_fd51da43c5104cf4 = "unsubscribe"
	// to client
	// data  = "data"
	__obf_4d6ea3f46b7e2c9c = "kickout"
	__obf_425ec3ee9a07a881 = "throw"
	__obf_8e6c449303e04ad8 = "greet"
	// popup = "popup"
)

// constants for server message
const (
	ErrInvalidMessage       = "Server: Invalid msg"
	ErrActionUnrecognizable = "Server: Action unrecognized"
)

type Topic string

// subscriber is a type for each string of topic and the clients that subscribe to it
type __obf_b08ccebf62d7531f map[Topic]Client

type ClientID string

// Client is a type that describe the clients' ID and their connection
type Client map[ClientID]*websocket.Conn

// Message is a struct for message to be sent by the client
type Message struct {
	Action string `json:"action"`
	Topic  `json:"topic"`
	Body   any `json:"body"`
}

type Option struct {
	// time period to send pings to client
	PingPeriod time.Duration `yaml:"ping-period"` // seconds, (pongWait * 9) / 10
	// time to read the next client's pong message
	PongWait time.Duration `yaml:"pong-wait"` // seconds, 60 * time.Second
	// time allowed to write a message to client
	WriteWait time.Duration `yaml:"write-wait"` // seconds, 10 * time.Second
	// max message size allowed
	MaxMessageSize int64 `yaml:"max-message-size"` // 512
}

type Handle func(ID ClientID, __obf_4d26230e13365e72 Topic, __obf_3fe025c030f856e9 string)

// Server is the struct to handle the Server functions & manage the subscriber
type Server struct {
	Option
	__obf_eecc7d886d9dcb14 *sync.WaitGroup
	__obf_922c595b3c5ec9c3 *sync.RWMutex
	__obf_66375fc7db98b5e6 *sync.RWMutex
	__obf_b08ccebf62d7531f __obf_b08ccebf62d7531f
	__obf_24dbb11615c97182 map[string]Handle
}

func NewPubsub(__obf_f6a22fd36b0e60c0 Option) *Server {
	return &Server{
		Option:                 __obf_f6a22fd36b0e60c0,
		__obf_eecc7d886d9dcb14: &sync.WaitGroup{},
		__obf_922c595b3c5ec9c3: &sync.RWMutex{},
		__obf_66375fc7db98b5e6: &sync.RWMutex{},
		__obf_b08ccebf62d7531f: make(__obf_b08ccebf62d7531f),
	}
}

func (__obf_c301d4a0b28bafec *Server) SetActionHandle(__obf_14058785daa2b25f string, __obf_6ebfbb12dd1284f3 Handle) {

	if __obf_c301d4a0b28bafec.__obf_24dbb11615c97182 == nil {
		__obf_c301d4a0b28bafec.__obf_24dbb11615c97182 = map[string]Handle{
			__obf_14058785daa2b25f: __obf_6ebfbb12dd1284f3,
		}
	} else {
		__obf_c301d4a0b28bafec.__obf_24dbb11615c97182[__obf_14058785daa2b25f] = __obf_6ebfbb12dd1284f3
	}
}

func (__obf_c301d4a0b28bafec *Server) Onlines(__obf_c7132fd45a037f93 Topic) int {
	return len(__obf_c301d4a0b28bafec.__obf_b08ccebf62d7531f[__obf_c7132fd45a037f93])
}

// Send simply sends message to the websocket client
func (__obf_c301d4a0b28bafec *Server) __obf_985bf35c48714f17(__obf_45329568b5aa0723 *websocket.Conn, __obf_14058785daa2b25f string, __obf_3fe025c030f856e9 any) {
	// send simple message

	// conn.WriteJSON(Message{Action: action, Body: msg})

	// bs, _ := tool.AnyToBytes(Message{Action: action, Body: msg})
	// conn.WriteMessage(websocket.BinaryMessage, bs)
	var __obf_680c3c9773facaf0 []byte
	__obf_680c3c9773facaf0, _ = util.Encode(Message{Action: __obf_14058785daa2b25f, Body: __obf_3fe025c030f856e9})

	__obf_c301d4a0b28bafec.__obf_922c595b3c5ec9c3.Lock()
	__obf_45329568b5aa0723.WriteMessage(websocket.TextMessage, __obf_680c3c9773facaf0)
	__obf_c301d4a0b28bafec.__obf_922c595b3c5ec9c3.Unlock()
}

func (__obf_c301d4a0b28bafec *Server) GetConn(__obf_c7132fd45a037f93 Topic, __obf_1a3b2d50cbc0bb42 ClientID) *websocket.Conn {
	__obf_3eeb596686cf85a1, __obf_690b61241d5b398e := __obf_c301d4a0b28bafec.__obf_b08ccebf62d7531f[__obf_c7132fd45a037f93]
	if !__obf_690b61241d5b398e {
		return nil
	}

	var __obf_45329568b5aa0723 *websocket.Conn
	if __obf_45329568b5aa0723, __obf_690b61241d5b398e = __obf_3eeb596686cf85a1[__obf_1a3b2d50cbc0bb42]; !__obf_690b61241d5b398e {
		return nil
	}
	return __obf_45329568b5aa0723
}

func (__obf_c301d4a0b28bafec *Server) SendTo(__obf_c7132fd45a037f93 Topic, __obf_1a3b2d50cbc0bb42 ClientID, __obf_14058785daa2b25f string, __obf_3fe025c030f856e9 any) {
	__obf_45329568b5aa0723 := __obf_c301d4a0b28bafec.GetConn(__obf_c7132fd45a037f93, __obf_1a3b2d50cbc0bb42)
	if __obf_45329568b5aa0723 != nil {
		__obf_c301d4a0b28bafec.__obf_985bf35c48714f17(__obf_45329568b5aa0723, __obf_14058785daa2b25f, __obf_3fe025c030f856e9)
	}
}

// SendWithWait sends message to the websocket client using wait group, allowing usage with goroutines
func (__obf_c301d4a0b28bafec *Server) SendWithWait(__obf_45329568b5aa0723 *websocket.Conn, __obf_14058785daa2b25f string, __obf_3fe025c030f856e9 any) {
	// send simple message
	// conn.WriteMessage(websocket.TextMessage, []byte(tool.EncodeString(msg)))
	__obf_c301d4a0b28bafec.__obf_985bf35c48714f17(__obf_45329568b5aa0723, __obf_14058785daa2b25f, __obf_3fe025c030f856e9)

	// set the task as done
	__obf_c301d4a0b28bafec.__obf_eecc7d886d9dcb14.Done()
}

func (__obf_c301d4a0b28bafec *Server) SendGreet(__obf_45329568b5aa0723 *websocket.Conn, __obf_3fe025c030f856e9 any) {
	__obf_c301d4a0b28bafec.__obf_985bf35c48714f17(__obf_45329568b5aa0723, __obf_8e6c449303e04ad8, __obf_3fe025c030f856e9)
}

func (__obf_c301d4a0b28bafec *Server) SendWithAction(__obf_45329568b5aa0723 *websocket.Conn, __obf_14058785daa2b25f string, __obf_3fe025c030f856e9 any) {
	__obf_c301d4a0b28bafec.__obf_985bf35c48714f17(__obf_45329568b5aa0723, __obf_14058785daa2b25f, __obf_3fe025c030f856e9)
}

// func (s *Server) SendData(conn *websocket.Conn, msg any) {
// 	s.send(conn, data, msg)
// }

func (__obf_c301d4a0b28bafec *Server) ThrowError(__obf_45329568b5aa0723 *websocket.Conn, __obf_4c0a65e7eb72a3b4 string) {
	__obf_c301d4a0b28bafec.__obf_985bf35c48714f17(__obf_45329568b5aa0723, __obf_425ec3ee9a07a881, __obf_4c0a65e7eb72a3b4)
}

// RemoveClient removes the clients from the server subscription map
func (__obf_c301d4a0b28bafec *Server) RemoveClient(__obf_1a3b2d50cbc0bb42 ClientID) {
	__obf_c301d4a0b28bafec.__obf_66375fc7db98b5e6.Lock()
	defer __obf_c301d4a0b28bafec.__obf_66375fc7db98b5e6.Unlock()
	// loop all topics
	for __obf_f01d727059f8b47c := range __obf_c301d4a0b28bafec.__obf_b08ccebf62d7531f {
		// delete the client from all the topic's client map
		delete(__obf_c301d4a0b28bafec.__obf_b08ccebf62d7531f[__obf_f01d727059f8b47c], __obf_1a3b2d50cbc0bb42)
	}
}

// ProcessMessage handle message according to the action type
func (__obf_c301d4a0b28bafec *Server) ProcessMessage(__obf_45329568b5aa0723 *websocket.Conn, __obf_1a3b2d50cbc0bb42 ClientID, __obf_3fe025c030f856e9 []byte) {
	// parse message
	__obf_6a35cef3f2dfe6ce := Message{}
	if __obf_4c0a65e7eb72a3b4 := json.Unmarshal(__obf_3fe025c030f856e9, &__obf_6a35cef3f2dfe6ce); __obf_4c0a65e7eb72a3b4 != nil {
		__obf_c301d4a0b28bafec.ThrowError(__obf_45329568b5aa0723, ErrInvalidMessage)
		return
	}

	// convert all action to lowercase and remove whitespace
	__obf_14058785daa2b25f := strings.TrimSpace(strings.ToLower(__obf_6a35cef3f2dfe6ce.Action))

	switch __obf_14058785daa2b25f {
	case __obf_d2b06147fc286e86:
		__obf_c301d4a0b28bafec.Publish(__obf_6a35cef3f2dfe6ce.Topic, __obf_6a35cef3f2dfe6ce.Body)

	case __obf_ac323d22c93b0b4e:
		__obf_c301d4a0b28bafec.Subscribe(__obf_45329568b5aa0723, __obf_1a3b2d50cbc0bb42, __obf_6a35cef3f2dfe6ce.Topic)
	case __obf_fd51da43c5104cf4:
		__obf_c301d4a0b28bafec.Unsubscribe(__obf_1a3b2d50cbc0bb42, __obf_6a35cef3f2dfe6ce.Topic)

	default:
		if __obf_38b05b3e85bb2579, __obf_a243795e2b7cacf4 := __obf_c301d4a0b28bafec.__obf_24dbb11615c97182[__obf_14058785daa2b25f]; __obf_a243795e2b7cacf4 {
			__obf_38b05b3e85bb2579(__obf_1a3b2d50cbc0bb42, __obf_6a35cef3f2dfe6ce.Topic, __obf_6a35cef3f2dfe6ce.Body.(string))
		} else {
			__obf_c301d4a0b28bafec.ThrowError(__obf_45329568b5aa0723, ErrActionUnrecognizable)
		}

	}
}

// Publish sends a message to all subscribing clients of a topic
func (__obf_c301d4a0b28bafec *Server) Publish(__obf_4d26230e13365e72 Topic, __obf_c9c5b49a5771b343 any) {

	// if topic does not exist, stop the process
	if _, __obf_690b61241d5b398e := __obf_c301d4a0b28bafec.__obf_b08ccebf62d7531f[__obf_4d26230e13365e72]; !__obf_690b61241d5b398e {
		return
	}

	// if topic exist
	__obf_3eeb596686cf85a1 := __obf_c301d4a0b28bafec.__obf_b08ccebf62d7531f[__obf_4d26230e13365e72]

	// send the message to the clients
	for _, __obf_45329568b5aa0723 := range __obf_3eeb596686cf85a1 {
		// add 1 job to wait group
		__obf_c301d4a0b28bafec.__obf_eecc7d886d9dcb14.Add(1)

		// send with goroutines
		go __obf_c301d4a0b28bafec.SendWithWait(__obf_45329568b5aa0723, __obf_d2b06147fc286e86, __obf_c9c5b49a5771b343)
	}

	// wait until all goroutines jobs done
	__obf_c301d4a0b28bafec.__obf_eecc7d886d9dcb14.Wait()
}

// Kickout: kick the older session out
func (__obf_c301d4a0b28bafec *Server) Kickout(__obf_c7132fd45a037f93 Topic, __obf_1a3b2d50cbc0bb42 ClientID) {
	__obf_45329568b5aa0723 := __obf_c301d4a0b28bafec.GetConn(__obf_c7132fd45a037f93, __obf_1a3b2d50cbc0bb42)
	if __obf_45329568b5aa0723 != nil {
		__obf_c301d4a0b28bafec.__obf_985bf35c48714f17(__obf_45329568b5aa0723, __obf_4d6ea3f46b7e2c9c, nil)
		__obf_45329568b5aa0723.Close()
		__obf_c301d4a0b28bafec.__obf_66375fc7db98b5e6.Lock()
		delete(__obf_c301d4a0b28bafec.__obf_b08ccebf62d7531f[__obf_c7132fd45a037f93], __obf_1a3b2d50cbc0bb42)
		__obf_c301d4a0b28bafec.__obf_66375fc7db98b5e6.Unlock()
	}
}

// Subscribe adds a client to a topic's client map
func (__obf_c301d4a0b28bafec *Server) Subscribe(__obf_45329568b5aa0723 *websocket.Conn, __obf_1a3b2d50cbc0bb42 ClientID, __obf_4d26230e13365e72 Topic) {
	__obf_c301d4a0b28bafec.__obf_66375fc7db98b5e6.Lock()
	defer __obf_c301d4a0b28bafec.__obf_66375fc7db98b5e6.Unlock()
	// if topic exist, check the client map
	if _, __obf_690b61241d5b398e := __obf_c301d4a0b28bafec.__obf_b08ccebf62d7531f[__obf_4d26230e13365e72]; __obf_690b61241d5b398e {
		__obf_3eeb596686cf85a1 := __obf_c301d4a0b28bafec.__obf_b08ccebf62d7531f[__obf_4d26230e13365e72]

		// if subbed, ok := s.actionHandle[subscribe]; ok {
		// 	subbed(id, topic, "")
		// }

		__obf_3eeb596686cf85a1[__obf_1a3b2d50cbc0bb42] = __obf_45329568b5aa0723
		return
	}

	// if topic does not exist, create a new topic
	__obf_a0ea53c5808ff5b3 := make(Client)
	__obf_c301d4a0b28bafec.__obf_b08ccebf62d7531f[__obf_4d26230e13365e72] = __obf_a0ea53c5808ff5b3

	// add the client to the topic
	__obf_c301d4a0b28bafec.__obf_b08ccebf62d7531f[__obf_4d26230e13365e72][__obf_1a3b2d50cbc0bb42] = __obf_45329568b5aa0723
}

// Unsubscribe removes a clients from a topic's client map
func (__obf_c301d4a0b28bafec *Server) Unsubscribe(__obf_1a3b2d50cbc0bb42 ClientID, __obf_4d26230e13365e72 Topic) {
	__obf_c301d4a0b28bafec.__obf_66375fc7db98b5e6.Lock()
	defer __obf_c301d4a0b28bafec.__obf_66375fc7db98b5e6.Unlock()
	// if topic exist, check the client map
	if _, __obf_690b61241d5b398e := __obf_c301d4a0b28bafec.__obf_b08ccebf62d7531f[__obf_4d26230e13365e72]; __obf_690b61241d5b398e {
		__obf_3eeb596686cf85a1 := __obf_c301d4a0b28bafec.__obf_b08ccebf62d7531f[__obf_4d26230e13365e72]

		// remove the client from the topic's client map
		delete(__obf_3eeb596686cf85a1, __obf_1a3b2d50cbc0bb42)
	}
}

// Pump just do readDump and writePump
func (__obf_c301d4a0b28bafec *Server) Pump(__obf_1a3b2d50cbc0bb42 ClientID, __obf_45329568b5aa0723 *websocket.Conn) {
	// create channel to signal client health
	__obf_01602bf362dbf019 := make(chan struct{})

	go __obf_c301d4a0b28bafec.__obf_d39f634f3f7b9af8(__obf_45329568b5aa0723, __obf_1a3b2d50cbc0bb42, __obf_01602bf362dbf019)
	__obf_c301d4a0b28bafec.__obf_018b4e744e76f55d(__obf_45329568b5aa0723, __obf_1a3b2d50cbc0bb42, __obf_01602bf362dbf019)
}

// readPump process incoming messages and set the settings
func (__obf_c301d4a0b28bafec *Server) __obf_018b4e744e76f55d(__obf_45329568b5aa0723 *websocket.Conn, __obf_1a3b2d50cbc0bb42 ClientID, __obf_01602bf362dbf019 chan<- struct{}) {
	// set limit, deadline to read & pong handler
	__obf_45329568b5aa0723.SetReadLimit(__obf_c301d4a0b28bafec.MaxMessageSize)
	__obf_45329568b5aa0723.SetReadDeadline(time.Now().Add(__obf_c301d4a0b28bafec.PongWait * time.Second))
	__obf_45329568b5aa0723.SetPongHandler(func(string) error {
		__obf_45329568b5aa0723.SetReadDeadline(time.Now().Add(__obf_c301d4a0b28bafec.PongWait * time.Second))
		return nil
	})
	// conn.SetCloseHandler(func(code int, text string) error {
	// 	s.RemoveClient(id)
	// 	// set health status to unhealthy by closing channel
	// 	close(done)
	// 	return nil
	// })

	// message handling
	for {
		// read incoming message
		_, __obf_3fe025c030f856e9, __obf_4c0a65e7eb72a3b4 := __obf_45329568b5aa0723.ReadMessage()
		// if error occured
		if __obf_4c0a65e7eb72a3b4 != nil {
			// remove from the client
			__obf_c301d4a0b28bafec.RemoveClient(__obf_1a3b2d50cbc0bb42)
			// set health status to unhealthy by closing channel
			close(__obf_01602bf362dbf019)
			// stop process
			break
		}

		// if no error, process incoming message
		__obf_c301d4a0b28bafec.ProcessMessage(__obf_45329568b5aa0723, __obf_1a3b2d50cbc0bb42, __obf_3fe025c030f856e9)
	}
}

// writePump sends ping to the client
func (__obf_c301d4a0b28bafec *Server) __obf_d39f634f3f7b9af8(__obf_45329568b5aa0723 *websocket.Conn, __obf_1a3b2d50cbc0bb42 ClientID, __obf_01602bf362dbf019 <-chan struct{}) {
	// create ping ticker
	__obf_ea1ef216c0e402f8 := time.NewTicker(__obf_c301d4a0b28bafec.PingPeriod * time.Second)
	defer __obf_ea1ef216c0e402f8.Stop()

	for {
		select {
		case <-__obf_ea1ef216c0e402f8.C:
			// send ping message
			__obf_4c0a65e7eb72a3b4 := __obf_45329568b5aa0723.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(__obf_c301d4a0b28bafec.WriteWait*time.Second))
			// log.Println("-----------------send ping message： ", err)
			if __obf_4c0a65e7eb72a3b4 != nil {
				// log.Println("-----------------send ping err: ", err)
				// if error sending ping, remove this client from the server
				__obf_c301d4a0b28bafec.RemoveClient(__obf_1a3b2d50cbc0bb42)
				// stop sending ping
				return
			}
		case <-__obf_01602bf362dbf019:
			// if process is done, stop sending ping
			return
		}
	}
}
