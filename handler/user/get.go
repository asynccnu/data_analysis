package user

import (
	. "github.com/asynccnu/data_analysis_service_v1/handler"
	"github.com/asynccnu/data_analysis_service_v1/model"
	"github.com/asynccnu/data_analysis_service_v1/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		SendError(c, errno.ErrUserNotFound, nil, err.Error())
		return
	}
	SendResponse(c, nil, user)
}
