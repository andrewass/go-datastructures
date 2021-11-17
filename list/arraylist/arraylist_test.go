package arraylist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayListOperations(t *testing.T) {
	arrayList := New()

	assert.Equal(t, 0, arrayList.Size())

	numbers := []int{67, 45, 2, 35, 62, 8, 15, 22, 34, 45}
	for _, number := range numbers {
		arrayList.Add(number)
	}

	assert.Equal(t, 10, arrayList.Size())
	assert.Equal(t, 67, arrayList.Get(0))
	assert.Equal(t, 45, arrayList.Get(9))

	arrayList.Remove(6)
	arrayList.Remove(0)
	arrayList.Remove(7)

	assert.Equal(t, 7, arrayList.Size())



	for arrayList.Size() > 0 {
		arrayList.Remove(0)
	}
}
