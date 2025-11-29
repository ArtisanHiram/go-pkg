package __obf_775ad2e2fb710130

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
	__obf_3d0cb88ba67024f3 =
	// from client
	"publish"
	__obf_20d4bfa70f8c734c = "subscribe"
	__obf_f6fc30e0b49f4d3d = "unsubscribe"
	__obf_078850448f7a5f5c =
	// to client
	// data  = "data"
	"kickout"
	__obf_96035d4f1f7782ff = "throw"
	__obf_55896f7b03405734 = "greet"
	// popup = "popup"
)

// constants for server message
const (
	ErrInvalidMessage       = "Server: Invalid msg"
	ErrActionUnrecognizable = "Server: Action unrecognized"
)

type Topic string

// subscriber is a type for each string of topic and the clients that subscribe to it
type __obf_73097b41cd7fdc67 map[Topic]Client

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

type Handle func(ID ClientID, __obf_061bc10d62c1be15 Topic, __obf_d29e6ec337e55cb9 string)

// Server is the struct to handle the Server functions & manage the subscriber
type Server struct {
	Option
	__obf_49c5d33eaffa7948 *sync.WaitGroup
	__obf_d904e86ddc31620b *sync.RWMutex
	__obf_5c354be353bd1b50 *sync.RWMutex
	__obf_73097b41cd7fdc67 __obf_73097b41cd7fdc67
	__obf_be388f3407943bfe map[string]Handle
}

func NewPubsub(__obf_8c010465edd03aaf Option) *Server {
	return &Server{
		Option: __obf_8c010465edd03aaf, __obf_49c5d33eaffa7948: &sync.WaitGroup{}, __obf_d904e86ddc31620b: &sync.RWMutex{}, __obf_5c354be353bd1b50: &sync.RWMutex{}, __obf_73097b41cd7fdc67: make(__obf_73097b41cd7fdc67),
	}
}

func (__obf_2f76acf7cd574a61 *Server) SetActionHandle(__obf_23c9418bbe890ab4 string, __obf_9676dfb154f2369a Handle) {

	if __obf_2f76acf7cd574a61.__obf_be388f3407943bfe == nil {
		__obf_2f76acf7cd574a61.__obf_be388f3407943bfe = map[string]Handle{__obf_23c9418bbe890ab4: __obf_9676dfb154f2369a}
	} else {
		__obf_2f76acf7cd574a61.__obf_be388f3407943bfe[__obf_23c9418bbe890ab4] = __obf_9676dfb154f2369a
	}
}

func (__obf_2f76acf7cd574a61 *Server) Onlines(__obf_3c3533d430576512 Topic) int {
	return len(__obf_2f76acf7cd574a61.

		// Send simply sends message to the websocket client
		__obf_73097b41cd7fdc67[__obf_3c3533d430576512])
}

func (__obf_2f76acf7cd574a61 *Server) __obf_49220c6602159988(__obf_7dda6c82dad254cd *websocket.Conn, __obf_23c9418bbe890ab4 string, __obf_d29e6ec337e55cb9 any) {
	// send simple message

	// conn.WriteJSON(Message{Action: action, Body: msg})

	// bs, _ := tool.AnyToBytes(Message{Action: action, Body: msg})
	// conn.WriteMessage(websocket.BinaryMessage, bs)
	var __obf_649110fe9c3f58bf []byte
	__obf_649110fe9c3f58bf, _ = util.Encode(Message{Action: __obf_23c9418bbe890ab4, Body: __obf_d29e6ec337e55cb9})
	__obf_2f76acf7cd574a61.__obf_d904e86ddc31620b.
		Lock()
	__obf_7dda6c82dad254cd.
		WriteMessage(websocket.TextMessage, __obf_649110fe9c3f58bf)
	__obf_2f76acf7cd574a61.__obf_d904e86ddc31620b.
		Unlock()
}

func (__obf_2f76acf7cd574a61 *Server) GetConn(__obf_3c3533d430576512 Topic, __obf_ecc3038172b00491 ClientID) *websocket.Conn {
	__obf_8d2a4cec32d37c09, __obf_4c0775848049782a := __obf_2f76acf7cd574a61.__obf_73097b41cd7fdc67[__obf_3c3533d430576512]
	if !__obf_4c0775848049782a {
		return nil
	}

	var __obf_7dda6c82dad254cd *websocket.Conn
	if __obf_7dda6c82dad254cd, __obf_4c0775848049782a = __obf_8d2a4cec32d37c09[__obf_ecc3038172b00491]; !__obf_4c0775848049782a {
		return nil
	}
	return __obf_7dda6c82dad254cd
}

