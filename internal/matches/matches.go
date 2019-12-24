package matches

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/dajvido/grrs/internal/cl"
	"github.com/pkg/errors"
)

// InFile looks for matches for provided in *cl.Args.Pattern pattern
// inside a file from provided in *cl.Args.Path path
// returns either all matching lines or errors when either fails to open a file,
// first I/O error from scanning a file, or a token too large to fit in the buffer.
func InFile(args *cl.Args) ([]string, error) {
	file, err := os.Open(args.Path)
	defer file.Close()
	if err != nil {
		return nil, errors.Wrap(err, args.Path)
	}
	lines, err := scanFileForMatches(args.Pattern, file)
	if err != nil {
		return nil, errors.Wrap(err, args.Path)
	}
	return lines, nil
}

// scanFileForMatches scan provided file line by line to match each line with provided pattern.
// It returns all the lines that match pattern or error from scanner which is either
// first I/O error from scanning a file, or a token too large to fit in the buffer.
func scanFileForMatches(pattern string, file io.Reader) ([]string, error) {
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := string(scanner.Bytes())
		if strings.Contains(line, pattern) {
			lines = append(lines, strings.TrimSpace(line))
		}
	}
	err := scanner.Err()
	if err != nil {
		return nil, errors.Wrap(err, pattern)
	}
	return lines, nil
}
