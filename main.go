package main

import (
	_ "proxy/routers"
	"github.com/astaxie/beego"	
)

//		Objects

//	URL					HTTP Verb				Functionality
//	/object				POST					Creating Objects
//	/object/<objectId>	GET						Retrieving Objects
//	/object/<objectId>	PUT						Updating Objects
//	/object				GET						Queries
//	/object/<objectId>	DELETE					Deleting Objects

func main() {
	beego.Run()
}
func init() {
	beego.CopyRequestBody=true
	beego.SetLogger("file", `{"filename":"logs/proxy.log"}`)
	//beego.SetLevel(beego.LevelError)
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogFuncCall(true)
	}
