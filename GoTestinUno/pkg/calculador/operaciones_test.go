package calculador

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	a, b := 10, 5
	expec := 15
	resultado := Add(a, b)
	assert.Equal(t, expec, resultado)
	if resultado != expec {
		t.Errorf("el resultado %d,el numero del resultado es distinto del esperado %d", resultado, expec)
	}
}
func TestSubStrac(t *testing.T) {
	a, b := 10, 5
	expec := 5
	resultado := Substrab(a, b)
	assert.Equal(t, expec, resultado)
	if resultado != expec {
		t.Errorf("el resultado %d,el numero del resultado es distinto del esperado %d", resultado, expec)
	}
}
func TestSort(t *testing.T) {
	list := []int{12, 3, 4, 3, 2, 4, 6}
	listSort := []int{2, 3, 3, 4, 4, 6, 12}
	result := Sort(list)
	assert.Equal(t, listSort, result)
}

func TestSplit(t *testing.T) {
	a, b := 3, 2
	expec := float64(1)
	result, err := Split(a, b)
	assert.Nil(t, err)
	assert.Equal(t, expec, result)

}

func TestSplitFail(t *testing.T) {
	a, b := 0, 2
	errorWait := fmt.Sprintf("ninguno de los numeros puede ser 0 : a %d, %d", a, b)
	_, err := Split(a, b)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, errorWait)

}
