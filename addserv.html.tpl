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
    <link href="css/add_stylesheet.css" rel="stylesheet">

</head>
<body>

    <div id="container">
        <div id="top_bar">
            <img id="logo" src="imgs/logo.png" alt="LinkUp">
            <a href="/home"><div id="home_button" style="animation: fade_in; animation-duration: 1s; animation-timing-function: ease-in-out;">Home</div></a>
        </div>

        <div id="title">
            <h1>{{.T}} Service Program</h1>
        </div>

        <form id="form" action="/newserv" method="post">
            <p>Name</p>
            <input name="name" value="{{.O.Pname}}">
            <p>Description</p>
            <input name="description" value="{{.O.Pdescription}}" spellcheck="true">
            <p>Website URL</p>
            <input name="url" value="{{.O.Purl}}">
            <input name="nskey" type="hidden" value="{{.O.Pkey}}"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form>

        <div id="sub_container">

            <div id="left_container">
                <div id="title">
                    <h2>Questions</h2>
                </div>

                <form id="form" action="/newserv">
                    <h3>Associate Question To Service</h3>
                    <p>Question ID</p>
                    <input name="nskey" type="hidden" value="{{.O.Pkey}}">
                    <input name="nprompt" type="number" value="{{.Qid}}"><br>
                    <input id="submit_button" value="Submit" type=submit>
                </form>
                <form id="qid_form" action="/searchqid" method="post">
                        <input name="skey" type="hidden" value="{{.O.Pkey}}">
                        <input id="submit_button_qid" type='submit' value="Search For Question's ID">
                </form>
                <hr>

                <h3>Remove Question From Service</h3>
                {{range .O.Pqlist}}
                <form id="form" action="delc" method="post">
                    {{.Pprompt}}
                    <input name="nckey" type="hidden" value="{{.Pkey}}"><br>
                    <input id="submit_button" value="Remove" type=submit>
                </form>
                {{end}}

            </div><!--End left_container-->
            <div id="right_container">

                <div id="title">
                    <h2>Criteria</h2>
                </div>

                <form id="form" action="/newserv" method="post">
                    <h3>Add Eligibility Criterion</h3>
                    <p>Text Answers</p>
                    <input name="regex" spellcheck="true">
                    <p>Lower Limit</p>
                    <input type="number" name="aval">
                    <p>Upper Limit</p>
                    <input type="number" name="bval">
                    <p>Invert Range</p>
                    <input type="checkbox" class="checkbox" value="1" name="inv">
                    <p>Yes/No Questions <br><select name="lval">
                        <option></option>
                        <option value="1">Yes</option>
                        <option value="0">No</option>
                    </select></p><br>
                    <p>Conjunctive</p>
                    <input type="checkbox" name="conj" value="1" class="checkbox">
                    <p>isnil:</p>
                    <input type="checkbox" value="1" class="checkbox" name="isnil">
                    <input name="nskey" type="hidden" value="{{.O.Pkey}}"><br>
                    <input id="submit_button" value="Submit" type=submit>
                </form><hr>

                <h3>Remove Criterion From Service</h3>
                {{range .O.Pclist}}
                <form id="form" action="delc" method="post">
                    ID: {{.Pkey}}
                    <input name="nckey" type="hidden" value="{{.Pkey}}"><br>
                    <input id="submit_button" value="Delete" type=submit>
                </form><hr>
                {{end}}
            </div><!--End right_container -->
        </div><!--End sub_container -->
    </div><!-- End container-->

</body>
</html>
