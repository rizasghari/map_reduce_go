package main

import "fmt"

func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
    contents := []string{
        "The quick brown fox",
        "jumps over the lazy dog",
        "The dog barks",
    }

    result := MapReduce(files, contents, WordCountMap, WordCountReduce)

    // Print results
    for word, count := range result {
        fmt.Printf("%s: %d\n", word, count)
    }
}