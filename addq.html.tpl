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
            <a href="/home"><div id="home_button" style="animation: fade_in; animation-duration: 1s; animation-timing-function: ease-in-out;">Home</div></a>
        </div>

        <div id="title">
            <h1>Add Question</h1>
        </div>

	{{.O.Pkey}}
        <form id="form" action="/newq" method="post">
            <!--TODO:oldvals-->
            <p>Question Prompt</p>
            <input name="prompt" value="{{.O.Pprompt}}" spellcheck="true" required>
            <input name="nqkey" value="{{.O.Pkey}}" type="hidden">
            </p><br>
            <input id="submit_button" value="Submit" type=submit>
        </form><hr>

        <div id="title">
            <h2>Validation Criteria</h2>
        </div>

        <form id="form" action="newq" method="post">
            <h3>Add Possible Answer</h3>
            <p>label</p><input name="vrtext" spellcheck="true">
            <input name="nqkey" value="{{.O.Pkey}}" type="hidden">
            <input id="submit_button" value="Submit" type=submit>
        </form>
        <hr>

        <p>Delete A Criterion</p>

        {{range .O.Pclist}}
	   <form id="form" action="delc" method="post">
           <input name="nckey" type="hidden" value="{{.Pkey}}"><br>
           <input id="submit_button" value="Delete" type=submit>
        </form><hr>
        {{end}}
	end



    </div>

</body>
</html>
