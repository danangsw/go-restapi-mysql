package post

import (
	"context"
	"database/sql"

	model "github.com/danangsw/go-restapi-mysql/model"
	pRepo "github.com/danangsw/go-restapi-mysql/repository"
)

// NewSQLPostRepo retunrs implement of post repository interface
func NewSQLPostRepo(Conn *sql.DB) pRepo.PostRepo {
	return &mysqlPostRepo{
		Conn: Conn,
	}
}

type mysqlPostRepo struct {
	Conn *sql.DB
}

func (m *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.Post, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*model.Post, 0)
	for rows.Next() {
		data := new(model.Post)

		err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.Content,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlPostRepo) Fetch(ctx context.Context, num int64) ([]*model.Post, error) {
	query := "Select id, title, content From posts limit ?"

	return m.fetch(ctx, query, num)
}

func (m *mysqlPostRepo) GetByID(ctx context.Context, id int64) (*model.Post, error) {
	query := "Select id, title, content From posts where id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &model.Post{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, model.ErrNotFound
	}

	return payload, nil
}

func (m *mysqlPostRepo) Create(ctx context.Context, p *model.Post) (int64, error) {
	query := "Insert posts SET title=?, content=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Title, p.Content)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlPostRepo) Update(ctx context.Context, p *model.Post) (*model.Post, error) {
	query := "Update posts set title=?, content=? where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Title,
		p.Content,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *mysqlPostRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From posts Where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
