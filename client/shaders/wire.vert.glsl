#version 330 

uniform mat4 u_mvp;

in vec3 a_pos;
in vec3 a_norm;

void main() {
    gl_Position = u_mvp * vec4(a_pos, 1.0);
}
