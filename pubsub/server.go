package __obf_e2ef7a91a3555fad

import (
	"encoding/json"
	"strings"
	"sync"
	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/gorilla/websocket"
)

// constants for action type
const (
	// from client
	__obf_3923868da733fc40 = "publish"
	__obf_6b2674e8747ecfd2 = "subscribe"
	__obf_07f6009aa575cda6 = "unsubscribe"
	// to client
	// data  = "data"
	__obf_c300000a524f0ac6 = "kickout"
	__obf_f5f0c148715276df = "throw"
	__obf_a04cd2282e01f146 = "greet"
	// popup = "popup"
)

// constants for server message
const (
	ErrInvalidMessage       = "Server: Invalid msg"
	ErrActionUnrecognizable = "Server: Action unrecognized"
)

type Topic string

// subscriber is a type for each string of topic and the clients that subscribe to it
type __obf_7282d94ba42b919f map[Topic]Client

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

type Handle func(ID ClientID, __obf_33593883a7613182 Topic, __obf_0f59921e2a2b409a string)

// Server is the struct to handle the Server functions & manage the subscriber
type Server struct {
	Option
	__obf_7beae969dc830b00 *sync.WaitGroup
	__obf_f7e591f1dcab1417 *sync.RWMutex
	__obf_ebe15b9f8fb713b2 *sync.RWMutex
	__obf_7282d94ba42b919f __obf_7282d94ba42b919f
	__obf_f9f17c492bc15196 map[string]Handle
}

func NewPubsub(__obf_b243066a46c37afc Option) *Server {
	return &Server{
		Option:                 __obf_b243066a46c37afc,
		__obf_7beae969dc830b00: &sync.WaitGroup{},
		__obf_f7e591f1dcab1417: &sync.RWMutex{},
		__obf_ebe15b9f8fb713b2: &sync.RWMutex{},
		__obf_7282d94ba42b919f: make(__obf_7282d94ba42b919f),
	}
}

func (__obf_3b8435971c7fa957 *Server) SetActionHandle(__obf_b0cfd848666e9e65 string, __obf_6e0827639b7d01de Handle) {

	if __obf_3b8435971c7fa957.__obf_f9f17c492bc15196 == nil {
		__obf_3b8435971c7fa957.__obf_f9f17c492bc15196 = map[string]Handle{
			__obf_b0cfd848666e9e65: __obf_6e0827639b7d01de,
		}
	} else {
		__obf_3b8435971c7fa957.__obf_f9f17c492bc15196[__obf_b0cfd848666e9e65] = __obf_6e0827639b7d01de
	}
}

func (__obf_3b8435971c7fa957 *Server) Onlines(__obf_b509d7d4759b6253 Topic) int {
	return len(__obf_3b8435971c7fa957.__obf_7282d94ba42b919f[__obf_b509d7d4759b6253])
}

// Send simply sends message to the websocket client
func (__obf_3b8435971c7fa957 *Server) __obf_003d36d93eb75bfd(__obf_2686046e06db786d *websocket.Conn, __obf_b0cfd848666e9e65 string, __obf_0f59921e2a2b409a any) {
	// send simple message

	// conn.WriteJSON(Message{Action: action, Body: msg})

	// bs, _ := tool.AnyToBytes(Message{Action: action, Body: msg})
	// conn.WriteMessage(websocket.BinaryMessage, bs)
	var __obf_99151bb402c1b0a1 []byte
	__obf_99151bb402c1b0a1, _ = util.Encode(Message{Action: __obf_b0cfd848666e9e65, Body: __obf_0f59921e2a2b409a})

	__obf_3b8435971c7fa957.__obf_f7e591f1dcab1417.Lock()
	__obf_2686046e06db786d.WriteMessage(websocket.TextMessage, __obf_99151bb402c1b0a1)
	__obf_3b8435971c7fa957.__obf_f7e591f1dcab1417.Unlock()
}

func (__obf_3b8435971c7fa957 *Server) GetConn(__obf_b509d7d4759b6253 Topic, __obf_8b0e4516a014d4e8 ClientID) *websocket.Conn {
	__obf_a9b27ce9600bd5c6, __obf_768bc0899b98f1b8 := __obf_3b8435971c7fa957.__obf_7282d94ba42b919f[__obf_b509d7d4759b6253]
	if !__obf_768bc0899b98f1b8 {
		return nil
	}

	var __obf_2686046e06db786d *websocket.Conn
	if __obf_2686046e06db786d, __obf_768bc0899b98f1b8 = __obf_a9b27ce9600bd5c6[__obf_8b0e4516a014d4e8]; !__obf_768bc0899b98f1b8 {
		return nil
	}
	return __obf_2686046e06db786d
}

