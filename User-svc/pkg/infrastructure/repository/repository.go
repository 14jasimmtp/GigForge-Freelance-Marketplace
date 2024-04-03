package repository

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/auth"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pkg/domain"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) RepoIfc {
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
		Is_active: false,
	}
	query := r.db.Create(&User).Scan(&User)
	if query.Error != nil {
		return nil, query.Error
	}

	return &User, nil

}

func (r *Repo) CheckIsUserActive(email string) domain.User {
	var res domain.User
	query := r.db.Where(&domain.User{Email: email}).Scan(&res)
	if query.Error != nil {
		return domain.User{}
	}
	return res
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

func (r *Repo) GetUser(email string) (*domain.UserModel, error) {
	var user domain.UserModel
	query := `SELECT * FROM users WHERE email = ? `
	err := r.db.Raw(query, email).Scan(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user doesn't exist with this email")
		}
		return nil, errors.New("something went wrong")

	}
	fmt.Println(user.ID)
	return &user, nil
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
	err := r.db.Raw(query,time.Now(),time.Now(), edu.School, edu.Course, edu.AreaOfStudy, edu.Date_Started, edu.Date_Ended, edu.Description, edu.UserId).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *Repo) UpdateEducation(edu *auth.UpdateEducationReq) (*domain.Freelancer_Education, error) {
	var res domain.Freelancer_Education

	query := `UPDATE freelancer_educations SET updated_at = ? school = ?,course = ?,area_of_study = ?,year_started = ?,year_ended = ?,description ? WHERE user_id = ? AND e_id = ?`
	err := r.db.Exec(query, time.Now() ,edu.School, edu.Course, edu.AreaOfStudy, edu.Date_Started, edu.Date_Ended, edu.Description, edu.UserId, edu.EducationId).Scan(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("no education found to update with this id")
		}
		return nil, err
	}
	return &res, nil
}

func (r *Repo) DeleteEducation(edu *auth.DeleteEducationReq) error {

	query := `DELETE FROM freelancer_educations WHERE user_id = ? AND e_id = ?`
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
	query := `INSERT INTO freelancer_descriptions(created_at,updated_at,Title,description,Hourly_rate,user_id) VALUES(?,?,?,?)`
	err := r.db.Exec(query, time.Now(), time.Now(), req.Title, req.Description, req.HourlyRate, req.UserId).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *Repo) UpdateProfileDescription(req *auth.UPDReq) (*domain.Freelancer_Description, error) {
	var res domain.Freelancer_Description
	query := `UPDATE freelancer_descriptions SET updated_at = ?, Title = ?,description = ?,Hourly_rate = ?)WHERE user_id = ?`
	err := r.db.Exec(query, time.Now(), req.Title, req.Description, req.HourlyRate, req.UserId).Scan(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("no education found to update with this id")
		}
		return nil, err
	}
	return &res, nil
}
