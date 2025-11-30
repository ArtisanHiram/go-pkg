package __obf_6145545ace35c006

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
	__obf_f5e6095368a8dca6 =
	// from client
	"publish"
	__obf_d1a7816d1ab11fcb = "subscribe"
	__obf_39b93939b512bfc3 = "unsubscribe"
	__obf_9882c7b74f9c9bd0 =
	// to client
	// data  = "data"
	"kickout"
	__obf_5cd60c70d696059f = "throw"
	__obf_6790234288627317 = "greet"
	// popup = "popup"
)

// constants for server message
const (
	ErrInvalidMessage       = "Server: Invalid msg"
	ErrActionUnrecognizable = "Server: Action unrecognized"
)

type Topic string

// subscriber is a type for each string of topic and the clients that subscribe to it
type __obf_a7efe3898ea66ffc map[Topic]Client

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

type Handle func(ID ClientID, __obf_577399377b1737a3 Topic, __obf_5be414915ded40a0 string)

// Server is the struct to handle the Server functions & manage the subscriber
type Server struct {
	Option
	__obf_b4ec0789d0002207 *sync.WaitGroup
	__obf_5094e51209b1cb40 *sync.RWMutex
	__obf_6e082324d0deec7d *sync.RWMutex
	__obf_a7efe3898ea66ffc __obf_a7efe3898ea66ffc
	__obf_b00d59b739d81a38 map[string]Handle
}

func NewPubsub(__obf_51b576d896706931 Option) *Server {
	return &Server{
		Option: __obf_51b576d896706931, __obf_b4ec0789d0002207: &sync.WaitGroup{}, __obf_5094e51209b1cb40: &sync.RWMutex{}, __obf_6e082324d0deec7d: &sync.RWMutex{}, __obf_a7efe3898ea66ffc: make(__obf_a7efe3898ea66ffc),
	}
}

func (__obf_76110f4857a6ae2c *Server) SetActionHandle(__obf_6a4de6b55dc7e622 string, __obf_24f7d0db77cdadc6 Handle) {

	if __obf_76110f4857a6ae2c.__obf_b00d59b739d81a38 == nil {
		__obf_76110f4857a6ae2c.__obf_b00d59b739d81a38 = map[string]Handle{__obf_6a4de6b55dc7e622: __obf_24f7d0db77cdadc6}
	} else {
		__obf_76110f4857a6ae2c.__obf_b00d59b739d81a38[__obf_6a4de6b55dc7e622] = __obf_24f7d0db77cdadc6
	}
}

func (__obf_76110f4857a6ae2c *Server) Onlines(__obf_b7afae5e9783bd02 Topic) int {
	return len(__obf_76110f4857a6ae2c.

		// Send simply sends message to the websocket client
		__obf_a7efe3898ea66ffc[__obf_b7afae5e9783bd02])
}

func (__obf_76110f4857a6ae2c *Server) __obf_056dad473188b1bf(__obf_5544fcc9eb8bca89 *websocket.Conn, __obf_6a4de6b55dc7e622 string, __obf_5be414915ded40a0 any) {
	// send simple message

	// conn.WriteJSON(Message{Action: action, Body: msg})

	// bs, _ := tool.AnyToBytes(Message{Action: action, Body: msg})
	// conn.WriteMessage(websocket.BinaryMessage, bs)
	var __obf_c3ce8d6aef60bd99 []byte
	__obf_c3ce8d6aef60bd99, _ = util.Encode(Message{Action: __obf_6a4de6b55dc7e622, Body: __obf_5be414915ded40a0})
	__obf_76110f4857a6ae2c.__obf_5094e51209b1cb40.
		Lock()
	__obf_5544fcc9eb8bca89.
		WriteMessage(websocket.TextMessage, __obf_c3ce8d6aef60bd99)
	__obf_76110f4857a6ae2c.__obf_5094e51209b1cb40.
		Unlock()
}

func (__obf_76110f4857a6ae2c *Server) GetConn(__obf_b7afae5e9783bd02 Topic, __obf_7a2df81d1dc821d1 ClientID) *websocket.Conn {
	__obf_aa9717628b8a945a, __obf_a19eb0063e60014b := __obf_76110f4857a6ae2c.__obf_a7efe3898ea66ffc[__obf_b7afae5e9783bd02]
	if !__obf_a19eb0063e60014b {
		return nil
	}

	var __obf_5544fcc9eb8bca89 *websocket.Conn
	if __obf_5544fcc9eb8bca89, __obf_a19eb0063e60014b = __obf_aa9717628b8a945a[__obf_7a2df81d1dc821d1]; !__obf_a19eb0063e60014b {
		return nil
	}
	return __obf_5544fcc9eb8bca89
}

