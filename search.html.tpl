<html><body>
Search for a {{.OB}}:<form action="{{.ACT}}" method="post"><input name="q"><input type="submit"></form><hr>
{{range .O}}
{{.Ptext}} <form action="/newq" method="post"><input name="qkey"type="hidden" value="{{.Pkey}}"><input type=submit></form>
{{end}}
