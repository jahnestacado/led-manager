package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/jahnestacado/cable"
)

type action struct {
	Sequence []int
	Duration time.Duration
}

var (
	pattern0 = []action{
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 100 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 1, 2, 0, 0, 0, 0}, 100 * time.Millisecond},
		action{[]int{0, 0, 0, 1, 1, 2, 2, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 1, 1, 1, 2, 2, 2, 0, 0}, 50 * time.Millisecond},
		action{[]int{0, 1, 1, 1, 1, 2, 2, 2, 2, 0}, 50 * time.Millisecond},
		action{[]int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2}, 50 * time.Millisecond},
		action{[]int{0, 1, 1, 1, 1, 2, 2, 2, 2, 0}, 50 * time.Millisecond},
		action{[]int{0, 0, 1, 1, 1, 2, 2, 2, 0, 0}, 50 * time.Millisecond},
		action{[]int{0, 0, 0, 1, 1, 2, 2, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 1, 2, 0, 0, 0, 0}, 100 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 100 * time.Millisecond},

		action{[]int{0, 0, 0, 0, 1, 2, 0, 0, 0, 0}, 100 * time.Millisecond},
		action{[]int{0, 0, 0, 1, 1, 2, 2, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 1, 1, 1, 2, 2, 2, 0, 0}, 50 * time.Millisecond},
		action{[]int{0, 1, 1, 1, 1, 2, 2, 2, 2, 0}, 50 * time.Millisecond},
		action{[]int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2}, 50 * time.Millisecond},
		action{[]int{0, 1, 1, 1, 1, 2, 2, 2, 2, 0}, 50 * time.Millisecond},
		action{[]int{0, 0, 1, 1, 1, 2, 2, 2, 0, 0}, 50 * time.Millisecond},
		action{[]int{0, 0, 0, 1, 1, 2, 2, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 1, 2, 0, 0, 0, 0}, 100 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 100 * time.Millisecond},

		action{[]int{0, 0, 0, 0, 1, 2, 0, 0, 0, 0}, 100 * time.Millisecond},
		action{[]int{0, 0, 0, 1, 1, 2, 2, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 1, 1, 1, 2, 2, 2, 0, 0}, 50 * time.Millisecond},
		action{[]int{0, 1, 1, 1, 1, 2, 2, 2, 2, 0}, 50 * time.Millisecond},
		action{[]int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2}, 50 * time.Millisecond},
		action{[]int{0, 1, 1, 1, 1, 2, 2, 2, 2, 0}, 50 * time.Millisecond},
		action{[]int{0, 0, 1, 1, 1, 2, 2, 2, 0, 0}, 50 * time.Millisecond},
		action{[]int{0, 0, 0, 1, 1, 2, 2, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 1, 2, 0, 0, 0, 0}, 100 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 100 * time.Millisecond},
	}
	pattern1 = []action{
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{2, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 0, 0, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 0, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 2, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 2, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 2, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 2, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 2}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 2}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, 80 * time.Millisecond},

		action{[]int{2, 0, 0, 0, 0, 0, 0, 0, 0, 1}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 0, 0, 0, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 0, 0, 0, 0, 0, 0, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 0, 0, 0, 0, 0, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 2, 0, 0, 0, 0, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 2, 0, 0, 0, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 2, 0, 0, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 2, 0, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 2, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 2, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1}, 80 * time.Millisecond},

		action{[]int{2, 0, 0, 0, 0, 0, 0, 0, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 0, 0, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 0, 0, 0, 0, 0, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 0, 0, 0, 0, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 2, 0, 0, 0, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 2, 0, 0, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 2, 0, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 2, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 2, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1}, 80 * time.Millisecond},

		action{[]int{2, 0, 0, 0, 0, 0, 0, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 0, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 0, 0, 0, 0, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 0, 0, 0, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 2, 0, 0, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 2, 0, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 2, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 2, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1}, 80 * time.Millisecond},

		action{[]int{2, 0, 0, 0, 0, 0, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 0, 0, 0, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 0, 0, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 2, 0, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 2, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 2, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1}, 80 * time.Millisecond},

		action{[]int{2, 0, 0, 0, 0, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 0, 0, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 0, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 2, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 2, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},

		action{[]int{2, 0, 0, 0, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 0, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},

		action{[]int{2, 0, 0, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},

		action{[]int{2, 0, 1, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 2, 1, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 2, 1, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},

		action{[]int{2, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{2, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 300 * time.Millisecond},
	}
	pattern2 = []action{
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 100 * time.Millisecond},
		action{[]int{2, 2, 2, 0, 0, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 2, 2, 2, 0, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 2, 2, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 2, 2, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 2, 2, 2, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 2, 2, 2, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 2, 2, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 2, 2, 2}, 160 * time.Millisecond},

		action{[]int{0, 0, 0, 0, 0, 0, 2, 2, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 2, 2, 2, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 2, 2, 2, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 2, 2, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 2, 2, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 2, 2, 2, 0, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{2, 2, 2, 0, 0, 0, 0, 0, 0, 0}, 160 * time.Millisecond},
		action{[]int{0, 2, 2, 2, 0, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 2, 2, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 2, 2, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 2, 2, 2, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 2, 2, 2, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 2, 2, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 2, 2, 2}, 160 * time.Millisecond},

		action{[]int{0, 0, 0, 0, 0, 0, 2, 2, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 2, 2, 2, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 2, 2, 2, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 2, 2, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 2, 2, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 2, 2, 2, 0, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{2, 2, 2, 0, 0, 0, 0, 0, 0, 0}, 160 * time.Millisecond},
		action{[]int{0, 2, 2, 2, 0, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 2, 2, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 2, 2, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 2, 2, 2, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 2, 2, 2, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 2, 2, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 2, 2, 2}, 160 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 2, 2, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 2, 2, 2, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 2, 2, 2, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 2, 2, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 2, 2, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 2, 2, 2, 0, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{2, 2, 2, 0, 0, 0, 0, 0, 0, 0}, 160 * time.Millisecond},
		action{[]int{0, 2, 2, 2, 0, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 2, 2, 2, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 2, 2, 0, 0, 0, 0}, 300 * time.Millisecond},
		action{[]int{0, 0, 0, 1, 1, 1, 0, 0, 0, 0}, 300 * time.Millisecond},
		action{[]int{0, 0, 0, 2, 2, 2, 0, 0, 0, 0}, 300 * time.Millisecond},
		action{[]int{0, 0, 0, 1, 1, 1, 0, 0, 0, 0}, 300 * time.Millisecond},
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 100 * time.Millisecond},
	}
	pattern3 = []action{
		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 0, 0, 2, 0}, 1 * time.Second},
		action{[]int{0, 2, 1, 0, 0, 0, 0, 0, 2, 0}, 150 * time.Millisecond},
		action{[]int{0, 2, 0, 1, 0, 0, 0, 0, 2, 0}, 150 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 1, 0, 0, 0, 2, 0}, 150 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 1, 0, 0, 2, 0}, 150 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 1, 0, 2, 0}, 150 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 0, 1, 2, 0}, 150 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 1, 0, 2, 0}, 150 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 1, 0, 0, 2, 0}, 150 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 1, 0, 0, 0, 2, 0}, 150 * time.Millisecond},
		action{[]int{0, 2, 0, 1, 0, 0, 0, 0, 2, 0}, 150 * time.Millisecond},
		action{[]int{0, 2, 1, 0, 0, 0, 0, 0, 2, 0}, 150 * time.Millisecond},

		action{[]int{0, 2, 0, 1, 0, 0, 0, 0, 2, 0}, 100 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 1, 0, 0, 0, 2, 0}, 100 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 1, 0, 0, 2, 0}, 100 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 1, 0, 2, 0}, 100 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 0, 1, 2, 0}, 100 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 1, 0, 2, 0}, 100 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 1, 0, 0, 2, 0}, 100 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 1, 0, 0, 0, 2, 0}, 100 * time.Millisecond},
		action{[]int{0, 2, 0, 1, 0, 0, 0, 0, 2, 0}, 100 * time.Millisecond},
		action{[]int{0, 2, 1, 0, 0, 0, 0, 0, 2, 0}, 100 * time.Millisecond},

		action{[]int{0, 2, 0, 1, 0, 0, 0, 0, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 1, 0, 0, 0, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 1, 0, 0, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 1, 0, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 0, 1, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 1, 0, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 1, 0, 0, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 1, 0, 0, 0, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 2, 0, 1, 0, 0, 0, 0, 2, 0}, 80 * time.Millisecond},
		action{[]int{0, 2, 1, 0, 0, 0, 0, 0, 2, 0}, 80 * time.Millisecond},

		action{[]int{0, 2, 0, 1, 0, 0, 0, 0, 2, 0}, 40 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 1, 0, 0, 0, 2, 0}, 40 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 1, 0, 0, 2, 0}, 40 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 1, 0, 2, 0}, 40 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 0, 1, 2, 0}, 40 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 0, 1, 0, 2, 0}, 40 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 0, 1, 0, 0, 2, 0}, 40 * time.Millisecond},
		action{[]int{0, 2, 0, 0, 1, 0, 0, 0, 2, 0}, 40 * time.Millisecond},
		action{[]int{0, 2, 0, 1, 0, 0, 0, 0, 2, 0}, 40 * time.Millisecond},
		action{[]int{0, 2, 1, 0, 0, 0, 0, 0, 2, 0}, 40 * time.Millisecond},

		action{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 80 * time.Millisecond},
	}
	patterns = [][]action{
		pattern0,
		pattern1,
		pattern2,
		pattern3,
	}
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "OK")
}

func generateRandomPatternNum(max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(max)
}
func mapNumberToLEDState(i int) (string, int) {
	switch i {
	case 1:
		{
			return "green", 1
		}
	case 2:
		{
			return "red", 1
		}
	default:
		{
			return "green", 0
		}
	}
}

func main() {

	cable.SetInterval(func() bool {
		data := patterns[generateRandomPatternNum(len(patterns))]
		len := len(data) - 1

		for i := 0; i <= len; i++ {
			entry := data[i]

			sequence := entry.Sequence

			for n, numericalState := range sequence {
				color, state := mapNumberToLEDState(numericalState)
				if i == 0 || data[i-1].Sequence[n] != numericalState {
					resp, err := http.Get(fmt.Sprintf("http://minion-%d:3333/%s/%d", n, color, state))
					resp.Body.Close()
					resp.Close = true
					if err != nil {
						fmt.Println(err)
					}
				}
			}
			time.Sleep(entry.Duration)

		}

		return true
	}, 1*time.Minute)

	blockForever()
}

func blockForever() {
	select {}
}
