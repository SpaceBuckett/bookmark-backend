package api

import (
	db "github.com/SpaceBuckett/bookmark-backend/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createBookmarkRequest struct {
	OwnerID int64  `json:"owner_id" binding:"required,min=1"`
	Title   string `json:"title" binding:"required"`
	Url     string `json:"url" binding:"required"`
}

func (server *Server) createBookmark(c *gin.Context) {
	var req createBookmarkRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateBookParams{
		OwnerID: req.OwnerID,
		Title:   req.Title,
		Url:     req.Url,
	}

	bookmark, err := server.store.CreateBook(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, bookmark)
}

type getBookmarkRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getBookmark(c *gin.Context) {
	var req getBookmarkRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	bookmark, err := server.store.GetBookMark(c, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, bookmark)
}

type getBookmarksByUserRequest struct {
	OwnerId int64 `uri:"owner_id" binding:"required,min=1"`
}

func (server *Server) getBookmarksByUser(c *gin.Context) {
	var req getBookmarksByUserRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	bookmarks, err := server.store.GetBookmarksByUser(c, req.OwnerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, bookmarks)
}
