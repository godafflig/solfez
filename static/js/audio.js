  // JavaScript pour gérer les interactions avec l'audio
  const audioElement = document.getElementById("audio-element");
  const playButton = document.getElementById("play-button");
  const pauseButton = document.getElementById("pause-button");

  playButton.addEventListener("click", () => {
      audioElement.play();
  });

  pauseButton.addEventListener("click", () => {
      audioElement.pause();
  });

  audioElement.addEventListener("play", () => {
      playButton.disabled = true;
      pauseButton.disabled = false;
  });

  audioElement.addEventListener("pause", () => {
      playButton.disabled = false;
      pauseButton.disabled = true;
  });
  // JavaScript
// JavaScript
document.addEventListener("DOMContentLoaded", function () {
    const radioButtons = document.querySelectorAll(".input-note");

    radioButtons.forEach(function (radioButton) {
        radioButton.addEventListener("change", function () {
            // Supprimez la classe 'selected' de tous les boutons radio
            radioButtons.forEach(function (button) {
                button.parentElement.classList.remove("selected");
                button.parentElement.classList.add("btn-note");

            });

            // Vérifiez si le bouton radio est coché
            if (this.checked) {
                // Ajoutez la classe 'selected' au bouton radio coché
                const selectedValue = this.value;
                // Vous pouvez maintenant interagir avec la valeur sélectionnée
                this.parentElement.classList.add("selected");
                this.parentElement.classList.remove("btn-note");
            }
        });
    });
});

