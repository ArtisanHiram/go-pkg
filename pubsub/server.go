package __obf_930479f7379c770a

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
	__obf_6b3343ac1efbd510 =
	// from client
	"publish"
	__obf_11f0a682b2a81a36 = "subscribe"
	__obf_db38c5035579e958 = "unsubscribe"
	__obf_ecae868bcef0b31b =
	// to client
	// data  = "data"
	"kickout"
	__obf_4830b988b5756c69 = "throw"
	__obf_c158aa75b4504f5f = "greet"
	// popup = "popup"
)

// constants for server message
const (
	ErrInvalidMessage       = "Server: Invalid msg"
	ErrActionUnrecognizable = "Server: Action unrecognized"
)

type Topic string

// subscriber is a type for each string of topic and the clients that subscribe to it
type __obf_40cbc99485fb1aed map[Topic]Client

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

type Handle func(ID ClientID, __obf_3dbcb45a3588c1ad Topic, __obf_ab4da9561b4ba64a string)

// Server is the struct to handle the Server functions & manage the subscriber
type Server struct {
	Option
	__obf_d999cbb8597915e8 *sync.WaitGroup
	__obf_8664bb301fb4a57e *sync.RWMutex
	__obf_e481f145f7e675fd *sync.RWMutex
	__obf_40cbc99485fb1aed __obf_40cbc99485fb1aed
	__obf_c819988506de6608 map[string]Handle
}

func NewPubsub(__obf_e2f7dfc07bc6ed8c Option) *Server {
	return &Server{
		Option: __obf_e2f7dfc07bc6ed8c, __obf_d999cbb8597915e8: &sync.WaitGroup{}, __obf_8664bb301fb4a57e: &sync.RWMutex{}, __obf_e481f145f7e675fd: &sync.RWMutex{}, __obf_40cbc99485fb1aed: make(__obf_40cbc99485fb1aed),
	}
}

func (__obf_46746c95581844c1 *Server) SetActionHandle(__obf_91352fa03244854d string, __obf_804125c2c626a4e9 Handle) {

	if __obf_46746c95581844c1.__obf_c819988506de6608 == nil {
		__obf_46746c95581844c1.__obf_c819988506de6608 = map[string]Handle{__obf_91352fa03244854d: __obf_804125c2c626a4e9}
	} else {
		__obf_46746c95581844c1.__obf_c819988506de6608[__obf_91352fa03244854d] = __obf_804125c2c626a4e9
	}
}

func (__obf_46746c95581844c1 *Server) Onlines(__obf_e4bbceb437bd7bf3 Topic) int {
	return len(__obf_46746c95581844c1.

		// Send simply sends message to the websocket client
		__obf_40cbc99485fb1aed[__obf_e4bbceb437bd7bf3])
}

func (__obf_46746c95581844c1 *Server) __obf_7881a29d5cd7969d(__obf_a71c54218f46a429 *websocket.Conn, __obf_91352fa03244854d string, __obf_ab4da9561b4ba64a any) {
	// send simple message

	// conn.WriteJSON(Message{Action: action, Body: msg})

	// bs, _ := tool.AnyToBytes(Message{Action: action, Body: msg})
	// conn.WriteMessage(websocket.BinaryMessage, bs)
	var __obf_ee5e47a3dfde1a46 []byte
	__obf_ee5e47a3dfde1a46, _ = util.Encode(Message{Action: __obf_91352fa03244854d, Body: __obf_ab4da9561b4ba64a})
	__obf_46746c95581844c1.__obf_8664bb301fb4a57e.
		Lock()
	__obf_a71c54218f46a429.
		WriteMessage(websocket.TextMessage, __obf_ee5e47a3dfde1a46)
	__obf_46746c95581844c1.__obf_8664bb301fb4a57e.
		Unlock()
}

func (__obf_46746c95581844c1 *Server) GetConn(__obf_e4bbceb437bd7bf3 Topic, __obf_d33c619bfb2a5966 ClientID) *websocket.Conn {
	__obf_6af57ade83d5e2ab, __obf_70fa623410319c33 := __obf_46746c95581844c1.__obf_40cbc99485fb1aed[__obf_e4bbceb437bd7bf3]
	if !__obf_70fa623410319c33 {
		return nil
	}

	var __obf_a71c54218f46a429 *websocket.Conn
	if __obf_a71c54218f46a429, __obf_70fa623410319c33 = __obf_6af57ade83d5e2ab[__obf_d33c619bfb2a5966]; !__obf_70fa623410319c33 {
		return nil
	}
	return __obf_a71c54218f46a429
}

