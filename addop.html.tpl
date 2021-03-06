<!DOCTYPE html>
<html>
<head>
    <title>LinkUP</title>
    <link rel="icon" href="imgs/logo.png" type="image/x-icon">
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <meta charset="UTF-8">
    <meta name="description" content="">
    <meta name="keywords" content="">

     <!-- Fonts -->
    <link href="https://fonts.googleapis.com/css?family=Roboto+Condensed|Nunito+Sans" rel="stylesheet">

    <!-- Stylesheet -->
    <link href="css/survey_stylesheet.css" rel="stylesheet">
    <link href="css/add_stylesheet.css" rel="stylesheet">
    <link href="css/popup_stylesheet.css" rel="stylesheet">

</head>
<body>

    <div id="container">
        <div id="top_bar">
            <img id="logo" src="imgs/logo.png" alt="LinkUp">
            <a href="/home"><div id="home_button" style="animation: {{.Anim}}; animation-duration: 1s; animation-timing-function: ease-in-out;">Home</div></a>
        </div>

        <div id="title" style="animation: form_{{.Anim}}; animation-duration: 1s; animation-timing-function: ease-in-out;">
            <h1>{{.T}} Interviewer<a href="#popup-one"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h1>
        </div>

        <div id="succ_msg">{{.Succ}}</div>

        <form id="form" action="/newop" method="post" style="animation: form_{{.Anim}}; animation-duration: 1s; animation-timing-function: ease-in-out;">
            <p>New Username</p>
            <input name="uname" value="{{.I.Puname}}" required>
            <p>New Password</p>
            <input type="password" name="pw" pattern="(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{6,}" title="Must contain at least one number and one uppercase and lowercase letter, and at least 6 or more characters" required>
            <p>Administrator</p>
            <input type="checkbox" class="checkbox" value="true" name="admin" {{if .I.Padmin}}checked{{end}}>
            <input type="hidden" name="nokey" value="{{.I.Pkey}}">
            <br>
            <input id="submit_button" value="Submit" type=submit>
        </form>

    </div><!-- End container -->

     <!-- Pop up 1 -->
    <div class="popup" id="popup-one" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>{{.T}} Interviewer</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                {{if .E}}
                <p class="popup-content">Change the username, password, and/or administrative privileges.</p>
                {{else}}
                <p class="popup-content">Create a new account for a person to be an interviewer or administrator.</p>
                {{end}}
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 1 -->


</body>
</html>
