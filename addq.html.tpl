<body><form action="/newq" method="post"> 
TODO:oldvals
		prompt:<input name="prompt" value="{{.O.Pprompt}}"><br>
		qtype:<input name="qtype" value="{{.O.Ptype}}"><br>
		<input type=submit>
</form><hr>
Criteria:<br>
{{range .O.Pclist}}
	<form action="delc" method="post">delete <input name="nckey" value="{{.Pkey}}"><input type=submit></form><hr>
{{end}}
<form action="addc"> add a criterion <input name="rp" type="hidden" value="/newq" ><input type=submit> </form>
</body>
