#version 320

precision lowp float;
precision lowp sampler3D;

uniform sampler3D tex;

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 norm;
in vec3 pass_color;
in vec3 pass_pos;

out vec4 output_color;

void main() {
  vec3 light_pos = (model * vec4(10, 10, -10, 1)).xyz;
  float specular_strength = 0.5;
  float diffuse_strength = 1.0;

  // diffuse lighting
  vec3 to_light_vec = normalize(pass_pos - light_pos);
  float brightness = max(dot(norm, to_light_vec), 0.0);
  float diffuse = brightness * diffuse_strength;

  // final color
  output_color = vec4(diffuse * pass_color + vec3(0.1), 1);
  output_color = vec4(1, 1, 0, 1);
}
