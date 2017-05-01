#version 330

uniform sampler2D tex;

in vec2 fragTexCoord;
out vec4 outColor;

void main() {
    outColor = vec4(1.0, 1.0, 1.0, 1.0);
}