func (__obf_2f76acf7cd574a61 *Server) SendTo(__obf_3c3533d430576512 Topic, __obf_ecc3038172b00491 ClientID, __obf_23c9418bbe890ab4 string, __obf_d29e6ec337e55cb9 any) {
	__obf_7dda6c82dad254cd := __obf_2f76acf7cd574a61.GetConn(__obf_3c3533d430576512, __obf_ecc3038172b00491)
	if __obf_7dda6c82dad254cd != nil {
		__obf_2f76acf7cd574a61.__obf_49220c6602159988(__obf_7dda6c82dad254cd, __obf_23c9418bbe890ab4,

			// SendWithWait sends message to the websocket client using wait group, allowing usage with goroutines
			__obf_d29e6ec337e55cb9)
	}
}

func (__obf_2f76acf7cd574a61 *Server) SendWithWait(__obf_7dda6c82dad254cd *websocket.Conn, __obf_23c9418bbe890ab4 string, __obf_d29e6ec337e55cb9 any) {
	__obf_2f76acf7cd574a61.
		// send simple message
		// conn.WriteMessage(websocket.TextMessage, []byte(tool.EncodeString(msg)))
		__obf_49220c6602159988(__obf_7dda6c82dad254cd,

			// set the task as done
			__obf_23c9418bbe890ab4, __obf_d29e6ec337e55cb9)
	__obf_2f76acf7cd574a61.__obf_49c5d33eaffa7948.
		Done()
}

func (__obf_2f76acf7cd574a61 *Server) SendGreet(__obf_7dda6c82dad254cd *websocket.Conn, __obf_d29e6ec337e55cb9 any) {
	__obf_2f76acf7cd574a61.__obf_49220c6602159988(__obf_7dda6c82dad254cd, __obf_55896f7b03405734, __obf_d29e6ec337e55cb9)
}

func (__obf_2f76acf7cd574a61 *Server) SendWithAction(__obf_7dda6c82dad254cd *websocket.Conn, __obf_23c9418bbe890ab4 string, __obf_d29e6ec337e55cb9 any) {
	__obf_2f76acf7cd574a61.__obf_49220c6602159988(__obf_7dda6c82dad254cd,

		// func (s *Server) SendData(conn *websocket.Conn, msg any) {
		// 	s.send(conn, data, msg)
		// }
		__obf_23c9418bbe890ab4, __obf_d29e6ec337e55cb9)
}

func (__obf_2f76acf7cd574a61 *Server) ThrowError(__obf_7dda6c82dad254cd *websocket.Conn, __obf_7df3d26406dcbc2a string) {
	__obf_2f76acf7cd574a61.__obf_49220c6602159988(__obf_7dda6c82dad254cd,

		// RemoveClient removes the clients from the server subscription map
		__obf_96035d4f1f7782ff, __obf_7df3d26406dcbc2a)
}

func (__obf_2f76acf7cd574a61 *Server) RemoveClient(__obf_ecc3038172b00491 ClientID) {
	__obf_2f76acf7cd574a61.__obf_5c354be353bd1b50.
		Lock()
	defer __obf_2f76acf7cd574a61.
		// loop all topics
		__obf_5c354be353bd1b50.Unlock()

	for __obf_619c1fa4cc0a3d41 := range __obf_2f76acf7cd574a61.
		// delete the client from all the topic's client map
		__obf_73097b41cd7fdc67 {

		delete(__obf_2f76acf7cd574a61.__obf_73097b41cd7fdc67[__obf_619c1fa4cc0a3d41], __obf_ecc3038172b00491)
	}
}

// ProcessMessage handle message according to the action type
func (__obf_2f76acf7cd574a61 *Server) ProcessMessage(__obf_7dda6c82dad254cd *websocket.Conn, __obf_ecc3038172b00491 ClientID, __obf_d29e6ec337e55cb9 []byte) {
	__obf_959e9947043954a4 := // parse message
		Message{}
	if __obf_7df3d26406dcbc2a := json.Unmarshal(__obf_d29e6ec337e55cb9, &__obf_959e9947043954a4); __obf_7df3d26406dcbc2a != nil {
		__obf_2f76acf7cd574a61.
			ThrowError(__obf_7dda6c82dad254cd, ErrInvalidMessage)
		return
	}
	__obf_23c9418bbe890ab4 := // convert all action to lowercase and remove whitespace
		strings.TrimSpace(strings.ToLower(__obf_959e9947043954a4.Action))

	switch __obf_23c9418bbe890ab4 {
	case __obf_3d0cb88ba67024f3:
		__obf_2f76acf7cd574a61.
			Publish(__obf_959e9947043954a4.Topic, __obf_959e9947043954a4.Body)

	case __obf_20d4bfa70f8c734c:
		__obf_2f76acf7cd574a61.
			Subscribe(__obf_7dda6c82dad254cd, __obf_ecc3038172b00491, __obf_959e9947043954a4.Topic)
	case __obf_f6fc30e0b49f4d3d:
		__obf_2f76acf7cd574a61.
			Unsubscribe(__obf_ecc3038172b00491, __obf_959e9947043954a4.Topic)

	default:
		if __obf_cb832dec4a9968b6, __obf_deb0bc5b4f44eeca := __obf_2f76acf7cd574a61.__obf_be388f3407943bfe[__obf_23c9418bbe890ab4]; __obf_deb0bc5b4f44eeca {
			__obf_cb832dec4a9968b6(__obf_ecc3038172b00491, __obf_959e9947043954a4.Topic, __obf_959e9947043954a4.Body.(string))
		} else {
			__obf_2f76acf7cd574a61.
				ThrowError(__obf_7dda6c82dad254cd, ErrActionUnrecognizable)
		}

	}
}