func (__obf_76110f4857a6ae2c *Server) SendTo(__obf_b7afae5e9783bd02 Topic, __obf_7a2df81d1dc821d1 ClientID, __obf_6a4de6b55dc7e622 string, __obf_5be414915ded40a0 any) {
	__obf_5544fcc9eb8bca89 := __obf_76110f4857a6ae2c.GetConn(__obf_b7afae5e9783bd02, __obf_7a2df81d1dc821d1)
	if __obf_5544fcc9eb8bca89 != nil {
		__obf_76110f4857a6ae2c.__obf_056dad473188b1bf(__obf_5544fcc9eb8bca89, __obf_6a4de6b55dc7e622,

			// SendWithWait sends message to the websocket client using wait group, allowing usage with goroutines
			__obf_5be414915ded40a0)
	}
}

func (__obf_76110f4857a6ae2c *Server) SendWithWait(__obf_5544fcc9eb8bca89 *websocket.Conn, __obf_6a4de6b55dc7e622 string, __obf_5be414915ded40a0 any) {
	__obf_76110f4857a6ae2c.
		// send simple message
		// conn.WriteMessage(websocket.TextMessage, []byte(tool.EncodeString(msg)))
		__obf_056dad473188b1bf(__obf_5544fcc9eb8bca89,

			// set the task as done
			__obf_6a4de6b55dc7e622, __obf_5be414915ded40a0)
	__obf_76110f4857a6ae2c.__obf_b4ec0789d0002207.
		Done()
}

func (__obf_76110f4857a6ae2c *Server) SendGreet(__obf_5544fcc9eb8bca89 *websocket.Conn, __obf_5be414915ded40a0 any) {
	__obf_76110f4857a6ae2c.__obf_056dad473188b1bf(__obf_5544fcc9eb8bca89, __obf_6790234288627317, __obf_5be414915ded40a0)
}

func (__obf_76110f4857a6ae2c *Server) SendWithAction(__obf_5544fcc9eb8bca89 *websocket.Conn, __obf_6a4de6b55dc7e622 string, __obf_5be414915ded40a0 any) {
	__obf_76110f4857a6ae2c.__obf_056dad473188b1bf(__obf_5544fcc9eb8bca89,

		// func (s *Server) SendData(conn *websocket.Conn, msg any) {
		// 	s.send(conn, data, msg)
		// }
		__obf_6a4de6b55dc7e622, __obf_5be414915ded40a0)
}

func (__obf_76110f4857a6ae2c *Server) ThrowError(__obf_5544fcc9eb8bca89 *websocket.Conn, __obf_ea1c3ce45f4fd63a string) {
	__obf_76110f4857a6ae2c.__obf_056dad473188b1bf(__obf_5544fcc9eb8bca89,

		// RemoveClient removes the clients from the server subscription map
		__obf_5cd60c70d696059f, __obf_ea1c3ce45f4fd63a)
}

func (__obf_76110f4857a6ae2c *Server) RemoveClient(__obf_7a2df81d1dc821d1 ClientID) {
	__obf_76110f4857a6ae2c.__obf_6e082324d0deec7d.
		Lock()
	defer __obf_76110f4857a6ae2c.
		// loop all topics
		__obf_6e082324d0deec7d.Unlock()

	for __obf_350a0aedb8a25ec8 := range __obf_76110f4857a6ae2c.
		// delete the client from all the topic's client map
		__obf_a7efe3898ea66ffc {

		delete(__obf_76110f4857a6ae2c.__obf_a7efe3898ea66ffc[__obf_350a0aedb8a25ec8], __obf_7a2df81d1dc821d1)
	}
}

// ProcessMessage handle message according to the action type
func (__obf_76110f4857a6ae2c *Server) ProcessMessage(__obf_5544fcc9eb8bca89 *websocket.Conn, __obf_7a2df81d1dc821d1 ClientID, __obf_5be414915ded40a0 []byte) {
	__obf_1217135181f2b020 := // parse message
		Message{}
	if __obf_ea1c3ce45f4fd63a := json.Unmarshal(__obf_5be414915ded40a0, &__obf_1217135181f2b020); __obf_ea1c3ce45f4fd63a != nil {
		__obf_76110f4857a6ae2c.
			ThrowError(__obf_5544fcc9eb8bca89, ErrInvalidMessage)
		return
	}
	__obf_6a4de6b55dc7e622 := // convert all action to lowercase and remove whitespace
		strings.TrimSpace(strings.ToLower(__obf_1217135181f2b020.Action))

	switch __obf_6a4de6b55dc7e622 {
	case __obf_f5e6095368a8dca6:
		__obf_76110f4857a6ae2c.
			Publish(__obf_1217135181f2b020.Topic, __obf_1217135181f2b020.Body)

	case __obf_d1a7816d1ab11fcb:
		__obf_76110f4857a6ae2c.
			Subscribe(__obf_5544fcc9eb8bca89, __obf_7a2df81d1dc821d1, __obf_1217135181f2b020.Topic)
	case __obf_39b93939b512bfc3:
		__obf_76110f4857a6ae2c.
			Unsubscribe(__obf_7a2df81d1dc821d1, __obf_1217135181f2b020.Topic)

	default:
		if __obf_89701418da2e3d63, __obf_51aa442b2366bca6 := __obf_76110f4857a6ae2c.__obf_b00d59b739d81a38[__obf_6a4de6b55dc7e622]; __obf_51aa442b2366bca6 {
			__obf_89701418da2e3d63(__obf_7a2df81d1dc821d1, __obf_1217135181f2b020.Topic, __obf_1217135181f2b020.Body.(string))
		} else {
			__obf_76110f4857a6ae2c.
				ThrowError(__obf_5544fcc9eb8bca89, ErrActionUnrecognizable)
		}

	}
}

