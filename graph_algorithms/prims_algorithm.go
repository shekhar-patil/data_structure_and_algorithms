package main

import (
    "container/heap"
    "fmt"
)

type MinHeap [][]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(value any)    { *h = append(*h, value.([]int)) }
func (h *MinHeap) Pop() any {
    l := len(*h)
    val := (*h)[l-1]
    *h = (*h)[:l-1]
    return val
}

func printMst(edges [][][]int) [][][]int {
    mst := make([][][]int, len(edges))
    h := &MinHeap{}
    heap.Init(h)

    for _, edge := range edges[0] {
        heap.Push(h, []int{0, edge[0], edge[1]}) // from, to, weight
    }

    visited := map[int]bool{0: true}

    for h.Len() > 0 {
        edge := heap.Pop(h).([]int)
        from, to, weight := edge[0], edge[1], edge[2]

        if visited[to] {
            continue
        }

        visited[to] = true
        mst[from] = append(mst[from], []int{to, weight})
        mst[to] = append(mst[to], []int{from, weight})

        for _, next := range edges[to] {
            if !visited[next[0]] {
                heap.Push(h, []int{to, next[0], next[1]})
            }
        }
    }

    return mst
}

func main() {
    edges := [][][]int{
        {{1, 12}, {4, 6}},        // Node 0
        {{0, 12}, {2, 8}},        // Node 1
        {{1, 8}, {3, 5}, {4, 3}}, // Node 2
        {{2, 5}, {4, 2}},         // Node 3
        {{0, 6}, {2, 3}, {3, 2}}, // Node 4
    }

    mst := printMst(edges)
    fmt.Println("Minimum Spanning Tree:")
    for i, connections := range mst {
        for _, conn := range connections {
            fmt.Printf("%d -- %d (weight %d)\n", i, conn[0], conn[1])
        }
    }
}