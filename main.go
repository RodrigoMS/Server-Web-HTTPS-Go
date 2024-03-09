package main // Declara o nome do pacote. Todo programa Go é feito de pacotes e o programa começa a rodar no pacote main.

import ( // Início do bloco de importação de pacotes.
	"crypto/tls" // Fornece funções para lidar com segurança TLS.
	"log"        // Fornece funções simples de registro de log.
	"net/http"   // Fornece funções para construir servidores HTTP.

	"github.com/gin-gonic/gin" // Gin é um framework web para construir APIs.
	"github.com/gorilla/csrf"  // Gorilla CSRF fornece proteção CSRF.
)

// defineRoutes define as rotas para o servidor
func defineRoutes(r *gin.Engine) { 
	r.GET("/", func(c *gin.Context) { // Define uma rota GET para a raiz ("/").
		// Retorna uma resposta JSON com uma mensagem e um token CSRF
		c.JSON(http.StatusOK, gin.H{"message": "GET request", "csrfToken": csrf.Token(c.Request)})
	})

	r.POST("/", csrfMiddleware, func(c *gin.Context) { // Define uma rota POST para a raiz ("/") com o middleware CSRF.
		// Retorna uma resposta JSON com uma mensagem
		c.JSON(http.StatusOK, gin.H{"message": "POST request"})
	})

	r.PUT("/", csrfMiddleware, func(c *gin.Context) { // Define uma rota PUT para a raiz ("/") com o middleware CSRF.
		// Retorna uma resposta JSON com uma mensagem
		c.JSON(http.StatusOK, gin.H{"message": "PUT request"})
	})

	r.DELETE("/", csrfMiddleware, func(c *gin.Context) { // Define uma rota DELETE para a raiz ("/") com o middleware CSRF.
		// Retorna uma resposta JSON com uma mensagem
		c.JSON(http.StatusOK, gin.H{"message": "DELETE request"})
	})

	r.PATCH("/", csrfMiddleware, func(c *gin.Context) { // Define uma rota PATCH para a raiz ("/") com o middleware CSRF.
		// Retorna uma resposta JSON com uma mensagem
		c.JSON(http.StatusOK, gin.H{"message": "PATCH request"})
	})
} 

// csrfMiddleware é um middleware que verifica a validade do token CSRF
func csrfMiddleware(c *gin.Context) {
	// Obtém o token CSRF do cabeçalho da solicitação
	csrfToken := c.GetHeader("X-CSRF-Token")
	// Se o token CSRF não estiver presente ou for inválido, a solicitação será rejeitada
	if csrfToken == "" || csrfToken != csrf.Token(c.Request) {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid CSRF token"})
		c.Abort() // Aborta a solicitação atual.
		return
	}
	c.Next() // Passa para o próximo middleware ou manipulador de rota.
}

func main() {
	// Configura o Gin para rodar em modo de produção
	gin.SetMode(gin.ReleaseMode)

	// Cria um novo roteador Gin
	r := gin.New()

	// Define as rotas
	defineRoutes(r)

	// Carrega o certificado e a chave TLS de arquivos locais
	certFile := "./certificate/localhost.crt"
	keyFile := "./certificate/localhost.key"

	cert, err := tls.LoadX509KeyPair(certFile, keyFile) // Carrega o par de chaves do certificado e do arquivo de chave.
	if err != nil { // Se houver um erro ao carregar o par de chaves...
		log.Printf("Erro ao carregar o par de chaves: %v", err) // Registra o erro.
		return // Termina a execução da função main.
	}

	// Cria um servidor HTTP personalizado
	server := &http.Server{
		Addr:    ":8080", // Define o endereço e a porta em que o servidor vai escutar.
		Handler: r, // Define o manipulador de solicitações HTTP.
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert}, // Define os certificados TLS.
		},
	} 

	err = server.ListenAndServeTLS(certFile, keyFile) // Inicia o servidor com suporte a TLS.
	if err != nil { // Se houver um erro ao iniciar o servidor...
		log.Printf("Erro ao iniciar o servidor: %v", err) // Registra o erro.
		return // Termina a execução da função main.
	}
}