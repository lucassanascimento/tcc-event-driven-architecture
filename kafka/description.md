
---

# 🟡 Kafka Example

Este tutorial guia você na execução de um exemplo de comunicação assíncrona utilizando **Apache Kafka**.

---

## 🛠️ Como Executar?

### 1️⃣ **Iniciar a Infraestrutura Necessária para o Kafka**

Primeiro, suba a infraestrutura do Kafka utilizando **Docker Compose**:

```bash
docker compose up -d
```

### 2️⃣ **Configurar o Tópico Kafka**

Agora, configure o tópico no Kafka para onde as mensagens serão enviadas e consumidas.

#### 🔗 Acessar o Container do Kafka

Entre no container do Kafka usando o comando abaixo:

```bash
docker exec -it kafka-kafka-1 /bin/bash
```

#### 📝 Criar o Tópico

Crie um tópico chamado `example-topic` com a configuração básica de partições e fator de replicação:

```bash
kafka-topics.sh --create --topic example-topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
```

#### ✅ Validar a Criação do Tópico

Verifique se o tópico foi criado com sucesso, listando os tópicos disponíveis no Kafka:

```bash
kafka-topics.sh --list --bootstrap-server localhost:9092
```

---

### 3️⃣ **Executar a Aplicação Golang**

Com o Kafka e o tópico configurados, execute a aplicação Go.

Entre no diretório `golang-example` e execute o seguinte comando:

> 🐹 **Dica**: Certifique-se de que as dependências do Golang estão instaladas corretamente usando `go mod tidy` antes de rodar o projeto.

Agora, execute o aplicativo com o comando:

```bash
go run .
```

---

### 🎯 Pronto!

Agora sua aplicação está rodando e se comunicando com o **Kafka** usando o tópico `example-topic`.

Caso tudo esteja configurado, você poderá ver algo parecido com isso em seu console:
```bash
2024/09/15 16:14:28 INFO Consumer started. Waiting for messages...
2024/09/15 16:14:29 INFO Message sent successfully
2024/09/15 16:14:29 INFO Received message: key=Key-A value="Hello Kafka"
2024/09/15 16:14:32 INFO Message sent successfully
2024/09/15 16:14:32 INFO Received message: key=Key-A value="Hello Kafka"
2024/09/15 16:14:35 INFO Message sent successfully
2024/09/15 16:14:35 INFO Received message: key=Key-A value="Hello Kafka"
2024/09/15 16:14:38 INFO Message sent successfully
2024/09/15 16:14:38 INFO Received message: key=Key-A value="Hello Kafka"
2024/09/15 16:14:41 INFO Message sent successfully
2024/09/15 16:14:41 INFO Received message: key=Key-A value="Hello Kafka"
```
---

### ℹ️ Dicas:

- Certifique-se de que o **Docker** e o **Golang** estão corretamente instalados antes de iniciar.
- O Kafka estará rodando em `localhost:9093`. Se você alterar as configurações, ajuste os parâmetros nos comandos de criação de tópicos e na aplicação Go.
- Para testar a comunicação, você pode enviar e receber mensagens entre o **Producer** e **Consumer** do Kafka, conforme definido no código.

---

### 📖 Links Úteis:

- **Documentação Oficial do Kafka**: [https://kafka.apache.org/documentation](https://kafka.apache.org/documentation)
---
