package mysql

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "time"
)

type Config struct {
    Host     string `yaml:"host"`
    Port     uint   `yaml:"port"`
    User     string `yaml:"user"`
    Password string `yaml:"password"`
    Database string `yaml:"database"`
    Charset  string `yaml:"charset"`
}

var Db *sql.DB

func (c Config) Init() (*sql.DB, error) {
    var err error
    sqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=%s",
        Config.User,
        Config.Password,
        Config.Host,
        Config.Port,
        Config.Database,
        Config.Charset,
    )
    Db, err = sql.Open("mysql", sqlDsn)
    if err != nil {
        log.Println("MySql Connect failure: ", err.Error())
        return nil, err
    }
    err = Db.Ping()
    if err != nil {
        log.Println("MySql Ping failure: ", err.Error())
        return nil, err
    }
    // Connection pool and timeouts
    // 连接池 和 超时
    Db.SetMaxOpenConns(150) // 最大打开连接数
    Db.SetMaxIdleConns(10)  // 最大空闲连接数
    // 连接过期时间 测试结果如下:
    // 1如不设过期时间 连接会一直不释放 连接池内连接数量为小于等于maxopen的数字
    // 2如设置了连接过期时间
    // 2.1 连接池内连接数量在连接过期后归零
    // 2.2 如之前连接数达到了最大打开连接数 连接池内连接数会依次经历: 由maxopen => maxidle => 0
    Db.SetConnMaxLifetime(time.Second * 5)
    return Db, nil
}