<!DOCTYPE html>
<html>
<head>
    <title>LinkUp</title>
    <link rel="icon" href="../imgs/chevron.png" type="image/x-icon">
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <meta charset="UTF-8">
    <meta name="description" content="">
    <meta name="keywords" content="">

     <!-- Fonts -->
    <link href="https://fonts.googleapis.com/css?family=Roboto+Condensed|Nunito+Sans" rel="stylesheet">

    <!-- Stylesheet -->
    <link href="../css/survey_stylesheet.css" rel="stylesheet">

</head>
<body>
    <div id="container">
        <div id="top_bar">
            <img id="logo" src="../imgs/logo.svg" alt="LinkUp">
        </div>

        <div id="title">
            <h1>Suggested Services</h1>
        </div>



        <div id="slist">
	       {{range .}}
	       <div id="suggestion"> <a href="{{.Purl}}">{{.Pname}}</a><p>{{.Pdesc}}</p> </div><br>
	       {{end}}
	   </div>

	   <a href="/qprompt" id="sug_services_button"><div id="submit_button" style="width: 180px; padding-top: 6px; height: 29px; margin-left: auto; margin-right: auto; margin-top: 50px;">Return</div></a>

    </div>

</body>
</html>
