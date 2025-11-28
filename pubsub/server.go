package __obf_4cffc5013cde63ef

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
	__obf_1970d9dac5727f8a = "publish"
	__obf_bb1dea58983986ea = "subscribe"
	__obf_0476c517057f5aaa = "unsubscribe"
	// to client
	// data  = "data"
	__obf_1d207b025d83722a = "kickout"
	__obf_c1b6f34f93a36855 = "throw"
	__obf_6798feecc572852e = "greet"
	// popup = "popup"
)

// constants for server message
const (
	ErrInvalidMessage       = "Server: Invalid msg"
	ErrActionUnrecognizable = "Server: Action unrecognized"
)

type Topic string

// subscriber is a type for each string of topic and the clients that subscribe to it
type __obf_a45bdaae470d2d2c map[Topic]Client

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

type Handle func(ID ClientID, __obf_f2a5175cf5dcde3d Topic, __obf_4319ef689546cf96 string)

// Server is the struct to handle the Server functions & manage the subscriber
type Server struct {
	Option
	__obf_2f39669a8db446fb *sync.WaitGroup
	__obf_59cdc4dc9193195f *sync.RWMutex
	__obf_06f72e3f799c807e *sync.RWMutex
	__obf_a45bdaae470d2d2c __obf_a45bdaae470d2d2c
	__obf_f62b446b7271564a map[string]Handle
}

func NewPubsub(__obf_d631513d6a6575cd Option) *Server {
	return &Server{
		Option:                 __obf_d631513d6a6575cd,
		__obf_2f39669a8db446fb: &sync.WaitGroup{},
		__obf_59cdc4dc9193195f: &sync.RWMutex{},
		__obf_06f72e3f799c807e: &sync.RWMutex{},
		__obf_a45bdaae470d2d2c: make(__obf_a45bdaae470d2d2c),
	}
}

func (__obf_52cee885659624e3 *Server) SetActionHandle(__obf_de4b287f8a08799c string, __obf_db38023dccecc711 Handle) {

	if __obf_52cee885659624e3.__obf_f62b446b7271564a == nil {
		__obf_52cee885659624e3.__obf_f62b446b7271564a = map[string]Handle{
			__obf_de4b287f8a08799c: __obf_db38023dccecc711,
		}
	} else {
		__obf_52cee885659624e3.__obf_f62b446b7271564a[__obf_de4b287f8a08799c] = __obf_db38023dccecc711
	}
}

func (__obf_52cee885659624e3 *Server) Onlines(__obf_a2aee5c6e65493b0 Topic) int {
	return len(__obf_52cee885659624e3.__obf_a45bdaae470d2d2c[__obf_a2aee5c6e65493b0])
}

// Send simply sends message to the websocket client
func (__obf_52cee885659624e3 *Server) __obf_48d579b59735f9d6(__obf_311e5e28d75f81c2 *websocket.Conn, __obf_de4b287f8a08799c string, __obf_4319ef689546cf96 any) {
	// send simple message

	// conn.WriteJSON(Message{Action: action, Body: msg})

	// bs, _ := tool.AnyToBytes(Message{Action: action, Body: msg})
	// conn.WriteMessage(websocket.BinaryMessage, bs)
	var __obf_6b3368afbd48c1ba []byte
	__obf_6b3368afbd48c1ba, _ = util.Encode(Message{Action: __obf_de4b287f8a08799c, Body: __obf_4319ef689546cf96})

	__obf_52cee885659624e3.__obf_59cdc4dc9193195f.Lock()
	__obf_311e5e28d75f81c2.WriteMessage(websocket.TextMessage, __obf_6b3368afbd48c1ba)
	__obf_52cee885659624e3.__obf_59cdc4dc9193195f.Unlock()
}

func (__obf_52cee885659624e3 *Server) GetConn(__obf_a2aee5c6e65493b0 Topic, __obf_a78dd0f2f16f429e ClientID) *websocket.Conn {
	__obf_ba834a58c499f354, __obf_d787f515f585ddf6 := __obf_52cee885659624e3.__obf_a45bdaae470d2d2c[__obf_a2aee5c6e65493b0]
	if !__obf_d787f515f585ddf6 {
		return nil
	}

	var __obf_311e5e28d75f81c2 *websocket.Conn
	if __obf_311e5e28d75f81c2, __obf_d787f515f585ddf6 = __obf_ba834a58c499f354[__obf_a78dd0f2f16f429e]; !__obf_d787f515f585ddf6 {
		return nil
	}
	return __obf_311e5e28d75f81c2
}

