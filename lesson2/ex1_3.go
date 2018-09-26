package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {

	s := [...]int{1,2,3,4,5,6,7,8,9,10}

	//ex. 4.3
	reverse(&s)
	log.Println(s) //2018/09/24 11:20:07 [10 9 8 7 6 5 4 3 2 1]
	reverse(&s)
	log.Println(s) //2018/09/24 11:20:07 [1 2 3 4 5 6 7 8 9 10]

	//ex. 4.4
	rotate(s[:], 2)
	log.Println(s) //2018/09/24 11:20:07 [3 4 5 6 7 8 9 10 1 2]
	rotate(s[:], 5)
	log.Println(s) //2018/09/24 11:20:07 [8 9 10 1 2 3 4 5 6 7]
	rotate(s[:], 3)
	log.Println(s) //2018/09/24 11:20:07 [1 2 3 4 5 6 7 8 9 10]

	wordfreq()

}

//ex. 4.3
func reverse(s *[10]int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//ex. 4.4
func rotate(s []int, shift int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
		if shift - i - 1 > i {
			s[i], s[shift - i - 1] = s[shift - i - 1], s[i]
		}
		if j > shift + i {
			s[shift + i], s[j] = s[j], s[shift + i]
		}
 		s[i], s[j] = s[j], s[i]
	}
}

//ex. 4.9
func wordfreq() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		verb := strings.ToLower(scanner.Text())
		counts[verb]++
	}

	if err := scanner.Err(); err != nil {
		log.Printf("reading error: %v\n", err)
	}

	for verb, frq := range counts {
		log.Printf("%s: %d\n", verb, frq)
	}
}
