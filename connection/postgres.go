package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DatabaseConnection()  {
	var err error
	// postgres://{user}:{password}@{host}:{port}/{database}
	dbUrl := "postgres://postgres:999@localhost:5432/personalWeb"
	Conn, err = pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		// fprintf ?
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
		os.Exit(1)
	}

	fmt.Println("Database Connected.")
}