#version 330 

const float SCALE_FACTOR = 1.01;

uniform mat4 u_mvp;

in vec3 a_pos;

void main() {
    vec3 scaledVertex = a_pos * SCALE_FACTOR;
    gl_Position = u_mvp * vec4(scaledVertex, 1.0);
}
