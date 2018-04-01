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

        <form action="/newserv" method="post">
            <p>Name</p>
            <input name="name" value="{{.O.Pname}}">
            <p>Description</p>
            <input name="description" value="{{.O.Pdescription}}">
            <p>Website URL</p>
            <input name="url" value="{{.O.Purl}}">
            <input id="submit_button" value="Submit" type=submit>
        </form><hr>

        Criteria:<br>
        {{range .O.Pclist}}
        <form action="delc" method="post">
            <p>delete</p>
            <input name="nckey" value="{{.Pkey}}">
            <input id="submit_button" value="Submit" type=submit>
        </form><hr>
        {{end}}

        <form action="addc">
            <p>Add A Criterion</p>
            <input name="rp" type="hidden" value="/neserv" >
            <input id="submit_button" value="Submit" type=submit>
        </form>
        <hr>

        Questions:<br>
        {{range .O.Pqlist}}
        <form action="delc" method="post">{{.Pprompt}}
            <p>delete</p>
            <input name="nckey" value="{{.Pkey}}">
            <input id="submit_button" value="Submit" type=submit>
        </form><hr>
        {{end}}

        <form action="/newserv">
            <p>Add A Question (Type The Prompt Here)</p>
            <input name="nprompt">
            <input id="submit_button" value="Submit" type=submit>
        </form>

    </div>

</body>
</html>
