#version 330

uniform sampler2D tex;

in vec3 col;
out vec4 outputColor;

void main() {
  outputColor = vec4(col, 1);
}
