// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: kv.sql

package dal

import "context"

const DeleteKVPostgresql = `-- name: DeleteKV :execrows
DELETE FROM "kv" WHERE "key" = $1
`

func (q *PostgresqlAccess) DeleteKV(ctx context.Context, key string) (int64, error) {
	result, err := q.db.ExecContext(ctx, DeleteKVPostgresql, key)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const IndexKVPostgresql = `-- name: IndexKV :many
SELECT key, value, create_at, update_at, create_by, update_by FROM "kv" ORDER BY "key" ASC LIMIT $1 OFFSET $2
`

func (q *PostgresqlAccess) IndexKV(ctx context.Context, arg *IndexKVParams) ([]*Kv, error) {
	rows, err := q.db.QueryContext(ctx, IndexKVPostgresql, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Kv
	for rows.Next() {
		var i Kv
		if err := rows.Scan(
			&i.Key,
			&i.Value,
			&i.CreateAt,
			&i.UpdateAt,
			&i.CreateBy,
			&i.UpdateBy,
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

const InsertKVPostgresql = `-- name: InsertKV :execrows
INSERT INTO "kv" ("key", "value", "create_at", "update_at", "create_by", "update_by")
VALUES ($1, $2, $3, $4, $5, $6)
`

func (q *PostgresqlAccess) InsertKV(ctx context.Context, arg *InsertKVParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, InsertKVPostgresql,
		arg.Key,
		arg.Value,
		arg.CreateAt,
		arg.UpdateAt,
		arg.CreateBy,
		arg.UpdateBy,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}