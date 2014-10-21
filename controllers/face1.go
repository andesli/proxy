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
//	"io/ioutil"
)

type FaceIController struct {
	beego.Controller
}

func (this *FaceIController) Post() {
	beego.Info("face1 method=post")

	re := this.Ctx.Input.Request
	//re.Header.Write(os.Stdout)
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
	//resp,_ := client.Post("https://www.china-clearing.com/Gateway/InterfaceI","application/x-www-form-urlencoded",buf)
	resp,_ := client.Post("https://106.39.51.19/Gateway/InterfaceI","application/x-www-form-urlencoded",buf)
	//resp,_ := client.Post("https://210.74.41.138/Gateway/InterfaceI","text/xml",buf)

	beego.Info("reques head: ")
	//resp.Request.Header.Write(os.Stdout)
	resp.Header.Write(os.Stdout)
	beego.Info("url: " + resp.Request.URL.String())

	buf2 := new(bytes.Buffer)
	buf2.ReadFrom(resp.Body)
    beego.Info("len=",len(buf2.Bytes()))
	beego.Info("response=",string(buf2.Bytes()))
	this.Ctx.ResponseWriter.Write(buf2.Bytes())
}

func (this *FaceIController) Get() {
		/*
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
	*/
	fmt.Println("get")
	re := this.Ctx.Input.Request
	fmt.Println(re.URL.String())
	//response, err := client.Do(re)
	buf := new(bytes.Buffer)
	buf.ReadFrom(re.Body)
	beego.Info("request body=",string(buf.Bytes()))
	
}

func (this *FaceIController) Put() {
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

func (this *FaceIController) Delete() {
	objectId := this.Ctx.Input.Params[":objectId"]
	models.Delete(objectId)
	this.Data["json"] = "delete success!"
	this.ServeJson()
}