func (__obf_3b8435971c7fa957 *Server) SendTo(__obf_b509d7d4759b6253 Topic, __obf_8b0e4516a014d4e8 ClientID, __obf_b0cfd848666e9e65 string, __obf_0f59921e2a2b409a any) {
	__obf_2686046e06db786d := __obf_3b8435971c7fa957.GetConn(__obf_b509d7d4759b6253, __obf_8b0e4516a014d4e8)
	if __obf_2686046e06db786d != nil {
		__obf_3b8435971c7fa957.__obf_003d36d93eb75bfd(__obf_2686046e06db786d, __obf_b0cfd848666e9e65, __obf_0f59921e2a2b409a)
	}
}

// SendWithWait sends message to the websocket client using wait group, allowing usage with goroutines
func (__obf_3b8435971c7fa957 *Server) SendWithWait(__obf_2686046e06db786d *websocket.Conn, __obf_b0cfd848666e9e65 string, __obf_0f59921e2a2b409a any) {
	// send simple message
	// conn.WriteMessage(websocket.TextMessage, []byte(tool.EncodeString(msg)))
	__obf_3b8435971c7fa957.__obf_003d36d93eb75bfd(__obf_2686046e06db786d, __obf_b0cfd848666e9e65, __obf_0f59921e2a2b409a)

	// set the task as done
	__obf_3b8435971c7fa957.__obf_7beae969dc830b00.Done()
}

func (__obf_3b8435971c7fa957 *Server) SendGreet(__obf_2686046e06db786d *websocket.Conn, __obf_0f59921e2a2b409a any) {
	__obf_3b8435971c7fa957.__obf_003d36d93eb75bfd(__obf_2686046e06db786d, __obf_a04cd2282e01f146, __obf_0f59921e2a2b409a)
}

func (__obf_3b8435971c7fa957 *Server) SendWithAction(__obf_2686046e06db786d *websocket.Conn, __obf_b0cfd848666e9e65 string, __obf_0f59921e2a2b409a any) {
	__obf_3b8435971c7fa957.__obf_003d36d93eb75bfd(__obf_2686046e06db786d, __obf_b0cfd848666e9e65, __obf_0f59921e2a2b409a)
}

// func (s *Server) SendData(conn *websocket.Conn, msg any) {
// 	s.send(conn, data, msg)
// }

func (__obf_3b8435971c7fa957 *Server) ThrowError(__obf_2686046e06db786d *websocket.Conn, __obf_4e555280590a47e8 string) {
	__obf_3b8435971c7fa957.__obf_003d36d93eb75bfd(__obf_2686046e06db786d, __obf_f5f0c148715276df, __obf_4e555280590a47e8)
}

// RemoveClient removes the clients from the server subscription map
func (__obf_3b8435971c7fa957 *Server) RemoveClient(__obf_8b0e4516a014d4e8 ClientID) {
	__obf_3b8435971c7fa957.__obf_ebe15b9f8fb713b2.Lock()
	defer __obf_3b8435971c7fa957.__obf_ebe15b9f8fb713b2.Unlock()
	// loop all topics
	for __obf_04d055bb934fd4ff := range __obf_3b8435971c7fa957.__obf_7282d94ba42b919f {
		// delete the client from all the topic's client map
		delete(__obf_3b8435971c7fa957.__obf_7282d94ba42b919f[__obf_04d055bb934fd4ff], __obf_8b0e4516a014d4e8)
	}
}

// ProcessMessage handle message according to the action type
func (__obf_3b8435971c7fa957 *Server) ProcessMessage(__obf_2686046e06db786d *websocket.Conn, __obf_8b0e4516a014d4e8 ClientID, __obf_0f59921e2a2b409a []byte) {
	// parse message
	__obf_f4b5b70225df7b66 := Message{}
	if __obf_4e555280590a47e8 := json.Unmarshal(__obf_0f59921e2a2b409a, &__obf_f4b5b70225df7b66); __obf_4e555280590a47e8 != nil {
		__obf_3b8435971c7fa957.ThrowError(__obf_2686046e06db786d, ErrInvalidMessage)
		return
	}

	// convert all action to lowercase and remove whitespace
	__obf_b0cfd848666e9e65 := strings.TrimSpace(strings.ToLower(__obf_f4b5b70225df7b66.Action))

	switch __obf_b0cfd848666e9e65 {
	case __obf_3923868da733fc40:
		__obf_3b8435971c7fa957.Publish(__obf_f4b5b70225df7b66.Topic, __obf_f4b5b70225df7b66.Body)

	case __obf_6b2674e8747ecfd2:
		__obf_3b8435971c7fa957.Subscribe(__obf_2686046e06db786d, __obf_8b0e4516a014d4e8, __obf_f4b5b70225df7b66.Topic)
	case __obf_07f6009aa575cda6:
		__obf_3b8435971c7fa957.Unsubscribe(__obf_8b0e4516a014d4e8, __obf_f4b5b70225df7b66.Topic)

	default:
		if __obf_ba89833a94db168e, __obf_47a34d5d4ec9c901 := __obf_3b8435971c7fa957.__obf_f9f17c492bc15196[__obf_b0cfd848666e9e65]; __obf_47a34d5d4ec9c901 {
			__obf_ba89833a94db168e(__obf_8b0e4516a014d4e8, __obf_f4b5b70225df7b66.Topic, __obf_f4b5b70225df7b66.Body.(string))
		} else {
			__obf_3b8435971c7fa957.ThrowError(__obf_2686046e06db786d, ErrActionUnrecognizable)
		}

	}
}

