package __obf_a5d9fffca83c4f6a

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
	__obf_2f3f009dc4fac642 =
	// from client
	"publish"
	__obf_d8c698cdc658ca07 = "subscribe"
	__obf_cc0e19ec4f3ed44d = "unsubscribe"
	__obf_f6f305c021ed890a =
	// to client
	// data  = "data"
	"kickout"
	__obf_5df85c1d11283e4c = "throw"
	__obf_a3edf515fa789509 = "greet"
	// popup = "popup"
)

// constants for server message
const (
	ErrInvalidMessage       = "Server: Invalid msg"
	ErrActionUnrecognizable = "Server: Action unrecognized"
)

type Topic string

// subscriber is a type for each string of topic and the clients that subscribe to it
type __obf_bcf6af72ce160e0b map[Topic]Client

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

type Handle func(ID ClientID, __obf_a9b0d551107d0e04 Topic, __obf_b158500b507af22f string)

// Server is the struct to handle the Server functions & manage the subscriber
type Server struct {
	Option
	__obf_f9d6f41eefa8c1a9 *sync.WaitGroup
	__obf_0408ee5d54fb8d15 *sync.RWMutex
	__obf_4e9f0e38c5ceae84 *sync.RWMutex
	__obf_bcf6af72ce160e0b __obf_bcf6af72ce160e0b
	__obf_9ceb92e4b771bdb3 map[string]Handle
}

func NewPubsub(__obf_e80de1f329e1fb59 Option) *Server {
	return &Server{
		Option: __obf_e80de1f329e1fb59, __obf_f9d6f41eefa8c1a9: &sync.WaitGroup{}, __obf_0408ee5d54fb8d15: &sync.RWMutex{}, __obf_4e9f0e38c5ceae84: &sync.RWMutex{}, __obf_bcf6af72ce160e0b: make(__obf_bcf6af72ce160e0b),
	}
}

func (__obf_cce4c87c6bca93ba *Server) SetActionHandle(__obf_707ded8bbcceafab string, __obf_56e82049f425aa97 Handle) {

	if __obf_cce4c87c6bca93ba.__obf_9ceb92e4b771bdb3 == nil {
		__obf_cce4c87c6bca93ba.__obf_9ceb92e4b771bdb3 = map[string]Handle{__obf_707ded8bbcceafab: __obf_56e82049f425aa97}
	} else {
		__obf_cce4c87c6bca93ba.__obf_9ceb92e4b771bdb3[__obf_707ded8bbcceafab] = __obf_56e82049f425aa97
	}
}

func (__obf_cce4c87c6bca93ba *Server) Onlines(__obf_d4dc6df1050c5bc5 Topic) int {
	return len(__obf_cce4c87c6bca93ba.

		// Send simply sends message to the websocket client
		__obf_bcf6af72ce160e0b[__obf_d4dc6df1050c5bc5])
}

func (__obf_cce4c87c6bca93ba *Server) __obf_345ae38d7326c139(__obf_bf6c77b07b73c2e3 *websocket.Conn, __obf_707ded8bbcceafab string, __obf_b158500b507af22f any) {
	// send simple message

	// conn.WriteJSON(Message{Action: action, Body: msg})

	// bs, _ := tool.AnyToBytes(Message{Action: action, Body: msg})
	// conn.WriteMessage(websocket.BinaryMessage, bs)
	var __obf_fef36ee98107ae41 []byte
	__obf_fef36ee98107ae41, _ = util.Encode(Message{Action: __obf_707ded8bbcceafab, Body: __obf_b158500b507af22f})
	__obf_cce4c87c6bca93ba.__obf_0408ee5d54fb8d15.
		Lock()
	__obf_bf6c77b07b73c2e3.
		WriteMessage(websocket.TextMessage, __obf_fef36ee98107ae41)
	__obf_cce4c87c6bca93ba.__obf_0408ee5d54fb8d15.
		Unlock()
}

