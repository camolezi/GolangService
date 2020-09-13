package data

import (
	"context"

	"github.com/camolezi/MicroservicesGolang/src/debug"
	"github.com/jackc/pgx/v4/pgxpool"
)

//this file handles starting and cofiguring the database connection

var databaseAccess databaseConnection

//InitializeDatabase is used for strating the database connection
func InitializeDatabase(config string, log debug.Logger) {
	databaseAccess = databaseConnection{dbconfig: config, log: log}
	databaseAccess.initialize()
}

//CloseDatabase is used to close the database connection
func CloseDatabase() {
	databaseAccess.close()
}

//databaseConnection handle database connection
type databaseConnection struct {
	dbconfig   string
	log        debug.Logger
	connection *pgxpool.Pool
}

//Initialize try to strat a database connection
func (d *databaseConnection) initialize() {
	//Just for test
	dbpool, err := pgxpool.Connect(context.Background(), d.dbconfig)
	if err != nil {
		d.log.Error().Fatalf("Unable to connect to database: %v\n", err)
	}

	d.connection = dbpool

	//Create a ping here
	d.log.Debug().Println("Database connected")

}

func (d *databaseConnection) close() {
	d.connection.Close()
}
