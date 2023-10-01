package main

import (
	"log"
	"os"

	_ "github.com/cheldontk/matcher/matchers"
	"github.com/cheldontk/matcher/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
