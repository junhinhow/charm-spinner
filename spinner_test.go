package spinner

import (
	"testing"
)

func TestAllSpinnersHaveFrames(t *testing.T) {
	for name, s := range All() {
		if len(s.Frames) == 0 {
			t.Errorf("spinner %s nao possui frames", name)
		}
		if s.Interval <= 0 {
			t.Errorf("spinner %s possui intervalo invalido: %v", name, s.Interval)
		}
	}
}

func TestRenderCycles(t *testing.T) {
	s := Dots
	total := len(s.Frames)
	for i := 0; i < total*3; i++ {
		result := s.Render(i)
		if result == "" {
			t.Errorf("Render(%d) retornou string vazia", i)
		}
	}
}

func TestAllCount(t *testing.T) {
	all := All()
	if len(all) < 30 {
		t.Errorf("esperado pelo menos 30 spinners, encontrado %d", len(all))
	}
}