func (__obf_46746c95581844c1 *Server) SendTo(__obf_e4bbceb437bd7bf3 Topic, __obf_d33c619bfb2a5966 ClientID, __obf_91352fa03244854d string, __obf_ab4da9561b4ba64a any) {
	__obf_a71c54218f46a429 := __obf_46746c95581844c1.GetConn(__obf_e4bbceb437bd7bf3, __obf_d33c619bfb2a5966)
	if __obf_a71c54218f46a429 != nil {
		__obf_46746c95581844c1.__obf_7881a29d5cd7969d(__obf_a71c54218f46a429, __obf_91352fa03244854d,

			// SendWithWait sends message to the websocket client using wait group, allowing usage with goroutines
			__obf_ab4da9561b4ba64a)
	}
}

func (__obf_46746c95581844c1 *Server) SendWithWait(__obf_a71c54218f46a429 *websocket.Conn, __obf_91352fa03244854d string, __obf_ab4da9561b4ba64a any) {
	__obf_46746c95581844c1.
		// send simple message
		// conn.WriteMessage(websocket.TextMessage, []byte(tool.EncodeString(msg)))
		__obf_7881a29d5cd7969d(__obf_a71c54218f46a429,

			// set the task as done
			__obf_91352fa03244854d, __obf_ab4da9561b4ba64a)
	__obf_46746c95581844c1.__obf_d999cbb8597915e8.
		Done()
}

func (__obf_46746c95581844c1 *Server) SendGreet(__obf_a71c54218f46a429 *websocket.Conn, __obf_ab4da9561b4ba64a any) {
	__obf_46746c95581844c1.__obf_7881a29d5cd7969d(__obf_a71c54218f46a429, __obf_c158aa75b4504f5f, __obf_ab4da9561b4ba64a)
}

func (__obf_46746c95581844c1 *Server) SendWithAction(__obf_a71c54218f46a429 *websocket.Conn, __obf_91352fa03244854d string, __obf_ab4da9561b4ba64a any) {
	__obf_46746c95581844c1.__obf_7881a29d5cd7969d(__obf_a71c54218f46a429,

		// func (s *Server) SendData(conn *websocket.Conn, msg any) {
		// 	s.send(conn, data, msg)
		// }
		__obf_91352fa03244854d, __obf_ab4da9561b4ba64a)
}

func (__obf_46746c95581844c1 *Server) ThrowError(__obf_a71c54218f46a429 *websocket.Conn, __obf_bdaff79c89ee94b7 string) {
	__obf_46746c95581844c1.__obf_7881a29d5cd7969d(__obf_a71c54218f46a429,

		// RemoveClient removes the clients from the server subscription map
		__obf_4830b988b5756c69, __obf_bdaff79c89ee94b7)
}

func (__obf_46746c95581844c1 *Server) RemoveClient(__obf_d33c619bfb2a5966 ClientID) {
	__obf_46746c95581844c1.__obf_e481f145f7e675fd.
		Lock()
	defer __obf_46746c95581844c1.
		// loop all topics
		__obf_e481f145f7e675fd.Unlock()

	for __obf_30c7fe172bf550e4 := range __obf_46746c95581844c1.
		// delete the client from all the topic's client map
		__obf_40cbc99485fb1aed {

		delete(__obf_46746c95581844c1.__obf_40cbc99485fb1aed[__obf_30c7fe172bf550e4], __obf_d33c619bfb2a5966)
	}
}

// ProcessMessage handle message according to the action type
func (__obf_46746c95581844c1 *Server) ProcessMessage(__obf_a71c54218f46a429 *websocket.Conn, __obf_d33c619bfb2a5966 ClientID, __obf_ab4da9561b4ba64a []byte) {
	__obf_2bbe5c58bc758a66 := // parse message
		Message{}
	if __obf_bdaff79c89ee94b7 := json.Unmarshal(__obf_ab4da9561b4ba64a, &__obf_2bbe5c58bc758a66); __obf_bdaff79c89ee94b7 != nil {
		__obf_46746c95581844c1.
			ThrowError(__obf_a71c54218f46a429, ErrInvalidMessage)
		return
	}
	__obf_91352fa03244854d := // convert all action to lowercase and remove whitespace
		strings.TrimSpace(strings.ToLower(__obf_2bbe5c58bc758a66.Action))

	switch __obf_91352fa03244854d {
	case __obf_6b3343ac1efbd510:
		__obf_46746c95581844c1.
			Publish(__obf_2bbe5c58bc758a66.Topic, __obf_2bbe5c58bc758a66.Body)

	case __obf_11f0a682b2a81a36:
		__obf_46746c95581844c1.
			Subscribe(__obf_a71c54218f46a429, __obf_d33c619bfb2a5966, __obf_2bbe5c58bc758a66.Topic)
	case __obf_db38c5035579e958:
		__obf_46746c95581844c1.
			Unsubscribe(__obf_d33c619bfb2a5966, __obf_2bbe5c58bc758a66.Topic)

	default:
		if __obf_280f249c768ccd2b, __obf_74480d8d2970c6a2 := __obf_46746c95581844c1.__obf_c819988506de6608[__obf_91352fa03244854d]; __obf_74480d8d2970c6a2 {
			__obf_280f249c768ccd2b(__obf_d33c619bfb2a5966, __obf_2bbe5c58bc758a66.Topic, __obf_2bbe5c58bc758a66.Body.(string))
		} else {
			__obf_46746c95581844c1.
				ThrowError(__obf_a71c54218f46a429, ErrActionUnrecognizable)
		}

	}
}