func (__obf_cce4c87c6bca93ba *Server) GetConn(__obf_d4dc6df1050c5bc5 Topic, __obf_5977179c2154d968 ClientID) *websocket.Conn {
	__obf_40791e423f904de6, __obf_828ba6b4a9b68722 := __obf_cce4c87c6bca93ba.__obf_bcf6af72ce160e0b[__obf_d4dc6df1050c5bc5]
	if !__obf_828ba6b4a9b68722 {
		return nil
	}

	var __obf_bf6c77b07b73c2e3 *websocket.Conn
	if __obf_bf6c77b07b73c2e3, __obf_828ba6b4a9b68722 = __obf_40791e423f904de6[__obf_5977179c2154d968]; !__obf_828ba6b4a9b68722 {
		return nil
	}
	return __obf_bf6c77b07b73c2e3
}

func (__obf_cce4c87c6bca93ba *Server) SendTo(__obf_d4dc6df1050c5bc5 Topic, __obf_5977179c2154d968 ClientID, __obf_707ded8bbcceafab string, __obf_b158500b507af22f any) {
	__obf_bf6c77b07b73c2e3 := __obf_cce4c87c6bca93ba.GetConn(__obf_d4dc6df1050c5bc5, __obf_5977179c2154d968)
	if __obf_bf6c77b07b73c2e3 != nil {
		__obf_cce4c87c6bca93ba.__obf_345ae38d7326c139(__obf_bf6c77b07b73c2e3, __obf_707ded8bbcceafab,

			// SendWithWait sends message to the websocket client using wait group, allowing usage with goroutines
			__obf_b158500b507af22f)
	}
}

func (__obf_cce4c87c6bca93ba *Server) SendWithWait(__obf_bf6c77b07b73c2e3 *websocket.Conn, __obf_707ded8bbcceafab string, __obf_b158500b507af22f any) {
	__obf_cce4c87c6bca93ba.
		// send simple message
		// conn.WriteMessage(websocket.TextMessage, []byte(tool.EncodeString(msg)))
		__obf_345ae38d7326c139(__obf_bf6c77b07b73c2e3,

			// set the task as done
			__obf_707ded8bbcceafab, __obf_b158500b507af22f)
	__obf_cce4c87c6bca93ba.__obf_f9d6f41eefa8c1a9.
		Done()
}

func (__obf_cce4c87c6bca93ba *Server) SendGreet(__obf_bf6c77b07b73c2e3 *websocket.Conn, __obf_b158500b507af22f any) {
	__obf_cce4c87c6bca93ba.__obf_345ae38d7326c139(__obf_bf6c77b07b73c2e3, __obf_a3edf515fa789509, __obf_b158500b507af22f)
}

func (__obf_cce4c87c6bca93ba *Server) SendWithAction(__obf_bf6c77b07b73c2e3 *websocket.Conn, __obf_707ded8bbcceafab string, __obf_b158500b507af22f any) {
	__obf_cce4c87c6bca93ba.__obf_345ae38d7326c139(__obf_bf6c77b07b73c2e3,

		// func (s *Server) SendData(conn *websocket.Conn, msg any) {
		// 	s.send(conn, data, msg)
		// }
		__obf_707ded8bbcceafab, __obf_b158500b507af22f)
}

func (__obf_cce4c87c6bca93ba *Server) ThrowError(__obf_bf6c77b07b73c2e3 *websocket.Conn, __obf_24d0a00a0810aa56 string) {
	__obf_cce4c87c6bca93ba.__obf_345ae38d7326c139(__obf_bf6c77b07b73c2e3,

		// RemoveClient removes the clients from the server subscription map
		__obf_5df85c1d11283e4c, __obf_24d0a00a0810aa56)
}

