package repository

import (
	"fmt"
	"paramount_school/internal/model"
)

func (p *Postgres) CreateStudent(student *model.Student) error {
	student.IsActive = false
	err := p.DB.Create(student).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) FindStudentByEmail(email string) (*model.Student, error) {
	//var student *model.Student
	student := &model.Student{}
	if err := p.DB.Where("email = ?", email).Preload("Transcripts").Preload("Certificates").
		Preload("Photos").Preload("LocalIds").Preload("Passports").First(&student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (p *Postgres) FindStudentByPhone(phone string) (*model.Student, error) {
	//var student *model.Student
	student := &model.Student{}
	if err := p.DB.Where("phone_number = ?", phone).Preload("Transcripts").Preload("Certificates").
		Preload("Photos").Preload("LocalIds").Preload("Passports").First(&student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (p *Postgres) UpdateStudentProfile(email string, student model.Student) error {
	err := p.DB.Model(model.Student{}).Where("email = ?", email).Updates(&student).Error
	if err != nil {
		fmt.Println("error updating food")
		return err
	}

	return nil
}
