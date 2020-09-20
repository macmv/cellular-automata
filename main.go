package main

import (
  "github.com/macmv/simple-gl"

  "github.com/go-gl/mathgl/mgl32"
)

func main() {
  gl.Init()
  window := gl.NewWindow("CA sim", 800, 600)
  defer window.Close()
  cubes := []*gl.Model{}
  amount := 8
  scale := float32(1) / float32(amount) * 2
  for x := 0; x < amount; x++ {
    for y := 0; y < amount; y++ {
      for z := 0; z < amount; z++ {
        fx := scale * float32(x - amount / 2)
        fy := scale * float32(y - amount / 2)
        fz := scale * float32(z - amount / 2)
        cube := gl.NewCube(scale / 2, scale / 2, scale / 2)
        cube.Transform = mgl32.Translate3D(fx, fy, fz).Mul4(cube.Transform)
        cubes = append(cubes, cube)
      }
    }
  }

  shader, err := gl.NewShader("vertex.glsl", "fragment.glsl")
  if err != nil {
    panic(err)
  }
  window.Use(shader)
  shader.LoadPerspective(window, 0.1, 10)
  window.Finish()

  rot := mgl32.Rotate3DY(0.01)

  for !window.Closed() {
    window.Use(shader)
    shader.LoadCamera(2, 3, -3)
    for _, c := range cubes {
      c.Transform = rot.Mat4().Mul4(c.Transform)
      window.Render(c)
    }
    window.Finish()

    window.Sync()
  }
}
