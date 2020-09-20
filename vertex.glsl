#version 430

uniform sampler3D tex;

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
  vec4 tex_val = texture(tex, vec3(float(x) / 100, float(y) / 100, float(z) / 100));
  float dist = sqrt(pow(x - 50, 2) + pow(y - 50, 2) + pow(z - 50, 2));
  col = tex_val.rgb;
  col.g += dist * 0.03;
  col.b *= dist * 0.01;
  if (tex_val.a > 0.5) {
    gl_Position = projection * camera * model * vec4(vec3(x, y, z) + vert, 1);
  } else {
    gl_Position = vec4(0);
  }
}

