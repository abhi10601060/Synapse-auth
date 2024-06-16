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
		"message": "Synapse Auth is Alive",
	})
}

func SignUp(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		log.Println("error in binding user is : ", err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	log.Println("received user is : ", user)

	if db.UserExist(&user) {
		c.JSON(409, gin.H{
			"message": "user exists already",
		})
		c.Abort()
		return
	}

	res := db.AddUser(&user)

	if res == -1 {
		c.JSON(400, gin.H{
			"message": "Internal Server Error",
		})
		c.Abort()
		return
	}

}

func Login(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		log.Println("error in binding user is : ", err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	log.Println("received user is : ", user)

	if !db.IsValidPassword(&user){
		c.JSON(201, gin.H{
			"message" : "Incorrect Credentials",
		})
		c.Abort()
		return
	}

	
}
