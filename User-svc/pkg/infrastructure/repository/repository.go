package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/auth"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/job"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/project"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pkg/domain"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
	job.UnimplementedJobserviceServer
	project.UnimplementedUserServiceServer
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) CreateUser(user *auth.UserSignupReq) (*domain.User, error) {
	User := domain.User{
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  user.Password,
		Country:   user.Country,
		Role:      user.Role,
		Is_active: true,
	}
	query := r.db.Create(&User).Scan(&User)
	if query.Error != nil {
		return nil, query.Error
	}

	return &User, nil

}

func (r *Repo) CheckIsUserActive(email string) (error) {
	var res bool
	query := r.db.Raw(`select is_active from users where email = ?`,email).Scan(&res)
	if query.Error != nil {
		return query.Error
	}
	if !res{
		return errors.New("user is blocked")
	}
	return nil
}

func (r *Repo) CheckUserExist(email, phone string) error {
	var count int
	query := r.db.Raw(`SELECT COUNT(*) FROM users WHERE email = ? OR phone = ?`, email, phone).Scan(&count)
	if query.Error != nil {
		return query.Error
	}
	if count == 0 {
		return nil
	}
	return errors.New("user already exist")
}

func (r *Repo) AdminLogin(email string) (*domain.Admin, error) {
	var details domain.Admin
	if err := r.db.Raw("SELECT * FROM admins WHERE email=?", email).Scan(&details).Error; err != nil {
		return nil, err
	}
	return &details, nil
}



func (r *Repo) GetUser(email string) (*domain.UserModel, error) {
	var user domain.User
	var count int
	query := r.db.Raw(`SELECT * FROM users WHERE email = ? `, email).Scan(&user)
	if query.Error != nil {
		return nil, errors.New("something went wrong")

	}
	query = r.db.Raw(`SELECT count(*) FROM users WHERE email = ? `, email).Scan(&count)
	if query.Error != nil {
		return nil, errors.New("something went wrong")

	}
	if count < 1{
		return nil, errors.New(`user doesn't exist`)
	}

	fmt.Println(user.ID)
	return &domain.UserModel{FirstName: user.FirstName,Email: user.Email,ID: user.ID,Password: user.Password,Role: user.Role}, nil
}

func (r *Repo) CheckOTP(email string) (int64, time.Time, error) {
	var expiration time.Time
	type OTPInfo struct {
		OTP        int64     `gorm:"column:otp"`
		Expiration time.Time `gorm:"column:expiration"`
	}
	var otpInfo OTPInfo
	if err := r.db.Raw("SELECT otp, expiration FROM otp_infos WHERE email = ? ORDER BY expiration DESC LIMIT 1;", email).Scan(&otpInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, expiration, errors.New("otp verification failed, no data found for this user-email")
		}
		return 0, expiration, errors.New("otp verification failed, error finding user data: " + err.Error())
	}

	return otpInfo.OTP, otpInfo.Expiration, nil
}

func (r *Repo) SaveOTP(otp int, email string, exp time.Time) error {
	query := `INSERT INTO otp_infos (email, otp, expiration) VALUES ($1, $2, $3)`
	err := r.db.Exec(query, email, otp, exp).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteOTP() error {
	query := "DELETE FROM otp_infos WHERE expiration < CURRENT_TIMESTAMP - INTERVAL '2 minutes';"
	err := r.db.Exec(query).Error
	if err != nil {
		return err
	}
	log.Println("expired otps deleted")
	return nil
}

func (r *Repo) AddEducation(edu *auth.AddEducationReq) (*domain.Freelancer_Education, error) {
	var res domain.Freelancer_Education
	query := `INSERT INTO freelancer_educations(created_at,updated_at,school,course,area_of_study,year_started,year_ended,description,user_id) Values(?,?,?,?,?,?,?,?,?)`
	err := r.db.Raw(query, time.Now(), time.Now(), edu.School, edu.Course, edu.AreaOfStudy, edu.Date_Started, edu.Date_Ended, edu.Description, edu.UserId).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *Repo) UpdateEducation(edu *auth.UpdateEducationReq) (*domain.Freelancer_Education, error) {
	var res domain.Freelancer_Education

	query := `UPDATE freelancer_educations SET updated_at = ?, school = ?,course = ?,area_of_study = ?,year_started = ?,year_ended = ?,description = ? WHERE user_id = ? AND id = ?`
	err := r.db.Raw(query, time.Now(), edu.School, edu.Course, edu.AreaOfStudy, edu.Date_Started, edu.Date_Ended, edu.Description, edu.UserId, edu.EducationId).Scan(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("no education found to update with this id")
		}
		return nil, err
	}
	return &res, nil
}

