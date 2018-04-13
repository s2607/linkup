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
            <a href="/home"><div id="home_button">Home</div></a>
        </div>

        <div id="title">
            <h1>Select Question</h1>
        </div>

        <!-- TODO: Box on left with services (make dropdown in top left) -->

        {{if .Empty}}
        <div class="no_questions">No Questions To Answer</div>
        {{else}}
        <div id="qlist">
	       {{range .Q}}
            <a href="/qprompt/{{.Pkey}}">{{.Pprompt}}</a><br><br>
	       {{end}}
        </div>
        {{end}}

        <a href="/sugs"><div id="sug_services_button" class="button_anim">Suggest Services</div></a>

    </div>
</body>
</html>
