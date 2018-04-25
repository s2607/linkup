<!DOCTYPE html>
<html>
<head>
    <title>LinkUP</title>
    <link rel="icon" href="imgs/chevron.png" type="image/x-icon">
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <meta charset="UTF-8">
    <meta name="description" content="">
    <meta name="keywords" content="">

     <!-- Fonts -->
    <link href="https://fonts.googleapis.com/css?family=Roboto+Condensed|Nunito+Sans" rel="stylesheet">

    <!-- Stylesheet -->
    <link href="css/survey_stylesheet.css" rel="stylesheet">
    <link href="css/search_stylesheet.css" rel="stylesheet">

</head>
<body>

    <div id="container">
        <div id="top_bar">
            <img id="logo" src="imgs/logo.png" alt="LinkUp">
            <a href="/home"><div id="home_button">Home</div></a>
        </div>

        <div id="left_container">
            <div id="left_title" style="width: 400px;">
                <h1>Search Service Programs</h1>
            </div>

            <form id="form" action="/searchs" method="post" style="animation: none;">
                <p>Enter Program</p>
                <input type="search" name="q" spellcheck="true"><br>
                <input id="submit_button" value="Search" type=submit>
            </form>

        </div>

        <div id="right_container">
            <div id="right_title">
                <h1>Results</h1>
            </div>

            {{range .}}
            <div id="title">
                <div id="result">{{.Pname}}</div>
                <div id="id">ID: {{.Pkey}}</div>
            </div>
            <form id="result_form" action="/newserv" method="post">
                <input name="nskey" type="hidden" value="{{.Pkey}}" spellcheck="true">
                <input id="submit_button" value="Edit" type=submit>
            </form>

            <div id="horiz_bar"></div>

            {{end}}

        </div>

    </div>
</body>
</html>
