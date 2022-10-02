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
	"strconv"
)

// PhoneLoginStudentHandler godoc
// @Summary      Student Login with phone_number and password
// @Description  After a student Login with phone_number and password, a code is sent to the student for the student to use in VerifyPhoneLoginStudentHandler. Also note that the access token gotten from this endpoint will serve as bearer token for VerifyPhoneLoginStudentHandler. Default password is 12345678
// @Tags         Student
// @Accept       json
// @Produce      json
// @Param UserLogin body model.PhoneLoginRequest true "phone_number, password"
// @Success      200  {string}  string "Enter the verification code sent to your phone"
// @Failure      400  {string}  string "bad request"
// @Failure      500  {string}  string "internal server error"
// @Router       /login/phone [post]
func (u *HTTPHandler) PhoneLoginStudentHandler(c *gin.Context) {

	studentLoginRequest := &model.PhoneLoginRequest{}

	err := c.ShouldBindJSON(&studentLoginRequest)
	if err != nil {
		helpers.Response(c, "bad request", 400, nil, []string{"bad request"})
		return
	}

	student, sqlErr := u.Repository.FindStudentByPhone(studentLoginRequest.PhoneNumber)

	if sqlErr != nil {
		helpers.Response(c, "Phone Number does not exists", http.StatusBadRequest, nil,
			[]string{"Bad Request"})
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

	if HashErr := bcrypt.CompareHashAndPassword([]byte(student.PasswordHash), []byte(studentLoginRequest.Password)); HashErr != nil {
		helpers.Response(c, "invalid Password", http.StatusBadRequest, nil, []string{"Bad Request"})
		return
	}
	verifyCode := helpers.GenerateVerificationCode()

	student.VerificationCode = verifyCode
	err = u.Repository.UpdateStudentProfile(student.Email, *student)
	if err != nil {
		helpers.Response(c, "Internal ServerError", http.StatusInternalServerError, nil, []string{"Error updating student profile"})
		return
	}

	msg := fmt.Sprintf("Here is your verification code %d", verifyCode)
	_, sendErr := u.SMS.SendSms(msg, student.PhoneNumber)
	if sendErr != nil {
		helpers.Response(c, "unable to send code", 500, nil,
			[]string{"unable to send code"})
		return
	}

	log.Printf("verification code is %v\n", verifyCode)

	// Generates access claims and refresh claims
	accessClaims, refreshClaims := middleware.GeneratePhoneNumberClaims(student.PhoneNumber)

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

	helpers.Response(c, "Enter the verification code sent to your phone", http.StatusOK, gin.H{
		"access_token":  *accToken,
		"refresh_token": *refreshToken,
	}, nil)

}

// VerifyPhoneLoginStudentHandler godoc
// @Summary      Verifies that student logging with phone number has the code to log into the App
// @Description  Uses the code sent while logging in with PhoneLoginStudentHandler to log student in. It uses the access token gotten from PhoneLoginStudentHandler as its bearer token and in return, provides another access token while using any STUDENT AUTHENTICATED ROUTE.
// @Tags         Student
// @Accept       json
// @Produce      json
// @Param PhoneLoginCode body model.PhoneLoginCode true "verification_code"
// @Success      200  {string}  string "login successful"
// @Failure      400  {string}  string "bad request"
// @Failure      500  {string}  string "internal server error"
// @Router       /mobile/phone/verify [post]
func (u *HTTPHandler) VerifyPhoneLoginStudentHandler(c *gin.Context) {
	contextStudent, err := u.GetStudentFromContext(c)
	if err != nil {
		helpers.Response(c, "Unauthorized", http.StatusUnauthorized, nil, []string{"unauthorized"})
		return
	}

	PhoneLoginRequest := &model.PhoneLoginCode{}

	err = c.ShouldBindJSON(&PhoneLoginRequest)
	if err != nil {
		helpers.Response(c, "bad request", 400, nil, []string{"bad request"})
		return
	}
	verificationCodeInt, changeErr := strconv.Atoi(PhoneLoginRequest.VerificationCode)
	if changeErr != nil {
		helpers.Response(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	if !helpers.CompareCode(contextStudent.VerificationCode, verificationCodeInt) {
		helpers.Response(c, "incorrect code", 400, nil, []string{"incorrect code"})
		return
	}

	// Generates access claims and refresh claims
	accessClaims, refreshClaims := middleware.GenerateClaims(contextStudent.Email)

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
		"user":          contextStudent,
		"access_token":  *accToken,
		"refresh_token": *refreshToken,
	}, nil)

}