// Publish sends a message to all subscribing clients of a topic
func (__obf_76110f4857a6ae2c *Server) Publish(__obf_577399377b1737a3 Topic, __obf_84c7a937fa9eda75 any) {

	// if topic does not exist, stop the process
	if _, __obf_a19eb0063e60014b := __obf_76110f4857a6ae2c.__obf_a7efe3898ea66ffc[__obf_577399377b1737a3]; !__obf_a19eb0063e60014b {
		return
	}
	__obf_aa9717628b8a945a := // if topic exist
		__obf_76110f4857a6ae2c.__obf_a7efe3898ea66ffc[__obf_577399377b1737a3]

	// send the message to the clients
	for _, __obf_5544fcc9eb8bca89 := range __obf_aa9717628b8a945a {
		__obf_76110f4857a6ae2c.
			// add 1 job to wait group
			__obf_b4ec0789d0002207.
			Add(1)

		// send with goroutines
		go __obf_76110f4857a6ae2c.SendWithWait(__obf_5544fcc9eb8bca89, __obf_f5e6095368a8dca6,

			// wait until all goroutines jobs done
			__obf_84c7a937fa9eda75)
	}
	__obf_76110f4857a6ae2c.__obf_b4ec0789d0002207.
		Wait()
}

// Kickout: kick the older session out
func (__obf_76110f4857a6ae2c *Server) Kickout(__obf_b7afae5e9783bd02 Topic, __obf_7a2df81d1dc821d1 ClientID) {
	__obf_5544fcc9eb8bca89 := __obf_76110f4857a6ae2c.GetConn(__obf_b7afae5e9783bd02, __obf_7a2df81d1dc821d1)
	if __obf_5544fcc9eb8bca89 != nil {
		__obf_76110f4857a6ae2c.__obf_056dad473188b1bf(__obf_5544fcc9eb8bca89, __obf_9882c7b74f9c9bd0, nil)
		__obf_5544fcc9eb8bca89.
			Close()
		__obf_76110f4857a6ae2c.__obf_6e082324d0deec7d.
			Lock()
		delete(__obf_76110f4857a6ae2c.__obf_a7efe3898ea66ffc[__obf_b7afae5e9783bd02], __obf_7a2df81d1dc821d1)
		__obf_76110f4857a6ae2c.__obf_6e082324d0deec7d.
			Unlock()
	}
}

// Subscribe adds a client to a topic's client map
func (__obf_76110f4857a6ae2c *Server) Subscribe(__obf_5544fcc9eb8bca89 *websocket.Conn, __obf_7a2df81d1dc821d1 ClientID, __obf_577399377b1737a3 Topic) {
	__obf_76110f4857a6ae2c.__obf_6e082324d0deec7d.
		Lock()
	defer __obf_76110f4857a6ae2c.
		// if topic exist, check the client map
		__obf_6e082324d0deec7d.Unlock()

	if _, __obf_a19eb0063e60014b := __obf_76110f4857a6ae2c.__obf_a7efe3898ea66ffc[__obf_577399377b1737a3]; __obf_a19eb0063e60014b {
		__obf_aa9717628b8a945a := __obf_76110f4857a6ae2c.__obf_a7efe3898ea66ffc[__obf_577399377b1737a3]
		__obf_aa9717628b8a945a[ // if subbed, ok := s.actionHandle[subscribe]; ok {
		// 	subbed(id, topic, "")
		// }

		__obf_7a2df81d1dc821d1] = __obf_5544fcc9eb8bca89
		return
	}
	__obf_c477170ebac14537 := // if topic does not exist, create a new topic
		make(Client)
	__obf_76110f4857a6ae2c.__obf_a7efe3898ea66ffc[__obf_577399377b1737a3] = __obf_c477170ebac14537

	// add the client to the topic
	__obf_76110f4857a6ae2c.__obf_a7efe3898ea66ffc[__obf_577399377b1737a3][__obf_7a2df81d1dc821d1] = __obf_5544fcc9eb8bca89
}

