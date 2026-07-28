gotool
=======
GoTool is a small and complete Golang tool set. It mainly refines and integrates the commonly used methods in daily development to avoid repeating the wheel and improve work efficiency. Each method is extracted from the author's work experience and previous projects.

If you feel OK, please click STAR
### English [简体中文](./README-zh.md)
# Please pay attention
- [https://github.com/druidcaesa/gotool](https://github.com/druidcaesa/gotool)
- [https://gitee.com/termites/gotool](https://gitee.com/termites/gotool)
## Please see the documentation for detailed use of 2021-7-9 updates

- Add file IO manipulation tool FileUtils
- Add CaptchaUtils, a captcha generation tool
- Add file directory compression and decompression tool ZipUtis
- String array tool StrArrayUtils

## 2021-7-10 update content, please refer to the document for detailed usage
-  Add DesensitizedUtils, a sensitive message desensitization tool
-  Add the menu tree data formatting tool TreeUtils
-  Add JSON output beautification tool PrettyUtils

### How to use gotool?

### installation

go get github.com/druidcaesa/gotool

go.mod github.com/druidcaesa/gotool

### Introduce

```go
import "github.com/druidcaesa/gotool"
```

StrUtils
=======
Golang is a string commonly used tool set, which basically covers the tools that are often used in development, and is currently in the process of misconduct.

#### 1、gotool.StrUtils.ReplacePlaceholder Placeholder replacement

```go
func TestStringReplacePlaceholder(t *testing.T) {
s := "Hello {},My is {}"
placeholder, err := gotool.StrUtils.ReplacePlaceholder(s, "work", "gotool")
if err == nil {
fmt.Println(placeholder)
}
}
//out
== = RUN   TestStringReplacePlaceholder
Hello work, My is gotool
--- PASS: TestStringReplacePlaceholder (0.00s)
PASS
```

#### 2、gotool.StrUtils.RemoveSuffix Remove the file extension to get the file name

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

#### 3、gotool.StrUtils.GetSuffix Get file extension

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

#### 4、gotool.StrUtils.HasEmpty Judge whether the string is not empty, I return true if it is empty

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

StrArrayUtils string Array manipulation tool
======

#### 1、gotool.StrArrayUtils.StringToInt64 Convert string array to int64 array, please make sure that the string arrays are all numbers before calling

```go
func TestStringToInt64(t *testing.T) {
//String array to int64
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

#### 2、gotool.StrArrayUtils.StringToInt32 Convert string array to int64 array, please make sure that the string arrays are all numbers before calling

```go
func TestStringToInt32(t *testing.T) {
//String array to int32
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

#### 3、gotool.StrArrayUtils.ArrayDuplication Array deduplication

```go
func TestArrayDuplication(t *testing.T) {
//string array deduplication
strings := []string{"hello", "word", "gotool", "word"}
fmt.Println("Before weight removal----------------->", strings)
duplication := gotool.StrArrayUtils.ArrayDuplication(strings)
fmt.Println("After weight removal----------------->", duplication)
}
//out
== = RUN   TestArrayDuplication
Before weight removal-----------------> [hello word gotool word]
After weight removal-----------------> [hello word gotool]
--- PASS: TestArrayDuplication (0.00s)
PASS
```

DateUtil
=======
Golang is a time operation tool set, which basically covers the tools frequently used in development, and is currently being improved.

#### 1、gotool.DateUtil.FormatToString Time is formatted as a string

```go
func TestFormatToString(t *testing.T) {
now := gotool.DateUtil.Now()
toString := gotool.DateUtil.FormatToString(&now, "YYYY-MM-DD hh:mm:ss")
fmt.Println(toString)
toString = gotool.DateUtil.FormatToString(&now, "YYYYMMDD hhmmss")
fmt.Println(toString)
}
//Year, month and day correspond to YYYY MM DD hour, minute and second hhmmss can be combined in any combination, such as YYYY hh YYYY-DD hh:mm, etc.
//out
== = RUN   TestFormatToString
2021-07-07 16:13:30
20210707 161330
--- PASS: TestFormatToString (0.00s)
PASS
```

#### 2、gotool.DateUtil.IsZero Determine whether the time is empty

```go
//Time is empty true otherwise false
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

#### 3、gotool.DateUtil.Now Get the current time is equivalent to time.Now(), so this method is also included in the tool for unification

#### 4、gotool.DateUtil.InterpretStringToTimestamp   The string is formatted into a time type

```go
//Parameter one is the time string to be formatted. Parameter two is the string format, which needs to be consistent with the format of the string to be formatted. 
//For example, 2021-6-4 corresponds to YYYY-MM-DD 2021.6.4 corresponds to YYYY.MM.DD
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

#### 5、gotool.DateUtil.UnixToTime Timestamp to time

```go
func TestUnixToTime(t *testing.T) {
unix := gotool.DateUtil.Now().Unix()
fmt.Println("Timestamp----------------------->", unix)
toTime := gotool.DateUtil.UnixToTime(unix)
fmt.Println(toTime)
}
//out
== = RUN   TestUnixToTime
Timestamp-----------------------> 1625645682
2021-07-07 16:14:42 +0800 CST
--- PASS: TestUnixToTime (0.00s)
PASS
```

#### 6、gotool.DateUtil.GetWeekDay Get the day of the week

```go
func TestGetWeekDay(t *testing.T) {
now := gotool.DateUtil.Now()
day := gotool.DateUtil.GetWeekDay(now)
fmt.Println("Today is-----------------week", day)
}
//out
== = RUN   TestGetWeekDay
Today is-----------------week 3
--- PASS: TestGetWeekDay (0.00s)
PASS
```

#### 7、gotool.DateUtil.MinuteAddOrSub,HourAddOrSub,DayAddOrSub Time calculation tool

```go
//Time calculation
func TestTimeAddOrSub(t *testing.T) {
now := gotool.DateUtil.Now()
fmt.Println("The time is now--------------------->", now)
sub := gotool.DateUtil.MinuteAddOrSub(now, 10)
fmt.Println("Minute calculation result-------------------->", sub)
sub = gotool.DateUtil.MinuteAddOrSub(now, -10)
fmt.Println("Minute calculation result-------------------->", sub)
sub = gotool.DateUtil.HourAddOrSub(now, 10)
fmt.Println("Hourly calculation result-------------------->", sub)
sub = gotool.DateUtil.HourAddOrSub(now, -10)
fmt.Println("Hourly calculation result-------------------->", sub)
sub = gotool.DateUtil.DayAddOrSub(now, 10)
fmt.Println("Day calculation result-------------------->", sub)
sub = gotool.DateUtil.DayAddOrSub(now, -10)
fmt.Println("Day calculation result-------------------->", sub)
}
//The time is now---------------------> 2021-07-07 11:18:17.8295757 +0800 CST m=+0.012278001
//Minute calculation result--------------------> 2021-07-07 11:28:17.8295757 +0800 CST m=+600.012278001
//Minute calculation result--------------------> 2021-07-07 11:08:17.8295757 +0800 CST m=-599.987721999
//Hourly calculation result--------------------> 2021-07-07 21:18:17.8295757 +0800 CST m=+36000.012278001
//Hourly calculation result--------------------> 2021-07-07 01:18:17.8295757 +0800 CST m=-35999.987721999
//Day calculation result--------------------> 2021-07-17 11:18:17.8295757 +0800 CST m=+864000.012278001
//Day calculation result--------------------> 2021-06-27 11:18:17.8295757 +0800 CST m=-863999.987721999
```

ConvertUtils Gregorian calendar to lunar calendar tool
============

#### 1、gotool.ConvertUtils.GregorianToLunarCalendar(Gregorian calendar to lunar calendar),GetLunarYearDays(Lunar calendar to Gregorian calendar),GetLunarYearDays(Get the number of days in the lunar calendar year)

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

#### 2、gotool.ConvertUtils.GetSolarMonthDays(2021,7) Get the number of days in a certain month in the Gregorian calendar, the number of days in July 2021

#### 3、gotool.ConvertUtils.IsLeapYear(2021) Get whether a certain year is Ruinian true or false or not

#### 4、gotool.ConvertUtils.GetLeapMonth(2021) Get the leap month month of a certain year

BcryptUtils Encryption and decryption tools
==========

#### 1、gotool.BcryptUtils.Generate Encryption processing, mostly used for database storage after password encryption, irreversible

#### 2、gotool.BcryptUtils.CompareHash Compared with the unencrypted password after encryption, it is mostly used for login authentication.

```go
func TestGenerate(t *testing.T) {
//Encrypt
generate := gotool.BcryptUtils.Generate("123456789")
fmt.Println(generate)
//Encrypted comparison
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

#### 3、gotool.BcryptUtils.MD5 md5 encryption

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

#### 4、gotool.BcryptUtils.GenRsaKey(Obtain public and private keys),

#### 5、RsaSignWithSha256(Sign),

#### 6、RsaVerySignWithSha256(verification),

#### 7、RsaEncrypt(Public key encryption),

#### 8、RsaDecrypt(Private key decryption)

```go
func TestRsa(t *testing.T) {
//rsa Key file generation
fmt.Println("-------------------------------Obtain RSA public and private keys-----------------------------------------")
prvKey, pubKey := gotool.BcryptUtils.GenRsaKey()
fmt.Println(string(prvKey))
fmt.Println(string(pubKey))

fmt.Println("-------------------------------Sign and verify operations-----------------------------------------")
var data = "I am the encrypted data remember me-------------------------------"
fmt.Println("Sign the message...")
signData := gotool.BcryptUtils.RsaSignWithSha256([]byte(data), prvKey)
fmt.Println("Signature information of the message： ", hex.EncodeToString(signData))
fmt.Println("\nVerify the signature information...")
if gotool.BcryptUtils.RsaVerySignWithSha256([]byte(data), signData, pubKey) {
fmt.Println("The signature information is verified successfully, and it is determined that it is the correct private key signature！！")
}

fmt.Println("-------------------------------Perform encryption and decryption operations-----------------------------------------")
ciphertext := gotool.BcryptUtils.RsaEncrypt([]byte(data), pubKey)
fmt.Println("Public key encrypted data：", hex.EncodeToString(ciphertext))
sourceData := gotool.BcryptUtils.RsaDecrypt(ciphertext, prvKey)
fmt.Println("Data decrypted by the private key：", string(sourceData))
}
//out
== = RUN   TestRsa
-------------------------------Obtain RSA public and private keys-----------------------------------------
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

-------------------------------Sign and verify operations-----------------------------------------
Sign the message...
Signature information of the message：  1fcf20c4fb548c8fc0906e369287feb84c861bf488d822d78a0eada23d1af66ed3a12e9440d7181b1748fd0ad805222cf2ce7ce1f6772b330ef11b717700ba26945dda9d749a5c4d8c108ede103c17bed92801a4c3fbc1ebf38d10bf4bf183713eeb7f429acc3dcc3812366a324737f756720f3f9e75ddda1e024a7698b89163

Verify the signature information...
The signature information is verified successfully, and it is determined that it is the correct private key signature！！
-------------------------------Perform encryption and decryption operations-----------------------------------------
Public key encrypted data： 637b05798c1cf95cfcc63adf228645c77f8e9a9f34b12b722e6938ded8c9560a0430171a4f940d3fb2d96bc6c470c80a817d81f4a2669b991adbff5b22b335129e514c921083ce3e64c1c876409d9b763d5e8e269797283ef951a364da6a59a1d8454391356cb0cd0808092e9dd7ac371f9247a43760f3c82b7ad26a32a7a807
Data decrypted by the private key: I am the encrypted data, remember me-------------------------------
--- PASS: TestRsa (0.02s)
PASS
```

ZipUtils
=======
Compression and decompression tools, single file compression, directory compression, or cross-directory compression

#### 1、gotool.ZipUtils.Compress

- files File array can be multi-directory files
- dest Compressed file storage address

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
fmt.Println("Compressed successfully")
}
}
//out
== = RUN   TestCompress
Compressed successfully
--- PASS: TestCompress (0.12s)
PASS
```

#### 2、gotool.ZipUtils.DeCompress

- zipFile Compressed package path
- dest Path to decompress

```go
func TestDeCompress(t *testing.T) {
compress, err := gotool.ZipUtils.DeCompress("/Users/fanyanan/Downloads/test.zip", "/Users/fanyanan/Downloads")
if err != nil {
t.Fatal(err)
}
if compress {
fmt.Println("Successfully decompressed")
}
}
//out
== = RUN   TestDeCompress
Successfully decompressed
--- PASS: TestDeCompress (0.44s)
PASS
```

FileUtils
=======

File operation tools, making io operations easier and more convenient, making io operations not so difficult

#### 1、gotool.FileUtils.Exists Determine whether the file or directory exists

```go
func TestFileExists(t *testing.T) {
//Determine whether the file or directory exists
exists := gotool.FileUtils.Exists("F:/go-test/test")
fmt.Println("Before creation------------------------>", exists)
err := os.MkdirAll("F:/go-test/test", os.ModePerm)
if err != nil {
t.Fatal(err)
}
exists = gotool.FileUtils.Exists("F:/go-test/test")
fmt.Println("After creation------------------------>", exists)
}
//out
== = RUN   TestFileExists
Before creation------------------------> false
After creation------------------------> true
--- PASS: TestFileExists (0.00s)
PASS
```

#### 2、gotool.FileUtils.IsDir Determine whether it is a directory

```go
func TestIsDir(t *testing.T) {
//Determine whether it is a directory
dir := gotool.FileUtils.IsDir("F:/go-test/test")
fmt.Println("Is it a directory--------------------->", dir)
dir = gotool.FileUtils.IsDir("F:/go-test/test/test.txt")
fmt.Println("Is it a directory--------------------->", dir)
}
//out
== = RUN   TestIsDir
Is it a directory---------------------> true
Is it a directory---------------------> false
--- PASS: TestIsDir (0.00s)
PASS
```

#### 3、gotool.FileUtils.IsFile Determine whether it is a file

```go
func TestIsFile(t *testing.T) {
//Determine whether it is a file
file := gotool.FileUtils.IsFile("F:/go-test/test")
fmt.Println("Is it a file--------------------->", file)
file = gotool.FileUtils.IsFile("F:/go-test/test/test.txt")
fmt.Println("Is it a file--------------------->", file)
}
//out
== = RUN   TestIsFile
Is it a file---------------------> false
Is it a file---------------------> true
--- PASS: TestIsFile (0.00s)
PASS
```

#### 4、gotool.FileUtils.RemoveFile Delete files or directories

```go
func TestRemove(t *testing.T) {
//Delete Files
file, err := gotool.FileUtils.RemoveFile("F:/go-test/test/test.txt")
if err != nil {
t.Fatal(err)
}
if file {
//Check if the file still exists
exists := gotool.FileUtils.Exists("F:/go-test/test/test.txt")
fmt.Println("Does the file exist------------------------>", exists)
}
}
//out
== = RUN   TestRemove
Does the file exist------------------------> false
--- PASS: TestRemove (0.00s)
PASS
```

#### 5、gotool.FileUtils.OpenFileWronly Open the file in write-only mode, it will be automatically created if it is not, write to the internal slave for testing

```go
func TestOpenFileWronly(t *testing.T) {
//Open a file in write-only mode and write 5 pieces of content. If the file does not exist, one will be created
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
//Flush actually writes the cached file to the file
write.Flush()
//Read files and write to the console
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

#### 6、gotool.FileUtils.OpenFileRdonly Open the file in read-only mode, read the content and output to the console

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

#### 7、gotool.FileUtils.OpenFileAppend Open the file and append the content after the file, if not, the file will be created automatically

```go
func TestOpenFileAppend(t *testing.T) {
//Open the file and append data after the file
path := "F:/go-test/test/test.txt"
str := "Additional content \n"
wronly, err := gotool.FileUtils.OpenFileAppend(path)
if err != nil {
t.Fatal(err)
}
defer wronly.Close()
write := bufio.NewWriter(wronly)
for i := 0; i < 5; i++ {
write.WriteString(str)
}
//Flush actually writes the cached file to the file
write.Flush()
//Read files and write to the console
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
Additional content
Additional content
Additional content
Additional content
Additional content
--- PASS: TestOpenFileAppend (0.00s)
PASS
```

#### 8、gotool.FileUtils.FileCopy File copy method

```go
func TestFileCopy(t *testing.T) {
//File copy function
path := "F:/go-test/test/test.txt"
copyPath := "F:/go-test/test/test.txt1"
//Copy money
exists := gotool.FileUtils.Exists(copyPath)
fmt.Println("Does the file exist before copying------------------>", exists)
//After copying
fileCopy, err := gotool.FileUtils.FileCopy(path, copyPath)
if err != nil {
t.Fatal(err)
}
if fileCopy {
exists := gotool.FileUtils.Exists(copyPath)
fmt.Println("Does the file exist before copying------------------>", exists)
}
}
//out
== = RUN   TestFileCopy
Does the file exist before copying------------------> false
Does the file exist before copying------------------> true
--- PASS: TestFileCopy (0.00s)
PASS
```

CaptchaUtils Verification code generation tool
======

#### 1、gotool.CaptchaUtils.GetRandStr Generate verification code strng string

```go
func TestCaptcha(t *testing.T) {
//Generate verification code string, which can be stored in redis for verification logic
str := gotool.CaptchaUtils.GetRandStr(6)
fmt.Println(str)
}
//out
== = RUN   TestCaptcha
Generate verification code string-------------------> qK5DME
--- PASS: TestCaptcha (0.01s)
PASS
```

#### 2、gotool.CaptchaUtils.ImgText Generate image []byte array, which can be converted into base64 or image file

```go
func TestCaptcha(t *testing.T) {
str := gotool.CaptchaUtils.GetRandStr(6)
fmt.Println("Generate verification code string------------------->", str)
//Generate image byte data, which can be converted into base64 or image file as needed
text := gotool.CaptchaUtils.ImgText(100, 40, str)
sourcestring := base64.StdEncoding.EncodeToString(text)
fmt.Println(sourcestring)
}
//out
== = RUN   TestCaptcha
Generate verification code string-------------------> qK5DME
iVBORw0KGgoAAAANSUhEUgAAAGQAAAAoCAYAAAAIeF9DAAAbcUlEQVR4nOx6B3Rc1bX2PuXeO6MZjUaj3nu15SbLHWNjcIuN8SMYjAMYnPZCebz3QscQWmiPFFoCoRhCCAaDzcPGxrjLlnuTrGJJltW7ZkbSaGZuO+dfd2TLCMtEAvK/ZK3stWbNle69Z87Z3y7f3udQVdd64V/yDyP4/3oC/5LB8i9A/sHkX4D8g8l3BqTX60FlNeXk+5nOBenqdKHve8x/BvnOgDi7u9Avf3ev9Mzbz0lev/c7K9Ht6kYP3PmYyfgU7thPv+t4/2zynQFJikli655b6zOur7v3BnPh8b3fWom6rqOiXQfJls92kJNHT2Gf1/9dp/d/Jowx5FN9qNvfg9o9HbixuxnXOGtxeXslOd5UQg7UHyF7aoro9qrdtKG7aQCH78UCg0xB/P5b75XnT59PHn/9CWnT3s3k3pX3KI5gOx/JOLrO4NihEuLp8aDImAhud4SM6P2hRNEUpOgqyEwFv+pDsiaDrClI1VSQdRlk1Y9UXw9SmQYyoXzgeV0BWZORX5NB0RVQNRUZ/zt/bXzLugrq+ec1GfyaP3BtfHSmDztaZEaksTeuezFg1N9rSBiXOUb/4Om/+P60/k3h+vtuMP/H8ruURZct1L7+3LOP/F5afuu1amJKPPvq/0vrKvCR0uNEUVQQ7BSO+E+QkwdKiHxukaomo/PXiiYj1VCyoWBVPqcIQ4kqUs8rjmmI88GYSowjE+dg5hyMb6GmDyDeDKAyqLGIoAv4OxsBRpibBRMXiQAilfq/iQDCub8l4xoLYAo8I8LlqdP08++iv1dhWFVfjR9 //UkpJDiEP7jqfiU2PIYZbnx4/3H826deFe2hdn7rvy9XJ8+YODCZf3vh5qDiV05htV0Da4EFQhfZAZu/nYIM8zQRCjYkgAVhCEaUWwCDiAmwdh/q/rIGaS4fEJMAgBCYU6I5jQjjVBR4SHoit0SF8SCbjQeHhXHTeeUSEUQqBpRoKFUULihbpCZuIiKYqAQEk28N6qU9ROcIPdUiQrOC4bYIhU+y6CMZOCMxna15/E3/+1vW0ptWrzStunqlesO867XW5jbcUNuEz1SehbTMZP5VQAqE8XqFVom5yGHShAn69GmTdYvJcs7S+hVx4VoAiUiAKEEWamIiEZDEOSeqgogqAyh+xDV58KQQABEtvLjwCFnX84ro7/GD1uMHhBDIbT0IE4JMViv0fHkCrA4Hz5t5uRY7JoRPunqOjvF395zhyKUBKfIQ9F6nELje0WOGmcE63Bsr80yJDXdww1JuWnijOjv/cv3JN58WtxRtpQ+uvE+Jio3iJcdLcfmpSnz8cAkZX5Cn97h7EXVR8Lp84AgP5fesvEvJzs1gmAytCKYpyO/rwS7FyUO8fUiT+xBwFojb519AmAI1WTkxWzk12TiRgoz/8fSpVo5+/wcBY4zisrJYZHIyaygvJ6LJxDsaGrDi9YLf40Hb331HCIuL4ye2bydX3rJSjcvOZmar9e8KzKUBSZMYt2JAnn79oz29BPaeDoLFoRr/zygFYsVvBKajx4UjbKGBZ+Kj4tgfH3zZ/9mejfTOF+6WUjIzuL3Wxk+XVeHiY6XYAMTv96P9hUcC9cykafl6qMPOz4PBGUNM8YLm70Wa34N02QNMlZEb6cjEEWiAuRF2iGTh1GzlRLICMQVzIpiGVF5USgpzxMTwPrcb2SMj+fyf/kyNSk6Wa4uLSXdHB3K2tqLyvYXE2dqKXS0t6OS2bbSzoQFPv+6HasHCRZo1NPTvBsqlAYkVGd+c5YXft4lovZOCEVgYAPrURdEWN+UrwlX+i0gFbEPHy42Hd9A5Y6dpieExA8AtnrlImzZuuv7MG89Kns0epDSpUF5ymlSfPqu3NLXiro4uJEoiZGQl8TAbBl9nLdb9nnPWP/hnNNpv/aEmBwsoXzSsf/ixO23CBL21pgZ7XC7U29mJkvPyeM706RrTdaTKMsxesUKtOnyYlO4tJMe++ILWnTplPCvKXj+a/5OfKMP9nZEKeeTRRx685F0r4TDHpsECuwatKkZn5X6+rAOg416CPnAKwAFBnpkBHczyerwe1NXrRilRg5lUkGSGuVOv1Jpam/CpYxW4sbkJp6XE8K6WFvTZ+u3UYEUrlk1lUXaKmNKHjNBkWL+hfMHq4FJIFDeHJzO/I5I7rFFcModwTCVAaGQllb+vD5Xv20cUWUbhCQk8PT8/kMsQxkAFAQRJgsjkZJ4+caIen53NyvftDXhPQ1k5zp42jRmedX6spjP1uKb4NKkpqSQJWSnDDulDyTcDcl4clMMiuwZTrTpUywS1qf3alzlC+z0EbXAJEEwAskwccP8tkQpw4PQJmp8+ehAZ0FU/0rwuPHF0Ktu4fhdtb3Ghiuqj5Mi+MtLb44P8/Ex2zdKZekRcPJNskdzsSGBSeBLQkEguBYVyIyz5MSAOHCxIuKRHlJ49Szbs20vzM7OGVJDZYoGDn20UetrbUUh4OB9/1dyL6LmR7AVRhLD4+ADbaqmpIc7mZkRFAfIun6UbrNHd5sR/eeZP0qY3PhZVv4KyC8Ywk8X8rUKa09mDyOpHVz+EYJg1TKzIYZlDhRwzQ2U+DO5zxY+HIbS9h6KtPRSiBQ6pErOagmD3yf3CqJhYBv5uJLuasa+zjhjfap8LYSYjn09BJ05UY38fApfTE0jGVy2Zo8+99lrVGhbNBbONYyoCRwAekLEI1IhcqAcUsIEIGA097+6+PvT+jm10zZbNYmFJMb32sssvUrY52MbLCgtpV1MTtoSE8OQxY1iwwzGkIjEhAVA8LieqOXaMGDnN2dKMQuOS+eY164UjW4uopmpgtlpgysLLtJEA0tXZjXZuP0QLdx2lhTuPUnL7w//1uB80qnCNqKATDXSsAwt8WEAVAUOBQaClSowvD9MgWuSo1Eegr5/dgFNDaJObst2d1GdvBQ+uB6W3C9swAFP9BgsCTyPBlR8INDghik9ZPEX7y5pPhd5uT3+K4ADl1afx6aoqgnXCo2MjDUdAohDwBOQHDenAgQICE6KXXHRdWyu+66UXTa1OJ3L29KAD5aV088GDdMn0GYOAaa+rww3lZdhYYExaOovNyLhkuBFNJhAlCYrWrxdURYGM/Im8rakXb39/k6j4++l1xvgcPW1slm4bosOgKioi5OIe7PEjFfSeu18w7dl1lBQXVxLyy9X3PwLAkaF8nTOscR2rXCcq14gBkp+r1McVaoAmwwXQVFCxmsG4fylmml1DtJJh5O2fB25jSNwMNLLLgWocXSh13GjNFBrHTGFJ3FkWhus/l0jrfoHYk0MYiZDZ4aJjgZlijEH3MzhTWou3b9tNv/h0u9Da1I462jsx6BxUzHA39+MwHBTIGWgID+nz+9HKZ582n21tQSEWCxCEoayuFkuiAJsO7KfXzZo9AIrs9aKKoiLq9XhQWGwcz5o8+RtrLYs9FLa/s0b09vQgR1IuP328irTXtwYmYYSxmJQ4NuuH8zRjHYPAkBV0eOs+WvjpDuFMcSXJLrgQxnt7Pei9NZsEwyCjohychpFgHzcA4RzpwALAGGbJzoHEODOsEjGmIl3XkKxriDMNgJ0zJgkAbrRA9zJgQi1GtJQh7GRA3AwktwVP2+0A0ocluN6soiCqJ/+Aar52HZ352EtP/rZbnLn4Gu116R2Q/TIYVDchOZ53dXahjo4uVFVRg2qq6gRbSDAIgsBHF4xikYlRECIGsYSkOHbTj5epX114eV0d/t3HH4kHy8twVGgoXzl/oepXZFRUeoqU1tbicJuNX7P6QfOGJ34d6Bul5OXpUlAQd7a04PaGemwwLEwuzdQ6GxtRZFISa6iowBXHzxB3h2fAIozwGZ+eyAgd/L6uaahk71Hy0e/+LLo7nGjczAK9vbEVR8ZHBxSYkhrPgm0W6O3xQHZOCgvQXgSIE4Q4Odf8ZbqKmHyO9/s8CORehAw2hQlwghE3XI9ShCQLIMEEnAoILALwsYC0LA6o2o/VZhXBhakRONJJeLLEIUVisbcSZp1u5s27ZNzdoYsLZ1/D129ei8IiQ2HZzVdrY/Pz9C837iL7iw7T4pOncK+nF3SVoV2b9pBz8ZPExEXxlLREPmP2lIDFf35gP91fVkbW7twRWNNPF1+tLpoyVctKSGQb9xfRv+7YJhyvqsbFZ86QRQ/eZzYJIqx77AlfeHw8b6+rA6/Ljdpqz6KYtPRLA9LQgP1eGVNLLLQ3dCEiUAgOtfFeVw9SFRXyZkzQOecIGdo6J2UHSsiudVsFV1tXALzOpnbkiAofuB8UZOJJSTHsVEk1tljNfKAO0bxuJPd0YCZ7QFflwbEAGS5pYVSycmKy9Bdd4sWJK+BpQRyxMUGI2f2Eb3RT1iRjZscQ+HRzxBoUoqdK2JImsoQwkbtOIzQ/bSHauO0TqCw7g6rO1gljZo0nt9x/E5/fMJd1tnaxrVu24d2FRViQKchumdvDQ/m4SaNZfGLsQMw/Xl2NX97wSaCzcOWEfH1Sdo4eG96/8EVTp2kp0TF8zRdbaGFJMTlUXk6SoqL40tUPmW/JydFPHzxIep1dqKOuHhu5ZCgwFJ8PudraUFujE0CKAoOeZ03I1XvdvcgAJHl0OpOCTPyrYBh5Y/fHX9Dqk6cHkkfqmAwme/1AQ6wDY0dE2jmlBAySMwCIr7Me60r/BlN/yyH4XMsheKDlcCnLuYDbVzwtyarD7VbFublB4M83ixGN5n6QKQIWgpGWJRLrqnDVlGrRGouCxDlTfgBbCjfA/u0H0bipY1BsUjREJkQGPtkF2Xy5awUrLz2FXW43yk3NYjaLXY9NjQvM6dE1b0trd/bXMIbUtrbiVc89a7pp7lxtYmaWPm30aH1USor+6C0r2RufbxR2HT9OjlZWkoTISIjKzAy8ZNQYXS1NQ9I2w+q7mpvxhhf/KOnYDlyngAmGuTcvUdf+z9ui4bXRSbEMfyVpO1s68OFt++mx7QcpYxcwjstIZJgOzjEhocEcYWSEtwuVelBUOtOVPvRNLYdvI6Hz47Xf+L8kd8Ai3fSyU0QNCsJdOohFPoCiBsGUa2azfxwjNx9YKGxFn+GK4koo/rKUjc8dpcYlx3Ijh3VzGcc7IlHCZXOgtqWWvPvpuzg82EHC4lbxurOtcKCslHR2dyNJEEBWVejs6UbdfX3wm48+FCbn5JJTtWe1pTNmaulxcezua69Txqal0063W4sOC2N58QlMU1XJsH5XazvWVBVR4UJ9425txWeLT+Kta94Te7oxAAkCk8UEly29UqWUgiCJIJokkEwSxCTHDWj+8JdFdPe6LwNgGIAZxhKbksCCQ2zcbAkapF97iI1zxkFVvwIIkYIMTxgxEJqmIUWRwWQ2B/YBvn7fcOGUmAS9LMrJxm/J8sJaJ0Uvt4kGRQ7cL/Nh83/VSleOYmx/yjy+teZztO3DIpqXOZYl/Xu84gUdSYgyG5ICY+fFZcOvbn0I/fb9F8U7H7td+vmNv9B+MGWqRgkmo1PSWOK5Crqj243+un0bLampwW1Op2C3BvPwkBBut1r57HHjNb+iIJMoBp6NTc9g9WWlgb5Vxb59xGwLDjibr6cHVR49Qkr3HqCNZ9oRI47AmrInjtKvvHGRWl9Rg1W/ElB4dGp/R8JgVDXFVXj7B5uFjsZWRAUKcWmJrKGyFltCgweeG2S0jmCuaTp0dLjRt96gamquxzV1lbiu4Qzu9fQEek2jcsbpM6ddXPFmxCSxisYaMj4lW+c/ClP50lAN/alDQGvaRThHlZNLrfgOdTJsg81Q110FO18rFTLiciF6sV0NDVC5C2IxW/jDqx6Qj5Yf017839dMyZHJ7OW77pLjI6I5wTjAJQzLnFcwSVv91pvSieoq/NR770qjklPYtFGjApTzPBhGkZcxcaJuAHJ400Z65thREhYbywxq21pTgxW/H7gQDiBFAGgA6eOy2Jzli9Ww6HBeX1EDzWcbMMIYsieO1o2xakqq8MY3PhI7m9pQXGoiS8xNZXKfDzXXNIDVZuUpo9IvotbWYAs35tvY0IpHDIgs+1HRoZ1ky/YNwqGjhcTg3F5fHwhUAJstlL79/kui2+1C69/b23f+nfSYRLb56B6BMYYC+woWzPndUQq/0aGhV9oF9JFTAI1DPA6HO61LYJ9SBll9Eqp6WxWZLkP4UvNAM89YtNbrDWwq5edM0F9Lf9H3zs614h3P32G+/d9+Ji+csSBgEAYwk3Ny2Vv33Odf8vADZpfHg0prz+LzgJwXhDEffflM/diXW6nhIT2dncjV0kI0tZ9Rh8YmgrvPAaqqB0LVgpVLlYzx2ToiGOoqzgaSRmJWClN8MrTWt6AtazaIJfuOB/Qy6/p5qiMyjG949QNR13TInjR6yDqHkP6c0ufxjXwLd/e+L+hTL9wnGbUKZwySElKZzRbKGdMR5wxOlZ/AJlMQLLhuomXzR0cCoJhFE3cEh/BGZytKDI+9ENYiBcYfi5P5reEqfqFVjNgK9PbgxbCSXQWh2AYdjZ3Q/rwulLQSGP0zq6r7ZJA73Ght0R5KMYYFs2ZqEdGRbNXcH8lTxkwmz7/8tLh53xf0wVX3KTHhMQwjxH2KjBy2EN7mdiNXby/6aqg6L9lTp+k3PfFkgFqWHzxANFmBiIR4rnET1Fc246KNhTQkzM4nzJmi5V85dSACNFfVYyMkmS0m3tXcjnd/vJWUHjxJ7JEOnj9nijb96iu06uMVpP70WWxU6Xkz8jXDoNDXNrvCwvsre5NJHBkgew9spy+/8YyoqDJkpY9iebkT9BXX/SxgSsHBIbyqphxXnynDr7/zG8lsDoJf/PIG86v/80GgCBudmKl3ezwIwocYOFli7KUkP5z0EvPzraL5kId0B2lg92GIcvaiHesEcX9tHx51i6Z+VnKIvFu0mxoV+JSCCbqDhSOKMc+ITmWvPvYHecPmdfhHq1eafrzkNvX6uddpvV4fMkKWseIIu50L9OIlU1HkebNmabqqotGzZg0o/Oi2/XTLn7cGqHTa2Gw29QcXemItNY2YCJQznYGz3Ykrjp4i+zftoUzXwajWZy+bp5otZn5qf7+3ZBeM0o0iEWF8UQ6x220BQPx+ZfiAtLY34z+vfVXs6GxFGWk57NYVdypjR0/U7SFhA2iPG12gJ8ansNr6avy/Wz4UZNmPiw7vpNMKZmtTssZelFsukrFBOnsv1afscgv6HxpF03GGmBUBMzPecUwjR52AwhaE6iUNdThIkqCspQmlJScFXjVzynuxjm5YdKM2u2C2/uQbgR1KQQhOZ8FBQZAUFc1iwhxGjrkkcSHC4O7x2bJq4u5wIUd0OJ+yaKaamZ87EHIwJeBs68RG7He2dKBDW/ZRURLBeOaqHy1SI+KimJHga09VYSO82qPC+aWajj6fjELswdDt7h3+uawt2z6hJWXHcbDVBjff8AtlXN6kQWCcl5LSo2TfwZ1U13XIHzdVb2ysG/HZL8/lJk16M8Xv/Y1d9r4e4st7AMnmCOC9tYCb1lmFnLhk3ifLUFRWSsi5VofB5ixIZF5QcGxULH/toVf8YRGZbPOBPdQn+yEhMpzPL5j8t43inCiygkxBZj7u8on6vFuWqDkFeV/bRtCgtrQ6sDZN1cDn8ULWxFx90aofKmExEYE51ZaeCTQuRZMIklniBkhD/ZaRr0JDgwPvDEtZBotqbmkIDD4+b7KeGJ/KQ2wXb2NWVpeS519aLXW52tGkCTP0tORMtmzpyhHtrnlVGcttLqy3uriWjXUWhnlQFGKTHpZke47E7N1RkN6aiwzL3HbkMH3mr38R21wuw1IRAcwtIDKPz4c+2rVTKGvsxF4VQ5BIQHaXQXFVybCPvIqSyBf9+IfKv931I2XODQvV4FDboPXWllaTkPB+HRgsKyEjmV21YrGaOTF34ECEv8+HnG1dWPHLkJh96Y0rQig4wvrzyLBCliAI4O5xBdrHxltxMYmDBtd1PVCLPPXCvZK724WyM0axSfmX6cuvXTUiMNSePtTR1Y5CdHLBVggGKSyEW20WNuM57j/6TLcYfiSSFuB8dKrzFPqkcA8tra3FS6ZN1wVKeZ/Ph07VnsVvbf5cYJxDRnwCs5hM8Itl16j3/v5+6YqJs/U7l9+umKXh7VkkZqUMyYwsIcFc9vW3mBIzk9mC25YOCmm6piO/zwftDa1IMkuQMjqDDbDMr49lNfPomDAWiITDmRQCBPljp+rHig+Qvr5e0LQLTVZVU1FXVzt64ZVfSZVnynB8bDJfsvBGdfH8wZ3YvyVM0ZCvzYktCAMFY9IIaHAQSOE2js6FJWJCvGB1iGL6wzxet/ky+mv+LO7sasNfth3GWw4dpFGhobzd5UIIYVA0FQqyc1h2fDx78a67A2dSC3In6r/760vidffdYH7wtvvlaWOmjuho01clLiORzbjmCtXZ2omuWLZAzZk8OKQRSnjFkVKCCYa49ETm7e6DSx0lKpg8St+390Qg+Q9rC5dSCrLih12FW4T2jhbc1tGCszJGM0k0waFjheTD9WuEnYWf0yCzBX5+2z3KSMHoF460Hi9QjgCLAphiwrhot3L0tb0FhBFETZJ0YkKQe3wSblRbUER4CG/ytiFKCAiUotTYODZzzFj9oRU3KbcuWDgwF1EQYeaEy/TMxAz267eelU5Vl+KJuROYSTSNeLZBwRYekxLHx88q0GNS4/jX90B8fT607b3PhK7mDmx404ylV2jBoZc+Guv3K+jokVIy7JOLRkG4s3AzffWt50SMMITY7DwlKZOdrTuNK8+UB2bz33f8Sh6VPZ7lZI75VpbHNB1xVQNiloYVTpp2++mJ3/aIFWolsozizDFP06hI+OTcHGaRzIE2ySXXo8jotY9fFzbu/Zzevfwu5XxB+X1JXfkZ8vGL74nFhcdI9sRR+n+//phfEC99BqCj3YU/+XAbHdFR0qaWenzg8C6yaes6wQhPBpOCQEFjhpuv/7mydNEKdSjm9feUrhKFHH7CLal9HLJXWbW0pUEqAB+yrzaUnK6rxI+99oQUYQ/nD/z4fiXaMTQTGql0d7rwi//xtOnMydN44apr1WX/ebP8t97xeLxoeKdOzoktOIRHRkSDIEi87PRJ4vf7UHRUHP/B3B9qP1v5S8Wo0P9/S1AU4dGTJZ1pCJKuMmuSDfOhtnYvJeH2MH7NrCV6Z3cnevSPj0uSIEJuWs6IxhhKDFrccrYJy14/WrDyGvU8Ff4mEUVh5Iet3d1d6KEn7zAZFDc8PIpNK5it3/nTB/8m+v8MUt/agJ94/SlJYxo88pOH5ZS45O/sLa62LiyaRG7kkeE8P2xAdF1DRoh6+Kk7TUdPFBFHaASfd8XV6jWLVmjhjsjvxc3/UWT9jg3Cq+v+KNx29Up1+fwbvgVB+fby/wIAAP//mC6VL2hDN4cAAAAASUVORK5CYII=
--- PASS: TestCaptcha (0.01s)
PASS

```

Logs Log printing tool
=======

#### 1、gotool.Logs.ErrorLog Exception log

#### 2、gotool.Logs.InfoLog Load log

#### 3、gotool.Logs.DebugLog Debug log

```go
func TestLogs(t *testing.T) {
gotool.Logs.ErrorLog().Println("Error Log test")
gotool.Logs.InfoLog().Println("Info Log test")
gotool.Logs.DebugLog().Println("Debug Log test")
}
//out
== = RUN   TestLogs
[ERROR] 2021/07/07 15:58:10 logs_test.go:9: Error Log test
[INFO] 2021/07/07 15:58:10 logs_test.go:10: Info Log test
[DEBUG] 2021/07/07 15:58:10 logs_test.go:11: Debug Log test
--- PASS: TestLogs (0.00s)
PASS
```

PageUtils Paging tool
=========

#### 1、gotool.PageUtils.Paginator Rainbow Pagination

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
//Description AfterPage next page beforePage previous page currPage current page pages page number totalPages total number of pages
```

IdUtils
=======
id generation tool, can generate string id and int type id, choose the generation rules you need according to your needs

#### 1、gotool.IdUtils.IdUUIDToTime According to the UUID rules generated by time, enter the parameter true to eliminate "-" false to retain "-"

```go
func TestUUID(t *testing.T) {
time, err := gotool.IdUtils.IdUUIDToTime(true)
if err == nil {
fmt.Println("Generate UUID removal based on time--------------------->'-'----->", time)
}
time, err = gotool.IdUtils.IdUUIDToTime(false)
if err == nil {
fmt.Println("Generate according to time without removing--------------------->'-'----->", time)
}
}
//out
== = RUN   TestUUID
Generate UUID removal based on time--------------------->'-'-----> 6fb94fe4dfd511ebbc4418c04d462680
Generate according to time without removing--------------------->'-'-----> 6fb9c783-dfd5-11eb-bc44-18c04d462680
--- PASS: TestUUID (0.00s)
PASS
```

#### 2、gotool.IdUtils.IdUUIDToRan
It is recommended to use this method based on the UUID generated by the random number, and there will be no duplication in concurrency. Parameter true to eliminate "-" false to retain "-"

```go
    time, err := gotool.IdUtils.IdUUIDToTime(true)
if err == nil {
fmt.Println("Generate UUID removal based on time--------------------->'-'----->", time)
}
time, err = gotool.IdUtils.IdUUIDToTime(false)
if err == nil {
fmt.Println("Generate according to time without removing--------------------->'-'----->", time)
}
//out
== = RUN   TestUUID
Generate UUID removal based on time--------------------->'-'-----> cf5bcdc585454cda95447aae186d14e6
Generate according to time without removing--------------------->'-'-----> 72035dba-d45f-480f-b1fd-508d1e036f71
--- PASS: TestUUID (0.00s)
PASS
```

#### 3、gotool.IdUtils.CreateCaptcha Generate a random number id, int type, enter the parameter int 1-18, if it exceeds 18, it will cause the int to exceed the length

```go
func TestCreateCaptcha(t *testing.T) {
captcha, err := gotool.IdUtils.CreateCaptcha(18)
if err != nil {
fmt.Println(err)
}
fmt.Println("18 bits------------------------------------------>", captcha)
captcha, err = gotool.IdUtils.CreateCaptcha(10)
if err != nil {
fmt.Println(err)
}
fmt.Println("10 bits------------------------------------------>", captcha)
}
//out
== = RUN   TestCreateCaptcha
18 bits------------------------------------------> 492457482855750014
10 bits------------------------------------------> 2855750014
--- PASS: TestCreateCaptcha (0.00s)
PASS
```

#### 4、gotool.IdUtils.GetIdWorkAccording to the timestamp, it is calculated to obtain the id of int64. The length is 16 bits

```go
func TestGetIdWork(t *testing.T) {
work := gotool.IdUtils.GetIdWork()
fmt.Println("Calculate according to the timestamp to obtain the int64 type id-------->", work)
}
//out
== = RUN   TestGetIdWork
Calculate according to the timestamp to obtain the int64 type id--------> 1625746675366450
--- PASS: TestGetIdWork (0.00s)
PASS
```

HttpUtils
=======

A simple "HTTP request" package for golang `GET` `POST` `DELETE` `PUT`

### How do we use HttpUtils?

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

#### Set request header

```go
req.SetHeaders(map[string]string{
"Content-Type": "application/x-www-form-urlencoded",
"Connection": "keep-alive",
})

req.SetHeaders(map[string]string{
"Source":"api",
})
```

#### Set cookies

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

#### Set timeout

```go
req.SetTimeout(5) //default 30s
```

#### Object-oriented operation mode

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

#### Query parameter

```go
resp, err := req.Get("http://127.0.0.1:8000")
resp, err := req.Get("http://127.0.0.1:8000",nil)
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest")
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest", "address=beijing")

OR

resp, err := gotool.HttpUtils.Get("http://127.0.0.1:8000")
resp, err := gotool.HttpUtils.Debug(true).SetHeaders(map[string]string{}).Get("http://127.0.0.1:8000")
```

#### Multi-parameter

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

### Agency model

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

### upload files

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

### Prompt mode

#### default false

```go
req.Debug(true)
```

#### Print to standard output：

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

Post JSON request

#### Set request header

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

#### Post request

```go
resp, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
"id":    10,
"title": "HttpRequest",
})
```

#### Print formatted JSON

```go
str, err := resp.Export()
if err != nil {
return
}
```

#### Ungroup JSON

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
The quick menu tree generation is realized in gotool, including the selected state, half-selected state of the menu node, and the search of the menu.

### The tool provides two methods

#### 1、gotool.TreeUtils.GenerateTree(nodes, selectedNodes []INode) (trees []Tree)

The `GenerateTree` custom structure implements the `INode` interface and calls this method to generate the tree structure.

#### 2、gotool.TreeUtils.FindRelationNode(nodes, allNodes []INode) (respNodes []INode) 

`FindRelationNode` queries all the parent and child nodes of the node in `nodes` in `allNodes` Return `respNodes` (including `nodes`, and all its parent and child nodes)

#### Test code
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

// Define our own menu object
type SystemMenu struct {
	Id    int    `json:"id"`    //id
	Pid   int    `json:"pid"`   //Upper menu id
	Name  string `json:"name"`  //Menu name
	Route string `json:"route"` //Page path
	Icon  string `json:"icon"`  //Icon path
}

// region Implement all interfaces of ITree
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
	// Here, FatherId is equal to 0 or FatherId is equal to its own Id to indicate the top-level root node
	return s.Pid == 0 || s.Pid == s.Id
}

// endregion

type SystemMenus []SystemMenu

// ConvertToINodeArray Convert the current array to the parent INode interface array
func (s SystemMenus) ConvertToINodeArray() (nodes []tree.INode) {
	for _, v := range s {
		nodes = append(nodes, v)
	}
	return
}

func TestGenerateTree(t *testing.T) {
	// Simulate to obtain all the menus in the database. In all other queries, first query all the data in the database and put them in the array.
	// The subsequent traversal and recursion are all performed in this allMenu instead of recursive query in the database to reduce database pressure.
	allMenu := []SystemMenu{
		{Id: 1, Pid: 0, Name: "SystemOverview", Route: "/systemOverview", Icon: "icon-system"},
		{Id: 2, Pid: 0, Name: "SystemConfiguration", Route: "/systemConfig", Icon: "icon-config"},

		{Id: 3, Pid: 1, Name: "assets", Route: "/asset", Icon: "icon-asset"},
		{Id: 4, Pid: 1, Name: "MovingRing", Route: "/pe", Icon: "icon-pe"},

		{Id: 5, Pid: 2, Name: "MenuConfiguration", Route: "/menuConfig", Icon: "icon-menu-config"},
		{Id: 6, Pid: 3, Name: "equipment", Route: "/device", Icon: "icon-device"},
		{Id: 7, Pid: 3, Name: "Cabinet", Route: "/device", Icon: "icon-device"},
	}

	// Complete cabinet generation tree
	resp := gotool.TreeUtils.GenerateTree(SystemMenus.ConvertToINodeArray(allMenu), nil)
	bytes, _ := json.MarshalIndent(resp, "", "\t")
	// Simulate selecting the'Assets' menu
	selectedNode := []SystemMenu{allMenu[2]}
	resp = gotool.TreeUtils.GenerateTree(SystemMenus.ConvertToINodeArray(allMenu), SystemMenus.ConvertToINodeArray(selectedNode))
	bytes, _ = json.Marshal(resp)
	// Simulate querying the'equipment' from the database
	device := []SystemMenu{allMenu[5]}
	// Query all parent nodes of the device
	respNodes := gotool.TreeUtils.FindRelationNode(SystemMenus.ConvertToINodeArray(device), SystemMenus.ConvertToINodeArray(allMenu))
	resp = gotool.TreeUtils.GenerateTree(respNodes, nil)
	bytes, _ = json.Marshal(resp)
	fmt.Println(string(pretty.Color(pretty.PrettyOptions(bytes, pretty.DefaultOptions), nil)))
}
```
- Output structure`
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
In data processing or cleaning, a lot of desensitization of private information may be involved, so gotool encapsulates some desensitization methods for commonly used information.

### Provide two methods

#### 1、gotool.DesensitizedUtils.DefaultDesensitized(str string) (result string)

- Desensitization of Chinese ID cards, mobile phones, and ID cards
- General mailbox desensitization

```go
func TestDesensitized(t *testing.T) {
	//mail desensitization
	mail := "testhello@gmail.com"
	star := gotool.DesensitizedUtils.DefaultDesensitized(mail)
	fmt.Println("mail-------------------->", star)
	//phone desensitization
	phone := "13333333333"
	desensitized := gotool.DesensitizedUtils.DefaultDesensitized(phone)
	fmt.Println("phone------------------->", desensitized)
}
//out
=== RUN   TestDesensitized
mail--------------------> tes***@gmail.com
phone-------------------> 133****3333
--- PASS: TestDesensitized (0.00s)
PASS
```

#### 1、gotool.DesensitizedUtils.CustomizeHash(str string, start int, end int) string

- Custom desensitization
- `str` data to be desensitized
- `start` start position
- `end` end position

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
PrettyUtils is a gotool package that provides  methods for formatting JSON for human readability, or to compact JSON for smaller payloads.

### Pretty
Using this example:
```json
{"name":  {"first":"Tom","last":"Anderson"},  "age":37,
"children": ["Sara","Alex","Jack"],
"fav.movie": "Deer Hunter", "friends": [
    {"first": "Janet", "last": "Murphy", "age": 44}
  ]}
```

The following code:
```go
result = pretty.Pretty(example)
```

Will format the json to:

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

Color will colorize the json for outputing to the screen.

```json
result = pretty.Color(json, nil)
```

Will add color to the result for printing to the terminal.
The second param is used for a customizing the style, and passing nil will use the default `pretty.TerminalStyle`.

## Ugly

The following code:
```go
result = pretty.Ugly(example)
```

Will format the json to:

```json
{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":["Sara","Alex","Jack"],"fav.movie":"Deer Hunter","friends":[{"first":"Janet","last":"Murphy","age":44}]}```
```

## Customized output

There's a `PrettyOptions(json, opts)` function which allows for customizing the output with the following options:

```go
type Options struct {
	// Width is an max column width for single line arrays
	// Default is 80
	Width int
	// Prefix is a prefix for all lines
	// Default is an empty string
	Prefix string
	// Indent is the nested indentation
	// Default is two spaces
	Indent string
	// SortKeys will sort the keys alphabetically
	// Default is false
	SortKeys bool
}
```

