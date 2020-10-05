package route

import (
	"WebRTCDemo/webserver/feedback"
	"WebRTCDemo/webserver/handler"

	"github.com/gin-gonic/gin"
)

// Router
func Router() *gin.Engine {
	//init
	emotions := feedback.New()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// router.Use(handler.TLSHandler())
	router.Static("/static/", "./static")

	//default index
	router.GET("/", handler.DefaultHomePageHandler)
	router.GET("/newEmotion", handler.GetEmotion(emotions))
	router.POST("/newEmotion", handler.AddEmotion(emotions))

	//handle socketio request
	router.GET("/socket.io/", handler.SocketIOServerHandler)
	router.POST("/socket.io/", handler.SocketIOServerHandler)
	router.Handle("WS", "/socket.io", handler.SocketIOServerHandler)
	router.Handle("WSS", "/socket.io", handler.SocketIOServerHandler)

	return router
}
