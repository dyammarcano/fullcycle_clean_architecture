[![CI and Test](https://github.com/dyammarcano/fullcycle_clean_architecture/actions/workflows/ci.yml/badge.svg)](https://github.com/dyammarcano/fullcycle_clean_architecture/actions/workflows/ci.yml)

# Desafio Clean Architecture

Olá devs!
Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:

- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
  Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando
docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada
serviço.

# Requisitos

- [Docker](https://www.docker.com/)
- [Golang](https://golang.org/) (Opcional)

# Arquitetura do projeto

![img.png](img.png)

```text
github.com/dyammarcano/fullcycle_clean_architecture
├───cmd
├───internal
│   ├───adapter
│   │   ├───grpc
│   │   └───http
│   ├───domain
│   ├───repository
│   │   └───migrations
│   └───usecase
└───pkg
│   ├───config
│   ├───grpc
│   │   ├───pb
│   │   └───proto
│   ├───logger
│   └───util
├── main.go
├── Dockerfile
├── docker-compose.yaml
└── README.md
```

# Executando o projeto

1. Clone o repositório

```bash
$ git clone github.com/dyammarcano/fullcycle_clean_architecture.git
```

2. Acesse a pasta do projeto

```bash
$ cd fullcycle_clean_architecture
```

3. Executar o projeto

```bash
# Comando para gerar as images e subir os containers com o banco de dados
$  docker-compose up --build
```

4. Acesse o endereço `http://localhost:8080/` para acessar a interface do GraphQL Playground, e o server http
5. Acesse o endereço `http://localhost:8081/` para acessar a interface do gRPC 

# Endpoints

- GraphQL Playground: `http://localhost:8080/`
- gRPC: `http://localhost:8081/`
- REST: `http://localhost:8080/order`

# gRPC Client

Para testar o gRPC, você pode usar o [evans](https://github.com/ktr0731/evans)

```bash
$ evans -r -p 8081 repl
```

```bash
127.0.0.1:8081> show package
+-------------------------+
|         PACKAGE         |
+-------------------------+
| fullcycle               |
| grpc.reflection.v1      |
| grpc.reflection.v1alpha |
+-------------------------+
```

```bash
127.0.0.1:8081> package fullcycle
```

```bash
fullcycle@127.0.0.1:8081> show service
+--------------+------------+-------------------+--------------------+
|   SERVICE    |    RPC     |   REQUEST TYPE    |   RESPONSE TYPE    |
+--------------+------------+-------------------+--------------------+
| OrderService | ListOrders | ListOrdersRequest | ListOrdersResponse |
+--------------+------------+-------------------+--------------------+
```

```bash
fullcycle@127.0.0.1:8081> show message
+--------------------+
|      MESSAGE       |
+--------------------+
| ListOrdersRequest  |
| ListOrdersResponse |
+--------------------+
```

```bash
fullcycle@127.0.0.1:8081> desc ListOrdersRequest
+-------+------+----------+
| FIELD | TYPE | REPEATED |
+-------+------+----------+
+-------+------+----------+
```

```bash
fullcycle@127.0.0.1:8081> desc ListOrdersResponse
+--------+----------------------+----------+
| FIELD  |         TYPE         | REPEATED |
+--------+----------------------+----------+
| orders | TYPE_MESSAGE (Order) | true     |
+--------+----------------------+----------+
```

```bash
fullcycle@127.0.0.1:8081> service OrderService
```

```bash
fullcycle.OrderService@127.0.0.1:8081> call ListOrders {}
```