// Publish sends a message to all subscribing clients of a topic
func (__obf_3b8435971c7fa957 *Server) Publish(__obf_33593883a7613182 Topic, __obf_ca67c2911c040490 any) {

	// if topic does not exist, stop the process
	if _, __obf_768bc0899b98f1b8 := __obf_3b8435971c7fa957.__obf_7282d94ba42b919f[__obf_33593883a7613182]; !__obf_768bc0899b98f1b8 {
		return
	}

	// if topic exist
	__obf_a9b27ce9600bd5c6 := __obf_3b8435971c7fa957.__obf_7282d94ba42b919f[__obf_33593883a7613182]

	// send the message to the clients
	for _, __obf_2686046e06db786d := range __obf_a9b27ce9600bd5c6 {
		// add 1 job to wait group
		__obf_3b8435971c7fa957.__obf_7beae969dc830b00.Add(1)

		// send with goroutines
		go __obf_3b8435971c7fa957.SendWithWait(__obf_2686046e06db786d, __obf_3923868da733fc40, __obf_ca67c2911c040490)
	}

	// wait until all goroutines jobs done
	__obf_3b8435971c7fa957.__obf_7beae969dc830b00.Wait()
}

// Kickout: kick the older session out
func (__obf_3b8435971c7fa957 *Server) Kickout(__obf_b509d7d4759b6253 Topic, __obf_8b0e4516a014d4e8 ClientID) {
	__obf_2686046e06db786d := __obf_3b8435971c7fa957.GetConn(__obf_b509d7d4759b6253, __obf_8b0e4516a014d4e8)
	if __obf_2686046e06db786d != nil {
		__obf_3b8435971c7fa957.__obf_003d36d93eb75bfd(__obf_2686046e06db786d, __obf_c300000a524f0ac6, nil)
		__obf_2686046e06db786d.Close()
		__obf_3b8435971c7fa957.__obf_ebe15b9f8fb713b2.Lock()
		delete(__obf_3b8435971c7fa957.__obf_7282d94ba42b919f[__obf_b509d7d4759b6253], __obf_8b0e4516a014d4e8)
		__obf_3b8435971c7fa957.__obf_ebe15b9f8fb713b2.Unlock()
	}
}

// Subscribe adds a client to a topic's client map
func (__obf_3b8435971c7fa957 *Server) Subscribe(__obf_2686046e06db786d *websocket.Conn, __obf_8b0e4516a014d4e8 ClientID, __obf_33593883a7613182 Topic) {
	__obf_3b8435971c7fa957.__obf_ebe15b9f8fb713b2.Lock()
	defer __obf_3b8435971c7fa957.__obf_ebe15b9f8fb713b2.Unlock()
	// if topic exist, check the client map
	if _, __obf_768bc0899b98f1b8 := __obf_3b8435971c7fa957.__obf_7282d94ba42b919f[__obf_33593883a7613182]; __obf_768bc0899b98f1b8 {
		__obf_a9b27ce9600bd5c6 := __obf_3b8435971c7fa957.__obf_7282d94ba42b919f[__obf_33593883a7613182]

		// if subbed, ok := s.actionHandle[subscribe]; ok {
		// 	subbed(id, topic, "")
		// }

		__obf_a9b27ce9600bd5c6[__obf_8b0e4516a014d4e8] = __obf_2686046e06db786d
		return
	}

	// if topic does not exist, create a new topic
	__obf_32074d47ca73c896 := make(Client)
	__obf_3b8435971c7fa957.__obf_7282d94ba42b919f[__obf_33593883a7613182] = __obf_32074d47ca73c896

	// add the client to the topic
	__obf_3b8435971c7fa957.__obf_7282d94ba42b919f[__obf_33593883a7613182][__obf_8b0e4516a014d4e8] = __obf_2686046e06db786d
}

