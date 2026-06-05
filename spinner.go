// Package spinner fornece uma colecao de 30+ animacoes de spinner para terminais.
package spinner

import (
	"time"

	"charm.land/lipgloss/v2"
)

// Spinner define uma animacao de spinner com frames e intervalo.
type Spinner struct {
	Frames   []string
	Interval time.Duration
}

// Style armazena o estilo lipgloss aplicado ao spinner.
var Style = lipgloss.NewStyle()

// Render retorna o frame na posicao indicada (com ciclo automatico).
func (s Spinner) Render(frame int) string {
	f := s.Frames[frame%len(s.Frames)]
	return Style.Render(f)
}

// Spinners predefinidos — cada um com frames Unicode distintos.
var (
	Dots = Spinner{
		Frames:   []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		Interval: 80 * time.Millisecond,
	}
	Line = Spinner{
		Frames:   []string{"-", "\\", "|", "/"},
		Interval: 130 * time.Millisecond,
	}
	Star = Spinner{
		Frames:   []string{"✶", "✸", "✹", "✺", "✹", "✷"},
		Interval: 70 * time.Millisecond,
	}
	Moon = Spinner{
		Frames:   []string{"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘"},
		Interval: 80 * time.Millisecond,
	}
	Earth = Spinner{
		Frames:   []string{"🌍", "🌎", "🌏"},
		Interval: 180 * time.Millisecond,
	}
	Clock = Spinner{
		Frames:   []string{"🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛"},
		Interval: 100 * time.Millisecond,
	}
	Bounce = Spinner{
		Frames:   []string{"⠁", "⠂", "⠄", "⠂"},
		Interval: 120 * time.Millisecond,
	}
	Bar = Spinner{
		Frames:   []string{"▏", "▎", "▍", "▌", "▋", "▊", "▉", "█", "▉", "▊", "▋", "▌", "▍", "▎", "▏"},
		Interval: 80 * time.Millisecond,
	}
	Circle = Spinner{
		Frames:   []string{"◐", "◓", "◑", "◒"},
		Interval: 120 * time.Millisecond,
	}
	Square = Spinner{
		Frames:   []string{"▖", "▘", "▝", "▗"},
		Interval: 120 * time.Millisecond,
	}
	Triangle = Spinner{
		Frames:   []string{"◢", "◣", "◤", "◥"},
		Interval: 120 * time.Millisecond,
	}
	Arrow = Spinner{
		Frames:   []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
		Interval: 100 * time.Millisecond,
	}
	Pulse = Spinner{
		Frames:   []string{"█", "▓", "▒", "░", "▒", "▓"},
		Interval: 120 * time.Millisecond,
	}
	Flip = Spinner{
		Frames:   []string{"_", "_", "_", "-", "`", "`", "'", "´", "-", "_", "_", "_"},
		Interval: 70 * time.Millisecond,
	}
	Grow = Spinner{
		Frames:   []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▂", "▁"},
		Interval: 80 * time.Millisecond,
	}
	Layer = Spinner{
		Frames:   []string{"─", "━", "═", "━"},
		Interval: 150 * time.Millisecond,
	}
	Toggle = Spinner{
		Frames:   []string{"⊶", "⊷"},
		Interval: 250 * time.Millisecond,
	}
	Balloon = Spinner{
		Frames:   []string{".", "o", "O", "°", "O", "o", "."},
		Interval: 140 * time.Millisecond,
	}
	Noise = Spinner{
		Frames:   []string{"▓", "▒", "░", "▒"},
		Interval: 100 * time.Millisecond,
	}
	Fading = Spinner{
		Frames:   []string{"░", "▒", "▓", "█", "▓", "▒", "░", " "},
		Interval: 120 * time.Millisecond,
	}
	Christmas = Spinner{
		Frames:   []string{"🎄", "🎅", "⛄", "🎁", "🔔", "⭐"},
		Interval: 200 * time.Millisecond,
	}
	Monkey = Spinner{
		Frames:   []string{"🙈", "🙉", "🙊"},
		Interval: 210 * time.Millisecond,
	}
	Hearts = Spinner{
		Frames:   []string{"💛", "💙", "💜", "💚", "❤️", "🧡"},
		Interval: 120 * time.Millisecond,
	}
	Weather = Spinner{
		Frames:   []string{"☀️", "🌤️", "⛅", "🌥️", "☁️", "🌧️", "⛈️", "🌩️", "🌧️", "☁️", "🌥️", "⛅", "🌤️"},
		Interval: 150 * time.Millisecond,
	}
	Soccer = Spinner{
		Frames:   []string{"⚽", "🥅", "👟", "💨"},
		Interval: 200 * time.Millisecond,
	}
	Matrix = Spinner{
		Frames:   []string{"ᚠ", "ᚡ", "ᚢ", "ᚣ", "ᚤ", "ᚥ", "ᚦ", "ᚧ", "ᚨ", "ᚩ"},
		Interval: 70 * time.Millisecond,
	}
	DNA = Spinner{
		Frames:   []string{"╔═╗", "╚═╝", "╔═╗", "═══", "╚═╝", "═══"},
		Interval: 160 * time.Millisecond,
	}
	Orbit = Spinner{
		Frames:   []string{"◜ ", " ◝", " ◞", "◟ "},
		Interval: 100 * time.Millisecond,
	}
	Rocket = Spinner{
		Frames:   []string{"🚀", "💨", "✨", "🌟", "✨", "💨"},
		Interval: 120 * time.Millisecond,
	}
	Fire = Spinner{
		Frames:   []string{"🔥", "🔥", "💥", "✨", "🔥", "💥"},
		Interval: 100 * time.Millisecond,
	}
)

// All retorna todos os spinners predefinidos em um mapa nome->spinner.
func All() map[string]Spinner {
	return map[string]Spinner{
		"Dots":      Dots,
		"Line":      Line,
		"Star":      Star,
		"Moon":      Moon,
		"Earth":     Earth,
		"Clock":     Clock,
		"Bounce":    Bounce,
		"Bar":       Bar,
		"Circle":    Circle,
		"Square":    Square,
		"Triangle":  Triangle,
		"Arrow":     Arrow,
		"Pulse":     Pulse,
		"Flip":      Flip,
		"Grow":      Grow,
		"Layer":     Layer,
		"Toggle":    Toggle,
		"Balloon":   Balloon,
		"Noise":     Noise,
		"Fading":    Fading,
		"Christmas": Christmas,
		"Monkey":    Monkey,
		"Hearts":    Hearts,
		"Weather":   Weather,
		"Soccer":    Soccer,
		"Matrix":    Matrix,
		"DNA":       DNA,
		"Orbit":     Orbit,
		"Rocket":    Rocket,
		"Fire":      Fire,
	}
}
