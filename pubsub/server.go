package __obf_73650e5c67babe39

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
	__obf_4d69e603d2296d75 =
	// from client
	"publish"
	__obf_317fedd18f3baea6 = "subscribe"
	__obf_6109636d122c33ab = "unsubscribe"
	__obf_9dad881b1e5a2a94 =
	// to client
	// data  = "data"
	"kickout"
	__obf_548b437f9bea80cc = "throw"
	__obf_235c8a6ac56bbad0 = "greet"
	// popup = "popup"
)

// constants for server message
const (
	ErrInvalidMessage       = "Server: Invalid msg"
	ErrActionUnrecognizable = "Server: Action unrecognized"
)

type Topic string

// subscriber is a type for each string of topic and the clients that subscribe to it
type __obf_90a259ae07de44fa map[Topic]Client

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

type Handle func(ID ClientID, __obf_4fbd419e634c2688 Topic, __obf_0052e3c7488be106 string)

// Server is the struct to handle the Server functions & manage the subscriber
type Server struct {
	Option
	__obf_29d5838d5b731984 *sync.WaitGroup
	__obf_6d445d549963ea9f *sync.RWMutex
	__obf_8fe2ef1b19fe43ad *sync.RWMutex
	__obf_90a259ae07de44fa __obf_90a259ae07de44fa
	__obf_4046692401f1241a map[string]Handle
}

func NewPubsub(__obf_a0fc3a30918a97c1 Option) *Server {
	return &Server{
		Option: __obf_a0fc3a30918a97c1, __obf_29d5838d5b731984: &sync.WaitGroup{}, __obf_6d445d549963ea9f: &sync.RWMutex{}, __obf_8fe2ef1b19fe43ad: &sync.RWMutex{}, __obf_90a259ae07de44fa: make(__obf_90a259ae07de44fa),
	}
}

func (__obf_4f8025047382d160 *Server) SetActionHandle(__obf_e6e2b26604ba9b9b string, __obf_87f7ea037d06ac36 Handle) {

	if __obf_4f8025047382d160.__obf_4046692401f1241a == nil {
		__obf_4f8025047382d160.__obf_4046692401f1241a = map[string]Handle{__obf_e6e2b26604ba9b9b: __obf_87f7ea037d06ac36}
	} else {
		__obf_4f8025047382d160.__obf_4046692401f1241a[__obf_e6e2b26604ba9b9b] = __obf_87f7ea037d06ac36
	}
}

func (__obf_4f8025047382d160 *Server) Onlines(__obf_ea4fca7fd07b07f6 Topic) int {
	return len(__obf_4f8025047382d160.

		// Send simply sends message to the websocket client
		__obf_90a259ae07de44fa[__obf_ea4fca7fd07b07f6])
}

func (__obf_4f8025047382d160 *Server) __obf_5eabb95660fc0d91(__obf_6d1c7064e94466ae *websocket.Conn, __obf_e6e2b26604ba9b9b string, __obf_0052e3c7488be106 any) {
	// send simple message

	// conn.WriteJSON(Message{Action: action, Body: msg})

	// bs, _ := tool.AnyToBytes(Message{Action: action, Body: msg})
	// conn.WriteMessage(websocket.BinaryMessage, bs)
	var __obf_5e6d7b588a0eb4ae []byte
	__obf_5e6d7b588a0eb4ae, _ = util.Encode(Message{Action: __obf_e6e2b26604ba9b9b, Body: __obf_0052e3c7488be106})
	__obf_4f8025047382d160.__obf_6d445d549963ea9f.
		Lock()
	__obf_6d1c7064e94466ae.
		WriteMessage(websocket.TextMessage, __obf_5e6d7b588a0eb4ae)
	__obf_4f8025047382d160.__obf_6d445d549963ea9f.
		Unlock()
}

func (__obf_4f8025047382d160 *Server) GetConn(__obf_ea4fca7fd07b07f6 Topic, __obf_2cdc7127d0e23d64 ClientID) *websocket.Conn {
	__obf_162f952545cfa04e, __obf_4c9a79bce95923ad := __obf_4f8025047382d160.__obf_90a259ae07de44fa[__obf_ea4fca7fd07b07f6]
	if !__obf_4c9a79bce95923ad {
		return nil
	}

	var __obf_6d1c7064e94466ae *websocket.Conn
	if __obf_6d1c7064e94466ae, __obf_4c9a79bce95923ad = __obf_162f952545cfa04e[__obf_2cdc7127d0e23d64]; !__obf_4c9a79bce95923ad {
		return nil
	}
	return __obf_6d1c7064e94466ae
}

