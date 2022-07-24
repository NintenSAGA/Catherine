package scene

import (
  "github.com/gdamore/tcell/v2"
  "push_blocks/src/graphics"
)

var texts = [][]string{
  {
    "Your name is Vincent.                           ",
    "",
    "Because of your getting too close with a girl   ",
    "called Catherine, you're cursed by your fiancee ",
    "Katherine.",
    "",
    "During a nightmare, you're stuck in a dungeon.  ",
    "",
    "To avoid being killed, you have to push all the ",
    "blocks to the correct places.                   ",
  },
  {
    "              Operation Manual              ",
    "                                            ",
    "   ○ You               ■ Unmovable Blocks   ",
    "   ◆ Object Blocks     ◇ Object Places      ",
    "   △ Magical Sand Clocks                    ",
    "                                            ",
    "*You can turn back time by entering 'I' when",
    " holding a Magical Sand Clock*              ",
  },
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

  text := texts[m.phase]
  tH, tW := len(text), len(text[0])
  graphics.DrawTileCenterRelative(m.s, text, graphics.NormalStyle, false,
    -(tW / 2), -(tH/2 + 2))

  m.flashPrinter.Show(func() {
    graphics.DrawTileCenterRelative(m.s, prompt2, graphics.NormalStyle, false,
      -(len(prompt2[0]) / 2), tH/2)
  })
}

func (m *Manual) IsFinished() bool {
  return m.phase >= len(texts)
}
