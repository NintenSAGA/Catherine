package scene

import (
  "github.com/gdamore/tcell/v2"
  "path"
  "push_blocks/src/graphics"
  "push_blocks/src/resource"
)

var figure = resource.ReadText(path.Join(resource.ASCIIArt, "background.txt"))
var title = resource.ReadText(path.Join(resource.ASCIIArt, "title.txt"))

var prompt1 = []string{
  "Press Enter to Start",
}

type Title struct {
  s      tcell.Screen
  finish bool

  flashPrinter graphics.FlashPrinter
}

func (t *Title) Init(s tcell.Screen) {
  t.s = s
  t.finish = false

  t.flashPrinter.Init("0.5s")
}

func (t *Title) LogicUpdate() {
  if got, key, _ := GetKey(t.s); got && key == tcell.KeyEnter {
    t.finish = true
  }
}

func (t *Title) drawSideBar() {
  w, h := t.s.Size()
  r := '*'
  graphics.DrawLine(t.s, r, graphics.NormalStyle, 1, 1, 1, h-2)
  graphics.DrawLine(t.s, r, graphics.NormalStyle, w-2, 1, w-2, h-2)
}

func (t *Title) Show() {
  graphics.DrawBackGround(t.s, graphics.NormalPattern, graphics.NormalStyle)
  t.drawSideBar()

  tH := len(title)
  tW := graphics.GetWidth(title)
  pW := graphics.GetWidth(prompt1)

  fH := len(figure)
  fW := graphics.GetWidth(figure)

  graphics.DrawTileCenterRelative(t.s, figure, graphics.NormalStyle, true, -(fW / 2), -(fH/2 + 2))
  graphics.DrawTileCenterRelative(t.s, title, graphics.NormalStyle, true, -(tW / 2), -(tH/2 + 2))

  t.flashPrinter.Show(func() {
    graphics.DrawTileCenterRelative(t.s, prompt1, graphics.NormalStyle, false, -(pW / 2), tH/2)
  })

}

func (t *Title) IsFinished() bool {
  return t.finish
}
