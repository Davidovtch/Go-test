package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 2)
	want := "aa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	/*b.N é o número de vezes que o benchmark é realizado
	ele é decidido pela aplicação*/
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
