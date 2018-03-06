package main

import (
	"database/sql"
	"log"
	"net/url"

	"github.com/urfave/cli"
)

const stmtCreate = `
CREATE TABLE IF NOT EXISTS products (
  id BYTES DEFAULT uuid_v4(),
  brand STRING,
  price FLOAT,
  url STRING,
  ref STRING,

  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS list_products (
  productId BYTES DEFAULT uuid_v4(),
  listId BYTES DEFAULT uuid_v4(),
  huntly_price FLOAT,

  PRIMARY KEY (productId, listId),
  UNIQUE INDEX byListID (productId, listId)
);

CREATE TABLE IF NOT EXISTS photos (
  id BYTES DEFAULT uuid_v4(),
  productId BYTES DEFAULT uuid_v4(),
  caption STRING,
  timestamp TIMESTAMP,

  PRIMARY KEY (id),
  UNIQUE INDEX byProductID (productId, timestamp)
);

CREATE TABLE IF NOT EXISTS list (
  id BYTES DEFAULT uuid_v4(),
  name STRING,
  memberId BYTES DEFAULT uuid_v4(),
  visibility BOOLEAN,
  views BIGINT,

  PRIMARY KEY (id),
  UNIQUE INDEX (id, name)
);

CREATE TABLE IF NOT EXISTS members (
  id BYTES DEFAULT uuid_v4(),
  memberId STRING NOT NULL,
  token STRING NOT NULL,

  PRIMARY KEY (id),
  UNIQUE INDEX (memberId),
  UNIQUE INDEX (token)
);
`

func createDB(c *cli.Context) {
	url := parseURL(c)
	db, err := sql.Open("postgres", url.String())
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = db.Close() }()

	const query = `CREATE DATABASE IF NOT EXISTS huntly`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(stmtCreate); err != nil {
		log.Fatal(err)
	}
}

func initDB(c *cli.Context) *sql.DB {
	url := parseURL(c)
	db, err := sql.Open("postgres", url.String())
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func parseURL(c *cli.Context) *url.URL {
	parsedURL, err := url.Parse("postgres://" +
		c.String("cockroach-user") +
		"@" + c.String("cockroach-host") +
		":" + c.String("cockroach-port") +
		"?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	parsedURL.Path = c.String("cockroach-db")
	return parsedURL
}
