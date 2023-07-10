package bookmarks

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/Philistino/fman/bookmarks/store"
	_ "modernc.org/sqlite"
)

//go:embed .sqlite/schema.sql
var ddl string

type Querier struct {
	*store.Queries
	db *sql.DB
}

func NewQueries(ctx context.Context, path string) (*Querier, error) {

	// connect
	con, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	// create tables
	if _, err := con.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	// prepare statements and return queries object
	quieries, err := store.Prepare(ctx, con)
	if err != nil {
		return nil, err
	}
	q := Querier{
		Queries: quieries,
		db:      con,
	}
	return &q, err
}

func (q *Querier) Close() error {
	q.Queries.Close()
	return q.db.Close()
}

func (q *Querier) CreateBookmarks(ctx context.Context, paths []string) error {
	tx, err := q.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()
	qtx := q.WithTx(tx)

	for _, path := range paths {
		if err := qtx.CreateBookmark(ctx, path); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (q *Querier) DeleteBookMarks(ctx context.Context, paths []string) error {
	tx, err := q.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()
	qtx := q.WithTx(tx)

	for _, path := range paths {
		if err := qtx.DeleteBookmark(ctx, path); err != nil {
			return err
		}
	}
	return tx.Commit()
}
