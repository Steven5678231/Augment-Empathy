package handler

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

var (
	server *socketio.Server
	err    error
)

const (
	MaxUserCnt = 2
)

type RoomInfo struct {
	Pair   map[string]string
	Roomid uint32
}

var dict map[string]string

type Msg struct {
	UserID    string   `json:"userID"`
	Text      string   `json:"text"`
	State     string   `json:"state"`
	Namespace string   `json:"namespace"`
	Rooms     []string `json:"rooms"`
}

type EmotionMsg struct {
	UserID  string    `json:"userID"`
	RoomID  string    `json:"roomID"`
	Type    string    `json:"typeOfFeedback"`
	Emotion []float32 `json:"value"`
}

func init() {
	server, err = socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(so socketio.Conn) error {
		log.Println("on connection, ID: ", so.ID())

		so.SetContext("")
		msg := Msg{so.ID(), "Connected", "notice", "", nil}
		so.Emit("res", msg)

		return nil
	})

	server.OnEvent("/", "join", func(so socketio.Conn, room string) {
		if server.RoomLen(so.Namespace(), room) >= MaxUserCnt {
			//Room is full
			so.Emit("full", room)
			return
		}

		so.Join(room)
		log.Println(so.ID(), " join ", room, so.Rooms())
		//broadcast to everyone
		server.BroadcastToRoom(so.Namespace(), room, "joined", room, so.ID())
	})

	//When someone leave
	server.OnEvent("/", "leave", func(so socketio.Conn, room string) {
		log.Println(so.ID(), " leave ", room, so.Namespace(), so.Rooms())
		server.BroadcastToRoom(so.Namespace(), room, "leaved", room, so.ID())

		so.Leave(room)
	})

	//When it's time to pop-up form
	server.OnEvent("/", "startPopup", func(so socketio.Conn, room string) {
		log.Println(so.ID(), " Have a Questionaire ", room, so.Namespace(), so.Rooms())
		server.BroadcastToRoom(so.Namespace(), room, "startPopup", room, so.ID())

	})
	//When it's time to stop pop-up form
	server.OnEvent("/", "stopPopup", func(so socketio.Conn, room string) {
		log.Println(so.ID(), " chat finished ", room, so.Namespace(), so.Rooms())
		server.BroadcastToRoom(so.Namespace(), room, "stopPopup", room, so.ID())
	})

	server.OnEvent("/", "message", func(so socketio.Conn, room string, msg interface{}) {
		//Forward to the other directly
		server.BroadcastToRoom(so.Namespace(), room, "message", room, so.ID(), msg)
	})

	server.OnEvent("/", "ready", func(so socketio.Conn, room string) {
		//Forward to the other directly
		server.BroadcastToRoom(so.Namespace(), room, "ready", room, so.ID())
	})

	server.OnEvent("/", "chat", func(so socketio.Conn, msg string) {
		res := Msg{so.ID(), "----" + msg, "normal", so.Namespace(), so.Rooms()}
		so.SetContext(res)
		log.Println("chat receive", msg, so.Namespace(), so.Rooms(), server.Rooms(so.Namespace()))
		rooms := so.Rooms()

		for _, room := range rooms {
			server.BroadcastToRoom(so.Namespace(), room, "res", res)
		}

	})

	go server.Serve()
}

func runDataLoop(newInfo EmotionMsg) {
	// extra info
	for {
		emotion := []float32{0, 0, 0, 1, 0, 0, 0, 0.5}
		feedback_type := "EDA Singal"
		res := EmotionMsg{newInfo.UserID, newInfo.RoomID, feedback_type, emotion}
		server.BroadcastToRoom("/", newInfo.RoomID, "newFeedback", newInfo.UserID, res)
		time.Sleep(1 * time.Second)
	}

}

func SocketIOServerHandler(c *gin.Context) {

	//server.OnEvent("/", "notice")
	if server != nil {
		log.Println("WebSocket server start...")
		server.ServeHTTP(c.Writer, c.Request)

	}
}
