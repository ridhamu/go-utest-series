package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// table test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		Name     string
		Request  string
		Expected string
	}{
		{
			Name:     "Muhammad",
			Request:  "muhammad",
			Expected: "Hello muhammad",
		},
		{
			Name:     "Ridha",
			Request:  "ridha",
			Expected: "Hello ridha",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := HelloWorld(test.Request)
			require.Equal(t, test.Expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Muhammad", func(t *testing.T) {
		result := HelloWorld("Muhammad")
		require.Equal(t, "Hello Muhammad", result, "result must be 'Hello Muhammad'")
	})
	t.Run("Ridha", func(t *testing.T) {
		result := HelloWorld("Ridha")
		require.Equal(t, "Hello Ridha", result, "result must be 'Hello Ridha'")
	})
}

func TestMain(m *testing.M) {
	// before running test
	fmt.Println("before running test")
	m.Run()
	fmt.Println("after running test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Test cannot run on MacOS")
	}
	result := HelloWorld("ridha")
	assert.Equal(t, "Hello ridha", result, "result should be 'Hello ridha'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("ridha")
	assert.Equal(t, "Hello ridha", result, "result should be 'Hello ridha'")
	fmt.Println("TestHelloWorldAssert done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("ridha")
	require.Equal(t, "Hello ridha", result, "result should be 'Hello ridha'")
	fmt.Println("TestHelloWorldAssert done") // should not be printed
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("muhammad")

	if result != "Hello muhammad" {
		// panic("result not 'Hello muhammad'")
		// t.Fail() // marked failed but still execute code programs
		t.Error("unmatch result 'Hello Muhammad'", "\tinstead get ", result)
	}

	fmt.Println("===TestHelloWorld done===") // this will be printed
}

func TestHelloWorldRidha(t *testing.T) {
	result := HelloWorld("ridha")

	if result != "Hello ridha" {
		// panic("result not 'Hello ridha'")
		// t.FailNow() // marked failed and finished
		t.Fatal("unmatch result")
	}
	fmt.Println("===TestHelloWorldRidha done===") // this will not be executed
}
