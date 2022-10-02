package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"paramount_school/internal/helpers"
	"paramount_school/internal/middleware"
	"paramount_school/internal/model"
)

// LoginStudentHandler godoc
// @Summary      Student Login with email and password
// @Description  Allows student access to use the app after logging in with email and password. Logging in provides an access token that will be used as bearer token while using any STUDENT AUTHENTICATED ROUTE. Default password is 12345678
// @Tags         Student
// @Accept       json
// @Produce      json
// @Param UserLogin body model.UserLogin true "email, password"
// @Success      200  {string}  string "login successful"
// @Failure      400  {string}  string "bad request"
// @Failure      500  {string}  string "internal server error"
// @Router       /student_login [post]
func (u *HTTPHandler) LoginStudentHandler(c *gin.Context) {
	student := &model.Student{}
	studentLoginRequest := &model.UserLogin{}

	err := c.ShouldBindJSON(&studentLoginRequest)
	if err != nil {
		helpers.Response(c, "bad request", 400, nil, []string{"bad request"})
		return
	}

	student, sqlErr := u.Repository.FindStudentByEmail(studentLoginRequest.Email)

	if sqlErr != nil {
		fmt.Println(sqlErr)
		helpers.Response(c, "bad request", http.StatusBadRequest, nil, []string{"email does not exists"})
		return
	}

	//if !student.IsActive {
	//	helpers.Response(c, "please activate your account", http.StatusBadRequest, nil, []string{"Bad Request"})
	//	return
	//}
	//
	//if student.IsBlock {
	//	helpers.Response(c, "you have been blocked", http.StatusBadRequest, nil, []string{"Bad Request"})
	//	return
	//}

	if err := bcrypt.CompareHashAndPassword([]byte(student.PasswordHash), []byte(studentLoginRequest.Password)); err != nil {
		helpers.Response(c, "invalid Password", http.StatusBadRequest, nil, []string{"Bad Request"})
		return
	}

	// Generates access claims and refresh claims
	accessClaims, refreshClaims := middleware.GenerateClaims(student.Email)

	secret := os.Getenv("JWT_SECRET")
	accToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
	if err != nil {
		log.Printf("token generation error err: %v\n", err)
		helpers.Response(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}

	refreshToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, refreshClaims, &secret)
	if err != nil {
		log.Printf("token generation error err: %v\n", err)
		helpers.Response(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}
	c.Header("refresh_token", *refreshToken)
	c.Header("access_token", *accToken)

	helpers.Response(c, "login successful", http.StatusOK, gin.H{
		"user":          student,
		"access_token":  *accToken,
		"refresh_token": *refreshToken,
	}, nil)

}
