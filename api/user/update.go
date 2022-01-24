package user

import (
	"strconv"

	. "github.com/GopherReady/ApiBackEnd/api/response"
	"github.com/GopherReady/ApiBackEnd/global"
	"github.com/GopherReady/ApiBackEnd/model"
	"github.com/GopherReady/ApiBackEnd/pkg/errno"
	"github.com/gin-gonic/gin"
)

// Update update a exist user account info.
func Update(c *gin.Context) {
	global.Logger.Infof("User Create function called.X-Request-Id %s ", GetReqID(c))
	// Get the user id from the url parameter.
	userId, _ := strconv.Atoi(c.Param("id"))

	// Binding the user data.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// We update the record based on the user id.
	u.Id = uint64(userId)

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Save changed fields.
	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
