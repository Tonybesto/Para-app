package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"paramount_school/internal/helpers"
	"paramount_school/internal/middleware"
	"paramount_school/internal/model"
	"path/filepath"
	"strings"
)

// SignUpHandler godoc
// @Summary      Student registers to use the app
// @Description  registers/creates student in the database by collecting form data in model.Student. Note: "files"  are to be uploaded in jpeg, png or pdf.
// @Tags         Student
// @Accept       json
// @Produce      json
// @Param Student body model.Student true "every field can be filled or omitted as form data as decided from the frontend"
// @Success      201  {string} string "account created successfully"
// @Failure      500  {string}  string "internal server error"
// @Failure      400  {string}  string "bad request"
// @Router       /register [post]
func (u *HTTPHandler) SignUpHandler(c *gin.Context) {
	form, err := c.MultipartForm()

	if err != nil {
		log.Printf("error parsing multipart form: %v", err)
		helpers.Response(c, "error parsing multipart form", 400, nil, []string{"bad request"})
		return
	}

	formTranscript := form.File["transcript"]
	var transcript []model.Transcript

	// upload the transcript to aws.
	for _, f := range formTranscript {
		file, err := f.Open()
		if err != nil {

		}
		fileExtension, ok := middleware.CheckSupportedFile(strings.ToLower(f.Filename))
		log.Printf(filepath.Ext(strings.ToLower(f.Filename)))
		fmt.Println(fileExtension)
		if ok {
			log.Println(fileExtension)
			helpers.Response(c, "Bad Request", 400, nil, []string{fileExtension + " image file type is not supported"})
			return
		}

		session, tempFileName, err := middleware.PreAWS(fileExtension, "product")
		if err != nil {
			log.Println("could not upload file", err)
		}

		url, err := u.AWS.UploadFileToS3(session, file, tempFileName, f.Size)
		if err != nil {
			helpers.Response(c, "internal server error", http.StatusInternalServerError, nil, []string{"an error occurred while uploading the image"})
			return
		}

		trans := model.Transcript{
			Url: url,
		}
		transcript = append(transcript, trans)
	}

	formCertificate := form.File["certificate"]
	var certificate []model.Certificate

	// upload the certificate to aws.
	for _, f := range formCertificate {
		file, err := f.Open()
		if err != nil {

		}
		fileExtension, ok := middleware.CheckSupportedFile(strings.ToLower(f.Filename))
		log.Printf(filepath.Ext(strings.ToLower(f.Filename)))
		fmt.Println(fileExtension)
		if ok {
			log.Println(fileExtension)
			helpers.Response(c, "Bad Request", 400, nil, []string{fileExtension + " image file type is not supported"})
			return
		}

		session, tempFileName, err := middleware.PreAWS(fileExtension, "product")
		if err != nil {
			log.Println("could not upload file", err)
		}

		url, err := u.AWS.UploadFileToS3(session, file, tempFileName, f.Size)
		if err != nil {
			helpers.Response(c, "internal server error", http.StatusInternalServerError, nil, []string{"an error occurred while uploading the image"})
			return
		}

		cert := model.Certificate{
			Url: url,
		}
		certificate = append(certificate, cert)
	}

	formPhoto := form.File["photo"]
	var photo []model.Photo

	// upload the photo to aws.
	for _, f := range formPhoto {
		file, err := f.Open()
		if err != nil {

		}
		fileExtension, ok := middleware.CheckSupportedFile(strings.ToLower(f.Filename))
		log.Printf(filepath.Ext(strings.ToLower(f.Filename)))
		fmt.Println(fileExtension)
		if ok {
			log.Println(fileExtension)
			helpers.Response(c, "Bad Request", 400, nil, []string{fileExtension + " image file type is not supported"})
			return
		}

		session, tempFileName, err := middleware.PreAWS(fileExtension, "product")
		if err != nil {
			log.Println("could not upload file", err)
		}

		url, err := u.AWS.UploadFileToS3(session, file, tempFileName, f.Size)
		if err != nil {
			helpers.Response(c, "internal server error", http.StatusInternalServerError, nil, []string{"an error occurred while uploading the image"})
			return
		}

		pic := model.Photo{
			Url: url,
		}
		photo = append(photo, pic)
	}

	formLocalID := form.File["id"]
	var ID []model.LocalId

	// upload the ID to aws.
	for _, f := range formLocalID {
		file, err := f.Open()
		if err != nil {

		}
		fileExtension, ok := middleware.CheckSupportedFile(strings.ToLower(f.Filename))
		log.Printf(filepath.Ext(strings.ToLower(f.Filename)))
		fmt.Println(fileExtension)
		if ok {
			log.Println(fileExtension)
			helpers.Response(c, "Bad Request", 400, nil, []string{fileExtension + " image file type is not supported"})
			return
		}

		session, tempFileName, err := middleware.PreAWS(fileExtension, "product")
		if err != nil {
			log.Println("could not upload file", err)
		}

		url, err := u.AWS.UploadFileToS3(session, file, tempFileName, f.Size)
		if err != nil {
			helpers.Response(c, "internal server error", http.StatusInternalServerError, nil, []string{"an error occurred while uploading the image"})
			return
		}

		id := model.LocalId{
			Url: url,
		}
		ID = append(ID, id)
	}

	formPassport := form.File["passport"]
	var passport []model.Passport

	// upload the passport to aws.
	for _, f := range formPassport {
		file, err := f.Open()
		if err != nil {

		}
		fileExtension, ok := middleware.CheckSupportedFile(strings.ToLower(f.Filename))
		log.Printf(filepath.Ext(strings.ToLower(f.Filename)))
		fmt.Println(fileExtension)
		if ok {
			log.Println(fileExtension)
			helpers.Response(c, "Bad Request", 400, nil, []string{fileExtension + " image file type is not supported"})
			return
		}

		session, tempFileName, err := middleware.PreAWS(fileExtension, "product")
		if err != nil {
			log.Println("could not upload file", err)
		}

		url, err := u.AWS.UploadFileToS3(session, file, tempFileName, f.Size)
		if err != nil {
			helpers.Response(c, "internal server error", http.StatusInternalServerError, nil, []string{"an error occurred while uploading the image"})
			return
		}

		pass := model.Passport{
			Url: url,
		}
		passport = append(passport, pass)
	}

	studentProfile := model.StudentProfile{
		FullName:    strings.TrimSpace(c.PostForm("full_name")),
		FatherName:  strings.TrimSpace(c.PostForm("father_name")),
		MotherName:  strings.TrimSpace(c.PostForm("mother_name")),
		DateOfBirth: strings.TrimSpace(c.PostForm("date_of_birth")),
		Gender:      strings.TrimSpace(c.PostForm("gender")),
		Email:       strings.TrimSpace(c.PostForm("email")),
	}

	studentContact := model.StudentContact{
		PhoneNumber:        strings.TrimSpace(c.PostForm("phone_number")),
		Citizenship:        strings.TrimSpace(c.PostForm("citizenship")),
		CountryOfResidence: strings.TrimSpace(c.PostForm("residence")),
		Region:             strings.TrimSpace(c.PostForm("region")),
		PostalCode:         strings.TrimSpace(c.PostForm("postal_code")),
		PassportNumber:     strings.TrimSpace(c.PostForm("passport_number")),
		IssueDate:          strings.TrimSpace(c.PostForm("issue_date")),
		ExpiryDate:         strings.TrimSpace(c.PostForm("expiry_date")),
		IdNumber:           strings.TrimSpace(c.PostForm("id_number")),
	}

	studentSchool := model.StudentHighSchool{
		HighSchoolCountry: strings.TrimSpace(c.PostForm("school_country")),
		City:              strings.TrimSpace(c.PostForm("school_city")),
		HighSchool: model.HighSchool{
			School: strings.TrimSpace(c.PostForm("school")),
		},
		GraduationDate: strings.TrimSpace(c.PostForm("graduation_date")),
		StudentID:      strings.TrimSpace(c.PostForm("student_id")),
		Majors:         strings.TrimSpace(c.PostForm("majors")),
		Language: model.Language{
			Language: strings.TrimSpace(c.PostForm("language")),
		},
	}

	student := model.Student{
		StudentProfile:    studentProfile,
		StudentContact:    studentContact,
		StudentHighSchool: studentSchool,
		Transcripts:       transcript,
		Certificates:      certificate,
		Photos:            photo,
		LocalIds:          ID,
		Passports:         passport,
	}

	validEmail := student.ValidateEmail()
	if !validEmail {
		helpers.Response(c, "bad request", 400, nil, []string{"input valid email"})
		return
	}

	_, emailErr := u.Repository.FindStudentByEmail(student.Email)
	if emailErr == nil {
		helpers.Response(c, "Email already exists", 400, nil, []string{"email exists"})
		return
	}
	_, phoneErr := u.Repository.FindStudentByPhone(student.PhoneNumber)
	if phoneErr == nil {
		helpers.Response(c, "phone number already exists", 400, nil, []string{"phone number exists"})
		return
	}

	student.Password = "12345678"

	if hashErr := student.HashPassword(); hashErr != nil {
		helpers.Response(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	log.Println(student.HighSchool)

	createErr := u.Repository.CreateStudent(&student)
	if createErr != nil {
		helpers.Response(c, "Unable to create user", 500, nil, []string{"unable to create user"})
		return
	}

	helpers.Response(c, "account created successfully", 201, nil, nil)
}
