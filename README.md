# goclean

[![Go Reference](https://pkg.go.dev/badge/github.com/sonnt85/goclean.svg)](https://pkg.go.dev/github.com/sonnt85/goclean)

Cross-platform terminal screen clearing and cursor control for Go.

## Installation

```bash
go get github.com/sonnt85/goclean
```

## Features

- Clear the terminal screen on Linux, macOS, and Windows
- Move the cursor to the top-left position
- Query terminal width and height
- Native Windows implementation via `kernel32.dll` (no ANSI codes)
- ANSI escape code implementation for Unix-like systems

## Usage

```go
package main

import (
    "fmt"
    "time"

    "github.com/sonnt85/goclean"
)

func main() {
    goclean.Clear()

    for {
        goclean.MoveTopLeft()

        w, h := goclean.Size()
        fmt.Printf("Width: %d Height: %d\n", w, h)
        fmt.Println(time.Now())

        time.Sleep(time.Second)
    }
}
```

## API

- `Clear()` — clears the terminal screen
- `MoveTopLeft()` — moves the cursor to position (0, 0)
- `Size() (int, int)` — returns terminal width and height in columns/rows

## Author

**sonnt85** — [thanhson.rf@gmail.com](mailto:thanhson.rf@gmail.com)

## License

MIT License - see [LICENSE](LICENSE) for details.
