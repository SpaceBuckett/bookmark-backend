package api

import (
	"fmt"
	db "github.com/SpaceBuckett/bookmark-backend/db/sqlc"
	"github.com/SpaceBuckett/bookmark-backend/token"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store      *db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker("12345678901234567890123456789012")
	if err != nil {
		return nil, fmt.Errorf("cannot create token %w", err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	router.POST("/userprofiles", server.createUserProfile)
	router.GET("/userprofiles", server.getAllProfiles)
	router.GET("/userprofiles/:id", server.getUserProfile)
	router.GET("/users/:owner_id/bookmarks", server.getBookmarksByUser)

	router.POST("/bookmarks", server.createBookmark)
	router.GET("/bookmarks/:id", server.getBookmark)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
