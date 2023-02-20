package database

import (
	model "jump-backend-interview/src/model"
)

func GetUsers() (*[]model.User, int) {
	db := GetDb()
	rows, res := db.Query("SELECT * FROM users")
	if res != nil {
		println(res.Error())
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
			println(res.Error())
			return nil, 500
		}
		users = append(users, user)
	}
	return &users, 200
}

func PostInvoice(invoice model.Invoice) int {
	db := GetDb()
	stmt, err := db.Prepare("INSERT INTO invoices (user_id, amount, label) VALUES ($1, $2, $3)")
	if err != nil {
		println(err.Error())
		return 500
	}
	defer stmt.Close()
	_, err = stmt.Exec(invoice.User_id, (invoice.Amount * 100), invoice.Label)
	if err != nil {
		println(err.Error())
		return 500
	}
	return 204
}

func PostTransaction(transaction model.Transaction) int {
	db := GetDb()
	var result bool

	db.QueryRow("SELECT NOT EXISTS(SELECT 1 FROM invoices WHERE id = $1)", transaction.Invoice_id).Scan(&result)
	println("result 404: ", result)
	if result {
		return 404
	}

	db.QueryRow("SELECT EXISTS(SELECT 1 FROM invoices WHERE status = 'paid' AND id = $1)", transaction.Invoice_id).Scan(&result)
	println("result 422: ", result)
	if result {
		return 422
	}

	db.QueryRow("SELECT NOT EXISTS(SELECT 1 FROM invoices WHERE id = $1 AND amount = $2)", transaction.Invoice_id, transaction.Amount*100).Scan(&result)
	println("result 400: ", result)
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
		println(err.Error())
		return 500
	}

	query = `WITH invoice_to_pay AS (
		SELECT i.id, i.status, i.amount, u.id AS user_id, u.balance
		FROM invoices i
		JOIN users u ON i.user_id = u.id
		WHERE i.id = $1
	)
	UPDATE invoices i
	SET status = s.status
	FROM invoice_to_pay itp
	JOIN invoice_status s ON s.status = 'paid'
	WHERE i.id = itp.id
	  AND itp.status = 'pending'
	  AND itp.amount = $2`
	_, err = db.Exec(query, transaction.Invoice_id, (transaction.Amount * 100))
	if err != nil {
		println(err.Error())
		return 500
	}
	return 204
}
