# Kafka Example

## Como executar?
### Execute a infraestrutura necessária para usar o Kafka:
```shell
docker compose up -d
```

### Configuração do Tópico:

Acessando container:
```shell
docker exec -it kafka-like /bin/bash
```

Criando o tópico: 
```shell
kafka-topics.sh --create --topic example-topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
```

Valide que o tópico foi criado com sucesso:
```shell
kafka-topics.sh --list --bootstrap-server localhost:9092
```

### Executando a aplicação:
Com todas as ferramentas instaladas, em seu terminal, na pasta  `golang-example`, execute o comando:

```shell
go run .
```
