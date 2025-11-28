package __obf_a574f3926693cb21

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
	__obf_f899bcf8c6c42f2d = "publish"
	__obf_c1901ed5836a09cc = "subscribe"
	__obf_89c31460018b2ca3 = "unsubscribe"
	// to client
	// data  = "data"
	__obf_581ea99dfe601d24 = "kickout"
	__obf_213abe9219323328 = "throw"
	__obf_3b0e17cb23388ecd = "greet"
	// popup = "popup"
)

// constants for server message
const (
	ErrInvalidMessage       = "Server: Invalid msg"
	ErrActionUnrecognizable = "Server: Action unrecognized"
)

type Topic string

// subscriber is a type for each string of topic and the clients that subscribe to it
type __obf_3f0233b2a328acbc map[Topic]Client

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

type Handle func(ID ClientID, __obf_2364a266770771d9 Topic, __obf_7426b86937588ed5 string)

// Server is the struct to handle the Server functions & manage the subscriber
type Server struct {
	Option
	__obf_b399cf1ec2bb53dc *sync.WaitGroup
	__obf_8adf20f54d30e2a4 *sync.RWMutex
	__obf_ffa4ea3a5142c1ad *sync.RWMutex
	__obf_3f0233b2a328acbc __obf_3f0233b2a328acbc
	__obf_b5531bfac0b94b0f map[string]Handle
}

func NewPubsub(__obf_a3e8d7d6b39dc1f0 Option) *Server {
	return &Server{
		Option:                 __obf_a3e8d7d6b39dc1f0,
		__obf_b399cf1ec2bb53dc: &sync.WaitGroup{},
		__obf_8adf20f54d30e2a4: &sync.RWMutex{},
		__obf_ffa4ea3a5142c1ad: &sync.RWMutex{},
		__obf_3f0233b2a328acbc: make(__obf_3f0233b2a328acbc),
	}
}

func (__obf_28eff38c1f84d0f8 *Server) SetActionHandle(__obf_510d4d28f4ff0f0f string, __obf_e750ceaf2e0f92d5 Handle) {

	if __obf_28eff38c1f84d0f8.__obf_b5531bfac0b94b0f == nil {
		__obf_28eff38c1f84d0f8.__obf_b5531bfac0b94b0f = map[string]Handle{
			__obf_510d4d28f4ff0f0f: __obf_e750ceaf2e0f92d5,
		}
	} else {
		__obf_28eff38c1f84d0f8.__obf_b5531bfac0b94b0f[__obf_510d4d28f4ff0f0f] = __obf_e750ceaf2e0f92d5
	}
}

func (__obf_28eff38c1f84d0f8 *Server) Onlines(__obf_43bb274d564864ca Topic) int {
	return len(__obf_28eff38c1f84d0f8.__obf_3f0233b2a328acbc[__obf_43bb274d564864ca])
}

// Send simply sends message to the websocket client
func (__obf_28eff38c1f84d0f8 *Server) __obf_90db7806db4ae65f(__obf_22407855e75bf514 *websocket.Conn, __obf_510d4d28f4ff0f0f string, __obf_7426b86937588ed5 any) {
	// send simple message

	// conn.WriteJSON(Message{Action: action, Body: msg})

	// bs, _ := tool.AnyToBytes(Message{Action: action, Body: msg})
	// conn.WriteMessage(websocket.BinaryMessage, bs)
	var __obf_d86ae2ddf811ca8f []byte
	__obf_d86ae2ddf811ca8f, _ = util.Encode(Message{Action: __obf_510d4d28f4ff0f0f, Body: __obf_7426b86937588ed5})

	__obf_28eff38c1f84d0f8.__obf_8adf20f54d30e2a4.Lock()
	__obf_22407855e75bf514.WriteMessage(websocket.TextMessage, __obf_d86ae2ddf811ca8f)
	__obf_28eff38c1f84d0f8.__obf_8adf20f54d30e2a4.Unlock()
}

func (__obf_28eff38c1f84d0f8 *Server) GetConn(__obf_43bb274d564864ca Topic, __obf_06348a4a0f4bf6b6 ClientID) *websocket.Conn {
	__obf_6988ea74fed3ec4d, __obf_e50b8d259c4ba792 := __obf_28eff38c1f84d0f8.__obf_3f0233b2a328acbc[__obf_43bb274d564864ca]
	if !__obf_e50b8d259c4ba792 {
		return nil
	}

	var __obf_22407855e75bf514 *websocket.Conn
	if __obf_22407855e75bf514, __obf_e50b8d259c4ba792 = __obf_6988ea74fed3ec4d[__obf_06348a4a0f4bf6b6]; !__obf_e50b8d259c4ba792 {
		return nil
	}
	return __obf_22407855e75bf514
}

