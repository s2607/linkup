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

        <div id="title" style="{{.A}}">
            <h1>{{.T}} Service Program<a href="#popup-one"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h1>
        </div>

        <div id="succ_msg" {{if .Err}}style="color: red;"{{end}}>{{.M}}</div>

        <form id="form" action="/newserv" method="post" style="{{.A}};">
            <p>Name</p>
            <input name="name" value="{{.O.Pname}}" required>
            <p>Description</p>
            <input name="description" value="{{.O.Pdescription}}" spellcheck="true" required>
            <p>Website URL <span id="url_format">(http://)</span></p>
            <input name="url" type="url" value="{{.O.Purl}}">
            <input name="nskey" type="hidden" value="{{.O.Pkey}}">
             <!-- The last input (with name editing) is only used for deciding which message to display by what the title is -->
            <input name="editing" type="hidden" value="{{.T}}"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form>

        {{if .E}} <!-- If editing service show editing fields -->

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
                    <h2>Questions For {{.O.Pname}}<a href="#popup-four"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h2>
                </div>

                {{if .Assoc}}<!--Show associate a question form if button clicked -->
                <!--Scrolls 1/3 of the way down page to be at top of this form -->
                <script>
                    window.onload = function(){
                        window.scrollTo(0, document.getElementById("sub_container").offsetTop);
                    }
                </script>
                <form id="form" action="/newserv" style="{{.A}};">
                    <h3 style="{{.A}};">Associate Question To Service<a href="#popup-two"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h3>
                    <p>Question ID</p>
                    <input name="nskey" type="hidden" value="{{.O.Pkey}}">
                    <input name="nprompt" type="number" value="{{.Qid}}" readonly required><br>
                    <input id="submit_button" value="Submit" type=submit>
                </form>

                {{else}}
                <form id="qid_form" action="/searchqid" method="post" style="{{.A}};">
                        <input name="skey" type="hidden" value="{{.O.Pkey}}">
                        <input id="submit_button_qid" type='submit' value="Search For Question">
                </form>
                {{end}}


                {{if .QList}} <!--If there ARE associated questions, show them -->
                <hr>
                <h3 style="{{.A}};">Associated Questions<a href="#popup-three"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h3>
                <div id="question_list" class="remove">
                    {{$Animation := .A}}
                    {{$ServiceKey := .O.Pkey}}
                    {{range .O.Pqlist}}
                    <div id="prompt"><b>{{.Pprompt}}</b></div>
                    <div id="qid">ID: {{.Pkey}}</div>
                    <div class="action" tabindex="1">Select An Action <span><img class="arrow" src="imgs/chevron.png"></span>
                    <div class="shown"><!-- This div appears with click of .action div right above -->
                    <form class="form_button" action="/newserv" method="post" style="{{$Animation}};">
                        <input name="questionid" value="{{.Pkey}}" type="hidden">
                        <input name="nskey" type="hidden" value="{{$ServiceKey}}">
                        <input id="submit_button_assoc" value="Add A Criterion" type="submit" style="margin-top: 20px;">
                    </form>
                    <form class="form_button" action="/newq" method="post" style="{{$Animation}}; margin-top: 0px;">
                        <input name="nqkey" value="{{.Pkey}}" type="hidden">
                        <input name="nskey" type="hidden" value="{{$ServiceKey}}">
                        <input name="editfromserv" type="hidden" value="true">
                        <input id="submit_button_assoc" value="Edit Question" type="submit">
                    </form>
                    <form id="form_button" action="delq" method="post" style="{{$Animation}}; margin-top: 0px;">
                        <input name="ikey" type="hidden" value="{{.Pkey}}">
                        <input name="nskey" type="hidden" value="{{$ServiceKey}}">
                        <input id="submit_button_assoc" value="Disassociate" type="submit">
                    </form>

                    </div><!-- End action div -->
                    </div><!-- End shown div-->
                    {{end}}<!--End Range-->
                </div>
                {{end}}<!--End .Qlist -->

            </div><!--End left_container-->

            {{if or .C .Nec}}
            <!-- Make two column layout appear as well as the title-->
            <style>
                #left_container
                {
                    width: 50%;
                }
            </style>

            <div id="right_container">

                <div id="title" style="{{.A}};">
                    <h2>Eligibility Criteria</h2>
                </div>

                {{if .C}}<!--Shows adding criterion form if "Add A Criterion" button clicked -->
                <!--Scrolls 1/3 of the way down page to be at top of this form -->
                <script>
                     window.onload = function(){
                        window.scrollTo(0, document.getElementById("sub_container").offsetTop);
                    }
                </script>
                <form id="form" action="/newserv" method="post" style="{{.A}};">
                    <h3 style="{{.A}};">Add Criterion To Question: <br>
                        {{.Q.Pprompt}}<a href="#popup-five"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h3>
                    <input name="qid" value="{{.Q.Pkey}}" type="hidden">

                    {{if .BoolQ}} <!-- if yes/no question -->
                    <p>Yes or No <br><select name="lval" required>
                        <option></option>
                        <option value="true">Yes</option>
                        <option value="false">No</option>
                    </select></p>
                    {{else}}
                        {{if .NumQ}} <!-- if number type -->
                        <p>Minimum<a href="#popup-six"><img class="popup_icon" src="imgs/popup_icon.png"/></a></p>
                        <input type="number" name="aval" step="any">
                        <p>Maximum<a href="#popup-seven"><img class="popup_icon" src="imgs/popup_icon.png"/></a></p>
                        <input type="number" name="bval" step="any">
                        <p>Only Allow Positives</p>
                        <input type="checkbox" class="checkbox" value="true" name="pos">
                        <p>No Decimals</p>
                        <input type="checkbox" class="checkbox" value="true" name="dec">

                    <!-- ********** THESE FIELDS ARE FULLY IMPLEMENTED AND WORKING BUT ARE NOT BEING
                        *********** SHOWN FOR SIMPLICITY OF USE AT THIS TIME
                        <p>Invert Range<a href="#popup-eight"><img class="popup_icon" src="imgs/popup_icon.png"/></a></p>
                        <input type="checkbox" class="checkbox" value="true" name="inv">
                        <p>Exclusive Range<a href="#popup-twelve"><img class="popup_icon" src="imgs/popup_icon.png"/></a></p>
                        <input type="checkbox" class="checkbox" value="true" name="exc">
                    -->

                        {{else}} <!--Would then have to be a text type -->

                        {{if .EmptyL}}<!--If there is no validation criteria display this -->
                        <br>
                        <p id="error_msg">Edit This Question To Add Validation Criteria Before Adding Eligibility Criteria.</p><br>
                        {{end}}<!-- End .EmptyList -->

                        <p>Select All Eligible Answers<a href="#popup-eleven"><img class="popup_icon" src="imgs/popup_icon.png"/></a></p>

                        <select style="min-width: 150px; height: 100px;" name="regex"  multiple required>
                            <option></option>
                            {{range .List}}
                            <option value="{{.}}">{{.}}</option>
                            {{end}}
                        </select>
                        <br>
                        {{end}}<!--End .NumQ-->
                    {{end}}<!--End .BoolQ -->

                    <br>
                    <p>Make Non-Optional<a href="#popup-nine"><img class="popup_icon" src="imgs/popup_icon.png"/></a></p>
                    <input type="checkbox" name="conj" value="true" class="checkbox" checked>
                    <input name="nskey" type="hidden" value="{{.O.Pkey}}"><br>
                    <input id="submit_button" value="Submit" type=submit>
                </form><hr>
                {{end}} <!--End .C -->

                {{if .Nec}} <!--if there ARE criterion show this form-->
                <h3 style="{{.A}};">Remove Criterion From Service<a href="#popup-ten"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h3>
                {{$Animation2 := .A}}
                {{$ServiceKey2 := .O.Pkey}}
                <div class="remove">
                    {{range .O.Pclist}}
                    <form id="form" action="delc" method="post" style="{{$Animation2}};">
                        <div id="criterion">Criterion For Question With ID: {{.Qkey}}</div>
                        <div id="value"><b>Value:</b> {{.Pvalue}}</div>
                        <input name="nckey" type="hidden" value="{{.Pkey}}">
                        <input name="nskey" type="hidden" value="{{$ServiceKey2}}"><br>
                        <input id="submit_button" value="Remove" type=submit>
                    </form><hr>
                    {{end}}
                </div>
                {{end}}<!--End .Nec -->
            </div><!--End right_container -->
            {{end}} <!--End or .C .Nec -->
        </div><!--End sub_container -->
        {{end}}<!--End of if .E-->
    </div><!-- End container-->



    <!-- Pop up 1 Add/Edit Service-->
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
    <!-- Pop up 2 Associate Question-->
    <div class="popup" id="popup-two" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Associate A Question</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content"><b>Steps:</b></p>
                <p class="popup-content">1. Click the <b><em>Search For Question Button.</em></b></p>
                <p class="popup-content">2. Search for the question and click <b><em>Add</em></b> or if you don't find the question, create a new one/</p>
                <p class="popup-content">3. Click <b><em>Submit</em></b> to associate it with this service.</p>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 2 -->
    <!-- Pop up 3 Remove A Question-->
    <div class="popup" id="popup-three" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Associated Questions</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Using the <b><em>Select An Action</em></b> menu: </p><br />
                <p class="popup-content">Click the <b><em>Add A Criterion</em></b> button to create the criteria for this question in order to qualify for this service.</p><br />
                <p class="popup-content">Click the <b><em>Edit Question</em></b> button to add validation criteria for this question which will serve as all the possibilities that this question could be answered.</p><br />
                <p class="popup-content">Click the <b><em>Disassociate</em></b> button to detach the question from this service.</p><br />
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 3 -->
    <!-- Pop up 4 Questions-->
    <div class="popup" id="popup-four" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Questions</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Click the <b><em>Search For Question</em></b> button to associate questions to this service.</p><br>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 4 -->
    <!-- Pop up 5 Yes/No-->
    <div class="popup" id="popup-five" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Add Criterion To Question</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Input the response that a person must have in order to qualify for this service.</p><br>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 5 -->
    <!-- Pop up 6 Lower Limit-->
    <div class="popup" id="popup-six" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Minimum</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Enter the lowest value that would allow a person to qualify. If the answer must be an exact value. Enter the same number in the <b><em>Maximum</em></b>. Leave the maximum blank if the qualifying range is to be the minimum value or higher.</p><br>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 6 -->
    <!-- Pop up 7 Upper Limit-->
    <div class="popup" id="popup-seven" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Maximum</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Enter the highest value that would allow a person to qualify. If the answer must be an exact value. Enter the same number as the <b><em>Minimum</em></b>. Leave the minimum blank if the qualifying range is to be the maximum value or lower.</p><br>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 7 -->
    <!-- Pop up 8 Invert Range-->
    <div class="popup" id="popup-eight" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Invert Range</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Check this box to make the qualifying values to be everything outside the given range.</p><br>
                <p class="popup-content">Ex: A range of 5-10 would become anything less than 5 or greater than 10.</p><br>
                <p class="popup-content"><em>Special Note:</em> An <b>inverted range</b> that is also an <b>exclusive range</b> would have a valid range of anything 5 or less or anything 10 or more. An entry of 5 would be valid.</p><br>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 8 -->
    <!-- Pop up 9 Conjunctive-->
    <div class="popup" id="popup-nine" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Non-Optional</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Leave this box checked if this criterion must be met in addition to other criterion. Deselect this box if this criterion must be met OR another criterion must be met.</p><br>
                <p class="popup-content">Ex: If a person must BOTH be over 18 AND own a car to qualify, leave this box checked for each of those criteria. But, if a person must EITHER be over 18 OR own a car to qualify, uncheck this box. </p><br>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 9 -->
    <!-- Pop up 10 Remove A Criterion-->
    <div class="popup" id="popup-ten" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Remove A Criterion</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Click the <b><em>Remove</em></b> button to delete the criterion from its question.</p><br>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 10 -->
    <!-- Pop up 11 Text Answer-->
    <div class="popup" id="popup-eleven" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Select Answers</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Select all answers from the options that would allow a person to qualify for this service. If the option does not appear in the list, edit the question and add it as validation criteria first.</p><br>
                <p class="popup-content">To select multiple options:  Press ctrl + click for Windows or command + click for Mac.</p><br>
            </div>
            <div class="popup-footer">
                <a href="#close" class="btn">Close</a>
            </div>
        </div>
    </div>
    <!-- End Pop up 11 -->
    <!-- Pop up 12 Exclusive Range-->
    <div class="popup" id="popup-twelve" aria-hidden="true">
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
    <!-- End Pop up 12 -->


</body>
</html>
