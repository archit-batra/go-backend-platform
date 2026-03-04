package wallet

import (
	"testing"

	"github.com/archit-batra/go-backend-platform/internal/testutil"
)

func TestTransferSuccess(t *testing.T) {

	db := testutil.SetupTestDB()
	defer db.Close()

	repo := NewRepository(db)

	// Clean tables
	db.Exec("DELETE FROM audit_logs")
	db.Exec("DELETE FROM wallets")
	db.Exec("DELETE FROM users")

	// Insert test users
	db.Exec("INSERT INTO users (id, name, email) VALUES (1, 'A', 'a@test.com')")
	db.Exec("INSERT INTO users (id, name, email) VALUES (2, 'B', 'b@test.com')")

	// Insert wallets
	db.Exec("INSERT INTO wallets (user_id, balance) VALUES (1, 1000)")
	db.Exec("INSERT INTO wallets (user_id, balance) VALUES (2, 0)")

	err := repo.Transfer(1, 2, 500)
	if err != nil {
		t.Fatalf("transfer failed: %v", err)
	}

	var balance int
	db.QueryRow("SELECT balance FROM wallets WHERE user_id = 1").Scan(&balance)

	if balance != 500 {
		t.Fatalf("expected 500, got %d", balance)
	}
}

func TestTransferInsufficientBalance(t *testing.T) {

	db := testutil.SetupTestDB()
	defer db.Close()

	repo := NewRepository(db)

	db.Exec("DELETE FROM wallets")
	db.Exec("DELETE FROM users")

	db.Exec("INSERT INTO users (id, name, email) VALUES (1, 'A', 'a@test.com')")
	db.Exec("INSERT INTO users (id, name, email) VALUES (2, 'B', 'b@test.com')")

	db.Exec("INSERT INTO wallets (user_id, balance) VALUES (1, 100)")
	db.Exec("INSERT INTO wallets (user_id, balance) VALUES (2, 0)")

	err := repo.Transfer(1, 2, 500)

	if err == nil {
		t.Fatalf("expected error for insufficient balance")
	}
}