func (__obf_52cee885659624e3 *Server) SendTo(__obf_a2aee5c6e65493b0 Topic, __obf_a78dd0f2f16f429e ClientID, __obf_de4b287f8a08799c string, __obf_4319ef689546cf96 any) {
	__obf_311e5e28d75f81c2 := __obf_52cee885659624e3.GetConn(__obf_a2aee5c6e65493b0, __obf_a78dd0f2f16f429e)
	if __obf_311e5e28d75f81c2 != nil {
		__obf_52cee885659624e3.__obf_48d579b59735f9d6(__obf_311e5e28d75f81c2, __obf_de4b287f8a08799c, __obf_4319ef689546cf96)
	}
}

// SendWithWait sends message to the websocket client using wait group, allowing usage with goroutines
func (__obf_52cee885659624e3 *Server) SendWithWait(__obf_311e5e28d75f81c2 *websocket.Conn, __obf_de4b287f8a08799c string, __obf_4319ef689546cf96 any) {
	// send simple message
	// conn.WriteMessage(websocket.TextMessage, []byte(tool.EncodeString(msg)))
	__obf_52cee885659624e3.__obf_48d579b59735f9d6(__obf_311e5e28d75f81c2, __obf_de4b287f8a08799c, __obf_4319ef689546cf96)

	// set the task as done
	__obf_52cee885659624e3.__obf_2f39669a8db446fb.Done()
}

func (__obf_52cee885659624e3 *Server) SendGreet(__obf_311e5e28d75f81c2 *websocket.Conn, __obf_4319ef689546cf96 any) {
	__obf_52cee885659624e3.__obf_48d579b59735f9d6(__obf_311e5e28d75f81c2, __obf_6798feecc572852e, __obf_4319ef689546cf96)
}

func (__obf_52cee885659624e3 *Server) SendWithAction(__obf_311e5e28d75f81c2 *websocket.Conn, __obf_de4b287f8a08799c string, __obf_4319ef689546cf96 any) {
	__obf_52cee885659624e3.__obf_48d579b59735f9d6(__obf_311e5e28d75f81c2, __obf_de4b287f8a08799c, __obf_4319ef689546cf96)
}

// func (s *Server) SendData(conn *websocket.Conn, msg any) {
// 	s.send(conn, data, msg)
// }

func (__obf_52cee885659624e3 *Server) ThrowError(__obf_311e5e28d75f81c2 *websocket.Conn, __obf_711c27fff4c33269 string) {
	__obf_52cee885659624e3.__obf_48d579b59735f9d6(__obf_311e5e28d75f81c2, __obf_c1b6f34f93a36855, __obf_711c27fff4c33269)
}

// RemoveClient removes the clients from the server subscription map
func (__obf_52cee885659624e3 *Server) RemoveClient(__obf_a78dd0f2f16f429e ClientID) {
	__obf_52cee885659624e3.__obf_06f72e3f799c807e.Lock()
	defer __obf_52cee885659624e3.__obf_06f72e3f799c807e.Unlock()
	// loop all topics
	for __obf_80b5fa46ceb694a9 := range __obf_52cee885659624e3.__obf_a45bdaae470d2d2c {
		// delete the client from all the topic's client map
		delete(__obf_52cee885659624e3.__obf_a45bdaae470d2d2c[__obf_80b5fa46ceb694a9], __obf_a78dd0f2f16f429e)
	}
}

// ProcessMessage handle message according to the action type
func (__obf_52cee885659624e3 *Server) ProcessMessage(__obf_311e5e28d75f81c2 *websocket.Conn, __obf_a78dd0f2f16f429e ClientID, __obf_4319ef689546cf96 []byte) {
	// parse message
	__obf_2838eba38f48cdab := Message{}
	if __obf_711c27fff4c33269 := json.Unmarshal(__obf_4319ef689546cf96, &__obf_2838eba38f48cdab); __obf_711c27fff4c33269 != nil {
		__obf_52cee885659624e3.ThrowError(__obf_311e5e28d75f81c2, ErrInvalidMessage)
		return
	}

	// convert all action to lowercase and remove whitespace
	__obf_de4b287f8a08799c := strings.TrimSpace(strings.ToLower(__obf_2838eba38f48cdab.Action))

	switch __obf_de4b287f8a08799c {
	case __obf_1970d9dac5727f8a:
		__obf_52cee885659624e3.Publish(__obf_2838eba38f48cdab.Topic, __obf_2838eba38f48cdab.Body)

	case __obf_bb1dea58983986ea:
		__obf_52cee885659624e3.Subscribe(__obf_311e5e28d75f81c2, __obf_a78dd0f2f16f429e, __obf_2838eba38f48cdab.Topic)
	case __obf_0476c517057f5aaa:
		__obf_52cee885659624e3.Unsubscribe(__obf_a78dd0f2f16f429e, __obf_2838eba38f48cdab.Topic)

	default:
		if __obf_4351965e1329fee9, __obf_91fc0b32162dbcc2 := __obf_52cee885659624e3.__obf_f62b446b7271564a[__obf_de4b287f8a08799c]; __obf_91fc0b32162dbcc2 {
			__obf_4351965e1329fee9(__obf_a78dd0f2f16f429e, __obf_2838eba38f48cdab.Topic, __obf_2838eba38f48cdab.Body.(string))
		} else {
			__obf_52cee885659624e3.ThrowError(__obf_311e5e28d75f81c2, ErrActionUnrecognizable)
		}

	}
}

