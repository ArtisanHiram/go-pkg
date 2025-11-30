package __obf_ec34ec234892f363

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
	__obf_e6179b909c60bce8 =
	// from client
	"publish"
	__obf_159b5dcb522f7573 = "subscribe"
	__obf_2b95785b545c80d5 = "unsubscribe"
	__obf_aa8bc8c18218d461 =
	// to client
	// data  = "data"
	"kickout"
	__obf_96943beaf59fdc5d = "throw"
	__obf_df61d2af207703a3 = "greet"
	// popup = "popup"
)

// constants for server message
const (
	ErrInvalidMessage       = "Server: Invalid msg"
	ErrActionUnrecognizable = "Server: Action unrecognized"
)

type Topic string

// subscriber is a type for each string of topic and the clients that subscribe to it
type __obf_2f0ce967d60c62ed map[Topic]Client

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

type Handle func(ID ClientID, __obf_8ee2ce372a7c506e Topic, __obf_953ef16555a7c447 string)

// Server is the struct to handle the Server functions & manage the subscriber
type Server struct {
	Option
	__obf_4083d61495f3f7fd *sync.WaitGroup
	__obf_f312721e9f08b589 *sync.RWMutex
	__obf_0bd0b09e8ecd4d8c *sync.RWMutex
	__obf_2f0ce967d60c62ed __obf_2f0ce967d60c62ed
	__obf_4dfb522986174228 map[string]Handle
}

func NewPubsub(__obf_415ff2851020bb70 Option) *Server {
	return &Server{
		Option: __obf_415ff2851020bb70, __obf_4083d61495f3f7fd: &sync.WaitGroup{}, __obf_f312721e9f08b589: &sync.RWMutex{}, __obf_0bd0b09e8ecd4d8c: &sync.RWMutex{}, __obf_2f0ce967d60c62ed: make(__obf_2f0ce967d60c62ed),
	}
}

func (__obf_e931fa0ee5f98570 *Server) SetActionHandle(__obf_74ca9ef814199c97 string, __obf_9931cd609cf732c9 Handle) {

	if __obf_e931fa0ee5f98570.__obf_4dfb522986174228 == nil {
		__obf_e931fa0ee5f98570.__obf_4dfb522986174228 = map[string]Handle{__obf_74ca9ef814199c97: __obf_9931cd609cf732c9}
	} else {
		__obf_e931fa0ee5f98570.__obf_4dfb522986174228[__obf_74ca9ef814199c97] = __obf_9931cd609cf732c9
	}
}

func (__obf_e931fa0ee5f98570 *Server) Onlines(__obf_b444d6aa8e668645 Topic) int {
	return len(__obf_e931fa0ee5f98570.

		// Send simply sends message to the websocket client
		__obf_2f0ce967d60c62ed[__obf_b444d6aa8e668645])
}

func (__obf_e931fa0ee5f98570 *Server) __obf_e59249a5fec6f81f(__obf_22772a0856af48e0 *websocket.Conn, __obf_74ca9ef814199c97 string, __obf_953ef16555a7c447 any) {
	// send simple message

	// conn.WriteJSON(Message{Action: action, Body: msg})

	// bs, _ := tool.AnyToBytes(Message{Action: action, Body: msg})
	// conn.WriteMessage(websocket.BinaryMessage, bs)
	var __obf_a348897cf40b60d8 []byte
	__obf_a348897cf40b60d8, _ = util.Encode(Message{Action: __obf_74ca9ef814199c97, Body: __obf_953ef16555a7c447})
	__obf_e931fa0ee5f98570.__obf_f312721e9f08b589.
		Lock()
	__obf_22772a0856af48e0.
		WriteMessage(websocket.TextMessage, __obf_a348897cf40b60d8)
	__obf_e931fa0ee5f98570.__obf_f312721e9f08b589.
		Unlock()
}

func (__obf_e931fa0ee5f98570 *Server) GetConn(__obf_b444d6aa8e668645 Topic, __obf_c1c775dee4a05460 ClientID) *websocket.Conn {
	__obf_bfb155cd8686f9d7, __obf_40606f172bd3fa79 := __obf_e931fa0ee5f98570.__obf_2f0ce967d60c62ed[__obf_b444d6aa8e668645]
	if !__obf_40606f172bd3fa79 {
		return nil
	}

	var __obf_22772a0856af48e0 *websocket.Conn
	if __obf_22772a0856af48e0, __obf_40606f172bd3fa79 = __obf_bfb155cd8686f9d7[__obf_c1c775dee4a05460]; !__obf_40606f172bd3fa79 {
		return nil
	}
	return __obf_22772a0856af48e0
}

