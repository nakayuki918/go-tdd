package iteration

import "testing"

func TestRepeat(t *testing.T) {
	s := []string{"Golang", "Java", "Python"}
	repeated := Repeat(s)
	expected := "GolangJavaPython"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat([]string{"Golang", "Java", "Python"})
	}
}
