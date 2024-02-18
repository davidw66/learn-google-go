package dog

import (
	"fmt"
	"log"
	"os"
	"testing"
)

// TestMain will only show output when a test fails
// In this example a file is created and then deleted at the end of the tests
func TestMain(m *testing.M) {
	log.Println("Stuff before test")
	os.Create("myTest.txt")

	// This runs all the tests
	exitVal := m.Run()

	log.Println("Stuff after test")
	os.Remove("myTest.txt")

	// Always have this line
	os.Exit(exitVal)
}

func TestYears(t *testing.T) {
	got := Years(7)
	want := 49

	if got != want {
		t.Errorf("\n Got: %v\nWant: %v\n", got, want)
	}
}

func ExampleYears() {
	fmt.Println(Years(7))
	// Output:
	// 49
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func TestYearsTwo(t *testing.T) {
	got := YearsTwo(7)
	want := 49

	if got != want {
		t.Errorf("\n Got: %v\nWant: %v\n", got, want)
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(7))
	// Output:
	// 49
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}
