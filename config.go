package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func initDB() {
	var err error

	// Configurar la cadena de conexión sin el parámetro extra_float_digits
	connStr := "user=tutti_cargo_admin password=8Qz$%+W=q17* dbname=tutti_cargo host=5.161.189.5 port=5432 sslmode=disable options='-c extra_float_digits=0'"
	db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxOpenConns(80) // Aumentar el número de conexiones abiertas
	db.SetMaxIdleConns(80) // Aumentar el número de conexiones en espera
}

/**
We set the maximum number of open connections to 80 and the maximum number of idle connections to 80.
This is because we want to increase the number of concurrent connections to the database.
We also set the extra_float_digits parameter to 0 in the connection string.
This is because the extra_float_digits parameter is set to 3 by default in the PostgreSQL database,
which can cause problems when inserting data into the database. By setting it to 0, we avoid these problems.
*/

/**
Why do we need to set max open connections and max idle connections?

The max open connections setting determines the maximum number of open connections to the database. This
setting is important because it can help prevent the database from becoming overloaded with too many open
connections. By setting a maximum number of open connections, we can limit the number of concurrent
connections to the database and prevent it from becoming overwhelmed.
*/
