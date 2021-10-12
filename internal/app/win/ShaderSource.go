package win

var (
	//顶点输入
	vertices = []float32{
		0.5, 0.5, 0.0,
		0.5, -0.5, 0.0,
		-0.5, -0.5, 0.0,
		-0.5, 0.5, 0.0,
	}
	//索引缓冲对象
	indices = []int32{
		// 注意索引从0开始!
		0, 1, 3, // 第一个三角形
		1, 2, 3, // 第二个三角形
	}
)

const (
	//顶点着色器
	vertexShaderSource = `
		#version 330 core
		layout (location = 0) in vec3 aPos;
		void main()
		{
			gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
		}	
	` + "\x00"
	//片段着色器
	fragmentShaderSource = `
		#version 330 core
		out vec4 FragColor;
		void main()
		{
			FragColor = vec4(1.0f, 1.0f, 1.0f, 1.0f);
		} 
	` + "\x00"
)
