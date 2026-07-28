package gotool

import (
	"github.com/druidcaesa/gotool/array"
	"github.com/druidcaesa/gotool/bcrypt"
	"github.com/druidcaesa/gotool/captcha"
	"github.com/druidcaesa/gotool/compression"
	"github.com/druidcaesa/gotool/convert"
	"github.com/druidcaesa/gotool/date"
	"github.com/druidcaesa/gotool/desensitized"
	"github.com/druidcaesa/gotool/id_utils"
	"github.com/druidcaesa/gotool/logs"
	"github.com/druidcaesa/gotool/openfile"
	"github.com/druidcaesa/gotool/page"
	"github.com/druidcaesa/gotool/pretty"
	"github.com/druidcaesa/gotool/request"
	"github.com/druidcaesa/gotool/str"
	"github.com/druidcaesa/gotool/tree"
)

// @Title tool
// @Description A simple, semantic and developer-friendly Golang tool development set
// @Page github.com/druidcaesa/gotool
// @Version v0.0.1
// @Author druidcaesa
// @Email hbsjzfynxm@gmail.com

var (
	StrArrayUtils     array.StrArray            //String 数据工具声明
	Logs              logs.Logs                 //log日志声明
	BcryptUtils       bcrypt.BcryptUtil         //加密工具声明
	DateUtil          date.Date                 //时间操作
	StrUtils          str.StrUtils              //字符串操作
	HttpUtils         request.Request           //http工具
	ConvertUtils      convert.Convert           //公历转农历相关操作
	PageUtils         page.Page                 //分页插件
	IdUtils           id_utils.IdUtils          //id生成工具
	ZipUtils          compression.ZipUtils      //压缩和解压缩工具
	FileUtils         openfile.FileUtils        //IO操作工具
	CaptchaUtils      captcha.Captcha           //验证码工具
	DesensitizedUtils desensitized.Desensitized //敏感数据脱敏
	TreeUtils         tree.Tree                 //菜单树结构化工具
	PrettyUtils       pretty.PrettyUtils        //JSON打印格式化工具
)
