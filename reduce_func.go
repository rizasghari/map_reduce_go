package main

func WordCountReduce(key string, values []int) int {
    sum := 0
    for _, count := range values {
        sum += count
    }
    return sum
}