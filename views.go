package gocontroller

const gamepadPage string = `
<!DOCTYPE html>
<html>
<head>
<meta content='width=device-width; initial-scale=1.0; maximum-scale=1.0; user-scalable=0;' name='viewport' />
<meta name="viewport" content="width=device-width" />
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
	background-color:#2c2c2c;
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
	height:15%;
	line-height:15%;
	width:100px;
	text-decoration:none;
	text-align:center;
	outline:none;
}
.buttonDirectional:hover {
	background-color:#505050;
}</style>
<script>
function httpGet(theUrl)
{
    var xmlHttp = null;

    xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", theUrl, false );
    xmlHttp.send( null );
}
</script>
</head>
<body style="background: #803030;">
<button class="buttonDirectional" type="button" style="left:20%;top:20%;" onclick="httpGet('buttonUP')" onmousedown="httpGet('buttonUPtypePRESS')" onmouseup="httpGet('buttonUPtypeRELEASE')" >UP</button>
<button class="buttonDirectional" type="button" style="left:20%;top:60%;" onclick="httpGet('buttonDOWN')" onmousedown="httpGet('buttonDOWNtypePRESS')" onmouseup="httpGet('buttonDOWNtypeRELEASE')" >DOWN</button>
<button class="buttonDirectional" type="button" style="left:10%;top:40%;" onclick="httpGet('buttonLEFT')" onmousedown="httpGet('buttonLEFTtypePRESS')" onmouseup="httpGet('buttonLEFTtypeRELEASE')" >LEFT</button>
<button class="buttonDirectional" type="button" style="left:30%;top:40%;" onclick="httpGet('buttonRIGHT')" onmousedown="httpGet('buttonRIGHTtypePRESS')" onmouseup="httpGet('buttonRIGHTtypeRELEASE')" >RIGHT</button>


<button  class="buttonDirectional" type="button" style="left:60%;top:40%;" onclick="httpGet('buttonA')">A</button>
<button  class="buttonDirectional" type="button" style="left:80%;top:40%;" onclick="httpGet('buttonB')">B</button>
<button  class="buttonDirectional" type="button" style="left:45%;top:10%;" onclick="httpGet('buttonSTART')">START</button>
</body>
</html>
`
