
<body><form action="/news" method="post"> 
		name:<input name="name" value="{{.O.Pname}}"><br>
		description:<input name="description" value="{{.O.Pdescription}}"><br>
		url:<input name="url" value="{{.O.Purl}}"><br>
		<input type=submit>
</form><hr>
Criteria:<br>
{{range .O.Pclist}}
	<form action="delc" method="post">delete <input name="nckey" value="{{.Pkey}}"><input type=submit></form><hr>
{{end}}
<form action="addc"> add a criterion <input name="rp" type="hidden" value="{{.A}}" ><input type=submit> </form>
<hr>
Questions:<br>
{{range .O.Pqlist}}
	<form action="delc" method="post">{{.Pprompt}} delete: <input name="nckey" value="{{.Pkey}}"><input type=submit></form><hr>
{{end}}
<form action="/news"> add a question (type the prompt here):<input name="nprompt"><input type=submit> </form>
</body>
