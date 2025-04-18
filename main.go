package main

// importando o pacote "config" para acessar as funções definidas nele
import (
	"go-api/config"
	"log"
	"net/http"
)

// Função chamada quando a aplicação é iniciada. É o ponto de entrada da aplicação
func main() {
	// Chama a função SetupDataBase do pacote config
	dbConnection := config.SetupDataBase()

	// Chama a função Close para fechar a conexão com o banco de dados
	// O defer garante que a função Close será chamada quando a função main terminar
	defer dbConnection.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))

}
