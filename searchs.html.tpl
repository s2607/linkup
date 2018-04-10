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
            <img id="logo" src="imgs/logo.svg" alt="LinkUp">
        </div>

        <div id="title" style="width: 400px;">
            <h1>Search Service Programs</h1>
        </div>

        <form id="form" action="/searchs" method="post">
            <p>Enter Program</p>
            <input type="search" name="q" ><br>
            <input id="submit_button" value="Search" type=submit>
        </form><hr>

        <div id="title">
            <h2>Results</h2>
        </div>

        {{range .}}
        <div id="title">
            {{.Pname}} @ {{.Pkey}}
        </div>
        <form id="form" action="/newserv" method="post">
            <p>Edit</p>
            <input name="nskey" value="{{.Pkey}}"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form><hr>
        {{end}}

    </div>
</body>
</html>
