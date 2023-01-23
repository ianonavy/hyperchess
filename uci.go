package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/freeeve/uci"
)

func engine(port string, isWhite bool, nextMove chan string, results chan string) {
	eng, err := uci.NewEngine("/usr/bin/nc", "localhost", port)
	if err != nil {
		log.Fatal(err)
	}

	// set some engine options
	eng.SetOptions(uci.Options{
		Hash:    128,
		Ponder:  false,
		OwnBook: true,
		MultiPV: 4,
	})

	moves := ""
	if !isWhite {
		moves = <-nextMove
	}

	for {
		// set the starting position
		fmt.Printf("moves %s\n", moves)
		eng.SetMoves(moves)

		// set some result filter options
		resultOpts := uci.HighestDepthOnly | uci.IncludeUpperbounds | uci.IncludeLowerbounds
		result, _ := eng.Go(0, "", 100, resultOpts)
		myMove := result.BestMove
		results <- myMove
		opponent := <-nextMove
		moves = fmt.Sprintf("%s %s %s", moves, myMove, opponent)
	}
}

func doUci(w http.ResponseWriter) {
	nextMove1 := make(chan string)
	results1 := make(chan string)
	nextMove2 := make(chan string)
	results2 := make(chan string)

	go engine("5556", true, nextMove1, results1)
	go engine("5557", false, nextMove2, results2)

	for {
		white := <-results1
		if white == "(none)" {
			return
		}
		w.Write([]byte(white + " "))
		fmt.Println(white)
		nextMove2 <- white
		black := <-results2
		if black == "(none)" {
			return
		}
		w.Write([]byte(black + " "))
		fmt.Println(black)
		nextMove1 <- black
	}
}
