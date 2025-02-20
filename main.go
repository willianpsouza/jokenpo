package main

import (
	"fmt"
	lDbHandler "jokenpo.provengo.io/internal/dbhandler"
	lGame "jokenpo.provengo.io/internal/game"
	lSetup "jokenpo.provengo.io/internal/setup"
	lUtils "jokenpo.provengo.io/internal/utils"
	"runtime"
)

func startGame() {

	lUtils.Check([]string{lSetup.DataPath})

	db, err := lDbHandler.InitializeDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	lGame.Start(db)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GC()
	startGame()
}
