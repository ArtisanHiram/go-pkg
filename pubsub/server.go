package __obf_59c9fb1d2f9088b5

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
	__obf_39835c703952f3c2 = "publish"
	__obf_731dcdb0a4a1c994 = "subscribe"
	__obf_328bf98e743fa9d4 = "unsubscribe"
	// to client
	// data  = "data"
	__obf_a1d10a9366d4b3ee = "kickout"
	__obf_c5420dcf22964d23 = "throw"
	__obf_a06b7b3666cf424f = "greet"
	// popup = "popup"
)

// constants for server message
const (
	ErrInvalidMessage       = "Server: Invalid msg"
	ErrActionUnrecognizable = "Server: Action unrecognized"
)

type Topic string

// subscriber is a type for each string of topic and the clients that subscribe to it
type __obf_2860800d947ae7bf map[Topic]Client

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

type Handle func(ID ClientID, __obf_7e0d9092cb776df8 Topic, __obf_4146a9eb2015e1b3 string)

// Server is the struct to handle the Server functions & manage the subscriber
type Server struct {
	Option
	__obf_643a4faf1c7c6cf4 *sync.WaitGroup
	__obf_2c84e3b2883b0d88 *sync.RWMutex
	__obf_4f41bbfca5df2450 *sync.RWMutex
	__obf_2860800d947ae7bf __obf_2860800d947ae7bf
	__obf_55c3bc2ebc749fce map[string]Handle
}

func NewPubsub(__obf_a28d1bcb5d76b500 Option) *Server {
	return &Server{
		Option:                 __obf_a28d1bcb5d76b500,
		__obf_643a4faf1c7c6cf4: &sync.WaitGroup{},
		__obf_2c84e3b2883b0d88: &sync.RWMutex{},
		__obf_4f41bbfca5df2450: &sync.RWMutex{},
		__obf_2860800d947ae7bf: make(__obf_2860800d947ae7bf),
	}
}

func (__obf_503cab0b4b53bc0f *Server) SetActionHandle(__obf_6bc72bb5dbe12311 string, __obf_bc815840d5f9bf0a Handle) {

	if __obf_503cab0b4b53bc0f.__obf_55c3bc2ebc749fce == nil {
		__obf_503cab0b4b53bc0f.__obf_55c3bc2ebc749fce = map[string]Handle{
			__obf_6bc72bb5dbe12311: __obf_bc815840d5f9bf0a,
		}
	} else {
		__obf_503cab0b4b53bc0f.__obf_55c3bc2ebc749fce[__obf_6bc72bb5dbe12311] = __obf_bc815840d5f9bf0a
	}
}

func (__obf_503cab0b4b53bc0f *Server) Onlines(__obf_c3c1c637aafbea4d Topic) int {
	return len(__obf_503cab0b4b53bc0f.__obf_2860800d947ae7bf[__obf_c3c1c637aafbea4d])
}

// Send simply sends message to the websocket client
func (__obf_503cab0b4b53bc0f *Server) __obf_8661bf645a560097(__obf_fd69d0f6ad4bcc50 *websocket.Conn, __obf_6bc72bb5dbe12311 string, __obf_4146a9eb2015e1b3 any) {
	// send simple message

	// conn.WriteJSON(Message{Action: action, Body: msg})

	// bs, _ := tool.AnyToBytes(Message{Action: action, Body: msg})
	// conn.WriteMessage(websocket.BinaryMessage, bs)
	var __obf_649015e3057954e2 []byte
	__obf_649015e3057954e2, _ = util.Encode(Message{Action: __obf_6bc72bb5dbe12311, Body: __obf_4146a9eb2015e1b3})

	__obf_503cab0b4b53bc0f.__obf_2c84e3b2883b0d88.Lock()
	__obf_fd69d0f6ad4bcc50.WriteMessage(websocket.TextMessage, __obf_649015e3057954e2)
	__obf_503cab0b4b53bc0f.__obf_2c84e3b2883b0d88.Unlock()
}