// Unsubscribe removes a clients from a topic's client map
func (__obf_3b8435971c7fa957 *Server) Unsubscribe(__obf_8b0e4516a014d4e8 ClientID, __obf_33593883a7613182 Topic) {
	__obf_3b8435971c7fa957.__obf_ebe15b9f8fb713b2.Lock()
	defer __obf_3b8435971c7fa957.__obf_ebe15b9f8fb713b2.Unlock()
	// if topic exist, check the client map
	if _, __obf_768bc0899b98f1b8 := __obf_3b8435971c7fa957.__obf_7282d94ba42b919f[__obf_33593883a7613182]; __obf_768bc0899b98f1b8 {
		__obf_a9b27ce9600bd5c6 := __obf_3b8435971c7fa957.__obf_7282d94ba42b919f[__obf_33593883a7613182]

		// remove the client from the topic's client map
		delete(__obf_a9b27ce9600bd5c6, __obf_8b0e4516a014d4e8)
	}
}

// Pump just do readDump and writePump
func (__obf_3b8435971c7fa957 *Server) Pump(__obf_8b0e4516a014d4e8 ClientID, __obf_2686046e06db786d *websocket.Conn) {
	// create channel to signal client health
	__obf_1513f9112e06169f := make(chan struct{})

	go __obf_3b8435971c7fa957.__obf_fcdb190d649dc077(__obf_2686046e06db786d, __obf_8b0e4516a014d4e8, __obf_1513f9112e06169f)
	__obf_3b8435971c7fa957.__obf_d4a4114f45959e76(__obf_2686046e06db786d, __obf_8b0e4516a014d4e8, __obf_1513f9112e06169f)
}

// readPump process incoming messages and set the settings
func (__obf_3b8435971c7fa957 *Server) __obf_d4a4114f45959e76(__obf_2686046e06db786d *websocket.Conn, __obf_8b0e4516a014d4e8 ClientID, __obf_1513f9112e06169f chan<- struct{}) {
	// set limit, deadline to read & pong handler
	__obf_2686046e06db786d.SetReadLimit(__obf_3b8435971c7fa957.MaxMessageSize)
	__obf_2686046e06db786d.SetReadDeadline(time.Now().Add(__obf_3b8435971c7fa957.PongWait * time.Second))
	__obf_2686046e06db786d.SetPongHandler(func(string) error {
		__obf_2686046e06db786d.SetReadDeadline(time.Now().Add(__obf_3b8435971c7fa957.PongWait * time.Second))
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
		_, __obf_0f59921e2a2b409a, __obf_4e555280590a47e8 := __obf_2686046e06db786d.ReadMessage()
		// if error occured
		if __obf_4e555280590a47e8 != nil {
			// remove from the client
			__obf_3b8435971c7fa957.RemoveClient(__obf_8b0e4516a014d4e8)
			// set health status to unhealthy by closing channel
			close(__obf_1513f9112e06169f)
			// stop process
			break
		}

		// if no error, process incoming message
		__obf_3b8435971c7fa957.ProcessMessage(__obf_2686046e06db786d, __obf_8b0e4516a014d4e8, __obf_0f59921e2a2b409a)
	}
}

// writePump sends ping to the client
func (__obf_3b8435971c7fa957 *Server) __obf_fcdb190d649dc077(__obf_2686046e06db786d *websocket.Conn, __obf_8b0e4516a014d4e8 ClientID, __obf_1513f9112e06169f <-chan struct{}) {
	// create ping ticker
	__obf_c8191df05df053c3 := time.NewTicker(__obf_3b8435971c7fa957.PingPeriod * time.Second)
	defer __obf_c8191df05df053c3.Stop()

	for {
		select {
		case <-__obf_c8191df05df053c3.C:
			// send ping message
			__obf_4e555280590a47e8 := __obf_2686046e06db786d.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(__obf_3b8435971c7fa957.WriteWait*time.Second))
			// log.Println("-----------------send ping message： ", err)
			if __obf_4e555280590a47e8 != nil {
				// log.Println("-----------------send ping err: ", err)
				// if error sending ping, remove this client from the server
				__obf_3b8435971c7fa957.RemoveClient(__obf_8b0e4516a014d4e8)
				// stop sending ping
				return
			}
		case <-__obf_1513f9112e06169f:
			// if process is done, stop sending ping
			return
		}
	}
}
