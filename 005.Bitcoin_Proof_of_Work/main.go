package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/AndersonQ/100-days-of-algorithms/005.Bitcoin_Proof_of_Work/block"
)

type cliArgs struct {
	S, File string
	Z       uint
	Verbose bool
}

func buildPrefix(n uint) string {
	prefix := ""

	for n != 0 {
		prefix += "0"
		n--
	}

	return prefix
}

func parseArgs() cliArgs {
	args := cliArgs{}

	flag.UintVar(&args.Z, "z", 4, "The amount of zeros the hash should start with")
	flag.StringVar(&args.S, "s", "Hello, world!", "The string to fount its proof of work (to mine). If s is defined, f will be ignored")
	flag.StringVar(&args.File, "f", "", "A file path to found the file proof of work (to mine). It is ignored if s is defined")
	flag.BoolVar(&args.Verbose, "v", false, "Verbose mode. Prints the the data used to find the proof of work. WARNING: Do not user verbose mode if the data is too long!")

	flag.Parse()

	return args
}

func main() {
	var b block.Block

	args := parseArgs()

	if args.File == "" {
		b = block.Mine([]byte(args.S), buildPrefix(args.Z))
	} else {
		dat, err := ioutil.ReadFile(args.File)
		if err != nil {
			panic(err)
		}

		b = block.Mine(dat, buildPrefix(args.Z))
	}

	if args.Verbose {
		fmt.Printf("Nonce: %d\nHash: %x\nData: '%s'\n", b.Nonce, b.Hash, string(b.Data))
	} else {
		fmt.Printf("Nonce: %d\nHash: %x\n", b.Nonce, b.Hash)
	}
}
