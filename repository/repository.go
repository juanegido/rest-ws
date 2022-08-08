package repository

import (
	"context"
	"rest-ws/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	InsertProduct(ctx context.Context, product *models.Product) error
	GetProductByID(ctx context.Context, id string) (*models.Product, error)
	GetProductByReference(ctx context.Context, reference string) (*models.Product, error)

	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func InsertProduct(ctx context.Context, product *models.Product) error {
	return implementation.InsertProduct(ctx, product)
}

func GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	return implementation.GetProductByID(ctx, id)
}

func GetProductByReference(ctx context.Context, reference string) (*models.Product, error) {
	return implementation.GetProductByReference(ctx, reference)
}

func Close() error {
	return implementation.Close()
}
