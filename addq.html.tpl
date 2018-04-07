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
            <h1>Add Question</h1>
        </div>

        <form id="form" action="/newq" method="post">
            <!--TODO:oldvals-->
            <p>prompt</p>
            <input name="prompt" value="{{.O.Pprompt}}">
            <p>qtype</p>
            <input name="qtype" value="{{.O.Ptype}}"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form><hr>

        <div id="title">
            <h2>Criteria</h2>
        </div>

        {{range .O.Pclist}}
	   <form id="form" action="delc" method="post">
           <p>Delete a Criterion</p>
           <input name="nckey" value="{{.Pkey}}"><br>
           <input id="submit_button" value="Submit" type=submit>
        </form><hr>
        {{end}}

        <form id="form" action="newq" method="post">
            <h3 style="color: #19396D;">Add A Criterion</h3>
            <p>Regex</p>
            <input name="regex">
            <p>aval</p>
            <input name="aval">
            <p>bval</p>
            <input name="bval">
            <p>lval:</p>
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

    </div>

</body>
</html>
