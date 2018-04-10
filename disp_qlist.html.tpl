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
            <h1>Select Question</h1>
        </div>

        <!-- TODO: Box on left with services (make dropdown in top left) -->

        <div id="qlist">
	       {{range .}}
	       <a href="/qprompt/{{.Pkey}}">{{.Pprompt}}</a><br><br>
	       {{end}}
        </div>

        <a href="/sugs" id="sug_services_button"><div id="submit_button" class="button_anim" style="width: 180px; padding-top: 6px; height: 29px; margin-left: auto; margin-right: auto; margin-top: 50px;">Suggest Services</div></a>

    </div>
</body>
</html>
