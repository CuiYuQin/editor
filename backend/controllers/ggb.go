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

	html := `<html>

	<head>
		<meta name=viewport content="width=device-width,initial-scale=1">
		<meta charset="utf-8" />
		<script src="https://www.geogebra.org/apps/deployggb.js"></script>
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
			width: 850,
			height: 420,
			enableRightClick: false,
			showToolBar: true,
			showAlgebraInput: true,
			showMenuBar: true,
			showToolBar: true,
			borderColor: "white",
			ggbBase64: ggb,
		};
		var applet = new GGBApplet(parameters, "5.0");
		window.onload = function () {
			applet.inject("ggbApplet");
		};
	</script>
	
	</html>`

	fmt.Println(html)

	ggbHtml, err := os.OpenFile(ggbName.(string), os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_TRUNC, os.ModeAppend|os.ModePerm)
	common.CheckError(err)
	defer ggbHtml.Close()
	ggbHtml.WriteString(html)

	responseData.Status = 201
	res, _ := json.Marshal(responseData)
	w.Write(res)
}
