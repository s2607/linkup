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
            <h1>Add Operator</h1>
        </div>

        <form id="form" action="/newop" method="post">
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
