package userapi

import (
	"mywallet/infrastructure/log"
	"mywallet/infrastructure/util"
	"mywallet/usecase/showalluser"
	"net/http"

	"github.com/gin-gonic/gin"
)

// showAllUserHandler ...
func (r *Controller) showAllUserHandler(inputPort showalluser.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.Context(c.Request.Context())

		var req showalluser.InportRequest

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(err))
			return
		}

		log.Info(ctx, util.MustJSON(res))
		c.JSON(http.StatusOK, NewSuccessResponse(res))

	}
}
