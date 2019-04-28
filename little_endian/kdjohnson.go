package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func endian(s string, base int) string {
	// prepend for easier byte splitting

	if base == 10 {
		b, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		x := fmt.Sprintf("%x", b)
		s = x
	}

	if len(s)%2 != 0 {
		var sb strings.Builder
		sb.WriteString("0")
		sb.WriteString(s)
		s = sb.String()
	}

	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	c := make([]string, len(s))
	copy(c, strings.Split(s, ""))
	if decoded[0] < decoded[1] {
		tmp, tmp2 := c[0], c[1]
		c[0], c[1] = c[2], c[3]
		c[2], c[3] = tmp, tmp2
	}

	if base == 10 {
		tmp := strings.Join(c[:], "")
		i, err := strconv.ParseInt(tmp, 16, 64)
		if err != nil {
			log.Fatal(err)
		}
		return fmt.Sprintf("%d", i)
	}

	return strings.Join(c[:], "")
}

func main() {
	fmt.Println(endian("35582", 10))
	fmt.Println(endian("8AFE", 16))
	fmt.Println(endian("332", 16))
}
