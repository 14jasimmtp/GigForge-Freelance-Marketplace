package repository

// import (
// 	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/admin"
// 	"gorm.io/gorm"
// )

// type AdminRepo struct {
// 	db *gorm.DB
// }

// type AdminIFc interface {
// 	AddSkill(req *admin.AddSkillReq) (int, error)
// 	BlockUser(userID int64) (int, error)
// 	UnBlockUser(userID int64) (int, error)
// }

// func NewAdminRepo(db *gorm.DB) AdminIFc {
// 	return &AdminRepo{db: db}
// }

// func (r *AdminRepo) AddSkill(req *admin.AddSkillReq)(int,error){
// 	query:=`INSERT INTO skills(skill,description) VALUES(?,?)`
// 	err:=r.db.Exec(query,req.Skill,req.Description).Error
// 	if err != nil {
// 		return 500,err
// 	}
// 	return 200,nil
// }

// func (r *AdminRepo) BlockUser(userID int64)(int,error){
// 	var active bool
// 	query:=`SELECT is_active FROM users where id = ?`
// 	err:=r.db.Raw(query,userID).Scan(&active).Error
// 	if err != nil{
// 		if err == gorm.ErrRecordNotFound{
// 			return 404,errors.New("user not found with this id")
// 		}
// 		return 500,errors.New("something went wrong")
// 	}
// 	if !active{
// 		return 409,errors.New("user already blocked")
// 	}
// 	query =`UPDATE users SET is_active = ? WHERE id = ?`
// 	err =r.db.Raw(query,userID).Scan(&active).Error
// 	if err != nil {
// 		return 500,errors.New("something went wrong")
// 	}

// 	return 200,nil

// }

// func (r *AdminRepo) UnBlockUser(userID int64)(int,error){
// 	var active bool
// 	query:=`SELECT is_active FROM users where id = ?`
// 	err:=r.db.Raw(query,userID).Scan(&active).Error
// 	if err != nil{
// 		if err == gorm.ErrRecordNotFound{
// 			return 404,errors.New("user not found with this id")
// 		}
// 		return 500,errors.New("something went wrong")
// 	}
// 	if active{
// 		return 409,errors.New("user already Unblocked")
// 	}
// 	query =`UPDATE users SET is_active = ? WHERE id = ?`
// 	err =r.db.Raw(query,userID).Scan(&active).Error
// 	if err != nil {
// 		return 500,errors.New("something went wrong")
// 	}

// 	return 200,nil
// }
