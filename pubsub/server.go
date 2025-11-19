package __obf_fda006bc08c20e81

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
	__obf_df0d54715a4c535f = "publish"
	__obf_398c119f585cda7e = "subscribe"
	__obf_2c9021b4df8a8db4 = "unsubscribe"
	// to client
	// data  = "data"
	__obf_0e2f438ba048fc30 = "kickout"
	__obf_e20dcf5adeb45bc7 = "throw"
	__obf_d24486f98c63848a = "greet"
	// popup = "popup"
)

// constants for server message
const (
	ErrInvalidMessage       = "Server: Invalid msg"
	ErrActionUnrecognizable = "Server: Action unrecognized"
)

type Topic string

// subscriber is a type for each string of topic and the clients that subscribe to it
type __obf_37ca54f725be081a map[Topic]Client

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

type Handle func(ID ClientID, __obf_0a50140327f02ddc Topic, __obf_3d0896c2890aa16f string)

// Server is the struct to handle the Server functions & manage the subscriber
type Server struct {
	Option
	__obf_c5a6aa2782f543d7 *sync.WaitGroup
	__obf_79570a500aa0701a *sync.RWMutex
	__obf_4e96be8ff1e25dbf *sync.RWMutex
	__obf_37ca54f725be081a __obf_37ca54f725be081a
	__obf_b4949e66449a9f3c map[string]Handle
}

func NewPubsub(__obf_ac9434c98a4452bf Option) *Server {
	return &Server{
		Option:                 __obf_ac9434c98a4452bf,
		__obf_c5a6aa2782f543d7: &sync.WaitGroup{},
		__obf_79570a500aa0701a: &sync.RWMutex{},
		__obf_4e96be8ff1e25dbf: &sync.RWMutex{},
		__obf_37ca54f725be081a: make(__obf_37ca54f725be081a),
	}
}

func (__obf_2478267a0a2ff751 *Server) SetActionHandle(__obf_d69511bf2896b013 string, __obf_a8bf0b9515030b99 Handle) {

	if __obf_2478267a0a2ff751.__obf_b4949e66449a9f3c == nil {
		__obf_2478267a0a2ff751.__obf_b4949e66449a9f3c = map[string]Handle{
			__obf_d69511bf2896b013: __obf_a8bf0b9515030b99,
		}
	} else {
		__obf_2478267a0a2ff751.__obf_b4949e66449a9f3c[__obf_d69511bf2896b013] = __obf_a8bf0b9515030b99
	}
}

func (__obf_2478267a0a2ff751 *Server) Onlines(__obf_dd1a6169b313a4d5 Topic) int {
	return len(__obf_2478267a0a2ff751.__obf_37ca54f725be081a[__obf_dd1a6169b313a4d5])
}

// Send simply sends message to the websocket client
func (__obf_2478267a0a2ff751 *Server) __obf_45022cf8cccf6baf(__obf_d971a8958298f1dd *websocket.Conn, __obf_d69511bf2896b013 string, __obf_3d0896c2890aa16f any) {
	// send simple message

	// conn.WriteJSON(Message{Action: action, Body: msg})

	// bs, _ := tool.AnyToBytes(Message{Action: action, Body: msg})
	// conn.WriteMessage(websocket.BinaryMessage, bs)
	var __obf_9d147980afc482a8 []byte
	__obf_9d147980afc482a8, _ = util.Encode(Message{Action: __obf_d69511bf2896b013, Body: __obf_3d0896c2890aa16f})

	__obf_2478267a0a2ff751.__obf_79570a500aa0701a.Lock()
	__obf_d971a8958298f1dd.WriteMessage(websocket.TextMessage, __obf_9d147980afc482a8)
	__obf_2478267a0a2ff751.__obf_79570a500aa0701a.Unlock()
}

func (__obf_2478267a0a2ff751 *Server) GetConn(__obf_dd1a6169b313a4d5 Topic, __obf_b1c87d8f10255f3d ClientID) *websocket.Conn {
	__obf_641c65eca55701d8, __obf_bf04a8a0978b2743 := __obf_2478267a0a2ff751.__obf_37ca54f725be081a[__obf_dd1a6169b313a4d5]
	if !__obf_bf04a8a0978b2743 {
		return nil
	}

	var __obf_d971a8958298f1dd *websocket.Conn
	if __obf_d971a8958298f1dd, __obf_bf04a8a0978b2743 = __obf_641c65eca55701d8[__obf_b1c87d8f10255f3d]; !__obf_bf04a8a0978b2743 {
		return nil
	}
	return __obf_d971a8958298f1dd
}

