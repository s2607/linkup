<body><form action="/newq" method="post"> 
		prompt:<input name="prompt" value="{{.O.Pprompt}}"><br>
		qtype:<input name="qtype" value="{{.O.Ptype}}"><br>
		<input type="hidden" name="nqkey" value="{{.O.Pkey}}">
		<input type=submit>
</form><hr>
Criteria:<br>
{{range .O.Pclist}}
	<form action="delc" method="post">delete <input name="nckey" value="{{.Pkey}}"><input type=submit></form><hr>
{{end}}
<form action="newq" method="post"> add a criterion 
		regex:<input name="regex"><br>
		aval:<input name="aval"><br>
		bval:<input name="bval"><br>
		lval:<input name="lval"><br>
		isnil:<input name="isnil"><br>
		inv:<input name="inv"><br>
		conj:<input name="conj"><br>
		<input type="hidden" name="nqkey" value="{{.O.Pkey}}">
		<input type=submit>
</form>
</body>
