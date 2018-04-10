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
        </div>

        <div id="title">
            <h1>Add Question</h1>
        </div>

        <form id="form" action="/newq" method="post">
            <!--TODO:oldvals-->
            <p>Question Prompt</p>
            <input name="prompt" value="{{.O.Pprompt}}" spellcheck="true">
            <p>Answer Type <br><span>
                <input type="radio" name="qtype" value="0">Text
                <input type="radio" name="qtype" value="1">Number
                <input type="radio" name="qtype" value="boolean">Any More?
                </span>
            </p><br>
            <input id="submit_button" value="Submit" type=submit>
        </form><hr>

        <div id="title">
            <h2>Criteria</h2>
        </div>

        <form id="form" action="newq" method="post">
            <h3>Add A Criterion</h3>
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
            <p>inv</p>
            <input type="checkbox" class="checkbox" value="1" name="inv">
            <p>conj</p>
            <input type="checkbox" name="conj" value="1" class="checkbox">
            <input type="hidden" name="nqkey" value="{{.O.Pkey}}"><br>
            <input id="submit_button" value="Submit" type=submit>
        </form>
        <hr>

        {{range .O.Pclist}}
	   <form id="form" action="delc" method="post">
           <p>Delete A Criterion</p>
           <input name="nckey" value="{{.Pkey}}"><br>
           <input id="submit_button" value="Submit" type=submit>
        </form><hr>
        {{end}}



    </div>

</body>
</html>
