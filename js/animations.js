const NUM_ANIM_ITEMS = 7;
const MEDIAQUERY = window.matchMedia( "(min-width: 900px)" );
var timeout = 1000;
var animCount = 1;
var elem;
var scrolledToBottom = false;


function fadeAnimation()
{
    switch(animCount){
        case 1:
            elem = document.getElementById('title');
            break;
        case 2:
            elem = document.getElementById('text_username');
            break;
        case 3:
            elem = document.getElementById('uname');
            break;
        case 4:
            elem = document.getElementById('text_password');
            break;
        case 5:
            elem = document.getElementById('pw');
            break;
        case 6:
            elem = document.getElementById('submit_button');
            break;
        default:
            return;
    }
   
    elem.classList.add("start_fade_and_slide_in");
    elem.style.visibility = 'visible';
    animCount++;
}

window.onbeforeunload = function () {
    document.getElementById('left_container').style.visibility = 'hidden';
    window.scrollTo(0, 0);
}

window.onload = function() {
    
    elem = document.getElementById('left_container');

    if(MEDIAQUERY.matches)
    {
        //If bigger than 900px

        //begin initial animation
        elem.classList.add("start_slide_in");
        elem.style.visibility = 'visible';

    }
    else
    {
        //If smaller than 900px

        //begin initial animation
        elem.classList.add("start_slide_in_small_screen");
        elem.style.visibility = 'visible';
        
        //set timeout to allow time for auto scroll down;
        timeout = 2000;

        //automatically scroll the page
        setTimeout(pageScroll, 1800);

    }
    
    //animate login items
    for(var i = 1; i < NUM_ANIM_ITEMS; i++) //start at 1 because 1 animation already played.
        {
            setTimeout(fadeAnimation, timeout);
            timeout = timeout + 100;
        }

}


function pageScroll() {
        //scroll the page down
        window.scrollBy(0,5);

        //checks if the scroll bar has hit the bottom
        window.onscroll = function(ev) {
        if ((window.innerHeight + window.scrollY) >= document.body.scrollHeight - 2) {
               scrolledToBottom = true;
            }
        };

        //stop scrolling if it has reached the bottom
        if(!scrolledToBottom)
        {
           scrolldelay = setTimeout(pageScroll,5);
        }

}
