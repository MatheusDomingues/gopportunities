package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o Router utilizando as configurações Default do gin
	router := gin.Default()
	// Definindo a rota
	initializeRoutes(router)
	// Estamos rodando a nossa api
	router.Run(":8080")
}
