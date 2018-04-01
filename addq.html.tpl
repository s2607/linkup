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

        <form action="/newq" method="post">
            <!--TODO:oldvals-->
            <p>prompt</p>
            <input name="prompt" value="{{.O.Pprompt}}">
            <p>qtype</p>
            <input name="qtype" value="{{.O.Ptype}}">
            <input id="submit_button" value="Submit" type=submit>
        </form><hr>

        Criteria:<br>
        {{range .O.Pclist}}
	   <form action="delc" method="post">
           <p>Delete a Criterion</p>
           <input name="nckey" value="{{.Pkey}}">
           <input id="submit_button" value="Submit" type=submit>
        </form><hr>
        {{end}}

        <form action="addc">
            <p>Add a Criterion</p>
            <input name="rp" type="hidden" value="/newq" >
            <input id="submit_button" value="Submit" type=submit>
        </form>

    </div>

</body>
</html>