func (__obf_e931fa0ee5f98570 *Server) SendTo(__obf_b444d6aa8e668645 Topic, __obf_c1c775dee4a05460 ClientID, __obf_74ca9ef814199c97 string, __obf_953ef16555a7c447 any) {
	__obf_22772a0856af48e0 := __obf_e931fa0ee5f98570.GetConn(__obf_b444d6aa8e668645, __obf_c1c775dee4a05460)
	if __obf_22772a0856af48e0 != nil {
		__obf_e931fa0ee5f98570.__obf_e59249a5fec6f81f(__obf_22772a0856af48e0, __obf_74ca9ef814199c97,

			// SendWithWait sends message to the websocket client using wait group, allowing usage with goroutines
			__obf_953ef16555a7c447)
	}
}

func (__obf_e931fa0ee5f98570 *Server) SendWithWait(__obf_22772a0856af48e0 *websocket.Conn, __obf_74ca9ef814199c97 string, __obf_953ef16555a7c447 any) {
	__obf_e931fa0ee5f98570.
		// send simple message
		// conn.WriteMessage(websocket.TextMessage, []byte(tool.EncodeString(msg)))
		__obf_e59249a5fec6f81f(__obf_22772a0856af48e0,

			// set the task as done
			__obf_74ca9ef814199c97, __obf_953ef16555a7c447)
	__obf_e931fa0ee5f98570.__obf_4083d61495f3f7fd.
		Done()
}

func (__obf_e931fa0ee5f98570 *Server) SendGreet(__obf_22772a0856af48e0 *websocket.Conn, __obf_953ef16555a7c447 any) {
	__obf_e931fa0ee5f98570.__obf_e59249a5fec6f81f(__obf_22772a0856af48e0, __obf_df61d2af207703a3, __obf_953ef16555a7c447)
}

func (__obf_e931fa0ee5f98570 *Server) SendWithAction(__obf_22772a0856af48e0 *websocket.Conn, __obf_74ca9ef814199c97 string, __obf_953ef16555a7c447 any) {
	__obf_e931fa0ee5f98570.__obf_e59249a5fec6f81f(__obf_22772a0856af48e0,

		// func (s *Server) SendData(conn *websocket.Conn, msg any) {
		// 	s.send(conn, data, msg)
		// }
		__obf_74ca9ef814199c97, __obf_953ef16555a7c447)
}

func (__obf_e931fa0ee5f98570 *Server) ThrowError(__obf_22772a0856af48e0 *websocket.Conn, __obf_ac19f08e3bdd5ddb string) {
	__obf_e931fa0ee5f98570.__obf_e59249a5fec6f81f(__obf_22772a0856af48e0,

		// RemoveClient removes the clients from the server subscription map
		__obf_96943beaf59fdc5d, __obf_ac19f08e3bdd5ddb)
}

func (__obf_e931fa0ee5f98570 *Server) RemoveClient(__obf_c1c775dee4a05460 ClientID) {
	__obf_e931fa0ee5f98570.__obf_0bd0b09e8ecd4d8c.
		Lock()
	defer __obf_e931fa0ee5f98570.
		// loop all topics
		__obf_0bd0b09e8ecd4d8c.Unlock()

	for __obf_a638f5a703a910dc := range __obf_e931fa0ee5f98570.
		// delete the client from all the topic's client map
		__obf_2f0ce967d60c62ed {

		delete(__obf_e931fa0ee5f98570.__obf_2f0ce967d60c62ed[__obf_a638f5a703a910dc], __obf_c1c775dee4a05460)
	}
}

// ProcessMessage handle message according to the action type
func (__obf_e931fa0ee5f98570 *Server) ProcessMessage(__obf_22772a0856af48e0 *websocket.Conn, __obf_c1c775dee4a05460 ClientID, __obf_953ef16555a7c447 []byte) {
	__obf_84642b03f8384ae3 := // parse message
		Message{}
	if __obf_ac19f08e3bdd5ddb := json.Unmarshal(__obf_953ef16555a7c447, &__obf_84642b03f8384ae3); __obf_ac19f08e3bdd5ddb != nil {
		__obf_e931fa0ee5f98570.
			ThrowError(__obf_22772a0856af48e0, ErrInvalidMessage)
		return
	}
	__obf_74ca9ef814199c97 := // convert all action to lowercase and remove whitespace
		strings.TrimSpace(strings.ToLower(__obf_84642b03f8384ae3.Action))

	switch __obf_74ca9ef814199c97 {
	case __obf_e6179b909c60bce8:
		__obf_e931fa0ee5f98570.
			Publish(__obf_84642b03f8384ae3.Topic, __obf_84642b03f8384ae3.Body)

	case __obf_159b5dcb522f7573:
		__obf_e931fa0ee5f98570.
			Subscribe(__obf_22772a0856af48e0, __obf_c1c775dee4a05460, __obf_84642b03f8384ae3.Topic)
	case __obf_2b95785b545c80d5:
		__obf_e931fa0ee5f98570.
			Unsubscribe(__obf_c1c775dee4a05460, __obf_84642b03f8384ae3.Topic)

	default:
		if __obf_31916cb2b1ad8a7c, __obf_0aebd580ef870cf7 := __obf_e931fa0ee5f98570.__obf_4dfb522986174228[__obf_74ca9ef814199c97]; __obf_0aebd580ef870cf7 {
			__obf_31916cb2b1ad8a7c(__obf_c1c775dee4a05460, __obf_84642b03f8384ae3.Topic, __obf_84642b03f8384ae3.Body.(string))
		} else {
			__obf_e931fa0ee5f98570.
				ThrowError(__obf_22772a0856af48e0, ErrActionUnrecognizable)
		}

	}
}

