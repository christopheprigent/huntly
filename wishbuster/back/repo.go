package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

// Get user with token
func getUserIDFromToken(db *sql.DB, token string) (string, error) {
	var id string
	// Get member id with token
	query := `SELECT id FROM members WHERE token = $1;`
	err := db.QueryRow(query, token).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return "", fmt.Errorf("error executing the query: %v", err)
	}

	if id == "" {
		return "", fmt.Errorf("user ID does not exist for token: %s", token)
	}

	return id, nil
}

// RepoCreateUser creates a user and returns a token
func RepoCreateUser(db *sql.DB, member Member) (Member, error) {
	var token string

	if member.ID == "" {
		return Member{}, errors.New("member ID is empty")
	}

	query := `SELECT token FROM members WHERE memberId = $1;`

	err := db.QueryRow(query, member.ID).Scan(&token)
	if token != "" {
		member.Token = token
		return member, nil
	}
	member.Token = GetMD5Hash(string(member.ID))

	query = `INSERT INTO members VALUES ($1, $2, $3);`
	_, err = db.Exec(query, uuid.NewV4().String(), member.ID, member.Token)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return Member{}, fmt.Errorf("cannot insert into db: %v", err)
	}
	log.Info("successfully inserted entry")

	private := List{
		Token:      member.Token,
		Name:       "Privée",
		Visibility: false,
	}

	RepoCreateList(db, private)

	public := List{
		Token:      member.Token,
		Name:       "Publique",
		Visibility: true,
	}

	RepoCreateList(db, public)

	return member, nil
}

//RepoCreateList create a list
func RepoCreateList(db *sql.DB, list List) (string, error) {
	var id string

	if list.Token == "" {
		return "", errors.New("Token is empty")
	}

	if list.Name == "" {
		return "", errors.New("Name is empty")
	}

	memberID, err := getUserIDFromToken(db, list.Token)
	if err != nil {
		return "", fmt.Errorf("cannot execute query: %v", err)
	}

	query := `SELECT id FROM list WHERE name = $1 AND memberId = $2;`
	err = db.QueryRow(query, list.Name, memberID).Scan(&id)
	if id != "" {
		return "", errors.New("List already exists for user")
	}

	uuid := uuid.NewV4().String()

	query = `INSERT INTO list VALUES ($1, $2, $3, $4, $5);`
	_, err = db.Exec(query, uuid, list.Name, memberID, list.Visibility, 0)
	if err != nil {
		return "", fmt.Errorf("cannot execute query: %v", err)
	}

	log.Info("successfully inserted list")

	return uuid, nil
}

// RepoGetList get list with all product
//TODO if token is unset get only list when visibility = true
func RepoGetList(db *sql.DB, list List) (List, error) {
	if list.ID == "" {
		return List{}, errors.New("ID is empty")
	}

	query := `SELECT name, visibility, views FROM list WHERE id = $1;`
	err := db.QueryRow(query, list.ID).Scan(&list.Name, &list.Visibility, &list.Views)
	if err != nil {
		return List{}, fmt.Errorf("error executing the query: %v", err)
	}

	query = `SELECT productid, huntly_price FROM list_products WHERE listid = $1;`
	rows, err := db.Query(query, list.ID)
	if err != nil && err != sql.ErrNoRows {
		return List{}, fmt.Errorf("error executing the query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		p := Product{}
		err = rows.Scan(&p.ID, &p.HuntlyPrice)
		if err != nil {
			return List{}, err
		}
		p.ListID = list.ID
		query = `SELECT brand, price, url, ref, id FROM products WHERE id = $1;`
		err = db.QueryRow(query, p.ID).Scan(&p.Brand, &p.Price, &p.URL, &p.Ref, &p.ID)
		if err != nil {
			return List{}, fmt.Errorf("error executing the query: %v", err)
		}
		query = `SELECT caption FROM photos WHERE productId = $1;`
		err = db.QueryRow(query, p.ID).Scan(&p.Picture)
		if err != nil && err != sql.ErrNoRows {
			return List{}, fmt.Errorf("error executing the query: %v", err)
		}
		list.Products = append(list.Products, p)
	}

	log.Info("successfully get list")
	return list, nil
}

//RepoGetLists get all lists of a member
func RepoGetLists(db *sql.DB, member Member) ([]List, error) {
	var lists []List

	if member.Token == "" {
		return nil, errors.New("token is empty")
	}

	// Get member id with token
	memberID, err := getUserIDFromToken(db, member.Token)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %v", err)
	}

	query := `SELECT name, visibility, views, id FROM list WHERE memberid = $1;`
	rows, err := db.Query(query, memberID)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("error executing the query: %v", err)
	}
	defer rows.Close()
	//for each list
	for rows.Next() {
		list := List{}
		err = rows.Scan(&list.Name, &list.Visibility, &list.Views, &list.ID)
		if err != nil {
			return nil, err
		}
		query = `SELECT productid, huntly_price FROM list_products WHERE listid = $1;`
		productRows, err := db.Query(query, list.ID)
		if err != nil && err != sql.ErrNoRows {
			return nil, fmt.Errorf("error executing the query: %v", err)
		}
		defer productRows.Close()

		//for each product of the list
		for productRows.Next() {
			p := Product{}
			err = productRows.Scan(&p.ID, &p.HuntlyPrice)
			if err != nil {
				return nil, err
			}
			p.ListID = list.ID
			query = `SELECT brand, price, url, ref FROM products WHERE id = $1;`
			err = db.QueryRow(query, p.ID).Scan(&p.Brand, &p.Price, &p.URL, &p.Ref)
			if err != nil {
				return nil, fmt.Errorf("error executing the query: %v", err)
			}
			query = `SELECT caption FROM photos WHERE productId = $1;`
			err = db.QueryRow(query, p.ID).Scan(&p.Picture)
			if err != nil && err != sql.ErrNoRows {
				return nil, fmt.Errorf("error executing the query: %v", err)
			}
			list.Products = append(list.Products, p)
		}
		lists = append(lists, list)
	}

	log.Info("successfully get lists")
	return lists, nil
}

