package scene

import (
  "github.com/gdamore/tcell/v2"
  "path"
  "push_blocks/src/graphics"
  "push_blocks/src/resource"
)

var ending = resource.ReadText(path.Join(resource.ASCIIArt, "ending.txt"))

var endingTexts = [][]string{
  ending,
}

type Ending struct {
  s     tcell.Screen
  phase int

  flashPrinter graphics.FlashPrinter
}

func (e *Ending) Init(s tcell.Screen) {
  e.s = s
  e.phase = 0

  e.flashPrinter.Init("0.5s")
}

func (e *Ending) LogicUpdate() {
  if got, key, _ := GetKey(e.s); got && key == tcell.KeyEnter {
    e.phase++
  }
}

func (e *Ending) Show() {
  if e.IsFinished() {
    return
  }

  graphics.DrawBackGround(e.s, graphics.NormalPattern, graphics.NormalStyle)

  text := endingTexts[e.phase]
  tH, tW := len(text), graphics.GetWidth(text)
  graphics.DrawTileCenterRelative(e.s, text, graphics.NormalStyle, false,
    -(tW / 2), -(tH/2 + 2))

  e.flashPrinter.Show(func() {
    graphics.DrawTileCenterRelative(e.s, prompt2, graphics.NormalStyle, false,
      -(len(prompt2[0]) / 2), tH/2+2)
  })
}

func (e *Ending) IsFinished() bool {
  return e.phase >= len(endingTexts)
}
