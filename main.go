package main

import (
  "net/http"
  _ "net/http/pprof"

  "github.com/macmv/simple-gl"

  gogl "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/mathgl/mgl32"
)

func main() {
  go func() {
    http.ListenAndServe("localhost:6060", nil)
  }()

  gl.Init()
  window := gl.NewWindow("CA sim", 1920, 1080)
  defer window.Close()

  size := 100

  w := NewWorld(
    size,
    NewRule(13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26),
    NewRule(13, 14, 17, 18, 19),
    2,
  )

  shader, err := gl.NewShader("vertex.glsl", "fragment.glsl")
  if err != nil {
    panic(err)
  }
  window.Use(shader)
  shader.LoadPerspective(window, 0.1, 10)
  window.Finish()

  t, err := gl.NewTextureFromFile("test.png")
  if err != nil {
    panic(err)
  }

  rot := mgl32.Rotate3DY(0.01).Mat4()

  world_trans := rot

  i := 0
  step := 0
  for !window.Closed() {
    i++
    if i > 2 {
      i = 0
      // w.Update()
      step++
    }

    window.Use(shader)
    shader.LoadCamera(2, 3, -3)
    world_trans = rot.Mul4(world_trans)
    // for l, p := range w.alive {
    //   w.cube.Transform = world_trans.Mul4(p.trans)
    //   v := mgl32.Vec3{float32(l.x) - 50, float32(l.y) - 50, float32(l.z) - 50}
    //   dist := v.Len() / 100
    //   w.cube.Color = mgl32.Vec3{1 - dist, 1, 1}
    //
    //   shader.StoreUniformMat4f("model", w.cube.Transform)
    //   shader.StoreUniform3f("color", w.cube.Color)
    //   break
    // }
    w.cube.Transform = world_trans.Mul4(w.locs[0][0][0].trans)

    shader.StoreUniformMat4f("model", w.cube.Transform)
    shader.StoreUniform3f("color", w.cube.Color)

    w.cube.Vao().Bind()

    gogl.ActiveTexture(gogl.TEXTURE0);
    t.Bind();
    gogl.DrawElementsInstanced(gogl.TRIANGLES, w.cube.Vao().Length(), gogl.UNSIGNED_INT, nil, int32(size * size * size))
    w.cube.Vao().Unbind()

    w.cube.Transform = world_trans.Mul4(w.locs[0][0][0].trans)
    v := mgl32.Vec3{0, 0, 0}
    dist := v.Len() / 100
    w.cube.Color = mgl32.Vec3{1 - dist, 1, 1}
    window.Finish()

    window.Sync()
  }
}
