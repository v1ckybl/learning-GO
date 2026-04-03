package main

import (
	"sync/atomic"
)

type GenerarID struct {
	contadorID uint64
}

func (g *GenerarID) generarID() uint64 {
	return atomic.AddUint64(&g.contadorID, 1)
}

var idGeneraaado = &GenerarID{}