func (__obf_2478267a0a2ff751 *Server) SendTo(__obf_dd1a6169b313a4d5 Topic, __obf_b1c87d8f10255f3d ClientID, __obf_d69511bf2896b013 string, __obf_3d0896c2890aa16f any) {
	__obf_d971a8958298f1dd := __obf_2478267a0a2ff751.GetConn(__obf_dd1a6169b313a4d5, __obf_b1c87d8f10255f3d)
	if __obf_d971a8958298f1dd != nil {
		__obf_2478267a0a2ff751.__obf_45022cf8cccf6baf(__obf_d971a8958298f1dd, __obf_d69511bf2896b013, __obf_3d0896c2890aa16f)
	}
}

// SendWithWait sends message to the websocket client using wait group, allowing usage with goroutines
func (__obf_2478267a0a2ff751 *Server) SendWithWait(__obf_d971a8958298f1dd *websocket.Conn, __obf_d69511bf2896b013 string, __obf_3d0896c2890aa16f any) {
	// send simple message
	// conn.WriteMessage(websocket.TextMessage, []byte(tool.EncodeString(msg)))
	__obf_2478267a0a2ff751.__obf_45022cf8cccf6baf(__obf_d971a8958298f1dd, __obf_d69511bf2896b013, __obf_3d0896c2890aa16f)

	// set the task as done
	__obf_2478267a0a2ff751.__obf_c5a6aa2782f543d7.Done()
}

func (__obf_2478267a0a2ff751 *Server) SendGreet(__obf_d971a8958298f1dd *websocket.Conn, __obf_3d0896c2890aa16f any) {
	__obf_2478267a0a2ff751.__obf_45022cf8cccf6baf(__obf_d971a8958298f1dd, __obf_d24486f98c63848a, __obf_3d0896c2890aa16f)
}

func (__obf_2478267a0a2ff751 *Server) SendWithAction(__obf_d971a8958298f1dd *websocket.Conn, __obf_d69511bf2896b013 string, __obf_3d0896c2890aa16f any) {
	__obf_2478267a0a2ff751.__obf_45022cf8cccf6baf(__obf_d971a8958298f1dd, __obf_d69511bf2896b013, __obf_3d0896c2890aa16f)
}

// func (s *Server) SendData(conn *websocket.Conn, msg any) {
// 	s.send(conn, data, msg)
// }

func (__obf_2478267a0a2ff751 *Server) ThrowError(__obf_d971a8958298f1dd *websocket.Conn, __obf_92845df685cfc9c8 string) {
	__obf_2478267a0a2ff751.__obf_45022cf8cccf6baf(__obf_d971a8958298f1dd, __obf_e20dcf5adeb45bc7, __obf_92845df685cfc9c8)
}

// RemoveClient removes the clients from the server subscription map
func (__obf_2478267a0a2ff751 *Server) RemoveClient(__obf_b1c87d8f10255f3d ClientID) {
	__obf_2478267a0a2ff751.__obf_4e96be8ff1e25dbf.Lock()
	defer __obf_2478267a0a2ff751.__obf_4e96be8ff1e25dbf.Unlock()
	// loop all topics
	for __obf_2d620b44160a0afe := range __obf_2478267a0a2ff751.__obf_37ca54f725be081a {
		// delete the client from all the topic's client map
		delete(__obf_2478267a0a2ff751.__obf_37ca54f725be081a[__obf_2d620b44160a0afe], __obf_b1c87d8f10255f3d)
	}
}

// ProcessMessage handle message according to the action type
func (__obf_2478267a0a2ff751 *Server) ProcessMessage(__obf_d971a8958298f1dd *websocket.Conn, __obf_b1c87d8f10255f3d ClientID, __obf_3d0896c2890aa16f []byte) {
	// parse message
	__obf_64c00b9205499a30 := Message{}
	if __obf_92845df685cfc9c8 := json.Unmarshal(__obf_3d0896c2890aa16f, &__obf_64c00b9205499a30); __obf_92845df685cfc9c8 != nil {
		__obf_2478267a0a2ff751.ThrowError(__obf_d971a8958298f1dd, ErrInvalidMessage)
		return
	}

	// convert all action to lowercase and remove whitespace
	__obf_d69511bf2896b013 := strings.TrimSpace(strings.ToLower(__obf_64c00b9205499a30.Action))

	switch __obf_d69511bf2896b013 {
	case __obf_df0d54715a4c535f:
		__obf_2478267a0a2ff751.Publish(__obf_64c00b9205499a30.Topic, __obf_64c00b9205499a30.Body)

	case __obf_398c119f585cda7e:
		__obf_2478267a0a2ff751.Subscribe(__obf_d971a8958298f1dd, __obf_b1c87d8f10255f3d, __obf_64c00b9205499a30.Topic)
	case __obf_2c9021b4df8a8db4:
		__obf_2478267a0a2ff751.Unsubscribe(__obf_b1c87d8f10255f3d, __obf_64c00b9205499a30.Topic)

	default:
		if __obf_98b4690372ce927d, __obf_7f12278359712ef2 := __obf_2478267a0a2ff751.__obf_b4949e66449a9f3c[__obf_d69511bf2896b013]; __obf_7f12278359712ef2 {
			__obf_98b4690372ce927d(__obf_b1c87d8f10255f3d, __obf_64c00b9205499a30.Topic, __obf_64c00b9205499a30.Body.(string))
		} else {
			__obf_2478267a0a2ff751.ThrowError(__obf_d971a8958298f1dd, ErrActionUnrecognizable)
		}

	}
}

