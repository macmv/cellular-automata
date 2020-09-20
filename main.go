package main

import (
  "github.com/macmv/simple-gl"

  "github.com/go-gl/mathgl/mgl32"
)

func main() {
  gl.Init()
  window := gl.NewWindow("CA sim", 800, 600)
  defer window.Close()
  cube := gl.NewCube(1, 2, 1)

  shader, err := gl.NewShader("vertex.glsl", "fragment.glsl")
  if err != nil {
    panic(err)
  }
  window.Use(shader)
  shader.LoadPerspective(window, 0.1, 10)
  window.Finish()

  rot := mgl32.Rotate3DX(0.04).Mul3(mgl32.Rotate3DY(0.04)).Mul3(mgl32.Rotate3DZ(0.04))

  for !window.Closed() {
    cube.Transform = cube.Transform.Mul4(rot.Mat4())

    window.Use(shader)
    shader.LoadCamera(2, 3, -3)
    window.Render(cube)
    window.Finish()

    window.Sync()
  }
}
