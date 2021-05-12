package controllers

import (
	"backend/common"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Ggb(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ggb请求:", r.Method)
	common.SetHeader(w)

	//请求的数据
	requestData := common.GetData(r)
	code := requestData["code"]
	ggbName := requestData["ggbName"]
	//返回的数据体
	responseData := common.ResponseData{}

	ggbHtml, err := os.OpenFile(ggbName.(string), os.O_WRONLY|os.O_CREATE|os.O_EXCL, os.ModeAppend|os.ModePerm)
	if err != nil {
		responseData.Status = 202
		responseData.Message = "已存在这个ggb"
		res, _ := json.Marshal(responseData)
		w.Write(res)
	}
	defer ggbHtml.Close()

	html := `<html>

	<head>
		<meta name=viewport content="width=device-width,initial-scale=1">
		<meta charset="utf-8" />
		<script src="./GeoGebra/deployggb.js"></script>
	</head>
	
	<body>
		<div id="ggbApplet"></div>
		<div id="ggbid">`
	html = html + code.(string)
	html = html + `</div>
	</body>
	<script type="text/javascript">
		var ggb = document.getElementById("ggbid").innerText;
		document.getElementById("ggbid").style.display = "none";
		var parameters = {
			width: 750,
			height: 400,
			enableRightClick: false,
			showToolBar: false,
			showMenuBar: false,
			showToolBar: false,
			borderColor: "white",
			"language":"en",
			ggbBase64: ggb,
		};
		var applet = new GGBApplet(parameters, "5.0");
		applet.setHTML5Codebase('./GeoGebra/HTML5/5.0/web3d/');
		window.onload = function () {
			applet.inject("ggbApplet");
		};
	</script>
	
	</html>`
	
	ggbHtml.WriteString(html)

	responseData.Status = 201
	responseData.Message = "ggb以生成"
	res, _ := json.Marshal(responseData)
	w.Write(res)
}
