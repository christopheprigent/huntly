package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	// Import postgres driver.
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/urfave/cli"
)

// Server is the main object, holding the router and db.
type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

func main() {
	app := cli.NewApp()
	app.Name = "huntly"
	app.Usage = "Huntly server"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		// cockroach db
		cli.StringFlag{
			Name:   "cockroach-db",
			Usage:  "cockroach db",
			EnvVar: "COCKROACH_DB",
			Value:  "huntly",
		},
		cli.StringFlag{
			Name:   "cockroach-user",
			Usage:  "cockroach user",
			EnvVar: "COCKROACH_USER",
			Value:  "root",
		},
		cli.StringFlag{
			Name:   "cockroach-password",
			Usage:  "cockroach password",
			EnvVar: "COCKROACH_PASSWORD",
		},
		cli.StringFlag{
			Name:   "cockroach-host",
			Usage:  "cockroach host",
			EnvVar: "COCKROACH_HOST",
			Value:  "localhost",
		},
		cli.IntFlag{
			Name:   "cockroach-port",
			Usage:  "cockroach port",
			EnvVar: "COCKROACH_PORT",
			Value:  26257,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "start",
			Usage:  "Start the huntly server",
			Flags:  app.Flags,
			Action: start,
		},
		{
			Name:   "init",
			Usage:  "Initialize the huntly database",
			Flags:  app.Flags,
			Action: create,
		},
		{
			Name:   "clean",
			Usage:  "Clean the huntly database",
			Flags:  app.Flags,
			Action: clean,
		},
	}
	app.Run(os.Args)
}

func start(c *cli.Context) {
	db := initDB(c)
	router := mux.NewRouter()

	server := &Server{
		DB:     db,
		Router: router,
	}

	server.initializeRoutes()
	server.Run()
}

func create(c *cli.Context) {
	createDB(c)
}

func clean(c *cli.Context) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("couldn't get working directory")
	}
	fmt.Println(dir)
	err = os.RemoveAll(dir + "/cockroach-data")
	if err != nil {
		log.Fatal("couldn't delete the content of data directory")
	}
}

// Run starts the huntly server.
func (s *Server) Run() {
	fmt.Println("huntly server listening on: :5050")
	log.Fatal(http.ListenAndServe(":5050", s.Router))
}
