package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"rest-ws/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	// SSL Disabled
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db: db}, nil
}

//Close connection to database
func (r *PostgresRepository) Close() error {
	return r.db.Close()
}

// InsertProduct inserts a new product into the database.
func (r *PostgresRepository) InsertProduct(ctx context.Context, product *models.Product) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO products (id, reference, name, description, price, quantity) VALUES ($1, $2, $3, $4, $5, $6)", product.Id, product.Reference, product.Name, product.Description, product.Price, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

// GetProductByID returns a product by its ID.
func (r *PostgresRepository) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	var product models.Product
	err := r.db.QueryRowContext(ctx, "SELECT id, reference, name, description, price, quantity FROM products WHERE id = $1", id).Scan(&product.Id, &product.Reference, &product.Name, &product.Description, &product.Price, &product.Quantity)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetProductByReference returns a product by its reference.
func (r *PostgresRepository) GetProductByReference(ctx context.Context, reference string) (*models.Product, error) {
	var product models.Product
	err := r.db.QueryRowContext(ctx, "SELECT id, reference, name, description, price, quantity FROM products WHERE reference = $1", reference).Scan(&product.Id, &product.Reference, &product.Name, &product.Description, &product.Price, &product.Quantity)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// InsertUser inserts a new user into the database.
func (r *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (id ,name, email, password) VALUES ($1, $2, $3, $4)", user.Id, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// GetUserById returns a user by its ID.
func (r *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRowContext(ctx, "SELECT id, name, email, password FROM users WHERE id = $1", id).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail returns a user by its email.
func (r *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRowContext(ctx, "SELECT id, name, email, password FROM users WHERE email = $1", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
