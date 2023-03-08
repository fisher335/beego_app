package services

import (
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
)

var FilterUser = func(c *context.Context) {
	ok := c.Input.Cookie("token")
	if len(ok) == 0 {
		c.Redirect(302, "/login")
	}
}

func init() {
	fmt.Println("===================================启动顺序测试============")
}
