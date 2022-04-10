package ziface

type IServer interface { //定义一个服务器接口
	Start() //启动服务器
	Stop()  //停止服务器
	Serve() //运行服务器
}
