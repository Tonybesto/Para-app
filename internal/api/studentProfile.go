package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paramount_school/internal/helpers"
)

// StudentProfileHandler godoc
// @Summary      Displays all the information of the student available in the database
// @Description  This displays whatever information of the student captured during registration. It is a STUDENT AUTHENTICATED ROUTE. Therefore, it uses the access token gotten from LoginStudentHandler and VerifyPhoneLoginStudentHandler as bearer token.
// @Tags         Student
// @Accept       json
// @Produce      json
// @Success      200  {string} string "found successfully"
// @Failure      500  {string}  string "internal server error"
// @Router       /get_profile [get]
func (u *HTTPHandler) StudentProfileHandler(c *gin.Context) {
	contextStudent, err := u.GetStudentFromContext(c)
	if err != nil {
		helpers.Response(c, "Unauthorized", http.StatusUnauthorized, nil, []string{"unauthorized"})
		return
	}

	student, err := u.Repository.FindStudentByEmail(contextStudent.Email)
	if err != nil {
		helpers.Response(c, "Internal Server Error", http.StatusInternalServerError, nil, []string{"Could not get student"})
		return
	}

	helpers.Response(c, "found successfully", 200, student, nil)
}
