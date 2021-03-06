<!DOCTYPE html>
<html>
<head>
    <title>LinkUP</title>
    <link rel="icon" href="imgs/logo.png" type="image/x-icon">
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <meta charset="UTF-8">
    <meta name="description" content="LinkUp">
    <meta name="keywords" content="LinkUp">
    
    <!-- CSS Stylesheets -->
    <link href="css/stylesheet.css" rel="stylesheet" />
    <!-- This is placed here so elements are shown and not hidden if login is invalid -->
    <link href="css/noscript_stylesheet.css" rel="stylesheet" />
    
    <!-- Fonts -->
    <link href="https://fonts.googleapis.com/css?family=Roboto+Condensed|Nunito+Sans" rel="stylesheet">

    
</head>
<body>
    
    <div id="container">

        <div id="left_container">
            <div id="logo_container">

                <img id="logo" src="imgs/logo.png" alt="LinkUp"/>

            </div>

        </div>

        <div id="right_container">

            <div id="login_container">
                <h1 id="title">LinkUP</h1>
		<div id="error" style="color:red;"> {{.err}} </div>
                <form id="login" action="/auth" method="post">
                    <p id="text_username">Username</p>
                    <input id="uname" name="uname" value="">
                    <p id="text_password">Password</p>
                    <input id="pw" type="password" name="pw" value="">
                    <input id="submit_button" name="submit" type="submit" value="Submit">
                </form>
            </div>

        </div>

    </div>

    
</body>
</html>
