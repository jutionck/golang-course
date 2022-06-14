package intro

import "testing"

func TestNewCubeService(t *testing.T) {
	t.Run("When the cube service return not nil", func(t *testing.T) {
		cubeMock := Cube{Side: 4}
		repo := NewCubeRepostory(cubeMock)
		useCase := NewCubeService(repo)
		if useCase.repo != repo {
			t.Error("Cube service error")
		}
	})
}

func TestGetCubeVolumeService(t *testing.T) {
	t.Run("When get cube volume with service correct", func(t *testing.T) {
		cubeMock := Cube{Side: 4}
		repo := NewCubeRepostory(cubeMock)
		useCase := NewCubeService(repo)
		expected := float64(64)
		actual, err := useCase.GetCubeVolume()
		if expected != actual {
			t.Errorf("The volume result should be %2.f", actual)
		}
		if err != nil {
			t.Error("Area result error")
		}
	})

	t.Run("When get cube volume with service in-correct", func(t *testing.T) {
		cubeMock := Cube{Side: -4}
		repo := NewCubeRepostory(cubeMock)
		useCase := NewCubeService(repo)
		actual, err := useCase.GetCubeVolume()
		if err == nil {
			t.Error("Numbers can't be negative")
		}
		if actual < 0 {
			t.Error("Numbers can't be negative")
		}
	})
}

func TestGetCubeAreService(t *testing.T) {
	t.Run("When get cube area with service correct", func(t *testing.T) {
		cubeMock := Cube{Side: 4}
		repo := NewCubeRepostory(cubeMock)
		useCase := NewCubeService(repo)
		expected := float64(96)
		actual, err := useCase.GetCubeArea()
		if expected != actual {
			t.Errorf("The Area result should be %2.f", actual)
		}
		if err != nil {
			t.Error("Area result error")
		}
	})

	t.Run("When get cube area with service in-correct", func(t *testing.T) {
		cubeMock := Cube{Side: -4}
		repo := NewCubeRepostory(cubeMock)
		useCase := NewCubeService(repo)
		actual, err := useCase.GetCubeArea()
		if err == nil {
			t.Error("Numbers can't be negative")
		}
		if actual < 0 {
			t.Error("Numbers can't be negative")
		}
	})
}
