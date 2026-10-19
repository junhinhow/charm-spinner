package spinner

import (
	"sort"
	"strings"
	"sync"
	"testing"

	"charm.land/lipgloss/v2"
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
	if len(all) < 40 {
		t.Errorf("esperado pelo menos 40 spinners, encontrado %d", len(all))
	}
}

// TestRenderCicloConfere que o frame volta ao início depois de um ciclo
// completo de frames.
func TestRenderCiclo(t *testing.T) {
	total := len(Dots.Frames)
	if got, want := Dots.Render(total), Dots.Render(0); got != want {
		t.Errorf("Render(%d) = %q, esperado %q (frame 0)", total, got, want)
	}
	if got, want := Dots.Render(total+3), Dots.Render(3); got != want {
		t.Errorf("Render(%d) = %q, esperado %q (frame 3)", total+3, got, want)
	}
}

// TestRenderFrameNegativo verifica que frames negativos não entram em panic e
// contam a partir do fim do ciclo.
func TestRenderFrameNegativo(t *testing.T) {
	total := len(Dots.Frames)
	if got, want := Dots.Render(-1), Dots.Frames[total-1]; got != want {
		t.Errorf("Render(-1) = %q, esperado %q", got, want)
	}
	if got, want := Dots.Render(-total), Dots.Frames[0]; got != want {
		t.Errorf("Render(-%d) = %q, esperado %q", total, got, want)
	}
	// Frame muito negativo também precisa ser seguro.
	if Dots.Render(-1000) == "" {
		t.Error("Render(-1000) retornou string vazia")
	}
}

// TestRenderSemFrames verifica que um spinner sem frames não entra em panic.
func TestRenderSemFrames(t *testing.T) {
	vazio := Spinner{Interval: 100}
	if got := vazio.Render(0); got != "" {
		t.Errorf("Render de um spinner sem frames = %q, esperado %q", got, "")
	}
	if got := vazio.Render(-5); got != "" {
		t.Errorf("Render de um spinner sem frames = %q, esperado %q", got, "")
	}
}

// TestRenderIndicesExtremos verifica que nenhum índice causa panic.
func TestRenderIndicesExtremos(t *testing.T) {
	for _, frame := range []int{-1000000, -1, 0, 1, 1000000} {
		for nome, s := range All() {
			if got := s.Render(frame); got == "" {
				t.Errorf("spinner %s: Render(%d) retornou string vazia", nome, frame)
			}
		}
	}
}

// TestAllNomesUnicos verifica que não há nomes repetidos no mapa.
func TestAllNomesUnicos(t *testing.T) {
	// Cada spinner predefinido precisa estar presente no mapa exatamente
	// uma vez, com o mesmo conteúdo.
	predefinidos := map[string]Spinner{
		"Dots": Dots, "Line": Line, "Star": Star, "Moon": Moon, "Earth": Earth,
		"Clock": Clock, "Bounce": Bounce, "Bar": Bar, "Circle": Circle,
		"Square": Square, "Triangle": Triangle, "Arrow": Arrow, "Pulse": Pulse,
		"Flip": Flip, "Grow": Grow, "Layer": Layer, "Toggle": Toggle,
		"Balloon": Balloon, "Noise": Noise, "Fading": Fading,
		"Christmas": Christmas, "Monkey": Monkey, "Hearts": Hearts,
		"Weather": Weather, "Soccer": Soccer, "Matrix": Matrix, "DNA": DNA,
		"Orbit": Orbit, "Rocket": Rocket, "Fire": Fire, "Ellipsis": Ellipsis,
		"Rectangle": Rectangle, "Diamond": Diamond, "Hexagon": Hexagon,
		"BouncingBar": BouncingBar, "Digits": Digits, "Quotes": Quotes,
		"Cat": Cat, "Money": Money, "Beach": Beach, "Music": Music,
		"Target": Target,
	}

	all := All()
	if len(all) != len(predefinidos) {
		t.Errorf("All() tem %d entradas, esperado %d", len(all), len(predefinidos))
	}
	for nome, esperado := range predefinidos {
		obtido, ok := all[nome]
		if !ok {
			t.Errorf("All() não contém o spinner %q", nome)
			continue
		}
		if len(obtido.Frames) != len(esperado.Frames) {
			t.Errorf("spinner %q: %d frames, esperado %d", nome, len(obtido.Frames), len(esperado.Frames))
		}
		for i := range esperado.Frames {
			if obtido.Frames[i] != esperado.Frames[i] {
				t.Errorf("spinner %q: frame %d = %q, esperado %q", nome, i, obtido.Frames[i], esperado.Frames[i])
			}
		}
	}
}

