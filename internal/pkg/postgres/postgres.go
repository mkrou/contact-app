package postgres

import (
	"context"
	"contact/internal/pkg/config"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-pg/pg/v10"
)

//NewPostgresDB will create new connection to DB
func NewPostgresDB(c config.Config, l log.Logger) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Postgres.Host, c.Postgres.Port),
		User:     c.Postgres.Username,
		Password: c.Postgres.Password,
		Database: c.Postgres.Database,
	})

	db.AddQueryHook(Hook{logger: l})

	return db
}

// Hook represents actions before and after query
type Hook struct {
	logger log.Logger
}

//BeforeQuery do nothing special
func (h Hook) BeforeQuery(ctx context.Context, _ *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

//AfterQuery will write debug log with formatted query
func (h Hook) AfterQuery(_ context.Context, q *pg.QueryEvent) error {
	text, err := q.FormattedQuery()
	if err != nil {
		return err
	}

	_ = level.Debug(h.logger).Log("SQL", string(text))

	return nil
}
