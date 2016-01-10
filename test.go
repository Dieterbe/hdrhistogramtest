//go:generate msgp

package main

import (
	"fmt"
	"github.com/codahale/hdrhistogram"
)

func record(h *hdrhistogram.Histogram, i int) {
	if err := h.RecordValue(int64(i)); err != nil {
		panic(err)
	}
}
func main() {
	// 0.1 ms to 600s= 6M 0.1ms
	// 1ms to 60s = 60k ms

	// 6M,  3 -> 0values: 14k  10M values: 17500
	// 6M,  2 -> 0values: 2300 10M values: 2700
	// 60k, 3 -> 0values: 7k   10M values: 10500
	// 60k, 2 -> 0values: 1400 10M values: 1800

	// for comparison: maxuint16 (2B) is 65k
	// every second, 1h worth = 3600 vals * 2 Bytes/val =7kB

	h := hdrhistogram.New(1, int64(60000), 2)

	//	for i := 0; i < M6; i++ {
	//		record(h, i)
	//	}
	//for i := 0; i < 2000000; i++ {
	//		record(h, 1234567/100)
	//	}
	//	for i := 0; i < 2000000; i++ {
	//		record(h, 1000/100)
	//	}
	//	for i := 0; i < 2000000; i++ {
	//		record(h, 200/100)
	//	}
	// 4M
	//	for i := 2000000; i < 6000000; i++ {
	//		record(h, i/100)
	//	}
	// 10 200
	// 20 200
	// 30 1000
	// 40 1000
	// 50 1M
	// 60 1M
	// 70 3M
	// 80 4M
	// 90
	// 100
	qs := []float64{
		10,
		20,
		30,
		40,
		50,
		60,
		70,
		//75,
		80,
		90,
		//95,
		//99,
		//99.9,
		//99.99,
		100,
	}

	fmt.Println("value at requested quantiltes:")

	for _, q := range qs {
		fmt.Println(q, h.ValueAtQuantile(q))
	}

	actual := h.CumulativeDistribution()

	fmt.Println("returned cumulative distribution:")
	for _, b := range actual {
		fmt.Printf("Count %10.d      Quantile %4.f ValueAt %10.d\n", b.Count, b.Quantile, b.ValueAt)
	}

	s := h.Export()
	snap := Snap{
		s.LowestTrackableValue,
		s.HighestTrackableValue,
		s.SignificantFigures,
		s.Counts,
	}

	in := make([]byte, 0)
	bin, err := snap.MarshalMsg(in)
	if err != nil {
		panic(err)
	}
	fmt.Println("serialized size", len(bin))

}

type Snap struct {
	LowestTrackableValue  int64
	HighestTrackableValue int64
	SignificantFigures    int64
	Counts                []int64
}
