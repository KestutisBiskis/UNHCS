package main

import (
	"fmt"
	"strings"
	"time"

	"mapreduceproject.eparker.dev/dataload"
	"mapreduceproject.eparker.dev/mapreduce"
)

// func mapFunc(doc dataload.Bucket) map[string]int {
// 	wordCounts := make(map[string]int)

// 	for _, word := range doc.Words {
// 		wordCounts[word]++
// 	}

// 	return wordCounts
// }

// func shuffleFunc(wordCounts map[string]int) []map[string]int {
// 	return []map[string]int{wordCounts}
// }

// func reduceFunc(wordCounts []map[string]int) map[string]int {
// 	totalWordCounts := make(map[string]int)

// 	for _, wordCount := range wordCounts {
// 		for word, count := range wordCount {
// 			totalWordCounts[word] += count
// 		}
// 	}

// 	return totalWordCounts
// }

type WordOccurrence struct {
	Word  string
	Count int
	URLs  map[string]int
}

func mapFunc(bucket dataload.Bucket) []WordOccurrence {
	wordMap := make(map[string]*WordOccurrence)

	for _, word := range bucket.Words {
		if occurrence, exists := wordMap[word]; exists {
			occurrence.Count++
			occurrence.URLs[bucket.URL]++
		} else {
			wordMap[word] = &WordOccurrence{
				Word:  word,
				Count: 1,
				URLs:  map[string]int{bucket.URL: 1},
			}
		}
	}

	result := make([]WordOccurrence, 0, len(wordMap))
	for _, occurrence := range wordMap {
		result = append(result, *occurrence)
	}
	return result
}

func shuffleFunc(occurrences []WordOccurrence) [][]WordOccurrence {
	return [][]WordOccurrence{occurrences}
}

func reduceFunc(occurrenceGroups [][]WordOccurrence) map[string]WordOccurrence {
	result := make(map[string]WordOccurrence)

	for _, group := range occurrenceGroups {
		for _, occurrence := range group {
			if existing, exists := result[occurrence.Word]; exists {
				existing.Count += occurrence.Count
				for url, count := range occurrence.URLs {
					existing.URLs[url] += count
				}
				result[occurrence.Word] = existing
			} else {
				result[occurrence.Word] = occurrence
			}
		}
	}

	return result
}

func finalReducer(occurrenceGroups []map[string]WordOccurrence) map[string]WordOccurrence {
	result := make(map[string]WordOccurrence)

	for _, group := range occurrenceGroups {
		for word, occurrence := range group {
			if existing, exists := result[word]; exists {
				existing.Count += occurrence.Count
				for url, count := range occurrence.URLs {
					existing.URLs[url] += count
				}
				result[word] = existing
			} else {
				result[word] = occurrence
			}
		}
	}

	return result
}

func formatBytes(bytes int) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%dB", bytes)
	}

	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f%cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func MainDocuments() {
	start := time.Now()
	var docs []dataload.Document = dataload.LoadAllDocuments(5000)
	fmt.Printf("Loaded %d document(s) in %s, %s\n", len(docs), time.Since(start), formatBytes(dataload.SizeOfDocuments(docs)))

	var buckets []dataload.Bucket

	for _, doc := range docs {
		buckets = append(buckets, doc.Buckets...)
	}

	start = time.Now()
	searchDB := mapreduce.MapReduce(mapFunc, shuffleFunc, reduceFunc, finalReducer, buckets)
	fmt.Printf("MapReduce completed in %s\n", time.Since(start))

	// var word string

	// for {
	// 	fmt.Print("Enter a word to find its count: ")
	// 	_, err := fmt.Scanln(&word)

	// 	if err != nil {
	// 		fmt.Println("Error reading from stdin:", err)
	// 		return
	// 	}

	// 	result, ok := searchDB[word]

	// 	if ok {
	// 		var urls []string

	// 		for url, count := range result.URLs {
	// 			urls = append(urls, fmt.Sprintf("- %s (%d)", url, count))
	// 		}

	// 		fmt.Printf("Word \"%s\" found %d times in %d document(s):\n%s\n", word, result.Count, len(result.URLs), strings.Join(urls, "\n"))

	// 	} else {
	// 		fmt.Printf("Word \"%s\" not found\n", word)
	// 	}
	// }

	// Find documents and sort based on an input phrase
	var phrase string

	for {
		fmt.Print("Enter a phrase to search for in the documents: ")
		_, err := fmt.Scanf("%s", &phrase)

		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			return
		}

		var words []string = strings.Fields(phrase)
		var results []WordOccurrence

		// Implement a thing to figure out the best results
	}
}

func MainNumberReducer() {
	var numbers []int = make([]int, 1024*1024)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i
	}

	start := time.Now()
	sum := mapreduce.MapReduce(
		func(n int) int { return n },
		func(n int) []int { return []int{n} },
		func(ns []int) int {
			sum := 0
			for _, n := range ns {
				sum += n
			}
			return sum
		},
		func(sums []int) int {
			sum := 0

			for _, s := range sums {
				sum += s
			}

			return sum
		},
		numbers,
	)

	fmt.Println("Time taken: ", time.Since(start))
	fmt.Println("Sum:", sum)
}

func main() {
	MainDocuments()
}
