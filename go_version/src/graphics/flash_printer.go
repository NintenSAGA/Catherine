package graphics

import "time"

type FlashPrinter struct {
  show     bool
  duration time.Duration
  timer    *time.Timer
}

func (fp *FlashPrinter) Init(rawTime string) {
  fp.show = true
  fp.duration, _ = time.ParseDuration(rawTime)
  fp.timer = time.NewTimer(fp.duration)
}

func (fp *FlashPrinter) Show(printFunc func()) {
  select {
  case <-fp.timer.C:
    fp.show = !fp.show
    fp.timer.Reset(fp.duration)
  default:
  }

  if fp.show {
    printFunc()
  }
}
