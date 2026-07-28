gotool
=======
gotool是一个小而全的Golang工具集，主要是将日常开发中常用的到方法进行提炼集成，避免重复造轮子，提高工作效率，每一个方法都是作者经过工作经验，和从以往的项目中提炼出来的。
如果您觉得这个工具好用，请给个star谢谢

### 简体中文 [English](README.md)
# 请关注
- [https://github.com/druidcaesa/gotool](https://github.com/druidcaesa/gotool)
- [https://gitee.com/termites/gotool](https://gitee.com/termites/gotool)

## 2021-7-9更新内容详细使用请看文档

- 添加文件IO操作工具FileUtils
- 添加验证码生成工具CaptchaUtils
- 添加文件目录压缩和解压缩工具ZipUtis
- 字符串数组工具StrArrayUtils

## 2021-7-10更新内容详细使用请看文档
 - 添加敏感消息脱敏工具DesensitizedUtils 
 - 添加菜单树数据格式化工具TreeUtils 
 - 添加JSON输出美化工具PrettyUtils

### 如何使用gotool呢？

### 安装

go get github.com/druidcaesa/gotool

go.mod github.com/druidcaesa/gotool

### 引入

```go
import "github.com/druidcaesa/gotool"
```

StrUtils
=======
golang一个string常用工具集，基本涵盖了开发中经常用到的工具，目前正在不端的完善中

#### 1、gotool.StrUtils.ReplacePlaceholder 占位符替换

```go
func TestStringReplacePlaceholder(t *testing.T) {
s := "你是我的{},我是你的{}"
placeholder, err := gotool.StrUtils.ReplacePlaceholder(s, "唯一", "所有")
if err == nil {
fmt.Println(placeholder)
}
}
//out
== = RUN   TestStringReplacePlaceholder
你是我的唯一, 我是你的所有
--- PASS: TestStringReplacePlaceholder (0.00s)
PASS
```

#### 2、gotool.StrUtils.RemoveSuffix 去除文件扩展名获取文件名

```go
func TestRemoveSuffix(t *testing.T) {
fullFilename := "test.txt"
suffix, _ := gotool.StrUtils.RemoveSuffix(fullFilename)
fmt.Println(suffix)
fullFilename = "/root/home/test.txt"
suffix, _ = gotool.StrUtils.RemoveSuffix(fullFilename)
fmt.Println(suffix)
}
//out
== = RUN   TestRemoveSuffix
test
test
--- PASS: TestRemoveSuffix (0.00s)
PASS
```

#### 3、gotool.StrUtils.GetSuffix 获取文件扩展名

```go
func TestGetSuffix(t *testing.T) {
fullFilename := "test.txt"
suffix, _ := gotool.StrUtils.GetSuffix(fullFilename)
fmt.Println(suffix)
fullFilename = "/root/home/test.txt"
suffix, _ = gotool.StrUtils.GetSuffix(fullFilename)
fmt.Println(suffix)
}
//out
== = RUN   TestGetSuffix
.txt
.txt
--- PASS: TestGetSuffix (0.00s)
PASS
```

#### 4、gotool.StrUtils.HasEmpty 判断字符串是否未空，我空返回ture

```go
func TestHasStr(t *testing.T) {
str := ""
empty := gotool.StrUtils.HasEmpty(str)
fmt.Println(empty)
str = "11111"
empty = gotool.StrUtils.HasEmpty(str)
fmt.Println(empty)
}
//out
== = RUN   TestHasStr
true
false
--- PASS: TestHasStr (0.00s)
PASS

```

StrArrayUtils string数组操作工具
======

#### 1、gotool.StrArrayUtils.StringToInt64 字符串数组转int64数组，调用前请确保字符串数组均为数字

```go
func TestStringToInt64(t *testing.T) {
//字符串数组转int64
strings := []string{"1", "23123", "232323"}
fmt.Println(reflect.TypeOf(strings[0]))
toInt64, err := gotool.StrArrayUtils.StringToInt64(strings)
if err != nil {
t.Fatal(err)
}
fmt.Println(reflect.TypeOf(toInt64[0]))
}
//out
== = RUN   TestStringToInt64
string
int64
--- PASS: TestStringToInt64 (0.00s)
PASS
```

#### 2、gotool.StrArrayUtils.StringToInt32 字符串数组转int64数组，调用前请确保字符串数组均为数字

```go
func TestStringToInt32(t *testing.T) {
//字符串数组转int64
strings := []string{"1", "23123", "232323"}
fmt.Println(reflect.TypeOf(strings[0]))
toInt64, err := gotool.StrArrayUtils.StringToInt32(strings)
if err != nil {
t.Fatal(err)
}
fmt.Println(reflect.TypeOf(toInt64[0]))
}
//out
== = RUN   TestStringToInt32
string
int32
--- PASS: TestStringToInt32 (0.00s)
PASS
```

#### 3、gotool.StrArrayUtils.ArrayDuplication 数组去重

```go
func TestArrayDuplication(t *testing.T) {
//string数组去重
strings := []string{"hello", "word", "gotool", "word"}
fmt.Println("去重前----------------->", strings)
duplication := gotool.StrArrayUtils.ArrayDuplication(strings)
fmt.Println("去重后----------------->", duplication)
}
//out
== = RUN   TestArrayDuplication
去重前-----------------> [hello word gotool word]
去重后-----------------> [hello word gotool]
--- PASS: TestArrayDuplication (0.00s)
PASS
```

DateUtil
=======
golang一个时间操作工具集，基本涵盖了开发中经常用到的工具，目前正在不端的完善中

#### 1、gotool.DateUtil.FormatToString 时间格式化成字符串

```go
func TestFormatToString(t *testing.T) {
now := gotool.DateUtil.Now()
toString := gotool.DateUtil.FormatToString(&now, "YYYY-MM-DD hh:mm:ss")
fmt.Println(toString)
toString = gotool.DateUtil.FormatToString(&now, "YYYYMMDD hhmmss")
fmt.Println(toString)
}
//年月日对应YYYY MM DD 时分秒 hhmmss 可进行任意组合 比如 YYYY  hh   YYYY-DD hh:mm 等
//out
== = RUN   TestFormatToString
2021-07-07 16:13:30
20210707 161330
--- PASS: TestFormatToString (0.00s)
PASS
```

#### 2、gotool.DateUtil.IsZero 判断时间是否为空

```go
//时间为空 true 否则 false
func TestDate_IsZero(t *testing.T) {
t2 := time.Time{}
zero := gotool.DateUtil.IsZero(t2)
fmt.Println(zero)
zero = gotool.DateUtil.IsZero(gotool.DateUtil.Now())
fmt.Println(zero)
}
//out
== = RUN   TestDate_IsZero
true
false
--- PASS: TestDate_IsZero (0.00s)
PASS
```

#### 3、gotool.DateUtil.Now 获取当前时间 等同于time.Now(),为了统一化所以将此方法也纳入到工具中

#### 4、gotool.DateUtil.InterpretStringToTimestamp 字符串格式化成时间类型

```go
//参数一 需要格式化的时间字符串 参数二 字符串格式，需要和需格式化字符串格式一致 
//如 2021-6-4 对应YYYY-MM-DD  2021.6.4 对应YYYY.MM.DD
func TestInterpretStringToTimestamp(t *testing.T) {
timestamp, err := gotool.DateUtil.InterpretStringToTimestamp("2021-05-04 15:12:59", "YYYY-MM-DD hh:mm:ss")
if err != nil {
gotool.Logs.ErrorLog().Println(err.Error())
}
fmt.Println(timestamp)
}
//out
== = RUN   TestInterpretStringToTimestamp
1620112379
--- PASS: TestInterpretStringToTimestamp (0.00s)
PASS
```

#### 5、gotool.DateUtil.UnixToTime 时间戳转时间

```go
func TestUnixToTime(t *testing.T) {
unix := gotool.DateUtil.Now().Unix()
fmt.Println("时间戳----------------------->", unix)
toTime := gotool.DateUtil.UnixToTime(unix)
fmt.Println(toTime)
}
//out
== = RUN   TestUnixToTime
时间戳-----------------------> 1625645682
2021-07-07 16:14:42 +0800 CST
--- PASS: TestUnixToTime (0.00s)
PASS
```

#### 6、gotool.DateUtil.GetWeekDay 获取星期几

```go
func TestGetWeekDay(t *testing.T) {
now := gotool.DateUtil.Now()
day := gotool.DateUtil.GetWeekDay(now)
fmt.Println("今天是-----------------周", day)
}
//out
== = RUN   TestGetWeekDay
今天是-----------------周 3
--- PASS: TestGetWeekDay (0.00s)
PASS
```

#### 7、gotool.DateUtil.MinuteAddOrSub,HourAddOrSub,DayAddOrSub 时间计算工具

```go
//时间计算
func TestTimeAddOrSub(t *testing.T) {
now := gotool.DateUtil.Now()
fmt.Println("现在时间是--------------------->", now)
sub := gotool.DateUtil.MinuteAddOrSub(now, 10)
fmt.Println("分钟计算结果-------------------->", sub)
sub = gotool.DateUtil.MinuteAddOrSub(now, -10)
fmt.Println("分钟计算结果-------------------->", sub)
sub = gotool.DateUtil.HourAddOrSub(now, 10)
fmt.Println("小时计算结果-------------------->", sub)
sub = gotool.DateUtil.HourAddOrSub(now, -10)
fmt.Println("小时计算结果-------------------->", sub)
sub = gotool.DateUtil.DayAddOrSub(now, 10)
fmt.Println("天计算结果-------------------->", sub)
sub = gotool.DateUtil.DayAddOrSub(now, -10)
fmt.Println("天计算结果-------------------->", sub)
}
//现在时间是---------------------> 2021-07-07 11:18:17.8295757 +0800 CST m=+0.012278001
//分钟计算结果--------------------> 2021-07-07 11:28:17.8295757 +0800 CST m=+600.012278001
//分钟计算结果--------------------> 2021-07-07 11:08:17.8295757 +0800 CST m=-599.987721999
//小时计算结果--------------------> 2021-07-07 21:18:17.8295757 +0800 CST m=+36000.012278001
//小时计算结果--------------------> 2021-07-07 01:18:17.8295757 +0800 CST m=-35999.987721999
//天计算结果--------------------> 2021-07-17 11:18:17.8295757 +0800 CST m=+864000.012278001
//天计算结果--------------------> 2021-06-27 11:18:17.8295757 +0800 CST m=-863999.987721999
```

ConvertUtils 公历转农历工具
============

#### 1、gotool.ConvertUtils.GregorianToLunarCalendar(公历转农历),GetLunarYearDays(农历转公历),GetLunarYearDays(获取农历这一年农历天数)

```go
func TestConvertTest(t *testing.T) {
calendar := gotool.ConvertUtils.GregorianToLunarCalendar(2020, 2, 1)
fmt.Println(calendar)
gregorian := gotool.ConvertUtils.LunarToGregorian(calendar[0], calendar[1], calendar[2], false)
fmt.Println(gregorian)
days := gotool.ConvertUtils.GetLunarYearDays(2021)
fmt.Println(days)
}
//[2020 1 8]
//[2020 2 1]
//354
```

#### 2、gotool.ConvertUtils.GetSolarMonthDays(2021,7)获取公历某月天数 2021年7月天数

#### 3、gotool.ConvertUtils.IsLeapYear(2021)获取某年是否是瑞年 true是 false不是

#### 4、gotool.ConvertUtils.GetLeapMonth(2021)获取某年闰月月份

BcryptUtils 加密和解密工具
==========

#### 1、gotool.BcryptUtils.Generate 加密处理，多用于密码进行加密后数据库存储使用，不可逆

#### 2、gotool.BcryptUtils.CompareHash 加密后和未加密密码对比，多用于登录验证使用

```go
func TestGenerate(t *testing.T) {
//进行加密
generate := gotool.BcryptUtils.Generate("123456789")
fmt.Println(generate)
//进行加密对比
hash := gotool.BcryptUtils.CompareHash(generate, "123456789")
fmt.Println(hash)
}
//out
== = RUN   TestGenerate
$2a$10$IACJ6zGuNuzaumrvDz58Q.vJUzz4JGqougYKrdCs48rQYIRjAXcU2
true
--- PASS: TestGenerate (0.11s)
PASS
```

#### 3、gotool.BcryptUtils.MD5 md5加密

```go
func TestMd5(t *testing.T) {
md5 := gotool.BcryptUtils.MD5("123456789")
fmt.Println(md5)
}
//out
== = RUN   TestMd5
25f9e794323b453885f5181f1b624d0b
--- PASS: TestMd5 (0.00s)
PASS
```

#### 4、gotool.BcryptUtils.GenRsaKey(获取公钥和私钥),

#### 5、RsaSignWithSha256(进行签名),

#### 6、RsaVerySignWithSha256(验证),

#### 7、RsaEncrypt(公钥加密),

#### 8、RsaDecrypt(私钥解密)

```go
func TestRsa(t *testing.T) {
//rsa 密钥文件产生
fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
prvKey, pubKey := gotool.BcryptUtils.GenRsaKey()
fmt.Println(string(prvKey))
fmt.Println(string(pubKey))

fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
var data = "我是被加密的数据记住我哦-------------------------------"
fmt.Println("对消息进行签名操作...")
signData := gotool.BcryptUtils.RsaSignWithSha256([]byte(data), prvKey)
fmt.Println("消息的签名信息： ", hex.EncodeToString(signData))
fmt.Println("\n对签名信息进行验证...")
if gotool.BcryptUtils.RsaVerySignWithSha256([]byte(data), signData, pubKey) {
fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
}

fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
ciphertext := gotool.BcryptUtils.RsaEncrypt([]byte(data), pubKey)
fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
sourceData := gotool.BcryptUtils.RsaDecrypt(ciphertext, prvKey)
fmt.Println("私钥解密后的数据：", string(sourceData))
}
//out
== = RUN   TestRsa
-------------------------------获取RSA公私钥-----------------------------------------
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCgHh1ZYFjlxrIJYjjWGFaLwC8Oov8KqyMtHa+GauF121dperr3
i46JyDoskoSBhbkmqv70LMNrjqVdttdIsC0BtH9ThWLBwKVLH56EqfzqlzClKZEh
WTNaddCSuxoZpN33mwS82DCjZe3d7nAPdEGD5pSBx6TVt5bG1c3NVAmcBQIDAQAB
AoGAWc5KO9T0R3xYYzb6Fer0r9GNEzKMxdkTE7zws/3Cky4BKyIxN6LIwbLSHinX
tCXioTOLaDyrJuqNCbEBsr1NoCIPxnswA5Jm5QDYO5x9aq4u8+SZ9KqLbMrO1JDS
ZV7Cbiblz1xNMfdVIvlVjD5RdEotbYTbHI2CZUoPsjHng8kCQQDHi6TJYJqvej8r
q46ZceuWHDgE81Wu16RrA/kZKi6MJAApQtfO/4HM6W/ImbCjZ2rSYxqnAlIg/GxA
dT6iJABjAkEAzWra06RyhGm3lk9j9Xxc0YPX6VX7qT5Iq6c7ry1JtTSeVOksKANG
elaNnCj8YYUgj7BeBBcMMvLd39hP1h06dwJAINTyHQwfB2ZGxImqocajq4QjF3Vu
EKF8dPsnXiOZmwdFW4Sa+30Av1VdRhU7gfc/FTSnKvlvx+ugaA6iao0f3wJBALa8
sTCH4WwUE8K+m4DeAkBMVn34BKnZg5JYcgrzcdemmJeW2rY5u6/HYbCi8WnboUzS
K8Dds/d7AJBKgTNLyx8CQBAeU0St3Vk6SJ6H71J8YtVxlRGDjS2fE0JfUBrpI/bg
r/QI8yMAMyFkt1YshN+UNWJcvR5SXQnyT/udnWJIdg4 =
-----END RSA PRIVATE KEY-----

-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCgHh1ZYFjlxrIJYjjWGFaLwC8O
ov8KqyMtHa+GauF121dperr3i46JyDoskoSBhbkmqv70LMNrjqVdttdIsC0BtH9T
hWLBwKVLH56EqfzqlzClKZEhWTNaddCSuxoZpN33mwS82DCjZe3d7nAPdEGD5pSB
x6TVt5bG1c3NVAmcBQIDAQAB
-----END PUBLIC KEY-----

-------------------------------进行签名与验证操作-----------------------------------------
对消息进行签名操作...
消息的签名信息：  1fcf20c4fb548c8fc0906e369287feb84c861bf488d822d78a0eada23d1af66ed3a12e9440d7181b1748fd0ad805222cf2ce7ce1f6772b330ef11b717700ba26945dda9d749a5c4d8c108ede103c17bed92801a4c3fbc1ebf38d10bf4bf183713eeb7f429acc3dcc3812366a324737f756720f3f9e75ddda1e024a7698b89163

对签名信息进行验证...
签名信息验证成功，确定是正确私钥签名！！
-------------------------------进行加密解密操作-----------------------------------------
公钥加密后的数据： 637b05798c1cf95cfcc63adf228645c77f8e9a9f34b12b722e6938ded8c9560a0430171a4f940d3fb2d96bc6c470c80a817d81f4a2669b991adbff5b22b335129e514c921083ce3e64c1c876409d9b763d5e8e269797283ef951a364da6a59a1d8454391356cb0cd0808092e9dd7ac371f9247a43760f3c82b7ad26a32a7a807
私钥解密后的数据： 我是被加密的数据记住我哦-------------------------------
--- PASS: TestRsa (0.02s)
PASS
```

ZipUtils
=======
压缩和解压工具，可单文件压缩也可进行目录压缩，或者跨目录压缩

#### 1、gotool.ZipUtils.Compress

- files 文件数组 可以是多目录文件
- dest 压缩文件存放地址

```go
func TestCompress(t *testing.T) {
open, err := os.Open("/Users/fanyanan/Downloads/gotool")
if err != nil {
t.Fatal(err)
}
files := []*os.File{open}
flag, err := gotool.ZipUtils.Compress(files, "/Users/fanyanan/Downloads/test.zip")
if err != nil {
t.Fatal(err)
}
if flag {
fmt.Println("压缩成功")
}
}
//out
== = RUN   TestCompress
压缩成功
--- PASS: TestCompress (0.12s)
PASS
```

#### 2、gotool.ZipUtils.DeCompress

- zipFile 压缩包路径
- dest 要解压的路径

```go
func TestDeCompress(t *testing.T) {
compress, err := gotool.ZipUtils.DeCompress("/Users/fanyanan/Downloads/test.zip", "/Users/fanyanan/Downloads")
if err != nil {
t.Fatal(err)
}
if compress {
fmt.Println("解压成功")
}
}
//out
== = RUN   TestDeCompress
解压成功
--- PASS: TestDeCompress (0.44s)
PASS
```

FileUtils
=======

文件操作工具，让做io操作更简单各个方便,让io操作不是那么难

#### 1、gotool.FileUtils.Exists 判断文件或者目录是否存在

```go
func TestFileExists(t *testing.T) {
//判断文件或目录是否存在
exists := gotool.FileUtils.Exists("F:/go-test/test")
fmt.Println("创建前------------------------>", exists)
err := os.MkdirAll("F:/go-test/test", os.ModePerm)
if err != nil {
t.Fatal(err)
}
exists = gotool.FileUtils.Exists("F:/go-test/test")
fmt.Println("创建后------------------------>", exists)
}
//out
== = RUN   TestFileExists
创建前------------------------> false
创建后------------------------> true
--- PASS: TestFileExists (0.00s)
PASS
```

#### 2、gotool.FileUtils.IsDir 判断是否是目录

```go
func TestIsDir(t *testing.T) {
//判断是否是目录
dir := gotool.FileUtils.IsDir("F:/go-test/test")
fmt.Println("是否是目录--------------------->", dir)
dir = gotool.FileUtils.IsDir("F:/go-test/test/test.txt")
fmt.Println("是否是目录--------------------->", dir)
}
//out
== = RUN   TestIsDir
是否是目录---------------------> true
是否是目录---------------------> false
--- PASS: TestIsDir (0.00s)
PASS
```

#### 3、gotool.FileUtils.IsFile 判断是否是文件

```go
func TestIsFile(t *testing.T) {
//判断是否是文件
file := gotool.FileUtils.IsFile("F:/go-test/test")
fmt.Println("是否是文件--------------------->", file)
file = gotool.FileUtils.IsFile("F:/go-test/test/test.txt")
fmt.Println("是否是文件--------------------->", file)
}
//out
== = RUN   TestIsFile
是否是文件---------------------> false
是否是文件---------------------> true
--- PASS: TestIsFile (0.00s)
PASS
```

#### 4、gotool.FileUtils.RemoveFile 删除文件或目录

```go
func TestRemove(t *testing.T) {
//删除文件
file, err := gotool.FileUtils.RemoveFile("F:/go-test/test/test.txt")
if err != nil {
t.Fatal(err)
}
if file {
//查看文件是否还存在
exists := gotool.FileUtils.Exists("F:/go-test/test/test.txt")
fmt.Println("文件是否存在------------------------>", exists)
}
}
//out
== = RUN   TestRemove
文件是否存在------------------------> false
--- PASS: TestRemove (0.00s)
PASS
```

#### 5、gotool.FileUtils.OpenFileWronly 只写模式打开文件，没有将自动创建，写入内从进行测试

```go
func TestOpenFileWronly(t *testing.T) {
//用只写模式打开一个文件，并且写入5条内容,若文件不存在将会创建一个
path := "F:/go-test/test/test.txt"
str := "hello word gotool \n"
wronly, err := gotool.FileUtils.OpenFileWronly(path)
if err != nil {
t.Fatal(err)
}
defer wronly.Close()
write := bufio.NewWriter(wronly)
for i := 0; i < 5; i++ {
write.WriteString(str)
}
//Flush将缓存的文件真正写入到文件中
write.Flush()
//读取文件写入到控制台
files, err := gotool.FileUtils.OpenFileRdonly(path)
if err != nil {
t.Fatal(err)
}
defer files.Close()
reader := bufio.NewReader(files)
for {
str, err := reader.ReadString('\n')
if err != nil {
break
}
fmt.Print(str)
}
}
//out
== = RUN   TestOpenFileWronly
hello word gotool
hello word gotool
hello word gotool
hello word gotool
hello word gotool
--- PASS: TestOpenFileWronly (0.00s)
PASS
```

#### 6、gotool.FileUtils.OpenFileRdonly 只读模式打开文件，读取内容输出到控制台

```go
func TestOpenFileRdonly(t *testing.T) {
path := "F:/go-test/test/test.txt"
files, err := gotool.FileUtils.OpenFileRdonly(path)
if err != nil {
t.Fatal(err)
}
defer files.Close()
reader := bufio.NewReader(files)
for {
str, err := reader.ReadString('\n')
if err != nil {
break
}
fmt.Print(str)
}
}
//out
== = RUN   TestOpenFileRdonly
hello word gotool
hello word gotool
hello word gotool
hello word gotool
hello word gotool
--- PASS: TestOpenFileRdonly (0.00s)
PASS
```

#### 7、gotool.FileUtils.OpenFileAppend 打开文件在文件后面追加内容，没有将自动创建文件

```go
func TestOpenFileAppend(t *testing.T) {
//打开文件在文件后面追加数据
path := "F:/go-test/test/test.txt"
str := "追加内容 \n"
wronly, err := gotool.FileUtils.OpenFileAppend(path)
if err != nil {
t.Fatal(err)
}
defer wronly.Close()
write := bufio.NewWriter(wronly)
for i := 0; i < 5; i++ {
write.WriteString(str)
}
//Flush将缓存的文件真正写入到文件中
write.Flush()
//读取文件写入到控制台
files, err := gotool.FileUtils.OpenFileRdonly(path)
if err != nil {
t.Fatal(err)
}
defer files.Close()
reader := bufio.NewReader(files)
for {
str, err := reader.ReadString('\n')
if err != nil {
break
}
fmt.Print(str)
}
}
//out
== = RUN   TestOpenFileAppend
hello word gotool
hello word gotool
hello word gotool
hello word gotool
hello word gotool
追加内容
追加内容
追加内容
追加内容
追加内容
--- PASS: TestOpenFileAppend (0.00s)
PASS
```

#### 8、gotool.FileUtils.FileCopy 文件复制方法

```go
func TestFileCopy(t *testing.T) {
//文件复制功能
path := "F:/go-test/test/test.txt"
copyPath := "F:/go-test/test/test.txt1"
//复制钱
exists := gotool.FileUtils.Exists(copyPath)
fmt.Println("复制前文件是否存在------------------>", exists)
//复制后
fileCopy, err := gotool.FileUtils.FileCopy(path, copyPath)
if err != nil {
t.Fatal(err)
}
if fileCopy {
exists := gotool.FileUtils.Exists(copyPath)
fmt.Println("复制前文件是否存在------------------>", exists)
}
}
//out
== = RUN   TestFileCopy
复制前文件是否存在------------------> false
复制前文件是否存在------------------> true
--- PASS: TestFileCopy (0.00s)
PASS
```

CaptchaUtils 验证码生成工具
======

#### 1、gotool.CaptchaUtils.GetRandStr 生成验证码strng串

```go
func TestCaptcha(t *testing.T) {
//生成验证码字符串，可以进行redis等存储用于验证逻辑
str := gotool.CaptchaUtils.GetRandStr(6)
fmt.Println(str)
}
//out
== = RUN   TestCaptcha
生成验证码字符串-------------------> qK5DME
--- PASS: TestCaptcha (0.01s)
PASS
```

#### 2、gotool.CaptchaUtils.ImgText 生成图片[]byte数组，可以转成base64或者image文件

```go
func TestCaptcha(t *testing.T) {
str := gotool.CaptchaUtils.GetRandStr(6)
fmt.Println("生成验证码字符串------------------->", str)
//生成图片byte数据，可以根据需要转成base64或者是image图片文件
text := gotool.CaptchaUtils.ImgText(100, 40, str)
sourcestring := base64.StdEncoding.EncodeToString(text)
fmt.Println(sourcestring)
}
//out
== = RUN   TestCaptcha
生成验证码字符串-------------------> qK5DME
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
--- PASS: TestCaptcha (0.01s)
PASS

```

Logs 日志打印工具
=======

#### 1、gotool.Logs.ErrorLog 异常日志

#### 2、gotool.Logs.InfoLog 加载日志

#### 3、gotool.Logs.DebugLog 调试日志

```go
func TestLogs(t *testing.T) {
gotool.Logs.ErrorLog().Println("Error日志测试")
gotool.Logs.InfoLog().Println("Info日志测试")
gotool.Logs.DebugLog().Println("Debug日志测试")
}
//out
== = RUN   TestLogs
[ERROR] 2021/07/07 15:58:10 logs_test.go:9: Error日志测试
[INFO] 2021/07/07 15:58:10 logs_test.go:10: Info日志测试
[DEBUG] 2021/07/07 15:58:10 logs_test.go:11: Debug日志测试
--- PASS: TestLogs (0.00s)
PASS
```

PageUtils 分页工具
=========

#### 1、gotool.PageUtils.Paginator 彩虹分页

```go
func TestPage(t *testing.T) {
paginator := gotool.PageUtils.Paginator(5, 20, 500)
fmt.Println(paginator)
}
//out
== = RUN   TestPage
map[AfterPage:6 beforePage:4 currPage:5 pages:[3 4 5 6 7] totalPages:25]
--- PASS: TestPage (0.00s)
PASS
//说明 AfterPage 下一页  beforePage前一页 currPage当前页 pages页码 totalPages总页数
```

IdUtils
=======
id生成工具，可生成字符串id和int类型id，根据需要选择自己需要的生成规则

#### 1、gotool.IdUtils.IdUUIDToTime 根据时间生成的UUID规则，入参 true消除“-”false保留“-”

```go
func TestUUID(t *testing.T) {
time, err := gotool.IdUtils.IdUUIDToTime(true)
if err == nil {
fmt.Println("根据时间生成UUID去除--------------------->'-'----->", time)
}
time, err = gotool.IdUtils.IdUUIDToTime(false)
if err == nil {
fmt.Println("根据时间生成不去除--------------------->'-'----->", time)
}
}
//out
== = RUN   TestUUID
根据时间生成UUID去除--------------------->'-'-----> 6fb94fe4dfd511ebbc4418c04d462680
根据时间生成不去除--------------------->'-'-----> 6fb9c783-dfd5-11eb-bc44-18c04d462680
--- PASS: TestUUID (0.00s)
PASS
```

#### 2、gotool.IdUtils.IdUUIDToRan 根据随机数生成的UUID推荐使用本方法，并发不会出现重复现象入，参 true消除“-”false保留“-”

```go
    time, err := gotool.IdUtils.IdUUIDToTime(true)
if err == nil {
fmt.Println("根据时间生成UUID去除--------------------->'-'----->", time)
}
time, err = gotool.IdUtils.IdUUIDToTime(false)
if err == nil {
fmt.Println("根据时间生成不去除--------------------->'-'----->", time)
}
//out
== = RUN   TestUUID
根据随机数生成UUID去除--------------------->'-'-----> cf5bcdc585454cda95447aae186d14e6
根据随机数生成不去除--------------------->'-'-----> 72035dba-d45f-480f-b1fd-508d1e036f71
--- PASS: TestUUID (0.00s)
PASS
```

#### 3、gotool.IdUtils.CreateCaptcha 生成随机数id，int类型，入参int 1-18，超过18后会造成int超过长度

```go
func TestCreateCaptcha(t *testing.T) {
captcha, err := gotool.IdUtils.CreateCaptcha(18)
if err != nil {
fmt.Println(err)
}
fmt.Println("18位------------------------------------------>", captcha)
captcha, err = gotool.IdUtils.CreateCaptcha(10)
if err != nil {
fmt.Println(err)
}
fmt.Println("10位------------------------------------------>", captcha)
}
//out
== = RUN   TestCreateCaptcha
18位------------------------------------------> 492457482855750014
10位------------------------------------------> 2855750014
--- PASS: TestCreateCaptcha (0.00s)
PASS
```

#### 4、gotool.IdUtils.GetIdWork根据时间戳在加以计算获取int64的id 长度16位

```go
func TestGetIdWork(t *testing.T) {
work := gotool.IdUtils.GetIdWork()
fmt.Println("根据时间戳在加以计算获取int64类型id-------->", work)
}
//out
== = RUN   TestGetIdWork
根据时间戳在加以计算获取int64类型id--------> 1625746675366450
--- PASS: TestGetIdWork (0.00s)
PASS
```

HttpUtils
=======

golang 的一个简单的“HTTP 请求”包 `GET` `POST` `DELETE` `PUT`

### 我们如何使用HttpUtils？

```go
resp, err := gotool.HttpUtils.Get("http://127.0.0.1:8000")
resp, err := gotool.HttpUtils.SetTimeout(5).Get("http://127.0.0.1:8000")
resp, err := gotool.HttpUtils.Debug(true).SetHeaders(map[string]string{}).Get("http://127.0.0.1:8000")

OR

req := gotool.HttpUtils.NewRequest()
req := gotool.HttpUtils.NewRequest().Debug(true).SetTimeout(5)
resp, err := req.Get("http://127.0.0.1:8000")
resp, err := req.Get("http://127.0.0.1:8000",nil)
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest")
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest", "address=beijing")

```

#### 设置请求头

```go
req.SetHeaders(map[string]string{
"Content-Type": "application/x-www-form-urlencoded",
"Connection": "keep-alive",
})

req.SetHeaders(map[string]string{
"Source":"api",
})
```

#### 设置Cookies

```go
req.SetCookies(map[string]string{
"name":"json",
"token":"",
})

OR

gotool.HttpUtils.SetCookies(map[string]string{
"age":"19",
}).Post()
```

#### 设置超时时间

```go
req.SetTimeout(5) //default 30s
```

#### 面向对象的操作模式

```go
req := gotool.HttpUtils.NewRequest().
Debug(true).
SetHeaders(map[string]string{
"Content-Type": "application/x-www-form-urlencoded",
}).SetTimeout(5)
resp, err := req.Get("http://127.0.0.1")

resp,err := gotool.HttpUtils.NewRequest().Get("http://127.0.0.1")
```

### GET

#### 查询参数

```go
resp, err := req.Get("http://127.0.0.1:8000")
resp, err := req.Get("http://127.0.0.1:8000",nil)
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest")
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest", "address=beijing")

OR

resp, err := gotool.HttpUtils.Get("http://127.0.0.1:8000")
resp, err := gotool.HttpUtils.Debug(true).SetHeaders(map[string]string{}).Get("http://127.0.0.1:8000")
```

#### 多参数

```go
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest", map[string]interface{}{
"name":  "jason",
"score": 100,
})
defer resp.Close()

body, err := resp.Body()
if err != nil {
return
}

return string(body)
```

### POST

```go
// Send nil
resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000")

// Send integer
resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000", 100)

// Send []byte
resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000", []byte("bytes data"))

// Send io.Reader
resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000", bytes.NewReader(buf []byte))
resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000", strings.NewReader("string data"))
resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000", bytes.NewBuffer(buf []byte))

// Send string
resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000", "title=github&type=1")

// Send JSON
resp, err := gotool.HttpUtils.JSON().Post("http://127.0.0.1:8000", "{\"id\":10,\"title\":\"HttpRequest\"}")

// Send map[string]interface{}{}
resp, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
"id":    10,
"title": "HttpRequest",
})
defer resp.Close()

body, err := resp.Body()
if err != nil {
return
}
return string(body)

resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000")
resp, err := gotool.HttpUtils.JSON().Post("http://127.0.0.1:8000", map[string]interface{}{"title":"github"})
resp, err := gotool.HttpUtils.Debug(true).SetHeaders(map[string]string{}).JSON().Post("http://127.0.0.1:8000","{\"title\":\"github\"}")
```

### 代理模式

```go
proxy, err := url.Parse("http://proxyip:proxyport")
if err != nil {
log.Println(err)
}

resp, err := gotool.HttpUtils.Proxy(http.ProxyURL(proxy)).Get("http://127.0.0.1:8000/ip")
defer resp.Close()

if err != nil {
log.Println("Request error：%v", err.Error())
}

body, err := resp.Body()
if err != nil {
log.Println("Get body error：%v", err.Error())
}
log.Println(string(body))
```

### 上传文件

Params: url, filename, fileinput

```go
resp, err := req.Upload("http://127.0.0.1:8000/upload", "/root/demo.txt", "uploadFile")
body, err := resp.Body()
defer resp.Close()
if err != nil {
return
}
return string(body)
```

### 提示模式

#### 默认 false

```go
req.Debug(true)
```

#### 以标准输出打印：

```go
[HttpRequest]
-------------------------------------------------------------------
Request: GET http: //127.0.0.1:8000?name=iceview&age=19&score=100
Headers: map[Content-Type:application/x-www-form-urlencoded]
Cookies: map[]
Timeout: 30s
ReqBody: map[age:19 score:100]
-------------------------------------------------------------------
```

## Json

Post JSON 请求

#### 设置请求头

```go
 req.SetHeaders(map[string]string{"Content-Type": "application/json"})
```

Or

```go
req.JSON().Post("http://127.0.0.1:8000", map[string]interface{}{
"id":    10,
"title": "github",
})

req.JSON().Post("http://127.0.0.1:8000", "{\"title\":\"github\",\"id\":10}")
```

#### Post 请求

```go
resp, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
"id":    10,
"title": "HttpRequest",
})
```

#### 打印格式化的 JSON

```go
str, err := resp.Export()
if err != nil {
return
}
```

#### 解组 JSON

```go
var u User
err := resp.Json(&u)
if err != nil {
return err
}

var m map[string]interface{}
err := resp.Json(&m)
if err != nil {
return err
}
```

TreeUtils
====== 
在gotool中实现了快捷菜单树的生成，包括菜单节点的选中状态、半选状态、菜单的搜索。

### 该工具提供了两种方法

#### 1、gotool.TreeUtils.GenerateTree(nodes, selectedNodes []INode) (trees []Tree)

`GenerateTree` 自定义的结构体实现 `INode` 接口后调用此方法生成树结构。

#### 2、gotool.TreeUtilsFindRelationNode(nodes, allNodes []INode) (respNodes []INode)

`FindRelationNode` 在 `allNodes` 中查询 `nodes` 中节点的所有父子节点 返回 `respNodes`(包含 `nodes` ， 跟其所有父子节点)

#### 测试代码
```go
package test

import (
	"encoding/json"
	"fmt"
	"github.com/druidcaesa/gotool"
	"github.com/druidcaesa/gotool/pretty"
	"github.com/druidcaesa/gotool/tree"
	"testing"
)

// 定义我们自己的菜单对象
type SystemMenu struct {
	Id    int    `json:"id"`    //id
	Pid   int    `json:"pid"`   //上级菜单id
	Name  string `json:"name"`  //菜单名
	Route string `json:"route"` //页面路径
	Icon  string `json:"icon"`  //图标路径
}

// region 实现ITree 所有接口
func (s SystemMenu) GetTitle() string {
	return s.Name
}

func (s SystemMenu) GetId() int {
	return s.Id
}

func (s SystemMenu) GetPid() int {
	return s.Pid
}

func (s SystemMenu) GetData() interface{} {
	return s
}

func (s SystemMenu) IsRoot() bool {
	// 这里通过FatherId等于0 或者 FatherId等于自身Id表示顶层根节点
	return s.Pid == 0 || s.Pid == s.Id
}

// endregion

type SystemMenus []SystemMenu

// ConvertToINodeArray 将当前数组转换成父类 INode 接口 数组
func (s SystemMenus) ConvertToINodeArray() (nodes []tree.INode) {
	for _, v := range s {
		nodes = append(nodes, v)
	}
	return
}

func TestGenerateTree(t *testing.T) {
	// 模拟获取数据库中所有菜单，在其它所有的查询中，也是首先将数据库中所有数据查询出来放到数组中，
	// 后面的遍历递归，都在这个 allMenu中进行，而不是在数据库中进行递归查询，减小数据库压力。
	allMenu := []SystemMenu{
		{Id: 1, Pid: 0, Name: "系统总览", Route: "/systemOverview", Icon: "icon-system"},
		{Id: 2, Pid: 0, Name: "系统配置", Route: "/systemConfig", Icon: "icon-config"},

		{Id: 3, Pid: 1, Name: "资产", Route: "/asset", Icon: "icon-asset"},
		{Id: 4, Pid: 1, Name: "动环", Route: "/pe", Icon: "icon-pe"},

		{Id: 5, Pid: 2, Name: "菜单配置", Route: "/menuConfig", Icon: "icon-menu-config"},
		{Id: 6, Pid: 3, Name: "设备", Route: "/device", Icon: "icon-device"},
		{Id: 7, Pid: 3, Name: "机柜", Route: "/device", Icon: "icon-device"},
	}

	// 生成完全树
	resp := gotool.TreeUtils.GenerateTree(SystemMenus.ConvertToINodeArray(allMenu), nil)
	bytes, _ := json.MarshalIndent(resp, "", "\t")
	//fmt.Println(string(pretty.Color(pretty.PrettyOptions(bytes, pretty.DefaultOptions), nil)))

	// 模拟选中 '资产' 菜单
	selectedNode := []SystemMenu{allMenu[2]}
	resp = gotool.TreeUtils.GenerateTree(SystemMenus.ConvertToINodeArray(allMenu), SystemMenus.ConvertToINodeArray(selectedNode))
	bytes, _ = json.Marshal(resp)
	//fmt.Println(string(pretty.Color(pretty.PrettyOptions(bytes, pretty.DefaultOptions), nil)))

	// 模拟从数据库中查询出 '设备'
	device := []SystemMenu{allMenu[5]}
	// 查询 设备 的所有父节点
	respNodes := gotool.TreeUtils.FindRelationNode(SystemMenus.ConvertToINodeArray(device), SystemMenus.ConvertToINodeArray(allMenu))
	resp = gotool.TreeUtils.GenerateTree(respNodes, nil)
	bytes, _ = json.Marshal(resp)
	fmt.Println(string(pretty.Color(pretty.PrettyOptions(bytes, pretty.DefaultOptions), nil)))
}

```
- 输出结果
```json
[
  {
    "title": "SystemOverview",
    "data": {
      "id": 1,
      "pid": 0,
      "name": "SystemOverview",
      "route": "/systemOverview",
      "icon": "icon-system"
    },
    "leaf": false,
    "checked": false,
    "partiallySelected": false,
    "children": [
      {
        "title": "assets",
        "data": {
          "id": 3,
          "pid": 1,
          "name": "assets",
          "route": "/asset",
          "icon": "icon-asset"
        },
        "leaf": false,
        "checked": false,
        "partiallySelected": false,
        "children": [
          {
            "title": "equipment",
            "data": {
              "id": 6,
              "pid": 3,
              "name": "equipment",
              "route": "/device",
              "icon": "icon-device"
            },
            "leaf": true,
            "checked": false,
            "partiallySelected": false,
            "children": null
          }
        ]
      }
    ]
  }
]
```
##### Thanks for this feature [azhengyongqin](https://github.com/azhengyongqin/golang-tree-menu)

DesensitizedUtils
======
在数据处理或清洗中，可能会涉及到很多隐私信息的脱敏，所gotool封装了一些常用信息的脱敏方法。
### 提供两个函数方法

#### 1、gotool.DesensitizedUtils.DefaultDesensitized(str string) (result string)

- 中国身份证、手机、身份证脱敏
- 一般邮箱脱敏

```go
func TestDesensitized(t *testing.T) {
	//邮箱脱敏
	mail := "testhello@gmail.com"
	star := gotool.DesensitizedUtils.DefaultDesensitized(mail)
	fmt.Println("mail-------------------->", star)
	//手机脱敏
	phone := "13333333333"
	desensitized := gotool.DesensitizedUtils.DefaultDesensitized(phone)
	fmt.Println("phone------------------->", desensitized)
}
//out
=== RUN   TestDesensitized
邮箱--------------------> tes***@gmail.com
手机-------------------> 133****3333
--- PASS: TestDesensitized (0.00s)
PASS
```

#### 1、gotool.DesensitizedUtils.CustomizeHash(str string, start int, end int) string

- 自定义脱敏
- `str`要脱敏的数据
- `start` 开始位置
- `end` 结束位置

```go
func TestDesensitized(t *testing.T) {
	//customize desensitization
	customize := "sadasdasdkljkldfjlkdjflkjsdf"
	hideStar := gotool.DesensitizedUtils.CustomizeHash(customize, 4, 14)
	fmt.Println("customize--------------->", hideStar)
}
//out
=== RUN   TestDesensitized
customize---------------> sada**********dfjlkdjflkjsdf
--- PASS: TestDesensitized (0.00s)
PASS
```
PrettyUtils
=======
PrettyUtils 是一个 gotool 包，它提供了格式化 JSON 以提高可读性的方法，或者为较小的有效负载压缩 JSON。
### Pretty
使用这个例子：
```json
{"name":  {"first":"Tom","last":"Anderson"},  "age":37,
"children": ["Sara","Alex","Jack"],
"fav.movie": "Deer Hunter", "friends": [
    {"first": "Janet", "last": "Murphy", "age": 44}
  ]}
```

以下代码：
```go
result = PrettyUtils.Pretty(example)
```

将 json 格式化为：

```json
{
  "name": {
    "first": "Tom",
    "last": "Anderson"
  },
  "age": 37,
  "children": ["Sara", "Alex", "Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {
      "first": "Janet",
      "last": "Murphy",
      "age": 44
    }
  ]
}
```

## Color

颜色将为输出到屏幕的 json 着色。
```json
result = PrettyUtils.Color(json, nil)
```

将为打印到终端的结果添加颜色。
第二个参数用于自定义样式，传递 nil 将使用默认的 `pretty.TerminalStyle`。
## Ugly

以下代码：
```go
result = PrettyUtils.Ugly(example)
```

将 json 格式化为：

```json
{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":["Sara","Alex","Jack"],"fav.movie":"Deer Hunter","friends":[{"first":"Janet","last":"Murphy","age":44}]}```
```

## 定制输出

有一个 `PrettyOptions(json, opts)` 函数，它允许使用以下选项自定义输出：
```go
type Options struct {
    // 宽度是单行数组的最大列宽
    // 默认值为 80
	Width int
    // 前缀是所有行的前缀
    // 默认为空字符串
	Prefix string
	// Indent 是嵌套的缩进
	// 默认为两个空格
	Indent string
	// SortKeys 将按字母顺序对键进行排序
	// 默认为假
	SortKeys bool
}
```


