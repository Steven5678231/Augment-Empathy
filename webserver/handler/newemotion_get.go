package handler

import (
	"WebRTCDemo/webserver/feedback"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmotion(newfeedback *feedback.EmotionsRepo) gin.HandlerFunc{
	return func(c *gin.Context){
		results := newfeedback.GetAll()
		c.JSON(http.StatusOK, results)
	}
}