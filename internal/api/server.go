package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tcscheurer/rentals/db/sqlc"
)

type Server struct {
	querier sqlc.Querier
	router *gin.Engine
}

func NewServer(q sqlc.Querier) *Server {
	s := Server{
		querier: q,
	}
	r := gin.Default()

	r.GET("/rentals", s.GetRentals)
	r.GET("/rentals/:id", s.GetRental)


	s.router = r
	return &s
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}