func (__obf_4f8025047382d160 *Server) SendTo(__obf_ea4fca7fd07b07f6 Topic, __obf_2cdc7127d0e23d64 ClientID, __obf_e6e2b26604ba9b9b string, __obf_0052e3c7488be106 any) {
	__obf_6d1c7064e94466ae := __obf_4f8025047382d160.GetConn(__obf_ea4fca7fd07b07f6, __obf_2cdc7127d0e23d64)
	if __obf_6d1c7064e94466ae != nil {
		__obf_4f8025047382d160.__obf_5eabb95660fc0d91(__obf_6d1c7064e94466ae, __obf_e6e2b26604ba9b9b,

			// SendWithWait sends message to the websocket client using wait group, allowing usage with goroutines
			__obf_0052e3c7488be106)
	}
}

func (__obf_4f8025047382d160 *Server) SendWithWait(__obf_6d1c7064e94466ae *websocket.Conn, __obf_e6e2b26604ba9b9b string, __obf_0052e3c7488be106 any) {
	__obf_4f8025047382d160.
		// send simple message
		// conn.WriteMessage(websocket.TextMessage, []byte(tool.EncodeString(msg)))
		__obf_5eabb95660fc0d91(__obf_6d1c7064e94466ae,

			// set the task as done
			__obf_e6e2b26604ba9b9b, __obf_0052e3c7488be106)
	__obf_4f8025047382d160.__obf_29d5838d5b731984.
		Done()
}

func (__obf_4f8025047382d160 *Server) SendGreet(__obf_6d1c7064e94466ae *websocket.Conn, __obf_0052e3c7488be106 any) {
	__obf_4f8025047382d160.__obf_5eabb95660fc0d91(__obf_6d1c7064e94466ae, __obf_235c8a6ac56bbad0, __obf_0052e3c7488be106)
}

func (__obf_4f8025047382d160 *Server) SendWithAction(__obf_6d1c7064e94466ae *websocket.Conn, __obf_e6e2b26604ba9b9b string, __obf_0052e3c7488be106 any) {
	__obf_4f8025047382d160.__obf_5eabb95660fc0d91(__obf_6d1c7064e94466ae,

		// func (s *Server) SendData(conn *websocket.Conn, msg any) {
		// 	s.send(conn, data, msg)
		// }
		__obf_e6e2b26604ba9b9b, __obf_0052e3c7488be106)
}

func (__obf_4f8025047382d160 *Server) ThrowError(__obf_6d1c7064e94466ae *websocket.Conn, __obf_61c9bc7c94aa4912 string) {
	__obf_4f8025047382d160.__obf_5eabb95660fc0d91(__obf_6d1c7064e94466ae,

		// RemoveClient removes the clients from the server subscription map
		__obf_548b437f9bea80cc, __obf_61c9bc7c94aa4912)
}

func (__obf_4f8025047382d160 *Server) RemoveClient(__obf_2cdc7127d0e23d64 ClientID) {
	__obf_4f8025047382d160.__obf_8fe2ef1b19fe43ad.
		Lock()
	defer __obf_4f8025047382d160.
		// loop all topics
		__obf_8fe2ef1b19fe43ad.Unlock()

	for __obf_60604e3b3dfe8815 := range __obf_4f8025047382d160.
		// delete the client from all the topic's client map
		__obf_90a259ae07de44fa {

		delete(__obf_4f8025047382d160.__obf_90a259ae07de44fa[__obf_60604e3b3dfe8815], __obf_2cdc7127d0e23d64)
	}
}

// ProcessMessage handle message according to the action type
func (__obf_4f8025047382d160 *Server) ProcessMessage(__obf_6d1c7064e94466ae *websocket.Conn, __obf_2cdc7127d0e23d64 ClientID, __obf_0052e3c7488be106 []byte) {
	__obf_2f6c3fd8eb0f80e4 := // parse message
		Message{}
	if __obf_61c9bc7c94aa4912 := json.Unmarshal(__obf_0052e3c7488be106, &__obf_2f6c3fd8eb0f80e4); __obf_61c9bc7c94aa4912 != nil {
		__obf_4f8025047382d160.
			ThrowError(__obf_6d1c7064e94466ae, ErrInvalidMessage)
		return
	}
	__obf_e6e2b26604ba9b9b := // convert all action to lowercase and remove whitespace
		strings.TrimSpace(strings.ToLower(__obf_2f6c3fd8eb0f80e4.Action))

	switch __obf_e6e2b26604ba9b9b {
	case __obf_4d69e603d2296d75:
		__obf_4f8025047382d160.
			Publish(__obf_2f6c3fd8eb0f80e4.Topic, __obf_2f6c3fd8eb0f80e4.Body)

	case __obf_317fedd18f3baea6:
		__obf_4f8025047382d160.
			Subscribe(__obf_6d1c7064e94466ae, __obf_2cdc7127d0e23d64, __obf_2f6c3fd8eb0f80e4.Topic)
	case __obf_6109636d122c33ab:
		__obf_4f8025047382d160.
			Unsubscribe(__obf_2cdc7127d0e23d64, __obf_2f6c3fd8eb0f80e4.Topic)

	default:
		if __obf_9dcf69ef75334550, __obf_6acb2478bc8a7500 := __obf_4f8025047382d160.__obf_4046692401f1241a[__obf_e6e2b26604ba9b9b]; __obf_6acb2478bc8a7500 {
			__obf_9dcf69ef75334550(__obf_2cdc7127d0e23d64, __obf_2f6c3fd8eb0f80e4.Topic, __obf_2f6c3fd8eb0f80e4.Body.(string))
		} else {
			__obf_4f8025047382d160.
				ThrowError(__obf_6d1c7064e94466ae, ErrActionUnrecognizable)
		}

	}
}

