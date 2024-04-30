package repository

import (
	"errors"
	"strconv"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pb/job"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/domain"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewJobRepo(db *gorm.DB) Repo {
	return Repo{DB: db}
}

func (r *Repo) PostJob(req *job.PostjobReq) error {
	jobs := domain.Jobs{
		Title:       req.Title,
		Description: req.Description,
		TimePeriod:  req.TimePeriod,
		Type:        req.Type,
		Budget:      req.Budget,
		Category:    req.Category,
		Client_id:   int(req.UserId),
	}
	err := r.DB.Create(&jobs).Error
	if err != nil {
		return err
	}

	for _, s := range req.Skills {
		err := r.DB.Create(&domain.JobSkills{Job_id: int(jobs.ID), Skill_id: int(s)}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repo) ViewProposalsForFreelancer(userID int) (*[]domain.Proposals, error) {
	var proposals []domain.Proposals
	query := `SELECT * From proposals where user_id = ?`
	err := r.DB.Raw(query, userID).Scan(&proposals).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("no proposals found. propose to any job to view proposals")
		}
		return nil, err
	}
	return &proposals, nil
}

func (r *Repo) Proposal(req *job.ProposalReq) error {
	jobID, _ := strconv.Atoi(req.JobId)
	UserId, _ := strconv.Atoi(req.UserId)
	err := r.DB.Create(&domain.Proposals{
		Cover_letter: req.CoverLetter,
		Budget:       req.Budget,
		JobId:        jobID,
		User_id:      UserId,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) FindJob(jobID string) error {
	var count int
	query := `SELECT count(*) FROM jobs WHERE id = ?`
	c := r.DB.Raw(query, jobID).Scan(&count)
	if count == 0 {
		return errors.New("no job found with this id")
	}
	err := c.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("no job found with this id")
		}
		return err
	}
	return nil
}

func (r *Repo) ViewProposalsForClients(client_id int) error {
	query := `SELECT * FROM proposals INNER JOIN jobs on jobs.id=proposals.job_id WHERE jobs.client_id = ?`
	err := r.DB.Raw(query, client_id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetCategory() (*job.GetCategoryRes, error) {
	var category *job.GetCategoryRes
	err := r.DB.Raw("SELECT * FROM categories").Scan(&category.Categories).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}
func (r *Repo) AddCategory(category *job.AddCategoryReq) (*job.AddCategoryRes, error) {
	err := r.DB.Raw("INSERT INTO categories (category) VALUES (?) RETURNING category", category.Category).Error
	if err != nil {
		return nil, errors.New(`something went wrong`)
	}
	return &job.AddCategoryRes{Status: 200, Response: "added category successfully"}, nil
}

func (r *Repo) GetMyJobs(user_id string) ([]*job.Job, error) {
	var jobs []domain.Jobs
	err := r.DB.Raw(`SELECT * FROM jobs WHERE client_id = ?`, user_id).Scan(&jobs).Error
	if err != nil {
		return nil, err // return the actual error instead of a generic one
	}

	var resultJobs []*job.Job
	for _, jobi := range jobs {
		var skills []int
		err := r.DB.Raw(`SELECT skill_id FROM job_skills WHERE job_id = ?`, jobi.ID).Scan(&skills).Error
		if err != nil {
			return nil, err
		}
		// jobSkills, err := r.GetSkillsForJob(skills) // Assuming you have a method GetSkillsForJob to fetch skills by IDs
		// if err != nil {
		// 	return nil, err
		// }
		// var category string
		// err = r.DB.Raw(`SELECT category FROM categories WHERE id = ?`, jobi.Category).Scan(&category).Error
		// if err != nil {
		// 	return nil, err
		// }
		resultJobs = append(resultJobs, &job.Job{
			ID:          int64(jobi.ID),
			Title:       jobi.Title,
			Description: jobi.Description,
			// Skills:      jobSkills,
			// Category:    category,
			TimePeriod:  jobi.TimePeriod,
			Type:        jobi.Type,
			Budget:      jobi.Budget,
		})
	}
	return resultJobs, nil
}
