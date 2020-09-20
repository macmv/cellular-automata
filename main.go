package main

import (
  "github.com/macmv/simple-gl"
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

  for !window.Closed() {
    window.Use(shader)
    window.Render(cube)
    window.Finish()

    window.Sync()
  }
}
