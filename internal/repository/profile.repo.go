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

/* Created Profile */
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

/* Get All Profile */
func (r *RepoProfile) GetAllProfile() (*models.Profiles , error){
	query := `SELECT * FROM public.profile `
	data := models.Profiles{}

	if err := r.Select(&data , query); err != nil {
		return nil , err
	}

	return &data , nil
}

/* Get Detail Profile */
func (r *RepoProfile) GetDetailProfile(id string) (*models.Profile, error) {
	query := `SELECT * FROM public.profile WHERE user_id = $1`
	row := r.QueryRow(query, id)

	var profile models.Profile
	if err := row.Scan(
		&profile.User_id,
		&profile.Display_name,
		&profile.First_name,
		&profile.Last_name,
		&profile.Birth_date,
		&profile.Image,
		&profile.Delivery_address,
		&profile.Role,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error retrieving profile: %w", err)
	}

	return &profile, nil
}

/* Edit Profile */
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

/* Delete Profiles */
func (r *RepoProfile) DeleteProfiles(id string) (string, error) {
    query := `DELETE FROM public.profile WHERE user_id = $1 RETURNING user_id`
    row := r.QueryRow(query, id)

    var deletedID string
    if err := row.Scan(&deletedID); err != nil {
        if err == sql.ErrNoRows {
            return "", fmt.Errorf("profile with ID %s not found", id)
        }
        return "", fmt.Errorf("error while deleting profile: %w", err)
    }

    return "delete successful", nil
}