func (__obf_28eff38c1f84d0f8 *Server) SendTo(__obf_43bb274d564864ca Topic, __obf_06348a4a0f4bf6b6 ClientID, __obf_510d4d28f4ff0f0f string, __obf_7426b86937588ed5 any) {
	__obf_22407855e75bf514 := __obf_28eff38c1f84d0f8.GetConn(__obf_43bb274d564864ca, __obf_06348a4a0f4bf6b6)
	if __obf_22407855e75bf514 != nil {
		__obf_28eff38c1f84d0f8.__obf_90db7806db4ae65f(__obf_22407855e75bf514, __obf_510d4d28f4ff0f0f, __obf_7426b86937588ed5)
	}
}

// SendWithWait sends message to the websocket client using wait group, allowing usage with goroutines
func (__obf_28eff38c1f84d0f8 *Server) SendWithWait(__obf_22407855e75bf514 *websocket.Conn, __obf_510d4d28f4ff0f0f string, __obf_7426b86937588ed5 any) {
	// send simple message
	// conn.WriteMessage(websocket.TextMessage, []byte(tool.EncodeString(msg)))
	__obf_28eff38c1f84d0f8.__obf_90db7806db4ae65f(__obf_22407855e75bf514, __obf_510d4d28f4ff0f0f, __obf_7426b86937588ed5)

	// set the task as done
	__obf_28eff38c1f84d0f8.__obf_b399cf1ec2bb53dc.Done()
}

func (__obf_28eff38c1f84d0f8 *Server) SendGreet(__obf_22407855e75bf514 *websocket.Conn, __obf_7426b86937588ed5 any) {
	__obf_28eff38c1f84d0f8.__obf_90db7806db4ae65f(__obf_22407855e75bf514, __obf_3b0e17cb23388ecd, __obf_7426b86937588ed5)
}

func (__obf_28eff38c1f84d0f8 *Server) SendWithAction(__obf_22407855e75bf514 *websocket.Conn, __obf_510d4d28f4ff0f0f string, __obf_7426b86937588ed5 any) {
	__obf_28eff38c1f84d0f8.__obf_90db7806db4ae65f(__obf_22407855e75bf514, __obf_510d4d28f4ff0f0f, __obf_7426b86937588ed5)
}

// func (s *Server) SendData(conn *websocket.Conn, msg any) {
// 	s.send(conn, data, msg)
// }

func (__obf_28eff38c1f84d0f8 *Server) ThrowError(__obf_22407855e75bf514 *websocket.Conn, __obf_68bacaf5c47d7561 string) {
	__obf_28eff38c1f84d0f8.__obf_90db7806db4ae65f(__obf_22407855e75bf514, __obf_213abe9219323328, __obf_68bacaf5c47d7561)
}

// RemoveClient removes the clients from the server subscription map
func (__obf_28eff38c1f84d0f8 *Server) RemoveClient(__obf_06348a4a0f4bf6b6 ClientID) {
	__obf_28eff38c1f84d0f8.__obf_ffa4ea3a5142c1ad.Lock()
	defer __obf_28eff38c1f84d0f8.__obf_ffa4ea3a5142c1ad.Unlock()
	// loop all topics
	for __obf_5cae698c57658033 := range __obf_28eff38c1f84d0f8.__obf_3f0233b2a328acbc {
		// delete the client from all the topic's client map
		delete(__obf_28eff38c1f84d0f8.__obf_3f0233b2a328acbc[__obf_5cae698c57658033], __obf_06348a4a0f4bf6b6)
	}
}

// ProcessMessage handle message according to the action type
func (__obf_28eff38c1f84d0f8 *Server) ProcessMessage(__obf_22407855e75bf514 *websocket.Conn, __obf_06348a4a0f4bf6b6 ClientID, __obf_7426b86937588ed5 []byte) {
	// parse message
	__obf_336914d70eb27de4 := Message{}
	if __obf_68bacaf5c47d7561 := json.Unmarshal(__obf_7426b86937588ed5, &__obf_336914d70eb27de4); __obf_68bacaf5c47d7561 != nil {
		__obf_28eff38c1f84d0f8.ThrowError(__obf_22407855e75bf514, ErrInvalidMessage)
		return
	}

	// convert all action to lowercase and remove whitespace
	__obf_510d4d28f4ff0f0f := strings.TrimSpace(strings.ToLower(__obf_336914d70eb27de4.Action))

	switch __obf_510d4d28f4ff0f0f {
	case __obf_f899bcf8c6c42f2d:
		__obf_28eff38c1f84d0f8.Publish(__obf_336914d70eb27de4.Topic, __obf_336914d70eb27de4.Body)

	case __obf_c1901ed5836a09cc:
		__obf_28eff38c1f84d0f8.Subscribe(__obf_22407855e75bf514, __obf_06348a4a0f4bf6b6, __obf_336914d70eb27de4.Topic)
	case __obf_89c31460018b2ca3:
		__obf_28eff38c1f84d0f8.Unsubscribe(__obf_06348a4a0f4bf6b6, __obf_336914d70eb27de4.Topic)

	default:
		if __obf_859f177148682f6a, __obf_e5af5b1a1544d854 := __obf_28eff38c1f84d0f8.__obf_b5531bfac0b94b0f[__obf_510d4d28f4ff0f0f]; __obf_e5af5b1a1544d854 {
			__obf_859f177148682f6a(__obf_06348a4a0f4bf6b6, __obf_336914d70eb27de4.Topic, __obf_336914d70eb27de4.Body.(string))
		} else {
			__obf_28eff38c1f84d0f8.ThrowError(__obf_22407855e75bf514, ErrActionUnrecognizable)
		}

	}
}

