package controllers

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"os"
	"time"
)

type CommonController struct {
	beego.Controller
}

//proxy http with post method  to the real https
func (this *CommonController) Post() {
	re := this.Ctx.Input.Request
	beego.Info("request method = post")

	beego.Info("body len=", fmt.Sprintf("%d", re.ContentLength))
	body := this.Ctx.Input.RequestBody
	//beego.Info("body=" + string(body))
	buf := new(bytes.Buffer)
	buf.Write(body)

	url := ProxyUrl(re)
	if url == "" {
		beego.Info("error request")
		return
	}
	beego.Info("new url=" + url)

	client := NewClient()
	resp, err := client.Post(url, "application/x-www-form-urlencoded", buf)
	if err != nil {
		beego.Info(err.Error())
		this.Ctx.ResponseWriter.Write([]byte(err.Error() + "\n"))
		return
	}

	beego.Info("reques head: ")
	resp.Header.Write(os.Stdout)

	buf2 := new(bytes.Buffer)
	buf2.ReadFrom(resp.Body)
	beego.Info("len=", len(buf2.Bytes()))
	beego.Info("response=", string(buf2.Bytes()))
	this.Ctx.ResponseWriter.Write(buf2.Bytes())
}

//proxy http with get method  to the real https
func (this *CommonController) Get() {
	re := this.Ctx.Input.Request
	beego.Info("request method = get")

	url := ProxyUrl(re)
	if url == "" {
		beego.Info("Please setting url map in conf/app.cfg at first")
		return
	}
	beego.Info("new url=" + url)

	client := NewClient()
	resp, err := client.Get(url)
	if err != nil {
		beego.Info(err.Error())
		this.Ctx.ResponseWriter.Write([]byte(err.Error() + "\n"))
		return
	}

	beego.Info("reques head: ")
	resp.Header.Write(os.Stdout)

	buf2 := new(bytes.Buffer)
	buf2.ReadFrom(resp.Body)
	beego.Info("len=", len(buf2.Bytes()))
	beego.Info("response=", string(buf2.Bytes()))
	this.Ctx.ResponseWriter.Write(buf2.Bytes())
}

// NewClient创建一个带有超时机制的https客户端
func NewClient() *http.Client {
	ct, _ := beego.AppConfig.Int64("connecttimeout")
	rt, _ := beego.AppConfig.Int64("readwritetimeout")

	cts := (time.Duration)(ct) * time.Second
	rts := (time.Duration)(rt) * time.Second

	conf := &Config{
		ConnectTimeout:   cts,
		ReadWriteTimeout: rts,
	}

	tr := &http.Transport{
		DisableKeepAlives: true,
		Dial:              TimeoutDialer(conf),
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	return client

}

//turn the proxy http url to the real https url
func ProxyUrl(r *http.Request) string {
	url := r.URL.String()
	beego.Info("request url = " + url)

	index := url[7:11]
	uri := url[11:]
	url = beego.AppConfig.String(index)
	if url == "" {
		return ""
	}
	url = url + uri
	return url
}
