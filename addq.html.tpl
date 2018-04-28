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
            <input name="nqkey" type="hidden" value="{{.O.Pkey}}">
            <input name="nskey" type="hidden" value="{{.S}}">
            <!-- The last input (with name editing) is only used for deciding which message to display by what the title is -->
            <input name="editing" type="hidden" value="{{.T}}"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form>

        {{if .E}}<!-- Show if editing a question -->

        {{if not .B}}<!-- If its a bool, there is no adding criteria option -->
        <hr>
        <div id="title" style="{{.A}};">
            <h2>Validation Criteria</h2>
        </div>

        <form id="form" action="/newq" method="post" style="{{.A}};">
            <h3>Add Possible Answer</h3>
            {{if .N}}
                <p>Lower Limit</p>
                <input type="number" name="aval" required>
                <p>Upper Limit</p>
                <input type="number" name="bval" required>
                <p>Invert Range</p>
                <input type="checkbox" class="checkbox" value="true" name="inv">
                <p>Allow Negatives</p>
                <input type="checkbox" class="checkbox" value="true" name="neg">
                <p>Allow Decimals</p>
                <input type="checkbox" class="checkbox" value="true" name="dec">
                {{else}}
                <p>Text Answers</p>
                <input name="regex" spellcheck="true">
            {{end}}


            <br>
            <p>Conjunctive</p>
            <input type="checkbox" name="conj" value="true" class="checkbox">
            <input name="nskey" type="hidden" value="{{.S}}">
            <input type="hidden" name="qkey" value="{{.O.Pkey}}"><!-- used to add criterion -->
            <input type="hidden" name="nqkey" value="{{.O.Pkey}}"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form>


        {{if .CList}}
        <hr>

        <div id="title" style="{{.A}}">
            <h3>Delete A Criterion</h3>
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

        {{end}}<!-- End .CList -->

        {{end}}<!-- End .B -->

        {{end}}<!--End .E -->


    </div>

</body>
</html>