func (r *Repo) DeleteEducation(edu *auth.DeleteEducationReq) error {

	query := `DELETE FROM freelancer_educations WHERE user_id = ? AND id = ?`
	err := r.db.Exec(query, edu.UserId, edu.EducationId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("no education found to update with this id")
		}
		return err
	}
	return nil
}

func (r *Repo) AddProfileDescription(req *auth.APDReq) (*domain.Freelancer_Description, error) {
	var res domain.Freelancer_Description
	q := `SELECT * FROM freelancers_descriptions where user_id = ?`
	if r.db.Raw(q, req.UserId).RowsAffected != 0 {
		return nil, errors.New("profile description already exist. update if you want to change")
	}
	query := `INSERT INTO freelancer_descriptions(created_at,updated_at,Title,description,Hourly_rate,user_id) VALUES(?,?,?,?,?,?)`
	err := r.db.Raw(query, time.Now(), time.Now(), req.Title, req.Description, req.HourlyRate, req.UserId).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *Repo) UpdateProfileDescription(req *auth.UPDReq) (*domain.Freelancer_Description, error) {
	var res domain.Freelancer_Description
	query := `UPDATE freelancer_descriptions SET updated_at = ?, Title = ?,description = ?,Hourly_rate = ? WHERE user_id = ?`
	err := r.db.Exec(query, time.Now(), req.Title, req.Description, req.HourlyRate, req.UserId).Scan(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("no education found to update with this id")
		}
		return nil, err
	}
	return &res, nil
}

func (r *Repo) AddExperience(edu *auth.ExpReq) error {
	var res domain.Freelancer_Experiences
	query := `INSERT INTO freelancer_experiences(created_at,updated_at,company,city,country,title,from_date,to_date,description,user_id) Values(?,?,?,?,?,?,?,?,?,?)`
	err := r.db.Raw(query, time.Now(), time.Now(), edu.Company, edu.City, edu.Country, edu.Title, edu.FromDate, edu.ToDate, edu.Description, edu.UserId).Scan(&res).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateExperience(edu *auth.ExpReq) error {
	var res domain.Freelancer_Education

	query := `UPDATE freelancer_experiences SET updated_at = ?, company = ?,city = ?,country = ?,title = ?,from_date = ?,to_date = ?,description = ? WHERE user_id = ? AND id = ?`
	err := r.db.Raw(query, time.Now(), edu.Company, edu.City, edu.Country, edu.Title, edu.FromDate, edu.ToDate, edu.Description, edu.UserId, edu.ExpId).Scan(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("no education found to update with this id")
		}
		return err
	}
	return nil
}

func (r *Repo) DeleteExperience(edu *auth.DltExpReq) error {

	query := `DELETE FROM freelancer_experiences WHERE user_id = ? AND id = ?`
	err := r.db.Exec(query, edu.UserId, edu.ExperienceId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("no experience found to update with this id")
		}
		return err
	}
	return nil
}

func (r *Repo) GetUserWithId(id string) (*auth.User, error) {
	var user domain.User
	query := `SELECT * FROM users WHERE id = ?`
	err := r.db.Raw(query, id).Scan(&user).Error
	if err != nil {
		return nil, errors.New("something went wrong")
	}
	return &auth.User{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Phone:     user.Phone,
		Email:     user.Email,
		Country:   user.Country,
	}, nil
}

func (r *Repo) GetProfileDescription(id string) (*auth.UPDReq, error) {
	var pd domain.Freelancer_Description
	query := `SELECT * FROM freelancer_descriptions WHERE user_id = ?`
	err := r.db.Raw(query, id).Scan(&pd).Error
	if err != nil {
		return nil, errors.New("something went wrong")
	}
	return &auth.UPDReq{
		Title:       pd.Title,
		Description: pd.Description,
		HourlyRate:  int64(pd.Hourly_rate),
	}, nil
}

