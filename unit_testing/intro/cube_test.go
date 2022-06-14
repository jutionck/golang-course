package intro

import "testing"

func TestCubeVolume_Success(t *testing.T) {
	cubeMock := Cube{Side: 4}
	expected := float64(64)
	actual, err := cubeMock.Volume()
	if expected != actual {
		t.Errorf("The result should be %2.f", actual)
	}

	if err != nil {
		t.Error("Volume result error")
	}
}

func TestCubeVolume_Failed(t *testing.T) {
	cubeMock := Cube{Side: -4}
	actual, err := cubeMock.Volume()

	if err == nil {
		t.Error("Numbers can't be negative")
	}

	if actual < 0 {
		t.Error("Numbers can't be negative")
	}

}

func TestCubeArea(t *testing.T) {
	t.Run("When the input side is corect", func(t *testing.T) {
		cubeMock := Cube{Side: 4}
		expected := float64(96)
		actual, err := cubeMock.Area()
		if expected != actual {
			t.Errorf("The result should be %2.f", actual)
		}

		if err != nil {
			t.Error("Area result error")
		}
	})

	t.Run("When the input side is negative number", func(t *testing.T) {
		cubeMock := Cube{Side: -4}
		actual, err := cubeMock.Area()
		if err == nil {
			t.Error("Numbers can't be negative")
		}
		if actual < 0 {
			t.Error("Numbers can't be negative")
		}
	})

}
