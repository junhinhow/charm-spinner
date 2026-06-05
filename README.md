# charm-spinner

Collection of 30+ beautiful terminal spinner animations for Go, built on top of [Lip Gloss](https://github.com/charmbracelet/lipgloss).

> **[Leia em Portugues](README.pt-br.md)**

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
| Grow | `▁ ▂ ▃ ▄ ▅ ▆ ▇ █` |
| Hearts | `💛 💙 💜 💚 ❤️` |
| Rocket | `🚀 💨 ✨ 🌟` |
| Fire | `🔥 💥 ✨` |
| ...and 13 more! | |

## Custom Spinners

```go
custom := spinner.Spinner{
    Frames:   []string{"🐱", "🐶", "🐭"},
    Interval: 200 * time.Millisecond,
}
```

## License

[MIT](LICENSE)
