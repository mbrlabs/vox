#version 330 

uniform mat4 u_mvp;

in vec3 a_pos;
in vec2 a_uvs;
in vec3 a_norm;

out vec2 fragTexCoord;

void main() {
    fragTexCoord = a_uvs;
    gl_Position = u_mvp * vec4(a_pos, 1.0);
}
