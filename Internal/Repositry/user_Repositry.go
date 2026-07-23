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
func (r *UserRepositry) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	q := "SELECT id,name,email,Hashpassword FROM users WHERE email=$1"
	row := r.db.QueryRowContext(ctx, q, email)
	var user model.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Hashpassword)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
