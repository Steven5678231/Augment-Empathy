package handler

import (
	"WebRTCDemo/webserver/feedback"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type addEmotionPostRequest struct {
	UserID  string    `json:"userID"`
	RoomID  string    `json:"roomID"`
	Emotions []feedback.EmotionDetail `json:"emotion_detail"`
}

func AddEmotion(newfeedback *feedback.EmotionsRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := feedback.EmotionMsg{}
		c.BindJSON(&requestBody)


		newemotion := feedback.EmotionMsg{
			UserID:  requestBody.UserID,
			RoomID:  requestBody.RoomID,
			Emotions: requestBody.Emotions,
		}

		fmt.Println(requestBody)
		mapTestData := make(map[string]interface{})
		jsonInfo, _ := json.Marshal(&requestBody)
		_ = json.Unmarshal(jsonInfo, &mapTestData)
		fmt.Println(mapTestData)

		var roomid = newemotion.RoomID
		server.BroadcastToRoom("/", roomid, "newFeedback", newemotion.UserID, newemotion.Emotions)
		// newfeedback.Add(newemotion)

		c.Status(http.StatusNoContent)
	}
}
