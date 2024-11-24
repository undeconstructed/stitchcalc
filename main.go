package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func quit(code int, msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg, args...)
	fmt.Fprintln(os.Stderr)
	os.Exit(code)
}

// var ktog = regexp.MustCompile(`^k\dtog$`)

func main() {
	in := bufio.NewReader(os.Stdin)

	ll := 0

	for {
		line, _, err := in.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			quit(1, "read: %s", err)
		}

		ss := strings.Split(string(line), ",")

		a, b := 0, 0

		for _, s := range ss {
			if s == "" {
				continue
			}

			s = strings.ToLower(s)
			s = strings.TrimSpace(s)

			switch {
			case s == "mk":
				a += 1
				b += 2
			case s == "k2tog":
				a += 2
				b += 1
			case s == "k3tog":
				a += 3
				b += 1
			case s == "psso":
				a += 0
				b -= 1
			case strings.HasPrefix(s, "sl"):
				n, err := strconv.Atoi(s[3:])
				if err != nil {
					quit(2, "parse: %s, %s", s, err)
				}

				a += n
				b += n
			case strings.HasPrefix(s, "k"):
				n, err := strconv.Atoi(s[1:])
				if err != nil {
					quit(2, "parse: %s, %s", s, err)
				}

				a += n
				b += n
			case s == "yfwd":
				b += 1
			case s == "skpo":
				a += 2
				b += 1
			default:
				quit(2, "unknown: %s", s)
			}
		}

		if ll == 0 {
			ll = a
			fmt.Println("length:", ll)
		}

		if a != b || ll != a {
			fmt.Printf("PROBLEM:")
		}

		fmt.Printf("%3d -> %3d\n", a, b)
	}
}
