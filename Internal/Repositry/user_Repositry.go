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
	q := `INSERT INTO users(name,email,Hashpassword) VALUES($1,$2,$3) RETURNING id`
	err := r.db.QueryRowContext(ctx, q, user.Name, user.Email, user.Hashpassword).Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepositry) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	q := `SELECT id,name,email,Hashpassword FROM users WHERE email=$1`
	row := r.db.QueryRowContext(ctx, q, email)
	var user model.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Hashpassword)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositry) SaveRefreshToken(ctx context.Context, token model.RefreshToken) error {
	q := `INSERT INTO refresh_tokens(user_id,token,expires_at) VALUES($1,$2,$3)`
	_, err := r.db.ExecContext(ctx, q, token.UserId, token.Token, token.ExpiresAt)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepositry) ValidateRefreshToken(ctx context.Context, token string) (model.RefreshToken, error) {
	q := `SELECT id,user_id,token,expires_at FROM refresh_tokens WHERE revoked=false AND token=$1`
	row := r.db.QueryRowContext(ctx, q, token)
	var RefreshToken model.RefreshToken
	err := row.Scan(&RefreshToken.Id, &RefreshToken.UserId, &RefreshToken.Token, &RefreshToken.ExpiresAt)
	return RefreshToken, err
}
func (r *UserRepositry) MakerevokedTrue(ctx context.Context, id int64) error {
	q := `UPDATE refresh_tokens SET revoked=true
	where id=$1`
	_, err := r.db.ExecContext(ctx, q, id)
	return err
}
func (r *UserRepositry) GetRevokedToken(ctx context.Context) ([]*int64, error) {
	query := `SELECT id from refresh_tokens
	where revoked=false
	AND expires_at < now()`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*int64
	for rows.Next() {
		var task int64
		err := rows.Scan(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
