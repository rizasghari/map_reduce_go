package main

import "strings"

func WordCountMap(filename string, contents string) []KeyValue {
    words := strings.Fields(contents)
    kvs := make([]KeyValue, 0, len(words))
    for _, word := range words {
        kvs = append(kvs, KeyValue{strings.ToLower(word), 1})
    }
    return kvs
}