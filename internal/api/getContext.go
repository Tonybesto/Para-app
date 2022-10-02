package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"paramount_school/internal/model"
)

func (u *HTTPHandler) GetStudentFromContext(c *gin.Context) (*model.Student, error) {
	contextUser, exists := c.Get("user")
	if !exists {
		return nil, fmt.Errorf("error getting user from context")
	}
	user, ok := contextUser.(*model.Student)
	if !ok {
		return nil, fmt.Errorf("an error occurred")
	}
	return user, nil
}

func (u *HTTPHandler) GetMobileStudentLoginFromContext(c *gin.Context) (*model.Student, error) {
	contextUser, exists := c.Get("student")
	if !exists {
		return nil, fmt.Errorf("error getting user from context")
	}
	user, ok := contextUser.(*model.Student)
	if !ok {
		return nil, fmt.Errorf("an error occurred")
	}
	return user, nil
}
