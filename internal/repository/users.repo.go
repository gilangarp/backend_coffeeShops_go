package repository

import (
	"gilangarp/backend_coffeeShops_go/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

func (r *RepoUser) RegisterUser(data *models.User) (string, error) {
	query := `
	INSERT INTO users
	(email, phone, password)
	VALUES
	 	($1 , $2 , $3)
	`
	_, err := r.Exec(query, data.Email , data.Phone , data.Password)
	if err != nil {
        return "", err
    }
	
	return "registration successful", nil
}



func (r *RepoUser) GetAllUser() (*models.Users, error){
	query := `SELECT * FROM public.users order by created_at DESC`
	data := models.Users{}

	if err := r.Select(&data , query); err != nil {
		return nil , err
	}

	return &data , nil
}