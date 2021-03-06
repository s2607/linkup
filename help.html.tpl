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
            <link href="css/help_stylesheet.css" rel="stylesheet">

    </head>
    <body>
        <div id="container">
            <div id="top_bar">
                <img id="logo" src="imgs/logo.png" alt="LinkUp">
                <a href="/home"><div id="home_button" style="animation: fade_in; animation-duration: 1s; animation-timing-function: ease-in-out;">Home</div></a>
            </div>

            <div id="title">
                <h1>How To Use LinkUP</h1>
            </div>

            <p id="general">For more specific instructions on items, click the <span> <img id="icon" src="imgs/popup_icon.png"> </span> icon next to the item.</p>

            <div id="interview_container">
                <div id="title">
                    <h2>Interview Process</h2>
                </div>

                <h3 class="heading">Starting An Interview</h3>

                <div class="content">Select 'Start An Interview' on the home page. After beginning the interview process, the first step is to add the identifying information of the person being interviewed. This allows for unique identification of them. Once entered, the interviewer may now either select the person if they were previously interviewed or add a brand new entry to the system.</div>

                <h3 class="heading">Answering Questions</h3>

                <div class="content">Select a service program from the drop down menu. Answer the questions that appear.</div>

                <h3 class="heading">Suggesting Services</h3>

                <div class="content">Once satisfied with the number of questions ansswered, click the 'Suggest Services' button. The services that the person qualifies for will be displayed. Many of the services are clickable and will go to the program's website for more information.</div>

            </div><!-- End interview_container -->

            {{if .Admin}}
            <hr>

            <div id="service_container">
                <div id="title">
                    <h2>Creating and Editing Services</h2>
                </div>

                <h3 class="heading">Adding An Interviewer</h3>

                <div class="content">Select ‘Add An Interviewer’ on the home page. Fill in the fields with a unique username and a password for the interviewer to use. After filling in the form, press ‘Submit’. The new user will be able to log into the program to begin interviewing people.</div>

                <h3 class="heading">Editing An Interviewer</h3>

                <div class="content">Select ‘Search For An Interviewer’ on the home page. Search for the username of the interviewer you wish to edit. Select that user and type in the new username and password. </div>

                <h3 class="heading">Adding A Service Program</h3>

                <div class="content">Select ‘Add A Service Program’ on the home page. The 'Name' field is the name of the charity, organization, or service program. The description is to give a little information on what the program is. The Website URL field is an optional field that will give a link to the programs website for more detailed information about the program.</div>

                <h3 class="heading">Editing A Service Program</h3>

                <div class="content">If you just added a program, you are immediately brought to the editing screen of that same service program. If you are at the home page, select ‘Search For A Service Program’ and search and select the program to edit. Questions may be added to the program by selecting one of the two options to add a question to the service. Once a question is associated to the service, criteria can be added to the question that is used to tell whether or not a person can qualify for the service.</div>

                <h3 class="heading">Adding A Question</h3>

                <div class="content">Select ‘Add A Question’ on the home page. Input the question prompt and the type of answer that it will take. The validation criteria is what all possible answers are for the question. This is not what would make a person qualify for a service program but rather all possible values that will be accepted without returning an error.</div>

                <h3 class="heading">Editing A Question</h3>

                <div class="content">Select 'Search For A Question' on the home page. Search and select the question and edit the prompt for the question.</div>

                <h3 class="heading">Adding A Database Entry</h3>

                <div class="content">Select ‘Add A Database Entry’ on the home page. This allows you to use SQL commands to interact with the database. Only use this interface if you know SQL commands and the schema of the database.</div>

            </div><!-- End service_container -->
            {{end}}


        </div>

    </body>
</html>
