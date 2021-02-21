package alg

import (
	"testing"
)

func TestFibArr(t *testing.T) {
	n := 5

	arr := FibArr(n)

	t.Log(arr)
}