// Publish sends a message to all subscribing clients of a topic
func (__obf_52cee885659624e3 *Server) Publish(__obf_f2a5175cf5dcde3d Topic, __obf_9727665ec4b557a9 any) {

	// if topic does not exist, stop the process
	if _, __obf_d787f515f585ddf6 := __obf_52cee885659624e3.__obf_a45bdaae470d2d2c[__obf_f2a5175cf5dcde3d]; !__obf_d787f515f585ddf6 {
		return
	}

	// if topic exist
	__obf_ba834a58c499f354 := __obf_52cee885659624e3.__obf_a45bdaae470d2d2c[__obf_f2a5175cf5dcde3d]

	// send the message to the clients
	for _, __obf_311e5e28d75f81c2 := range __obf_ba834a58c499f354 {
		// add 1 job to wait group
		__obf_52cee885659624e3.__obf_2f39669a8db446fb.Add(1)

		// send with goroutines
		go __obf_52cee885659624e3.SendWithWait(__obf_311e5e28d75f81c2, __obf_1970d9dac5727f8a, __obf_9727665ec4b557a9)
	}

	// wait until all goroutines jobs done
	__obf_52cee885659624e3.__obf_2f39669a8db446fb.Wait()
}

// Kickout: kick the older session out
func (__obf_52cee885659624e3 *Server) Kickout(__obf_a2aee5c6e65493b0 Topic, __obf_a78dd0f2f16f429e ClientID) {
	__obf_311e5e28d75f81c2 := __obf_52cee885659624e3.GetConn(__obf_a2aee5c6e65493b0, __obf_a78dd0f2f16f429e)
	if __obf_311e5e28d75f81c2 != nil {
		__obf_52cee885659624e3.__obf_48d579b59735f9d6(__obf_311e5e28d75f81c2, __obf_1d207b025d83722a, nil)
		__obf_311e5e28d75f81c2.Close()
		__obf_52cee885659624e3.__obf_06f72e3f799c807e.Lock()
		delete(__obf_52cee885659624e3.__obf_a45bdaae470d2d2c[__obf_a2aee5c6e65493b0], __obf_a78dd0f2f16f429e)
		__obf_52cee885659624e3.__obf_06f72e3f799c807e.Unlock()
	}
}

// Subscribe adds a client to a topic's client map
func (__obf_52cee885659624e3 *Server) Subscribe(__obf_311e5e28d75f81c2 *websocket.Conn, __obf_a78dd0f2f16f429e ClientID, __obf_f2a5175cf5dcde3d Topic) {
	__obf_52cee885659624e3.__obf_06f72e3f799c807e.Lock()
	defer __obf_52cee885659624e3.__obf_06f72e3f799c807e.Unlock()
	// if topic exist, check the client map
	if _, __obf_d787f515f585ddf6 := __obf_52cee885659624e3.__obf_a45bdaae470d2d2c[__obf_f2a5175cf5dcde3d]; __obf_d787f515f585ddf6 {
		__obf_ba834a58c499f354 := __obf_52cee885659624e3.__obf_a45bdaae470d2d2c[__obf_f2a5175cf5dcde3d]

		// if subbed, ok := s.actionHandle[subscribe]; ok {
		// 	subbed(id, topic, "")
		// }

		__obf_ba834a58c499f354[__obf_a78dd0f2f16f429e] = __obf_311e5e28d75f81c2
		return
	}

	// if topic does not exist, create a new topic
	__obf_46815bc2c1ffdf7d := make(Client)
	__obf_52cee885659624e3.__obf_a45bdaae470d2d2c[__obf_f2a5175cf5dcde3d] = __obf_46815bc2c1ffdf7d

	// add the client to the topic
	__obf_52cee885659624e3.__obf_a45bdaae470d2d2c[__obf_f2a5175cf5dcde3d][__obf_a78dd0f2f16f429e] = __obf_311e5e28d75f81c2
}

