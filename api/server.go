package api

import (
	"database/sql"

	db "github.com/eldersoon/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
)



type Server struct {
	query *sql.DB
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

func (server *Server) Start(address string, conn *sql.DB) error {
	server.query = conn
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}