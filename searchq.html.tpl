<!DOCTYPE html>
<html>
<head>
    <title>Untitled Document</title>
    <meta charset="UTF-8">
    <meta name="description" content="">
    <meta name="keywords" content="">
</head>
<body>

    <form action="/searchq" method="post">
		prompt (regex):<input name="q" >
		<input type=submit>
    </form><hr>
    results:<br>
    {{range .}}
	   {{.Pprompt}} @ {{.Pkey}}<form action="/newq" method="post">edit<input name="nqkey" value="{{.Pkey}}"><input type=submit></form><hr>
    {{end}}

</body>
</html>
