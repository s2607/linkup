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
	    <input name="nskey" type="hidden" value="{{.O.Pkey}}">
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

	<!--Stephen-->
	<form action="/newserv" method="post"> add a criterion  
			<input name="nskey" type="hidden" value="{{.O.Pkey}}">
			qprompt:<input name="qprompt" ><br>
			regex:<input name="regex"><br>
			aval:<input name="aval"><br>
			bval:<input name="bval"><br>
			lval:<input name="lval"><br>
			isnil:<input name="isnil"><br>
			inv:<input name="inv"><br>
			conj:<input name="conj"><br>
			<input type=submit>
	</form>
        <hr>
	<!--<<Stephen-->

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