func (r *Repo) GetEducations(id string) ([]*auth.Education, error) {
	var educations []domain.Freelancer_Education
	query := `SELECT * FROM freelancer_educations where user_id = ?`
	err := r.db.Raw(query, id).Scan(&educations).Error
	if err != nil {
		return nil, errors.New("something went wrong")
	}
	var edu []*auth.Education
	for _, l := range educations {
		res := &auth.Education{
			EducationId:  int64(l.ID),
			School:       l.School,
			Description:  l.Description,
			AreaOfStudy:  l.Area_Of_Study,
			Course:       l.Course,
			Date_Started: l.Year_Started,
			Date_Ended:   l.Year_Ended,
		}
		edu = append(edu, res)
	}

	return edu, nil
}

func (r *Repo) GetExperience(id string) ([]*auth.ExpReq, error) {
	var educations []domain.Freelancer_Experiences
	query := `SELECT * FROM freelancer_experiences where user_id = ?`
	err := r.db.Raw(query, id).Scan(&educations).Error
	if err != nil {
		return nil, errors.New("something went wrong")
	}
	var exp []*auth.ExpReq
	for _, l := range educations {
		res := &auth.ExpReq{
			Company:     l.Company,
			City:        l.City,
			Country:     l.Country,
			Title:       l.Title,
			FromDate:    l.FromDate,
			ToDate:      l.ToDate,
			Description: l.Description,
		}
		exp = append(exp, res)
	}

	return exp, nil
}

func (r *Repo) UpdatePassword(password, email string) error {
	query := `UPDATE users SET password = ? WHERE email = ?`
	err := r.db.Exec(query, password, email).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("no user found ")
		}
		return err
	}
	return nil
}

func (r *Repo) AddSkill(req *auth.AddSkillReq) (int, error) {
	skill:=domain.Skill{Skill: req.Skill,Description: req.Description}
	err := r.db.Create(&skill).Error
	if err != nil {
		return 500, err
	}
	return 200, nil
}

func (r *Repo) GetSkills(user_id string) ([]string,error){
	var skill []string
	query:=r.db.Raw(`SELECT skills.skill FROM skills INNER JOIN freelancer_skills on skills.id = freelancer_skills.skill_id WHERE freelancer_skills.freelancer_id = ?`,user_id).Scan(&skill)
	if query.Error != nil{
		return nil,errors.New(`no skills found`)
	}
	return skill,nil
}


func (r *Repo) BlockUser(userID string) (int, error) {
	var active bool
	query := `SELECT is_active FROM users where id = ?`
	err := r.db.Raw(query, userID).Scan(&active).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 404, errors.New("user not found with this id")
		}
		return 500, errors.New("something went wrong")
	}
	if !active {
		return 409, errors.New("user already blocked")
	}
	query = `UPDATE users SET is_active = false WHERE id = ?`
	err = r.db.Raw(query, userID).Scan(&active).Error
	if err != nil {
		return 500, errors.New("something went wrong")
	}

	return 200, nil

}

func (r *Repo) UnBlockUser(userID string) (int, error) {
	var active bool
	query := `SELECT is_active FROM users where id = ?`
	err := r.db.Raw(query, userID).Scan(&active).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 404, errors.New("user not found with this id")
		}
		return 500, errors.New("something went wrong")
	}
	if active {
		return 409, errors.New("user already Unblocked")
	}
	query = `UPDATE users SET is_active = true WHERE id = ?`
	err = r.db.Raw(query, userID).Scan(&active).Error
	if err != nil {
		return 500, errors.New("something went wrong")
	}

	return 200, nil
}

