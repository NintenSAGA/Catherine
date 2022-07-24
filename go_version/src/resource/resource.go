package resource

import (
  "embed"
  "strings"
)

const (
  Levels   = "levels"
  ASCIIArt = "ascii_art"
)

//go:embed **
var F embed.FS

func ReadText(name string) []string {
  b, _ := F.ReadFile(name)
  raw := string(b)
  return strings.Split(raw, "\n")
}
