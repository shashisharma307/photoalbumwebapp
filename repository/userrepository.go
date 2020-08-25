package repository

import (
	"fmt"
	"albumwebapp/models"
	"github.com/jinzhu/gorm"
)

type UserRepositoryError struct {
	error
}




type UserReposityory struct {
	DB *gorm.DB
}

func GetUserRespository(db *gorm.DB) UserReposityory{
	return UserReposityory {DB: db}
}

func (u *UserReposityory) GetAll() ([]models.User, error){
	var users []models.User
	u.DB.Find(&users)
	if len(users) == 0 || users ==nil{
		return users, &UserRepositoryError{fmt.Errorf("no record found")}
	}else{
		return users, nil

	}
	err := fmt.Errorf("Server error")
	return nil, &UserRepositoryError{err}
}

func (u *UserReposityory) GetByID(id int) (models.User, error){
	var user models.User
	d := u.DB.Find(user, id)
	if d.RowsAffected == 0{
		return  user, &UserRepositoryError{fmt.Errorf("can not create user")}
	}else{
		return user, nil
	}
}

func (u *UserReposityory) CountEmailId(email string) (bool, error){
	var user models.User
	var count int64

	u.DB.Where("email = ?", email).Find(&user).Count(&count)
	if count > 0 {
		return false, &UserRepositoryError{fmt.Errorf("email alrady taken")}
	}
	return true, nil
}

func (u *UserReposityory) GetEmailById(email string) (models.User, error){
	var user models.User
	u.DB.Where("email = ?", email).Find(&user)
	if user.UserId != 0{
		return user, nil
	}
	return user, &UserRepositoryError{fmt.Errorf("user is not registered")}
}

func (u *UserReposityory) Save(user models.User) (models.User,error){
	d := u.DB.Save(&user)
	if d.Error !=nil{
		return  user, &UserRepositoryError{fmt.Errorf(d.Error.Error())}
	}else{
		return user, nil
	}
	return  user, &UserRepositoryError{fmt.Errorf("can not create user")}
}

func (u UserReposityory) Delete(user models.User){
	u.DB.Delete(&user)
}

