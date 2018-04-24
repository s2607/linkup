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
            <h1>Answer Question</h1>
        </div>

        {{if .B}}
        <form id="form" method="post">
            <div>{{.Q.Pprompt}}</div>
            <br>
            <select style="width: 75px;" name="qanswer" required>
                <option></option>
                <option value="1">Yes</option>
                <option value="0">No</option>
            </select>
            <br><br>
            <input id="submit_button" type=submit value="Submit">
        </form>
        {{else}}
        {{if .N}}
        <form id="form" method="post">
            <div>{{.Q.Pprompt}}</div>
            <br>
            <input name="qanswer" type="number" pattern="[0-9]" value="" spellcheck="true" required>
            <br>
            <input id="submit_button" type=submit value="Submit">
        </form>
        {{else}}
        <form id="form" method="post">
            <div>{{.Q.Pprompt}}</div>
            <br>
            <input name="qanswer" type="text" value="" spellcheck="true" required>
            <br>
            <input id="submit_button" type=submit value="Submit">
        </form>
        {{end}}
        {{end}}

    </div>
</body>
</html>
