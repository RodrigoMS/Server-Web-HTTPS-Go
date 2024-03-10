package main

// Importando os pacotes necessários
import (
	"crypto/tls" // Pacote para lidar com segurança TLS
	"log" // Pacote para registrar mensagens
	"net/http" // Pacote para lidar com solicitações HTTP
	//"os" // Pacote para lidar com operações do sistema operacional

	"github.com/gin-gonic/gin" // Framework web Gin
	"github.com/gorilla/csrf" // Pacote para lidar com tokens CSRF
	"github.com/gorilla/sessions" // Pacote para lidar com sessões
)

// Recuperando a chave de autenticação do ambiente (Recomendado)
//var AUTH_KEY = os.Getenv("AUTH_KEY")
var AUTH_KEY = "y3lYZf54jx2o6M5E93utK0RH5x3Nh/LU8AYtw8WVqdI=" 

// Inicializando o armazenamento de sessões com a chave de autenticação
var store = sessions.NewCookieStore([]byte(AUTH_KEY))

// Função para lidar com solicitações POST, PUT, DELETE e PATCH
func handleRequest(c *gin.Context) {
	// Recuperando a sessão
	session, err := store.Get(c.Request, "session-name")
	// Se houver um erro, retorne um erro interno do servidor
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get session"})
		c.Abort()
		return
	}

	// Definindo o valor CSRF na sessão
	session.Values["csrf"] = csrf.Token(c.Request)
	// Salvando a sessão
	err = session.Save(c.Request, c.Writer)
	// Se houver um erro, retorne um erro interno do servidor
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save session"})
		c.Abort()
		return
	}

	// Retornando uma resposta de sucesso com o método da solicitação e o token CSRF
	c.JSON(http.StatusOK, gin.H{"message": c.Request.Method + " request", "csrfToken": session.Values["csrf"]})
}

// Função para lidar com solicitações GET
func handleGet(c *gin.Context) {
	// Recuperando a sessão
	session, _ := store.Get(c.Request, "session-name")

	// Definindo o valor CSRF na sessão
	session.Values["csrf"] = csrf.Token(c.Request)

	// Salvando a sessão
	err := session.Save(c.Request, c.Writer)
	// Se houver um erro, retorne um erro interno do servidor
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save session"})
		c.Abort()
		return
	}

	// Retornando uma resposta de sucesso com o método da solicitação e o token CSRF
	c.JSON(http.StatusOK, gin.H{"message": c.Request.Method + " request", "csrfToken": session.Values["csrf"]})
	return
}

// Função para definir as rotas do servidor
func defineRoutes(r *gin.Engine) {
	r.GET("/", handleGet) // Rota GET para a raiz
	r.POST("/", csrfMiddleware, handleRequest) // Rota POST para a raiz com middleware CSRF
	r.PUT("/", csrfMiddleware, handleRequest) // Rota PUT para a raiz com middleware CSRF
	r.DELETE("/", csrfMiddleware, handleRequest) // Rota DELETE para a raiz com middleware CSRF
	r.PATCH("/", csrfMiddleware, handleRequest) // Rota PATCH para a raiz com middleware CSRF
}

// Middleware para lidar com tokens CSRF
func csrfMiddleware(c *gin.Context) {
	// Recuperando a sessão
	session, err := store.Get(c.Request, "session-name")
	// Se houver um erro, registre e retorne um erro interno do servidor
	if err != nil {
		log.Printf("Failed to get session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get session"})
		c.Abort()
		return
	}

	// Recuperando o token CSRF do cabeçalho
	csrfToken := c.GetHeader("X-CSRF-Token")
	// Se o token CSRF for inválido, retorne um erro de proibido
	if csrfToken == "" || csrfToken != session.Values["csrf"] {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid CSRF token"})
		c.Abort()
		return
	}

	// Continue para o próximo middleware ou manipulador
	c.Next()
}

// Função principal
func main() {
	// Configurando o modo de lançamento para o Gin
	gin.SetMode(gin.ReleaseMode)
	// Inicializando uma nova instância do Gin
	r := gin.New()

	// Inicializando o middleware CSRF
	csrfMiddleware := csrf.Protect([]byte(AUTH_KEY))

	// Adicionando o middleware CSRF ao Gin
	r.Use(func(c *gin.Context) {
		csrfMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
		})).ServeHTTP(c.Writer, c.Request)
		c.Next()
	})

	// Definindo as rotas
	defineRoutes(r)

	// Caminhos para o certificado e a chave
	certFile := "../certificate/localhost.crt"
	keyFile := "../certificate/localhost.key"

	// Carregando o par de chaves
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	// Se houver um erro, registre e retorne
	if err != nil {
		log.Printf("Erro ao carregar o par de chaves: %v", err)
		return
	}

	// Configurando o servidor
	server := &http.Server{
		Addr:    ":8080", // Endereço e porta do servidor
		Handler: r, // Manipulador para lidar com as solicitações
		TLSConfig: &tls.Config{ // Configuração TLS
			Certificates: []tls.Certificate{cert}, // Certificados TLS
			NextProtos:   []string{"h2", "http/1.1"}, // Protocolos suportados
		},
	}

	// Iniciando o servidor
	err = server.ListenAndServeTLS(certFile, keyFile)
	// Se houver um erro, registre e retorne
	if err != nil {
		log.Printf("Erro ao iniciar o servidor: %v", err)
		return
	}
}