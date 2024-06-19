package api

import (
	"database/sql"

	db "github.com/eldersoon/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
)



type Server struct {
	query *sql.DB
	store db.Store
	Router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
		Router: gin.Default(),
	}

	server.Router.POST("/account", server.createAccount)
	server.Router.GET("/account/:id", server.getAccount)
	server.Router.GET("/accounts", server.listAccounts)

	return server
}

func (server *Server) Start(address string, conn *sql.DB) error {
	// Add this query for extras queries needed
	server.query = conn
	return server.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}