func (__obf_503cab0b4b53bc0f *Server) GetConn(__obf_c3c1c637aafbea4d Topic, __obf_9e9c36dea22a3093 ClientID) *websocket.Conn {
	__obf_a3442935ee9ce758, __obf_4f157bf85f6e8f3b := __obf_503cab0b4b53bc0f.__obf_2860800d947ae7bf[__obf_c3c1c637aafbea4d]
	if !__obf_4f157bf85f6e8f3b {
		return nil
	}

	var __obf_fd69d0f6ad4bcc50 *websocket.Conn
	if __obf_fd69d0f6ad4bcc50, __obf_4f157bf85f6e8f3b = __obf_a3442935ee9ce758[__obf_9e9c36dea22a3093]; !__obf_4f157bf85f6e8f3b {
		return nil
	}
	return __obf_fd69d0f6ad4bcc50
}

func (__obf_503cab0b4b53bc0f *Server) SendTo(__obf_c3c1c637aafbea4d Topic, __obf_9e9c36dea22a3093 ClientID, __obf_6bc72bb5dbe12311 string, __obf_4146a9eb2015e1b3 any) {
	__obf_fd69d0f6ad4bcc50 := __obf_503cab0b4b53bc0f.GetConn(__obf_c3c1c637aafbea4d, __obf_9e9c36dea22a3093)
	if __obf_fd69d0f6ad4bcc50 != nil {
		__obf_503cab0b4b53bc0f.__obf_8661bf645a560097(__obf_fd69d0f6ad4bcc50, __obf_6bc72bb5dbe12311, __obf_4146a9eb2015e1b3)
	}
}

// SendWithWait sends message to the websocket client using wait group, allowing usage with goroutines
func (__obf_503cab0b4b53bc0f *Server) SendWithWait(__obf_fd69d0f6ad4bcc50 *websocket.Conn, __obf_6bc72bb5dbe12311 string, __obf_4146a9eb2015e1b3 any) {
	// send simple message
	// conn.WriteMessage(websocket.TextMessage, []byte(tool.EncodeString(msg)))
	__obf_503cab0b4b53bc0f.__obf_8661bf645a560097(__obf_fd69d0f6ad4bcc50, __obf_6bc72bb5dbe12311, __obf_4146a9eb2015e1b3)

	// set the task as done
	__obf_503cab0b4b53bc0f.__obf_643a4faf1c7c6cf4.Done()
}

func (__obf_503cab0b4b53bc0f *Server) SendGreet(__obf_fd69d0f6ad4bcc50 *websocket.Conn, __obf_4146a9eb2015e1b3 any) {
	__obf_503cab0b4b53bc0f.__obf_8661bf645a560097(__obf_fd69d0f6ad4bcc50, __obf_a06b7b3666cf424f, __obf_4146a9eb2015e1b3)
}

func (__obf_503cab0b4b53bc0f *Server) SendWithAction(__obf_fd69d0f6ad4bcc50 *websocket.Conn, __obf_6bc72bb5dbe12311 string, __obf_4146a9eb2015e1b3 any) {
	__obf_503cab0b4b53bc0f.__obf_8661bf645a560097(__obf_fd69d0f6ad4bcc50, __obf_6bc72bb5dbe12311, __obf_4146a9eb2015e1b3)
}

// func (s *Server) SendData(conn *websocket.Conn, msg any) {
// 	s.send(conn, data, msg)
// }

func (__obf_503cab0b4b53bc0f *Server) ThrowError(__obf_fd69d0f6ad4bcc50 *websocket.Conn, __obf_7ea45b848e69f2a5 string) {
	__obf_503cab0b4b53bc0f.__obf_8661bf645a560097(__obf_fd69d0f6ad4bcc50, __obf_c5420dcf22964d23, __obf_7ea45b848e69f2a5)
}

