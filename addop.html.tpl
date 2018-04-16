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
            <a href="/home"><div id="home_button" style="animation: {{.Anim}}; animation-duration: 1s; animation-timing-function: ease-in-out;">Home</div></a>
        </div>

        <div id="title" style="animation: form_{{.Anim}}; animation-duration: 1s; animation-timing-function: ease-in-out;">
            <h1>Add Interviewer</h1>
        </div>


        <div id="succ_msg">{{.Succ}}</div>

        <form id="form" action="/newop" method="post" style="animation: form_{{.Anim}}; animation-duration: 1s; animation-timing-function: ease-in-out;">
            <p>New Username</p>
            <input name="uname">
            <p>New Password</p>
            <input type="password" name="pw">
            <p>Service Key</p>
            <input name="skey"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form>

    </div>


</body>
</html>
