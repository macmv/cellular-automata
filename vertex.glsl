#version 430

uniform sampler3D tex;

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

layout (location = 0) in vec3 pos;
layout (location = 1) in vec2 uv;
layout (location = 2) in vec3 in_color;

out vec3 color;

bool test(int x, int y, int z) {
  vec4 tex_val = texture(tex, vec3(float(x) / 100, float(y) / 100, float(z) / 100));
  return tex_val.a > 0;
}

void main() {
  int x = gl_InstanceID % 100;
  int y = gl_InstanceID / 100 % 100;
  int z = gl_InstanceID / 10000;
  vec4 tex_val = texture(tex, vec3(float(x) / 100, float(y) / 100, float(z) / 100));
  float dist = sqrt(pow(x - 50, 2) + pow(y - 50, 2) + pow(z - 50, 2));
  color = tex_val.rgb;
  color.r = tex_val.a * 50 + 0.5;

  int num = 0;
  // ambient occlusion
  if (test(x - 1, y, z)) { num++; }
  if (test(x + 1, y, z)) { num++; }
  if (test(x, y - 1, z)) { num++; }
  if (test(x, y + 1, z)) { num++; }
  if (test(x, y, z - 1)) { num++; }
  if (test(x, y, z + 1)) { num++; }
  color -= num * 0.1;

  if (tex_val.a > 0) {
    gl_Position = projection * camera * model * vec4((vec3(x, y, z) - 50) * 2 + pos, 1);
  } else {
    gl_Position = vec4(0);
  }
}