// RemoveClient removes the clients from the server subscription map
func (__obf_503cab0b4b53bc0f *Server) RemoveClient(__obf_9e9c36dea22a3093 ClientID) {
	__obf_503cab0b4b53bc0f.__obf_4f41bbfca5df2450.Lock()
	defer __obf_503cab0b4b53bc0f.__obf_4f41bbfca5df2450.Unlock()
	// loop all topics
	for __obf_5d60d438af93eb7b := range __obf_503cab0b4b53bc0f.__obf_2860800d947ae7bf {
		// delete the client from all the topic's client map
		delete(__obf_503cab0b4b53bc0f.__obf_2860800d947ae7bf[__obf_5d60d438af93eb7b], __obf_9e9c36dea22a3093)
	}
}

// ProcessMessage handle message according to the action type
func (__obf_503cab0b4b53bc0f *Server) ProcessMessage(__obf_fd69d0f6ad4bcc50 *websocket.Conn, __obf_9e9c36dea22a3093 ClientID, __obf_4146a9eb2015e1b3 []byte) {
	// parse message
	__obf_25ddb005fa6dbf39 := Message{}
	if __obf_7ea45b848e69f2a5 := json.Unmarshal(__obf_4146a9eb2015e1b3, &__obf_25ddb005fa6dbf39); __obf_7ea45b848e69f2a5 != nil {
		__obf_503cab0b4b53bc0f.ThrowError(__obf_fd69d0f6ad4bcc50, ErrInvalidMessage)
		return
	}

	// convert all action to lowercase and remove whitespace
	__obf_6bc72bb5dbe12311 := strings.TrimSpace(strings.ToLower(__obf_25ddb005fa6dbf39.Action))

	switch __obf_6bc72bb5dbe12311 {
	case __obf_39835c703952f3c2:
		__obf_503cab0b4b53bc0f.Publish(__obf_25ddb005fa6dbf39.Topic, __obf_25ddb005fa6dbf39.Body)

	case __obf_731dcdb0a4a1c994:
		__obf_503cab0b4b53bc0f.Subscribe(__obf_fd69d0f6ad4bcc50, __obf_9e9c36dea22a3093, __obf_25ddb005fa6dbf39.Topic)
	case __obf_328bf98e743fa9d4:
		__obf_503cab0b4b53bc0f.Unsubscribe(__obf_9e9c36dea22a3093, __obf_25ddb005fa6dbf39.Topic)

	default:
		if __obf_5b9f7913e7ed374d, __obf_bebf2bccd814e692 := __obf_503cab0b4b53bc0f.__obf_55c3bc2ebc749fce[__obf_6bc72bb5dbe12311]; __obf_bebf2bccd814e692 {
			__obf_5b9f7913e7ed374d(__obf_9e9c36dea22a3093, __obf_25ddb005fa6dbf39.Topic, __obf_25ddb005fa6dbf39.Body.(string))
		} else {
			__obf_503cab0b4b53bc0f.ThrowError(__obf_fd69d0f6ad4bcc50, ErrActionUnrecognizable)
		}

	}
}

// Publish sends a message to all subscribing clients of a topic
func (__obf_503cab0b4b53bc0f *Server) Publish(__obf_7e0d9092cb776df8 Topic, __obf_28250613366d0d9b any) {

	// if topic does not exist, stop the process
	if _, __obf_4f157bf85f6e8f3b := __obf_503cab0b4b53bc0f.__obf_2860800d947ae7bf[__obf_7e0d9092cb776df8]; !__obf_4f157bf85f6e8f3b {
		return
	}

	// if topic exist
	__obf_a3442935ee9ce758 := __obf_503cab0b4b53bc0f.__obf_2860800d947ae7bf[__obf_7e0d9092cb776df8]

	// send the message to the clients
	for _, __obf_fd69d0f6ad4bcc50 := range __obf_a3442935ee9ce758 {
		// add 1 job to wait group
		__obf_503cab0b4b53bc0f.__obf_643a4faf1c7c6cf4.Add(1)

		// send with goroutines
		go __obf_503cab0b4b53bc0f.SendWithWait(__obf_fd69d0f6ad4bcc50, __obf_39835c703952f3c2, __obf_28250613366d0d9b)
	}

	// wait until all goroutines jobs done
	__obf_503cab0b4b53bc0f.__obf_643a4faf1c7c6cf4.Wait()
}

