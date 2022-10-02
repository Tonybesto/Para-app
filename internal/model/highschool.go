package model

type HighSchool struct {
	Model
	School  string `json:"school"`
	Country string `json:"country"`
	City    string `json:"city"`
	Address string `json:"address"`
	Postal  string `json:"postal"`
}

type Language struct {
	Language string `json:"language"`
}

type Transcript struct {
	StudentId uint   `json:"student_Id"`
	Url       string `json:"url"`
}

type Certificate struct {
	StudentId uint   `json:"student_Id"`
	Url       string `json:"url"`
}

type Photo struct {
	StudentId uint   `json:"student_Id"`
	Url       string `json:"url"`
}

type LocalId struct {
	StudentId uint   `json:"student_Id"`
	Url       string `json:"url"`
}

type Passport struct {
	StudentId uint   `json:"student_Id"`
	Url       string `json:"url"`
}
