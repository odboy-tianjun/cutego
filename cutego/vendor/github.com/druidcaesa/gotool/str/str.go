package str

import (
	"errors"
	"path"
	"strings"
	"unicode/utf8"
)

type StrUtils struct {
}

// ReplacePlaceholder 字符串占位符替换 占位符为"{}"  String placeholder replacement The placeholder is "{}"
func (u StrUtils) ReplacePlaceholder(s string, a ...interface{}) (string, error) {
	split := strings.Split(s, "{}")
	if len(split)-1 == len(a) {
		for _, item := range a {
			s = strings.Replace(s, "{}", item.(string), 1)
		}
	} else {
		return "", errors.New("Please check whether the number of placeholders matches the number of parameters," + s)
	}
	return s, nil
}

// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func (u *StrUtils) Replace(s, old, new string, n int) string {
	if old == new || n == 0 {
		return s // avoid allocation
	}

	// Compute number of replacements.
	if m := strings.Count(s, old); m == 0 {
		return s // avoid allocation
	} else if n < 0 || m < n {
		n = m
	}

	// Apply replacements to buffer.
	var b strings.Builder
	b.Grow(len(s) + n*(len(new)-len(old)))
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(s[start:], old)
		}
		b.WriteString(s[start:j])
		b.WriteString(new)
		start = j + len(old)
	}
	b.WriteString(s[start:])
	return b.String()
}

// HasEmpty It is to give some strings and return true if there are empty ones, which is often used to judge whether many fields are empty
func (*StrUtils) HasEmpty(s string) bool {
	if s == "" || len(s) == 0 {
		return true
	}
	return false
}

func (*StrUtils) HasNotEmpty(s string) bool {
	if s == "" || len(s) == 0 {
		return false
	}
	return true
}

// RemoveSuffix 去掉文件扩展名，直接获取文件名称
//Remove the file extension and get the file name directly
func (u *StrUtils) RemoveSuffix(str string) (string, error) {
	if u.HasEmpty(str) {
		return "", errors.New("Parameter  is an empty string")
	}
	filenameWithSuffix := path.Base(str)
	fileSuffix := path.Ext(filenameWithSuffix)
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly, nil
}

// GetSuffix 获取文件扩展名
// Get file extension
func (u *StrUtils) GetSuffix(str string) (string, error) {
	if u.HasEmpty(str) {
		return "", errors.New("Parameter  is an empty string")
	}
	filenameWithSuffix := path.Base(str)
	fileSuffix := path.Ext(filenameWithSuffix)
	return fileSuffix, nil
}
