package main

import (
	"aoc/internal/utils"
	"fmt"
	"log"
	"os"

	"github.com/quartercastle/vector"
)

func DrawCityMapToFile(cityMap *CityMap, filename string) {
	output := ""
	for y := range cityMap.MapWidth {
		line := ""
		for x := range cityMap.MapWidth {
			vec := vector.Vector{float64(x), float64(y)}
			s := "."
			if utils.ContainsVector(cityMap.Antinodes, vec) {
				s = "#"
			}
			// for o, antennas := range cityMap.Antennas {
			// 	if utils.ContainsVector(antennas, vec) {
			// 		s = o
			// 	}
			// }
			line = line + s
		}
		output = fmt.Sprintf("%s\n%s", output, line)
	}

	err := os.WriteFile(filename, []byte(output), 0666)
	if err != nil {
		log.Fatal("Unable to write to file")
	}
}
