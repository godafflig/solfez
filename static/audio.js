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