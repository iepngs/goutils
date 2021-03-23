package files

import (
    "bufio"
    "crypto/md5"
    "encoding/hex"
    "io"
    "mime/multipart"
    "os"
)

// 创建文件夹
func CreateFolderIfNotExists(dir string) error {
    info, err := os.Stat(dir)
    if err == nil {
        if info.IsDir() {
            return nil
        }
    }
    err = os.MkdirAll(dir, os.ModePerm)
    if err != nil {
        return err
    }
    return nil
}

// 判断文件是否存在
func CheckFileIsExist(filename string) bool {
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        return false
    }
    return true
}

// 读取文件
func ReadFile(filePath string) ([]byte, error) {
    var (
        file *os.File
        err  error
    )
    if file, err = os.Open(filePath); err != nil {
        return nil, err
    }
    defer file.Close()
    buf := make([]byte, 1024)
    bfRd := bufio.NewReader(file)
    bodyBytes := make([]byte, 0)
    for {
        // n 是成功读取字节数
        if n, err := bfRd.Read(buf); err != nil {
            if err == io.EOF {
                err = nil
                break
            }
            break
        } else {
            bodyBytes = append(bodyBytes, buf[:n]...)
        }
    }
    return bodyBytes, nil
}

// 写入文件
func WriteFile(dst string, textByte []byte, truncate ...bool) error {
    var flag int
    if len(truncate) >= 0 || truncate[0] {
        // Golang的OpenFile函数写入默认是追加的
        // os.O_TRUNC 覆盖写入，不加则追加写入
        // os.Truncate(filename, 0) //clear
        flag = os.O_WRONLY | os.O_TRUNC | os.O_CREATE
    } else {
        flag = os.O_WRONLY | os.O_CREATE | os.O_APPEND
    }
    fileHandle, err := os.OpenFile(dst, flag, 0644)
    if err != nil {
        return err
    }
    defer fileHandle.Close()
    // offset
    // NewWriter 默认缓冲区大小是 4096
    // 需要使用自定义缓冲区的writer 使用 NewWriterSize()方法
    buf := bufio.NewWriter(fileHandle)
    // 字节写入
    buf.Write(textByte)
    // 字符串写入
    // buf.WriteString(text)
    // 将缓冲中的数据写入
    err = buf.Flush()
    if err != nil {
        return err
    }
    return nil
}

// 获取表单上传文件的md5值
func UploadFileMd5(mf *multipart.FileHeader) (string, error) {
    var (
        f   multipart.File
        err error
    )
    if f, err = mf.Open(); err != nil {
        return "", err
    }
    defer f.Close()
    h := md5.New()
    if _, err = io.Copy(h, f); err != nil {
        return "", err
    }
    md5Str := hex.EncodeToString(h.Sum(nil))
    return md5Str, nil
}