// Publish sends a message to all subscribing clients of a topic
func (__obf_4f8025047382d160 *Server) Publish(__obf_4fbd419e634c2688 Topic, __obf_eafec46ef1212797 any) {

	// if topic does not exist, stop the process
	if _, __obf_4c9a79bce95923ad := __obf_4f8025047382d160.__obf_90a259ae07de44fa[__obf_4fbd419e634c2688]; !__obf_4c9a79bce95923ad {
		return
	}
	__obf_162f952545cfa04e := // if topic exist
		__obf_4f8025047382d160.__obf_90a259ae07de44fa[__obf_4fbd419e634c2688]

	// send the message to the clients
	for _, __obf_6d1c7064e94466ae := range __obf_162f952545cfa04e {
		__obf_4f8025047382d160.
			// add 1 job to wait group
			__obf_29d5838d5b731984.
			Add(1)

		// send with goroutines
		go __obf_4f8025047382d160.SendWithWait(__obf_6d1c7064e94466ae, __obf_4d69e603d2296d75,

			// wait until all goroutines jobs done
			__obf_eafec46ef1212797)
	}
	__obf_4f8025047382d160.__obf_29d5838d5b731984.
		Wait()
}

// Kickout: kick the older session out
func (__obf_4f8025047382d160 *Server) Kickout(__obf_ea4fca7fd07b07f6 Topic, __obf_2cdc7127d0e23d64 ClientID) {
	__obf_6d1c7064e94466ae := __obf_4f8025047382d160.GetConn(__obf_ea4fca7fd07b07f6, __obf_2cdc7127d0e23d64)
	if __obf_6d1c7064e94466ae != nil {
		__obf_4f8025047382d160.__obf_5eabb95660fc0d91(__obf_6d1c7064e94466ae, __obf_9dad881b1e5a2a94, nil)
		__obf_6d1c7064e94466ae.
			Close()
		__obf_4f8025047382d160.__obf_8fe2ef1b19fe43ad.
			Lock()
		delete(__obf_4f8025047382d160.__obf_90a259ae07de44fa[__obf_ea4fca7fd07b07f6], __obf_2cdc7127d0e23d64)
		__obf_4f8025047382d160.__obf_8fe2ef1b19fe43ad.
			Unlock()
	}
}

// Subscribe adds a client to a topic's client map
func (__obf_4f8025047382d160 *Server) Subscribe(__obf_6d1c7064e94466ae *websocket.Conn, __obf_2cdc7127d0e23d64 ClientID, __obf_4fbd419e634c2688 Topic) {
	__obf_4f8025047382d160.__obf_8fe2ef1b19fe43ad.
		Lock()
	defer __obf_4f8025047382d160.
		// if topic exist, check the client map
		__obf_8fe2ef1b19fe43ad.Unlock()

	if _, __obf_4c9a79bce95923ad := __obf_4f8025047382d160.__obf_90a259ae07de44fa[__obf_4fbd419e634c2688]; __obf_4c9a79bce95923ad {
		__obf_162f952545cfa04e := __obf_4f8025047382d160.__obf_90a259ae07de44fa[__obf_4fbd419e634c2688]
		__obf_162f952545cfa04e[ // if subbed, ok := s.actionHandle[subscribe]; ok {
		// 	subbed(id, topic, "")
		// }

		__obf_2cdc7127d0e23d64] = __obf_6d1c7064e94466ae
		return
	}
	__obf_eecbceb3dd3edbb2 := // if topic does not exist, create a new topic
		make(Client)
	__obf_4f8025047382d160.__obf_90a259ae07de44fa[__obf_4fbd419e634c2688] = __obf_eecbceb3dd3edbb2

	// add the client to the topic
	__obf_4f8025047382d160.__obf_90a259ae07de44fa[__obf_4fbd419e634c2688][__obf_2cdc7127d0e23d64] = __obf_6d1c7064e94466ae
}

