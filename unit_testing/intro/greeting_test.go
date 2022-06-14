package intro

import "testing"

func TestGreeting(t *testing.T) {
	expected := "Selamat Pagi"
	actual := Greeting()
	if expected != actual {
		t.Error("Seharusnya mencetak 'Selamat Pagi'")
	}
}

func TestSayHi(t *testing.T) {
	expected := "Hai salam kenal"
	actual := SayHi()
	if expected != actual {
		t.Error("Seharusnya mencetak 'Hai salam kenal'")
	}
}

func BenchmarkGreetingTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greeting()
	}
}
