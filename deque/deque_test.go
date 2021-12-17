package deque

import (
	"testing"
)

func TestDeque_AddLast(t *testing.T) {
	d := New[int]()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)
	d.AddLast(4)
	i := 1
	for !d.Empty() {
		first := d.RemoveFirst()
		if i != first {
			t.Errorf("AddLast() = %v, want %v", first, i)
		}
		i++
	}
}
