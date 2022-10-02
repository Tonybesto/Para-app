package model

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"unicode"
)

type Student struct {
	Model
	StudentProfile
	StudentContact
	StudentHighSchool
	Transcripts      []Transcript  `json:"transcripts" gorm:"oneToMany"`
	Certificates     []Certificate `json:"certificates" gorm:"oneToMany"`
	Photos           []Photo       `json:"photos" gorm:"oneToMany"`
	LocalIds         []LocalId     `json:"local_ids" gorm:"oneToMany"`
	Passports        []Passport    `json:"passport" gorm:"oneToMany"`
	IsActive         bool          `json:"is_active"`
	Status           string        `json:"status"`
	Token            string        `json:"token"`
	IsBlock          bool          `json:"is_block"`
	VerificationCode int           `json:"verification_code"`
}

type StudentProfile struct {
	FullName     string `json:"full_name"`
	FatherName   string `json:"father_name"`
	MotherName   string `json:"mother_name"`
	DateOfBirth  string `json:"date_of_birth"`
	Gender       string `json:"gender"`
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"password,omitempty" gorm:"-"`
	PasswordHash string `json:"password_hash"`
}

type StudentContact struct {
	PhoneNumber        string `json:"phone_number" gorm:"unique"`
	Citizenship        string `json:"citizenship"`
	CountryOfResidence string `json:"country_of_residence"`
	Region             string `json:"region"`
	PostalCode         string `json:"postal_code"`
	PassportNumber     string `json:"passport_number"`
	IssueDate          string `json:"issue_date"`
	ExpiryDate         string `json:"expiry_date"`
	IdNumber           string `json:"id_number"`
}

type StudentHighSchool struct {
	HighSchoolCountry string `json:"high_school_country"`
	City              string `json:"city"`
	HighSchool        HighSchool
	GraduationDate    string `json:"graduation_date"`
	StudentID         string `json:"student_id"`
	Majors            string `json:"majors"`
	Language          Language
}

func (s *Student) ValidateEmail() bool {
	emailRegexp := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegexp.MatchString(s.Email)

}

func (s *Student) IsValid(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(password) >= 8 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func (s *Student) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	s.PasswordHash = string(hashedPassword)
	return nil
}
