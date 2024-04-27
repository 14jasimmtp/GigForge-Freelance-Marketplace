package repository

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pb"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return Repo{db: db}
}

const (
	StatusOK                  = http.StatusOK
	StatusBadRequest          = http.StatusBadRequest
	StatusInternalServerError = http.StatusInternalServerError
)

func (r *Repo) AddSingleProject(project *pb.AddSingleProjectReq) (int, error) {
	var id int
	insertProjectQuery := `INSERT INTO projects(created_at,updated_at,title, description, category) VALUES (?, ?, ?, ?, ?) RETURNING id`
	err := r.db.Raw(insertProjectQuery, time.Now(),time.Now(), project.Title, project.Description, project.Category).Scan(&id).Error
	if err != nil {
		log.Printf("Error inserting project: %v", err)
		return StatusInternalServerError, errors.New("failed to add project")
	}

	insertSingleProjectQuery := `INSERT INTO single_projects(created_at,updated_at,project_id, price, deliver_days, number_of_revisions) VALUES (?, ?, ?, ?, ?, ?)`
	err = r.db.Exec(insertSingleProjectQuery, time.Now(),time.Now(), id, project.Price, project.DeliveryDays, project.NumberOfRevisions).Error
	if err != nil {
		log.Printf("Error inserting single project: %v", err)
		return StatusBadRequest, errors.New("failed to add single project")
	}

	log.Println("Project added successfully")
	return StatusOK, nil
}

// func (r *Repo) AddTierProject(project *pb.AddProjectReq) error{

// }
