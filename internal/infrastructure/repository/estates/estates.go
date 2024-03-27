package estates

import (
	"api/internal/core"
	"context"

	"github.com/jmoiron/sqlx"
)

type estates struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *estates {
	return &estates{db: db}
}

func (r *estates) Add(ctx context.Context, estate *core.Estate) error {
	return nil
}

func (r *estates) GetAll(ctx context.Context) (*[]core.Estate, error) {
	var output []core.Estate
	query := "SELECT * FROM estates.estates"
	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var estate core.Estate
		if err := rows.StructScan(&estate); err != nil {
			return nil, err
		}
		output = append(output, estate)
	}

	return &output, nil
}

func (r *estates) GetOne(ctx context.Context, id int) (*core.Estate, error) {
	var output *core.Estate
	query := "SELECT * FROM estates.estates WHERE id = $1"
	err := r.db.QueryRowx(query, id).Scan(&output)
	return output, err
}
