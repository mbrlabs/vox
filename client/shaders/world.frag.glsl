#version 330

uniform vec3 u_color;

in vec2 fragTexCoord;
out vec4 outColor;

void main() {
    outColor = vec4(u_color, 1.0);
}
