//Ecological Bin Packing
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var a, b, c, d, e, f, g, h, i int
	_, err := fmt.Scan(&a, &b, &c, &d, &e, &f, &g, &h, &i) //get the 9 digits
	for err == nil {
		line := []int{a, b, c, d, e, f, g, h, i}

		bins := [][]int{line[0:3], line[3:6], line[6:9]} //split into 3 bins

		bottle_colors := map[string]int{ //assign B G C color code to indices
			"B": 0,
			"G": 1,
			"C": 2,
		}

		colored_bins := [][]string{ //permutations of color patterns
			{"B", "G", "C"},
			{"B", "C", "G"},
			{"G", "C", "B"},
			{"G", "B", "C"},
			{"C", "B", "G"},
			{"C", "G", "B"},
		}

		total_moves := []int{}
		for _, pattern := range colored_bins { //get the 1st pattern from the permutations ex. ["B" "G" "C"]
			moved_bottles := 0
			for i, letter := range pattern { //get the index and letter from the patter ex. "B" from ["B" "G" "C"]
				for j, bin := range bins { //get the index and one bin from our 3 bins ex. [1, 2, 3]
					if i == j { //match the letters to the bins by using i (index from letters) and j (index from bins) -> B:[1,2,3] G:[4,5,6] C:[7,8,9]
						for k, num := range bin { //get the index, and number of bottles from a bin
							if k != bottle_colors[letter] { //color matching: if the index of the bottle doesnt match the index of the bottle color map -> we can move the bottle
								moved_bottles += num
							}
						}
					}
				}
			}
			total_moves = append(total_moves, moved_bottles) //append all moves to one list
		}

		/*
			map an index to the minimum moves
			ex.
			from the list of total move [10 20 30]
			index  value
			0       10
			1       20
			2       30
			3       10

			we get the index of the minimum value and do a mapping:
			{0: 10, 3: 10,}
		*/
		m := make(map[int]int)
		min_moves := total_moves[0]
		for index, move := range total_moves {
			if move <= min_moves {
				min_moves = move
				m[index] = min_moves
			}
		}
		/*
			get the keys of the minimum values
			and make them into a list
			ex.
			from: {0: 10, 3: 10,}
			to: [0 3]
		*/
		pattern_index := []int{}
		for key, value := range m {
			if min_moves == value {
				pattern_index = append(pattern_index, key)
			}
		}
		/*
			use the list of indices [0 3] as indices of the
			colored bins permutation
			ex. [0 3] for index 0
			colored_bins[0] == ["B", "G", "C"]
			Join the list of strings into one string -> BGC
			then append to answers list
			the answers list will be sorted to provide the
			alphabetically first string if there are the same min moves
		*/
		answers := []string{}
		for _, value := range pattern_index {
			ans := strings.Join(colored_bins[value], "")
			answers = append(answers, ans)
		}
		sort.Strings(answers)
		fmt.Println(answers[0], min_moves)
		_, err = fmt.Scan(&a, &b, &c, &d, &e, &f, &g, &h, &i)
	}
}

//for cmd: go run main.go < 102.in > 102.out
// for powershell: type .\102.in | go run main.go > 102.out
//refer to https://www.udebug.com/ to verify inputs
