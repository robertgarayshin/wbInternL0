package config

const (
	user     = "user=robert_admin"
	password = "password=passwd"
	dbname   = "dbname=wb_intern_db"
	sslmode  = "sslmode=disable"
)

const ConnStr = user + " " + password + " " + dbname + " " + sslmode
