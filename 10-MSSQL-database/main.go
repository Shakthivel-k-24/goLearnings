//############################################################################
// sample for native querying without orm
//############################################################################

// package main

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"time"

// 	_ "github.com/microsoft/go-mssqldb"
// )

// type sample struct {
// 	a int
// 	b int
// }

// func main() {
// 	conn_str := "server=20.244.89.51;user id=RNDUser;password=JGDHealth@2022;port=1433;database=HealthPortDev;"
// 	db, err := sql.Open("sqlserver", conn_str)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	ctx := context.Background()
// 	if err = db.PingContext(ctx); err != nil {
// 		fmt.Println("Error pinging database: ", err)
// 	}
// 	fmt.Println("Connected to MSSQL!")
// 	fmt.Println("Hello there ")
// 	// query := "INSERT INTO AGW VALUES (@a,@b)"
// 	// var id int
// 	// err2 := db.QueryRowContext(context.Background(), query, sql.Named("a", 1), sql.Named("b", 5)).Scan((&id))
// 	// fmt.Println(id)
// 	// fmt.Println(err2)
// 	crx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	// var samp sample
// 	rows, err := db.QueryContext(crx, "SELECT * FROM AGW;")
// 	for rows.Next() {
// 		var a int
// 		var b int
// 		err := rows.Scan(&a, &b)
// 		if err != nil {
// 			fmt.Println("error", err.Error())
// 		}
// 		fmt.Println(a, b)
// 	}

//		defer cancel()
//	}
package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"
)

type sample struct {
	A int `db:"a"`
	B int `db:"b"`
}

func main() {
	conn_str := "server=20.244.89.51;user id=RNDUser;password=JGDHealth@2022;port=1433;database=HealthPortDev;"
	db, err := sqlx.Open("sqlserver", conn_str)
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	if err = db.Ping(); err != nil {
		fmt.Println("Ping:", err)
	}
	var samps []sample
	err = db.Select(&samps, "SELECT * FROM AGW")
	if err != nil {
		fmt.Println(err)
	}
	for _, val := range samps {
		fmt.Println(val)
	}
}
