package models

// As tags json:"..." dizem ao Go como os campos da struct devem aparecer no JSON,
// tanto ao converter para JSON quanto ao ler de JSON. Elas definem
// os nomes das chaves no JSON.
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

const (
	TableName        = "tasks"
	CreateTableQuery = `CREATE TABLE IF NOT EXISTS tasks (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		status BOOLEAN NOT NULL DEFAULT false
	)`
)
