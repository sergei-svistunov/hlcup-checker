package main

import (
	"flag"
	"github.com/sergei-svistunov/hlcup-checker/phase"
	"log"
)

var (
	argServer    = flag.String("server", "http://localhost:8080", "Server address")
	argDataDir   = flag.String("data-dir", "./", "Directory with ammo and answers")
	argPhase     = flag.Int("phase", 1, "Number of phase")
	argMaxErrors = flag.Int("max-errors", 0, "Number of phase")
)

func main() {
	flag.Parse()

	if *argPhase < 1 || *argPhase > 3 {
		log.Fatalf("Invalid phase ID: %d", *argPhase)
	}

	p := phase.New(*argDataDir, 1)
	p.Check(*argServer, phase.CheckOpts{
		MaxErrors: *argMaxErrors,
	})
}
