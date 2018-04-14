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
            <link href="css/actions_stylesheet.css" rel="stylesheet">

    </head>
    <body>
        <div id="container">
            <div id="top_bar" style="animation: extend_bar; animation-duration: 1s; animation-timing-function: ease-in-out;">
                <img id="logo" src="imgs/logo.png" alt="LinkUp">
            </div>

            <div id="title">
                <h1>Select An Action</h1>
                <h3 id="user">{{.wel}}</h3>
                <a id="new_login" href="/oplogin.html">Login As Different User</a>
                <br /><br /><br />
                <a href="/addresponder.html" id="start_interview">Start An Interview</a>
            </div>

            <div id="content_container">

                <div id="left_container">
                    <h2>Add</h2>
				    <a href="/newop">Add An Interviewer</a><br /><br />
				    <a href="/newserv">Add A Service Program</a><br /><br />
				    <a href="/newq">Add A Question</a><br /><br />
                    <a href="/sql">Add A Database Entry</a>
                </div>

                <div id="right_container">
                    <h2>Search</h2>
				    <a href="/searchq">Search For A Question</a><br /><br />
				    <a href="/searcho">Search For An Interviewer</a><br /><br />
				    <a href="/searchs">Search For A Service Program</a><br /><br />

                </div>
            </div>
        </div>
</body>
</html>