// Publish sends a message to all subscribing clients of a topic
func (__obf_46746c95581844c1 *Server) Publish(__obf_3dbcb45a3588c1ad Topic, __obf_c73341db5271c21e any) {

	// if topic does not exist, stop the process
	if _, __obf_70fa623410319c33 := __obf_46746c95581844c1.__obf_40cbc99485fb1aed[__obf_3dbcb45a3588c1ad]; !__obf_70fa623410319c33 {
		return
	}
	__obf_6af57ade83d5e2ab := // if topic exist
		__obf_46746c95581844c1.__obf_40cbc99485fb1aed[__obf_3dbcb45a3588c1ad]

	// send the message to the clients
	for _, __obf_a71c54218f46a429 := range __obf_6af57ade83d5e2ab {
		__obf_46746c95581844c1.
			// add 1 job to wait group
			__obf_d999cbb8597915e8.
			Add(1)

		// send with goroutines
		go __obf_46746c95581844c1.SendWithWait(__obf_a71c54218f46a429, __obf_6b3343ac1efbd510,

			// wait until all goroutines jobs done
			__obf_c73341db5271c21e)
	}
	__obf_46746c95581844c1.__obf_d999cbb8597915e8.
		Wait()
}

// Kickout: kick the older session out
func (__obf_46746c95581844c1 *Server) Kickout(__obf_e4bbceb437bd7bf3 Topic, __obf_d33c619bfb2a5966 ClientID) {
	__obf_a71c54218f46a429 := __obf_46746c95581844c1.GetConn(__obf_e4bbceb437bd7bf3, __obf_d33c619bfb2a5966)
	if __obf_a71c54218f46a429 != nil {
		__obf_46746c95581844c1.__obf_7881a29d5cd7969d(__obf_a71c54218f46a429, __obf_ecae868bcef0b31b, nil)
		__obf_a71c54218f46a429.
			Close()
		__obf_46746c95581844c1.__obf_e481f145f7e675fd.
			Lock()
		delete(__obf_46746c95581844c1.__obf_40cbc99485fb1aed[__obf_e4bbceb437bd7bf3], __obf_d33c619bfb2a5966)
		__obf_46746c95581844c1.__obf_e481f145f7e675fd.
			Unlock()
	}
}

// Subscribe adds a client to a topic's client map
func (__obf_46746c95581844c1 *Server) Subscribe(__obf_a71c54218f46a429 *websocket.Conn, __obf_d33c619bfb2a5966 ClientID, __obf_3dbcb45a3588c1ad Topic) {
	__obf_46746c95581844c1.__obf_e481f145f7e675fd.
		Lock()
	defer __obf_46746c95581844c1.
		// if topic exist, check the client map
		__obf_e481f145f7e675fd.Unlock()

	if _, __obf_70fa623410319c33 := __obf_46746c95581844c1.__obf_40cbc99485fb1aed[__obf_3dbcb45a3588c1ad]; __obf_70fa623410319c33 {
		__obf_6af57ade83d5e2ab := __obf_46746c95581844c1.__obf_40cbc99485fb1aed[__obf_3dbcb45a3588c1ad]
		__obf_6af57ade83d5e2ab[ // if subbed, ok := s.actionHandle[subscribe]; ok {
		// 	subbed(id, topic, "")
		// }

		__obf_d33c619bfb2a5966] = __obf_a71c54218f46a429
		return
	}
	__obf_1a13c001bb12b580 := // if topic does not exist, create a new topic
		make(Client)
	__obf_46746c95581844c1.__obf_40cbc99485fb1aed[__obf_3dbcb45a3588c1ad] = __obf_1a13c001bb12b580

	// add the client to the topic
	__obf_46746c95581844c1.__obf_40cbc99485fb1aed[__obf_3dbcb45a3588c1ad][__obf_d33c619bfb2a5966] = __obf_a71c54218f46a429
}

