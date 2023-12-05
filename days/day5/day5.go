package day5

import (
	"AdventOfCode/common"
	"log"
	"math"
	"strconv"
	"strings"
)

func parseMap(almanac []string, curidx int) ([][]int64, int) {
	var m [][]int64
	i := 0
	for curidx < len(almanac) && almanac[curidx] != "" {
		m = append(m, []int64{})
		for _, vr := range strings.Fields(almanac[curidx]) {
			v, _ := strconv.ParseInt(vr, 10, 32)
			m[i] = append(m[i], v)
		}
		i += 1
		curidx += 1
	}
	return m, curidx + 2
}

func findLocation(maps [][][]int64, seed int64) int64 {
	for _, m := range maps {
		for _, c := range m {
			if c[1] <= seed && seed < c[1]+c[2] {
				seed = seed - c[1] + c[0]
				break
			}
		}
	}
	return seed
}

func Part1() {
	almanac := common.ReadInputFile("./days/day5/part1.in")
	var seeds []int64
	for _, seed := range strings.Fields(strings.Split(almanac[0], ":")[1]) {
		seed, _ := strconv.ParseInt(seed, 10, 64)
		seeds = append(seeds, seed)
	}
	var maps [][][]int64
	var m [][]int64
	idx := 3
	for i := 0; i < 7; i++ {
		m, idx = parseMap(almanac, idx)
		maps = append(maps, m)
	}
	var minLocation int64 = math.MaxInt
	for _, seed := range seeds {
		minLocation = min(findLocation(maps, seed), minLocation)

	}
	log.Printf("Final answer: %d", minLocation)
}

func Part2() {
	almanac := common.ReadInputFile("./days/day5/part1.in")
	var seedRanges [][]int64
	seedsRaw := strings.Fields(strings.Split(almanac[0], ":")[1])
	for i := 0; i < len(seedsRaw); i += 2 {
		seedStart, _ := strconv.ParseInt(seedsRaw[i], 10, 64)
		seedRange, _ := strconv.ParseInt(seedsRaw[i+1], 10, 64)
		seedRanges = append(seedRanges, []int64{seedStart, seedRange})
	}
	var maps [][][]int64
	var m [][]int64
	idx := 3
	for i := 0; i < 7; i++ {
		m, idx = parseMap(almanac, idx)
		maps = append(maps, m)
	}

	var minLocation int64 = math.MaxInt64
	for i, seedRange := range seedRanges {
		log.Printf("%d/%d: %v", i+1, len(seedRanges), seedRange)
		for seed := seedRange[0]; seed < seedRange[0]+seedRange[1]; seed++ {
			minLocation = min(findLocation(maps, seed), minLocation)
		}
	}
	log.Printf("Final answer: %d", minLocation)
}
