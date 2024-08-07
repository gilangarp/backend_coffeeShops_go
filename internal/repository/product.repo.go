package repository

import (
	"database/sql"
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

/* Created Product */
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

/* Get All Product */
func (r *RepoProduct) GetAllProduct(params *models.Filter) (*models.Products, error) {
    query := `
        SELECT p.id , p.product_name , p.img_product , p.product_price , p.product_description , c.categorie_name , p.created_at 
		FROM 
			public.product p 
			JOIN public.category c ON p.category_id = c.id 

    `
    
    values := []interface{}{}
    whereClauses := []string{}
    whereClauses = append(whereClauses, "p.is_deleted = FALSE")

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

/* Get Detail Product */
func (r *RepoProduct) GetDetailProduct(id string) (*models.Product, error) {
	query := `SELECT p.product_name ,p.img_product ,p.product_price ,p.product_description ,c.categorie_name , p.created_at  FROM public.product p 
    join category c on p.category_id = c.id 
    WHERE p.id = $1`
	row := r.QueryRow(query, id)

	var product models.Product
	if err := row.Scan(
		&product.Product_name,
		&product.Img_product,
		&product.Product_price,
		&product.Product_description,
		&product.Categorie_name,
		&product.Created_at,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error retrieving profile: %w", err)
	}

	return &product, nil
}

/* Edit Product */
func (r *RepoProduct) EditProduct(data *models.EditProduct, id string) (*models.EditProduct, error) {
    query := `UPDATE product SET `
    var values []interface{}
    condition := false

    if data.Product_name != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`product_name = $%d`, len(values)+1)
        values = append(values, data.Product_name)
        condition = true
    }

    if data.Img_product != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`img_product = $%d`, len(values)+1)
        values = append(values, data.Img_product)
        condition = true
    }

    if data.Product_price > 0 {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`product_price = $%d`, len(values)+1)
        values = append(values, data.Product_price)
        condition = true
    }

    if data.Product_description != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`product_description = $%d`, len(values)+1)
        values = append(values, data.Product_description)
        condition = true
    }

    if data.Category_id > 0 {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`category_id = $%d`, len(values)+1)
        values = append(values, data.Category_id)
        condition = true
    }

    if !condition {
        return nil, fmt.Errorf("no fields to update")
    }

    query += fmt.Sprintf(` WHERE id = $%d RETURNING product_name, img_product, product_price, product_description, category_id`, len(values)+1)
    values = append(values, id)

    row := r.DB.QueryRow(query, values...)
    var product models.EditProduct
    err := row.Scan(
        &product.Product_name,
        &product.Img_product,
        &product.Product_price,
        &product.Product_description,
        &product.Category_id,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("product with id = %s not found", id)
        }
        return nil, fmt.Errorf("query execution error: %w", err)
    }

    return &product, nil
}

/* Delete Product */
func (r *RepoProduct) DeleteProduct(id string) (string, error){
    query := `UPDATE product SET is_deleted = true WHERE id = $1`
    result, err := r.Exec(query, id)
    if err != nil {
        return "", fmt.Errorf("error while delete product: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return "", fmt.Errorf("error while fetching affected rows: %w", err)
    }

    if rowsAffected == 0 {
        return "", fmt.Errorf("product with ID %s not found", id)
    }

    return "Delete successful", nil
}
