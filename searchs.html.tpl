<html>
<body><form action="/searchs" method="post"> 
		name (regex):<input name="q" >
		<input type=submit>
</form><hr>
results:<br>
{{range .}}
	{{.Pname}} @ {{.Pkey}}<form action="/newserv" method="post">edit<input name="nskey" value="{{.Pkey}}"><input type=submit></form><hr>
{{end}}
</body>
</html>