// Publish sends a message to all subscribing clients of a topic
func (__obf_28eff38c1f84d0f8 *Server) Publish(__obf_2364a266770771d9 Topic, __obf_b28ebc07da5fc12b any) {

	// if topic does not exist, stop the process
	if _, __obf_e50b8d259c4ba792 := __obf_28eff38c1f84d0f8.__obf_3f0233b2a328acbc[__obf_2364a266770771d9]; !__obf_e50b8d259c4ba792 {
		return
	}

	// if topic exist
	__obf_6988ea74fed3ec4d := __obf_28eff38c1f84d0f8.__obf_3f0233b2a328acbc[__obf_2364a266770771d9]

	// send the message to the clients
	for _, __obf_22407855e75bf514 := range __obf_6988ea74fed3ec4d {
		// add 1 job to wait group
		__obf_28eff38c1f84d0f8.__obf_b399cf1ec2bb53dc.Add(1)

		// send with goroutines
		go __obf_28eff38c1f84d0f8.SendWithWait(__obf_22407855e75bf514, __obf_f899bcf8c6c42f2d, __obf_b28ebc07da5fc12b)
	}

	// wait until all goroutines jobs done
	__obf_28eff38c1f84d0f8.__obf_b399cf1ec2bb53dc.Wait()
}

// Kickout: kick the older session out
func (__obf_28eff38c1f84d0f8 *Server) Kickout(__obf_43bb274d564864ca Topic, __obf_06348a4a0f4bf6b6 ClientID) {
	__obf_22407855e75bf514 := __obf_28eff38c1f84d0f8.GetConn(__obf_43bb274d564864ca, __obf_06348a4a0f4bf6b6)
	if __obf_22407855e75bf514 != nil {
		__obf_28eff38c1f84d0f8.__obf_90db7806db4ae65f(__obf_22407855e75bf514, __obf_581ea99dfe601d24, nil)
		__obf_22407855e75bf514.Close()
		__obf_28eff38c1f84d0f8.__obf_ffa4ea3a5142c1ad.Lock()
		delete(__obf_28eff38c1f84d0f8.__obf_3f0233b2a328acbc[__obf_43bb274d564864ca], __obf_06348a4a0f4bf6b6)
		__obf_28eff38c1f84d0f8.__obf_ffa4ea3a5142c1ad.Unlock()
	}
}

// Subscribe adds a client to a topic's client map
func (__obf_28eff38c1f84d0f8 *Server) Subscribe(__obf_22407855e75bf514 *websocket.Conn, __obf_06348a4a0f4bf6b6 ClientID, __obf_2364a266770771d9 Topic) {
	__obf_28eff38c1f84d0f8.__obf_ffa4ea3a5142c1ad.Lock()
	defer __obf_28eff38c1f84d0f8.__obf_ffa4ea3a5142c1ad.Unlock()
	// if topic exist, check the client map
	if _, __obf_e50b8d259c4ba792 := __obf_28eff38c1f84d0f8.__obf_3f0233b2a328acbc[__obf_2364a266770771d9]; __obf_e50b8d259c4ba792 {
		__obf_6988ea74fed3ec4d := __obf_28eff38c1f84d0f8.__obf_3f0233b2a328acbc[__obf_2364a266770771d9]

		// if subbed, ok := s.actionHandle[subscribe]; ok {
		// 	subbed(id, topic, "")
		// }

		__obf_6988ea74fed3ec4d[__obf_06348a4a0f4bf6b6] = __obf_22407855e75bf514
		return
	}

	// if topic does not exist, create a new topic
	__obf_397936648cfeb89d := make(Client)
	__obf_28eff38c1f84d0f8.__obf_3f0233b2a328acbc[__obf_2364a266770771d9] = __obf_397936648cfeb89d

	// add the client to the topic
	__obf_28eff38c1f84d0f8.__obf_3f0233b2a328acbc[__obf_2364a266770771d9][__obf_06348a4a0f4bf6b6] = __obf_22407855e75bf514
}

