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

        <div id="succ_msg">{{.M}}</div>

        <form id="form" action="/newserv" method="post" style="{{.A}};">
            <p>Name</p>
            <input name="name" value="{{.O.Pname}}" required>
            <p>Description</p>
            <input name="description" value="{{.O.Pdescription}}" spellcheck="true" required>
            <p>Website URL <span id="url_format">(http://)</span></p>
            <input name="url" type="url" value="{{.O.Purl}}">
            <input name="nskey" type="hidden" value="{{.O.Pkey}}"><br>
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
                    <h2>Questions<a href="#popup-four"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h2>
                </div>

                {{if .Assoc}}
                <form id="form" action="/newserv" style="{{.A}};">
                    <h3>Associate Question To Service<a href="#popup-two"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h3>
                    <p>Question ID</p>
                    <input name="nskey" type="hidden" value="{{.O.Pkey}}">
                    <input name="nprompt" type="number" value="{{.Qid}}" required><br>
                    <input id="submit_button" value="Submit" type=submit>
                </form>
                <form id="qid_form" action="/searchqid" method="post" style="{{.A}};">
                        <input name="skey" type="hidden" value="{{.O.Pkey}}">
                        <input id="submit_button_qid" type='submit' value="Search For Question ID">
                </form>
                {{else}}
                <a href="/newq"><div id="newq">Create A New Question</div></a>

                <form id="form" action="/newserv" style="{{.A}}; margin-top: 20px;">
                    <input type="hidden" name="nskey" value="{{.O.Pkey}}">
                    <input type="hidden" name="assoc" value="true">
                    <input type="submit" id="submit_button_assoc" value="Associate A Question">
                </form>
                {{end}}
                <hr>

                {{if .QList}} <!--If there ARE associated questions, show them -->
                <h3>Question List<a href="#popup-three"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h3>
                <div id="question_list" class="remove">
                    {{$Animation := .A}}
                    {{$ServiceKey := .O.Pkey}}
                    {{range .O.Pqlist}}
                    <form id="form" action="delq" method="post" style="{{$Animation}};">
                        <div id="prompt">{{.Pprompt}}</div>
                        <div id="qid">ID: {{.Pkey}}</div>
                        <input name="ikey" type="hidden" value="{{.Pkey}}">
                        <input name="nskey" type="hidden" value="{{$ServiceKey}}">
                        <input id="submit_button" value="Remove" type=submit>
                    </form>
                    <form id="form" action="/newserv" method="post" style="{{$Animation}}; margin-top: 0px;">
                        <input name="questionid" value="{{.Pkey}}" type="hidden">
                        <input name="nskey" type="hidden" value="{{$ServiceKey}}">
                        <input id="submit_button_assoc" value="Add A Criterion" type="submit">
                    </form>
                    {{end}}
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
                <form id="form" action="/newserv" method="post" style="{{.A}};">
                    <h3>Add Criterion To Question: <br>
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
                        <p>Lower Limit<a href="#popup-six"><img class="popup_icon" src="imgs/popup_icon.png"/></a></p>
                        <input type="number" name="aval" required>
                        <p>Upper Limit<a href="#popup-seven"><img class="popup_icon" src="imgs/popup_icon.png"/></a></p>
                        <input type="number" name="bval" required>
                        <p>Invert Range<a href="#popup-eight"><img class="popup_icon" src="imgs/popup_icon.png"/></a></p>
                        <input type="checkbox" class="checkbox" value="true" name="inv">

                        {{else}} <!--Would then have to be a text type -->
                        <p>Text Answer</p>
                        <input name="regex" spellcheck="true" required>
                        {{end}}<!--End .NumQ-->
                    {{end}}<!--End .BoolQ -->

                    <br>
                    <p>Conjunctive<a href="#popup-nine"><img class="popup_icon" src="imgs/popup_icon.png"/></a></p>
                    <input type="checkbox" name="conj" value="true" class="checkbox" checked>
                    <input name="nskey" type="hidden" value="{{.O.Pkey}}"><br>
                    <input id="submit_button" value="Submit" type=submit>
                </form><hr>
                {{end}} <!--End .C -->

                {{if .Nec}} <!--if there ARE criterion show this form-->
                <h3>Remove Criterion From Service<a href="#popup-ten"><img class="popup_icon" src="imgs/popup_icon.png"/></a></h3>
                {{$Animation2 := .A}}
                <div class="remove">
                    {{range .O.Pclist}}
                    <form id="form" action="delc" method="post" style="{{$Animation2}};">
                        <div id="criterion">Criterion For Question With ID: {{.Qkey}}</div>
                        <div id="value">Value: {{.Pvalue}}</div>
                        <input name="nckey" type="hidden" value="{{.Pkey}}"><br>
                        <input id="submit_button" value="Delete" type=submit>
                    </form>
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
    <!-- Pop up 3 Remove A Question-->
    <div class="popup" id="popup-three" aria-hidden="true">
        <div class="popup-dialog">
            <div class="popup-header">
                <h2>Question List</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Click the <b><em>Remove</em></b> button to unassociate the question with this service.</p><br>
                <p class="popup-content">Click the <b><em>Add A Criterion</em></b> button to create the criteria for this question in order to qualify for this service.</p>
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
                <p class="popup-content">Click the <b><em>Create A New Question</em></b> button to add a brand new question to the system. You then must go back and search for your service to return here.</p><br>
                <p class="popup-content">Click the <b><em>Associate A Question</em></b> button to associate an existing question to this service.</p>
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
                <h2>Lower Limit</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Enter the lowest value that would allow a person to qualify. If the answer must be an exact value. Enter the same number in the <b><em>Upper Limit</em></b>.</p><br>
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
                <h2>Upper Limit</h2>
                <a href="#close" class="btn-close" aria-hidden="true">×</a>
            </div>
            <div class="popup-body">
                <p class="popup-content">Enter the highest value that would allow a person to qualify. If the answer must be an exact value. Enter the same number as the <b><em>Lower Limit</em></b>.</p><br>
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
                <h2>Conjunctive</h2>
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


</body>
</html>
