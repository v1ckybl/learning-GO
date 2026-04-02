package main

import (
	"sync/atomic"
)

var contadorID uint64

func generarID() uint64 {
	return atomic.AddUint64(&contadorID, 1)
}
