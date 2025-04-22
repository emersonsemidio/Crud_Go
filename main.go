package main

// importando o pacote "config" para acessar as funções definidas nele
import (
	"go-api/config"
	"go-api/handlers"
	"go-api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Função chamada quando a aplicação é iniciada. É o ponto de entrada da aplicação
func main() {
	// Chama a função SetupDataBase do pacote config
	dbConnection := config.SetupDataBase()

	/* As funções em Go podem retornar mais de um valor,
	abaixo, retorna 2 e pedi para ignorar o 1°	*/
	_, err := dbConnection.Exec(models.CreateTableQuery)

	if err != nil {
		log.Fatal(err)
	}

	// Chama a função Close para fechar a conexão com o banco de dados
	// O defer garante que a função Close será chamada quando a função main terminar
	defer dbConnection.Close()

	router := mux.NewRouter()

	taskHandler := handlers.NewTaskHandler(dbConnection)

	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
