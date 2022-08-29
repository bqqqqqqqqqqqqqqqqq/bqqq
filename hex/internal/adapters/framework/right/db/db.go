package db

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	//connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failur:%v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failuer:%v", err)
	}

	return &Adapter{db: db}, nil
}

func (da Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db close failure:%v", err)
	}
}

func (da Adapter) AddToHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arith_history").Columns("data", "answer", "operation").
		Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}
	return nil
}
