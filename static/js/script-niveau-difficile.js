var Elnote = document.getElementById("Elnote");

var IDCorrect = replaceHashWithT(Elnote.getAttribute("value"));

var CorrectNote = document.getElementById(IDCorrect);

function replaceHashWithT(inputString) {
    var modifiedString = inputString.replace(/#/g, 'T');
    return modifiedString;
}

var selectedNotes = [];

var pianoKeys = document.querySelectorAll('.NoteDefault');
var sharpKeys = document.querySelectorAll('.NoteDieseDefault');

function illuminateKey(key) {
          var index = selectedNotes.indexOf(key.getAttribute("id"));
          if (index > -1) {
              selectedNotes.splice(index, 1);
              key.style.background = "";
          } else {
              if (selectedNotes.length < 3) {
                  selectedNotes.push(key.getAttribute("id"));
                  key.style.background = "red";
              }
          }
          var submit = document.getElementById("valider-btn");
          submit.setAttribute("name", "answer");
          submit.setAttribute("value", selectedNotes.map(convertToFrench).join(', '));
      }
    
      pianoKeys.forEach(function(key) {
          key.addEventListener('click', function(e) {
              if (!key.getAttribute("id").includes("T")) {
                  e.stopPropagation();
                  illuminateKey(key);
              }
          });
      });
      
      sharpKeys.forEach(function(key) {
          key.addEventListener('click', function(e) {
              e.stopPropagation();
              illuminateKey(key);
          });
      });


function convertToFrench(note){
  const noteMap = {
    "4C": "Do4eme",
    "4CT": "Do#4eme",
    "4D": "Ré4eme",
    "4DT": "Ré#4eme",
    "4E": "Mi4eme",
    "4F": "Fa4eme",
    "4FT": "Fa#4eme",
    "4G": "Sol4eme",
    "4GT": "Sol#4eme",
    "4A": "La4eme",
    "4AT": "La#4eme",
    "4B": "Si4eme",
    "5C": "Do5eme",
    "5CT": "Do#5eme",
    "5D": "Ré5eme",
    "5DT": "Ré#5eme",
    "5E": "Mi5eme",
    "5F": "Fa5eme",
    "5FT": "Fa#5eme",
    "5G": "Sol5eme",
    "5GT": "Sol#5eme",
    "5A": "La5eme",
    "5AT": "La#5eme",
    "5B": "Si5eme"
  };

  const noteFrancaise = noteMap[note];
  return noteFrancaise;  
}