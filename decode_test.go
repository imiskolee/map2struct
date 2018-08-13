package map2struct

import (
	"fmt"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {

	m := map[string]interface{}{
		"A" : 1,
		"B" : 1.1,
		"C" : "1.3",
		"D" : true,
		"E" : map[string]interface {} {
			"G" : 1,
			"I" : "2018-01-01 00:00:00",
			"J" : "2018-01-01 00:00:00",
			"K" : 1.1,
			"L" : "abcd",
		},
		"M" : []interface{} {1,2,3},
		"N" : []interface{} {"4","5","6"},
		"o" : 1,
	}

	var parsed struct {
		A int
		B float64
		C string
		D bool
		E *struct {
			G int
			I time.Time
			J *time.Time
			K float32
			L string
		}
		M []int
		N []string
		O int `json:"o"`
	}

	err := Decode(m,&parsed)
	fmt.Print(err,parsed,parsed.E)
}