func (__obf_cce4c87c6bca93ba *Server) RemoveClient(__obf_5977179c2154d968 ClientID) {
	__obf_cce4c87c6bca93ba.__obf_4e9f0e38c5ceae84.
		Lock()
	defer __obf_cce4c87c6bca93ba.
		// loop all topics
		__obf_4e9f0e38c5ceae84.Unlock()

	for __obf_057d2db022ba84c1 := range __obf_cce4c87c6bca93ba.
		// delete the client from all the topic's client map
		__obf_bcf6af72ce160e0b {

		delete(__obf_cce4c87c6bca93ba.__obf_bcf6af72ce160e0b[__obf_057d2db022ba84c1], __obf_5977179c2154d968)
	}
}

// ProcessMessage handle message according to the action type
func (__obf_cce4c87c6bca93ba *Server) ProcessMessage(__obf_bf6c77b07b73c2e3 *websocket.Conn, __obf_5977179c2154d968 ClientID, __obf_b158500b507af22f []byte) {
	__obf_cc794d5e862baeaf := // parse message
		Message{}
	if __obf_24d0a00a0810aa56 := json.Unmarshal(__obf_b158500b507af22f, &__obf_cc794d5e862baeaf); __obf_24d0a00a0810aa56 != nil {
		__obf_cce4c87c6bca93ba.
			ThrowError(__obf_bf6c77b07b73c2e3, ErrInvalidMessage)
		return
	}
	__obf_707ded8bbcceafab := // convert all action to lowercase and remove whitespace
		strings.TrimSpace(strings.ToLower(__obf_cc794d5e862baeaf.Action))

	switch __obf_707ded8bbcceafab {
	case __obf_2f3f009dc4fac642:
		__obf_cce4c87c6bca93ba.
			Publish(__obf_cc794d5e862baeaf.Topic, __obf_cc794d5e862baeaf.Body)

	case __obf_d8c698cdc658ca07:
		__obf_cce4c87c6bca93ba.
			Subscribe(__obf_bf6c77b07b73c2e3, __obf_5977179c2154d968, __obf_cc794d5e862baeaf.Topic)
	case __obf_cc0e19ec4f3ed44d:
		__obf_cce4c87c6bca93ba.
			Unsubscribe(__obf_5977179c2154d968, __obf_cc794d5e862baeaf.Topic)

	default:
		if __obf_27dd535af2d5ccda, __obf_5fc6f3ddc5863649 := __obf_cce4c87c6bca93ba.__obf_9ceb92e4b771bdb3[__obf_707ded8bbcceafab]; __obf_5fc6f3ddc5863649 {
			__obf_27dd535af2d5ccda(__obf_5977179c2154d968, __obf_cc794d5e862baeaf.Topic, __obf_cc794d5e862baeaf.Body.(string))
		} else {
			__obf_cce4c87c6bca93ba.
				ThrowError(__obf_bf6c77b07b73c2e3, ErrActionUnrecognizable)
		}

	}
}

// Publish sends a message to all subscribing clients of a topic
func (__obf_cce4c87c6bca93ba *Server) Publish(__obf_a9b0d551107d0e04 Topic, __obf_dc8848cc6ec179bf any) {

	// if topic does not exist, stop the process
	if _, __obf_828ba6b4a9b68722 := __obf_cce4c87c6bca93ba.__obf_bcf6af72ce160e0b[__obf_a9b0d551107d0e04]; !__obf_828ba6b4a9b68722 {
		return
	}
	__obf_40791e423f904de6 := // if topic exist
		__obf_cce4c87c6bca93ba.__obf_bcf6af72ce160e0b[__obf_a9b0d551107d0e04]

	// send the message to the clients
	for _, __obf_bf6c77b07b73c2e3 := range __obf_40791e423f904de6 {
		__obf_cce4c87c6bca93ba.
			// add 1 job to wait group
			__obf_f9d6f41eefa8c1a9.
			Add(1)

		// send with goroutines
		go __obf_cce4c87c6bca93ba.SendWithWait(__obf_bf6c77b07b73c2e3, __obf_2f3f009dc4fac642,

			// wait until all goroutines jobs done
			__obf_dc8848cc6ec179bf)
	}
	__obf_cce4c87c6bca93ba.__obf_f9d6f41eefa8c1a9.
		Wait()
}

