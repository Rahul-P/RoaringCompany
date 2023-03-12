package roaringcompany

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	kitlog "github.com/go-kit/log"

	_ "github.com/microsoft/go-mssqldb"
)

// RecordFound is an struct, used to read ID in before Update/Delete requests.
type RecordFound struct {
	ID uint64
}

var db *sql.DB
var err error
var counts int64

func GetDbConn(context context.Context, logger kitlog.Logger, server, user, password, database, msdbname string, port int) (*sql.DB, error) {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	for {
		connection, err := openDb(connString, context, logger, msdbname)
		if err != nil {
			logger.Log(msdbname + " Database not yet ready ...")
			counts++
		} else {
			logger.Log("Connected to " + msdbname + " Database!")
			return connection, nil
		}

		if counts > TEN {
			logger.Log("Stepping back post 10 retries: ", err)
			return nil, err
		}

		logger.Log("Backing off for two seconds....")
		time.Sleep(TWO * time.Second)
		continue
	}
}

func Validate(context context.Context, db *sql.DB) error {
	// Check if database connection is Not null.
	if db == nil {
		return DatabaseIssue
	}

	// Check if database is alive.
	err := db.PingContext(context)
	if err != nil {
		return err
	}
	return nil
}

func GetTest() error {
	return nil
}

func openDb(dsn string, context context.Context, logger kitlog.Logger, msdbname string) (*sql.DB, error) {

	// Create connection pool
	db, err = sql.Open("sqlserver", dsn)
	if err != nil {
		logger.Log("Error creating connection to "+msdbname+" DB pool: ", err.Error())
		return nil, err
	}

	err = db.PingContext(context)
	if err != nil {
		logger.Log("Error PINGING to "+msdbname+" DB: ", err.Error())
		return nil, err
	}

	return db, nil
}