// Unsubscribe removes a clients from a topic's client map
func (__obf_52cee885659624e3 *Server) Unsubscribe(__obf_a78dd0f2f16f429e ClientID, __obf_f2a5175cf5dcde3d Topic) {
	__obf_52cee885659624e3.__obf_06f72e3f799c807e.Lock()
	defer __obf_52cee885659624e3.__obf_06f72e3f799c807e.Unlock()
	// if topic exist, check the client map
	if _, __obf_d787f515f585ddf6 := __obf_52cee885659624e3.__obf_a45bdaae470d2d2c[__obf_f2a5175cf5dcde3d]; __obf_d787f515f585ddf6 {
		__obf_ba834a58c499f354 := __obf_52cee885659624e3.__obf_a45bdaae470d2d2c[__obf_f2a5175cf5dcde3d]

		// remove the client from the topic's client map
		delete(__obf_ba834a58c499f354, __obf_a78dd0f2f16f429e)
	}
}

// Pump just do readDump and writePump
func (__obf_52cee885659624e3 *Server) Pump(__obf_a78dd0f2f16f429e ClientID, __obf_311e5e28d75f81c2 *websocket.Conn) {
	// create channel to signal client health
	__obf_bc3b801d855ef435 := make(chan struct{})

	go __obf_52cee885659624e3.__obf_05fb4f84bbd0249c(__obf_311e5e28d75f81c2, __obf_a78dd0f2f16f429e, __obf_bc3b801d855ef435)
	__obf_52cee885659624e3.__obf_1d23d3f9daa25221(__obf_311e5e28d75f81c2, __obf_a78dd0f2f16f429e, __obf_bc3b801d855ef435)
}

// readPump process incoming messages and set the settings
func (__obf_52cee885659624e3 *Server) __obf_1d23d3f9daa25221(__obf_311e5e28d75f81c2 *websocket.Conn, __obf_a78dd0f2f16f429e ClientID, __obf_bc3b801d855ef435 chan<- struct{}) {
	// set limit, deadline to read & pong handler
	__obf_311e5e28d75f81c2.SetReadLimit(__obf_52cee885659624e3.MaxMessageSize)
	__obf_311e5e28d75f81c2.SetReadDeadline(time.Now().Add(__obf_52cee885659624e3.PongWait * time.Second))
	__obf_311e5e28d75f81c2.SetPongHandler(func(string) error {
		__obf_311e5e28d75f81c2.SetReadDeadline(time.Now().Add(__obf_52cee885659624e3.PongWait * time.Second))
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
		_, __obf_4319ef689546cf96, __obf_711c27fff4c33269 := __obf_311e5e28d75f81c2.ReadMessage()
		// if error occured
		if __obf_711c27fff4c33269 != nil {
			// remove from the client
			__obf_52cee885659624e3.RemoveClient(__obf_a78dd0f2f16f429e)
			// set health status to unhealthy by closing channel
			close(__obf_bc3b801d855ef435)
			// stop process
			break
		}

		// if no error, process incoming message
		__obf_52cee885659624e3.ProcessMessage(__obf_311e5e28d75f81c2, __obf_a78dd0f2f16f429e, __obf_4319ef689546cf96)
	}
}

// writePump sends ping to the client
func (__obf_52cee885659624e3 *Server) __obf_05fb4f84bbd0249c(__obf_311e5e28d75f81c2 *websocket.Conn, __obf_a78dd0f2f16f429e ClientID, __obf_bc3b801d855ef435 <-chan struct{}) {
	// create ping ticker
	__obf_776899b26bf670d4 := time.NewTicker(__obf_52cee885659624e3.PingPeriod * time.Second)
	defer __obf_776899b26bf670d4.Stop()

	for {
		select {
		case <-__obf_776899b26bf670d4.C:
			// send ping message
			__obf_711c27fff4c33269 := __obf_311e5e28d75f81c2.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(__obf_52cee885659624e3.WriteWait*time.Second))
			// log.Println("-----------------send ping message： ", err)
			if __obf_711c27fff4c33269 != nil {
				// log.Println("-----------------send ping err: ", err)
				// if error sending ping, remove this client from the server
				__obf_52cee885659624e3.RemoveClient(__obf_a78dd0f2f16f429e)
				// stop sending ping
				return
			}
		case <-__obf_bc3b801d855ef435:
			// if process is done, stop sending ping
			return
		}
	}
}
