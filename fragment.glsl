#version 330

in vec3 col;
out vec4 outputColor;

void main() {
  outputColor = vec4(col, 1);
}
