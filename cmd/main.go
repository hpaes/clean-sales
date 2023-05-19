package main

import (
	"clean-sales/internal/app/dtos"
	"clean-sales/internal/app/usecases"
	"clean-sales/internal/infra/repositories"
	testfixture "clean-sales/testFixture"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file::memory:")
	if err != nil {
		log.Fatal(err)
	}

	if err := testfixture.PrepDb(db); err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	productRepository := repositories.NewProductRepositoryImpl(db)
	couponRepository := repositories.NewCouponRepositoryImpl(db)
	orderRepository := repositories.NewOrderRepositoryImpl(db)

	checkoutUseCase := usecases.NewCheckoutUseCaseImpl(productRepository, couponRepository, orderRepository)

	r.Use((middleware.Logger))

	r.Route(("/checkout"), func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			var input dtos.CheckoutInputDto
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			output, err := checkoutUseCase.Execute(&input)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())
				return
			}
			json.NewEncoder(w).Encode(output)
		})
	})

	http.ListenAndServe(":9090", r)
}
