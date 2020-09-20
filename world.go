package main

import (
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
}

func NewWorld(size int) *World {
  w := World{}
  w.size = size
  w.locs = [][][]*point{}
  w.alive = make(map[loc]*point)

  scale := float32(1) / float32(size) * 2
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
  a := loc{size / 2, size / 2, size / 2}
  b := loc{size / 2 + 1, size / 2, size / 2}
  c := loc{size / 2, size / 2 + 1, size / 2}
  d := loc{size / 2 + 1, size / 2 + 1, size / 2}
  w.locs[size / 2][size / 2][size / 2].state = 1
  w.locs[size / 2 + 1][size / 2][size / 2].state = 1
  w.locs[size / 2][size / 2 + 1][size / 2].state = 1
  w.locs[size / 2 + 1][size / 2 + 1][size / 2].state = 1
  w.alive[a] = w.locs[size / 2][size / 2][size / 2]
  w.alive[b] = w.locs[size / 2 + 1][size / 2][size / 2]
  w.alive[c] = w.locs[size / 2][size / 2 + 1][size / 2]
  w.alive[d] = w.locs[size / 2 + 1][size / 2 + 1][size / 2]
  return &w
}

func (w *World) Update() {
  new_locs := make(map[loc]int)
  for x, slice := range w.locs {
    for y, row := range slice {
      for z, p := range row {
        l := loc{x, y, z}
        nearby := w.find_nearby(l)
        if p.state == 0 && (nearby == 4 || nearby == 4) {
          new_locs[l] = 1
          w.alive[l] = p
        } else if p.state != 0 && (nearby < 2 || nearby > 4) {
          new_locs[l] = 0
          delete(w.alive, l)
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
        p := w.locs[x][y][z]
        if p.state != 0 {
          total++
        }
      }
    }
  }
  return total
}
