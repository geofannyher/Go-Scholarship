package models

import (
	"context"
	"time"
)

type User struct {
	ID        int64     `json:"id" form:"id"`
	Name      string    `json:"name" form:"name" binding:"required"`
	Image     string    `json:"image" form:"image"`
	Email     string    `json:"email" form:"email" binding:"required,email"`
	Password  string    `json:"password" form:"password" binding:"required,min=4"`
	Role      string    `json:"role" form:"role"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

type Login struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=4"`
}

type UserRepository interface {
	Fetch(ctx context.Context) ([]User, error)
	FetchById(ctx context.Context, id int64) (User, error)
	Create(ctx context.Context, u *User) (User, error)
	Update(ctx context.Context, id int64, u *User) (User, error)
	Delete(ctx context.Context, id int64) error
	Login(ctx context.Context, l *Login) (User, error)
	Register(ctx context.Context, u *User) (User, error)
}
