#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 vert;
out vec3 pos;

void main() {
  pos = vert;
  gl_Position = projection * camera * model * vec4(vert, 1);
}