// Publish sends a message to all subscribing clients of a topic
func (__obf_2f76acf7cd574a61 *Server) Publish(__obf_061bc10d62c1be15 Topic, __obf_942b44fd5dc9199c any) {

	// if topic does not exist, stop the process
	if _, __obf_4c0775848049782a := __obf_2f76acf7cd574a61.__obf_73097b41cd7fdc67[__obf_061bc10d62c1be15]; !__obf_4c0775848049782a {
		return
	}
	__obf_8d2a4cec32d37c09 := // if topic exist
		__obf_2f76acf7cd574a61.__obf_73097b41cd7fdc67[__obf_061bc10d62c1be15]

	// send the message to the clients
	for _, __obf_7dda6c82dad254cd := range __obf_8d2a4cec32d37c09 {
		__obf_2f76acf7cd574a61.
			// add 1 job to wait group
			__obf_49c5d33eaffa7948.
			Add(1)

		// send with goroutines
		go __obf_2f76acf7cd574a61.SendWithWait(__obf_7dda6c82dad254cd, __obf_3d0cb88ba67024f3,

			// wait until all goroutines jobs done
			__obf_942b44fd5dc9199c)
	}
	__obf_2f76acf7cd574a61.__obf_49c5d33eaffa7948.
		Wait()
}

// Kickout: kick the older session out
func (__obf_2f76acf7cd574a61 *Server) Kickout(__obf_3c3533d430576512 Topic, __obf_ecc3038172b00491 ClientID) {
	__obf_7dda6c82dad254cd := __obf_2f76acf7cd574a61.GetConn(__obf_3c3533d430576512, __obf_ecc3038172b00491)
	if __obf_7dda6c82dad254cd != nil {
		__obf_2f76acf7cd574a61.__obf_49220c6602159988(__obf_7dda6c82dad254cd, __obf_078850448f7a5f5c, nil)
		__obf_7dda6c82dad254cd.
			Close()
		__obf_2f76acf7cd574a61.__obf_5c354be353bd1b50.
			Lock()
		delete(__obf_2f76acf7cd574a61.__obf_73097b41cd7fdc67[__obf_3c3533d430576512], __obf_ecc3038172b00491)
		__obf_2f76acf7cd574a61.__obf_5c354be353bd1b50.
			Unlock()
	}
}

// Subscribe adds a client to a topic's client map
func (__obf_2f76acf7cd574a61 *Server) Subscribe(__obf_7dda6c82dad254cd *websocket.Conn, __obf_ecc3038172b00491 ClientID, __obf_061bc10d62c1be15 Topic) {
	__obf_2f76acf7cd574a61.__obf_5c354be353bd1b50.
		Lock()
	defer __obf_2f76acf7cd574a61.
		// if topic exist, check the client map
		__obf_5c354be353bd1b50.Unlock()

	if _, __obf_4c0775848049782a := __obf_2f76acf7cd574a61.__obf_73097b41cd7fdc67[__obf_061bc10d62c1be15]; __obf_4c0775848049782a {
		__obf_8d2a4cec32d37c09 := __obf_2f76acf7cd574a61.__obf_73097b41cd7fdc67[__obf_061bc10d62c1be15]
		__obf_8d2a4cec32d37c09[ // if subbed, ok := s.actionHandle[subscribe]; ok {
		// 	subbed(id, topic, "")
		// }

		__obf_ecc3038172b00491] = __obf_7dda6c82dad254cd
		return
	}
	__obf_49ac747a9efe2039 := // if topic does not exist, create a new topic
		make(Client)
	__obf_2f76acf7cd574a61.__obf_73097b41cd7fdc67[__obf_061bc10d62c1be15] = __obf_49ac747a9efe2039

	// add the client to the topic
	__obf_2f76acf7cd574a61.__obf_73097b41cd7fdc67[__obf_061bc10d62c1be15][__obf_ecc3038172b00491] = __obf_7dda6c82dad254cd
}

