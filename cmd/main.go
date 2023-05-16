package main

import (
	"clean-sales/internal/dtos"
	"clean-sales/internal/entities"
	"clean-sales/internal/infra/repositories"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file::memory:")
	if err != nil {
		log.Fatal(err)
	}

	if err := prepDB(db); err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Use((middleware.Logger))

	r.Route(("/checkout"), func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			var input dtos.CheckoutInputDto
			var output dtos.CheckoutOutputDto
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// Validate CPF
			err := entities.Validate(input.Cpf)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			productRepository := repositories.NewProductRepositoryImpl(db)
			couponRepository := repositories.NewCouponRepositoryImpl(db)

			var checkDuplicate []string

			if (len(input.Items)) > 0 {
				for _, item := range input.Items {
					if item.Quantity < 0 {
						w.WriteHeader(http.StatusBadRequest)
						w.Write([]byte("invalid quantity"))
						return
					}

					product, err := productRepository.GetProduct(item.IdProduct)
					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
						w.Write([]byte(err.Error()))
						return
					}

					for _, v := range checkDuplicate {
						if v == product.IdProduct {
							w.WriteHeader(http.StatusBadRequest)
							w.Write([]byte("product duplicated"))
							return
						}
					}

					checkDuplicate = append(checkDuplicate, product.IdProduct)
					output.Total += product.Price * float64(item.Quantity)
				}
			}

			if input.Coupon != "" {
				coupon, err := couponRepository.GetCoupon(input.Coupon)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(err.Error()))
					return
				}
				output.Total -= output.Total * (coupon.Discount / 100)
			}

			json.NewEncoder(w).Encode(output)
		})
	})

	http.ListenAndServe(":9090", r)
}

func prepDB(db *sql.DB) error {

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS coupons (code text PRIMARY KEY NOT NULL, discount DECIMAL(10,2), expire_at DATE);")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products (id_product text PRIMARY KEY NOT NULL,description text, price DECIMAL(10,2), width DECIMAL(10,2), height DECIMAL(10,2), length DECIMAL(10,2), weight DECIMAL(10,2));")
	if err != nil {
		return err
	}

	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS orders (id_order text PRIMARY KEY NOT NULL,cpf text, code text, total DECIMAL(10,2));")
	// if err != nil {
	// 	return err
	// }

	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS items (id_order text, id_product text, price DECIMAL(10,2), quantity integer, PRIMARY KEY (id_order, id_product), FOREIGN KEY (id_order) REFERENCES orders(id_order), FOREIGN KEY (id_product) REFERENCES product(id_product));")
	// if err != nil {
	// 	return err
	// }

	_, err = db.Exec("INSERT INTO products (id_product, description, price, width, height, length, weight) VALUES (1, 'A', 1000, 100, 30, 10, 3);")
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO products (id_product, description, price, width, height, length, weight) VALUES (2, 'B', 5000, 50, 50, 50, 22);")
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO products (id_product, description, price, width, height, length, weight) VALUES (3, 'C', 30, 10, 10, 10, 0.9);")
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO products (id_product, description, price, width, height, length, weight) VALUES (4, 'D', 1000, -100, 30, 10, 3.0);")
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO products (id_product, description, price, width, height, length, weight) VALUES (5, 'E', 1000, 100, 30, 10, -3);")
	if err != nil {
		return err
	}

	validDate := time.Now().AddDate(1, 0, 0).Format("2006-01-02")
	invalidDate := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")

	_, err = db.Exec("INSERT INTO coupons (code, discount, expire_at) VALUES ('CUPOM10',10.00, '" + validDate + "');")
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO coupons (code, discount, expire_at) VALUES ('EXPIRED',10.00, '" + invalidDate + "');")
	if err != nil {
		return err
	}
	return nil
}
