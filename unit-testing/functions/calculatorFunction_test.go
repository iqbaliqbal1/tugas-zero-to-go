package functions



import (
	"testing"

	"github.com/stretchr/testify/assert"

)
// test addition this is for testing addition
func TestAddition (t *testing.T)  {
	// if A + B = C 
	// expect 1 + 2 equals 3
	testResult := Addition(1, 2)

	assert.Equal(t, 3, testResult)
}

// test multiple this is for testing addition
func TestMultiple(t *testing.T){
	// if A * B = C
	// expect 1 * 2 equals 2
	testResult := Multiple(1, 2)

	assert.Equal(t, 2, testResult)
} 

func TestSubtraction (t *testing.T) {
	// if A - B = C
	// expect 5 - 2 equals 3
	testResult := Subtraction(5, 2)

	assert.Equal(t, 3, testResult)
}

func TestDivision (t *testing.T) {
	// if A / B = C
	// expect 5.0 / 2.0 equals 2.5
	testResult := Division(5.0, 2.0)

	assert.Equal(t, 2.5, testResult)
} 