// Kickout: kick the older session out
func (__obf_cce4c87c6bca93ba *Server) Kickout(__obf_d4dc6df1050c5bc5 Topic, __obf_5977179c2154d968 ClientID) {
	__obf_bf6c77b07b73c2e3 := __obf_cce4c87c6bca93ba.GetConn(__obf_d4dc6df1050c5bc5, __obf_5977179c2154d968)
	if __obf_bf6c77b07b73c2e3 != nil {
		__obf_cce4c87c6bca93ba.__obf_345ae38d7326c139(__obf_bf6c77b07b73c2e3, __obf_f6f305c021ed890a, nil)
		__obf_bf6c77b07b73c2e3.
			Close()
		__obf_cce4c87c6bca93ba.__obf_4e9f0e38c5ceae84.
			Lock()
		delete(__obf_cce4c87c6bca93ba.__obf_bcf6af72ce160e0b[__obf_d4dc6df1050c5bc5], __obf_5977179c2154d968)
		__obf_cce4c87c6bca93ba.__obf_4e9f0e38c5ceae84.
			Unlock()
	}
}

// Subscribe adds a client to a topic's client map
func (__obf_cce4c87c6bca93ba *Server) Subscribe(__obf_bf6c77b07b73c2e3 *websocket.Conn, __obf_5977179c2154d968 ClientID, __obf_a9b0d551107d0e04 Topic) {
	__obf_cce4c87c6bca93ba.__obf_4e9f0e38c5ceae84.
		Lock()
	defer __obf_cce4c87c6bca93ba.
		// if topic exist, check the client map
		__obf_4e9f0e38c5ceae84.Unlock()

	if _, __obf_828ba6b4a9b68722 := __obf_cce4c87c6bca93ba.__obf_bcf6af72ce160e0b[__obf_a9b0d551107d0e04]; __obf_828ba6b4a9b68722 {
		__obf_40791e423f904de6 := __obf_cce4c87c6bca93ba.__obf_bcf6af72ce160e0b[__obf_a9b0d551107d0e04]
		__obf_40791e423f904de6[ // if subbed, ok := s.actionHandle[subscribe]; ok {
		// 	subbed(id, topic, "")
		// }

		__obf_5977179c2154d968] = __obf_bf6c77b07b73c2e3
		return
	}
	__obf_e304d7f0de86d5a8 := // if topic does not exist, create a new topic
		make(Client)
	__obf_cce4c87c6bca93ba.__obf_bcf6af72ce160e0b[__obf_a9b0d551107d0e04] = __obf_e304d7f0de86d5a8

	// add the client to the topic
	__obf_cce4c87c6bca93ba.__obf_bcf6af72ce160e0b[__obf_a9b0d551107d0e04][__obf_5977179c2154d968] = __obf_bf6c77b07b73c2e3
}

// Unsubscribe removes a clients from a topic's client map
func (__obf_cce4c87c6bca93ba *Server) Unsubscribe(__obf_5977179c2154d968 ClientID, __obf_a9b0d551107d0e04 Topic) {
	__obf_cce4c87c6bca93ba.__obf_4e9f0e38c5ceae84.
		Lock()
	defer __obf_cce4c87c6bca93ba.
		// if topic exist, check the client map
		__obf_4e9f0e38c5ceae84.Unlock()

	if _, __obf_828ba6b4a9b68722 := __obf_cce4c87c6bca93ba.__obf_bcf6af72ce160e0b[__obf_a9b0d551107d0e04]; __obf_828ba6b4a9b68722 {
		__obf_40791e423f904de6 := __obf_cce4c87c6bca93ba.__obf_bcf6af72ce160e0b[__obf_a9b0d551107d0e04]

		// remove the client from the topic's client map
		delete(__obf_40791e423f904de6,

			// Pump just do readDump and writePump
			__obf_5977179c2154d968)
	}
}

