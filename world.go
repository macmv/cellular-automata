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
  alive map[loc]*point
  survives *Rule
  born *Rule
  max int
}

func NewWorld(size int, survives, born *Rule, max int) *World {
  w := World{}
  w.size = size
  w.locs = [][][]*point{}
  w.alive = make(map[loc]*point)
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
  (&w).SetArea(0, 0, 0, size - 1, size - 1, size - 1, true)
  for i := 0; i < 30000; i++ {
    x := rand.Intn(size - 3)
    y := rand.Intn(size - 3)
    z := rand.Intn(size - 3)
    (&w).SetArea(x, y, z, x + 1, y + 1, z + 1, false)
  }
  return &w
}

func (w *World) SetArea(x1, y1, z1, x2, y2, z2 int, val bool) {
  v := 0
  if val {
    v = w.max - 1
  }
  for x := x1; x <= x2; x++ {
    for y := y1; y <= y2; y++ {
      for z := z1; z <= z2; z++ {
        w.locs[x][y][z].state = v
        if val {
          w.alive[loc{x, y, z}] = w.locs[x][y][z]
        } else {
          delete(w.alive, loc{x, y, z})
        }
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
          w.alive[l] = p
          new_locs[l] = w.max
        } else if p.state == 1 && !w.survives.True(nearby) { // alive, and does not survive
          delete(w.alive, l)
          new_locs[l] = 0
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
