package repository

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pb"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pkg/domain"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/utils/round"
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
	err := r.db.Raw(insertProjectQuery, time.Now(), time.Now(), project.Title, project.Description, project.Category, project.UserId).Scan(&id).Error
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
	err := r.db.Raw(query, time.Now(), req.Title, req.Description, req.Category, req.ProjectId, req.UserId).Error
	if err != nil {
		fmt.Println(err)
		return StatusInternalServerError, errors.New("error while updating project")
	}
	query = `UPDATE single_project set updated_at = ?, price = ?, deliver_days = ?, number_of_revisions = ? where project_id = ?`
	errs := r.db.Exec(query, time.Now(), req.Price, req.DeliveryDays, req.NumberOfRevisions, req.ProjectId).Error
	if errs != nil {
		fmt.Println(err)
		return StatusInternalServerError, errors.New("error while updating project")
	}
	return StatusOK, nil
}

func (r *Repo) DeleteProject(req *pb.RemProjectReq) error {
	query := `Delete from projects where id = ? and user_id = ?`
	err := r.db.Raw(query, req.ProjectId, req.UserId).Error
	if err != nil {
		return errors.New(`can't delete project.Not yours`)
	}
	return nil
}

func (r *Repo) ListProjects() ([]*pb.Project, error) {
	var project []domain.Project
	var res []*pb.Project

	query := r.db.Raw(`SELECT * from projects`).Scan(&project)
	if query.Error != nil {
		return nil, errors.New(`something went wrong`)
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	errChan := make(chan error, len(project))

	for _, p := range project {
		wg.Add(1)
		go func(p domain.Project) {
			defer wg.Done()
			var singleProject domain.SingleProject

			query2 := r.db.Raw(`SELECT * from single_projects where project_id = ?`, p.ID).Scan(&singleProject)
			if query2.Error != nil {
				errChan <- errors.New(`something went wrong`)
				return
			}

			h := pb.Project{
				ID:                int32(p.ID),
				Title:             p.Title,
				Description:       p.Description,
				Price:             singleProject.Price,
				DeliveryDays:      int32(singleProject.DeliverDays),
				NumberOfRevisions: int32(singleProject.NumberOfRevisions),
			}

			mu.Lock()
			res = append(res, &h)
			mu.Unlock()
		}(p)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return nil, <-errChan
	}

	return res, nil
}
func (r *Repo) ListOneProject(id string) (*pb.Project, error) {
	var project domain.Project
	var singleProject domain.SingleProject
	query := r.db.Raw(`SELECT * from projects where id = ?`, id).Scan(&project)
	if query.Error != nil {
		return nil, errors.New(`something went wrong`)
	}
	query2 := r.db.Raw(`SELECT * from single_projects where project_id = ?`, id).Scan(&singleProject)
	if query2.Error != nil {
		return nil, errors.New(`something went wrong`)
	}
	h := pb.Project{
		ID:                int32(project.ID),
		Title:             project.Title,
		Description:       project.Description,
		Price:             singleProject.Price,
		DeliveryDays:      int32(singleProject.DeliverDays),
		NumberOfRevisions: int32(singleProject.NumberOfRevisions),
	}
	return &h, nil
}

func (r *Repo) ListMyProject(user_id string) ([]*pb.Project, error) {
	var project []domain.Project
	var singleProject domain.SingleProject
	var res []*pb.Project

	query := r.db.Raw(`SELECT * from projects where user_id = ?`, user_id).Scan(&project)
	if query.Error != nil {
		return nil, errors.New(`something went wrong`)
	}
	errchan:=make(chan error,len(project))
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, p := range project {
		wg.Add(1)
		go func(p domain.Project){
			defer wg.Done()
			query2 := r.db.Raw(`SELECT * from single_projects where project_id = ?`, p.ID).Scan(&singleProject)
			if query2.Error != nil {
				errchan <- errors.New(`something went wrong`)
				return
			}
			h := pb.Project{
				ID:                int32(p.ID),
				Title:             p.Title,
				Description:       p.Description,
				Price:             singleProject.Price,
				DeliveryDays:      int32(singleProject.DeliverDays),
				NumberOfRevisions: int32(singleProject.NumberOfRevisions),
			}
			mu.Lock()
			res = append(res, &h)
			mu.Unlock()
		}(p)
	}
	wg.Wait()
	close(errchan)
	if len(errchan) > 0{
		return nil, <-errchan
	}
	return res, nil
}

func (r *Repo) GetProjectOrder(orderID string) (*domain.ProjectOrders, error) {
	var order domain.ProjectOrders
	q := r.db.Raw(`SELECT * FROM project_orders WHERE id = ?`, orderID).Scan(&order)
	if q.Error != nil {
		return nil, errors.New(`something went wrong`)
	}
	if q.RowsAffected == 0 {
		return nil, errors.New(`no orders found with this id`)
	}
	return &order, nil
}

func (r *Repo) UpdateOrderPaymentStatus(orderID string) (*domain.ProjectOrders, error) {
	var order domain.ProjectOrders
	q := r.db.Raw(`SELECT * FROM project_orders WHERE id = ?`, orderID).Scan(&order)
	if q.Error != nil {
		return nil, errors.New(`something went wrong`)
	}
	if q.RowsAffected == 0 {
		return nil, errors.New(`no orders found with this id`)
	}
	if order.Payment_status == "paid" {
		return nil, errors.New(`payment already paid`)
	}

	j := r.db.Exec(`UPDATE project_orders SET payment_status = 'paid' WHERE id = ?`, orderID)
	if j.Error != nil {
		return nil, errors.New(`something went wrong`)
	}
	return &order, nil
}

func (r *Repo) CheckProjectActiveAndExist(projectID string) (*domain.SingleProject, *domain.Project, error) {
	var project domain.SingleProject
	q := r.db.Raw(`SELECT * FROM single_projects WHERE project_id = ?`, projectID).Scan(&project)
	if q.Error != nil {
		return nil, nil, q.Error
	}
	if q.RowsAffected < 1 {
		return nil, nil, errors.New(`no orders found with this id`)
	}
	var pro domain.Project
	query := r.db.Raw(`SELECT * FROM projects WHERE id = ?`, projectID).Scan(&pro)
	if query.Error != nil {
		return nil, nil, q.Error
	}
	if query.RowsAffected < 1 {
		return nil, nil, errors.New(`no orders found with this id`)
	}
	return &project, &pro, nil
}

func (r *Repo) OrderProject(project *domain.SingleProject, pro *domain.Project, userId string) error {
	uid, err := strconv.Atoi(userId)
	if err != nil {
		return errors.New(`error in userid`)
	}
	freelancerFee := round.RoundToTwoDecimalPlaces(float64(project.Price) * 0.80)
	marketPlaceFee := round.RoundToTwoDecimalPlaces(float64(project.Price) * 0.20)
	q := r.db.Create(&domain.ProjectOrders{FreelancerID: pro.User_id, ClientID: uid, ProjectID: int(pro.ID), Payment_status: "unpaid", Delivery_status: "pending", FreelancerFee: freelancerFee, MarketplaceFee: marketPlaceFee})
	if q.Error != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}
