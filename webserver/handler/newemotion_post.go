package handler

import (
	"WebRTCDemo/webserver/feedback"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addEmotionPostRequest struct {
	UserID  string      `json:"userID"`
	RoomID  string      `json:"roomID"`
	Type 	string 		`json:"type"`
	Emotion []float32	`json:"value"`

}

func AddEmotion(newfeedback *feedback.EmotionsRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := addEmotionPostRequest{}
		c.Bind(&requestBody)

		newemotion := feedback.EmotionMsg{
			UserID: requestBody.UserID,
			RoomID: requestBody.RoomID,
			Type: requestBody.Type,
			Emotion: requestBody.Emotion,

		}

		var roomid = newemotion.RoomID
		server.BroadcastToRoom("/",roomid,"newFeedback",newemotion.UserID,newemotion)
		newfeedback.Add(newemotion)

		c.Status(http.StatusNoContent)
	}
}
