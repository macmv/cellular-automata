#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;
uniform vec3 color;

in vec3 vert;
in vec2 uv;
in vec3 normal;
out vec3 col;

void main() {
  col = color;
  col.y += vert.y * 0.2;
  gl_Position = projection * camera * model * vec4(vert, 1);
}

