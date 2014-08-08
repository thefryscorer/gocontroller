package gocontroller

import "fmt"

const layoutHead string = `
<!DOCTYPE html>
<html>
<head>
<meta content='width=device-width; initial-scale=1.0; maximum-scale=1.0; user-scalable=0;' name='viewport' />
<meta name="viewport" content="width=device-width" />
<style type="text/css">
%v
</style>
<script>
function httpGet(theUrl) {
    var xmlHttp = null;

    xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", theUrl, false );
    xmlHttp.send( null );
}

$(document).dblclick(function (e) {
		e.preventDefault();
});
</script>
</head>
<body>
`

const layoutFoot string = `
</body>
</html>
`

const DefaultCSS string = `
body {
	background:
	radial-gradient(#151515 5px, transparent 5px) 0px 0px, 
	radial-gradient(#222 5px, transparent 5px) 0px 23px,
	radial-gradient(#222 5px, transparent 5px) 25px 0px,
	radial-gradient(#1b1b1b 10px, transparent 10px),
	radial-gradient(#1d1d1d 25%, #1a1a1a 25%, #1a1a1a 50%, transparent 50%, transparent 75%, #242424 5%, #242424);
	background-size: 30px 30px;
	background-color: #131313;
}
.button {
	background-color:#2c2c2c;
	border-radius: 42px;
	text-indent:0;
	display:block;
	position:absolute;
	color:#ffffff;
	font-family:Arial;
	font-size:15px;
	font-weight:bold;
	font-style:normal;
	height:15%;
	line-height:15%;
	width:10%;
	text-decoration:none;
	text-align:center;
	outline:none;
}
`

const buttonTemplate string = `
<button class="button" type="button" style="left:%v%%;top:%v%%;%v" onclick="httpGet('button%v')" onmousedown="httpGet('button%[4]vtypePRESS')" onmouseup="httpGet('button%[4]vtypeRELEASE')" onmouseout="httpGet('button%[4]vtypeRELEASE')" ontouchstart="httpGet('button%[4]vtypePRESS')" ontouchend="httpGet('button%[4]vtypeRELEASE')" ontouchcancel="httpGet('button%[4]vtypeRELEASE')" >%v</button>
`

type Layout struct {
	Style   string
	Buttons []Button
}

func (l Layout) String() string {
	out := ""
	out += fmt.Sprintf(layoutHead, l.Style)
	for _, btn := range l.Buttons {
		out += btn.String()
	}
	out += layoutFoot
	return out
}

var DefaultLayout Layout = Layout{Style: DefaultCSS, Buttons: []Button{
	{Left: 20, Top: 20, Key: "Up"},
	{Left: 20, Top: 60, Key: "Down"},
	{Left: 10, Top: 40, Key: "Left"},
	{Left: 30, Top: 40, Key: "Right"},
	{Left: 60, Top: 40, Key: "A", Color: "#75B34D"},
	{Left: 80, Top: 40, Key: "B", Color: "#7895D1"},
	{Left: 45, Top: 10, Key: "Start"},
}}
