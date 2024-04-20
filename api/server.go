package api

import (
	"github.com/gin-gonic/gin"
	"sqlcs.sqlc.dev/app/sqlcs"
)

type Server struct {
	store  *sqlcs.Store
	router *gin.Engine
}

func NewServer(store *sqlcs.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.POST("/transfers", server.createTransfer)
	router.POST("/signup",server.CreateUser)
	router.POST("/signin",server.getUser)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
