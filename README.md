# charm-spinner

Collection of 40+ beautiful terminal spinner animations for Go, built on top of [Lip Gloss](https://github.com/charmbracelet/lipgloss).

> **[Leia em Português](README.pt-br.md)**

## Install

```bash
go get github.com/junhinhow/charm-spinner@latest
```

## Usage

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
        fmt.Printf("\r%s Loading...", s.Render(i))
        time.Sleep(s.Interval)
    }
    fmt.Println()
}
```

`Render` cycles through the frames automatically: any index is valid, and
negative indexes count from the end of the cycle. A spinner without frames
renders an empty string instead of panicking.

## Available Spinners

| Name | Preview |
|------|---------|
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

Use `spinner.All()` to get a map with all the spinners, `spinner.Names()` to
get their names in alphabetical order and `spinner.Clone()` to get an
independent copy of the map.

## Custom Spinners

```go
custom := spinner.Spinner{
    Frames:   []string{"🐱", "🐶", "🐭"},
    Interval: 200 * time.Millisecond,
}
```

## Styles

Every spinner carries its own `lipgloss.Style`. When the style is not set, the
package default is used:

```go
import "charm.land/lipgloss/v2"

custom := spinner.Spinner{
    Frames:   []string{"🐱", "🐶", "🐭"},
    Interval: 200 * time.Millisecond,
    Style:    lipgloss.NewStyle().Bold(true),
}
```

The package default can be changed at any time -- even while other goroutines
are rendering -- with `SetDefaultStyle`:

```go
spinner.SetDefaultStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("205")))
```

> The exported `spinner.Style` variable is only the *initial* default: it is
> copied into the package when it is loaded. Assigning to it directly while
> another goroutine renders is a data race -- use `SetDefaultStyle` instead.

## License

[MIT](LICENSE)
