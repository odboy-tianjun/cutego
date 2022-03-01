package websocket

import (
	"cutego/pkg/common"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// 连接例子
// <script>
// var ws = new WebSocket("ws://127.0.0.1:21366/websocket?user=admin&code=notice");
// // 连接打开时触发
// ws.onopen = function(evt) {
// console.log("Connection open ...");
// ws.send("Hello WebSockets!");
// };
// // 接收到消息时触发
// ws.onmessage = function(evt) {
// console.log("Received Message: " + evt.data);
// };
// // 连接关闭时触发
// ws.onclose = function(evt) {
// console.log("Connection closed.");
// };
// </script>

// 定义回调接口(消息类型、内容)
type OnReceiveMessage func(messageType int, content []byte) error

// 分割符号
const SignalSplitSymbol = "=_="

// 用户名 <--> websocket通道
var OnlineUserMap = make(map[string]*websocket.Conn)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// webSocket请求ping 返回pong
func HandleWebSocketMessage(c *gin.Context) {
	userValue, userExist := c.GetQuery("user")
	codeValue, codeExist := c.GetQuery("code")

	if !userExist && !codeExist {
		return
	}

	// 升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	// defer ws.Close()
	cacheKey := userValue + SignalSplitSymbol + codeValue
	// 如果存在则踢掉之前的通道
	if OnlineUserMap[cacheKey] != nil {
		// 函数结束后关闭
		tmpConn := OnlineUserMap[cacheKey]
		defer tmpConn.Close()
	}
	OnlineUserMap[cacheKey] = ws
	// 回收监听消息
	go ListenWebSocketMessage(userValue, codeValue, HandleAdminNotice)
}

// 回调函数的具体实现
func HandleAdminNotice(messageType int, content []byte) error {
	common.InfoLog("messageType=%d\n", messageType)
	common.InfoLog("content=%s\n", string(content))
	return nil
}

// 监听到消息
// onReceiveMessage为函数提供的回调接口, 让外部去实现
func ListenWebSocketMessage(user string, code string, onReceiveMessage OnReceiveMessage) {
	cacheKey := user + SignalSplitSymbol + code
	ws := OnlineUserMap[cacheKey]
	if ws != nil {
		for {
			// 读取ws中的数据
			mt, message, err := ws.ReadMessage()
			if err != nil {
				break
			}
			// 处理心跳
			if string(message) == "ping" {
				message = []byte("pong")
				// 写入ws数据
				err = ws.WriteMessage(mt, message)
				if err != nil {
					break
				}
				continue
			}
			// 处理接收到的数据
			err = onReceiveMessage(mt, message)
			if err != nil {
				break
			}
		}
	}
}