// TestAllDeterministico verifica que All() devolve sempre o mesmo conteúdo.
func TestAllDeterministico(t *testing.T) {
	primeiro := All()
	segundo := All()
	if len(primeiro) != len(segundo) {
		t.Fatalf("All() devolveu tamanhos diferentes: %d e %d", len(primeiro), len(segundo))
	}
	for nome, s := range primeiro {
		outro, ok := segundo[nome]
		if !ok {
			t.Fatalf("All() não é determinístico: %q ausente na segunda chamada", nome)
		}
		if len(s.Frames) != len(outro.Frames) {
			t.Fatalf("All() não é determinístico no spinner %q", nome)
		}
		for i := range s.Frames {
			if s.Frames[i] != outro.Frames[i] {
				t.Fatalf("All() não é determinístico no spinner %q, frame %d", nome, i)
			}
		}
	}
}

// TestCloneIndependente verifica que Clone devolve uma cópia modificável sem
// afetar o mapa original.
func TestCloneIndependente(t *testing.T) {
	copia := Clone()
	if len(copia) != len(All()) {
		t.Fatalf("Clone() tem %d entradas, esperado %d", len(copia), len(All()))
	}
	copia["Dots"] = Spinner{Frames: []string{"X"}}
	if All()["Dots"].Frames[0] != Dots.Frames[0] {
		t.Error("Clone() não devolveu uma cópia independente")
	}
}

// TestNamesOrdenados verifica que Names devolve os nomes em ordem alfabética.
func TestNamesOrdenados(t *testing.T) {
	nomes := Names()
	if len(nomes) != len(All()) {
		t.Fatalf("Names() tem %d entradas, esperado %d", len(nomes), len(All()))
	}
	if !sort.StringsAreSorted(nomes) {
		t.Errorf("Names() não está em ordem alfabética: %v", nomes)
	}
}

// TestStyleCustomizado verifica que o estilo definido no spinner é aplicado.
func TestStyleCustomizado(t *testing.T) {
	custom := Spinner{
		Frames:   []string{"a"},
		Interval: 100,
		Style:    lipgloss.NewStyle().Bold(true),
	}
	got := custom.Render(0)
	if !strings.Contains(got, "\x1b[1m") {
		t.Errorf("o estilo próprio do spinner não foi aplicado: %q", got)
	}
}

// TestStylePadrao verifica que o estilo padrão do pacote é aplicado quando o
// spinner não define um estilo próprio.
func TestStylePadrao(t *testing.T) {
	original := DefaultStyle()
	defer SetDefaultStyle(original)

	SetDefaultStyle(lipgloss.NewStyle().Bold(true))
	got := Dots.Render(0)
	if !strings.Contains(got, "\x1b[1m") {
		t.Errorf("o estilo padrão do pacote não foi aplicado: %q", got)
	}

	// Um spinner com estilo próprio ignora o estilo padrão.
	own := Spinner{Frames: []string{"a"}, Style: lipgloss.NewStyle().Italic(true)}
	if got := own.Render(0); !strings.Contains(got, "\x1b[3m") {
		t.Errorf("o estilo próprio deveria prevalecer sobre o padrão: %q", got)
	}
}

// TestRenderNaoCorreDataRace verifica que a renderização não lê o global
// mutável Style. Execute com -race para que a corrida seja detectada.
func TestRenderNaoCorreDataRace(t *testing.T) {
	defer func() { Style = lipgloss.NewStyle() }()

	var wg sync.WaitGroup
	stop := make(chan struct{})

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
					_ = Dots.Render(1)
				}
			}
		}()
	}

	for i := 0; i < 500; i++ {
		Style = lipgloss.NewStyle().Bold(i%2 == 0)
	}

	close(stop)
	wg.Wait()
}

// TestSetDefaultStyleConcorrente verifica que trocar o estilo padrão durante
// a renderização é seguro. Execute com -race.
func TestSetDefaultStyleConcorrente(t *testing.T) {
	original := DefaultStyle()
	defer SetDefaultStyle(original)

	var wg sync.WaitGroup
	stop := make(chan struct{})

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
					_ = Dots.Render(1)
				}
			}
		}()
	}

	for i := 0; i < 500; i++ {
		SetDefaultStyle(lipgloss.NewStyle().Bold(i%2 == 0))
	}

	close(stop)
	wg.Wait()
}