// Publish sends a message to all subscribing clients of a topic
func (__obf_e931fa0ee5f98570 *Server) Publish(__obf_8ee2ce372a7c506e Topic, __obf_0a8d4a931632dbe3 any) {

	// if topic does not exist, stop the process
	if _, __obf_40606f172bd3fa79 := __obf_e931fa0ee5f98570.__obf_2f0ce967d60c62ed[__obf_8ee2ce372a7c506e]; !__obf_40606f172bd3fa79 {
		return
	}
	__obf_bfb155cd8686f9d7 := // if topic exist
		__obf_e931fa0ee5f98570.__obf_2f0ce967d60c62ed[__obf_8ee2ce372a7c506e]

	// send the message to the clients
	for _, __obf_22772a0856af48e0 := range __obf_bfb155cd8686f9d7 {
		__obf_e931fa0ee5f98570.
			// add 1 job to wait group
			__obf_4083d61495f3f7fd.
			Add(1)

		// send with goroutines
		go __obf_e931fa0ee5f98570.SendWithWait(__obf_22772a0856af48e0, __obf_e6179b909c60bce8,

			// wait until all goroutines jobs done
			__obf_0a8d4a931632dbe3)
	}
	__obf_e931fa0ee5f98570.__obf_4083d61495f3f7fd.
		Wait()
}

// Kickout: kick the older session out
func (__obf_e931fa0ee5f98570 *Server) Kickout(__obf_b444d6aa8e668645 Topic, __obf_c1c775dee4a05460 ClientID) {
	__obf_22772a0856af48e0 := __obf_e931fa0ee5f98570.GetConn(__obf_b444d6aa8e668645, __obf_c1c775dee4a05460)
	if __obf_22772a0856af48e0 != nil {
		__obf_e931fa0ee5f98570.__obf_e59249a5fec6f81f(__obf_22772a0856af48e0, __obf_aa8bc8c18218d461, nil)
		__obf_22772a0856af48e0.
			Close()
		__obf_e931fa0ee5f98570.__obf_0bd0b09e8ecd4d8c.
			Lock()
		delete(__obf_e931fa0ee5f98570.__obf_2f0ce967d60c62ed[__obf_b444d6aa8e668645], __obf_c1c775dee4a05460)
		__obf_e931fa0ee5f98570.__obf_0bd0b09e8ecd4d8c.
			Unlock()
	}
}

// Subscribe adds a client to a topic's client map
func (__obf_e931fa0ee5f98570 *Server) Subscribe(__obf_22772a0856af48e0 *websocket.Conn, __obf_c1c775dee4a05460 ClientID, __obf_8ee2ce372a7c506e Topic) {
	__obf_e931fa0ee5f98570.__obf_0bd0b09e8ecd4d8c.
		Lock()
	defer __obf_e931fa0ee5f98570.
		// if topic exist, check the client map
		__obf_0bd0b09e8ecd4d8c.Unlock()

	if _, __obf_40606f172bd3fa79 := __obf_e931fa0ee5f98570.__obf_2f0ce967d60c62ed[__obf_8ee2ce372a7c506e]; __obf_40606f172bd3fa79 {
		__obf_bfb155cd8686f9d7 := __obf_e931fa0ee5f98570.__obf_2f0ce967d60c62ed[__obf_8ee2ce372a7c506e]
		__obf_bfb155cd8686f9d7[ // if subbed, ok := s.actionHandle[subscribe]; ok {
		// 	subbed(id, topic, "")
		// }

		__obf_c1c775dee4a05460] = __obf_22772a0856af48e0
		return
	}
	__obf_30b145a2f36870d0 := // if topic does not exist, create a new topic
		make(Client)
	__obf_e931fa0ee5f98570.__obf_2f0ce967d60c62ed[__obf_8ee2ce372a7c506e] = __obf_30b145a2f36870d0

	// add the client to the topic
	__obf_e931fa0ee5f98570.__obf_2f0ce967d60c62ed[__obf_8ee2ce372a7c506e][__obf_c1c775dee4a05460] = __obf_22772a0856af48e0
}

