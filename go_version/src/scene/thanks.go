package scene

import (
  "github.com/gdamore/tcell/v2"
  "path"
  "push_blocks/src/graphics"
  "push_blocks/src/resource"
)

var thanks = resource.ReadText(path.Join(resource.ASCIIArt, "thanks.txt"))

type Thanks struct {
  s      tcell.Screen
  finish bool

  flashPrinter graphics.FlashPrinter
}

func (t *Thanks) Init(s tcell.Screen) {
  t.s = s
  t.finish = false

  t.flashPrinter.Init("0.5s")
}

func (t *Thanks) LogicUpdate() {
  if got, key, _ := GetKey(t.s); got && key == tcell.KeyEnter {
    t.finish = true
  }
}

func (t *Thanks) drawSideBar() {
  w, h := t.s.Size()
  r := '*'
  graphics.DrawLine(t.s, r, graphics.NormalStyle, 1, 1, 1, h-2)
  graphics.DrawLine(t.s, r, graphics.NormalStyle, w-2, 1, w-2, h-2)
}

func (t *Thanks) Show() {
  graphics.DrawBackGround(t.s, graphics.NormalPattern, graphics.NormalStyle)
  t.drawSideBar()

  tH := len(thanks)
  tW := graphics.GetWidth(thanks)

  graphics.DrawTileCenterRelative(t.s, thanks, graphics.NormalStyle, true, -(tW / 2), -(tH / 2))
}

func (t *Thanks) IsFinished() bool {
  return t.finish
}
