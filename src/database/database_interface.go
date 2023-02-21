package database

import (
	model "jump-backend-interview/src/model"
	"log"
)

// We get a database connection, query the database, and return the results
func GetUsers() (*[]model.User, int) {
	db := GetDb()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Default().Println(err.Error())
		return nil, 500
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		var balance int
		err := rows.Scan(&user.Id, &user.First_name, &user.Last_name, &balance)
		user.Balance = float64(balance) / 100
		if err != nil {
			log.Default().Println(err.Error())
			return nil, 500
		}
		users = append(users, user)
	}
	return &users, 200
}

// It takes an invoice, checks if the user exists, and if so, inserts the invoice into the database
func PostInvoice(invoice model.Invoice) int {
	db := GetDb()
	var result bool

	db.QueryRow("SELECT NOT EXISTS(SELECT 1 FROM users WHERE id = $1)", invoice.User_id).Scan(&result)
	if result {
		return 404
	}

	stmt, err := db.Prepare("INSERT INTO invoices (user_id, amount, label) VALUES ($1, $2, $3)")
	if err != nil {
		log.Default().Println(err.Error())
		return 500
	}
	defer stmt.Close()
	_, err = stmt.Exec(invoice.User_id, (invoice.Amount * 100), invoice.Label)
	if err != nil {
		log.Default().Println(err.Error())
		return 500
	}
	return 204
}

// It updates the user's balance and the invoice's status if the invoice exists, the invoice is
// pending, and the amount is correct
func PostTransaction(transaction model.Transaction) int {
	db := GetDb()
	var result bool

	db.QueryRow("SELECT NOT EXISTS(SELECT 1 FROM invoices WHERE id = $1)", transaction.Invoice_id).Scan(&result)
	if result {
		return 404
	}

	db.QueryRow("SELECT EXISTS(SELECT 1 FROM invoices WHERE status = 'paid' AND id = $1)", transaction.Invoice_id).Scan(&result)
	if result {
		return 422
	}

	db.QueryRow("SELECT NOT EXISTS(SELECT 1 FROM invoices WHERE id = $1 AND amount = $2)", transaction.Invoice_id, transaction.Amount*100).Scan(&result)
	if result {
		return 400
	}

	query := `UPDATE users
	SET balance = balance + invoices.amount
	FROM invoices
	WHERE invoices.id = $1
	  AND invoices.amount = $2
	  AND invoices.user_id = users.id
	  AND invoices.status = 'pending'`

	_, err := db.Exec(query, transaction.Invoice_id, (transaction.Amount * 100))
	if err != nil {
		log.Default().Println(err.Error())
		return 500
	}

	query = `
	UPDATE invoices i
	SET status = 'paid'
	WHERE i.id = $1`
	_, err = db.Exec(query, transaction.Invoice_id)
	if err != nil {
		log.Default().Println(err.Error())
		return 500
	}
	return 204
}
