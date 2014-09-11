package binoculars

import (
	"fmt"
	"testing"
)

type m struct {
	A string
	B string
}

type mAccessor struct {
	Property string
}

func (p mAccessor) Get(x Any) Any {
	y := x.(m)
	switch p.Property {
	case "A":
		return y.A
	default:
		return y.B
	}
}

func (p mAccessor) Set(x Any, y Any) Any {
	z := x.(m)
	switch p.Property {
	case "A":
		return m{A: y.(string), B: z.B}
	default:
		return m{A: z.A, B: y.(string)}
	}
}

func defaults() m {
	return m{
		A: "A",
		B: "B",
	}
}

func Benchmark_Testing_AccessorLens(b *testing.B) {
	x := defaults()

	var r m
	for i := 0; i < b.N; i++ {
		accessor := mAccessor{Property: "A"}
		l := AccessorLens(accessor).Run(x)
		r = l.Set(fmt.Sprintf("%s%v", l.Get(), i)).(m)
	}

	b.Log(r)
}

func Benchmark_Testing_ObjectLens(b *testing.B) {
	x := defaults()

	var r m
	for i := 0; i < b.N; i++ {
		l := ObjectLens("A").Run(x)
		r = l.Set(fmt.Sprintf("%s%v", l.Get(), i)).(m)
	}

	b.Log(r)
}
