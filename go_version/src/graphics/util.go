package graphics

import (
  "github.com/gdamore/tcell/v2"
  "golang.org/x/sys/unix"
  "syscall"
)

func GetTerminalInfo() (row uint16, col uint16) {
  winSize, err := unix.IoctlGetWinsize(syscall.Stdin, syscall.TIOCGWINSZ)

  if err != nil {
    row, col = 0, 0
  } else {
    row, col = winSize.Row, winSize.Col
  }

  return
}

func DrawBackGround(s tcell.Screen, pattern [3][3]rune, style tcell.Style) {
  w, h := s.Size()

  // 1. Corners
  s.SetContent(0, 0, pattern[0][0], nil, style)
  s.SetContent(w-1, 0, pattern[0][2], nil, style)
  s.SetContent(0, h-1, pattern[2][0], nil, style)
  s.SetContent(w-1, h-1, pattern[2][2], nil, style)

  // 2. Borders
  DrawLine(s, pattern[0][1], style, 1, 0, w-2, 0)
  DrawLine(s, pattern[2][1], style, 1, h-1, w-2, h-1)
  DrawLine(s, pattern[1][0], style, 0, 1, 0, h-2)
  DrawLine(s, pattern[1][2], style, w-1, 1, w-1, h-2)

  // 3. Content
  for x := 1; x < w-1; x++ {
    for y := 1; y < h-1; y++ {
      s.SetContent(x, y, pattern[1][1], nil, style)
    }
  }
}

func DrawLine(s tcell.Screen, r rune, style tcell.Style, fromX int, fromY int, toX int, toY int) {
  if !(fromX <= toX && fromY <= toY) {
    return
  }
  x, y := fromX, fromY
  for x != toX || y != toY {
    s.SetContent(x, y, r, nil, style)

    if x < toX {
      x++
    }
    if y < toY {
      y++
    }
  }
  s.SetContent(x, y, r, nil, style)
}

func DrawTileCenterRelative(s tcell.Screen, content []string, style tcell.Style, alpha bool, offsetX int, offsetY int) {
  w, h := s.Size()
  cOffsetX, cOffsetY := w/2, h/2

  cOffsetX += offsetX
  cOffsetY += offsetY

  DrawTile(s, content, style, cOffsetX, cOffsetY, alpha)
}

func DrawTile(s tcell.Screen, content []string, style tcell.Style, offsetX int, offsetY int, alpha bool) {
  for y, row := range content {
    for x, r := range row {
      if alpha && r == ' ' {
        continue
      }
      s.SetContent(offsetX+x, offsetY+y, r, nil, style)
    }
  }
}
