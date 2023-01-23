package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/freeeve/uci"
)

func doUci(w http.ResponseWriter) {
	eng, err := uci.NewEngine("/opt/homebrew/bin/stockfish")
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

	// set the starting position
	eng.SetFEN("rnb4r/ppp1k1pp/3bp3/1N3p2/1P2n3/P3BN2/2P1PPPP/R3KB1R b KQ - 4 11")

	// set some result filter options
	resultOpts := uci.HighestDepthOnly | uci.IncludeUpperbounds | uci.IncludeLowerbounds
	results, _ := eng.GoDepth(10, resultOpts)

	// print it (String() goes to pretty JSON for now)
	fmt.Fprintln(w, results.String())
}
