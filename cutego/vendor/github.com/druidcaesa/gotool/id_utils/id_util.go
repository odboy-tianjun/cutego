package id_utils

import (
	"errors"
	"fmt"
	"github.com/druidcaesa/gotool/str"
	"github.com/google/uuid"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type IdUtils struct {
	s str.StrUtils
}

// IdUUIDToTime 根据时间生成UUID true去除“-”，false不去除
func (i IdUtils) IdUUIDToTime(flag bool) (string, error) {
	u1, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	if flag {
		replace := strings.Replace(u1.String(), "-", "", -1)
		return replace, err
	}
	return u1.String(), err
}

// IdUUIDToRan V4 基于随机数 true去除“-”，false不去除
func (i IdUtils) IdUUIDToRan(flag bool) string {
	u4 := uuid.New()
	if flag {
		return strings.Replace(u4.String(), "-", "", -1)
	}
	return u4.String()
}

// CreateCaptcha 生成整形随机数，可根据传入参数生成对应位数 入参数为1-18
// CreateCaptcha Generate an integer random number, the corresponding number of digits can be generated according to the input parameter. The input parameter is 1-18
func (i IdUtils) CreateCaptcha(n int) (int64, error) {
	if n < 0 {
		n = 1
	}
	pow10 := math.Pow10(n)
	strOne := "%0{}v"
	strTow := "%{}v"
	var str string
	if n < 10 && n > 0 {
		str, _ = i.s.ReplacePlaceholder(strOne, strconv.Itoa(n))
	} else if n >= 10 && n < 19 {
		str, _ = i.s.ReplacePlaceholder(strTow, strconv.Itoa(n))
	} else {
		return 0, errors.New("Please input 1-18, when the parameter exceeds 18, the int will exceed the length")
	}
	sprintf := fmt.Sprintf(str, rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(int64(pow10)))
	parseInt, err := strconv.ParseInt(sprintf, 10, 64)
	return parseInt, err
}

// GetIdWork 根据时间戳在加以计算获取int64的id
func (i IdUtils) GetIdWork() int64 {
	unix := time.Now().Unix()
	itoa := strconv.Itoa(int(unix))
	captcha, _ := i.CreateCaptcha(6)
	parseInt, _ := strconv.ParseInt(itoa+strconv.Itoa(int(captcha)), 10, 64)
	return parseInt
}
