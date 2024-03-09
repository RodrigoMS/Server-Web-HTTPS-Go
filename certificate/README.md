Os arquivos gerados pelo código têm os seguintes propósitos:

- **myCA.key:** Esta é a chave privada da sua Autoridade Certificadora (CA). Ela é usada para assinar certificados. Deve ser mantida em segredo.
- **myCA.pem:** Este é o certificado da sua Autoridade Certificadora (CA). Ele contém a chave pública correspondente à chave privada myCA.key. Este certificado é usado para verificar se um certificado foi assinado por sua CA.
- **localhost.key:** Esta é a chave privada para o seu servidor "localhost". Ela é usada para criar uma Solicitação de Assinatura de Certificado (CSR) e para estabelecer conexões seguras com o servidor.
- **localhost.csr:** Esta é a Solicitação de Assinatura de Certificado (CSR) para o seu servidor "localhost". Ela contém informações sobre o servidor (como o nome do domínio) e a chave pública do servidor. A CSR é enviada para a CA para solicitar um certificado.
- **localhost.crt:** Este é o certificado para o seu servidor "localhost", assinado pela sua CA. Ele contém a chave pública do servidor e informações sobre o servidor. Os clientes podem verificar a assinatura deste certificado usando o certificado da CA para confirmar que o certificado é confiável.

Por favor, note que você deve manter as chaves privadas (**myCA.key** e **localhost.key**) em segurança para evitar qualquer acesso não autorizado. Se alguém obtiver acesso à sua chave privada, eles poderiam potencialmente se passar por seu servidor ou CA.

**Gerar Certificado válido HTTPS**

Unidade certificadora gratuita - https://letsencrypt.org/pt-br
OBS: Apenas quando já tiver o domínio público para a aplicação ou site.


**Exemplo adicional**

Estes comandos geram um certificado autoassinado e a chave correspondente que são válidos por 365 dias.

openssl genpkey -algorithm RSA -out chave_privada.key

openssl req -new -x509 -key chave_privada.key -out certificado.crt -days 365


