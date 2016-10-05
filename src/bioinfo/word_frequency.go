package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

type WordFrequency struct {
	Word      string
	Frequency int
}

func (w WordFrequency) String() string {
	return fmt.Sprintf("%s: %d", w.Word, w.Frequency)
}

type ByFreq []*WordFrequency

func (a ByFreq) Len() int           { return len(a) }
func (a ByFreq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFreq) Less(i, j int) bool { return a[i].Frequency > a[j].Frequency }

func getKeys(hash map[string]int) (keys []string) {
	for key := range hash {
		keys = append(keys, key)
	}
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findWordFreq(base string, length int) []*WordFrequency {
	r := make(map[string]int)
	for c := range base {
		if c+length <= len(base) {
			word := string(base[c : c+length])
			r[word]++
		}
	}

	words := getKeys(r)
	wordFreqArray := make([]*WordFrequency, len(words))

	for i, w := range words {
		wf := &WordFrequency{Word: w, Frequency: r[w]}
		wordFreqArray[i] = wf
	}
	sort.Sort(ByFreq(wordFreqArray))

	return wordFreqArray
}
func main() {
	if len(os.Args) == 3 {
		file := os.Args[1]
		tmp := os.Args[2]
		wLength, err := strconv.Atoi(tmp)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		b, err := ioutil.ReadFile(file)
		check(err)
		str := string(b)

		fmt.Println(findWordFreq(str, wLength))
	} else {
		fmt.Println("Need file and word length.")
	}
}
