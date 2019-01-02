package nvm

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

// RandomLine returns whatever
func RandomLine(path string) (int, string) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	count := 0
	var line string

	// Count lines
	for {
		_, eof_error := reader.ReadString('\n')
		if eof_error != nil {
			break
		}
		count += 1
	}

	// Random-ness
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	line_num := r.Intn(count)

	// Back to the top of the file
	file.Seek(0, 0)

	// Count lines
	for i := 0; i < line_num+1; i++ {
		line, _ = reader.ReadString('\n')
		line = line[0 : len(line)-1]
	}

	return line_num, line
}
