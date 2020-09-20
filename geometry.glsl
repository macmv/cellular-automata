#version 430
layout (triangles) in;
layout (triangle_strip, max_vertices = 3) out;

uniform sampler3D tex;

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 color[];
out vec3 norm;
out vec3 pass_color;
out vec3 pass_pos;

void set_pos(int i) {
  gl_Position = gl_in[i].gl_Position;
  pass_color = color[i];
  pass_pos = (model * gl_in[i].gl_Position).xyz;
  EmitVertex();
}

void main() {
  vec3 a = (model * gl_in[0].gl_Position).xyz;
  vec3 b = (model * gl_in[1].gl_Position).xyz;
  vec3 c = (model * gl_in[2].gl_Position).xyz;

  // pass normal to vertex shader
  norm = normalize(cross(b - a, c - a));

  set_pos(0);
  set_pos(1);
  set_pos(2);
  EndPrimitive();
}

