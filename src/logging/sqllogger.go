package logging

import (
	"database/sql"
	"fmt"
	"go-backend/api"
)

type PostgresTransactionLogger struct {
	events chan<- Event
	errors <-chan error
	db     *sql.DB
}

func (l *PostgresTransactionLogger) WritePut(key string, value string) {
	l.events <- Event{EventType: EventPut, Key: key, Value: value}
}

func (l *PostgresTransactionLogger) WriteDelete(key string) {
	l.events <- Event{EventType: EventDelete, Key: key}
}

func (l *PostgresTransactionLogger) Err() <-chan error {
	return l.errors
}

type PostgresDBParams struct {
	dbName   string
	host     string
	user     string
	password string
}

func NewPostgresTransactionLogger(config PostgresDBParams) (TransactionLogger, error) {

	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password=%s",
		config.host, config.dbName, config.user, config.password)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	logger := &PostgresTransactionLogger{db: db}

	/* exists, err := logger.verifyTableExists()
	if err != nil {
		return nil, fmt.Errorf("failed to verify table exists: %w", err)
	}
	if !exists {
		if err = logger.createTable(); err != nil {
			return nil, fmt.Errorf("failed to create table: %w", err)
		}
	} */

	return logger, nil
}

func (l *PostgresTransactionLogger) Run() {
	events := make(chan Event, 16)
	l.events = events

	errors := make(chan error, 1)
	l.errors = errors

	go func() {
		query := "INSERT INTO transactions (event_type, key, value) VALUES ($1, $2, $3)"

		for e := range events {
			_, err := l.db.Exec(
				query,
				e.EventType, e.Key, e.Value)

			if err != nil {
				errors <- err
			}
		}
	}()
}

func (l *PostgresTransactionLogger) ReadEvents() (<-chan Event, <-chan error) {
	outEvent := make(chan Event)
	outError := make(chan error, 1)

	go func() {
		defer close(outEvent)
		defer close(outError)

		query := "SELECT sequence, event_type, key, value FROM transactions ORDER BY sequence"

		rows, err := l.db.Query(query)
		if err != nil {
			outError <- fmt.Errorf("sql query error: %w", err)
			return
		}

		defer rows.Close()

		e := Event{}

		for rows.Next() {

			err = rows.Scan(
				&e.Sequence, &e.EventType,
				&e.Key, &e.Value)

			if err != nil {
				outError <- fmt.Errorf("error reading row: %w", err)
			}

			outEvent <- e
		}

		err = rows.Err()
		if err != nil {
			outError <- fmt.Errorf("transaction log read failure: %w", err)
		}
	}()

	return outEvent, outError
}

func InitializeSQLTransactionLog() error {
	var err error

	var config = PostgresDBParams{
		dbName:   "",
		host:     "",
		user:     "",
		password: "",
	}

	Logger, err = NewPostgresTransactionLogger(config)
	if err != nil {
		return fmt.Errorf("failed to create event logger: %w", err)
	}

	events, errors := Logger.ReadEvents()
	e, ok := Event{}, true

	for ok && err == nil {
		select {
		case err, ok = <-errors:
		case e, ok = <-events:
			switch e.EventType {
			case EventDelete:
				err = api.Delete(e.Key)
			case EventPut:
				err = api.Put(e.Key, e.Value)
			}
		}
	}

	Logger.Run()

	return err
}