// Unsubscribe removes a clients from a topic's client map
func (__obf_4f8025047382d160 *Server) Unsubscribe(__obf_2cdc7127d0e23d64 ClientID, __obf_4fbd419e634c2688 Topic) {
	__obf_4f8025047382d160.__obf_8fe2ef1b19fe43ad.
		Lock()
	defer __obf_4f8025047382d160.
		// if topic exist, check the client map
		__obf_8fe2ef1b19fe43ad.Unlock()

	if _, __obf_4c9a79bce95923ad := __obf_4f8025047382d160.__obf_90a259ae07de44fa[__obf_4fbd419e634c2688]; __obf_4c9a79bce95923ad {
		__obf_162f952545cfa04e := __obf_4f8025047382d160.__obf_90a259ae07de44fa[__obf_4fbd419e634c2688]

		// remove the client from the topic's client map
		delete(__obf_162f952545cfa04e,

			// Pump just do readDump and writePump
			__obf_2cdc7127d0e23d64)
	}
}

func (__obf_4f8025047382d160 *Server) Pump(__obf_2cdc7127d0e23d64 ClientID, __obf_6d1c7064e94466ae *websocket.Conn) {
	__obf_8473167a3e3572b8 := // create channel to signal client health
		make(chan struct{})

	go __obf_4f8025047382d160.__obf_b2ca6dfff84a0bb6(__obf_6d1c7064e94466ae, __obf_2cdc7127d0e23d64, __obf_8473167a3e3572b8)
	__obf_4f8025047382d160.__obf_fa5c0c2475f3c95b(__obf_6d1c7064e94466ae,

		// readPump process incoming messages and set the settings
		__obf_2cdc7127d0e23d64, __obf_8473167a3e3572b8)
}

func (__obf_4f8025047382d160 *Server) __obf_fa5c0c2475f3c95b(__obf_6d1c7064e94466ae *websocket.Conn, __obf_2cdc7127d0e23d64 ClientID, __obf_8473167a3e3572b8 chan<- struct{}) {
	__obf_6d1c7064e94466ae.
		// set limit, deadline to read & pong handler
		SetReadLimit(__obf_4f8025047382d160.MaxMessageSize)
	__obf_6d1c7064e94466ae.
		SetReadDeadline(time.Now().Add(__obf_4f8025047382d160.PongWait * time.Second))
	__obf_6d1c7064e94466ae.
		SetPongHandler(func(string) error {
			__obf_6d1c7064e94466ae.
				SetReadDeadline(time.Now().Add(__obf_4f8025047382d160.PongWait * time.Second))
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
		_, __obf_0052e3c7488be106, __obf_61c9bc7c94aa4912 := __obf_6d1c7064e94466ae.ReadMessage()
		// if error occured
		if __obf_61c9bc7c94aa4912 != nil {
			__obf_4f8025047382d160.
				// remove from the client
				RemoveClient(__obf_2cdc7127d0e23d64)
			// set health status to unhealthy by closing channel
			close(__obf_8473167a3e3572b8)
			// stop process
			break
		}
		__obf_4f8025047382d160.

			// if no error, process incoming message
			ProcessMessage(__obf_6d1c7064e94466ae,

				// writePump sends ping to the client
				__obf_2cdc7127d0e23d64, __obf_0052e3c7488be106)
	}
}

func (__obf_4f8025047382d160 *Server) __obf_b2ca6dfff84a0bb6(__obf_6d1c7064e94466ae *websocket.Conn, __obf_2cdc7127d0e23d64 ClientID, __obf_8473167a3e3572b8 <-chan struct{}) {
	__obf_e58b54839d81bc88 := // create ping ticker
		time.NewTicker(__obf_4f8025047382d160.PingPeriod * time.Second)
	defer __obf_e58b54839d81bc88.Stop()

	for {
		select {
		case <-__obf_e58b54839d81bc88.C:
			__obf_61c9bc7c94aa4912 := // send ping message
				__obf_6d1c7064e94466ae.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(__obf_4f8025047382d160.WriteWait*time.Second))
			// log.Println("-----------------send ping message： ", err)
			if __obf_61c9bc7c94aa4912 != nil {
				__obf_4f8025047382d160.
					// log.Println("-----------------send ping err: ", err)
					// if error sending ping, remove this client from the server
					RemoveClient(__obf_2cdc7127d0e23d64)
				// stop sending ping
				return
			}
		case <-__obf_8473167a3e3572b8:
			// if process is done, stop sending ping
			return
		}
	}
}
