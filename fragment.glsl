#version 330

uniform sampler2D tex;

in vec3 pos;
out vec4 outputColor;

void main() {
  outputColor = vec4(pos, 1);
}
