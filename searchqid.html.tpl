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

            <div id="left_title">
                <h1>Search Questions</h1>
            </div>

            <form id="form" action="/searchqid" method="post" style="animation: none;">
                <p>Enter Question</p>
                <input type="hidden" name="skey" value="{{.K}}">
                <input type="search" name="q" spellcheck="true"><br>
                <input id="submit_button" value="Search" type=submit>
            </form>

        </div>

        <div id="right_container">

            <div id="right_title" style="animation: none;">
                <h1>Results</h1>
            </div>

            {{$key := .K}}
            {{range .Q}}
            <div id="title">
                <div id="result">{{.Pprompt}}</div>
                <div id="id">ID: {{.Pkey}}</div>
            </div>

            <form id="result_form" action="/newserv" method="post">
                <input name="nqkey" type="hidden" value="{{.Pkey}}">
                <input name="nskey" type="hidden" value="{{$key}}">
                <input name="assoc" type="hidden" value="true">
                <input id="submit_button" value="Add" type=submit>
            </form>

            <div id="horiz_bar"></div>

            {{end}}

        </div>



    </div>

</body>
</html>
