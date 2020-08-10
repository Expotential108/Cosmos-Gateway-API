// Package main exposes Rosetta API for Cosmos SDK as a standalone Gateway.
package main

import (
	"log"

	"github.com/tendermint/cosmos-rosetta-gateway/crg/interface/cli/crg/cmd"
)

func main() {
	if err := cmd.New().Execute(); err != nil {
		log.Fatal(err)
	}
}
