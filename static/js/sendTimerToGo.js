const elapsedTime = localStorage.getItem("elapsedTime");


fetch(`/getLocalStorage?elapsedTime=${elapsedTime}`)
    .then(response => response.text())
    .then(data => {
      console.log(data);
    })
    .catch(error => {
      console.error("Erreur lors de la communication avec le serveur : ", error);
    });