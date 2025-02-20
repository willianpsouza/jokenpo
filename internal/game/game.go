package game

import (
	"bufio"
	"syscall"

	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"database/sql"
	eUUID "github.com/google/uuid"
	"golang.org/x/term"
	lDbHandler "jokenpo.provengo.io/internal/dbhandler"
	lRandom "jokenpo.provengo.io/internal/randomize"
)

func gameReport(db *sql.DB, userName string) string {
	format := "%16s %16s %16s %16s %16s\n"
	ret := fmt.Sprintf(format, "Player 1", "Player 2", "P. 1 Choice", "P. 2 Choice", "Result")

	results, err := lDbHandler.GetResultsFromUser(db, userName)
	if err != nil {
		_ = fmt.Errorf("error getting results from user %s: %v", userName, err)
	}
	for _, result := range results {
		_, msgWinner := determineWinner(result.Player1Choice, result.Player2Choice)
		ret += fmt.Sprintf(
			format,
			result.Player1Name,
			result.Player2Name,
			lDbHandler.Choices[result.Player1Choice],
			lDbHandler.Choices[result.Player2Choice],
			msgWinner,
		)
	}
	return ret
}

func getUserName(player string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the player name \033[1m%s\033[0m: ", player)
	name, _ := reader.ReadString('\n')
	return strings.TrimSpace(name)
}

func getUserChoice(player string, timeout time.Duration) string {
	fmt.Printf("%s, Choose: \n 1 - Rock 🪨\n 2 - Paper 📄\n 3 - Scissors ✂️\n Option: \n", player)
	choiceChan := make(chan string)
	go func() {
		input := getSecretChoice()
		choiceChan <- strings.TrimSpace(input)
	}()

	select {
	case choice := <-choiceChan:
		if _, exists := lDbHandler.Choices[choice]; exists {
			return choice
		}
		fmt.Println("Try Again...")
		return getUserChoice(player, timeout)
	case <-time.After(timeout):
		fmt.Println("Times up !")
		return "timeout"
	}
}

func getSecretChoice() string {
	fd := syscall.Stdin
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return ""
	}
	defer func() {
		_ = term.Restore(fd, oldState)
	}()

	bytePassword, err := term.ReadPassword(fd)
	if err != nil {
		return ""
	}
	return string(bytePassword)
}

func getComputerChoice() string {
	rand.New(rand.NewSource(lRandom.PseudoPrimeGenerator()))
	baseOptions := []string{"1", "2", "3"}
	var keys []string
	for i := 0; i < 8192; i++ {
		keys = append(keys, baseOptions[rand.Intn(len(baseOptions))])
	}
	return keys[rand.Intn(len(keys))]
}

func determineWinner(player1, player2 string) (string, string) {
	if player1 == player2 {
		return "Draw", "Tied"
	}

	switch player1 {
	case "1":
		if player2 == "3" {
			return player1, "Stone Wins"
		} else {
			return player2, "Stone Losses"
		}
	case "2":
		if player2 == "1" {
			return player1, "Paper Wins"
		} else {
			return player2, "Paper Losses"
		}
	case "3":
		if player2 == "2" {
			return player1, "Scissors Wins"
		} else {
			return player2, "Scissors Losses"
		}
	}
	return "---ERR---", "---ERR---"
}

func winnerMessages(player1Name, player2Name, player1Choice, player2Choice string) string {
	winner, messageWinner := determineWinner(player1Choice, player2Choice)
	if winner == player1Choice {
		messageWinner = fmt.Sprintf("Player %s wins, %s", player1Name, messageWinner)
	} else if winner == player2Choice {
		messageWinner = fmt.Sprintf("Player %s wins, %s", player2Name, messageWinner)
	}
	return messageWinner
}

func Start(db *sql.DB) {
	defaultTimeout := time.Second * 5
	Computers := []string{
		"IBM 386At",
		"IBM DeepBlue",
		"Comodore 64",
		"Apple II",
		"ZX Spectrum",
	}

	Welcome := []string{
		"Welcome to Joken Po CLI!",
		"Choose your option:",
		"1 - Human vs Computer ",
		"2 - Human vs Human ",
		"3 - List Your results",
	}

	reader := bufio.NewReader(os.Stdin)
	for _, msg := range Welcome {
		fmt.Println(msg)
	}
	fmt.Print("Options: ")
	modeInput, _ := reader.ReadString('\n')
	mode := strings.TrimSpace(modeInput)

	if mode == "1" || mode == "2" {
		var (
			player1Name   string
			player1Choice string
			player2Name   string
			player2Choice string
			winner        string
			winnerMessage string
		)

		if mode == "1" {
			player1Name = getUserName("Player")
			player1Choice = getUserChoice(player1Name, defaultTimeout)
			player2Name = Computers[rand.Intn(len(Computers))]
			player2Choice = getComputerChoice()
			fmt.Printf("Computer Chooses: %s\n", lDbHandler.Choices[player2Choice])
		}

		if mode == "2" {
			player1Name = getUserName("Player 1")
			player2Name = getUserName("Player 2")
			player1Choice = getUserChoice(player1Name, defaultTimeout)
			player2Choice = getUserChoice(player2Name, defaultTimeout)
		}

		winner, winnerMessage = determineWinner(player1Choice, player2Choice)
		winnerMessage = winnerMessages(player1Name, player2Name, player1Choice, player2Choice)

		fmt.Println(winnerMessage)

		lDbHandler.SaveGameResult(db, lDbHandler.GameResult{
			Player1Name:   player1Name,
			Player1Choice: player1Choice,
			Player2Name:   player2Name,
			Player2Choice: player2Choice,
			Winner:        winner,
			UniqueID:      eUUID.NewString(),
			Timestamp:     time.Now().UnixMilli(),
		})
	} else if mode == "3" {
		player1Name := getUserName("Player 1")

		fmt.Printf(gameReport(db, player1Name))
	} else {
		fmt.Println("Invalid option")
	}
}
