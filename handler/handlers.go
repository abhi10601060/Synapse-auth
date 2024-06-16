package handler

import (
	"log"
	"net/http"
	"synapse/auth/db"
	"synapse/auth/model"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message" : "Synapse Auth is Alive",
	})
}

func SignUp(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil{
		log.Println("error in binding user is : ", err)
		c.JSON(400, gin.H{
			"error" : err.Error(),
		})
		c.Abort()
	}
	log.Println("received user is : ", user)
	db.AddUser(&user)
}

func Login(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil{
		log.Println("error in binding user is : ", err)
		c.JSON(400, gin.H{
			"error" : err.Error(),
		})
		c.Abort()
	}
	log.Println("received user is : ", user)
}