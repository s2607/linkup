<!DOCTYPE html>
<html>
<head>
    <title>LinkUP</title>
    <link rel="icon" href="{{.P}}imgs/chevron.png" type="image/x-icon">
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <meta charset="UTF-8">
    <meta name="description" content="">
    <meta name="keywords" content="">

    <!-- Fonts -->
    <link href="https://fonts.googleapis.com/css?family=Roboto+Condensed|Nunito+Sans" rel="stylesheet">

    <!-- Stylesheet -->
    <link href="{{.P}}css/survey_stylesheet.css" rel="stylesheet">
    <link href="{{.P}}css/qlist_stylesheet.css" rel="stylesheet">

</head>
<body>
    <div id="container">
        <div id="top_bar">
            <img id="logo" src="{{.P}}imgs/logo.png" alt="LinkUp">
            <a href="/home"><div id="home_button">Home</div></a>
        </div>

        <div id="left_container">

            <div id="left_title">
                <h1>Find Questions</h1>
            </div>

            <form id="form" action="/qprompt" method="post">
                <p>Select Service Program</p>
                <select name="service" value="select">
                    <option value=""></option>
                    {{range .S}}
                    <option value="{{.Pkey}}">{{.Pname}}</option>
                    {{end}}
                </select><br />
                <input id="find_button" value="Find Questions" type="submit">
            </form>

            <a href="/sugs"><div id="sug_services_button" class="button_anim">Suggest Services</div></a>

        </div>

        <div id="right_container">

            <div id="right_title">
                <h1>Select Question</h1>
                <h3 id="serv_name">{{.Name}}</h3>
            </div>

            {{if .Empty}}
            <div class="no_questions">No Questions To Answer</div>
            {{else}}
            <div id="qlist">
                {{range .Q}}
                <a href="/qprompt/{{.Pkey}}">{{.Pprompt}}</a><br><br>
	           {{end}}
            </div>
            {{end}}

        </div>


    </div>
</body>
</html>
