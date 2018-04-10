<!DOCTYPE html>
<html>
<head>
    <title>LinkUp</title>
    <link rel="icon" href="imgs/chevron.png" type="image/x-icon">
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <meta charset="UTF-8">
    <meta name="description" content="">
    <meta name="keywords" content="">

     <!-- Fonts -->
    <link href="https://fonts.googleapis.com/css?family=Roboto+Condensed|Nunito+Sans" rel="stylesheet">

    <!-- Stylesheet -->
    <link href="css/survey_stylesheet.css" rel="stylesheet">
</head>
<body>

    <div id="container">
        <div id="top_bar">
            <img id="logo" src="imgs/logo.png" alt="LinkUp">
        </div>

        <div id="title">
            <h1>Search for a {{.OB}}</h1>
        </div>

        <form id="form" action="{{.ACT}}" method="post">
            <input type="search" name="q"><br>
            <input id="submit_button" value="Search" type="submit">
        </form><hr>

        {{range .O}}
        <div id="title">
            {{.Ptext}}
        </div>
        <form id="form" action="/newq" method="post">
            <input name="qkey"type="hidden" value="{{.Pkey}}">
            <input id="submit_button" value="Submit" type=submit>
        </form>
        {{end}}

    </div>

</body>
</html>
