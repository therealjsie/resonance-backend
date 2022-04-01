package internals

import "testing"

func TestDatabaseConnectionString(t *testing.T) {
	config := Config{
		Host:     "host",
		Port:     0,
		Username: "username",
		Password: "password",
		Database: "database",
	}

	connection_string := config.DatabaseConnectionString()

	if connection_string != "postgres://username:password@host:0/database?sslmode=disable" {
		t.Fail()
	}
}
