// Package spinner fornece uma coleção de 40+ animações de spinner para
// terminais.
package spinner

import (
	"reflect"
	"sort"
	"sync/atomic"
	"time"

	"charm.land/lipgloss/v2"
)

// Spinner define uma animação de spinner com frames, intervalo e estilo.
type Spinner struct {
	Frames   []string
	Interval time.Duration

	// Style é o estilo lipgloss aplicado aos frames. Quando o estilo não
	// define nenhuma propriedade, vale o estilo padrão do pacote, ajustável
	// com SetDefaultStyle.
	Style lipgloss.Style
}

// Style é o estilo padrão aplicado aos spinners que não definem um estilo
// próprio.
//
// O valor é copiado para o estilo padrão interno durante a inicialização do
// pacote. Para trocar o estilo padrão em tempo de execução use
// SetDefaultStyle: atribuir diretamente a esta variável enquanto outro
// goroutine renderiza é uma corrida de dados.
var Style = lipgloss.NewStyle()

// estiloPadrao guarda o estilo padrão atual de forma segura para uso
// concorrente.
var estiloPadrao atomic.Pointer[lipgloss.Style]

func init() {
	s := Style
	estiloPadrao.Store(&s)
}

// SetDefaultStyle define o estilo padrão dos spinners.
//
// É seguro chamar em tempo de execução, inclusive com renderizações
// concorrentes.
func SetDefaultStyle(s lipgloss.Style) {
	estiloPadrao.Store(&s)
}

// DefaultStyle devolve o estilo padrão atual dos spinners.
func DefaultStyle() lipgloss.Style {
	if p := estiloPadrao.Load(); p != nil {
		return *p
	}
	return Style
}

// Render retorna o frame na posição indicada (com ciclo automático).
//
// Índices fora do intervalo são ajustados para o ciclo: um frame negativo
// conta a partir do fim, e um índice maior que o número de frames volta ao
// começo. Spinners sem frames devolvem uma string vazia em vez de entrar em
// panic.
func (s Spinner) Render(frame int) string {
	frames := s.Frames
	if len(frames) == 0 {
		return ""
	}
	indice := frame % len(frames)
	if indice < 0 {
		indice += len(frames)
	}
	return s.style().Render(frames[indice])
}

// style devolve o estilo efetivo do spinner: o próprio, quando definido, ou
// o estilo padrão do pacote.
func (s Spinner) style() lipgloss.Style {
	if !reflect.ValueOf(&s.Style).Elem().IsZero() {
		return s.Style
	}
	return DefaultStyle()
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
	Ellipsis = Spinner{
		Frames:   []string{".", "..", "..."},
		Interval: 120 * time.Millisecond,
	}
	Rectangle = Spinner{
		Frames:   []string{"▛", "▜", "▙", "▟"},
		Interval: 110 * time.Millisecond,
	}
	Diamond = Spinner{
		Frames:   []string{"◆", "◇"},
		Interval: 180 * time.Millisecond,
	}
	Hexagon = Spinner{
		Frames:   []string{"⬢", "⬡"},
		Interval: 180 * time.Millisecond,
	}
	BouncingBar = Spinner{
		Frames:   []string{"[    ]", "[=   ]", "[==  ]", "[=== ]", "[ ===]", "[  ==]", "[   =]"},
		Interval: 100 * time.Millisecond,
	}
	Digits = Spinner{
		Frames:   []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
		Interval: 90 * time.Millisecond,
	}
	Quotes = Spinner{
		Frames:   []string{"'", "\"", "❞", "❟"},
		Interval: 150 * time.Millisecond,
	}
	Cat = Spinner{
		Frames:   []string{"😺", "😸", "😹", "😻", "😼", "😽"},
		Interval: 130 * time.Millisecond,
	}
	Money = Spinner{
		Frames:   []string{"💰", "💵", "💴", "💶"},
		Interval: 200 * time.Millisecond,
	}
	Beach = Spinner{
		Frames:   []string{"🏖️", "🌴", "🌊", "☀️"},
		Interval: 220 * time.Millisecond,
	}
	Music = Spinner{
		Frames:   []string{"♪", "♫", "♬"},
		Interval: 140 * time.Millisecond,
	}
	Target = Spinner{
		Frames:   []string{"◎", "◉", "⊙", "◉"},
		Interval: 130 * time.Millisecond,
	}
)

// predefinidos indexa todos os spinners predefinidos. O mapa é construído
// uma única vez, na inicialização do pacote.
var predefinidos = map[string]Spinner{
	"Dots":        Dots,
	"Line":        Line,
	"Star":        Star,
	"Moon":        Moon,
	"Earth":       Earth,
	"Clock":       Clock,
	"Bounce":      Bounce,
	"Bar":         Bar,
	"Circle":      Circle,
	"Square":      Square,
	"Triangle":    Triangle,
	"Arrow":       Arrow,
	"Pulse":       Pulse,
	"Flip":        Flip,
	"Grow":        Grow,
	"Layer":       Layer,
	"Toggle":      Toggle,
	"Balloon":     Balloon,
	"Noise":       Noise,
	"Fading":      Fading,
	"Christmas":   Christmas,
	"Monkey":      Monkey,
	"Hearts":      Hearts,
	"Weather":     Weather,
	"Soccer":      Soccer,
	"Matrix":      Matrix,
	"DNA":         DNA,
	"Orbit":       Orbit,
	"Rocket":      Rocket,
	"Fire":        Fire,
	"Ellipsis":    Ellipsis,
	"Rectangle":   Rectangle,
	"Diamond":     Diamond,
	"Hexagon":     Hexagon,
	"BouncingBar": BouncingBar,
	"Digits":      Digits,
	"Quotes":      Quotes,
	"Cat":         Cat,
	"Money":       Money,
	"Beach":       Beach,
	"Music":       Music,
	"Target":      Target,
}

// All retorna todos os spinners predefinidos em um mapa nome->spinner.
//
// O mapa é compartilhado entre as chamadas e não deve ser modificado. Use
// Clone para obter uma cópia independente.
func All() map[string]Spinner {
	return predefinidos
}

// Clone devolve uma cópia independente do mapa com todos os spinners
// predefinidos.
func Clone() map[string]Spinner {
	copia := make(map[string]Spinner, len(predefinidos))
	for nome, s := range predefinidos {
		copia[nome] = s
	}
	return copia
}

// Names devolve os nomes de todos os spinners predefinidos, em ordem
// alfabética.
func Names() []string {
	nomes := make([]string, 0, len(predefinidos))
	for nome := range predefinidos {
		nomes = append(nomes, nome)
	}
	sort.Strings(nomes)
	return nomes
}
