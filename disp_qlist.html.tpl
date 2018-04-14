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
    <link href="css/qlist_stylesheet.css" rel="stylesheet">

</head>
<body>
    <div id="container">
        <div id="top_bar">
            <img id="logo" src="imgs/logo.png" alt="LinkUp">
            <a href="/home"><div id="home_button">Home</div></a>
        </div>

        <div id="left_container">

            <div id="left_title">
                <h1>Find Questions</h1>
            </div>

            <form id="form" action="/qprompt" method="post">
                <p>Select Service Program</p>
                <select>
                    <option></option>
                    <option>Test</option>
                    <option>Test2</option>
                    <option>Test3</option>
                </select><br />
                <input id="find_button" value="Find Questions" type="submit">
            </form>

            <a href="/sugs"><div id="sug_services_button" class="button_anim">Suggest Services</div></a>

        </div>

        <div id="right_container">

            <div id="right_title">
                <h1>Select Question</h1>
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
