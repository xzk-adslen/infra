package base

import (
	"context"
	"database/sql"
	"log"

	"github.com/tietang/dbx"
)

const TX = "tx"

//Dao Data Access Object)
type BaseDao struct {
	TX *sql.Tx
}

func (d *BaseDao) SetTx(tx *sql.Tx) {
	d.TX = tx
}

type txFunc func(*dbx.TxRunner) error

func Tx(fn func(*dbx.TxRunner) error) error {
	return TxContext(context.Background(), fn)
}

func TxContext(ctx context.Context, fn func(runner *dbx.TxRunner) error) error {
	return DbxDatabase().Tx(fn)
}

func WithValueContext(parent context.Context, runner *dbx.TxRunner) context.Context {
	return context.WithValue(parent, TX, runner)
}

func ExecuteContext(ctx context.Context, fn func(*dbx.TxRunner) error) error {
	tx, ok := ctx.Value(TX).(*dbx.TxRunner)
	if !ok || tx == nil {
		log.Panic("是否在事务函数块中使用？")
	}
	return fn(tx)
}
