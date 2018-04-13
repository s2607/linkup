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
            <img id="logo" src="../imgs/logo.png" alt="LinkUp">
            <a href="/home"><div id="home_button">Home</div></a>
        </div>

        <div id="title">
            <h1>Suggested Services</h1>
        </div>

        {{if .Empty}}
        <div class="service_sug">No Matching Service Programs</div>
        {{else}}
        <div id="slist">
	       {{range .S}}
	       <div id="suggestion"> <a class="sug_link" href="{{.Purl}}">{{.Pname}}</a><p>{{.Pdesc}}</p> </div><br>
	       {{end}}
	   </div>
        {{end}}

	   <a href="/qprompt"><div id="return_button" class="button_anim">Return</div></a>

    </div>

</body>
</html>
