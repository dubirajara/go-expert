## Desafio Clean-Architercture

 Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço

## Paso a paso

1. **Iniciar a aplicação com Docker Compose**

    ```bash
    docker compose up -d
    ```

    Deve esperar que levante todos os serviços:
    - Build binario go multistage e inicia a aplicação
    - RabbitMQ
    - MysqlDB
    - Executa migrações


## Testando a aplicação

### API REST:

- **Criar order**: execute o arquivo `/api/create_order.http`.
- **Listar orders**: execute o arquivo `/api/list_order.http`.


### GRPC:

```bash
evans --proto internal/infra/grpc/protofiles/order.proto repl
```


### GraphQL
Abre o playground de GraphQL em [http://localhost:8080/](http://localhost:8080/).
 - **Criar order**:
    ```
    mutation createOrder {
        createOrder(input: {id:"42", Price: 250.24, Tax: 12.2}){
            id,
            Price,
            Tax,
            FinalPrice
        }
    }
    ```
- **Listar orders**:
    ```
    query listOrder {
        orders(input: {page: 0, pageSize: 100}) {
            id,
            Price,
            Tax,
            FinalPrice 
        }
    }
    ```