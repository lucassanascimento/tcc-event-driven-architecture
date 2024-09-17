# RabbitMQ Example

## Como executar?
### Execute a infraestrutura necessária para usar o RabbitMQ:
```shell
docker compose up -d
```

### Executando a aplicação:
Com todas as ferramentas instaladas, em seu terminal, na pasta  `golang-example`, execute o comando:

```shell
go run .
```



---

# 🟡 RabbitMQ Example

Este tutorial guia você na execução de um exemplo de comunicação assíncrona utilizando **RabbitMQ**.

---

## 🛠️ Como Executar?

### 1️⃣ **Iniciar a Infraestrutura Necessária para o RabbitMQ**

Primeiro, suba a infraestrutura do RabbitMQ utilizando **Docker Compose**:

```bash
docker compose up -d
```

#### ✅ Validar que o container está rodando:

Verifique se o container foi criado com sucesso, listando todos os containers disponíveis:

```bash
docker ps
```

---

### 3️⃣ **Executar a Aplicação Golang**

Com o RabbitMQ executando utilizando o docker, execute a aplicação Go.

Entre no diretório (Ex: `golang-simple-example`) e execute o seguinte comando:

> 🐹 **Dica**: Certifique-se de que as dependências do Golang estão instaladas corretamente usando `go mod tidy` antes de rodar o projeto.

Agora, execute o aplicativo com o comando:

```bash
go run .
```

---

### 🎯 Pronto!

Agora sua aplicação está rodando e se comunicando com o **RabbitMQ** usando a fila `example-queue`.

Caso tudo esteja configurado, você poderá ver algo parecido com isso em seu console:
```bash
2024/09/15 16:32:08 INFO Message sent successfully
2024/09/15 16:32:08 INFO Consumer started. Waiting for messages...
2024/09/15 16:32:08 INFO Received message body="Hello RabbitMQ"
2024/09/15 16:32:10 INFO Message sent successfully
2024/09/15 16:32:10 INFO Received message body="Hello RabbitMQ"
2024/09/15 16:32:12 INFO Message sent successfully
2024/09/15 16:32:12 INFO Received message body="Hello RabbitMQ"
2024/09/15 16:32:14 INFO Message sent successfully
2024/09/15 16:32:14 INFO Received message body="Hello RabbitMQ"
2024/09/15 16:32:16 INFO Message sent successfully
2024/09/15 16:32:16 INFO Received message body="Hello RabbitMQ"
2024/09/15 16:32:18 INFO Message sent successfully
2024/09/15 16:32:18 INFO Received message body="Hello RabbitMQ"
```
---

### ℹ️ Dicas:

- Certifique-se de que o **Docker** e o **Golang** estão corretamente instalados antes de iniciar.
- O RabbitMQ estará rodando em `localhost:5672`. Se você alterar as configurações, ajuste os parâmetros nos comandos de criação de tópicos e na aplicação Go.
- Para testar a comunicação, você pode enviar e receber mensagens entre o **Producer** e **Consumer** do RabbitMQ, conforme definido no código.

---

### 📖 Links Úteis:

- **Documentação Oficial do RabbitMQ**: [https://www.rabbitmq.com/docs](https://www.rabbitmq.com/docs)

---

