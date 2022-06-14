package intro

import "testing"

func TestNewCubeRepository(t *testing.T) {
	t.Run("When the cube repository return not nil", func(t *testing.T) {
		cubeMock := Cube{Side: 4}
		repo := NewCubeRepostory(cubeMock)
		if repo == nil {
			t.Error("Cube repository error")
		}
	})
}

func TestCubeRepositoryVolume(t *testing.T) {
	t.Run("When the input side is corect", func(t *testing.T) {
		cubeMock := Cube{Side: 4}
		repo := NewCubeRepostory(cubeMock)
		expected := float64(64)
		actual, err := repo.Volume()
		if expected != actual {
			t.Errorf("The volume result should be %2.f", actual)
		}
		if err != nil {
			t.Error("Area result error")
		}
	})

	t.Run("When the input side is negative number", func(t *testing.T) {
		cubeMock := Cube{Side: -4}
		repo := NewCubeRepostory(cubeMock)
		actual, err := repo.Volume()
		if err == nil {
			t.Error("Numbers can't be negative")
		}
		if actual < 0 {
			t.Error("Numbers can't be negative")
		}
	})
}

func TestCubeRepositoryArea(t *testing.T) {
	t.Run("When the input side is corect", func(t *testing.T) {
		cubeMock := Cube{Side: 4}
		repo := NewCubeRepostory(cubeMock)
		expected := float64(96)
		actual, err := repo.Area()
		if expected != actual {
			t.Errorf("The Area result should be %2.f", actual)
		}
		if err != nil {
			t.Error("Area result error")
		}
	})

	t.Run("When the input side is negative number", func(t *testing.T) {
		cubeMock := Cube{Side: -4}
		repo := NewCubeRepostory(cubeMock)
		actual, err := repo.Area()
		if err == nil {
			t.Error("Numbers can't be negative")
		}
		if actual < 0 {
			t.Error("Numbers can't be negative")
		}
	})
}
