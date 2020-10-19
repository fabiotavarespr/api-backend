# Documentação

## Sobre

Este projeto implementa um backend escrito em GO.

## Executando o projeto com Docker/Docker-Compose

### Iniciar o projeto

No raiz do projeto deve executar o seguinte comando:

```sh
docker-compose up
```

ou 

```sh
make docker-up
```

O comando criar a imagem do backend com o código compilado e inicializa o banco de dados.

### Para o projeto

```sh
docker-compose down
```

ou 

```sh
make docker-down
```

## Executando o projeto com GO

### Iniciar o banco de dados

```sh
make database-up
```

ou

```sh
docker run --rm --name db-backend -d -p 3306:3306 -v ./.sql-scripts/:/docker-entrypoint-initdb.d/ -e MYSQL_DATABASE=backend -e MYSQL_ROOT_PASSWORD=passwd123 mysql:8.0.21
```

### Parar o banco de dados

```sh
make database-down
```

ou

```sh
docker stop db-backend
```

### Iniciar o projeto

```sh
go run main.go --log-level=debug --server-hostname=0.0.0.0 --server-port=8081 --database-hostname=localhost --database-username=root --database-password=passwd123 --database-name=backend
```

## Variáveis de ambiente

| Variáveis              | Padrão       | Descrição                                       |
|:----------------------:|:------------:|:-----------------------------------------------:|
| log-level              | info         | debug, info, warning, error                     |
| server-hostname        | 0.0.0.0      | IP do servidor backend                          |
| server-port            | 8080         | Porta do servidor backend                       |
| database-hostname      | 127.0.0.1    | IP do servidor de banco de dados                |
| database-port          | 3306         | Porta do servidor de banco de dados             |
| database-username      | root         | Usuário do servidor de banco de dados           |
| database-password      | passwd123    | Senha do servidor de banco de dados             |
| database-name          | backend      | Nome da base do servidor de banco de dados      |

## Examplos de requisições

### Criando uma conta

```sh
curl -X POST localhost:8080/accounts -d '{"document_number": "999999999"}' -H 'Content-Type: application/json'
```

### Consultando uma conta

```sh
curl -X GET localhost:8080/accounts/1
```

### Criando uma transação

```sh
curl -X POST localhost:8080/transactions -d '{"account_id": 1, "operation_type_id": 4, "amount": 100.00}' -H 'Content-Type: application/json'
```