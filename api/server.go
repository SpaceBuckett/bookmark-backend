package api

import (
	db "github.com/SpaceBuckett/bookmark-backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	router.POST("/userprofiles", server.createUserProfile)
	router.GET("/userprofiles", server.getAllProfiles)
	router.GET("/userprofiles/:id", server.getUserProfile)
	router.GET("/users/:owner_id/bookmarks", server.getBookmarksByUser)

	router.POST("/bookmarks", server.createBookmark)
	router.GET("/bookmarks/:id", server.getBookmark)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
