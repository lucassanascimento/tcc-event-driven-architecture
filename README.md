# Event-driven architecture: Mapeamento sistemático da Literatura
Este repositório tem como propósito disponibilizar tutoriais sobre a implementação de ferramentas utilizadas na comunicação entre microsserviços.

## Ferramentas necessárias:

* Docker e Docker Compose (para instalar, visite o site oficial: https://docs.docker.com/compose/install/);
* Golang (Acesse a documentação oficial para instalação: https://go.dev/doc/install);

## Arquitetura base apresentadas nos exemplos:

### Kafka Example
![architecture](./kafka-architecture.png)

### RabbitMQ Example
![architecture](./rabbitMQ-architecture.png)

## Estrutura do Repositório:
```
.
├── kafka
│   ├── compose.yml
│   ├── description.md
│   ├── golang-example
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── main.go
│   └── kafka.excalidraw.json
├── kafka-architecture.png
├── LICENSE
├── rabbitMQ
│   ├── compose.yml
│   ├── description.md
│   ├── golang-example
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── main.go
│   └── rabbitMQ.excalidraw.json
├── rabbitMQ-architecture.png
└── README.md
```