// Kickout: kick the older session out
func (__obf_503cab0b4b53bc0f *Server) Kickout(__obf_c3c1c637aafbea4d Topic, __obf_9e9c36dea22a3093 ClientID) {
	__obf_fd69d0f6ad4bcc50 := __obf_503cab0b4b53bc0f.GetConn(__obf_c3c1c637aafbea4d, __obf_9e9c36dea22a3093)
	if __obf_fd69d0f6ad4bcc50 != nil {
		__obf_503cab0b4b53bc0f.__obf_8661bf645a560097(__obf_fd69d0f6ad4bcc50, __obf_a1d10a9366d4b3ee, nil)
		__obf_fd69d0f6ad4bcc50.Close()
		__obf_503cab0b4b53bc0f.__obf_4f41bbfca5df2450.Lock()
		delete(__obf_503cab0b4b53bc0f.__obf_2860800d947ae7bf[__obf_c3c1c637aafbea4d], __obf_9e9c36dea22a3093)
		__obf_503cab0b4b53bc0f.__obf_4f41bbfca5df2450.Unlock()
	}
}

// Subscribe adds a client to a topic's client map
func (__obf_503cab0b4b53bc0f *Server) Subscribe(__obf_fd69d0f6ad4bcc50 *websocket.Conn, __obf_9e9c36dea22a3093 ClientID, __obf_7e0d9092cb776df8 Topic) {
	__obf_503cab0b4b53bc0f.__obf_4f41bbfca5df2450.Lock()
	defer __obf_503cab0b4b53bc0f.__obf_4f41bbfca5df2450.Unlock()
	// if topic exist, check the client map
	if _, __obf_4f157bf85f6e8f3b := __obf_503cab0b4b53bc0f.__obf_2860800d947ae7bf[__obf_7e0d9092cb776df8]; __obf_4f157bf85f6e8f3b {
		__obf_a3442935ee9ce758 := __obf_503cab0b4b53bc0f.__obf_2860800d947ae7bf[__obf_7e0d9092cb776df8]

		// if subbed, ok := s.actionHandle[subscribe]; ok {
		// 	subbed(id, topic, "")
		// }

		__obf_a3442935ee9ce758[__obf_9e9c36dea22a3093] = __obf_fd69d0f6ad4bcc50
		return
	}

	// if topic does not exist, create a new topic
	__obf_55cac5838dbd68a6 := make(Client)
	__obf_503cab0b4b53bc0f.__obf_2860800d947ae7bf[__obf_7e0d9092cb776df8] = __obf_55cac5838dbd68a6

	// add the client to the topic
	__obf_503cab0b4b53bc0f.__obf_2860800d947ae7bf[__obf_7e0d9092cb776df8][__obf_9e9c36dea22a3093] = __obf_fd69d0f6ad4bcc50
}

// Unsubscribe removes a clients from a topic's client map
func (__obf_503cab0b4b53bc0f *Server) Unsubscribe(__obf_9e9c36dea22a3093 ClientID, __obf_7e0d9092cb776df8 Topic) {
	__obf_503cab0b4b53bc0f.__obf_4f41bbfca5df2450.Lock()
	defer __obf_503cab0b4b53bc0f.__obf_4f41bbfca5df2450.Unlock()
	// if topic exist, check the client map
	if _, __obf_4f157bf85f6e8f3b := __obf_503cab0b4b53bc0f.__obf_2860800d947ae7bf[__obf_7e0d9092cb776df8]; __obf_4f157bf85f6e8f3b {
		__obf_a3442935ee9ce758 := __obf_503cab0b4b53bc0f.__obf_2860800d947ae7bf[__obf_7e0d9092cb776df8]

		// remove the client from the topic's client map
		delete(__obf_a3442935ee9ce758, __obf_9e9c36dea22a3093)
	}
}

