#version 330

uniform sampler2D tex;

in vec2 fragTexCoord;
out vec4 outColor;

void main() {
    outColor = vec4(0, 0.8, 0.9, 1.0);
}
