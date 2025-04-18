package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // Importa o driver do MySQL para o Go
	"github.com/joho/godotenv"
)

// As funções devem ser declaradas com Letras maiusculas para que possam ser acessadas de outros pacotes
// O *sql.DB é um ponteiro para a estrutura sql.DB, que representa uma conexão com o banco de dados
// também é o tipo de retorno da função SetupDataBase
// Pega as variáveis de ambiente do arquivo .env
// Inicializa a conexão com o BD
func SetupDataBase() *sql.DB {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	fmt.Println("Connection String", connectionString)

	dbConnection, err := sql.Open("mysql", connectionString)

	if err != nil {
		fmt.Println(err)
	}

	err = dbConnection.Ping()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso")

	return dbConnection
}
