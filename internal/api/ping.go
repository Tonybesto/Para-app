package api

import (
	"github.com/gin-gonic/gin"
	"paramount_school/internal/helpers"
	"paramount_school/internal/model"
)

//PingHandler is for testing the connections
func (u *HTTPHandler) PingHandler(c *gin.Context) {
	data := &model.Student{}

	// healthcheck
	helpers.Response(c, "pong", 200, data, nil)
}
