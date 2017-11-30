package tests

import (
	"fmt"
	fnTest "github.com/fnproject/fn/test"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	s := fnTest.SetupDefaultSuite()
	result := m.Run()
	fnTest.Cleanup()
	s.Cancel()
	if result == 0 {
		fmt.Fprintln(os.Stdout, "😀  👍  🎗")
	}
	os.Exit(result)
}
