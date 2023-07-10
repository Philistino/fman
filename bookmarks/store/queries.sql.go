// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: queries.sql

package store

import (
	"context"
)

const createBookmark = `-- name: CreateBookmark :exec
INSERT
    or IGNORE INTO bookmarks (path)
VALUES (?)
`

func (q *Queries) CreateBookmark(ctx context.Context, path string) error {
	_, err := q.exec(ctx, q.createBookmarkStmt, createBookmark, path)
	return err
}

const deleteBookmark = `-- name: DeleteBookmark :exec
DELETE FROM bookmarks
WHERE path = ?
`

func (q *Queries) DeleteBookmark(ctx context.Context, path string) error {
	_, err := q.exec(ctx, q.deleteBookmarkStmt, deleteBookmark, path)
	return err
}

const getBookmarks = `-- name: GetBookmarks :many
SELECT path
FROM bookmarks
`

func (q *Queries) GetBookmarks(ctx context.Context) ([]string, error) {
	rows, err := q.query(ctx, q.getBookmarksStmt, getBookmarks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var path string
		if err := rows.Scan(&path); err != nil {
			return nil, err
		}
		items = append(items, path)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
