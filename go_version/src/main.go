package main

import (
  "bytes"
  "encoding/csv"
  "github.com/gdamore/tcell/v2"
  "log"
  "path"
  "push_blocks/src/resource"
  "push_blocks/src/scene"
  "runtime"
  "sort"
  "strings"
)

func main() {
  /* 1. Launch */
  switch {
  case runtime.GOOS == "windows":
    log.Fatal("Windows is not supported yet")
  }

  /* 2. Screen setup */
  s, _ := tcell.NewScreen()
  if err := s.Init(); err != nil {
    log.Fatalf("%v\r\n%v", err, "This program should run in a terminal")
  }

  /* 3. Load scenes */
  // 3.1 Title and manual
  scenes := []scene.Scene{
    &scene.Title{},
    &scene.Manual{},
  }

  // 3.2 resource.Levels
  levelDir := resource.Levels
  levelCount := 0

  if entries, err := resource.F.ReadDir(levelDir); err != nil {
    log.Fatalf(err.Error())

  } else {
    sort.Slice(entries, func(i, j int) bool {
      return entries[i].Name() < entries[j].Name()
    })

    for _, entry := range entries {
      name := entry.Name()
      if strings.HasPrefix(name, "level") && strings.HasSuffix(name, ".csv") {
        raw, _ := resource.F.ReadFile(path.Join(levelDir, entry.Name()))
        data, _ := csv.NewReader(bytes.NewReader(raw)).ReadAll()
        scenes = append(scenes, &scene.Level{Data: data, LevelID: levelCount + 1})
        levelCount++
      }
    }
  }

  // 3.3 Ending
  scenes = append(scenes, &scene.Ending{}, &scene.Thanks{})

  for _, eachScene := range scenes {
    eachScene.Init(s)
    for !eachScene.IsFinished() {
      eachScene.LogicUpdate()
      eachScene.Show()
      s.Show()
    }
  }

  s.Fini()
}