// Unsubscribe removes a clients from a topic's client map
func (__obf_46746c95581844c1 *Server) Unsubscribe(__obf_d33c619bfb2a5966 ClientID, __obf_3dbcb45a3588c1ad Topic) {
	__obf_46746c95581844c1.__obf_e481f145f7e675fd.
		Lock()
	defer __obf_46746c95581844c1.
		// if topic exist, check the client map
		__obf_e481f145f7e675fd.Unlock()

	if _, __obf_70fa623410319c33 := __obf_46746c95581844c1.__obf_40cbc99485fb1aed[__obf_3dbcb45a3588c1ad]; __obf_70fa623410319c33 {
		__obf_6af57ade83d5e2ab := __obf_46746c95581844c1.__obf_40cbc99485fb1aed[__obf_3dbcb45a3588c1ad]

		// remove the client from the topic's client map
		delete(__obf_6af57ade83d5e2ab,

			// Pump just do readDump and writePump
			__obf_d33c619bfb2a5966)
	}
}

func (__obf_46746c95581844c1 *Server) Pump(__obf_d33c619bfb2a5966 ClientID, __obf_a71c54218f46a429 *websocket.Conn) {
	__obf_f65def8805e02050 := // create channel to signal client health
		make(chan struct{})

	go __obf_46746c95581844c1.__obf_8f5ae69eda59e60f(__obf_a71c54218f46a429, __obf_d33c619bfb2a5966, __obf_f65def8805e02050)
	__obf_46746c95581844c1.__obf_ff9d0012d002f917(__obf_a71c54218f46a429,

		// readPump process incoming messages and set the settings
		__obf_d33c619bfb2a5966, __obf_f65def8805e02050)
}

func (__obf_46746c95581844c1 *Server) __obf_ff9d0012d002f917(__obf_a71c54218f46a429 *websocket.Conn, __obf_d33c619bfb2a5966 ClientID, __obf_f65def8805e02050 chan<- struct{}) {
	__obf_a71c54218f46a429.
		// set limit, deadline to read & pong handler
		SetReadLimit(__obf_46746c95581844c1.MaxMessageSize)
	__obf_a71c54218f46a429.
		SetReadDeadline(time.Now().Add(__obf_46746c95581844c1.PongWait * time.Second))
	__obf_a71c54218f46a429.
		SetPongHandler(func(string) error {
			__obf_a71c54218f46a429.
				SetReadDeadline(time.Now().Add(__obf_46746c95581844c1.PongWait * time.Second))
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
		_, __obf_ab4da9561b4ba64a, __obf_bdaff79c89ee94b7 := __obf_a71c54218f46a429.ReadMessage()
		// if error occured
		if __obf_bdaff79c89ee94b7 != nil {
			__obf_46746c95581844c1.
				// remove from the client
				RemoveClient(__obf_d33c619bfb2a5966)
			// set health status to unhealthy by closing channel
			close(__obf_f65def8805e02050)
			// stop process
			break
		}
		__obf_46746c95581844c1.

			// if no error, process incoming message
			ProcessMessage(__obf_a71c54218f46a429,

				// writePump sends ping to the client
				__obf_d33c619bfb2a5966, __obf_ab4da9561b4ba64a)
	}
}

func (__obf_46746c95581844c1 *Server) __obf_8f5ae69eda59e60f(__obf_a71c54218f46a429 *websocket.Conn, __obf_d33c619bfb2a5966 ClientID, __obf_f65def8805e02050 <-chan struct{}) {
	__obf_9ec6294e945a8089 := // create ping ticker
		time.NewTicker(__obf_46746c95581844c1.PingPeriod * time.Second)
	defer __obf_9ec6294e945a8089.Stop()

	for {
		select {
		case <-__obf_9ec6294e945a8089.C:
			__obf_bdaff79c89ee94b7 := // send ping message
				__obf_a71c54218f46a429.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(__obf_46746c95581844c1.WriteWait*time.Second))
			// log.Println("-----------------send ping message： ", err)
			if __obf_bdaff79c89ee94b7 != nil {
				__obf_46746c95581844c1.
					// log.Println("-----------------send ping err: ", err)
					// if error sending ping, remove this client from the server
					RemoveClient(__obf_d33c619bfb2a5966)
				// stop sending ping
				return
			}
		case <-__obf_f65def8805e02050:
			// if process is done, stop sending ping
			return
		}
	}
}
