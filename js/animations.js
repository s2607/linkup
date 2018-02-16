const NUM_ANIM_ITEMS = 7;
var timeout = 1000;
var animCount = 1;
var elem;


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


window.onload = function() {
    
    elem = document.getElementById('left_container');
    elem.classList.add("start_slide_in");
    elem.style.visibility = 'visible';

    for(var i = 1; i < NUM_ANIM_ITEMS; i++)
    {
        setTimeout(fadeAnimation, timeout);
        timeout = timeout + 100;
    }
        
    
}
