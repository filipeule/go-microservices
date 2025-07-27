# GO MICROSERVICES

Este projeto foi desenvolvido como parte do curso **Working with Microservices in Go (Golang)**, ministrado pelo professor **Trevor Sawler**.

O objetivo é estudar a construção de microsserviços em Go, criando serviços independentes, de baixa acoplagem, que se comunicam entre si e com um front-end via API REST. A comunicação entre os microsserviços é feita utilizando **JSON**, **RPC**, **gRPC**, e também via **mensageria AMQP** com **RabbitMQ**.

---

## Microsserviços desenvolvidos

- **Front-end**: interface simples para testes de funcionalidades  
- **Broker**: ponto de entrada principal para o cluster de microsserviços  
- **Serviço de Autenticação**: usa **PostgreSQL** para autenticar usuários  
- **Serviço de Logs**: registra atividades em um banco **MongoDB**  
- **Listener**: escuta mensagens no **RabbitMQ** e executa ações baseadas no payload  
- **Serviço de E-mail**: recebe payloads em JSON, formata e envia e-mails  

---

## Como executar o projeto

> Pré-requisitos: [Docker](https://www.docker.com/) e [Make](https://www.gnu.org/software/make/) instalados.

### Passos:

1. Clone este repositório
2. Na pasta *project*, execute o comando:
* No Linux:
    ```bash
    make up_build
* No Windows
    ```bash
    make -f Makefile.windows up_build