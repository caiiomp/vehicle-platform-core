# vehicle-platform-core
Este repositório contém a API para gestão de veículos, permitindo a criação e atualização dos mesmos.

## Funcionalidades

- **Cadastro de veículos:** Permite o cadastro de veículos à venda (marca, modelo, ano, cor e preço).
- **Edição de veículos:** Permite a edição dos dados de veículos cadastrados.

## Tecnologias Utilizadas

- **Go (Golang):** Para o desenvolvimento da API de veículos.
- **MongoDB:** Para o armazenamento dos dados de veículos.
- **Gin:** Framework web para o desenvolvimento da API.
- **Docker Compose:** Para o setup do serviço e suas dependências via Docker.

## Como Rodar o Projeto Localmente

### 1. Pré-requisitos

Certifique-se de que você tem as seguintes dependências instaladas:

- **Go (Golang)** versão 1.18 ou superior
- **Git** para clonar o repositório
- **Docker** e **Docker Compose**

### 2. Configuração para rodar o serviço localmente com Docker Compose

1. Clone o repositório:

    ```bash
    git clone git@github.com:caiiomp/vehicle-platform-core.git
    ```

2. Na raiz do projeto instale as dependências do Go:

    ```bash
    go mod tidy
    ```

3. Na raiz do projeto, inicie o serviço e suas dependências `docker`:

    ```bash
    # Cria uma rede compartilhada do docker para a integração entre os serviços
    docker network create shared_network

    # Faz o build das nossas imagens
    docker compose build

    # Sobe os serviços
    docker compose up -d
    ```

### 3. Testando o serviço

Use **Postman**, **Insomnia**, **cURL** ou qualquer outro cliente **HTTP** para testar os endpoints:

- `POST /vehicles` - Cadastrar um novo veículo
- `PATCH /vehicles/:entity_id` - Editar um veículo existente

Os testes unitários e os testes de integração podem ser executados da seguinte forma respectivamente:
```bash
    go test ./... -v
    go test -tags=integration -v ./...
```

## Documentação (Swagger)

Para acessar a documentação do serviço, acessar o seguinte endpoint: 
```
http://localhost:4001/swagger/index.html
```