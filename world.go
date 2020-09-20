package main

import (
  "math/rand"
  "github.com/macmv/simple-gl"

  "github.com/go-gl/mathgl/mgl32"
)

type loc struct {
  x, y, z int
}

type point struct {
  trans mgl32.Mat4
  state int
}

type World struct {
  size int
  cube *gl.Model
  locs [][][]*point
  survives *Rule
  born *Rule
  max int
}

func NewWorld(size int, survives, born *Rule, max int) *World {
  w := World{}
  w.size = size
  w.locs = [][][]*point{}
  w.survives = survives
  w.born = born
  w.max = max

  scale := float32(1) / float32(size) * 3
  w.cube = gl.NewCube(scale / 2, scale / 2, scale / 2)
  for x := 0; x < size; x++ {
    slice := [][]*point{}
    for y := 0; y < size; y++ {
      row := []*point{}
      for z := 0; z < size; z++ {
        fx := scale * float32(x - size / 2)
        fy := scale * float32(y - size / 2)
        fz := scale * float32(z - size / 2)
        trans := mgl32.Translate3D(fx, fy, fz).Mul4(w.cube.Transform)
        row = append(row, &point{trans: trans, state: 0})
      }
      slice = append(slice, row)
    }
    w.locs = append(w.locs, slice)
  }
  // (&w).SetArea(0, 0, 0, size - 1, size - 1, size - 1, true)
  (&w).SetArea(size / 2 - 5, size / 2 - 5, size / 2 - 5, 10, 10, 10, func() bool { return rand.Intn(10) < 9 })
  // for i := 0; i < 80000; i++ {
  //   x := rand.Intn(size - 1)
  //   y := rand.Intn(size - 1)
  //   z := rand.Intn(size - 1)
  //   (&w).SetArea(x, y, z, x + 1, y + 1, z + 1, false)
  // }
  return &w
}

func (w *World) SetArea(x1, y1, z1, w1, h1, d1 int, val func() bool) {
  for x := x1; x <= x1 + w1; x++ {
    for y := y1; y <= y1 + h1; y++ {
      for z := z1; z <= z1 + d1; z++ {
        v := 0
        if val() {
          v = w.max - 1
        }
        w.locs[x][y][z].state = v
      }
    }
  }
}

func (w *World) Update() {
  new_locs := make(map[loc]int)
  for x, slice := range w.locs {
    for y, row := range slice {
      for z, p := range row {
        l := loc{x, y, z}
        nearby := w.find_nearby(l)
        if p.state == 0 && w.born.True(nearby) { // dead, and is born
          new_locs[l] = w.max - 1
        } else if p.state == w.max - 1 && !w.survives.True(nearby) { // alive, and does not survive
          new_locs[l] = w.max - 2
        } else if p.state != 0 {
          new_locs[l] = p.state - 1
        }
      }
    }
  }
  for l, v := range new_locs {
    w.locs[l.x][l.y][l.z].state = v
  }
}

func (w *World) find_nearby(l loc) int {
  total := 0
  for x := l.x - 1; x < l.x + 2; x++ {
    if x < 0 || x >= len(w.locs) {
      continue
    }
    for y := l.y - 1; y < l.y + 2; y++ {
      if y < 0 || y >= len(w.locs) {
        continue
      }
      for z := l.z - 1; z < l.z + 2; z++ {
        if z < 0 || z >= len(w.locs) {
          continue
        }
        if x == l.x && y == l.y && z == l.z {
          continue
        }
        p := w.locs[x][y][z]
        if p.state != 0 {
          total++
        }
      }
    }
  }
  return total
}

func (w *World) TextureData() []uint8 {
  data := make([]uint8, w.size * w.size * w.size * 4)
  for x := 0; x < w.size; x++ {
    for y := 0; y < w.size; y++ {
      for z := 0; z < w.size; z++ {
        data[(x * w.size * w.size + y * w.size + z) * 4 + 0] = 0
        data[(x * w.size * w.size + y * w.size + z) * 4 + 1] = 0
        data[(x * w.size * w.size + y * w.size + z) * 4 + 2] = 255
        data[(x * w.size * w.size + y * w.size + z) * 4 + 3] = uint8(w.locs[x][y][z].state)
      }
    }
  }
  return data
}
