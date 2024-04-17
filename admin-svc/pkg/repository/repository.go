package repository

import (
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/admin-svc/pkg/domain"
	"gorm.io/gorm"
)

type AdminRepo struct {
	db *gorm.DB
}

type RepoIfc interface {
	AdminLogin(email string) (*domain.Admin, error)
}

func NewAdminRepo(db *gorm.DB) RepoIfc {
	return &AdminRepo{db: db}
}

func (r *AdminRepo) AdminLogin(email string) (*domain.Admin, error) {
	var details domain.Admin
	if err := r.db.Raw("SELECT * FROM admins WHERE email=?", email).Scan(&details).Error; err != nil {
		return nil, err
	}
	return &details, nil
}
