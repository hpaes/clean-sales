package testfixture

import (
	"database/sql"
	"time"
)

func PrepDb(db *sql.DB) error {
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
