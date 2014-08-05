package gocontroller

var gamepadPage string = `
<!DOCTYPE html>
<html>
<head>
<script>
function httpGet(theUrl)
{
    var xmlHttp = null;

    xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", theUrl, false );
    xmlHttp.send( null );
    return xmlHttp.responseText;
}
</script>
<style type="text/css">
.buttonDirectional {
	background-color:#b8e356;
	-webkit-border-top-left-radius:42px;
	-moz-border-radius-topleft:42px;
	border-top-left-radius:42px;
	-webkit-border-top-right-radius:42px;
	-moz-border-radius-topright:42px;
	border-top-right-radius:42px;
	-webkit-border-bottom-right-radius:42px;
	-moz-border-radius-bottomright:42px;
	border-bottom-right-radius:42px;
	-webkit-border-bottom-left-radius:42px;
	-moz-border-radius-bottomleft:42px;
	border-bottom-left-radius:42px;
	text-indent:0;
	display:block;
	position:absolute;
	color:#ffffff;
	font-family:Arial;
	font-size:15px;
	font-weight:bold;
	font-style:normal;
	height:95px;
	line-height:95px;
	width:100px;
	text-decoration:none;
	text-align:center;
	text-shadow:1px 1px 0px #86ae47;
}
.buttonDirectional:hover {
	background-color:#a5cc52;
}</style>
<script>
function httpGet(theUrl)
{
    var xmlHttp = null;

    xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", theUrl, false );
    xmlHttp.send( null );
    return xmlHttp.responseText;
}
</script>
</head>
<body>
<button class="buttonDirectional" type="button" style="left:30%;top:20%;" onclick="httpGet('buttonUP')">UP</button>
<button  class="buttonDirectional" type="button" style="left:30%;top:60%;" onclick="httpGet('buttonDOWN')">DOWN</button>
<button  class="buttonDirectional" type="button" style="left:20%;top:40%;" onclick="httpGet('buttonLEFT')">LEFT</button>
<button  class="buttonDirectional" type="button" style="left:40%;top:40%;" onclick="httpGet('buttonRIGHT')">RIGHT</button>
<button  class="buttonDirectional" type="button" style="left:60%;top:40%;" onclick="httpGet('buttonA')">A</button>
<button  class="buttonDirectional" type="button" style="left:70%;top:40%;" onclick="httpGet('buttonB')">B</button>
<button  class="buttonDirectional" type="button" style="left:50%;top:10%;" onclick="httpGet('buttonSTART')">START</button>
</body>
</html>
`
