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
        <div id="top_bar">
            <img id="logo" src="imgs/logo.png" alt="LinkUp">
            <a href="/home"><div id="home_button" style="animation: fade_in; {{.A}}; animation-duration: 1s; animation-timing-function: ease-in-out;">Home</div></a>
        </div>

        <div id="title" style="{{.A}};">
            <h1>{{.T}} Question</h1>
        </div>

        <div id="succ_msg">{{.M}}</div>

        <form id="form" action="/newq" method="post" style="{{.A}};">
            <!--TODO:oldvals-->
            <p>Question Prompt</p>
            <input name="prompt" value="{{.O.Pprompt}}" spellcheck="true" required>
            <p>Answer Type </p>
            <select name="qtype" required style="width: 100px;">
                <option></option>
                <option value="0">Text</option>
                <option value="1">Number</option>
                <option value="2">Yes/No</option>
            </select>
            <br>
            <input id="submit_button" value="Submit" type=submit>
        </form><hr>

        <div id="title" style="{{.A}};">
            <h2>Validation Criteria</h2>
        </div>

        <form id="form" action="/newq" method="post" style="{{.A}};">
            <h3>Add Possible Answer</h3>
            <p>Text Answer</p>
            <input name="regex" spellcheck="true">
            <p>Lower Limit</p>
            <input type="number" name="aval">
            <p>Upper Limit</p>
            <input type="number" name="bval">
            <p>Yes/No Questions <br><select name="lval">
                <option></option>
                <option value="1">Yes</option>
                <option value="0">No</option>
            </select></p><br>
            <p>isnil:</p>
            <input type="checkbox" value="1" class="checkbox" name="isnil">
            <p>Invert</p>
            <input type="checkbox" class="checkbox" value="1" name="inv">
            <p>Conjunctive</p>
            <input type="checkbox" name="conj" value="1" class="checkbox">
            <input type="hidden" name="nqkey" value="{{.O.Pkey}}"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form>
        <hr>

        <div id="title" style="{{.A}}">
            <h3>Delete A Criterion</h3>
        </div>

        {{$Animation := .A}}
        {{range .O.Pclist}}
	   <form id="form" action="delc" method="post" style="{{$Animation}};">
           <input name="nckey" type="hidden" value="{{.Pkey}}"><br>
           <input id="submit_button" value="Delete" type=submit>
        </form><hr>
        {{end}}



    </div>

</body>
</html>
