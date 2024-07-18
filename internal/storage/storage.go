package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

type Storage struct {
	db *sql.DB
}

func New(storagePath string) *Storage {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		log.Fatal("failed to open the database: %v", err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER PRIMARY KEY,
		title TEXT NO NULL UNIQUE,
		done BOOLEAN NO NULL);
	CREATE INDEX IF NOT EXISTS idx_done ON tasks(done);
	`)
	if err != nil {
		log.Fatal("failed to prepare the statement: %v", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal("failed to execute the statement: %v", err)
	}

	return &Storage{db: db}
}

func (s *Storage) GetAllTasks() ([]Task, error) {
	stmt, err := s.db.Prepare("SELECT * FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare the statement: %w", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("failed to execute the statement: %w", err)
	}
	if rows == nil {
		return nil, fmt.Errorf("failed to get the rows: %v", rows)
	}

	var tasks []Task
	for rows.Next() {
		var task Task

		err = rows.Scan(&task.ID, &task.Title, &task.Done)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *Storage) AddTask(title string, done bool) error {
	stmt, err := s.db.Prepare("INSERT INTO tasks(title, done) VALUES(?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare the statement: %w", err)
	}

	_, err = stmt.Exec(title, done)
	if err != nil {
		return fmt.Errorf("failed to execute the statement: %w", err)
	}

	return nil
}