func (r *Repo) UpdateProfilePhoto(userID,url string) error{
	query := `UPDATE users SET profile_url = ? WHERE id = ?`
	err:=r.db.Raw(query,url,userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetJobsSkills(ctx context.Context,req *job.Req) (*job.Res,error){
	var skill []string
	query:=`SELECT skill FROM skills WHERE id = ?`
	for _,id:=range req.Skill{
		var skil string
		err:=r.db.Raw(query,id).Scan(&skil).Error
		if err != nil {
			return nil,err
		}
		skill = append(skill, skil)
	}
	return &job.Res{Skill: skill},nil
}

func (r *Repo) CheckSkillsExist(skills []int64) error{
	var count int64
	if err := r.db.Model(&domain.Skill{}).Where("id IN ?", skills).Count(&count).Error; err != nil {
		return err
	}
	if count != int64(len(skills)) {
		return errors.New("some skills do not exist")
	}
	return nil
}

func (r *Repo) UpdateSkillUserProfile(user_id string, skill []int64) ([]string,error){
	var skills []string
	uid,_:=strconv.Atoi(user_id)
	for _,s:=range skill{
		query:=r.db.Where("freelancer_id = ? AND skill_id = ?",user_id,s).Attrs(&domain.Freelancer_skills{Freelancer_id: uid,Skill_id: int(s)}).FirstOrCreate(&domain.Freelancer_skills{})
		if query.Error != nil{
			return nil,errors.New(`something went wrong`)
		}
		var skil string
		err:=r.db.Raw(`SELECT skill FROM skills WHERE id = ?`,s).Scan(&skil).Error
		if err != nil {
			return nil,errors.New(`something went wrong`)
		}
		skills = append(skills, skil)
	}
	
	return skills,nil
}



func (r *Repo) CheckUserOnboardStatus(user_id string) error{
	var status bool
	query:=r.db.Raw(`SELECT paypal_onboard_status FROM users where id = ?`,user_id).Scan(&status)
	if status{
		return errors.New("user already onboarded to paypal")
	}
	if query.Error != nil {
		return errors.New(`something went wrong`)
	}
	if query.RowsAffected == 0{
		return errors.New("no user found with this id")
	}
	return nil
}

func (r *Repo) GetFreelancerPaypalEmail(ctx context.Context,req *job.Preq)(*job.Pres,error){
	var email string
	query:=r.db.Raw(`SELECT email  FROM freelancer_paypals where user_id = ?`,req.UserID).Scan(&email)
	if query.RowsAffected == 0{
		return &job.Pres{Error: "user doesn't added paypal account"},nil
	}
	if query.Error != nil {
		fmt.Println(query.Error)
		return &job.Pres{Error: "something went wrong"}, nil
	}
	return &job.Pres{Email: email},nil
}

func (r *Repo) GetFreelancerPaypalEmails(ctx context.Context,req *project.Preq)(*project.Pres,error){
	var email string
	query:=r.db.Raw(`SELECT email  FROM freelancer_paypals where user_id = ?`,req.UserID).Scan(&email)
	if query.RowsAffected == 0{
		return &project.Pres{Error: "user doesn't added paypal account"},nil
	}
	if query.Error != nil {
		fmt.Println(query.Error)
		return &project.Pres{Error: "something went wrong"}, nil
	}
	return &project.Pres{Email: email},nil
}

func (r *Repo) AddPaymentEmail(userID string,email string) (error){
	userid,err:=strconv.Atoi(userID)
	if err != nil {
		return errors.New(`user id is not integer.enter correctly`)
	}
	r.db.Create(&domain.Freelancer_paypal{UserID: uint(userid),Email: email})
	return nil
}

func (r *Repo) CheckFreelancerExist(userID int32) error{
	var count int 
	q:=r.db.Raw(`SELECT count(*) from users where id = ?`,userID).Scan(&count)
	if q.Error != nil {
		return errors.New(`something went wrong`)
	}
	if count == 1{
		return nil
	}
	return errors.New(`no freelancer found with this id`)
}

func (r *Repo) CheckContractWithFreelancerAndClient(freelancerID int32,clientID int64) error{
	var count int
	q:=r.db.Raw(`SELECT count(*) from contracts where freelancer_id = ? AND client_id = ?`,freelancerID,clientID).Scan(&count)
	if q.Error != nil {
		return errors.New(`something went wrong`)
	}
	if count >= 1{
		return nil
	}
	return errors.New(`you have no relation with this freelancer to review`)
}

func (r *Repo) AddReviewForFreelancer(req *auth.ReviewFlancerReq) error{
	err:=r.db.Create(&domain.FreelancerReview{Rating: int(req.Rating),Review: req.Review,Freelancer_id: int(req.FreelancerId),Client_id: int(req.ClientId)}).Error
	if err != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}

func (r *Repo) UpdateCmpDtails(details *auth.UpdCompDtlReq) error{
	Dtls:=&domain.ClientCompany{
		ClientID: int(details.UserId),
		CompanyName: details.CompanyName,
		NumberOfEmployees: int(details.NumberOfEmployees),
		Website: details.Website,
		Tagline: details.Tagline,
		Industry: details.Industry,
	}
	query:=r.db.Where("client_id = ?",details.UserId).Attrs(Dtls).FirstOrCreate(Dtls)
	if query.Error != nil{
		return errors.New(`something went wrong`)
	}
	return nil
}

func (r *Repo) UpdateCompContact(details *auth.UpdCompContReq) error{
	Cont:=&domain.CompanyAddress{
		ClientID: int(details.UserId),
		OwnerName: details.OwnerName,
		Phone: details.Phone,
		Country: details.Address.Country,
		State: details.Address.State,
		District: details.Address.District,
		City: details.Address.City,
		Pincode: details.Address.Pincode,
	}
	query:=r.db.Where("client_id = ?",details.UserId).Attrs(Cont).FirstOrCreate(&domain.CompanyAddress{})
	if query.Error != nil{
		return errors.New(`something went wrong`)
	}
	return nil
}

func (r *Repo) GetCompanyDetails(userID int32) (*auth.UpdCompDtlReq,error){
	var client domain.ClientCompany
	query:=r.db.Raw(`SELect * From client_companies Where client_id = ?`,userID).Scan(&client)
	if query.Error!= nil{
		return nil,errors.New(`something went wrong`)
	}
	if query.RowsAffected == 0 {
		return nil,nil
	}
	return &auth.UpdCompDtlReq{
		CompanyName: client.CompanyName,
		Website: client.Website,
		NumberOfEmployees: int32(client.NumberOfEmployees),
		Tagline: client.Tagline,
		Industry: client.Industry,
	},nil
}

func (r *Repo) ContactDetails(userID int32) (*auth.UpdCompContReq,error){
	var client domain.CompanyAddress
	query:=r.db.Raw(`SELect * From company_addresses Where client_id = ?`,userID).Scan(&client)
	if query.Error!= nil{
		return nil,errors.New(`something went wrong`)
	}
	if query.RowsAffected == 0 {
		return nil,nil
	}
	return &auth.UpdCompContReq{
		OwnerName: client.OwnerName,
		Phone: client.Phone,
		Address: &auth.Address{
			Country: client.Country,
			State: client.State,
			District: client.District,
			City: client.City,
			Pincode: client.Pincode,
		},
	},nil
}

func (r *Repo) GetReviews(userID string) ([]*auth.Reviews,error){
	var reviews []domain.FreelancerReview
	query:=r.db.Raw(`select * from freelancer_reviews where freelancer_id = ?`,userID).Scan(&reviews)
	if query.Error!= nil{
		return nil,errors.New(`something went wrong`)
	}
	if query.RowsAffected == 0 {
		return nil,errors.New(`no reviews found`)
	}
	var rr []*auth.Reviews
	for _,r:=range reviews{
		rr = append(rr, &auth.Reviews{
			Review: r.Review,
			Rating: int32(r.Rating),
			ClientID: strconv.Itoa(r.Client_id),
		})
	}
	return rr,nil
}

func (r *Repo) CheckPaypalEmailAdded(ctx context.Context,req *job.CReq) (*job.CRes, error) {
	var count int
	query:=r.db.Raw(`SELECT count(*) FROM freelancer_paypals WHERE user_id = ?`,req.UserId).Scan(&count)
	if query.Error != nil {
		return &job.CRes{Status: http.StatusFailedDependency,Error: "something went wrong",Exist: false},nil
	}
	if count < 1{
		return &job.CRes{Status: http.StatusNotFound,Error: "add payment email to post a job",Exist: false},nil
	}
	return &job.CRes{Status: http.StatusOK,Exist: true},nil
}

func (r *Repo) GetTalents(q string) ([]*auth.Talents,error){
	var users []domain.User
	query:=r.db.Raw(`SELECT * FROM users WHERE role = 'freelancer'`).Scan(&users)
	if query.Error != nil {
		return nil,errors.New(`something went wrong`)
	}
	var talents []*auth.Talents
	for _,u:=range users{
		var d domain.Freelancer_Description
		var sk []int
		var sh string
		var skill []string
		r.db.Raw(`select * From freelancer_descriptions Where user_id = ?`,u.ID).Scan(&d)
		r.db.Raw(`select skill_id From freelancer_skills Where freelancer_id = ?`,u.ID).Scan(&sk)
		for _,s:=range sk{
		r.db.Raw(`select skill From skills Where id = ?`,s).Scan(&sh)
		skill = append(skill, sh)
		}

		talents = append(talents, &auth.Talents{
			FreelancerId: int32(u.ID),
			Name: u.FirstName,
			Title: d.Title,
			Description: d.Description,
			HourlyRate: fmt.Sprintf("%d",d.Hourly_rate),
			Skills: skill,
		})
	}
	return talents,nil
}