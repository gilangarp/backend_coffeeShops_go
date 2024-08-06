package models

import "time"

/* var schemaProduct = `CREATE TABLE product (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_name VARCHAR(200) NOT NULL,
    img_product VARCHAR(200) NOT NULL,
    product_price INT NOT NULL,
    product_description VARCHAR(200) NOT NULL,
    category_id INT,
    favorite BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

)` */


type Product struct {
	Id                  string `db:"id" json:"id"`
	Product_name        string `db:"product_name" json:"product_name"`
	Img_product         string `db:"img_product" json:"img_product"`
	Product_price       int    `db:"product_price" json:"product_price"`
	Product_description string `db:"product_description" json:"product_description" params:"product_description"`
	Categorie_name       interface{}   `db:"categorie_name" json:"categorie_name"`
	Favorite            bool   `db:"favorite" json:"favorite"`
	Created_at          *time.Time `db:"created_at" json:"created_at"`
	Updated_at          *time.Time `db:"updated_at" json:"updated_at"`
}

type Filter struct{
	Category string `form:"category"`
	Favorite string `form:"favorite"`
	SearchText string `form:"SearchText"`
	Limit int `form:"limit"`
	Page int `form:"page"`
}

type Products []Product