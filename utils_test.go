package pterm_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/gookit/color"
	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

var printables = []interface{}{"Hello, World!", 1337, true, false, -1337, 'c', 1.5, "\\", "%s"}

// testPrintContains can be used to test Print methods.
func testPrintContains(t *testing.T, logic func(w io.Writer, a interface{})) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s := captureStdout(func(w io.Writer) {
				logic(w, printable)
			})
			assert.Contains(t, s, fmt.Sprint(printable))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s := captureStdout(func(w io.Writer) {
				logic(w, printable)
			})
			assert.Contains(t, s, fmt.Sprint(printable))
		})
		pterm.EnableStyling()
	}
}

// testPrintfContains can be used to test Printf methods.
func testPrintfContains(t *testing.T, logic func(w io.Writer, format string, a interface{})) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s := captureStdout(func(w io.Writer) {
				logic(w, "Hello, %v!", printable)
			})
			assert.Contains(t, s, fmt.Sprintf("Hello, %v!", fmt.Sprint(printable)))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s := captureStdout(func(w io.Writer) {
				logic(w, "Hello, %v!", printable)
			})
			assert.Contains(t, s, fmt.Sprintf("Hello, %v!", fmt.Sprint(printable)))
		})
		pterm.EnableStyling()
	}
}

// testPrintflnContains can be used to test Printfln methods.
func testPrintflnContains(t *testing.T, logic func(w io.Writer, format string, a interface{})) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testPrintfContains(t, logic)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testPrintfContains(t, logic)
		})
		pterm.EnableStyling()
	}
}

// testPrintlnContains can be used to test Println methods.
func testPrintlnContains(t *testing.T, logic func(w io.Writer, a interface{})) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testPrintContains(t, logic)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testPrintContains(t, logic)
		})
		pterm.EnableStyling()
	}
}

// testSprintContains can be used to test Sprint methods.
func testSprintContains(t *testing.T, logic func(a interface{}) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			assert.Contains(t, logic(printable), fmt.Sprint(printable))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			assert.Contains(t, logic(printable), fmt.Sprint(printable))
		})
		pterm.EnableStyling()
	}
}

// testSprintContainsWithoutError can be used to test Sprint methods which return an error.
func testSprintContainsWithoutError(t *testing.T, logic func(a interface{}) (string, error)) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s, err := logic(printable)
			assert.Contains(t, s, fmt.Sprint(printable))
			assert.NoError(t, err)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s, err := logic(printable)
			assert.Contains(t, s, fmt.Sprint(printable))
			assert.NoError(t, err)
		})
		pterm.EnableStyling()
	}
}

// testSprintfContains can be used to test Sprintf methods.
func testSprintfContains(t *testing.T, logic func(format string, a interface{}) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			assert.Contains(t, logic("Hello, %v!", printable), fmt.Sprintf("Hello, %v!", printable))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			assert.Contains(t, logic("Hello, %v!", printable), fmt.Sprintf("Hello, %v!", printable))
		})
		pterm.EnableStyling()
	}
}

// testSprintflnContains can be used to test Sprintfln methods.
func testSprintflnContains(t *testing.T, logic func(format string, a interface{}) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testSprintfContains(t, logic)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testSprintfContains(t, logic)
		})
		pterm.EnableStyling()
	}
}

// testSprintlnContains can be used to test Sprintln methods.
func testSprintlnContains(t *testing.T, logic func(a interface{}) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testSprintContains(t, logic)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testSprintContains(t, logic)
		})
		pterm.EnableStyling()
	}
}

// testDoesOutput can be used to test if something is outputted to stdout.
func testDoesOutput(t *testing.T, logic func(w io.Writer)) {
	assert.NotEmpty(t, captureStdout(logic))
	pterm.DisableStyling()
	assert.NotEmpty(t, captureStdout(logic))
	pterm.EnableStyling()
}

// testEmpty checks that a function does not return a string.
func testEmpty(t *testing.T, logic func(a interface{}) string) {
	for _, printable := range printables {
		assert.Empty(t, logic(printable))
		pterm.DisableStyling()
		assert.Empty(t, logic(printable))
		pterm.EnableStyling()
	}
}

// testDoesNotOutput can be used, to test that something does not output anything to stdout.
func testDoesNotOutput(t *testing.T, logic func(w io.Writer)) {
	assert.Empty(t, captureStdout(logic))
	pterm.DisableStyling()
	assert.Empty(t, captureStdout(logic))
	pterm.EnableStyling()
}

// captureStdout captures everything written to the terminal and returns it as a string.
func captureStdout(f func(w io.Writer)) string {
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	color.SetOutput(w)

	f(w)

	_ = w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = originalStdout
	color.SetOutput(w)
	_ = r.Close()

	return string(out)
}

func proxyToDevNull() {
	pterm.SetDefaultOutput(os.NewFile(0, os.DevNull))
}