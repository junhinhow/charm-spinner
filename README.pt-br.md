# charm-spinner

Coleção de mais de 40 animações de spinner para terminal em Go, construído sobre [Lip Gloss](https://github.com/charmbracelet/lipgloss).

> **[Read in English](README.md)**

## Instalação

```bash
go get github.com/junhinhow/charm-spinner@latest
```

## Uso

```go
package main

import (
    "fmt"
    "time"

    spinner "github.com/junhinhow/charm-spinner"
)

func main() {
    s := spinner.Dots
    for i := 0; i < 40; i++ {
        fmt.Printf("\r%s Carregando...", s.Render(i))
        time.Sleep(s.Interval)
    }
    fmt.Println()
}
```

`Render` percorre os frames automaticamente: qualquer índice é válido e índices
negativos contam a partir do fim do ciclo. Um spinner sem frames devolve uma
string vazia em vez de entrar em panic.

## Spinners Disponíveis

| Nome | Prévia |
|------|--------|
| Dots | `⠋ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠏` |
| Line | `- \ \| /` |
| Star | `✶ ✸ ✹ ✺ ✹ ✷` |
| Moon | `🌑 🌒 🌓 🌔 🌕 🌖 🌗 🌘` |
| Earth | `🌍 🌎 🌏` |
| Clock | `🕐 🕑 🕒 ... 🕛` |
| Bounce | `⠁ ⠂ ⠄ ⠂` |
| Bar | `▏ ▎ ▍ ▌ ▋ ▊ ▉ █` |
| Circle | `◐ ◓ ◑ ◒` |
| Square | `▖ ▘ ▝ ▗` |
| Triangle | `◢ ◣ ◤ ◥` |
| Arrow | `← ↖ ↑ ↗ → ↘ ↓ ↙` |
| Pulse | `█ ▓ ▒ ░` |
| Flip | `_ - ` ` ' ´` |
| Grow | `▁ ▂ ▃ ▄ ▅ ▆ ▇ █` |
| Layer | `─ ━ ═ ━` |
| Toggle | `⊶ ⊷` |
| Balloon | `. o O ° O o .` |
| Noise | `▓ ▒ ░ ▒` |
| Fading | `░ ▒ ▓ █ ▓ ▒ ░` |
| Christmas | `🎄 🎅 ⛄ 🎁 🔔 ⭐` |
| Monkey | `🙈 🙉 🙊` |
| Hearts | `💛 💙 💜 💚 ❤️ 🧡` |
| Weather | `☀️ 🌤️ ⛅ 🌥️ ☁️ 🌧️ ⛈️` |
| Soccer | `⚽ 🥅 👟 💨` |
| Matrix | `ᚠ ᚡ ᚢ ᚣ ᚤ ᚥ` |
| DNA | `╔═╗ ╚═╝ ═══` |
| Orbit | `◜ ◝ ◞ ◟` |
| Rocket | `🚀 💨 ✨ 🌟` |
| Fire | `🔥 🔥 💥 ✨` |
| Ellipsis | `. .. ...` |
| Rectangle | `▛ ▜ ▙ ▟` |
| Diamond | `◆ ◇` |
| Hexagon | `⬢ ⬡` |
| BouncingBar | `[    ] [=   ] [==  ] [=== ]` |
| Digits | `0 1 2 3 4 5 6 7 8 9` |
| Quotes | `' " ❞ ❟` |
| Cat | `😺 😸 😹 😻 😼 😽` |
| Money | `💰 💵 💴 💶` |
| Beach | `🏖️ 🌴 🌊 ☀️` |
| Music | `♪ ♫ ♬` |
| Target | `◎ ◉ ⊙ ◉` |

Use `spinner.All()` para obter um mapa com todos os spinners, `spinner.Names()`
para obter os nomes em ordem alfabética e `spinner.Clone()` para obter uma
cópia independente do mapa.

## Spinners Customizados

```go
custom := spinner.Spinner{
    Frames:   []string{"🐱", "🐶", "🐭"},
    Interval: 200 * time.Millisecond,
}
```

## Estilos

Cada spinner carrega o próprio `lipgloss.Style`. Quando o estilo não é
definido, vale o estilo padrão do pacote:

```go
import "charm.land/lipgloss/v2"

custom := spinner.Spinner{
    Frames:   []string{"🐱", "🐶", "🐭"},
    Interval: 200 * time.Millisecond,
    Style:    lipgloss.NewStyle().Bold(true),
}
```

O estilo padrão do pacote pode ser trocado a qualquer momento -- inclusive
enquanto outras goroutines renderizam -- com `SetDefaultStyle`:

```go
spinner.SetDefaultStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("205")))
```

> A variável exportada `spinner.Style` é apenas o estilo padrão *inicial*: ela
> é copiada para o pacote durante a inicialização. Atribuir a ela diretamente
> enquanto outro goroutine renderiza é uma corrida de dados -- use
> `SetDefaultStyle`.

## Licença

[MIT](LICENSE)
