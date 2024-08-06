package repository

import (
	"fmt"
	"gilangarp/backend_coffeeShops_go/internal/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type RepoProduct struct {
	*sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

func (r *RepoProduct) CreatedProduct(data *models.Product) (string, error) {
    query := `
    INSERT INTO public.product(
        product_name,
        img_product,
        product_price,
        product_description,
        category_id
        )
    VALUES
        ($1, $2, $3, $4, $5)
    `

    _, err := r.Exec(query, data.Product_name, data.Img_product, data.Product_price, data.Product_description, data.Category_id)
    if err != nil {
        return "", err
    }

    return "1 data product created", nil
}


func (r *RepoProduct) GetAllProduct(params *models.Filter) (*models.Products, error) {
    query := `
        SELECT p.id , p.product_name , p.img_product , p.product_price , p.product_description , c.categorie_name , p.created_at 
		FROM 
			public.product p 
			JOIN public.category c ON p.category_id = c.id 

    `
    
    values := []interface{}{}
    whereClauses := []string{}

	if params.Promo {
		query += ` inner join public.promo prm on p.id = prm.product_id `;
	}

    if params.SearchText  != "" {
        whereClauses = append(whereClauses, fmt.Sprintf("p.product_name ILIKE $%d", len(values)+1))
        values = append(values, fmt.Sprintf("%%%s%%", params.SearchText))
    }

    if params.Category != "" {
        whereClauses = append(whereClauses, fmt.Sprintf("c.categorie_name = $%d", len(values)+1))
        values = append(values, params.Category)
    }

	if params.Limit > 0 && params.Page > 0 {
		limit := params.Limit
		offset := (params.Page - 1) * limit
		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(values)+1, len(values)+2)
		values = append(values, limit, offset)
	}
	

    if len(whereClauses) > 0 {
        query += " WHERE " + strings.Join(whereClauses, " AND ")
    }

   var data models.Products

    if err := r.Select(&data, query, values...); err != nil {
        return nil, err
    }

	return &data, nil
}

