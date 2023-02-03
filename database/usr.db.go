package database

import (
	"time"
	modelsusr "todo-api/models/models-usr"

	"gorm.io/gorm"
)

func (i *instacePostgres) SingUp(usr *modelsusr.SingUp_Request) (uint, error) {
	u := modelsusr.User{
		Model: gorm.Model{
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
		Name:     usr.Name,
		Lastname: usr.Lastname,
		Email:    usr.Email,
		Pass:     usr.Pass,
		Position: usr.Position,
		Status:   usr.Status,
	}

	data := i.db.Create(&u)
	return u.ID, data.Error
}

func (i *instacePostgres) GetUsrById(id uint) (*modelsusr.GetUsr_ID, error) {
	var usr modelsusr.GetUsr_ID
	err := i.db.Table("users").Select("id, name, lastname, email, status, position").First(&usr, id).Error
	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (i *instacePostgres) GetUsrByEmail(email string) (*modelsusr.GetUsr_Email, error) {

	var usr modelsusr.GetUsr_Email
	err := i.db.Table("users").Select("id, position, pass").Where("email = ?", email).Scan(&usr)
	if err.Error != nil {
		return nil, err.Error
	}
	return &usr, nil
}
