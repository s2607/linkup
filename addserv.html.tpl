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

        <div id="title">
            <h1>Add Service</h1>
        </div>

        <form id="form" action="/newserv" method="post">
            <p>Name</p>
            <input name="name" value="{{.O.Pname}}">
            <p>Description</p>
            <input name="description" value="{{.O.Pdescription}}">
            <p>Website URL</p>
            <input name="url" value="{{.O.Purl}}"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form><hr>

        <div id="title">
            <h2>Criteria</h2>
        </div>

        {{range .O.Pclist}}
        <form id="form" action="delc" method="post">
            <p>delete</p>
            <input name="nckey" value="{{.Pkey}}"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form><hr>
        {{end}}

        <form id="form" action="newq" method="post">
            <h3 style="color: #19396D;">Add a Criterion</h3>
            <p>Regex (only for text)</p>
            <input name="regex">
            <p>Lower Limit (only for numerics)</p>
            <input name="aval">
            <p>Upper Limit (only for numerics)</p>
            <input name="bval">
            <p>lval</p>
            <input name="lval">
            <p>isnil:</p>
            <input name="isnil">
            <p>inv</p>
            <input name="inv">
            <p>conj</p>
            <input name="conj">
            <input type="hidden" name="nqkey" value="{{.O.Pkey}}"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form>
        <hr>

        <div id="title">
            <h2>Questions</h2>
        </div>
        {{range .O.Pqlist}}
        <form id="form" action="delc" method="post">{{.Pprompt}}
            <p>delete</p>
            <input name="nckey" value="{{.Pkey}}"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form><hr>
        {{end}}

        <form id="form" action="/newserv">
            <p>Add A Question (Type The Prompt Here)</p>
            <input name="nprompt"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form>

    </div>

</body>
</html>
