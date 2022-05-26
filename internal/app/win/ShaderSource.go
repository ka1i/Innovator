package win

var (
	//顶点输入
	vertices = []float32{
		// pos    // tex
		0.0, 1.0, 0.0, 1.0,
		1.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 1.0,
		1.0, 1.0, 1.0, 1.0,
		1.0, 0.0, 1.0, 0.0,
	}
)

const (
	//顶点着色器
	vertexShaderSource = `
		#version 330 core
		layout (location = 0) in vec4 vertex;

		out vec2 TexCoords;

		uniform mat4 projection;
		uniform mat4 model;

		void main()
		{
			TexCoords = vertex.zw;
			gl_Position = projection * model * vec4(vertex.xy, 0.0, 1.0);
		}	
	` + "\x00"
	//片段着色器
	fragmentShaderSource = `
		#version 330 core
		out vec4 fragmentColor;
		in vec2 TexCoords;

		uniform sampler2D imgtexture;

		void main()
		{
			fragmentColor = texture(imgtexture, TexCoords);
		} 
	` + "\x00"
)
