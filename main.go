package main

import (
  "net/http"
  _ "net/http/pprof"

  gogl "golang.org/x/mobile/gl"

  "github.com/macmv/simple-gl"
  "github.com/macmv/simple-gl/core"
  "github.com/macmv/simple-gl/event"

  "github.com/go-gl/mathgl/mgl32"
)

func main() {
  go func() {
    http.ListenAndServe("localhost:6060", nil)
  }()

  size := 100

  survives := NewRule(2, 3, 4, 5)
  born := NewRule(6, 7)

  rot := mgl32.Rotate3DY(0.01).Mat4()

  world_trans := rot

  var w *World
  var t core.Texture
  var shader *core.Shader
  gl.On(gl.START, func(c core.Core) {
    w = NewWorld(
      c,
      size,
      survives,
      born,
      2,
    )
    // window := gl.NewWindow("CA sim", 1920, 1080)
    // defer window.Close()

    var err error
    shader, err = core.NewShader(c.Gl(), "vertex.glsl", "fragment.glsl")
    if err != nil {
      panic(err)
    }
    c.Window().Use(shader)
    shader.LoadPerspective(c.Window(), 0.1, 10)
    c.Window().Finish()

    t = c.NewTexture2DFromData(100, 100 * 100, w.TextureData())
  })
  i := 0
  step := 0
  gl.On(gl.DRAW, func(e event.Draw, c core.Core) {
    i++
    if i > 2 {
      i = 0
      w.Update()
      step++
    }

    c.Window().Use(shader)
    shader.LoadCamera(3, 3, -3)
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

    t.Data(w.TextureData())

    scale := 1 / float32(size) * 2
    w.cube.Transform = world_trans.Mul4(mgl32.Scale3D(scale, scale, scale))

    shader.StoreUniformMat4f("model", w.cube.Transform)
    shader.StoreUniform3f("color", w.cube.Color)

    w.cube.Vao().Bind()

    c.Gl().ActiveTexture(gogl.TEXTURE0)
    t.Bind()
    c.Gl().DrawElements(gogl.TRIANGLES, w.cube.Vao().Length(), gogl.UNSIGNED_INT, 0)
    w.cube.Vao().Unbind()

    w.cube.Transform = world_trans.Mul4(w.locs[0][0][0].trans)
    v := mgl32.Vec3{0, 0, 0}
    dist := v.Len() / 100
    w.cube.Color = mgl32.Vec3{1 - dist, 1, 1}
    c.Window().Finish()
  })

  gl.Main()
}
