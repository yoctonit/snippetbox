package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/yoctonit/snippetbox/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:Password1@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// next: 4.8 Multiple-record SQL queries

// -- Create a new UTF-8 `snippetbox` database.
// CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

// -- Switch to using the `snippetbox` database.
// USE snippetbox;

// -- Create a `snippets` table.
// CREATE TABLE snippets (
//     id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
//     title VARCHAR(100) NOT NULL,
//     content TEXT NOT NULL,
//     created DATETIME NOT NULL,
//     expires DATETIME NOT NULL
// );

// -- Add an index on the created column.
// CREATE INDEX idx_snippets_created ON snippets(created);

// -- Add some dummy records (which we'll use in the next couple of chapters).
// INSERT INTO snippets (title, content, created, expires) VALUES (
//     'An old silent pond',
//     'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
//     UTC_TIMESTAMP(),
//     DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
// );

// INSERT INTO snippets (title, content, created, expires) VALUES (
//     'Over the wintry forest',
//     'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
//     UTC_TIMESTAMP(),
//     DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
// );

// INSERT INTO snippets (title, content, created, expires) VALUES (
//     'First autumn morning',
//     'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
//     UTC_TIMESTAMP(),
//     DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
// );

// CREATE USER 'web'@'localhost';
// GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'localhost';
// -- Important: Make sure to swap 'pass' with a password of your own choosing.
// ALTER USER 'web'@'localhost' IDENTIFIED BY 'Password1';
