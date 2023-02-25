package core

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gosuri/uilive"
	"go.etcd.io/etcd/api/v3/mvccpb"
)

const (
	barChar = "*"
)

type Stats struct {
	Largest  int
	Smallest int
	Average  int
	Count    int
	Total    int

	sizes       []int
	countLock   sync.RWMutex
	sizeToCount map[int]int
}

type Report interface {
	Results() chan<- []*mvccpb.KeyValue
	Run() <-chan string
	DynamicOutput()
}

type SizeOf func(*mvccpb.KeyValue) int

type report struct {
	results     chan []*mvccpb.KeyValue
	stats       Stats
	bucketCount int
	sizeOf      SizeOf
	writer      *uilive.Writer
	processOver atomic.Bool
	dynamicOnce sync.Once
}

func (r *report) Results() chan<- []*mvccpb.KeyValue { return r.results }

func (r *report) Run() <-chan string {
	r.writer.Start()
	donec := make(chan string, 1)
	go func() {
		defer close(donec)
		r.processResults()
		if r.stats.Count <= 0 {
			_, _ = fmt.Fprintln(r.writer, "empty data")
		}
		r.finalString()
		r.writer.Stop()
	}()

	return donec
}

func (r *report) processResults() {
	for res := range r.results {
		r.processResult(res)
	}
	r.processOver.Store(true)
	time.Sleep(time.Millisecond * 100)
}

func (r *report) processResult(res []*mvccpb.KeyValue) {
	l := len(res)
	if l == 0 {
		return
	}
	r.stats.Count += l
	for _, kv := range res {
		s := r.sizeOf(kv)
		r.stats.Smallest = Min(r.stats.Smallest, s)
		r.stats.Largest = Max(r.stats.Largest, s)
		r.stats.Total += s
		r.stats.countLock.Lock()
		_, ok := r.stats.sizeToCount[s]
		if !ok {
			r.stats.sizes = append(r.stats.sizes, s)
		}
		r.stats.sizeToCount[s] += 1
		r.stats.countLock.Unlock()
	}
	r.stats.Average = r.stats.Total / r.stats.Count
}

func (r *report) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Summary:\n")
	buffer.WriteString(fmt.Sprintf("  Count:\t%d.\n", r.stats.Count))
	buffer.WriteString(fmt.Sprintf("  Total:\t%s.\n", ReadableSize(r.stats.Total)))
	buffer.WriteString(fmt.Sprintf("  Smallest:\t%s.\n", ReadableSize(r.stats.Smallest)))
	buffer.WriteString(fmt.Sprintf("  Largest:\t%s.\n", ReadableSize(r.stats.Largest)))
	buffer.WriteString(fmt.Sprintf("  Average:\t%s.\n", ReadableSize(r.stats.Average)))

	sort.Ints(r.stats.sizes)
	buffer.WriteString(r.histogram())
	r.stats.countLock.RLock()
	buffer.WriteString(PrintPercent(r.stats.sizes, r.stats.sizeToCount))
	r.stats.countLock.RUnlock()

	return buffer.String()
}

func (r *report) DynamicOutput() {
	r.dynamicOnce.Do(func() {
		go func() {
			for {
				if r.processOver.Load() {
					return
				}
				r.dynamicString()
				time.Sleep(time.Millisecond * 100)
			}
		}()
	})
}

func (r *report) dynamicString() {
	if r.stats.Count <= 0 {
		return
	}

	_, _ = fmt.Fprint(r.writer, r.String())
	_ = r.writer.Flush()
}

func (r *report) finalString() {
	if r.stats.Count <= 0 {
		return
	}

	_, _ = fmt.Fprint(r.writer.Bypass(), r.String())
}

func (r *report) histogram() string {
	buckets := make([]int, r.bucketCount+1)
	counts := make([]int, r.bucketCount+1)
	bs := (r.stats.Largest - r.stats.Smallest) / r.bucketCount
	for i := 0; i < r.bucketCount; i++ {
		buckets[i] = r.stats.Smallest + bs*i
	}
	buckets[r.bucketCount] = r.stats.Largest

	var bi int
	var max int
	for i := 0; i < len(r.stats.sizes); {
		s := r.stats.sizes[i]
		if s <= buckets[bi] {
			i++
			r.stats.countLock.RLock()
			counts[bi] += r.stats.sizeToCount[s]
			r.stats.countLock.RUnlock()
			if max < counts[bi] {
				max = counts[bi]
			}
		} else if bi < len(buckets)-1 {
			bi++
		}
	}
	var buffer bytes.Buffer
	buffer.WriteString("\nSize histogram:\n")
	for i := 0; i < len(buckets); i++ {
		var barLen int
		if max > 0 {
			barLen = counts[i] * 40 / max
		}
		buffer.WriteString(fmt.Sprintf("  %s [%d]\t|%v\n", ReadableSize(buckets[i]), counts[i], strings.Repeat(barChar, barLen)))
	}
	return buffer.String()
}

func NewReport(bc int, of SizeOf) Report {
	return &report{
		results: make(chan []*mvccpb.KeyValue),
		stats: Stats{
			Largest:     -1,
			Smallest:    math.MaxInt32,
			sizeToCount: make(map[int]int),
		},
		bucketCount: bc,
		sizeOf:      of,
		writer:      uilive.New(),
	}
}