// Unsubscribe removes a clients from a topic's client map
func (__obf_2f76acf7cd574a61 *Server) Unsubscribe(__obf_ecc3038172b00491 ClientID, __obf_061bc10d62c1be15 Topic) {
	__obf_2f76acf7cd574a61.__obf_5c354be353bd1b50.
		Lock()
	defer __obf_2f76acf7cd574a61.
		// if topic exist, check the client map
		__obf_5c354be353bd1b50.Unlock()

	if _, __obf_4c0775848049782a := __obf_2f76acf7cd574a61.__obf_73097b41cd7fdc67[__obf_061bc10d62c1be15]; __obf_4c0775848049782a {
		__obf_8d2a4cec32d37c09 := __obf_2f76acf7cd574a61.__obf_73097b41cd7fdc67[__obf_061bc10d62c1be15]

		// remove the client from the topic's client map
		delete(__obf_8d2a4cec32d37c09,

			// Pump just do readDump and writePump
			__obf_ecc3038172b00491)
	}
}

func (__obf_2f76acf7cd574a61 *Server) Pump(__obf_ecc3038172b00491 ClientID, __obf_7dda6c82dad254cd *websocket.Conn) {
	__obf_0844912856c85deb := // create channel to signal client health
		make(chan struct{})

	go __obf_2f76acf7cd574a61.__obf_5acbc4430329c4d7(__obf_7dda6c82dad254cd, __obf_ecc3038172b00491, __obf_0844912856c85deb)
	__obf_2f76acf7cd574a61.__obf_d563db09284e7b97(__obf_7dda6c82dad254cd,

		// readPump process incoming messages and set the settings
		__obf_ecc3038172b00491, __obf_0844912856c85deb)
}

func (__obf_2f76acf7cd574a61 *Server) __obf_d563db09284e7b97(__obf_7dda6c82dad254cd *websocket.Conn, __obf_ecc3038172b00491 ClientID, __obf_0844912856c85deb chan<- struct{}) {
	__obf_7dda6c82dad254cd.
		// set limit, deadline to read & pong handler
		SetReadLimit(__obf_2f76acf7cd574a61.MaxMessageSize)
	__obf_7dda6c82dad254cd.
		SetReadDeadline(time.Now().Add(__obf_2f76acf7cd574a61.PongWait * time.Second))
	__obf_7dda6c82dad254cd.
		SetPongHandler(func(string) error {
			__obf_7dda6c82dad254cd.
				SetReadDeadline(time.Now().Add(__obf_2f76acf7cd574a61.PongWait * time.Second))
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
		_, __obf_d29e6ec337e55cb9, __obf_7df3d26406dcbc2a := __obf_7dda6c82dad254cd.ReadMessage()
		// if error occured
		if __obf_7df3d26406dcbc2a != nil {
			__obf_2f76acf7cd574a61.
				// remove from the client
				RemoveClient(__obf_ecc3038172b00491)
			// set health status to unhealthy by closing channel
			close(__obf_0844912856c85deb)
			// stop process
			break
		}
		__obf_2f76acf7cd574a61.

			// if no error, process incoming message
			ProcessMessage(__obf_7dda6c82dad254cd,

				// writePump sends ping to the client
				__obf_ecc3038172b00491, __obf_d29e6ec337e55cb9)
	}
}

func (__obf_2f76acf7cd574a61 *Server) __obf_5acbc4430329c4d7(__obf_7dda6c82dad254cd *websocket.Conn, __obf_ecc3038172b00491 ClientID, __obf_0844912856c85deb <-chan struct{}) {
	__obf_fcedb1756357f8d6 := // create ping ticker
		time.NewTicker(__obf_2f76acf7cd574a61.PingPeriod * time.Second)
	defer __obf_fcedb1756357f8d6.Stop()

	for {
		select {
		case <-__obf_fcedb1756357f8d6.C:
			__obf_7df3d26406dcbc2a := // send ping message
				__obf_7dda6c82dad254cd.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(__obf_2f76acf7cd574a61.WriteWait*time.Second))
			// log.Println("-----------------send ping message： ", err)
			if __obf_7df3d26406dcbc2a != nil {
				__obf_2f76acf7cd574a61.
					// log.Println("-----------------send ping err: ", err)
					// if error sending ping, remove this client from the server
					RemoveClient(__obf_ecc3038172b00491)
				// stop sending ping
				return
			}
		case <-__obf_0844912856c85deb:
			// if process is done, stop sending ping
			return
		}
	}
}
