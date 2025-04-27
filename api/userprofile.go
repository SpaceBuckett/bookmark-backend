package api

import (
	db "github.com/SpaceBuckett/bookmark-backend/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createUserProfileRequest struct {
	Username       string `json:"username" binding:"required,alphanum"`
	Email          string `json:"email" binding:"required,email"`
	HashedPassword string `json:"hashed_password" binding:"required"`
}

func (server *Server) createUserProfile(c *gin.Context) {
	var req createUserProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserAccountParams{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: req.HashedPassword,
	}

	userAccount, err := server.store.CreateUserAccount(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, userAccount)
}

type listUserProfilesRequest struct {
	PageId   int64 `form:"page_id" binding:"required,min=1"`
	PageSize int64 `form:"page_size" binding:"required,min=1,max=10"`
}

func (server *Server) getAllProfiles(c *gin.Context) {
	var req listUserProfilesRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAllUserAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}

	userProfiles, err := server.store.GetAllUserAccounts(c, arg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, userProfiles)
}

type getUserProfileRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUserProfile(c *gin.Context) {
	var req getUserProfileRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userProfile, err := server.store.GetUserAccount(c, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, userProfile)

}
