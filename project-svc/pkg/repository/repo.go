package repository

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pb"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pkg/domain"
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
	insertProjectQuery := `INSERT INTO projects(created_at,updated_at,title, description, category,user_id) VALUES (?, ?, ?, ?, ?,?) RETURNING id`
	err := r.db.Raw(insertProjectQuery, time.Now(), time.Now(), project.Title, project.Description, project.Category,project.UserId).Scan(&id).Error
	if err != nil {
		log.Printf("Error inserting project: %v", err)
		return StatusInternalServerError, errors.New("failed to add project")
	}

	insertSingleProjectQuery := `INSERT INTO single_projects(created_at,updated_at,project_id, price, deliver_days, number_of_revisions) VALUES (?, ?, ?, ?, ?, ?)`
	err = r.db.Exec(insertSingleProjectQuery, time.Now(), time.Now(), id, project.Price, project.DeliveryDays, project.NumberOfRevisions).Error
	if err != nil {
		log.Printf("Error inserting single project: %v", err)
		return StatusBadRequest, errors.New("failed to add single project")
	}

	log.Println("Project added successfully")
	return StatusOK, nil
}

// func (r *Repo) AddTierProject(project *pb.AddProjectReq) error{

// }

func (r *Repo) EditSingleProject(req *pb.EditSingleProjectReq) (int, error) {
	query := `UPDATE projects set updated_at = ? , title = ? , description = ? ,category = ? where id = ? and user_id = ?`
	err := r.db.Raw(query, time.Now(), req.Title, req.Description, req.Category,req.ProjectId,req.UserId).Error
	if err != nil {
		fmt.Println(err)
		return StatusInternalServerError, errors.New("error while updating project")
	}
	query= `UPDATE single_project set updated_at = ?, price = ?, deliver_days = ?, number_of_revisions = ? where project_id = ?`
	errs:= r.db.Exec(query,time.Now(),req.Price,req.DeliveryDays,req.NumberOfRevisions,req.ProjectId).Error
	if errs != nil {
		fmt.Println(err)
		return StatusInternalServerError, errors.New("error while updating project")
	}
	return StatusOK,nil
}

func (r *Repo) DeleteProject(req *pb.RemProjectReq) (error){
	query:=`Delete from projects where id = ? and user_id = ?`
	err:=r.db.Raw(query,req.ProjectId,req.UserId).Error
	if err != nil {
		return errors.New(`can't delete project.Not yours`)
	}
	return nil
}

func (r *Repo) ListProjects() ([]*pb.Project,error){
	var project []domain.Project
	var singleProject domain.SingleProject
	var res []*pb.Project

	query:=r.db.Raw(`SELECT * from projects `).Scan(&project)
	if query.Error != nil {
		return nil,errors.New(`something went wrong`)
	}
	
	for _,p:=range project{
		query2:=r.db.Raw(`SELECT * from single_projects where project_id = ?`,p.ID).Scan(&singleProject)
		if query2.Error != nil {
			return nil,errors.New(`something went wrong`)
		}
		h:=pb.Project{
			ID: int32(p.ID),
			Title: p.Title,
			Description: p.Description,
			Price: singleProject.Price,
			DeliveryDays: int32(singleProject.DeliverDays),
			NumberOfRevisions: int32(singleProject.NumberOfRevisions),
		}
		res = append(res, &h)

	}
	return res, nil
}

func (r *Repo) ListOneProject(id string) (*pb.Project,error){
	var project domain.Project
	var singleProject domain.SingleProject
	query:=r.db.Raw(`SELECT * from projects where id = ?`,id).Scan(&project)
	if query.Error != nil {
		return nil,errors.New(`something went wrong`)
	}
	query2:=r.db.Raw(`SELECT * from single_project where project_id = ?`,id).Scan(&singleProject)
		if query2.Error != nil {
			return nil,errors.New(`something went wrong`)
		}
		h:=pb.Project{
			ID: int32(project.ID),
			Title: project.Title,
			Description: project.Description,
			Price: singleProject.Price,
			DeliveryDays: int32(singleProject.DeliverDays),
			NumberOfRevisions: int32(singleProject.NumberOfRevisions),
		}
		return &h,nil
}

func (r *Repo) ListMyProject(user_id string) ([]*pb.Project,error){
	var project []domain.Project
	var singleProject domain.SingleProject
	var res []*pb.Project

	query:=r.db.Raw(`SELECT * from projects where user_id = ?`,user_id).Scan(&project)
	if query.Error != nil {
		return nil,errors.New(`something went wrong`)
	}
	
	for _,p:=range project{
		query2:=r.db.Raw(`SELECT * from single_projects where project_id = ?`,p.ID).Scan(&singleProject)
		if query2.Error != nil {
			return nil,errors.New(`something went wrong`)
		}
		h:=pb.Project{
			ID: int32(p.ID),
			Title: p.Title,
			Description: p.Description,
			Price: singleProject.Price,
			DeliveryDays: int32(singleProject.DeliverDays),
			NumberOfRevisions: int32(singleProject.NumberOfRevisions),
		}
		res = append(res, &h)

	}
	return res, nil
}