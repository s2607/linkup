<html>
<body><form action="/searcho" method="post"> 
		uname (regex):<input name="q" >
		<input type=submit>
</form><hr>
results:<br>
{{range .}}
	{{.Puname}} @ {{.Pkey}}<form action="/newop" method="post">edit<input name="nskey" value="{{.Pkey}}"><input type=submit></form><hr>
{{end}}
</body>
</html>