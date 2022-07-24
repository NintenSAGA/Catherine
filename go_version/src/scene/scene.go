package scene

import (
  "github.com/gdamore/tcell/v2"
  "os"
)

type Scene interface {
  Init(s tcell.Screen)
  LogicUpdate()
  Show()
  IsFinished() bool
}

func GetKey(s tcell.Screen) (bool, tcell.Key, rune) {
  if !s.HasPendingEvent() {
    return false, 0, 0
  }

  ev := s.PollEvent()
  switch ev := ev.(type) {
  case *tcell.EventKey:
    if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
      // Exit
      s.Fini()
      os.Exit(0)
    }
    return true, ev.Key(), ev.Rune()
  }
  return false, 0, 0
}
