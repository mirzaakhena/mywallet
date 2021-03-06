package userapi

import (
	"mywallet/application/apperror"
	"mywallet/infrastructure/log"
	"mywallet/infrastructure/util"
	"mywallet/usecase/registeruser"
	"net/http"

	"github.com/gin-gonic/gin"
)

// registerUserHandler ...
func (r *Controller) registerUserHandler(inputPort registeruser.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.Context(c.Request.Context())

		var req registeruser.InportRequest
		if err := c.BindJSON(&req); err != nil {
			newErr := apperror.FailUnmarshalResponseBodyError
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(newErr))
			return
		}

		log.Info(ctx, util.MustJSON(req))

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
