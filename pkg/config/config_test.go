package config

import (
	"errors"
	"testing"
)

func TestNewConfigPostgres(t *testing.T) {
	config := NewConfig(
		WithDriver(PostgresNamedDriver),
		WithDatabase("test"),
		WithUser("root"),
		WithPassword("root"),
		WithHost("localhost"),
		WithPort(5432),
	)

	if err := config.Validate(); err != nil {
		t.Error("expected nil, got ", err)
	}

	config.Db.Host = "" // reset host to empty string, to test validation

	if err := config.Validate(); err != nil && !errors.Is(err, ErrInvalidHost) {
		t.Error("expected (db host is required), got ", err)
	}
}

func TestNewConfigMySQL(t *testing.T) {
	config := NewConfig(
		WithDriver(MySQLNamedDriver),
		WithDatabase("test"),
		WithUser("root"),
		WithPassword("root"),
		WithHost("localhost"),
		WithPort(3306),
	)

	if err := config.Validate(); err != nil {
		t.Error("expected nil, got ", err)
	}
}

func TestNewConfigSQLite(t *testing.T) {
	config := NewConfig(
		WithDriver(SQLiteNamedDriver),
		WithSqliteDB("test", "test.db"),
	)

	if err := config.Validate(); err != nil {
		t.Error("expected nil, got ", err)
	}
}

func TestNewConfigBoltDB(t *testing.T) {
	config := NewConfig(
		WithDriver(BoltNamedDriver),
		WithBoltDB("test", "test.db"),
	)

	if err := config.Validate(); err != nil {
		t.Error("expected nil, got ", err)
	}
}

func TestNewConfigOracle(t *testing.T) {
	config := NewConfig(
		WithDriver(OracleNamedDriver),
		WithDatabase("test"),
		WithUser("root"),
		WithPassword("root"),
		WithHost("localhost"),
		WithPort(1521),
	)

	if err := config.Validate(); err != nil {
		t.Error("expected nil, got ", err)
	}
}
