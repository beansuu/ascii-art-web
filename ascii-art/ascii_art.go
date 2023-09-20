package ascii_art

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const SymbolSize = 8

func MakeSymbolMap(style string) (map[rune][]string, error) {
	symbolMap := make(map[rune][]string)
	file, err := ioutil.ReadFile(style)
	if err != nil {
		return nil, err
	}

	artSymbols := strings.Split(string(file), "\n")
	for symbol := ' '; symbol <= '~'; symbol++ {
		j, slice := int(symbol-' '), make([]string, SymbolSize)
		for i := range slice {
			slice[i] = artSymbols[j*(SymbolSize+1)+i+1]
		}
		symbolMap[symbol] = slice
	}

	return symbolMap, nil
}

func PrintArt(input string, style string) (string, error) {
	symbolMap, err := MakeSymbolMap(style)
	if err != nil {
		return "", err
	}

	var result strings.Builder
	for _, line := range strings.Split(input, "\\n") {
		for i := 0; i < SymbolSize; i++ {
			for _, symbol := range line {
				result.WriteString(symbolMap[symbol][i])
			}
			result.WriteString("\n")
		}
		if line == "" {
			result.WriteString("\n")
		}
	}
	return result.String(), nil
}

//