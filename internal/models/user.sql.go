// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: user.sql

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(id,name,type,team_id)
values($1,$2,$3,$4)
Returning id, name, type, team_id, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	ID     pgtype.Text `db:"id" json:"id"`
	Name   pgtype.Text `db:"name" json:"name"`
	Type   pgtype.Text `db:"type" json:"type"`
	TeamID pgtype.Text `db:"team_id" json:"team_id"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (*User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Type,
		arg.TeamID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.TeamID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE from users where id = $1
Returning id, name, type, team_id, created_at, updated_at, deleted_at
`

func (q *Queries) DeleteUser(ctx context.Context, id pgtype.Text) (*User, error) {
	row := q.db.QueryRow(ctx, deleteUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.TeamID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const findUserByID = `-- name: FindUserByID :one
SELECT id, name, type, team_id, created_at, updated_at, deleted_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) FindUserByID(ctx context.Context, id pgtype.Text) (*User, error) {
	row := q.db.QueryRow(ctx, findUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.TeamID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getListUser = `-- name: GetListUser :many
SELECT id, name, type, team_id, created_at, updated_at, deleted_at FROM users
offset $1 limit $2
`

type GetListUserParams struct {
	Offset int32 `db:"offset" json:"offset"`
	Limit  int32 `db:"limit" json:"limit"`
}

func (q *Queries) GetListUser(ctx context.Context, arg GetListUserParams) ([]*User, error) {
	rows, err := q.db.Query(ctx, getListUser, arg.Offset, arg.Limit)
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
			&i.Type,
			&i.TeamID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users set name = $1
where id = $2
Returning id, name, type, team_id, created_at, updated_at, deleted_at
`

type UpdateUserParams struct {
	Name pgtype.Text `db:"name" json:"name"`
	ID   pgtype.Text `db:"id" json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (*User, error) {
	row := q.db.QueryRow(ctx, updateUser, arg.Name, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.TeamID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
