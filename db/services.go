package db

func InsertLogInDb(dni string) {
	dbConnection.Exec("INSERT INTO logs (dni, estado, fecha) values ($1, $2, now());", dni, "EN PROCESO")
}

// insert wallet in db
func InsertWalletInDb(dni string) {
	dbConnection.Exec("INSERT INTO wallet (dni, fecha) values ($1, now());", dni)
}
