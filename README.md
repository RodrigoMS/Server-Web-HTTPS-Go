# Servidor Web HTTPS Go

Este projeto é um servidor web escrito em Go que utiliza o framework Gin para construir APIs. Ele implementa várias medidas de segurança para garantir a integridade e a segurança dos dados.

## Executando o Servidor

1. **Inicie o servidor**: Execute o programa Go fornecido. Certifique-se de que os arquivos de certificado e chave estão no local correto (`"./certificate/localhost.crt"` e `"./certificate/localhost.key"`).

#### Pré-requisitos

Antes de iniciar o servidor, é necessário ter o OpenSSL instalado em seu sistema. O OpenSSL é usado para criar os certificados necessários para o servidor HTTPS. Se você não tiver o OpenSSL instalado, pode baixá-lo e instalá-lo a partir do [site oficial do OpenSSL](https://www.openssl.org/source/).

Há um programa para criar os arquivos `"./certificate/localhost.crt"` e `"./certificate/localhost.key"`, navegue até o diretório **certificate** que contém o arquivo **criar_certificado.go** e execute o comando `go run criar_certificado.go`. Isso criará os arquivos de certificado na pasta certificate.

### Compilação

Para iniciar o servidor, execute o comando `go run main.go` ou `build main.go` no diretório raiz do projeto. O servidor começará a escutar no endereço e na porta especificados no código.

- Cria um executável na pasta temp do sistema operacional e o coloca em execução.
```
go run main.go
```

- Cria um executável dentro da pasta do projeto, sendo necessário executar após o fim da compilação.
```
go mod init example.com
```
```
go mod tidy
```
```
build main.go
```
**Linux**
```
./main
```
ou **Windows**
```
main.exe
```

### Acessando a Página HTTPS

1. **Abra um navegador da web**: Use o navegador de sua escolha.
2. **Digite o endereço do servidor**: Na barra de endereços do navegador, digite `https://localhost:8080` e pressione Enter.
3. **Aceite o aviso de certificado**: Como estamos usando um certificado autoassinado, você verá um aviso sobre o certificado do servidor não ser confiável. Você pode optar por continuar mesmo assim para acessar a página.

### Enviando Solicitações

Agora você pode enviar solicitações GET, POST, PUT, DELETE e PATCH para a raiz ("/"). As respostas serão retornadas como JSON. Para solicitações POST, PUT, DELETE e PATCH, você precisará fornecer um token CSRF válido no cabeçalho "X-CSRF-Token".

Por favor, note que este é um servidor de exemplo e não deve ser usado em produção sem a devida consideração para segurança e conformidade. Além disso, o uso de certificados autoassinados pode apresentar riscos de segurança e deve ser evitado em um ambiente de produção. Para produção, você deve usar um certificado emitido por uma autoridade de certificação confiável.

Entendi, você gostaria de adicionar o certificado ao seu navegador para que ele reconheça o certificado como válido. Aqui estão as etapas gerais para adicionar um certificado autoassinado aos navegadores mais populares:

## Aceitando e Registrando o Certificado

### Chrome

1. Abra o Chrome e acesse `chrome://settings`.
2. Role para baixo e clique em "Avançado".
3. Em "Privacidade e segurança", clique em "Gerenciar certificados".
4. Na guia "Autoridades", clique em "Importar".
5. Navegue até o arquivo `.pem` e clique em "Abrir".
6. Marque a caixa "Confiar neste certificado para identificar sites" e clique em "OK".

### Edge

1. Abra o Edge e clique nos três pontos no canto superior direito.
2. Selecione "Configurações" > "Privacidade, pesquisa e serviços".
3. Em "Segurança", clique em "Gerenciar certificados".
4. Na guia "Autoridades de Certificação Raiz Confiáveis", clique em "Importar".
5. Siga o assistente para importar o arquivo `.pem`.

### Safari

1. Abra o Safari e clique em "Safari" na barra de menu, depois em "Preferências".
2. Vá para a guia "Privacidade" e clique em "Gerenciar certificados do site".
3. Clique no sinal de mais (+) e navegue até o arquivo `.pem`.
4. Selecione o arquivo e clique em "Adicionar".

### Firefox

1. Abra o Firefox e clique nos três traços no canto superior direito.
2. Selecione "Opções" > "Privacidade e Segurança".
3. Role para baixo até "Certificados" e clique em "Ver certificados".
4. Na guia "Autoridades", clique em "Importar".
5. Navegue até o arquivo `.pem` e clique em "Abrir".
6. Marque a caixa "Confiar neste CA para identificar sites" e clique em "OK".

**Por favor, note que essas instruções podem variar dependendo da versão do seu navegador. Além disso, o uso de certificados autoassinados pode apresentar riscos de segurança e deve ser evitado em um ambiente de produção. Para produção, você deve usar um certificado emitido por uma autoridade de certificação confiável.**

## Segurança

### TLS (Transport Layer Security)

O servidor suporta conexões seguras TLS. Ele carrega um par de chaves de certificado e chave de um arquivo local, garantindo que todas as comunicações entre o cliente e o servidor sejam criptografadas e seguras contra interceptação.

### Proteção CSRF (Cross-Site Request Forgery)

O servidor usa o pacote Gorilla CSRF para fornecer proteção contra ataques CSRF. Ele verifica a validade do token CSRF em cada solicitação recebida. Se o token CSRF não estiver presente ou for inválido, a solicitação será rejeitada.

### Rotas Protegidas

Todas as rotas POST, PUT, DELETE e PATCH são protegidas pelo middleware CSRF, garantindo que apenas solicitações legítimas possam realizar essas ações.

### HTTP/2 (Performance)

HTTP/2 é a segunda versão principal do protocolo HTTP (Hyper Text Transfer Protocol), que forma a base da World Wide Web. Ele foi desenvolvido para resolver alguns problemas que existiam na versão 1.1, principalmente com relação à performance.

### Modo de Produção

O servidor é configurado para rodar no modo de produção do Gin, o que significa que ele não exibirá mensagens de erro detalhadas que poderiam potencialmente expor informações sensíveis.


## Informações adicionais

- **Roteamento:** O servidor define várias rotas (GET, POST, PUT, DELETE, PATCH) na função defineRoutes. Cada rota responde com uma mensagem JSON indicando o tipo de solicitação recebida.

- **Proteção CSRF:** O servidor usa tokens CSRF (Cross-Site Request Forgery) para proteger contra ataques CSRF. Um token CSRF é gerado para cada solicitação GET e é necessário para solicitações POST, PUT, DELETE e PATCH. Se o token CSRF não estiver presente ou for inválido, a solicitação será rejeitada.

- **Modo de Produção:** O servidor é executado no modo de produção (gin.ReleaseMode), o que significa que a saída de log é minimizada.

- **Servidor HTTPS:** O servidor é configurado para usar HTTPS, o que significa que todas as comunicações entre o cliente e o servidor são criptografadas. O servidor carrega um certificado e uma chave privada de arquivos locais e usa esses para configurar a conexão TLS.

## Sugestões de melhorias de segurança do servidor

- **Gerenciamento de Erros:** O programa atualmente termina com log.Fatal ou panic se encontrar um erro ao carregar o par de chaves ou iniciar o servidor. Em vez disso, você pode querer lidar com esses erros de uma maneira que seja mais adequada para o seu aplicativo.

- **Certificado e Chave TLS:** O certificado e a chave estão sendo carregados de arquivos locais. Em um ambiente de produção, você pode querer usar um serviço de gerenciamento de segredos para armazenar e recuperar esses.

- **Política de Segurança de Conteúdo (CSP):** Implementar uma Política de Segurança de Conteúdo pode ajudar a prevenir vários tipos de ataques, como Cross Site Scripting (XSS).

- **Atualizações de Segurança:** Certifique-se de manter todas as suas dependências atualizadas para garantir que você esteja protegido contra quaisquer vulnerabilidades conhecidas.

- **Limitar o Tamanho do Corpo da Solicitação:** Para proteger contra ataques de negação de serviço (DoS), você pode querer limitar o tamanho do corpo da solicitação.

- **Autenticação e Autorização:** Se o seu servidor estiver lidando com dados sensíveis, você pode querer implementar algum tipo de autenticação e autorização.

**Lembre-se, a segurança é um campo complexo e em constante evolução, e estas são apenas algumas sugestões gerais. Recomenda-se trabalhar com um especialista em segurança para garantir que seu aplicativo esteja adequadamente protegido.**