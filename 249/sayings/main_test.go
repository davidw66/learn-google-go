package sayings

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	got := fmt.Sprint(Greet("James"))
	want := "Welcome my dear James."

	if got != want {
		t.Errorf("\n Got: %v \nWant: %v", got, want)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome my dear James.
}

// To run all benchmarks go test -bench .
// To run only greet go test -bench Greet

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
