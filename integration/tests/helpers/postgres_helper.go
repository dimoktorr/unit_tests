package helpers

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"log"
)

type db interface {
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

type transactionKey struct{}

type PostgresHelper struct {
	pgx     *pgxpool.Pool
	scanAPI *pgxscan.API
}

func NewPostgresHelper(conn *pgxpool.Pool, scanApi *pgxscan.API) *PostgresHelper {
	return &PostgresHelper{
		pgx:     conn,
		scanAPI: scanApi,
	}
}

func (o *PostgresHelper) conn(ctx context.Context) db {
	if tx, ok := ctx.Value(transactionKey{}).(pgx.Tx); ok {
		return tx
	}

	return o.pgx
}

func (o *PostgresHelper) Exec(ctx context.Context, query string, args []interface{}) error {
	commandTag, err := o.conn(ctx).Exec(ctx, query, args...)
	log.Default().Print("query: %s, number of rows affected: %d", query, commandTag.RowsAffected())

	return err
}

func (o *PostgresHelper) QueryRows(ctx context.Context, dst interface{}, query string, args []interface{}) error {
	rows, err := o.conn(ctx).Query(ctx, query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return err
		}
		return err
	}

	log.Default().Print("query rows: %s, rows affected: %d", query, rows.CommandTag().RowsAffected())
	defer rows.Close()

	err = o.scanAPI.ScanAll(dst, rows)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}
	return nil
}

func (o *PostgresHelper) QueryRow(ctx context.Context, dst interface{}, query string, args []interface{}) error {
	row := o.conn(ctx).QueryRow(ctx, query, args...)
	log.Default().Print("query row: %s", query)

	return row.Scan(dst)
}

func (o *PostgresHelper) QueryOne(ctx context.Context, dst interface{}, query string, args []interface{}) error {
	rows, err := o.conn(ctx).Query(ctx, query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return err
		}
		return err
	}
	log.Default().Print("query one: %s, rows affected: %d", query, rows.CommandTag().RowsAffected())
	defer rows.Close()

	err = o.scanAPI.ScanOne(dst, rows)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}
	return nil
}
