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
    <link href="css/search_stylesheet.css" rel="stylesheet">

</head>
<body>
    <div id="container">
        <div id="top_bar">
            <img id="logo" src="imgs/logo.png" alt="LinkUp">
        </div>

        <div id="left_container">

            <div id="left_title">
                <h1>Search Questions</h1>
            </div>

            <form id="form" action="/searchq" method="post" style="animation: none;">
                <p>Enter Question</p>
                <input type="search" name="q" spellcheck="true"><br>
                <input id="submit_button" value="Search" type=submit>
            </form>

        </div>

        <div id="right_container">

            <div id="right_title" style="animation: none;">
                <h1>Results</h1>
            </div>

            {{range .}}
            <div id="title">
                {{.Pprompt}} @ {{.Pkey}}
            </div>

            <form id="form" action="/newq" method="post">
                <p>Edit</p>
                <input name="nqkey" value="{{.Pkey}}" spellcheck="true"><br>
                <input id="submit_button" value="Submit" type=submit>
            </form><hr>
            {{end}}

        </div>



    </div>

</body>
</html>
