package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paramount_school/internal/helpers"
	"paramount_school/internal/model"
)

// StudentUpdateHandler godoc
// @Summary      This is used to update student's profile
// @Description  After registering and logging in with a default password of 12345678. Students can change their password (or any field from model.Student as provided by the frontend). It is a STUDENT AUTHENTICATED ROUTE. Therefore, it uses the access token gotten from LoginStudentHandler and VerifyPhoneLoginStudentHandler as bearer token.
// @Tags         Student
// @Accept       json
// @Produce      json
// @Param Student body model.Student true "password"
// @Success      201  {string} string "Successfully Updated"
// @Failure      500  {string}  string "internal server error"
// @Failure      400  {string}  string "bad request"
// @Router       /update_profile [put]
func (u *HTTPHandler) StudentUpdateHandler(c *gin.Context) {

	contextStudent, err := u.GetStudentFromContext(c)
	if err != nil {
		helpers.Response(c, "Unauthorized", http.StatusUnauthorized, nil, []string{"unauthorized"})
		return
	}

	var student model.Student
	err = c.ShouldBindJSON(&student)
	if err != nil {
		helpers.Response(c, "bad request", 400, nil, []string{"validation error"})
		return
	}

	validPassword := student.IsValid(student.Password)
	if !validPassword {
		helpers.Response(c, "bad request", 400, nil, []string{"password must have upper, lower case, number, special character and length not less than 8 characters"})
		return
	}

	if hashErr := student.HashPassword(); hashErr != nil {
		helpers.Response(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	err = u.Repository.UpdateStudentProfile(contextStudent.Email, student)
	if err != nil {
		helpers.Response(c, "Internal ServerError", http.StatusInternalServerError, nil, []string{"Error updating student profile"})
		return
	}

	helpers.Response(c, "Successfully Updated", 200, student, nil)

}
