# go-smart-config

这是一个【零配置】的配置文件加载模块和命令行参数处理模块

本模块中使用了 cobra 做命令行参数解析（当然主要是用的是 pflag）
并且通过 viper 绑定了配置文件支持。

## 主要特性

主要特性如下（主要是相对于标准的 flag 模块而言）：

* 既可以通过配置文件加载配置信息，也可以通过命令行来覆盖配置文件中的内容
* 完全独立的模块，使用起来极其简单，几乎不需要配置和调用就可以使用

### 支持通过命令行参数向程序传递值

* 支持命令行参数，支持 POSIX 风格的长短名
* 命令行参数 --help 打印出的帮助当中参数顺序可控，不会自动排序，更加具有可读性
* 支持通过小数点来分割多重数据结构的配置参数
* 命令行参数的帮助文本支持占位符，例如 --config FILENAME 这种较为友好的显示方式
* 支持 1h2m3s 这样格式的时间值

### 支持通过配置文件向程序传递值

* 配置文件同时支持 json/yaml 格式，不需要任何额外的编码
* 配置文件可以通过嵌套的方式来支持多重数据结构
* 按照 viper 的说法，配置文件还支持 etcd 等现代化云计算设施
* 支持 1h2m3s 这样格式的时间值
* 支持配置文件变更检测和自动重新加载
* 支持根据程序默认值生成配置文件，用户可以在生成的配置文件基础上进行修改

## 安装方法

```shell
go get -u "github.com/flw-cn/go-smart-config"
```

## 使用方法

导入模块:

```go
import (
    smartConfig "github.com/flw-cn/go-smart-config"
)
```

声明你的配置信息数据结构:

```go
type ServiceOneConfig struct {
    // 注意 struct tag 里的 flag
    IP   string `flag:"H|127.0.0.1|Listen {IP}"`
    Port int    `flag:"p|8080|Listen {Port}"`
}

type ServiceTwoConfig struct {
    // ...
}
type ServiceThreeConfig struct {
    // ...
}

type MainConfig struct {
    Debug bool `flag:"v|false|debug mode"`
    One   ServiceOneConfig
    Two   ServiceTwoConfig
    Three ServiceThreeConfig
}
```

加载配置:

```go
config := &MainConfig{}
smartConfig.LoadConfig("example", "1.0", config)
```

监测配置文件更新：

```go
select {
case <-smartConfig.ConfigChanged():
    // 监测到配置文件修改
}

```

定制 struct tag key 名称:

```go
type ServiceOneConfig struct {
    IP   string `myflag:"H|127.0.0.1|Listen {IP}"`
    Port int    `myflag:"p|8080|Listen {Port}"`
}

smartConfig.StructTagKey = "myflag"
smartConfig.LoadConfig("example", "1.0", config)
```

## 一个完整的例子

```go
package main

import (
    "fmt"
    "time"

    smartConfig "github.com/flw-cn/go-smart-config"
)

// 实际应用中建议将 ServiceOneConfig 等类型的定义放到具体的业务模块当中
type ServiceOneConfig struct {
    IP   string `flag:"H|127.0.0.1|Listen {IP}"`
    Port int    `flag:"p|8080|Listen {Port}"`
}

// 实际应用中建议将 ServiceTwoConfig 等类型的定义放到具体的业务模块当中
type ServiceTwoConfig struct {
    Foo bool   `flag:"|true|help message for foo"`
    Bar string `flag:"|blablabla|help message for bar"`
}

// 实际应用中建议将 ServiceThreeConfig 等类型的定义放到具体的业务模块当中
type ServiceThreeConfig struct {
    Hello int32         `flag:"|100|help message for hello"`
    World time.Duration `flag:"|30s|help message for world"`
}

type MainConfig struct {
    Debug bool `flag:"v|false|debug mode"`
    One   ServiceOneConfig
    Two   ServiceTwoConfig
    Three ServiceThreeConfig
}

func main() {
    config := &MainConfig{}
    smartConfig.LoadConfig("example", "1.0", config)

    fmt.Printf("config: %#v\n", config)

    // 然后就可以根据配置信息启动相应模块了
    // one := NewServiceOne(config.One)
    // go one.Run()

    // two := NewServiceTwo(config.Two)
    // go two.Run()

    // three := NewServiceThree(config.Three)
    // go three.Run()

    for {
        select {
        case <-smartConfig.ConfigChanged():
            // 监测到配置文件修改
            fmt.Printf("new config: %#v\n", config)
        default:
            time.Sleep(1 * time.Second)
        }
    }
}
```

## 更多内容参见

1. https://github.com/spf13/cobra
2. https://github.com/spf13/viper
3. https://github.com/spf13/pflag
