# charm-spinner

Colecao de 30+ animacoes de spinner para terminal em Go, construido sobre [Lip Gloss](https://github.com/charmbracelet/lipgloss).

> **[Read in English](README.md)**

## Instalacao

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

## Spinners Disponiveis

Mais de 30 spinners predefinidos: Dots, Line, Star, Moon, Earth, Clock, Bounce, Bar, Circle, Square, Triangle, Arrow, Pulse, Flip, Grow, Layer, Toggle, Balloon, Noise, Fading, Christmas, Monkey, Hearts, Weather, Soccer, Matrix, DNA, Orbit, Rocket, Fire.

Use `spinner.All()` para obter um mapa com todos os spinners.

## Spinners Customizados

```go
custom := spinner.Spinner{
    Frames:   []string{"🐱", "🐶", "🐭"},
    Interval: 200 * time.Millisecond,
}
```

## Licenca

[MIT](LICENSE)
