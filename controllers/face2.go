package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"proxy/models"
	"bytes"
	"net/http"
	"crypto/tls"
	"fmt"
	"os"
)

type ResponseInfo struct {
}

type FaceIIController struct {
	beego.Controller
}

func (this *FaceIIController) Post() {
    beego.Info("face2 method=post")
	re := this.Ctx.Input.Request

	beego.Info(re.URL.String())
	beego.Info("body len=",fmt.Sprintf("%d",re.ContentLength))
	body := this.Ctx.Input.RequestBody
	beego.Info("body=" + string(body))
	buf := new(bytes.Buffer)
	buf.Write(body)

	//beego.Info("##sos##"+buf.String())

	tr := &http.Transport{
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//resp,_ := client.Post("https://test.china-clearing.com/Gateway/InterfaceII","text/html",buf)
	resp,_ := client.Post("https://www.china-clearing.com/Gateway/InterfaceII","application/x-www-form-urlencoded",buf)
	//resp,_ := client.Post("https://1.202.139.206/Gateway/InterfaceII","application/x-www-form-urlencoded",buf)

	beego.Info("let go")
    resp.Request.Write(os.Stdout)
    resp.Header.Write(os.Stdout)
	beego.Info("the request to send clearing:")
	//resp.Write(os.Stdout)
	buf2 := new(bytes.Buffer)
	buf2.ReadFrom(resp.Body)
    beego.Info("len=",len(buf2.Bytes()))
	beego.Info("response=",string(buf2.Bytes()))
	this.Ctx.Request.Header.Set("content-type","text/plain")
	this.Ctx.ResponseWriter.Write(buf2.Bytes())
}

func (this *FaceIIController) Get() {
	objectId := this.Ctx.Input.Params[":objectId"]
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = ob
		}
	} else {
		obs := models.GetAll()
		this.Data["json"] = obs
	}
	this.ServeJson()
}

func (this *FaceIIController) Put() {
	objectId := this.Ctx.Input.Params[":objectId"]
	var ob models.Object
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = "update success!"
	}
	this.ServeJson()
}

func (this *FaceIIController) Delete() {
	objectId := this.Ctx.Input.Params[":objectId"]
	models.Delete(objectId)
	this.Data["json"] = "delete success!"
	this.ServeJson()
}
