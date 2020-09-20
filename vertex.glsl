#version 430

uniform sampler2D tex;

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;
uniform vec3 color;

in vec3 vert;
in vec2 uv;
in vec3 normal;
out vec3 col;

void main() {
  int x = gl_InstanceID % 100;
  int y = gl_InstanceID / 100 % 100;
  int z = gl_InstanceID / 10000;
  vec4 tex_val = texture(tex, vec2(float(x) / 100, float(y) / 100));
  col = tex_val.rgb;
  if (tex_val.r > 0.5) {
    gl_Position = projection * camera * model * vec4(vec3(x, y, z) + vert, 1);
  } else {
    gl_Position = vec4(0);
  }
}

