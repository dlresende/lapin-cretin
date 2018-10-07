package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	cfenv "github.com/dlresende/lapin-cretin/cfenv"
)

func main() {
	nbOfConn, nbOfChPerConn, err := ParseArgs(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	uri := cfenv.GetAmqPUri()
	forever := make(chan interface{})
	CreateGhostChannels(uri, nbOfConn, nbOfChPerConn, forever, 1)
}

func ParseArgs(args []string) (int, int, error) {
	if len(args) != 3 {
		return -1, -1, errors.New(fmt.Sprintf("usage: %s <number of connections> <number of channels per connection>\n", args[0]))
	}

	nbOfConn, err := strconv.Atoi(args[1])

	if err != nil {
		return -1, -1, errors.New(fmt.Sprintf("%s is not an integer\n", args[1]))
	}

	nbOfChPerConn, err := strconv.Atoi(args[2])

	if err != nil {
		return -1, -1, errors.New(fmt.Sprintf("%s is not an integer\n", args[2]))
	}

	return nbOfConn, nbOfChPerConn, nil
}
