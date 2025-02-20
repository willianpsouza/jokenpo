package dbhandler

import (
	"testing"
	"time"
)

func TestInitializeDatabase(t *testing.T) {
	db, err := InitializeDatabase()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	defer func() { _ = db.Close() }()

	var tableName string
	err = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='results';").Scan(&tableName)
	if err != nil {
		t.Errorf("Table 'results' was not created: %v", err)
	}
}

func TestSaveGameResult(t *testing.T) {
	db, err := InitializeDatabase()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() { _ = db.Close() }()

	game := GameResult{
		Player1Name:   "Alice",
		Player1Choice: "1",
		Player2Name:   "Bob",
		Player2Choice: "2",
		Winner:        "Alice",
		Timestamp:     time.Now().Unix(),
		UniqueID:      "test-uuid-123",
	}

	SaveGameResult(db, game)

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM results WHERE uuid = ?", game.UniqueID).Scan(&count)
	if err != nil {
		t.Errorf("Failed to query database: %v", err)
	}

	if count == 0 {
		t.Errorf("Expected game result to be inserted, but it was not found")
	}
}

func TestGetResultsFromUser(t *testing.T) {
	db, err := InitializeDatabase()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() { _ = db.Close() }()

	game := GameResult{
		Player1Name:   "Charlie",
		Player1Choice: "3",
		Player2Name:   "Dave",
		Player2Choice: "1",
		Winner:        "Charlie",
		Timestamp:     time.Now().Unix(),
		UniqueID:      "test-uuid-456",
	}
	SaveGameResult(db, game)

	results, err := GetResultsFromUser(db, "Charlie")
	if err != nil {
		t.Errorf("Failed to retrieve game results: %v", err)
	}

	if len(results) == 0 {
		t.Errorf("Expected at least one game result for 'Charlie', but got none")
	}
}
