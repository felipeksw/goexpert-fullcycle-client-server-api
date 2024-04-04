package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
	"os"	
)

type exchangeReq struct {
	Bid string `json:"bid"`
}

func main() {

	// criação do arquivo de Log
	file, err := os.OpenFile("client.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("Log inicializado")

	// Recupera a cotação do dólar
	exchange, err := geExchange()
	if err != nil {
		log.Fatal(err)
	}

	// Grava em arquivo a cotação
	err = addRecord(exchange)
	if err != nil {
		log.Fatal(err)
	}
	
	ex, _ := json.Marshal(exchange)
	log.Print("Requisição realizada com sucesso:")
	log.Println(string(ex))
}

// Adicione uma linha no arquivo de log
func addRecord(exch *exchangeReq) error {
	
	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	
	rec := []byte("Dólar: "+ string(exch.Bid) + "\n")
	_, err = file.Write(rec)
	if err != nil {
		return err
	}

	return nil
}

// Recupera as informações da cotação atual
func geExchange() (*exchangeReq, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var c exchangeReq
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
