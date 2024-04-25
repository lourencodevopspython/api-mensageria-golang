package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/lourencodevopspython/api_mensageria/internal/infra/akafka"
	"github.com/lourencodevopspython/api_mensageria/internal/infra/repository"
	"github.com/lourencodevopspython/api_mensageria/internal/infra/web"
	"github.com/lourencodevopspython/api_mensageria/internal/usecase"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3310)/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductUsecase := usecase.NewCreateProductsUseCase(repository)
	listProductsUsecase := usecase.NewListProductsUseCase(repository)

	prooductHandlers := web.NewProductHandlers(createProductUsecase, listProductsUsecase)
	r := chi.NewRouter()
	r.Post("/products", prooductHandlers.CreateProductHandler)
	r.Get("/products", prooductHandlers.ListProductsHandler)

	go http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			//logar o erro
		}
		_, err = createProductUsecase.Execute(dto)
	}
}
