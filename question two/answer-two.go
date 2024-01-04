package main

import (
	"container/heap"
	"fmt"
)

type CharFrequency struct {
  char rune
  freq int
}

type PriorityQueue []*CharFrequency

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
  return pq[i].freq > pq[j].freq
}

func (pq PriorityQueue) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
  item := x.(*CharFrequency)
  *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
  old := *pq
  n := len(old)
  item := old[n-1]
  *pq = old[0 : n-1]
  return item
}

func reorganizeString(s string) string {
  freqMap := make(map[rune]int)
  for _, char := range s {
    freqMap[char]++
  }

  pq := make(PriorityQueue, len(freqMap))
  i := 0
  for char, freq := range freqMap {
    pq[i] = &CharFrequency{char, freq}
    i++
  }
  heap.Init(&pq)

  result := []rune{}
  var lastChar *CharFrequency
  for pq.Len() > 0 {
    currentChar := heap.Pop(&pq).(*CharFrequency)
    result = append(result, currentChar.char)
    currentChar.freq--

    if lastChar != nil && lastChar.freq > 0 {
      heap.Push(&pq, lastChar)
    }

    if currentChar.freq > 0 {
      lastChar = currentChar
    } else {
      lastChar = nil
    }
  }

  if lastChar != nil && lastChar.freq > 0 {
    return ""
  }

  return string(result)
}

func main() {
  fmt.Println("Example 1:", reorganizeString("aab"))
  fmt.Println("Example 2:", reorganizeString("aaabbz"))
}