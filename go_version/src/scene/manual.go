package scene

import (
  "github.com/gdamore/tcell/v2"
  "path"
  "push_blocks/src/graphics"
  "push_blocks/src/resource"
)

var intro = resource.ReadText(path.Join(resource.ASCIIArt, "introduction.txt"))
var manual = resource.ReadText(path.Join(resource.ASCIIArt, "manual.txt"))

var manualTexts = [][]string{
  intro,
  manual,
}

var prompt2 = []string{
  "Press Enter to Continue",
}

type Manual struct {
  s     tcell.Screen
  phase int

  flashPrinter graphics.FlashPrinter
}

func (m *Manual) Init(s tcell.Screen) {
  m.s = s
  m.phase = 0

  m.flashPrinter.Init("0.5s")
}

func (m *Manual) LogicUpdate() {
  if got, key, _ := GetKey(m.s); got && key == tcell.KeyEnter {
    m.phase++
  }
}

func (m *Manual) Show() {
  if m.IsFinished() {
    return
  }

  graphics.DrawBackGround(m.s, graphics.NormalPattern, graphics.NormalStyle)

  text := manualTexts[m.phase]
  tH, tW := len(text), graphics.GetWidth(text)
  graphics.DrawTileCenterRelative(m.s, text, graphics.NormalStyle, false,
    -(tW / 2), -(tH/2 + 2))

  m.flashPrinter.Show(func() {
    graphics.DrawTileCenterRelative(m.s, prompt2, graphics.NormalStyle, false,
      -(len(prompt2[0]) / 2), tH/2+2)
  })
}

func (m *Manual) IsFinished() bool {
  return m.phase >= len(manualTexts)
}
