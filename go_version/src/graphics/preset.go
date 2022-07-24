package graphics

import "github.com/gdamore/tcell/v2"

var NormalPattern = [3][3]rune{
  {tcell.RuneULCorner, tcell.RuneHLine, tcell.RuneURCorner},
  {tcell.RuneVLine, ' ', tcell.RuneVLine},
  {tcell.RuneLLCorner, tcell.RuneHLine, tcell.RuneLRCorner},
}

var NormalStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
