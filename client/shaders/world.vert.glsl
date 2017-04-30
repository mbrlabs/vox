#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 a_pos;
in vec3 a_norm;
in vec2 a_uvs;

out vec2 fragTexCoord;

void main() {
    fragTexCoord = a_uvs;
    gl_Position = projection * camera * model * vec4(a_pos, 1);
}