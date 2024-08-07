package repository

import (
	"database/sql"
	"fmt"
	"gilangarp/backend_coffeeShops_go/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoProfile struct {
	*sqlx.DB
}

func NewProfile(db *sqlx.DB) *RepoProfile {
	return &RepoProfile{db}
}

func (r *RepoProfile) CreatedProfile(data *models.Profile, id string) (string, error) {
	query := `
	INSERT INTO profile (
    	user_id,
    	display_name,
    	first_name,
    	last_name,
    	birth_date,
    	image,
    	delivery_address,
    	role
	) VALUES
	 ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.Exec(query, id, data.Display_name, data.First_name, data.Last_name, data.Birth_date, data.Image, data.Delivery_address, data.Role)
	if err != nil {
		return "", fmt.Errorf("failed to create profile: %w", err)
	}

	return "1 data profile created", nil
}


func (r *RepoProfile) EditProfile(data *models.Profile , id string) (*models.Profile, error) {
    query := `UPDATE profile SET `
    var values []interface{}
    condition := false

    if data.Display_name != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`display_name = $%d`, len(values)+1)
        values = append(values, data.Display_name)
        condition = true
    }

    if data.First_name != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`first_name = $%d`, len(values)+1)
        values = append(values, data.First_name)
        condition = true
    }

    if data.Last_name != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`last_name = $%d`, len(values)+1)
        values = append(values, data.Last_name)
        condition = true
    }

    if data.Birth_date != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`birth_date = $%d`, len(values)+1)
        values = append(values, data.Birth_date)
        condition = true
    }

    if data.Image != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`image = $%d`, len(values)+1)
        values = append(values, data.Image)
        condition = true
    }

    if data.Delivery_address != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`delivery_address = $%d`, len(values)+1)
        values = append(values, data.Delivery_address)
        condition = true
    }

    if data.Role != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`role = $%d`, len(values)+1)
        values = append(values, data.Role)
        condition = true
    }

    if !condition {
        return nil, fmt.Errorf("no fields to update")
    }

    query += fmt.Sprintf(` WHERE user_id = $%d RETURNING *`, len(values)+1)
    values = append(values, id)
	fmt.Print("ini dari value:", values)
	fmt.Print("ini dari query:", query)

    row := r.DB.QueryRow(query, values...)
    var profile models.Profile
    err := row.Scan(
        &profile.User_id,
        &profile.Display_name,
        &profile.First_name,
        &profile.Last_name,
        &profile.Birth_date,
        &profile.Image,
        &profile.Delivery_address,
        &profile.Role,
    )
	
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user with id = %s not found", id)
        }
        return nil, fmt.Errorf("query execution error: %w", err)
    }

    return &profile, nil
}
