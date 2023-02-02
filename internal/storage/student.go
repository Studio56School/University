package storage

import (
	"context"
	"fmt"
	"github.com/Studio56School/university/internal/config"
	"github.com/Studio56School/university/internal/model"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func NewRepository(conf *config.Config, logger *zap.Logger) (*Repo, error) {
	pgDB, err := ConnectDB(conf)
	if err != nil {
		logger.Sugar().Error("Unable to connect")
		return nil, err
	}
	return &Repo{DB: pgDB}, nil
}

type Repo struct {
	DB *pgx.Conn
}

func (r *Repo) CreateUser(user model.User) (int, error) {

	var id int

	query := `INSERT INTO public.users
	(name, username, password)
	VALUES ($1, $2, $3) RETURNING id`

	err := r.DB.QueryRow(context.Background(), query, user.Name, user.Username, user.Password).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error occurred while creating users in system: %v", err)
	}

	return id, nil
}

func (r *Repo) GetUser(username, password string) (model.User, error) {
	var user model.User

	query := `SELECT id, name FROM public.users
	WHERE username=$1 AND password=$2`

	row := r.DB.QueryRow(context.Background(), query, username, password)

	err := row.Scan(&user.Id, &user.Name)

	if err != nil {
		return model.User{}, fmt.Errorf("error occurred while creating users in system: %v", err)
	}

	return user, err
}