func (__obf_cce4c87c6bca93ba *Server) Pump(__obf_5977179c2154d968 ClientID, __obf_bf6c77b07b73c2e3 *websocket.Conn) {
	__obf_fe603e5ba804932d := // create channel to signal client health
		make(chan struct{})

	go __obf_cce4c87c6bca93ba.__obf_1bd67c94399a7725(__obf_bf6c77b07b73c2e3, __obf_5977179c2154d968, __obf_fe603e5ba804932d)
	__obf_cce4c87c6bca93ba.__obf_f1fd3dc05013b489(__obf_bf6c77b07b73c2e3,

		// readPump process incoming messages and set the settings
		__obf_5977179c2154d968, __obf_fe603e5ba804932d)
}

func (__obf_cce4c87c6bca93ba *Server) __obf_f1fd3dc05013b489(__obf_bf6c77b07b73c2e3 *websocket.Conn, __obf_5977179c2154d968 ClientID, __obf_fe603e5ba804932d chan<- struct{}) {
	__obf_bf6c77b07b73c2e3.
		// set limit, deadline to read & pong handler
		SetReadLimit(__obf_cce4c87c6bca93ba.MaxMessageSize)
	__obf_bf6c77b07b73c2e3.
		SetReadDeadline(time.Now().Add(__obf_cce4c87c6bca93ba.PongWait * time.Second))
	__obf_bf6c77b07b73c2e3.
		SetPongHandler(func(string) error {
			__obf_bf6c77b07b73c2e3.
				SetReadDeadline(time.Now().Add(__obf_cce4c87c6bca93ba.PongWait * time.Second))
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
		_, __obf_b158500b507af22f, __obf_24d0a00a0810aa56 := __obf_bf6c77b07b73c2e3.ReadMessage()
		// if error occured
		if __obf_24d0a00a0810aa56 != nil {
			__obf_cce4c87c6bca93ba.
				// remove from the client
				RemoveClient(__obf_5977179c2154d968)
			// set health status to unhealthy by closing channel
			close(__obf_fe603e5ba804932d)
			// stop process
			break
		}
		__obf_cce4c87c6bca93ba.

			// if no error, process incoming message
			ProcessMessage(__obf_bf6c77b07b73c2e3,

				// writePump sends ping to the client
				__obf_5977179c2154d968, __obf_b158500b507af22f)
	}
}

func (__obf_cce4c87c6bca93ba *Server) __obf_1bd67c94399a7725(__obf_bf6c77b07b73c2e3 *websocket.Conn, __obf_5977179c2154d968 ClientID, __obf_fe603e5ba804932d <-chan struct{}) {
	__obf_a0790ce73acd2329 := // create ping ticker
		time.NewTicker(__obf_cce4c87c6bca93ba.PingPeriod * time.Second)
	defer __obf_a0790ce73acd2329.Stop()

	for {
		select {
		case <-__obf_a0790ce73acd2329.C:
			__obf_24d0a00a0810aa56 := // send ping message
				__obf_bf6c77b07b73c2e3.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(__obf_cce4c87c6bca93ba.WriteWait*time.Second))
			// log.Println("-----------------send ping message： ", err)
			if __obf_24d0a00a0810aa56 != nil {
				__obf_cce4c87c6bca93ba.
					// log.Println("-----------------send ping err: ", err)
					// if error sending ping, remove this client from the server
					RemoveClient(__obf_5977179c2154d968)
				// stop sending ping
				return
			}
		case <-__obf_fe603e5ba804932d:
			// if process is done, stop sending ping
			return
		}
	}
}
