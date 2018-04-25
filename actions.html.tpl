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
            <link href="css/popup_stylesheet.css" rel="stylesheet">

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
                    <h2>Add<a href="#popup-one"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h2>
				    <a href="/newop">Add An Interviewer</a><br /><br />
				    <a href="/newserv">Add A Service Program</a><br /><br />
				    <a href="/newq">Add A Question</a><br /><br />
                    <a href="/sql">Add A Database Entry</a>
                </div>

                <div id="right_container">
                    <h2>Search To Edit<a href="#popup-two"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h2>
				    <a href="/searchq">Search For A Question</a><br /><br />
				    <a href="/searcho">Search For An Interviewer</a><br /><br />
				    <a href="/searchs">Search For A Service Program</a><br /><br />

                </div>
            </div>
        </div><!--End container -->

        <!-- Pop up 1 -->
    <div class="popup" id="popup-one" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Add</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Add new interviewers, services, or questions.</p>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 1 -->
    <!-- Pop up 2 -->
    <div class="popup" id="popup-two" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Search</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Edit existing interviewers, services, or questions by searching for them and clicking Edit.</p>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 2 -->
</body>
</html>
