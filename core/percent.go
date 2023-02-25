package core

import (
	"bytes"
	"fmt"
)

var pctls = []float64{10, 25, 50, 75, 90, 95, 99}

func percentiles(sizes []int, sizeToCount map[int]int) (data []int) {
	data = make([]int, len(pctls))
	j := 0
	n := 0
	for _, count := range sizeToCount {
		n += count
	}
	cur := 0
	for i := 0; i < len(sizes) && j < len(pctls); i++ {
		c := sizeToCount[sizes[i]]
		current := float64(cur+c) * 100.0 / float64(n)
		if current >= pctls[j] {
			data[j] = sizes[i]
			j++
		}
		cur += c
	}
	return data
}

func PrintPercent(sizes []int, sizeToCount map[int]int) string {
	data := percentiles(sizes, sizeToCount)
	var buffer bytes.Buffer
	buffer.WriteString("\nSize distribution:\n")
	for i := 0; i < len(pctls); i++ {
		buffer.WriteString(fmt.Sprintf("  %d%% in %s.\n", int(pctls[i]), ReadableSize(data[i])))
	}
	return buffer.String()
}
