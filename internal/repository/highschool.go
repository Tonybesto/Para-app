package repository

import "paramount_school/internal/model"

func (p *Postgres) CreateHighSchool(highSchool *model.HighSchool) error {
	err := p.DB.Create(highSchool).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) FindHighSchools() ([]model.HighSchool, error) {
	//var student *model.Student
	var highSchools []model.HighSchool
	if err := p.DB.Find(&highSchools).Error; err != nil {
		return nil, err
	}
	return highSchools, nil
}
