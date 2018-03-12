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

        <div id="qlist">
	       {{range .}}
	       <a href="/qprompt/{{.Pkey}}">{{.Pprompt}}</a><br><br>
	       {{end}}
        </div>

        <a href="/sugs"><div id="submit_button" style="width: 200px;">Suggest Services</div></a>

    </div>
</body>
</html>
