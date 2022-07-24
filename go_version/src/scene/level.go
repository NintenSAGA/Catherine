package scene

import (
  "fmt"
  "github.com/emirpasic/gods/stacks/linkedliststack"
  "github.com/gdamore/tcell/v2"
  "path"
  "push_blocks/src/graphics"
  "push_blocks/src/resource"
  "unicode"
)

const (
  upperPromptF    = "Level:%v\t\tSteps:%v \tItems:%v"
  lowerPrompt     = "◇Move:W,S,A,D  ◇Restart:R  ◇Exit:esc  ◇Item:I"
  nextLevelPrompt = "Press Enter to Continue"

  cellNone   = "0"
  cellPlayer = "1"
  cellBlock  = "2"
  cellWall   = "3"
  cellDest   = "4"
  cellItem   = "7"

  phaseInit    = 1
  phaseCleared = 2
  phaseNext    = 3
)

var clearPrompt = resource.ReadText(path.Join(resource.ASCIIArt, "cleared.txt"))

var symbolMap = map[string]rune{
  cellNone:   ' ',
  cellPlayer: '○',
  cellBlock:  '◆',
  cellWall:   tcell.RuneBoard,
  cellDest:   '◇',
  cellItem:   '△',
}

type Position struct {
  x int
  y int
}

type Push struct {
  prevPos Position
  vec     []int
}

type Level struct {
  s          tcell.Screen
  Data       [][]string
  backupData [][]string

  phase int

  LevelID int
  step    int
  numItem int

  playerInfo  Position
  history     *linkedliststack.Stack
  targetInfos []Position

  flashPrinter graphics.FlashPrinter
}

func (l *Level) Init(s tcell.Screen) {
  l.s = s
  l.step = 0
  l.numItem = 0
  l.phase = phaseInit
  l.flashPrinter.Init("1s")

  l.backupData = make([][]string, 0, len(l.Data))
  for i, row := range l.Data {
    dup := make([]string, len(row))
    copy(dup, row)
    l.backupData = append(l.backupData, dup)

    for j, cell := range row {
      switch cell {
      case cellPlayer:
        l.playerInfo.x, l.playerInfo.y = j, i
        row[j] = cellNone
      case cellDest:
        l.targetInfos = append(l.targetInfos, Position{x: j, y: i})
        row[j] = cellNone
      }
    }
  }

  l.history = linkedliststack.New()
}

func (l *Level) actionUpdate(r rune) {
  var v []int = nil

  switch unicode.ToUpper(r) {
  case 'W':
    v = []int{0, -1}
  case 'S':
    v = []int{0, 1}
  case 'A':
    v = []int{-1, 0}
  case 'D':
    v = []int{1, 0}
  case 'R':
    l.Data = l.backupData
    l.Init(l.s)
  case 'I':
    if l.numItem > 0 && l.history.Size() > 0 {
      l.numItem--
      // Block should track back too!
      val, _ := l.history.Pop()
      push := val.(Push)
      l.playerInfo = push.prevPos
      from := &l.Data[push.prevPos.y+push.vec[1]][push.prevPos.x+push.vec[0]]
      to := &l.Data[push.prevPos.y+(push.vec[1])*2][push.prevPos.x+(push.vec[0])*2]
      *(to) = cellNone
      *(from) = cellBlock
    }
  case '+':
    // Debug only, should be deleted in release version
    l.phase = phaseCleared
  }

  if v == nil {
    return
  }

  nxtPos := Position{x: l.playerInfo.x + v[0], y: l.playerInfo.y + v[1]}
  cellPtr := &l.Data[nxtPos.y][nxtPos.x]

  switch *(cellPtr) {
  case cellWall:
    return
  case cellBlock:
    nxtCellPtr := &l.Data[nxtPos.y+v[1]][nxtPos.x+v[0]]
    if *(nxtCellPtr) == cellNone {
      *(cellPtr) = cellNone
      *(nxtCellPtr) = cellBlock
      l.history.Push(Push{prevPos: l.playerInfo, vec: v})
    }
  case cellItem:
    l.numItem++
    *(cellPtr) = cellNone
  }

  if *(cellPtr) == cellNone {
    l.playerInfo = nxtPos
    l.step++
  }
}

func (l *Level) LogicUpdate() {
  if got, key, r := GetKey(l.s); got {
    switch l.phase {
    case phaseInit:
      l.actionUpdate(r)
    case phaseCleared:
      if key == tcell.KeyEnter {
        l.phase = phaseNext
      }
    }
  }
}

func (l *Level) placeAt(x int, y int, cell string) {
  l.s.SetContent(x, y, symbolMap[cell], nil, graphics.NormalStyle)
}

func (l *Level) mapRender() {
  sw, sh := l.s.Size()
  h, w := len(l.Data), len(l.Data[0])
  offsetX, offsetY := sw/2-w/2, sh/2-h/2

  for _, targetInfo := range l.targetInfos {
    l.placeAt(offsetX+targetInfo.x, offsetY+targetInfo.y, cellDest)
  }

  for i, row := range l.Data {
    for j, cell := range row {
      if cell != cellNone {
        l.placeAt(offsetX+j, offsetY+i, cell)
      }
    }
  }

  l.placeAt(offsetX+l.playerInfo.x, offsetY+l.playerInfo.y, cellPlayer)
}

func (l *Level) textCenterAligned(text string, offsetY int) {
  l.tileCenterAligned([]string{text}, offsetY)
}

func (l *Level) tileCenterAligned(tile []string, offsetY int) {
  graphics.DrawTileCenterRelative(l.s, tile, graphics.NormalStyle, false,
    -(graphics.GetWidth(tile) / 2), offsetY)
}

func (l *Level) Show() {
  _, h := l.s.Size()
  graphics.DrawBackGround(l.s, graphics.NormalPattern, graphics.NormalStyle)

  upperPrompt := fmt.Sprintf(upperPromptF, l.LevelID, l.step, l.numItem)
  l.textCenterAligned(upperPrompt, -(h / 2))
  l.textCenterAligned(lowerPrompt, h/2-1)

  switch l.phase {
  case phaseInit:
    l.mapRender()
  case phaseCleared:
    l.tileCenterAligned(clearPrompt, -(len(clearPrompt)/2 + 2))
    l.flashPrinter.Show(func() {
      l.textCenterAligned(nextLevelPrompt, 2)
    })
  }
}

func (l *Level) IsFinished() (result bool) {
  result = false

  switch l.phase {
  case phaseInit:
    for _, t := range l.targetInfos {
      if l.Data[t.y][t.x] != cellBlock {
        return
      }
    }
    l.phase = phaseCleared
  case phaseCleared:
  case phaseNext:
    result = true
  }

  return
}
