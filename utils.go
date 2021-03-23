package goutils

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "log"
    "os"
    "runtime"
    "strings"
)

// 检测err不为nil时记录日志
func ErrIsNotNil(err error, format ...string) bool {
    if err == nil {
        return false
    }
    if len(format) > 0 {
        log.Printf(format[0], err.Error())
    }
    CatchError(2, err)
    return true
}

// 异常捕获
func CatchError(skip int, err error) {
    if err == nil {
        return
    }
    pc, file, line, ok := runtime.Caller(skip)
    errorMessage := err.Error()
    if ok {
        //获取函数名
        pcName := runtime.FuncForPC(pc).Name()
        //pcName = strings.Join(strings.Split(pcName, "/")[1:], "/")
        dir,_ := os.Getwd()
        sep := "/"
        if file[0] != 47 {
            dir = strings.Replace(dir, "\\", "/", -1)
        }
        file = strings.TrimPrefix(file, dir+sep)
        errorMessage = fmt.Sprintf("【Error】%s:%d %s: %s", file, line, pcName, errorMessage)
    }
    log.Println(errorMessage)
}

// 获取字符串的md5值
func GetMd5String(s string) string {
    h := md5.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}
