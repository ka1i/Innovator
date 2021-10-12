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
		out vec4 vertexColor;

		void main()
		{
			gl_Position = vec4(aPos, 1.0);
			vertexColor = vec4(1.0f, 1.0f, 1.0f, 1.0f);
		}	
	` + "\x00"
	//片段着色器
	fragmentShaderSource = `
		#version 330 core
		out vec4 fragmentColor;
		in vec4 vertexColor; // 从顶点着色器传来的输入变量（名称相同、类型相同）

		uniform vec4 ourColor;

		void main()
		{
			fragmentColor = ourColor; // vertexColor;
		} 
	` + "\x00"
)
