package repository

import (
	"database/sql"
	"fmt"
	"gilangarp/backend_coffeeShops_go/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

/* Register User */
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

/* Get All User */
func (r *RepoUser) GetAllUser() (*models.Users, error){
	query := `SELECT * FROM public.users WHERE is_deleted = FALSE`
	data := models.Users{}

	if err := r.Select(&data , query); err != nil {
		return nil , err
	}

	return &data , nil
}

func (r *RepoUser) GetDetailUser(id string) (*models.UserDetail, error) {
    query := `SELECT email, phone, created_at, updated_at FROM public.users WHERE id = $1 AND is_deleted = FALSE`
    data := models.UserDetail{}

    if err := r.Get(&data, query, id); err != nil {
        return nil, err
    }

    return &data, nil
}

/* Edit Users */
func (r *RepoUser) EditUsers(data *models.User, id string) (*models.User, error) {
    query := `UPDATE users SET `
    var values []interface{}
    condition := false

    if data.Email != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`email = $%d`, len(values)+1)
        values = append(values, data.Email)
        condition = true
    }

    if data.Phone != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`phone = $%d`, len(values)+1)
        values = append(values, data.Phone)
        condition = true
    }

    if data.Password != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`password = $%d`, len(values)+1)
        values = append(values, data.Password)
        condition = true
    }

    if !condition {
        return nil, fmt.Errorf("no fields to update")
    }

    // Ensure the WHERE clause is correctly appended
    query += fmt.Sprintf(` WHERE id = $%d RETURNING email, phone, password`, len(values)+1)
    values = append(values, id)

    row := r.DB.QueryRow(query, values...)
    var user models.User
    err := row.Scan(
        &user.Email,
        &user.Phone,
        &user.Password,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user with id = %s not found", id)
        }
        return nil, fmt.Errorf("query execution error: %w", err)
    }

    return &user, nil
}

func (r *RepoUser) DeleteUser(id string) (string, error){
    query := `UPDATE users SET is_deleted = true WHERE id = $1`
    result, err := r.Exec(query, id)
    if err != nil {
        return "", fmt.Errorf("error while delete delete: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return "", fmt.Errorf("error while fetching affected rows: %w", err)
    }

    if rowsAffected == 0 {
        return "", fmt.Errorf("delete with ID %s not found", id)
    }

    return "Delete successful", nil
}
