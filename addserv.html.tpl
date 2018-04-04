
<body><form action="/newserv" method="post"> 
		name:<input name="name" value="{{.O.Pname}}"><br>
		description:<input name="description" value="{{.O.Pdescription}}"><br>
		url:<input name="url" value="{{.O.Purl}}"><br>
		<input name="nskey" type="hidden" value="{{.O.Pkey}}">
		<input type=submit>
</form><hr>
Criteria:<br>
{{range .O.Pclist}}
	<form action="delc" method="post">delete <input name="nckey" value="{{.Pkey}}"><input type=submit></form><br>
{{end}}
<form action="/newserv" method="post"> add a criterion  
		<input name="nskey" type="hidden" value="{{.O.Pkey}}">
		qprompt:<input name="qprompt" ><br>
		regex:<input name="regex"><br>
		aval:<input name="aval"><br>
		bval:<input name="bval"><br>
		lval:<input name="lval"><br>
		isnil:<input name="isnil"><br>
		inv:<input name="inv"><br>
		conj:<input name="conj"><br>
		<input type=submit>
</form>
<hr>
Questions:<br>
{{range .O.Pqlist}}
	<form action="delc" method="post">{{.Pprompt}} delete: <input name="nckey" value="{{.Pkey}}"><input type=submit></form><br>
{{end}}
<form action="/newserv" method="post"> add a question (type the prompt here):<input name="nprompt"><input type=submit> </form>
</body>
