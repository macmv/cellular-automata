package main

import (
  "net/http"
  _ "net/http/pprof"
  "github.com/macmv/simple-gl"

  "github.com/go-gl/mathgl/mgl32"
)

func main() {
  go func() {
    http.ListenAndServe("localhost:6060", nil)
  }()

  gl.Init()
  window := gl.NewWindow("CA sim", 800, 600)
  defer window.Close()

  w := NewWorld(100, nil)

  shader, err := gl.NewShader("vertex.glsl", "fragment.glsl")
  if err != nil {
    panic(err)
  }
  window.Use(shader)
  shader.LoadPerspective(window, 0.1, 10)
  window.Finish()

  rot := mgl32.Rotate3DY(0.01).Mat4()

  world_trans := rot

  i := 0
  step := 0
  for !window.Closed() {
    i++
    if i > 2 {
      i = 0
      w.Update()
      step++
    }

    window.Use(shader)
    shader.LoadCamera(2, 3, -3)
    world_trans = rot.Mul4(world_trans)
    for l, p := range w.alive {
      w.cube.Transform = world_trans.Mul4(p.trans)
      v := mgl32.Vec3{float32(l.x) - 50, float32(l.y) - 50, float32(l.z) - 50}
      dist := v.Len() / 100
      w.cube.Color = mgl32.Vec3{1 - dist, 1, 1}
      window.Render(w.cube)
    }
    window.Finish()

    window.Sync()
  }
}
