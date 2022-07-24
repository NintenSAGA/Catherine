package main

import (
  "github.com/gdamore/tcell/v2"
  "log"
  "push_blocks/src/scene"
  "runtime"
)

func main() {
  switch {
  case runtime.GOOS == "windows":
    log.Fatal("Windows is not supported yet")
  }

  s, _ := tcell.NewScreen()
  if err := s.Init(); err != nil {
    log.Fatalf("%v\r\n%v", err, "This program should run in a terminal")
  }

  scenes := []scene.Scene{
    &scene.Title{},
    &scene.Manual{},
  }

  for _, eachScene := range scenes {
    eachScene.Init(s)
    for !eachScene.IsFinished() {
      eachScene.LogicUpdate()
      eachScene.Show()
      s.Show()
    }
  }

  s.Fini()
}
