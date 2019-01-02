package main

import (
	"flag"
	"github.com/whatever/nvm"
	"log"
)

func main() {
	var path = flag.String("path", "", "/path/to/file")
	flag.Parse()

	hits := make(map[string]int)

	// ...
	for i := 0; i < 10000; i++ {
		_, line := nvm.RandomLine(*path)
		hits[line] += 1
	}

	for k, v := range hits {
		log.Printf("%s -> %d\n", k, v)
	}

	log.Printf("Hit LEN = %d", len(hits))
}
