# FullCycle - Pós Go Expert Desafio - Client-Server-API

## Pré requisito

* Ter o "go 1.20" instalado e configurado 
* Fazer o clone desse repositório

## Servidor
* Para executar o servidor execute esse comandos:
    * [user@local]$ cd ./goexpert-fullcycle-client-server-api/cmd/http-server
    * [user@local]$ go run main.go &
* Arquivo do banco de dados SqlLite:
    * ./goexpert-fullcycle-client-server-api/cmd/http-server/server.db
* Arquivo de log da aplicação:
    * ./goexpert-fullcycle-client-server-api/cmd/http-server/server.log
* Arquivo do código fonte:
    * [./goexpert-fullcycle-client-server-api/cmd/http-server/main.go](./cmd/http-server/main.go)

Atenção: os arquivos server.log e server.db serão criados na primeira execução da aplicação.

## Cliente
* Para executar o servidor execute esse comandos:
    * [user@local]$ cd ./goexpert-fullcycle-client-server-api/cmd/http-client
    * [user@local]$ go run main.go &
* Arquivo com as cotações realizadas:
    * ./goexpert-fullcycle-client-server-api/cmd/http-client/cotacao.txt
* Arquivo de log da aplicação:
    * ./goexpert-fullcycle-client-server-api/cmd/http-client/client.log
* Arquivo do código fonte:
    * [./goexpert-fullcycle-client-server-api/cmd/http-client/main.go](./cmd/http-client/main.go)

Atenção: os arquivos client.log e cotacao.txt serão criados na primeira execução da aplicação.