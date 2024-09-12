package main

import "sync"

func MapReduce(files []string, contents []string, mapF MapFunction, reduceF ReduceFunction) map[string]int {
    intermediate := make(chan KeyValue, 10000)
    var wg sync.WaitGroup

    // Map phase
    for i, file := range files {
        wg.Add(1)
        go func(file string, content string) {
            defer wg.Done()
            kvs := mapF(file, content)
            for _, kv := range kvs {
                intermediate <- kv
            }
        }(file, contents[i])
    }

    go func() {
        wg.Wait()
        close(intermediate)
    }()

    // Shuffle phase
    intermediateMap := make(map[string][]int)
    for kv := range intermediate {
        intermediateMap[kv.Key] = append(intermediateMap[kv.Key], kv.Value)
    }

    // Reduce phase
    result := make(map[string]int)
    for key, values := range intermediateMap {
        result[key] = reduceF(key, values)
    }

    return result
}
