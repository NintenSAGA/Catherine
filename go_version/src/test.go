package main

import (
  "bytes"
  "embed"
  "encoding/csv"
  "fmt"
  "path"
)

//go:embed resource/**
var ff embed.FS

func main() {
  dir := "resource/levels"
  entries, _ := ff.ReadDir(dir)
  for _, entry := range entries {
    fBytes, _ := ff.ReadFile(path.Join(dir, entry.Name()))
    csvReader := csv.NewReader(bytes.NewReader(fBytes))
    a, _ := csvReader.ReadAll()
    fmt.Println(a)
  }
}
