package repository

import (
	"fmt"
	"gilangarp/backend_coffeeShops_go/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoFavorite struct {
	*sqlx.DB
}

func NewFavorite(db *sqlx.DB) *RepoFavorite {
	return &RepoFavorite{db}
}

/* Created Favorite */
func (r *RepoFavorite) CreatedFavorite(data *models.Favorite , id string) (string, error){
	query := `
	INSERT INTO public.favorite (
    	user_id ,
    	product_id 
		)
	VALUES
	 	( $1 , $2 )`

	_, err := r.Exec(query, id , data.Product_id)
	if err != nil {
	 return "", fmt.Errorf("failed to create favorite: %w", err)
	}

	return "1 data favorite created", nil
}

/* Get Detail Favorite */
func (r *RepoFavorite) GetDetailFavorite(id string) (*models.FavoriteGets, error) {
    query := `SELECT 
        p2.display_name, 
        p.product_name, 
        p.product_price,
        p.img_product
    FROM public.favorite f
    JOIN product p ON f.product_id = p.id
    JOIN profile p2 ON f.user_id = p2.user_id
    JOIN users u ON f.user_id = u.id
    WHERE u.id = $1;`

    var data models.FavoriteGets

    if err := r.Select(&data, query, id); err != nil {
        return nil, err
    }

    return &data, nil
}

/* Delete Favorite */
func (r *RepoFavorite) DeleteFavorite(id string) (string, error) {
    query := `DELETE FROM public.favorite WHERE id = $1`
    result, err := r.Exec(query, id)
    if err != nil {
        return "", fmt.Errorf("error while deleting favorite: %w", err)
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return "", fmt.Errorf("error while fetching affected rows: %w", err)
    }
    if rowsAffected == 0 {
        return "", fmt.Errorf("product favorite with ID %s not found", id)
    }

    return "delete successful", nil
}
