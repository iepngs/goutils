package rand

import (
    "bytes"
    "math/rand"
    "strings"
    "time"
)

// 生成[0, n)的随机数
func RandomNumber(n int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(n)
}

// 生成指定类型，长度为length的随机字符串
func RandomString(randLength int, randType string) string {
    var (
        num   string = "0123456789"
        lower string = "abcdefghijklmnopqrstuvwxyz"
        upper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    )
    b := bytes.Buffer{}
    if strings.Contains(randType, "0") {
        b.WriteString(num)
    }
    if strings.Contains(randType, "a") {
        b.WriteString(lower)
    }
    if strings.Contains(randType, "A") {
        b.WriteString(upper)
    }
    var str = b.String()
    var strLen = len(str)
    if strLen == 0 {
        return ""
    }
    rand.Seed(time.Now().UnixNano())
    b = bytes.Buffer{}
    for i := 0; i < randLength; i++ {
        b.WriteByte(str[rand.Intn(strLen)])
    }
    return b.String()
}
