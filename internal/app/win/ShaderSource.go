package win

var (
	//顶点输入
	vertices = []float32{
		//-- 位置 --   ---- 颜色 ----  - 纹理坐标 -
		0.5, 0.5, 0.0, 1.0, 0.0, 0.0, 1.0, 1.0,
		0.5, -0.5, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0,
		-0.5, -0.5, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0,
		-0.5, 0.5, 0.0, 1.0, 1.0, 0.0, 0.0, 1.0,
	}
	// //索引缓冲对象
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
		layout (location = 0) in vec3 aPos;   // 位置变量的属性位置值为 0 
		layout (location = 1) in vec3 aColor; // 颜色变量的属性位置值为 1
		layout (location = 2) in vec2 aTexCoord;

		out vec3 vertexColor;
		out vec2 TexCoord;

		void main()
		{
			gl_Position = vec4(aPos, 1.0);
			vertexColor = aColor;
			TexCoord = aTexCoord;
		}	
	` + "\x00"
	//片段着色器
	fragmentShaderSource = `
		#version 330 core
		out vec4 fragmentColor;
		in vec3 vertexColor; // 从顶点着色器传来的输入变量（名称相同、类型相同）
		in vec2 TexCoord;

		uniform sampler2D texture1;
		uniform sampler2D texture2;

		void main()
		{
			fragmentColor = mix(texture(texture1, TexCoord), texture(texture2, TexCoord), 0.37) * vec4(vertexColor, 1.0);
		} 
	` + "\x00"
)
