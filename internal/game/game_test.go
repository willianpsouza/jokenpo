package game

import (
	"jokenpo.provengo.io/internal/dbhandler"
	"strings"
	"testing"
)

func TestDetermineWinner(t *testing.T) {
	tests := []struct {
		player1 string
		player2 string
		winner  string
		message string
	}{
		{"1", "3", "1", "Stone Wins"},
		{"2", "1", "2", "Paper Wins"},
		{"3", "2", "3", "Scissors Wins"},
		{"1", "1", "Draw", "Tied"},
		{"2", "2", "Draw", "Tied"},
		{"3", "3", "Draw", "Tied"},
	}

	for _, test := range tests {
		gotWinner, gotMessage := determineWinner(test.player1, test.player2)
		if gotWinner != test.winner || gotMessage != test.message {
			t.Errorf("For %s vs %s, expected (%s, %s) but got (%s, %s)",
				test.player1, test.player2, test.winner, test.message, gotWinner, gotMessage)
		}
	}
}

func TestGetComputerChoice(t *testing.T) {
	validChoices := map[string]bool{"1": true, "2": true, "3": true}

	for i := 0; i < 100; i++ {
		choice := getComputerChoice()
		if !validChoices[choice] {
			t.Errorf("Invalid computer choice: %s", choice)
		}
	}
}

func TestGameReport(t *testing.T) {
	db, err := dbhandler.InitializeDatabase()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() { _ = db.Close() }()

	game := dbhandler.GameResult{
		Player1Name:   "Alice",
		Player1Choice: "1",
		Player2Name:   "Bob",
		Player2Choice: "3",
		Winner:        "Alice",
		Timestamp:     1234567890,
		UniqueID:      "test-uuid-123",
	}
	dbhandler.SaveGameResult(db, game)

	report := gameReport(db, "Alice")
	if !strings.Contains(report, "Alice") || !strings.Contains(report, "Bob") {
		t.Errorf("Game report did not include expected players: %s", report)
	}
}

func TestWinnerMessages(t *testing.T) {
	tests := []struct {
		player1Name   string
		player2Name   string
		player1Choice string
		player2Choice string
		expected      string
	}{
		{"Alice", "Bob", "1", "3", "Player Alice wins, Stone Wins"},
		{"Charlie", "Dave", "2", "1", "Player Charlie wins, Paper Wins"},
		{"Eve", "Frank", "3", "2", "Player Eve wins, Scissors Wins"},
	}

	for _, test := range tests {
		got := winnerMessages(test.player1Name, test.player2Name, test.player1Choice, test.player2Choice)
		if got != test.expected {
			t.Errorf("Expected message: %s, got: %s", test.expected, got)
		}
	}
}
