package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gin/models"
	"strings"
)

type ClienteRepository interface {
	GetAll(ctx context.Context, nameFilter, emailFilter string) ([]models.Cliente, error)
	GetByID(ctx context.Context, id string) (*models.Cliente, error)
	Create(ctx context.Context, cliente models.Cliente) (*models.Cliente, error)
	Update(ctx context.Context, id string, cliente models.Cliente) error
	Delete(ctx context.Context, id string) error
}

type postgresRepository struct {
	DB *sql.DB
}

func NewClienteRepository(db *sql.DB) ClienteRepository {
	return &postgresRepository{DB: db}
}

func (r *postgresRepository) GetAll(ctx context.Context, nameFilter, emailFilter string) ([]models.Cliente, error) {
	query := "SELECT id, name, email FROM clientes"
	arguments := []string{}
	args := []interface{}{}
	argId := 1
	if nameFilter != "" {
		arguments = append(arguments, fmt.Sprintf("name ILIKE $%d", argId))
		args = append(args, fmt.Sprintf("%%%s%%", nameFilter))
		argId++
	}
	if emailFilter != "" {
		arguments = append(arguments, fmt.Sprintf("email ILIKE $%d", argId))
		args = append(args, fmt.Sprintf("%%%s%%", emailFilter))
		argId++
	}
	if len(arguments) > 0 {
		query += " WHERE " + strings.Join(arguments, " AND ")
	}

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []models.Cliente
	for rows.Next() {
		var cliente models.Cliente
		if err := rows.Scan(&cliente.ID, &cliente.Name, &cliente.Email); err != nil {
			return nil, err
		}
		clientes = append(clientes, cliente)
	}
	return clientes, nil
}

func (r *postgresRepository) GetByID(ctx context.Context, id string) (*models.Cliente, error) {
	var cliente models.Cliente
	query := "SELECT id, name, email FROM clientes WHERE id=$1"
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&cliente.ID, &cliente.Name, &cliente.Email)
	if err != nil {
		return nil, err
	}
	return &cliente, nil
}

func (r *postgresRepository) Create(ctx context.Context, cliente models.Cliente) (*models.Cliente, error) {
	query := "INSERT INTO clientes (name, email) VALUES ($1, $2) RETURNING id"
	err := r.DB.QueryRowContext(ctx, query, cliente.Name, cliente.Email).Scan(&cliente.ID)
	if err != nil {
		return nil, err
	}
	return &cliente, nil
}

func (r *postgresRepository) Update(ctx context.Context, id string, cliente models.Cliente) error {
	query := "UPDATE clientes SET name=$1, email=$2 WHERE id=$3"
	result, err := r.DB.ExecContext(ctx, query, cliente.Name, cliente.Email, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *postgresRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM clientes WHERE id=$1"
	result, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
