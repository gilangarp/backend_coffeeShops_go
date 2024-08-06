package models

import "time"

type User struct {
	ID        	string    `db:"id" json:"id"`
	Email     	string    `db:"email" json:"email"`
	Phone     	string    `db:"phone" json:"phone"`
	Password  	string    `db:"password" json:"password"`
	Is_deleted	*bool     `db:"is_deleted" json:"is_deleted"`
	CreatedAt 	time.Time `db:"created_at" json:"created_at"`
	UpdatedAt 	time.Time `db:"updated_at" json:"updated_at"`
}

type Users []User