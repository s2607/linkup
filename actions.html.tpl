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
            <div id="top_bar" style="animation: extend_bar; animation-duration: 1s; animation-timing-function: ease-in-out;">
                <img id="logo" src="imgs/logo.png" alt="LinkUp">
            </div>

            <div id="title">
                <h1>Select An Action</h1>
                <h3>{{.wel}}</h3>
            </div>

            <div style="text-align: center; margin-top: 50px; animation: form_fade_in; animation-duration: 1s; animation-timing-function: ease-in-out;">
				<a href="/addresponder.html">Add A Responder</a><br /><br />
				<a href="/newop">Add An Interviewer</a><br /><br />
				<a href="/newserv">Add A Service Program</a><br /><br />
				<a href="/newq">Add A Question</a><br /><br />
				<a href="/searchq">Search For A Question</a><br /><br />
				<a href="/searcho">Search For An Interviewer</a><br /><br />
				<a href="/searchs">Search For A Service Program</a><br />
            </div>
        </div>
</body>
</html>
