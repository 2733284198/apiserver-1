package user

import (
	h "apiserver/internal/app/handler"
	model "apiserver/internal/app/model/db"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
)

// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /v1/user [post]
func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		h.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	//// Validate the data.
	//if err := u.Validate(); err != nil {
	//	h.SendResponse(c, errno.ErrValidation, nil)
	//	return
	//}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		h.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		h.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	h.SendResponse(c, nil, rsp)
}
