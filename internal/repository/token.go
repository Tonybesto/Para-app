package repository

import (
	"log"
	"paramount_school/internal/model"
	"time"
)

// TokenInBlacklist checks if token is already in the blacklist collection
func (p *Postgres) TokenInBlacklist(token *string) bool {
	tok := &model.Blacklist{}
	if err := p.DB.Where("token = ?", token).First(&tok).Error; err != nil {
		return false
	}

	return true
}

// AddTokenToBlacklist adds used token to blacklist
func (p *Postgres) AddTokenToBlacklist(email string, token string) error {
	blacklisted := model.Blacklist{}
	blacklisted.Token = token
	blacklisted.Email = email
	blacklisted.CreatedAt = time.Now()

	err := p.DB.Create(&blacklisted).Error
	if err != nil {
		log.Println("error in ad token to blacklist")
		return err
	}
	log.Println("token added to blacklist")
	return nil

}