// Pump just do readDump and writePump
func (__obf_503cab0b4b53bc0f *Server) Pump(__obf_9e9c36dea22a3093 ClientID, __obf_fd69d0f6ad4bcc50 *websocket.Conn) {
	// create channel to signal client health
	__obf_3a68c97d6a98d0fd := make(chan struct{})

	go __obf_503cab0b4b53bc0f.__obf_37428b6a73be04dc(__obf_fd69d0f6ad4bcc50, __obf_9e9c36dea22a3093, __obf_3a68c97d6a98d0fd)
	__obf_503cab0b4b53bc0f.__obf_c38c02edfdc6014c(__obf_fd69d0f6ad4bcc50, __obf_9e9c36dea22a3093, __obf_3a68c97d6a98d0fd)
}

// readPump process incoming messages and set the settings
func (__obf_503cab0b4b53bc0f *Server) __obf_c38c02edfdc6014c(__obf_fd69d0f6ad4bcc50 *websocket.Conn, __obf_9e9c36dea22a3093 ClientID, __obf_3a68c97d6a98d0fd chan<- struct{}) {
	// set limit, deadline to read & pong handler
	__obf_fd69d0f6ad4bcc50.SetReadLimit(__obf_503cab0b4b53bc0f.MaxMessageSize)
	__obf_fd69d0f6ad4bcc50.SetReadDeadline(time.Now().Add(__obf_503cab0b4b53bc0f.PongWait * time.Second))
	__obf_fd69d0f6ad4bcc50.SetPongHandler(func(string) error {
		__obf_fd69d0f6ad4bcc50.SetReadDeadline(time.Now().Add(__obf_503cab0b4b53bc0f.PongWait * time.Second))
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
		_, __obf_4146a9eb2015e1b3, __obf_7ea45b848e69f2a5 := __obf_fd69d0f6ad4bcc50.ReadMessage()
		// if error occured
		if __obf_7ea45b848e69f2a5 != nil {
			// remove from the client
			__obf_503cab0b4b53bc0f.RemoveClient(__obf_9e9c36dea22a3093)
			// set health status to unhealthy by closing channel
			close(__obf_3a68c97d6a98d0fd)
			// stop process
			break
		}

		// if no error, process incoming message
		__obf_503cab0b4b53bc0f.ProcessMessage(__obf_fd69d0f6ad4bcc50, __obf_9e9c36dea22a3093, __obf_4146a9eb2015e1b3)
	}
}

// writePump sends ping to the client
func (__obf_503cab0b4b53bc0f *Server) __obf_37428b6a73be04dc(__obf_fd69d0f6ad4bcc50 *websocket.Conn, __obf_9e9c36dea22a3093 ClientID, __obf_3a68c97d6a98d0fd <-chan struct{}) {
	// create ping ticker
	__obf_150b94c66250d5d3 := time.NewTicker(__obf_503cab0b4b53bc0f.PingPeriod * time.Second)
	defer __obf_150b94c66250d5d3.Stop()

	for {
		select {
		case <-__obf_150b94c66250d5d3.C:
			// send ping message
			__obf_7ea45b848e69f2a5 := __obf_fd69d0f6ad4bcc50.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(__obf_503cab0b4b53bc0f.WriteWait*time.Second))
			// log.Println("-----------------send ping message： ", err)
			if __obf_7ea45b848e69f2a5 != nil {
				// log.Println("-----------------send ping err: ", err)
				// if error sending ping, remove this client from the server
				__obf_503cab0b4b53bc0f.RemoveClient(__obf_9e9c36dea22a3093)
				// stop sending ping
				return
			}
		case <-__obf_3a68c97d6a98d0fd:
			// if process is done, stop sending ping
			return
		}
	}
}
