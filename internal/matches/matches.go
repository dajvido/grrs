package matches

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/dajvido/grrs/internal/cl"
)

func InFile(args *cl.Args) ([]string, error) {
	file, err := os.Open(args.Path)
	defer file.Close()
	if err != nil {
		return []string{}, err
	}
	return scanFileForMatches(args.Pattern, file)
}

func scanFileForMatches(pattern string, file io.Reader) ([]string, error) {
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := string(scanner.Bytes())
		if strings.Contains(line, pattern) {
			lines = append(lines, strings.TrimSpace(line))
		}
	}
	return lines, scanner.Err()
}