// Unsubscribe removes a clients from a topic's client map
func (__obf_28eff38c1f84d0f8 *Server) Unsubscribe(__obf_06348a4a0f4bf6b6 ClientID, __obf_2364a266770771d9 Topic) {
	__obf_28eff38c1f84d0f8.__obf_ffa4ea3a5142c1ad.Lock()
	defer __obf_28eff38c1f84d0f8.__obf_ffa4ea3a5142c1ad.Unlock()
	// if topic exist, check the client map
	if _, __obf_e50b8d259c4ba792 := __obf_28eff38c1f84d0f8.__obf_3f0233b2a328acbc[__obf_2364a266770771d9]; __obf_e50b8d259c4ba792 {
		__obf_6988ea74fed3ec4d := __obf_28eff38c1f84d0f8.__obf_3f0233b2a328acbc[__obf_2364a266770771d9]

		// remove the client from the topic's client map
		delete(__obf_6988ea74fed3ec4d, __obf_06348a4a0f4bf6b6)
	}
}

// Pump just do readDump and writePump
func (__obf_28eff38c1f84d0f8 *Server) Pump(__obf_06348a4a0f4bf6b6 ClientID, __obf_22407855e75bf514 *websocket.Conn) {
	// create channel to signal client health
	__obf_fcb4a979b7fabda6 := make(chan struct{})

	go __obf_28eff38c1f84d0f8.__obf_1d772446583b0fee(__obf_22407855e75bf514, __obf_06348a4a0f4bf6b6, __obf_fcb4a979b7fabda6)
	__obf_28eff38c1f84d0f8.__obf_ce6fd58bae61c6c3(__obf_22407855e75bf514, __obf_06348a4a0f4bf6b6, __obf_fcb4a979b7fabda6)
}

// readPump process incoming messages and set the settings
func (__obf_28eff38c1f84d0f8 *Server) __obf_ce6fd58bae61c6c3(__obf_22407855e75bf514 *websocket.Conn, __obf_06348a4a0f4bf6b6 ClientID, __obf_fcb4a979b7fabda6 chan<- struct{}) {
	// set limit, deadline to read & pong handler
	__obf_22407855e75bf514.SetReadLimit(__obf_28eff38c1f84d0f8.MaxMessageSize)
	__obf_22407855e75bf514.SetReadDeadline(time.Now().Add(__obf_28eff38c1f84d0f8.PongWait * time.Second))
	__obf_22407855e75bf514.SetPongHandler(func(string) error {
		__obf_22407855e75bf514.SetReadDeadline(time.Now().Add(__obf_28eff38c1f84d0f8.PongWait * time.Second))
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
		_, __obf_7426b86937588ed5, __obf_68bacaf5c47d7561 := __obf_22407855e75bf514.ReadMessage()
		// if error occured
		if __obf_68bacaf5c47d7561 != nil {
			// remove from the client
			__obf_28eff38c1f84d0f8.RemoveClient(__obf_06348a4a0f4bf6b6)
			// set health status to unhealthy by closing channel
			close(__obf_fcb4a979b7fabda6)
			// stop process
			break
		}

		// if no error, process incoming message
		__obf_28eff38c1f84d0f8.ProcessMessage(__obf_22407855e75bf514, __obf_06348a4a0f4bf6b6, __obf_7426b86937588ed5)
	}
}

// writePump sends ping to the client
func (__obf_28eff38c1f84d0f8 *Server) __obf_1d772446583b0fee(__obf_22407855e75bf514 *websocket.Conn, __obf_06348a4a0f4bf6b6 ClientID, __obf_fcb4a979b7fabda6 <-chan struct{}) {
	// create ping ticker
	__obf_a6bd6b9e1e7f2762 := time.NewTicker(__obf_28eff38c1f84d0f8.PingPeriod * time.Second)
	defer __obf_a6bd6b9e1e7f2762.Stop()

	for {
		select {
		case <-__obf_a6bd6b9e1e7f2762.C:
			// send ping message
			__obf_68bacaf5c47d7561 := __obf_22407855e75bf514.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(__obf_28eff38c1f84d0f8.WriteWait*time.Second))
			// log.Println("-----------------send ping message： ", err)
			if __obf_68bacaf5c47d7561 != nil {
				// log.Println("-----------------send ping err: ", err)
				// if error sending ping, remove this client from the server
				__obf_28eff38c1f84d0f8.RemoveClient(__obf_06348a4a0f4bf6b6)
				// stop sending ping
				return
			}
		case <-__obf_fcb4a979b7fabda6:
			// if process is done, stop sending ping
			return
		}
	}
}