// Publish sends a message to all subscribing clients of a topic
func (__obf_2478267a0a2ff751 *Server) Publish(__obf_0a50140327f02ddc Topic, __obf_f668a68ae34c3a90 any) {

	// if topic does not exist, stop the process
	if _, __obf_bf04a8a0978b2743 := __obf_2478267a0a2ff751.__obf_37ca54f725be081a[__obf_0a50140327f02ddc]; !__obf_bf04a8a0978b2743 {
		return
	}

	// if topic exist
	__obf_641c65eca55701d8 := __obf_2478267a0a2ff751.__obf_37ca54f725be081a[__obf_0a50140327f02ddc]

	// send the message to the clients
	for _, __obf_d971a8958298f1dd := range __obf_641c65eca55701d8 {
		// add 1 job to wait group
		__obf_2478267a0a2ff751.__obf_c5a6aa2782f543d7.Add(1)

		// send with goroutines
		go __obf_2478267a0a2ff751.SendWithWait(__obf_d971a8958298f1dd, __obf_df0d54715a4c535f, __obf_f668a68ae34c3a90)
	}

	// wait until all goroutines jobs done
	__obf_2478267a0a2ff751.__obf_c5a6aa2782f543d7.Wait()
}

// Kickout: kick the older session out
func (__obf_2478267a0a2ff751 *Server) Kickout(__obf_dd1a6169b313a4d5 Topic, __obf_b1c87d8f10255f3d ClientID) {
	__obf_d971a8958298f1dd := __obf_2478267a0a2ff751.GetConn(__obf_dd1a6169b313a4d5, __obf_b1c87d8f10255f3d)
	if __obf_d971a8958298f1dd != nil {
		__obf_2478267a0a2ff751.__obf_45022cf8cccf6baf(__obf_d971a8958298f1dd, __obf_0e2f438ba048fc30, nil)
		__obf_d971a8958298f1dd.Close()
		__obf_2478267a0a2ff751.__obf_4e96be8ff1e25dbf.Lock()
		delete(__obf_2478267a0a2ff751.__obf_37ca54f725be081a[__obf_dd1a6169b313a4d5], __obf_b1c87d8f10255f3d)
		__obf_2478267a0a2ff751.__obf_4e96be8ff1e25dbf.Unlock()
	}
}

// Subscribe adds a client to a topic's client map
func (__obf_2478267a0a2ff751 *Server) Subscribe(__obf_d971a8958298f1dd *websocket.Conn, __obf_b1c87d8f10255f3d ClientID, __obf_0a50140327f02ddc Topic) {
	__obf_2478267a0a2ff751.__obf_4e96be8ff1e25dbf.Lock()
	defer __obf_2478267a0a2ff751.__obf_4e96be8ff1e25dbf.Unlock()
	// if topic exist, check the client map
	if _, __obf_bf04a8a0978b2743 := __obf_2478267a0a2ff751.__obf_37ca54f725be081a[__obf_0a50140327f02ddc]; __obf_bf04a8a0978b2743 {
		__obf_641c65eca55701d8 := __obf_2478267a0a2ff751.__obf_37ca54f725be081a[__obf_0a50140327f02ddc]

		// if subbed, ok := s.actionHandle[subscribe]; ok {
		// 	subbed(id, topic, "")
		// }

		__obf_641c65eca55701d8[__obf_b1c87d8f10255f3d] = __obf_d971a8958298f1dd
		return
	}

	// if topic does not exist, create a new topic
	__obf_76a9cb033e9848a7 := make(Client)
	__obf_2478267a0a2ff751.__obf_37ca54f725be081a[__obf_0a50140327f02ddc] = __obf_76a9cb033e9848a7

	// add the client to the topic
	__obf_2478267a0a2ff751.__obf_37ca54f725be081a[__obf_0a50140327f02ddc][__obf_b1c87d8f10255f3d] = __obf_d971a8958298f1dd
}

