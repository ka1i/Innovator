package win

var (
	//顶点输入
	vertices = []float32{
		//-- 位置- -纹理坐标-
		1, 1, 0.0, 1.0, 0.0,
		1, -1, 0.0, 1.0, 1.0,
		-1, -1, 0.0, 0.0, 1.0,
		-1, 1, 0.0, 0.0, 0.0,
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
		layout (location = 0) in vec3 aPos;
		layout (location = 1) in vec2 aTexCoord;

		out vec2 uv;

		void main()
		{
			gl_Position = vec4(aPos, 1.0);
			uv = aTexCoord;
		}	
	` + "\x00"
	//片段着色器
	fragmentShaderSource = `
		#version 330 core
		out vec4 fragmentColor;
		in vec2 uv;

		uniform sampler2D texture1;

		out vec4 base;
		uniform vec4 background;

		void main()
		{
			base = texture(texture1, uv);
			fragmentColor = vec4(mix(background.xyz, base.rgb, max(background.w, base.a)), 1.0);
		} 
	` + "\x00"
)
