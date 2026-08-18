package stdlib

import (
	"context"
	"database/sql/driver"
	"github.com/jackc/pgx/v5"
)

// ... existing code ...

func (c *conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	// Ensure the transaction is bound to the context lifecycle
	// If the context is cancelled, we must ensure the underlying connection
	// is notified or the query is interrupted.
	tx, err := c.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	return &txWrapper{tx: tx, ctx: ctx}, nil
}

// ... existing code ...