// Unsubscribe removes a clients from a topic's client map
func (__obf_e931fa0ee5f98570 *Server) Unsubscribe(__obf_c1c775dee4a05460 ClientID, __obf_8ee2ce372a7c506e Topic) {
	__obf_e931fa0ee5f98570.__obf_0bd0b09e8ecd4d8c.
		Lock()
	defer __obf_e931fa0ee5f98570.
		// if topic exist, check the client map
		__obf_0bd0b09e8ecd4d8c.Unlock()

	if _, __obf_40606f172bd3fa79 := __obf_e931fa0ee5f98570.__obf_2f0ce967d60c62ed[__obf_8ee2ce372a7c506e]; __obf_40606f172bd3fa79 {
		__obf_bfb155cd8686f9d7 := __obf_e931fa0ee5f98570.__obf_2f0ce967d60c62ed[__obf_8ee2ce372a7c506e]

		// remove the client from the topic's client map
		delete(__obf_bfb155cd8686f9d7,

			// Pump just do readDump and writePump
			__obf_c1c775dee4a05460)
	}
}

func (__obf_e931fa0ee5f98570 *Server) Pump(__obf_c1c775dee4a05460 ClientID, __obf_22772a0856af48e0 *websocket.Conn) {
	__obf_47bb1929ca1b844d := // create channel to signal client health
		make(chan struct{})

	go __obf_e931fa0ee5f98570.__obf_2b376bb83f1aa820(__obf_22772a0856af48e0, __obf_c1c775dee4a05460, __obf_47bb1929ca1b844d)
	__obf_e931fa0ee5f98570.__obf_e6a5041040ead410(__obf_22772a0856af48e0,

		// readPump process incoming messages and set the settings
		__obf_c1c775dee4a05460, __obf_47bb1929ca1b844d)
}

func (__obf_e931fa0ee5f98570 *Server) __obf_e6a5041040ead410(__obf_22772a0856af48e0 *websocket.Conn, __obf_c1c775dee4a05460 ClientID, __obf_47bb1929ca1b844d chan<- struct{}) {
	__obf_22772a0856af48e0.
		// set limit, deadline to read & pong handler
		SetReadLimit(__obf_e931fa0ee5f98570.MaxMessageSize)
	__obf_22772a0856af48e0.
		SetReadDeadline(time.Now().Add(__obf_e931fa0ee5f98570.PongWait * time.Second))
	__obf_22772a0856af48e0.
		SetPongHandler(func(string) error {
			__obf_22772a0856af48e0.
				SetReadDeadline(time.Now().Add(__obf_e931fa0ee5f98570.PongWait * time.Second))
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
		_, __obf_953ef16555a7c447, __obf_ac19f08e3bdd5ddb := __obf_22772a0856af48e0.ReadMessage()
		// if error occured
		if __obf_ac19f08e3bdd5ddb != nil {
			__obf_e931fa0ee5f98570.
				// remove from the client
				RemoveClient(__obf_c1c775dee4a05460)
			// set health status to unhealthy by closing channel
			close(__obf_47bb1929ca1b844d)
			// stop process
			break
		}
		__obf_e931fa0ee5f98570.

			// if no error, process incoming message
			ProcessMessage(__obf_22772a0856af48e0,

				// writePump sends ping to the client
				__obf_c1c775dee4a05460, __obf_953ef16555a7c447)
	}
}

func (__obf_e931fa0ee5f98570 *Server) __obf_2b376bb83f1aa820(__obf_22772a0856af48e0 *websocket.Conn, __obf_c1c775dee4a05460 ClientID, __obf_47bb1929ca1b844d <-chan struct{}) {
	__obf_08e181ade1bcfc2b := // create ping ticker
		time.NewTicker(__obf_e931fa0ee5f98570.PingPeriod * time.Second)
	defer __obf_08e181ade1bcfc2b.Stop()

	for {
		select {
		case <-__obf_08e181ade1bcfc2b.C:
			__obf_ac19f08e3bdd5ddb := // send ping message
				__obf_22772a0856af48e0.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(__obf_e931fa0ee5f98570.WriteWait*time.Second))
			// log.Println("-----------------send ping message： ", err)
			if __obf_ac19f08e3bdd5ddb != nil {
				__obf_e931fa0ee5f98570.
					// log.Println("-----------------send ping err: ", err)
					// if error sending ping, remove this client from the server
					RemoveClient(__obf_c1c775dee4a05460)
				// stop sending ping
				return
			}
		case <-__obf_47bb1929ca1b844d:
			// if process is done, stop sending ping
			return
		}
	}
}
