// Code generated by hertztool.

package main

import (
	"code.byted.org/middleware/hertz/byted"
	"github.com/dirac-lee/ddd_demo_parking/biz/app"
)

func main() {
	app.Init()
	byted.Init()
	r := byted.Default()

	register(r)
	r.Spin()
}