// Unsubscribe removes a clients from a topic's client map
func (__obf_2478267a0a2ff751 *Server) Unsubscribe(__obf_b1c87d8f10255f3d ClientID, __obf_0a50140327f02ddc Topic) {
	__obf_2478267a0a2ff751.__obf_4e96be8ff1e25dbf.Lock()
	defer __obf_2478267a0a2ff751.__obf_4e96be8ff1e25dbf.Unlock()
	// if topic exist, check the client map
	if _, __obf_bf04a8a0978b2743 := __obf_2478267a0a2ff751.__obf_37ca54f725be081a[__obf_0a50140327f02ddc]; __obf_bf04a8a0978b2743 {
		__obf_641c65eca55701d8 := __obf_2478267a0a2ff751.__obf_37ca54f725be081a[__obf_0a50140327f02ddc]

		// remove the client from the topic's client map
		delete(__obf_641c65eca55701d8, __obf_b1c87d8f10255f3d)
	}
}

// Pump just do readDump and writePump
func (__obf_2478267a0a2ff751 *Server) Pump(__obf_b1c87d8f10255f3d ClientID, __obf_d971a8958298f1dd *websocket.Conn) {
	// create channel to signal client health
	__obf_14b49b73906fed14 := make(chan struct{})

	go __obf_2478267a0a2ff751.__obf_b3175f834b15edbf(__obf_d971a8958298f1dd, __obf_b1c87d8f10255f3d, __obf_14b49b73906fed14)
	__obf_2478267a0a2ff751.__obf_dbcd9d5a988e234f(__obf_d971a8958298f1dd, __obf_b1c87d8f10255f3d, __obf_14b49b73906fed14)
}

// readPump process incoming messages and set the settings
func (__obf_2478267a0a2ff751 *Server) __obf_dbcd9d5a988e234f(__obf_d971a8958298f1dd *websocket.Conn, __obf_b1c87d8f10255f3d ClientID, __obf_14b49b73906fed14 chan<- struct{}) {
	// set limit, deadline to read & pong handler
	__obf_d971a8958298f1dd.SetReadLimit(__obf_2478267a0a2ff751.MaxMessageSize)
	__obf_d971a8958298f1dd.SetReadDeadline(time.Now().Add(__obf_2478267a0a2ff751.PongWait * time.Second))
	__obf_d971a8958298f1dd.SetPongHandler(func(string) error {
		__obf_d971a8958298f1dd.SetReadDeadline(time.Now().Add(__obf_2478267a0a2ff751.PongWait * time.Second))
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
		_, __obf_3d0896c2890aa16f, __obf_92845df685cfc9c8 := __obf_d971a8958298f1dd.ReadMessage()
		// if error occured
		if __obf_92845df685cfc9c8 != nil {
			// remove from the client
			__obf_2478267a0a2ff751.RemoveClient(__obf_b1c87d8f10255f3d)
			// set health status to unhealthy by closing channel
			close(__obf_14b49b73906fed14)
			// stop process
			break
		}

		// if no error, process incoming message
		__obf_2478267a0a2ff751.ProcessMessage(__obf_d971a8958298f1dd, __obf_b1c87d8f10255f3d, __obf_3d0896c2890aa16f)
	}
}

// writePump sends ping to the client
func (__obf_2478267a0a2ff751 *Server) __obf_b3175f834b15edbf(__obf_d971a8958298f1dd *websocket.Conn, __obf_b1c87d8f10255f3d ClientID, __obf_14b49b73906fed14 <-chan struct{}) {
	// create ping ticker
	__obf_670a344344b0f5cb := time.NewTicker(__obf_2478267a0a2ff751.PingPeriod * time.Second)
	defer __obf_670a344344b0f5cb.Stop()

	for {
		select {
		case <-__obf_670a344344b0f5cb.C:
			// send ping message
			__obf_92845df685cfc9c8 := __obf_d971a8958298f1dd.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(__obf_2478267a0a2ff751.WriteWait*time.Second))
			// log.Println("-----------------send ping message： ", err)
			if __obf_92845df685cfc9c8 != nil {
				// log.Println("-----------------send ping err: ", err)
				// if error sending ping, remove this client from the server
				__obf_2478267a0a2ff751.RemoveClient(__obf_b1c87d8f10255f3d)
				// stop sending ping
				return
			}
		case <-__obf_14b49b73906fed14:
			// if process is done, stop sending ping
			return
		}
	}
}
