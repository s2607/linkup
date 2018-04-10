<!DOCTYPE html>
<html>
<head>
    <title>Untitled Document</title>
    <meta charset="UTF-8">
    <meta name="description" content="">
    <meta name="keywords" content="">
</head>
<body>

    Search for a {{.OB}}:<form action="{{.ACT}}" method="post"><input name="q"><input type="submit"></form><hr>
    {{range .O}}
    {{.Ptext}} <form action="/newq" method="post"><input name="qkey"type="hidden" value="{{.Pkey}}"><input type=submit></form>
    {{end}}

</body>
</html>
