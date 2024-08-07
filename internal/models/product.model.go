package models

import "time"

type Product struct {
	Id                  string `db:"id" json:"id"`
	Product_name        string `db:"product_name" json:"product_name"`
	Img_product         string `db:"img_product" json:"img_product"`
	Product_price       int    `db:"product_price" json:"product_price"`
	Product_description string `db:"product_description" json:"product_description" params:"product_description"`
	Categorie_name      string   `db:"categorie_name" json:"categorie_name"`
	Category_id        int  `db:"category_id" json:"category_id"`
	Created_at          *time.Time `db:"created_at" json:"created_at"`
	Updated_at          *time.Time `db:"updated_at" json:"updated_at"`
}

type EditProduct struct {
	Product_name        string `db:"product_name" json:"product_name"`
	Img_product         string `db:"img_product" json:"img_product"`
	Product_price       int    `db:"product_price" json:"product_price"`
	Product_description string `db:"product_description" json:"product_description" params:"product_description"`
	Category_id        int  `db:"category_id" json:"category_id"`
}

type Filter struct{
	Category string `form:"category"`
	Favorite string `form:"favorite"`
	SearchText string `form:"SearchText"`
	Limit int `form:"limit"`
	Page int `form:"page"`
	Promo bool 
}

type Products []Product