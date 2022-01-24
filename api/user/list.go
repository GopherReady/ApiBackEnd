package user

import (
	. "github.com/GopherReady/ApiBackEnd/api/response"
	"github.com/GopherReady/ApiBackEnd/pkg/errno"
	"github.com/GopherReady/ApiBackEnd/service"
	"github.com/gin-gonic/gin"
)

// List list the users in the database.
func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