// Unsubscribe removes a clients from a topic's client map
func (__obf_76110f4857a6ae2c *Server) Unsubscribe(__obf_7a2df81d1dc821d1 ClientID, __obf_577399377b1737a3 Topic) {
	__obf_76110f4857a6ae2c.__obf_6e082324d0deec7d.
		Lock()
	defer __obf_76110f4857a6ae2c.
		// if topic exist, check the client map
		__obf_6e082324d0deec7d.Unlock()

	if _, __obf_a19eb0063e60014b := __obf_76110f4857a6ae2c.__obf_a7efe3898ea66ffc[__obf_577399377b1737a3]; __obf_a19eb0063e60014b {
		__obf_aa9717628b8a945a := __obf_76110f4857a6ae2c.__obf_a7efe3898ea66ffc[__obf_577399377b1737a3]

		// remove the client from the topic's client map
		delete(__obf_aa9717628b8a945a,

			// Pump just do readDump and writePump
			__obf_7a2df81d1dc821d1)
	}
}

func (__obf_76110f4857a6ae2c *Server) Pump(__obf_7a2df81d1dc821d1 ClientID, __obf_5544fcc9eb8bca89 *websocket.Conn) {
	__obf_f785aaf051f900fe := // create channel to signal client health
		make(chan struct{})

	go __obf_76110f4857a6ae2c.__obf_ab0704dff73ec9c9(__obf_5544fcc9eb8bca89, __obf_7a2df81d1dc821d1, __obf_f785aaf051f900fe)
	__obf_76110f4857a6ae2c.__obf_8c21f84ec4f94cb5(__obf_5544fcc9eb8bca89,

		// readPump process incoming messages and set the settings
		__obf_7a2df81d1dc821d1, __obf_f785aaf051f900fe)
}

func (__obf_76110f4857a6ae2c *Server) __obf_8c21f84ec4f94cb5(__obf_5544fcc9eb8bca89 *websocket.Conn, __obf_7a2df81d1dc821d1 ClientID, __obf_f785aaf051f900fe chan<- struct{}) {
	__obf_5544fcc9eb8bca89.
		// set limit, deadline to read & pong handler
		SetReadLimit(__obf_76110f4857a6ae2c.MaxMessageSize)
	__obf_5544fcc9eb8bca89.
		SetReadDeadline(time.Now().Add(__obf_76110f4857a6ae2c.PongWait * time.Second))
	__obf_5544fcc9eb8bca89.
		SetPongHandler(func(string) error {
			__obf_5544fcc9eb8bca89.
				SetReadDeadline(time.Now().Add(__obf_76110f4857a6ae2c.PongWait * time.Second))
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
		_, __obf_5be414915ded40a0, __obf_ea1c3ce45f4fd63a := __obf_5544fcc9eb8bca89.ReadMessage()
		// if error occured
		if __obf_ea1c3ce45f4fd63a != nil {
			__obf_76110f4857a6ae2c.
				// remove from the client
				RemoveClient(__obf_7a2df81d1dc821d1)
			// set health status to unhealthy by closing channel
			close(__obf_f785aaf051f900fe)
			// stop process
			break
		}
		__obf_76110f4857a6ae2c.

			// if no error, process incoming message
			ProcessMessage(__obf_5544fcc9eb8bca89,

				// writePump sends ping to the client
				__obf_7a2df81d1dc821d1, __obf_5be414915ded40a0)
	}
}

func (__obf_76110f4857a6ae2c *Server) __obf_ab0704dff73ec9c9(__obf_5544fcc9eb8bca89 *websocket.Conn, __obf_7a2df81d1dc821d1 ClientID, __obf_f785aaf051f900fe <-chan struct{}) {
	__obf_f6d8b177eedda256 := // create ping ticker
		time.NewTicker(__obf_76110f4857a6ae2c.PingPeriod * time.Second)
	defer __obf_f6d8b177eedda256.Stop()

	for {
		select {
		case <-__obf_f6d8b177eedda256.C:
			__obf_ea1c3ce45f4fd63a := // send ping message
				__obf_5544fcc9eb8bca89.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(__obf_76110f4857a6ae2c.WriteWait*time.Second))
			// log.Println("-----------------send ping message： ", err)
			if __obf_ea1c3ce45f4fd63a != nil {
				__obf_76110f4857a6ae2c.
					// log.Println("-----------------send ping err: ", err)
					// if error sending ping, remove this client from the server
					RemoveClient(__obf_7a2df81d1dc821d1)
				// stop sending ping
				return
			}
		case <-__obf_f785aaf051f900fe:
			// if process is done, stop sending ping
			return
		}
	}
}
