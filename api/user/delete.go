package user

import (
	"strconv"

	. "github.com/GopherReady/ApiBackEnd/api/response"
	"github.com/GopherReady/ApiBackEnd/model"
	"github.com/GopherReady/ApiBackEnd/pkg/errno"
	"github.com/gin-gonic/gin"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	if c.Request.Method != "DELETE" {
		SendResponse(c, errno.ErrorUseDeleteMethod, nil)
		return
	}
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
