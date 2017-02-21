package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Task is a struct containing Task data
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TaskCollection is collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		err := rows.Scan(&task.ID, &task.Name)

		if err != nil {
			panic(err)
		}
		result.Tasks = append(result.Tasks, task)
	}
	return result
}

func PutTask(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO tasks(name) VALUES(?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
