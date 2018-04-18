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
    <link href="css/popup_stylesheet.css" rel="stylesheet">

</head>
<body>

    <div id="container">
        <div id="top_bar">
            <img id="logo" src="imgs/logo.png" alt="LinkUp">
            <a href="/home"><div id="home_button" style="animation: fade_in; animation-duration: 1s; animation-timing-function: ease-in-out;">Home</div></a>
        </div>

        <div id="title">
            <h1>{{.T}} Service Program<a href="#popup-one"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h1>
        </div>

        <form id="form" action="/newserv" method="post">
            <p>Name</p>
            <input name="name" value="{{.O.Pname}}" required>
            <p>Description</p>
            <input name="description" value="{{.O.Pdescription}}" spellcheck="true" required>
            <p>Website URL</p>
            <input name="url" type="url" value="{{.O.Purl}}">
            <input name="nskey" type="hidden" value="{{.O.Pkey}}"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form>

        <div id="sub_container">

            <div id="left_container">
                <div id="title">
                    <h2>Questions</h2>
                </div>

                <form id="form" action="/newserv">
                    <h3>Associate Question To Service<a href="#popup-two"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h3>
                    <p>Question ID</p>
                    <input name="nskey" type="hidden" value="{{.O.Pkey}}">
                    <input name="nprompt" type="number" value="{{.Qid}}" required><br>
                    <input id="submit_button" value="Submit" type=submit>
                </form>
                <form id="qid_form" action="/searchqid" method="post">
                        <input name="skey" type="hidden" value="{{.O.Pkey}}">
                        <input id="submit_button_qid" type='submit' value="Search For Question ID">
                </form>
                <hr>

                <h3>Remove Question From Service<a href="#popup-three"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h3>
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
                    <h2>Eligibility Criteria</h2>
                </div>

                <form id="form" action="/newserv" method="post">
                    <h3>Add Criterion</h3>
                    <p>Question ID</p>
                    <input name="qid" value="{{.Qid}}">
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
                    <input name="nqkey" type="hidden" value="{{.O.Pkey}}"><br>
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

    <!-- Pop up 1 -->
    <div class="popup" id="popup-one" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>{{.T}} Service Program</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Enter the name of your charity, organization, or non-profit and provide a short description of it. If you wish to include a link to your webpage, include the full web address including the 'http://'.</p>
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
                <h2>Associate A Question</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content"><b>Steps:</b></p>
                <p class="popup-content">1. Click the Search For Question ID Button.</p>
                <p class="popup-content">2. Search for the question and click Add</p>
                <p class="popup-content">3. Click Submit</p>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 2 -->
    <!-- Pop up 3 -->
    <div class="popup" id="popup-three" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Remove A Question</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Click the Remove button to unassociate the question with this service.</p>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 2 -->

</body>
</html>
