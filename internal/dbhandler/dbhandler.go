package dbhandler

import (
	"database/sql"
	"fmt"
	eUUID "github.com/google/uuid"
	lEncrypt "jokenpo.provengo.io/internal/encrypt"
	lSetup "jokenpo.provengo.io/internal/setup"
	_ "modernc.org/sqlite"
	"time"
)

type GameResult struct {
	UniqueID      string
	Player1Name   string
	Player1Choice string
	Player2Name   string
	Player2Choice string
	Winner        string
	Timestamp     int64
}

type GameResultReport struct {
	Player1Name   string
	Player2Name   string
	Player1Choice string
	Player2Choice string
	Winner        string
	Timestamp     int64
}

var (
	Choices = map[string]string{
		"1": "Stone 🪨",
		"2": "Paper 📄",
		"3": "Scissors ✂️",
	}
)

func SaveGameResult(db *sql.DB, result GameResult) {
	var lastKeySaved string

	func() string {
		query := "select signature from results order by id desc limit 1;"
		if err := db.QueryRow(query, "").Scan(&lastKeySaved); err != nil {
			lastKeySaved = lSetup.DefaultEncKey
		}
		return lastKeySaved
	}()

	strConcat := fmt.Sprintf(
		lSetup.BlockConcat,
		result.Player1Name,
		result.Player1Choice,
		result.Player2Name,
		result.Player2Choice,
		result.UniqueID,
		result.Timestamp,
		lastKeySaved,
	)
	strSignature := lEncrypt.CalculateChecksum([]byte(strConcat), lSetup.DefaultAlgo)
	insertQuery := "INSERT INTO results (player1_name, player1_choice, player2_name, player2_choice, winner, signature, uuid, timestamp) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(insertQuery,
		result.Player1Name,
		result.Player1Choice,
		result.Player2Name,
		result.Player2Choice,
		result.Winner,
		strSignature,
		result.UniqueID,
		result.Timestamp,
	)
	if err != nil {
		fmt.Println("fail to save record:", err)
	}
}

func GetResultsFromUser(db *sql.DB, userID string) ([]GameResultReport, error) {
	var results []GameResultReport
	selectQuery := "SELECT player1_name, player2_name, player1_choice, player2_choice, winner, timestamp FROM results WHERE player1_name = ?"
	rows, err := db.Query(selectQuery, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var r GameResultReport
		if err := rows.Scan(&r.Player1Name, &r.Player2Name, &r.Player1Choice, &r.Player2Choice, &r.Winner, &r.Timestamp); err != nil {
			return nil, err
		} else {

			results = append(results, r)
		}
	}
	defer func() {
		_ = rows.Close()
	}()
	return results, nil
}

func InitializeDatabase() (*sql.DB, error) {

	db, err := sql.Open(lSetup.SqlDriver, lSetup.SqliteFilename)
	if err != nil {
		return nil, fmt.Errorf("error trying to open Database: %s, Driver %s, Error: %v", lSetup.SqliteFilename, lSetup.SqlDriver, err)
	}

	createQuery := "" +
		"CREATE TABLE IF NOT EXISTS results " +
		"(id INTEGER PRIMARY KEY AUTOINCREMENT, " +
		"player1_name TEXT, " +
		"player1_choice TEXT, " +
		"player2_name TEXT, " +
		"player2_choice TEXT, " +
		"winner TEXT, " +
		"signature TEXT, " +
		"uuid TEXT," +
		"timestamp INTEGER);" +
		"CREATE INDEX IF NOT EXISTS idx_results_player_name ON results(player1_name);"

	_, err = db.Exec(createQuery)

	if err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("error on create table: %v", err)
	}

	func(db *sql.DB) {
		var totalLines int
		query := "SELECT COUNT(*) FROM results"
		if err := db.QueryRow(query).Scan(&totalLines); err != nil {
			fmt.Println("error on count result table:", err)
		}

		if totalLines == 0 {
			SaveGameResult(db, GameResult{
				Player1Name:   "player1",
				Player1Choice: "1",
				Player2Name:   "player2",
				Player2Choice: "2",
				Winner:        "player1",
				Timestamp:     time.Now().Unix(),
				UniqueID:      eUUID.NewString(),
			})
		}
	}(db)
	return db, nil
}
