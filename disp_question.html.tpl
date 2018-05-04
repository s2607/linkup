<!DOCTYPE html>
<html>
<head>
    <title>LinkUP</title>
    <link rel="icon" href="../imgs/logo.png" type="image/x-icon">
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <meta charset="UTF-8">
    <meta name="description" content="">
    <meta name="keywords" content="">

    <!-- Fonts -->
    <link href="https://fonts.googleapis.com/css?family=Roboto+Condensed|Nunito+Sans" rel="stylesheet">

    <!-- Stylesheet -->
    <link href="../css/survey_stylesheet.css" rel="stylesheet">

</head>
<body>
    <div id="container">
        <div id="top_bar">
            <img id="logo" src="../imgs/logo.png" alt="LinkUp">
            <a href="/home"><div id="home_button">Home</div></a>
        </div>

        <div id="title" style="{{.A}};">
            <h1>Answer Question</h1>
        </div>

        <div id="error_msg">
            {{.M}}
            <br />
            {{$Question := .Q}}
            {{if ne .M ""}}<!-- show criterion for question if there is an error -->
                Valid Responses Are: <br />
                {{range .Q.Pclist}}
                    {{$Question.Pvalue .}}<br />
                {{end}}
            {{end}}
        </div>

        {{if .B}}
        <form id="form" method="post" style="{{.A}};">
            <div>{{.Q.Pprompt}}</div>
            <br>
            <select style="width: 75px;" name="qanswer" required>
                <option></option>
                <option value="1">Yes</option>
                <option value="0">No</option>
            </select>
            <br><br>
            <input id="submit_button" type=submit value="Submit">
        </form>
        {{else}}
        {{if .N}}
        <form id="form" method="post" style="{{.A}};">
            <div>{{.Q.Pprompt}}</div>
            <br>
            <input name="qanswer" type="number" value="" step="any" required>
            <br><br>
            <input id="submit_button" type=submit value="Submit">
        </form>
        {{else}}
        <form id="form" method="post" style="{{.A}};">
            <div>{{.Q.Pprompt}}</div>
            <br>
            {{if .EmptyL}}
            <input name="qanswer" type="text" value="" spellcheck="true" required>
            {{else}}
            <select style="min-width: 120px;" name="qanswer" required>
                <option></option>
                {{range .S}}
                <option value="{{.}}">{{.}}</option>
                {{end}}
            </select>
            {{end}}
            <br><br>
            <input id="submit_button" type=submit value="Submit">
        </form>
        {{end}}
        {{end}}

    </div>
</body>
</html>
