// +build ignore

package main

import (
	"fmt"
	"time"

	"github.com/flw-cn/go-smartConfig"
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
	smartConfig.VersionDetail = "a\nlong\nlong\nversion\ndetail"
	smartConfig.LoadConfig("example", "1.0", config)

	fmt.Printf("config: %#v\n", config)

	// one := NewServiceOne(config.One)
	// go one.Run()

	// two := NewServiceTwo(config.Two)
	// go two.Run()

	// three := NewServiceThree(config.Three)
	// go three.Run()

	for {
		select {
		case <-smartConfig.ConfigChanged():
			fmt.Printf("new config: %#v\n", config)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
