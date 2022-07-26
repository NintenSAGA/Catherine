package graphics

import (
  "github.com/gdamore/tcell/v2"
  "golang.org/x/sys/unix"
  "math"
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

func GetWidth(texts []string) int {
  width := math.MinInt
  for _, row := range texts {
    if w := len([]rune(row)); w > width {
      width = w
    }
  }
  return width
}

func DrawBox(s tcell.Screen, pattern [3][3]rune, style tcell.Style, alpha bool, fromX int, fromY int, w int, h int) {
  // 1. Corners
  s.SetContent(fromX+0, fromY+0, pattern[0][0], nil, style)
  s.SetContent(fromX+w-1, fromY+0, pattern[0][2], nil, style)
  s.SetContent(fromX+0, fromY+h-1, pattern[2][0], nil, style)
  s.SetContent(fromX+w-1, fromY+h-1, pattern[2][2], nil, style)

  // 2. Borders
  DrawLine(s, pattern[0][1], style, fromX+1, fromY+0, fromX+w-2, fromY+0)
  DrawLine(s, pattern[2][1], style, fromX+1, fromY+h-1, fromX+w-2, fromY+h-1)
  DrawLine(s, pattern[1][0], style, fromX+0, fromY+1, fromX+0, fromY+h-2)
  DrawLine(s, pattern[1][2], style, fromX+w-1, fromY+1, fromX+w-1, fromY+h-2)

  // 3. Content
  if cell := pattern[1][1]; !alpha || cell != ' ' {
    for x := 1; x < w-1; x++ {
      for y := 1; y < h-1; y++ {
        s.SetContent(fromX+x, fromY+y, cell, nil, style)
      }
    }
  }
}

func DrawBackGround(s tcell.Screen, pattern [3][3]rune, style tcell.Style) {
  w, h := s.Size()
  DrawBox(s, pattern, style, false, 0, 0, w, h)
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

  DrawTile(s, content, style, alpha, cOffsetX, cOffsetY)
}

func DrawTile(s tcell.Screen, content []string, style tcell.Style, alpha bool, offsetX int, offsetY int) {
  for y, row := range content {
    row := []rune(row)
    for x, r := range row {
      if alpha && r == ' ' {
        continue
      }
      s.SetContent(offsetX+x, offsetY+y, r, nil, style)
    }
  }
}
