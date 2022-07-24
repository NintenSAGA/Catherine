package graphics

import "fmt"

// Functions for ANSI Escape Codes

const (
  EraseScreen = "2J"

  SetHidden   = "8m"
  UnsetHidden = "28m"
)

func ESC(code string) {
  fmt.Print("\x1b[" + code)
}