//RepoAddProduct add a product to list
func RepoAddProduct(db *sql.DB, product Product) (string, error) {
	productID := uuid.NewV4().String()

	if product.ListID == "" {
		return "", fmt.Errorf("list ID not specified")
	}

	// Create Product
	query := `INSERT INTO products VALUES ($1, $2, $3, $4, $5);`
	_, err := db.Exec(query, productID, product.Brand, product.Price, product.URL, product.Ref)
	if err != nil {
		return "", fmt.Errorf("cannot execute query: %v", err)
	}

	// Insert photo
	query = `INSERT INTO photos VALUES ($1, $2, $3, $4);`
	_, err = db.Exec(query, uuid.NewV4().String(), productID, product.Picture, time.Now())
	if err != nil {
		return "", fmt.Errorf("cannot execute query: %v", err)
	}

	// TODO: Verify that list exists before inserting

	// Create Product/List association
	query = `INSERT INTO list_products VALUES ($1, $2, $3);`
	_, err = db.Exec(query, productID, product.ListID, product.HuntlyPrice)
	if err != nil {
		return "", fmt.Errorf("cannot execute query: %v", err)
	}

	log.Info("product successfully added")
	return productID, nil
}

//RepoDeleteProduct delete a product from a list
func RepoDeleteProduct(db *sql.DB, product Product) error {
	if product.ListID == "" {
		return fmt.Errorf("list ID not specified")
	}
	if product.ID == "" {
		return fmt.Errorf("product ID not specified")
	}

	// Delete Product
	query := `DELETE FROM list_products WHERE (listid, productid) = ($1, $2)`
	_, err := db.Exec(query, product.ListID, product.ID)
	if err != nil {
		return fmt.Errorf("cannot execute query: %v", err)
	}

	log.Info("product successfully deleted")
	return nil
}

//RepoDeleteList delete a list
func RepoDeleteList(db *sql.DB, list List) error {
	if list.ID == "" {
		return fmt.Errorf("list ID not specified")
	}

	if list.Token == "" {
		return errors.New("token is empty")
	}

	// Get member id with token
	memberID, err := getUserIDFromToken(db, list.Token)
	if err != nil {
		return fmt.Errorf("cannot execute query: %v", err)
	}

	query := `SELECT name, visibility, views FROM list WHERE (memberid, id) = ($1, $2);`
	err = db.QueryRow(query, memberID, list.ID).Scan(&list.Name, &list.Visibility, &list.Views)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("list does not exist")
		}
		return fmt.Errorf("error executing the query: %v", err)
	}

	// Delete Products
	query = `DELETE FROM list_products WHERE listid = $1`
	_, err = db.Exec(query, list.ID)
	if err != nil {
		return fmt.Errorf("cannot execute query: %v", err)
	}
	// Delete List
	query = `DELETE FROM list WHERE id = $1`
	_, err = db.Exec(query, list.ID)
	if err != nil {
		return fmt.Errorf("cannot execute query: %v", err)
	}

	log.Info("list successfully deleted")
	return nil
}

//RepoUpdateList update a list
func RepoUpdateList(db *sql.DB, list List) error {
	if list.ID == "" {
		return fmt.Errorf("list ID not specified")
	}

	if list.Name == "" {
		return fmt.Errorf("list Name not specified")
	}

	// if list.Visibility == nil {
	// 	query := `SELECT visibility FROM list WHERE id = $1;`
	// 	err := db.QueryRow(query, list.ID).Scan( &list.Visibility)
	// 	if err != nil {
	// 		return List{}, fmt.Errorf("error executing the query: %v", err)
	// 	}
	// }

	if list.Token == "" {
		return errors.New("token is empty")
	}

	// Get member id with token
	memberID, err := getUserIDFromToken(db, list.Token)
	if err != nil {
		return fmt.Errorf("cannot execute query: %v", err)
	}

	query := `UPDATE list SET name = $1, visibility = $2 FROM list WHERE (memberid, id) = ($3, $4);`
	_, err = db.Exec(query, list.Name, list.Visibility, memberID, list.ID)
	if err != nil {
		return fmt.Errorf("error executing the query: %v", err)
	}

	log.Info("list successfully updated")
	return nil
}

//RepoGetProduct delete a list
//TODO get product without ID
func RepoGetProduct(db *sql.DB, product Product) (Product, error) {
	if product.ID == "" {
		return Product{}, fmt.Errorf("product ID not specified")
	}

	// Select product
	query := `SELECT COUNT(listid) FROM list_products WHERE productid = $1;`
	err := db.QueryRow(query, product.ID).Scan(&product.Count)
	if err != nil {
		return Product{}, fmt.Errorf("cannot execute query: %v", err)
	}

	log.Info("count product successfully get")
	return product, nil
}
