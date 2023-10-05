var Elnote = document.getElementById("Elnote");

var IDCorrect = replaceHashWithT(Elnote.getAttribute("value"));

var CorrectNote = document.getElementById(IDCorrect);
var t = document.getElementById("4D")
console.log(t)
illuminateNote(CorrectNote);

/* for (var i = 0; i < Blanche.length; i++) {
    
    Blanche[i].onmouseenter = function(e){
        Blanche[i].style.boxShadow = "inset 0 0 20px 5px rgb(255, 115, 0)";
        console.log("okokok")
        };
}
for (var i = 0; i < Noir.length; i++) {
Noir[i].addEventListener("mouseenter", function(e){
    e.target.style.boxShadow = "inset 0 0 20px 5px rgb(255, 115, 0)";
}); 
}*/
function illuminateNote(object){
    console.log(object)
object.style.background = "red"
}

function replaceHashWithT(inputString) {
    var modifiedString = inputString.replace(/#/g, 'T');
    return modifiedString;
}