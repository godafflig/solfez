var Elnote = document.getElementById("Elnote");

var IDCorrect = replaceHashWithT(Elnote.getAttribute("value"));

var CorrectNote = document.getElementById(IDCorrect);
var t = document.getElementById("4D")
illuminateNote(CorrectNote);

function illuminateNote(object){
object.style.background = "red"
}

function replaceHashWithT(inputString) {
    var modifiedString = inputString.replace(/#/g, 'T');
    return modifiedString;
}