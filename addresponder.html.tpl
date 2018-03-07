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
    <link href="css/addresponder_stylesheet.css" rel="stylesheet">

</head>
<body>
    <div id="container">
        <div id="top_bar">
            <img id="logo" src="imgs/logo.svg" alt="LinkUp">
        </div>

        <div id="title">
            <h1>Identifying Questions</h1>
        </div>

        <div id="error" style="color:red; text-align: center; margin-bottom: -20px;">{{.err}}</div>

        <form id="form" action="/newr" method="post">
            <p>First Name</p>
            <input name="fname">
            <p>Last Name</p>
            <input name="lname">
            <p>Date Of Birth</p>
		    <input name="dob">
            <p>Zip Code</p>
            <input name="zip"><br>
		  <input id="submit_button" type=submit>
        </form>

    </div>

</body>
</html>
