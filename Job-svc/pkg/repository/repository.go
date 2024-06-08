package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pb/job"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pb/user"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/domain"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/utils/round"
	"gorm.io/gorm"
)

type Repo struct {
	DB  *gorm.DB
	job user.JobserviceClient
}

func NewJobRepo(db *gorm.DB, job user.JobserviceClient) Repo {
	return Repo{DB: db, job: job}
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

func (r *Repo) CheckJobExist(job_id string,userID int64) error{
	var count int
	query:=r.DB.Raw(`SELECT count(*) FROM jobs WHERE id = ? AND client_id = ?`,job_id,userID).Scan(&count)
	if query.Error != nil {
		fmt.Println(query.Error)
		return errors.New(`something went wrong`)
	}
	if count < 1{
		return errors.New(`no jobs found with the given id`)
	}
	return nil
}

func (r *Repo) EditJobPost(Post *job.EditjobReq) error{
	JobPost:=domain.Jobs{
		Title: Post.Title,
		Description: Post.Description,
		TimePeriod: Post.TimePeriod,
		Type: Post.Type,
		Category: Post.Category,
		Budget: Post.Budget,
		Client_id: int(Post.UserId),
	}
	query:=r.DB.Model(&domain.Jobs{}).Where("id = ?",Post.JobID).Updates(JobPost)
	if query.Error != nil {
		fmt.Println(query.Error)
		return errors.New(`something went wrong`)
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

func (r *Repo) GetCategory(query string) ([]*job.Category, error) {
	var category []*job.Category
	q := r.DB.Raw("SELECT id AS ID,category AS Category FROM categories WHERE category ILIKE %?% OR ? = ''", query).Scan(&category)
	if q.RowsAffected == 0 {
		return nil, errors.New(`no categories found`)
	}
	if q.Error != nil {
		return nil, errors.New(`something went wrong`)
	}
	return category, nil
}
func (r *Repo) AddCategory(category *job.AddCategoryReq) (*job.AddCategoryRes, error) {
	categori := domain.Category{Category: category.Category}
	err := r.DB.Create(&categori).Error
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
		skills := []int64{}
		err := r.DB.Raw(`SELECT skill_id FROM job_skills WHERE job_id = ?`, jobi.ID).Scan(&skills).Error
		if err != nil {
			return nil, err
		}
		jobSkills, err := r.job.GetJobsSkills(context.Background(), &user.Req{Skill: skills}) // Assuming you have a method GetSkillsForJob to fetch skills by IDs
		if err != nil {
			return nil, err
		}
		var category string
		err = r.DB.Raw(`SELECT category FROM categories WHERE id = ?`, jobi.Category).Scan(&category).Error
		if err != nil {
			return nil, err
		}
		resultJobs = append(resultJobs, &job.Job{
			ID:          int64(jobi.ID),
			Title:       jobi.Title,
			Description: jobi.Description,
			Skills:      jobSkills.Skill,
			Category:    category,
			TimePeriod:  jobi.TimePeriod,
			Type:        jobi.Type,
			Budget:      jobi.Budget,
		})
	}
	return resultJobs, nil
}

func (r *Repo) SendOffer(req *job.SendOfferReq) (*job.SendOfferRes, error) {
	err := r.DB.Create(&domain.Offer{
		Budget:        req.Budget,
		Offer_letter:  req.OfferLetter,
		Freelancer_id: int(req.FreelancerId),
		Starting_time: req.StartingTime,
		Client_id:     int(req.ClientId),
		Job_id:        int(req.JobId),
		Status:        "pending",
	}).Error
	if err != nil {
		println(err)
		return nil, errors.New(`something went wrong`)
	}
	return &job.SendOfferRes{Status: 200, Response: "offer letter sended successfully to freelancer"}, nil
}

func (r *Repo) GetOffers(user_id, status string) ([]*job.Offer, error) {
	var offers []*job.Offer
	query := r.DB.Raw(`SELECT id AS offer_id,client_id,job_id,budget,offer_letter,starting_time,status FROM offers WHERE freelancer_id = $1 AND (status = $2 OR $2 = '')`, user_id, status).Scan(&offers)
	if query.RowsAffected == 0 {
		return nil, errors.New(`no offers found`)
	}
	if query.Error != nil {
		return nil, errors.New(`something went wrong`)
	}
	return offers, nil
}

func (r *Repo) AcceptOffer(id string) error {
	var status string
	query := r.DB.Raw(`SELECT status from offers where id = ?`, id).Scan(&status)
	if query.RowsAffected == 0 {
		return errors.New("no offers found with this id enter correct id")
	}
	if query.Error != nil {
		print(query.Error)
		return errors.New(`something went wrong`)
	}
	if status == "accepted" {
		return errors.New("offer already accepted")
	}
	err := r.DB.Exec(`UPDATE offers set status = 'accepted' where id = ?`, id).Error
	if err != nil {
		print(err)
		return errors.New("something went wrong")
	}
	return nil
}

func (r *Repo) CreateContract(id string) (int, string, float32, error) {
	var offer domain.Offer
	fmt.Println(id)
	query := `SELECT * from offers where id = ?`
	err := r.DB.Raw(query, id).Scan(&offer).Error
	if err != nil {
		fmt.Println(err,"error")
		return 0, "", 0, errors.New(`something went wrong`)
	}
	var Type string
	query = `select type from jobs where id = ?`
	err = r.DB.Raw(query, offer.Job_id).Scan(&Type).Error
	if err != nil {
		return 0, "", 0, errors.New(`no job found with this id`)
	}
	fmt.Println(offer.Starting_time)
	start_date, err := time.Parse("02-01-2006", offer.Starting_time)
	if err!= nil {
		fmt.Println(err)
		return 0, "", 0, err
	}
	fmt.Println(start_date)
	contract := &domain.Contract{
		Client_id:      offer.Client_id,
		Start_date:     start_date,
		Freelancer_id:  offer.Freelancer_id,
		Job_id:         offer.Job_id,
		Paid_amount:    0,
		Pending_amount: int(offer.Budget),
		Type:           Type,
		Status:         "active",
	}
	err = r.DB.Create(contract).Error
	if err != nil {
		return 0, "", 0, errors.New(`error while creating contract`)
	}
	return int(contract.ID), contract.Type, offer.Budget, nil
}

func (r *Repo) SendFixedInvoice(id int, types string, budget float32) error {
    freelancerFee := round.RoundToTwoDecimalPlaces(float64(budget) * 0.80)
    marketPlaceFee := round.RoundToTwoDecimalPlaces(float64(budget) * 0.20)
fmt.Println(freelancerFee,marketPlaceFee)
    invoice := &domain.Invoice{
        Freelancer_fee:  freelancerFee,
        MarketPlace_fee: marketPlaceFee,
		Start_date: time.Time{},
		End_date: time.Time{},
        Status:          "unpaid",
        ContractID:      id,
    }
    err := r.DB.Create(invoice).Error
    if err != nil {
        fmt.Println(err)
        return errors.New("something went wrong")
    }
    return nil
}

func (r *Repo) CheckContractActive(id int32) (domain.Contract, error) {
	var contract domain.Contract
	query := r.DB.Raw(`select * from contracts where id = ?`, id).Scan(&contract)
	if query.RowsAffected == 0 {
		return domain.Contract{}, errors.New("no contracts found with this id")
	}
	if query.Error != nil {
		return domain.Contract{}, errors.New(`error fetching contract`)
	}
	return contract, nil
}

func (r *Repo) GetLastInvoiceDetails(contract_id int32) (domain.Invoice, error) {
	var Invoice domain.Invoice
	query := r.DB.Raw("select * from invoices where contract_id = ?", contract_id).Scan(&Invoice)
	if query.Error != nil {
		return domain.Invoice{}, errors.New("something went wrong")
	}
	return Invoice, nil
}

func (r *Repo) SendHourlyInvoice(id int, types string, budget float32, Hours float32, Start_date, End_date time.Time) error {
	amount := budget * Hours
	f:=round.RoundToTwoDecimalPlaces(float64(amount) * 0.80)
	m:=round.RoundToTwoDecimalPlaces(float64(budget) * 0.20)
	invoice := &domain.Invoice{
		Freelancer_fee:  f,
		MarketPlace_fee: m,
		Start_date:      Start_date,
		End_date:        End_date,
		Status:          "unpaid",
		ContractID:      id,
	}
	err := r.DB.Create(invoice)
	if err != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}

func (r *Repo) GetJobs() ([]*job.Job, error) {
	var jobs []domain.Jobs
	query := r.DB.Raw(`SELECT * FROM jobs `).Scan(&jobs)
	if query.Error != nil {
		return nil, errors.New(`error while fetching jobs`) // return the actual error instead of a generic one
	}
	if query.RowsAffected == 0{
		return nil, errors.New("no jobs found ")
	}

	var resultJobs []*job.Job
	for _, jobi := range jobs {
		var skills []int64
		err := r.DB.Raw(`SELECT skill_id FROM job_skills WHERE job_id = ?`, jobi.ID).Scan(&skills).Error
		if err != nil {
			return nil, err
		}
		jobSkills, err := r.job.GetJobsSkills(context.Background(), &user.Req{Skill: skills}) // Assuming you have a method GetSkillsForJob to fetch skills by IDs
		if err != nil {
			return nil, err
		}
		var category string
		err = r.DB.Raw(`SELECT category FROM categories WHERE id = ?`, jobi.Category).Scan(&category).Error
		if err != nil {
			return nil, err
		}
		resultJobs = append(resultJobs, &job.Job{
			ID:          int64(jobi.ID),
			Title:       jobi.Title,
			Description: jobi.Description,
			Skills:      jobSkills.Skill,
			Category:    category,
			TimePeriod:  jobi.TimePeriod,
			Type:        jobi.Type,
			Budget:      jobi.Budget,
		})
	}
	return resultJobs, nil
}

func (r *Repo) GetJob(id string) (*job.Job, error) {
	var jobs domain.Jobs
	query := r.DB.Raw(`SELECT * FROM jobs where id = ? `, id).Scan(&jobs)
	err := query.Error
	if err != nil {
		return nil, errors.New(`error while fetching jobs`) // return the actual error instead of a generic one
	}

	if query.RowsAffected == 0 {
		return nil, errors.New(`no job found with this id`)
	}
	var skills []int64
	err = r.DB.Raw(`SELECT skill_id FROM job_skills WHERE job_id = ?`, jobs.ID).Scan(&skills).Error
	if err != nil {
		return nil, err
	}
	jobSkills, err := r.job.GetJobsSkills(context.Background(), &user.Req{Skill: skills}) // Assuming you have a method GetSkillsForJob to fetch skills by IDs
	if err != nil {
		return nil, err
	}
	resultJobs := &job.Job{
		ID:          int64(jobs.ID),
		Title:       jobs.Title,
		Description: jobs.Description,
		Skills:      jobSkills.Skill,
		//Category:    category,
		TimePeriod: jobs.TimePeriod,
		Type:       jobs.Type,
		Budget:     jobs.Budget,
	}

	return resultJobs, nil

}

func (r *Repo) SearchJobs(category, paytype, query string, fixedRate, HourlyRate []string) ([]*job.Job, int32, error) {
	q := "%" + query + "%"
	cat, _ := strconv.Atoi(category)
	var jobs []domain.Jobs
	r.DB.Raw(`
	SELECT * FROM jobs 
	WHERE 
    (type = $1 OR $1 = '') AND 
    (category = $2 OR $2 = 0) AND 
    title ILIKE $3`,
		paytype, cat, q).Scan(&jobs)

	var resultJobs []*job.Job
	for _, jobi := range jobs {
		var skills []int64
		err := r.DB.Raw(`SELECT skill_id FROM job_skills WHERE job_id = ?`, jobi.ID).Scan(&skills).Error
		if err != nil {
			return nil, 500, err
		}
		jobSkills, err := r.job.GetJobsSkills(context.Background(), &user.Req{Skill: skills}) // Assuming you have a method GetSkillsForJob to fetch skills by IDs
		if err != nil {
			return nil, 500, err
		}
		var category string
		err = r.DB.Raw(`SELECT category FROM categories WHERE id = ?`, jobi.Category).Scan(&category).Error
		if err != nil {
			return nil, 500, err
		}
		resultJobs = append(resultJobs, &job.Job{
			ID:          int64(jobi.ID),
			Title:       jobi.Title,
			Description: jobi.Description,
			Skills:      jobSkills.Skill,
			Category:    category,
			TimePeriod:  jobi.TimePeriod,
			Type:        jobi.Type,
			Budget:      jobi.Budget,
		})
	}
	return resultJobs, 200, nil
}

func (r *Repo) FindJobExistOfClient(job_id, client_id string) error {
	var jobs *domain.Jobs
	r.DB.Raw(`select * from jobs WHERE id = $1 and client_id = $2`, job_id, client_id).Scan(&jobs)
	if jobs == nil {
		return errors.New(`no job found with this id of client`)
	}
	return nil
}

func (r *Repo) GetJobProposals(job_id string) ([]*job.Proposals, error) {
	var proposals []domain.Proposals
	query := r.DB.Raw(`SELECT * From proposals WHERE job_id = ?`, job_id).Scan(&proposals)
	if query.Error != nil {
		return nil, errors.New(`something went wrong`)
	}
	if query.RowsAffected == 0 {
		return nil, errors.New(`no proposals found for this job`)
	}
	var res []*job.Proposals
	for _, prop := range proposals {

		p := &job.Proposals{
			UserId:      int32(prop.User_id),
			Coverletter: prop.Cover_letter,
			Budget:      prop.Budget,
		}
		res = append(res, p)
	}
	return res, nil
}

func (r *Repo) GetInvoiceWithID(id string)(*domain.Invoice,error){
	var invoice domain.Invoice
	q:=r.DB.Raw("SELECT * FROM invoices WHERE id = ?",id).Scan(&invoice)
	if q.RowsAffected == 0{
		fmt.Println("no invoices found with this id")
		return nil,errors.New(`no invoices found with this id`)
	}
	if q.Error != nil {
		fmt.Println(q.Error)
		return nil,errors.New(`something went wrong`)
	}
	if invoice.Status == "paid"{
		return nil, errors.New("already paid")
	}
	return &invoice,nil
}

func (r *Repo) UpdateInvoicePaymentStatus(invoiceID string)(*domain.Invoice,error){
	var invoice domain.Invoice
	query:=r.DB.Raw(`SELECT * FROM invoices WHERE id = ?`,invoiceID).Scan(&invoice)
	if query.Error != nil {
		return nil,errors.New(`something went wrong`)
	}
	if invoice.Status=="unpaid"{
		r.DB.Exec(`UPDATE invoices SET status = 'paid' WHERE id = ?`,invoiceID)
	}
	return &invoice,nil
}

func (r *Repo) UpdateContractDetails(contractID int,freelancerFee,MarketplaceFee float64) error{
	err:=r.DB.Exec(`UPDATE contracts SET paid_amount=paid_amount+$1, pending_amount = pending_amount - $1 WHERE id = $2`,freelancerFee+MarketplaceFee,contractID).Error
	if err != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}

func (r *Repo) GetAllContracts(client_id int64) ([]domain.Contract,error){
	var contracts []domain.Contract
	q:=r.DB.Raw(`SELECT * FROM contracts WHERE client_id = ?`,client_id).Scan(&contracts)
	if q.Error != nil {
		return nil, errors.New(`something went wrong`)
	}
	if q.RowsAffected == 0 {
		return nil, errors.New(`no contracts.Make contract with freelancer to show contracts`)
	}
	return contracts,nil
}

func (r *Repo) GetOneContract(contractID string,clientID int64) (*domain.Contract,error){
	var contracts domain.Contract
	q:=r.DB.Raw(`SELECT * FROM contracts WHERE client_id = ? AND id = ?`,clientID,contractID).Scan(&contracts)
	if q.Error != nil {
		return nil, errors.New(`something went wrong`)
	}
	if q.RowsAffected == 0 {
		return nil, errors.New(`no contracts.Make contract with freelancer to show contracts`)
	}
	return &contracts,nil
}

func (r *Repo) GetInvoices(client_id int64,contractID string) ([]domain.Invoice,error){
	var invoices []domain.Invoice
	fmt.Println("1")
	q:=r.DB.Raw(`SELECT * From invoices where contract_id = ?`,contractID).Scan(&invoices)
	if q.Error!= nil {
		return nil,errors.New(`something went wrong`)
	}
	fmt.Println("2")
	if q.RowsAffected == 0{
		return nil, errors.New(`no invoices found`)
	}
	return invoices,nil
}

func (r *Repo) StoreAttachmentUrl(url string, req *job.AddAttachmentReq) error{
	cid,_:=strconv.Atoi(req.ContractID)
	attachment:= &domain.Attachment{
		ContractID: cid,
		UploadTime: time.Now(),
		AttachmentURL: url,
		Description: req.Description,
	}
	query:=r.DB.Create(attachment)
	if query.Error != nil{
		return errors.New(`something went wrong`)
	}
	return nil
}

func (r *Repo) GetAttachments(contractId string) ([]*job.Attachment,error){
	var attachments []domain.Attachment
	query:=r.DB.Raw(`select * From attachments where contract_id = ?`,contractId).Scan(&attachments)
	if query.Error != nil {
		return nil,errors.New(`something went wrong`)
	}
	if query.RowsAffected == 0{
		return nil,errors.New(`no attachments found`)
	}
	var aa []*job.Attachment
	for _,a:=range attachments{
		aa = append(aa, &job.Attachment{
			ContractID: int32(a.ContractID),
			AttachmentUrl: a.AttachmentURL,
			Description: a.Description,
			PostedDate: a.UploadTime.Format("02-01-2006"),
		})
	}
	return aa,nil
}
