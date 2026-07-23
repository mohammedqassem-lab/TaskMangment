package repositry

import (
	model "TaskMangment/Internal/Model"
	"context"
	"database/sql"
)

type UserRepositry struct {
	db *sql.DB
}

func GetNewUserRepositry(db *sql.DB) *UserRepositry {
	return &UserRepositry{
		db: db,
	}
}
func (r *UserRepositry) Create(ctx context.Context, user *model.User) error {
	q := "INSERT INTO users(name,email,Hashpassword) VALUES($1,$2,$3) RETURNING id"
	err := r.db.QueryRowContext(ctx, q, user.Name, user.Email, user.Hashpassword).Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}
