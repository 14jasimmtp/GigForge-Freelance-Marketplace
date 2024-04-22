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
	jobs:=domain.Jobs{
		Title:       req.Title,
		Description: req.Description,
		TimePeriod:  req.TimePeriod,
		Level:       req.FreelancerLevel,
		Budget:      req.Budget,
		Category:    req.Category,
		Client_id: int(req.UserId),
	}
	err := r.DB.Create(&jobs).Error
	if err != nil {
		return err
	}

	for _,s:=range req.Skills{
		err := r.DB.Create(&domain.JobSkills{Job_id: int(jobs.ID),Skill_id: int(s)}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repo) ViewProposalsForFreelancer(userID int) (*[]domain.Proposals,error){
	var proposals []domain.Proposals
	query:=`SELECT * From proposals where user_id = ?`
	err:=r.DB.Raw(query,userID).Scan(&proposals).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil,errors.New("no proposals found. propose to any job to view proposals")
		}
		return nil,err
	}
	return &proposals,nil
}

func (r *Repo) Proposal(req *job.ProposalReq) error {
	jobID,_:=strconv.Atoi(req.JobId)
	err := r.DB.Create(&domain.Proposals{
		Cover_letter: req.CoverLetter,
		Budget:       req.Budget,
		JobId:        jobID,
		User_id: int(req.UserId),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) FindJob(jobID string) error {
	var count int
	query := `SELECT count(*) FROM jobs WHERE id = ?`
	c:= r.DB.Raw(query, jobID).Scan(&count)
	if count == 0{
		return errors.New("no job found with this id")
	}
	err:=c.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("no job found with this id")
		}
		return err
	}
	return nil
}

func (r *Repo) ViewProposalsForClients(client_id int) error{
	query:=`SELECT * FROM proposals INNER JOIN jobs on jobs.id=proposals.job_id WHERE jobs.client_id = ?`
	err := r.DB.Raw(query,client_id).Error
	if err != nil {
		return err
	}
	return nil
}


func (r *Repo) 