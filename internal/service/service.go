package service

import (
	"context"
	"github/eugene-krivtsov/idler-email/internal/model/dto"
	"github/eugene-krivtsov/idler-email/internal/repository/mongo"
	postgres_repository "github/eugene-krivtsov/idler-email/internal/repository/postgres"
	"github/eugene-krivtsov/idler-email/pkg/mail"
)

type Mails interface {
	Send(ctx context.Context, toEmail, path string) error
}

type Users interface {
	GetByEmail(ctx context.Context, email string) (*dto.UserDTO, error)
}

type Services struct {
	MailService Mails
	UserService Users
}

type ServicesDepends struct {
	Sender               mail.Sender
	PostgresRepositories *postgres_repository.Repositories
	MongoRepositories    *mongo_repository.Repositories
}

func NewServices(deps ServicesDepends) *Services {
	mailService := NewMailService(deps.Sender, deps.MongoRepositories.Mails, deps.PostgresRepositories.Users)
	userService := NewUserService(deps.PostgresRepositories.Users)

	return &Services{
		MailService: mailService,
		UserService: userService,
	}
}
