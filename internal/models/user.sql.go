// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package models

import (
	"context"

	"github.com/jackc/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(id,name)
values($1,$2)
Returning id, name, bio, updated_at
`

type CreateUserParams struct {
	ID   pgtype.Text `db:"id" json:"id"`
	Name pgtype.Text `db:"name" json:"name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (*User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.ID, arg.Name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.UpdatedAt,
	)
	return &i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE from users where id = $1
Returning id, name, bio, updated_at
`

func (q *Queries) DeleteUser(ctx context.Context, id pgtype.Text) (*User, error) {
	row := q.db.QueryRowContext(ctx, deleteUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.UpdatedAt,
	)
	return &i, err
}

const findUserByID = `-- name: FindUserByID :one
SELECT id, name, bio, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) FindUserByID(ctx context.Context, id pgtype.Text) (*User, error) {
	row := q.db.QueryRowContext(ctx, findUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.UpdatedAt,
	)
	return &i, err
}

const getFunction = `-- name: GetFunction :many
select  from test()
`

type GetFunctionRow struct {
}

func (q *Queries) GetFunction(ctx context.Context) ([]*GetFunctionRow, error) {
	rows, err := q.db.QueryContext(ctx, getFunction)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetFunctionRow
	for rows.Next() {
		var i GetFunctionRow
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getListUser = `-- name: GetListUser :many
SELECT id, name, bio, updated_at FROM users
offset $1 limit $2
`

type GetListUserParams struct {
	Offset int32 `db:"offset" json:"offset"`
	Limit  int32 `db:"limit" json:"limit"`
}

func (q *Queries) GetListUser(ctx context.Context, arg GetListUserParams) ([]*User, error) {
	rows, err := q.db.QueryContext(ctx, getListUser, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users set name = $1
where id = $2
Returning id, name, bio, updated_at
`

type UpdateUserParams struct {
	Name pgtype.Text `db:"name" json:"name"`
	ID   pgtype.Text `db:"id" json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (*User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.Name, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.UpdatedAt,
	)
	return &i, err
}
