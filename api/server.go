package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gyu-young-park/go_blog/db"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}
	server.setUpRouter()
	return server
}

func (server *Server) setUpRouter() {
	router := gin.Default()
	server.router = router
	router.GET("/users", server.getAllUserData)
	router.GET("/user/:id", server.getUserData)
	router.POST("/user", server.registerUser)
	router.POST("/delete/user", server.deleteUser)
	router.POST("/update/user", server.updateUserInfo)
}

func (server *Server) StartServer(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func sendErrorMessage(status int, message string, err error, c *gin.Context) {
	log.Printf("failed to [%v] [%v]\n", message, err.Error())
	c.JSON(status, errResponse(err))
}
