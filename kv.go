package main

type KeyValue struct {
    Key   string
    Value int
}

type MapFunction func(filename string, contents string) []KeyValue
type ReduceFunction func(key string, values []int) int