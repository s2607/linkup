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
            <a href="/home"><div id="home_button" style="animation: fade_in; {{.A}}; animation-duration: 1s; animation-timing-function: ease-in-out;">Home</div></a>
        </div>

        <div id="title" style="{{.A}};">
            <h1>{{.T}} Question</h1>
        </div>

        <div id="succ_msg">{{.M}}</div>

        {{if .Back}} <!-- if they came directly from the service, show this -->
        <form id="form" action="/newserv" method="post" style="animation: none; margin-top: 10px;">
            <input name="nskey" type="hidden" value="{{.S}}">
            <input name="nqkey" type="hidden" value="{{.O.Pkey}}">
            <input id="submit_button_serv" value="Back To Service Program" type=submit>
        </form>
        {{end}}

        <form id="form" action="/newq" method="post" style="{{.A}};">
            <p>Question Prompt</p>
            <input name="prompt" value="{{.O.Pprompt}}" spellcheck="true" required>
            {{if not .E}}<!-- hide type editing-->
            <p>Answer Type </p>
            <select name="qtype" required style="width: 100px;">
                <option></option>
                <option value="0">Text</option>
                <option value="1">Number</option>
                <option value="2">Yes/No</option>
            </select>
            {{end}}
            {{if .E}}<!-- retains type when editing question -->
            <input name="qtype" type="hidden" value="{{.O.Ptype}}">
            {{end}}
            <input name="nqkey" type="hidden" value="{{.O.Pkey}}">
            <input name="nskey" type="hidden" value="{{.S}}">
            <!-- The last input (with name editing) is only used for deciding which message to display by what the title is -->
            <input name="editing" type="hidden" value="{{.T}}"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form>

        {{if .E}}<!-- Show if editing a question -->

        {{if not .B}}<!-- If its a bool, there is no adding criteria option -->
        <!-- Make single column layout appear -->
        <style>
            #left_container
            {
                width: 100%;
            }
        </style>

        <div id="sub_container">

            <div id="left_container">

                <div id="title" style="{{.A}};">
                    <h2>Validation Criteria</h2>
                </div>

                <form id="form" action="/newq" method="post" style="{{.A}};">
                <h3>Add Possible Answer</h3>
                {{if .N}}
                    <p>Lower Limit<a href="#popup-two"><img class="popup_icon" src="imgs/popup_icon.png"/></a></p>
                    <input type="number" name="aval" required>
                    <p>Upper Limit<a href="#popup-three"><img class="popup_icon" src="imgs/popup_icon.png"/></a></p>
                    <input type="number" name="bval" required>
                    <p>Invert Range<a href="#popup-four"><img class="popup_icon" src="imgs/popup_icon.png"/></a></p>
                    <input type="checkbox" class="checkbox" value="true" name="inv">
                    <p>Exclusive Range<a href="#popup-six"><img class="popup_icon" src="imgs/popup_icon.png"/></a></p>
                    <input type="checkbox" class="checkbox" value="true" name="exc">
                    <p>Only Allow Positives</p>
                    <input type="checkbox" class="checkbox" value="true" name="pos">
                    <p>No Decimals</p>
                    <input type="checkbox" class="checkbox" value="true" name="dec">
                    {{else}}
                    <p>Text Answers<a href="#popup-one"><img class="popup_icon" src="imgs/popup_icon.png"/></a></p>
                    <input name="regex" spellcheck="true">
                {{end}}


                <br>

                <input name="nskey" type="hidden" value="{{.S}}">
                <input type="hidden" name="qkey" value="{{.O.Pkey}}"><!-- used to add criterion -->
                <input type="hidden" name="nqkey" value="{{.O.Pkey}}"><br>
                <input id="submit_button" value="Submit" type=submit>
                </form>
            </div>

        {{if .CList}}
            <!-- Make single column layout appear -->
        <style>
            #left_container
            {
                width: 50%;
            }
        </style>
            <div id="right_container">

                <div id="title" style="{{.A}}">
                    <h3>Delete A Criterion<a href="#popup-five"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h3>
                </div>

                {{$Animation := .A}}
                {{$QKey := .O.Pkey}}
                {{$Question := .O}}
                {{$SKey := .S}}
                {{range .O.Pclist}}
	           <form id="form" action="delqc" method="post" style="{{$Animation}};">
                   <p style="margin-bottom: 0px;">Criterion Value: {{$Question.Pvalue .}}</p>
                   <input name="nskey" type="hidden" value="{{$SKey}}">
                   <input name="nqkey" type="hidden" value="{{$QKey}}">
                   <input name="nckey" type="hidden" value="{{.Pkey}}"><br>
                   <input id="submit_button" value="Delete" type="submit">
                </form><hr>
                {{end}}<!-- End Range -->
            </div>
        {{end}}<!-- End .CList -->

        </div>

        {{end}}<!-- End .B -->

        {{end}}<!--End .E -->


    </div><!--End container -->

    <!-- Pop up 1 Text Answer-->
    <div class="popup" id="popup-one" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Text Answers</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Enter all valid answers into this text box. Answers are case <em>insensitive</em>. If there are multiple options for answers, separate them with a | character.</p><br>
                <p class="popup-content">Ex: male|female|other is valid input for having options for answers of male, female, or other with any capitalizations.</p><br>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 1 -->
    <!-- Pop up 2 Lower Limit-->
    <div class="popup" id="popup-two" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Lower Limit</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Enter the lowest value that is allowed as response for this question. If the answer must be an exact value. Enter the same number in the <b><em>Upper Limit</em></b>.</p><br>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 2 -->
    <!-- Pop up 3 Upper Limit-->
    <div class="popup" id="popup-three" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Upper Limit</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Enter the highest value that is allowed as response for this question. If the answer must be an exact value. Enter the same number as the <b><em>Lower Limit</em></b>.</p><br>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 3 -->
    <!-- Pop up 4 Invert Range-->
    <div class="popup" id="popup-four" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Invert Range</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Check this box to make the valid values to be everything outside the given range.</p><br>
                <p class="popup-content">Ex: A range of 5-10 would become anything less than 5 or greater than 10.</p><br>
                <p class="popup-content"><em>Special Note:</em> An <b>inverted range</b> that is also an <b>exclusive range</b> would have a valid range of anything 5 or less or anything 10 or more. An entry of 5 would be valid.</p><br>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 4 -->
    <!-- Pop up 5 Remove A Criterion-->
    <div class="popup" id="popup-five" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Delete A Validation Criterion</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Click the <b><em>Delete</em></b> button to delete the valid response from this question.</p><br>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 5 -->
    <!-- Pop up 6 Exclusive Range-->
    <div class="popup" id="popup-six" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Exclusive Range</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Check this box to make the valid values exclude the upper and lower limits of the range. </p><br>
                <p class="popup-content">Ex: A range of 5-10 would not include 5 nor 10 but be anything in between them. An entry of 5 would be considered invalid.</p><br>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 6 -->


</body>
</html>
