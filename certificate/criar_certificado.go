package main // Declara o nome do pacote. Todo programa Go é feito de pacotes e o programa começa a rodar no pacote main.

import (
	"log"    // Fornece funções simples de registro de log.
	"os/exec" // Fornece funções para executar comandos externos.
)

func main() { // Esta é a função principal que é executada quando o programa é iniciado.
	commands := []string{ // Declara uma variável chamada commands que é uma fatia de strings. Cada string na fatia é um comando que será executado.
		"openssl genrsa -out myCA.key 2048", // Gera uma chave RSA privada de 2048 bits e a salva em um arquivo chamado myCA.key.
		`openssl req -x509 -new -nodes -key myCA.key -sha256 -days 1825 -subj "/C=BR/ST=country/L=city/O=YourOrganization/OU=YourUnit/CN=yourdomain.com" -out myCA.pem`, // Gera um novo certificado autoassinado X.509 usando a chave privada do arquivo myCA.key. O certificado é válido por 1825 dias e é salvo em um arquivo chamado myCA.pem.
		"openssl genrsa -out localhost.key 2048", // Gera uma nova chave RSA privada de 2048 bits e a salva em um arquivo chamado localhost.key.
		`openssl req -new -key localhost.key -subj "/C=BR/ST=country/L=city/O=YourOrganization/OU=YourUnit/CN=localhost" -out localhost.csr`, // Gera uma nova solicitação de assinatura de certificado (CSR) usando a chave privada do arquivo localhost.key e a salva em um arquivo chamado localhost.csr.
		"openssl x509 -req -in localhost.csr -CA myCA.pem -CAkey myCA.key -CAcreateserial -out localhost.crt -days 1825 -sha256", // Gera um novo certificado X.509 baseado na CSR do arquivo localhost.csr, assinado pela CA do arquivo myCA.pem e myCA.key. O certificado é válido por 1825 dias e é salvo em um arquivo chamado localhost.crt.
	}

	for _, cmd := range commands { // Início de um loop que percorre cada comando na fatia commands.
		err := exec.Command("sh", "-c", cmd).Run() // Executa o comando atual do loop. Se houver um erro ao executar o comando, ele será retornado e armazenado na variável err.
		if err != nil { // Se houver um erro (ou seja, err não é nil), o programa registra o erro e termina.
			log.Fatalf("Falha ao executar o comando %s: %v", cmd, err)
		}
	} // Conclusão do loop for e da função main.
}

