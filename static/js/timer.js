function updateTimer() {
          const timerElement = document.getElementById("timer");
          let seconds = parseInt(timerElement.innerText);
          
          if (seconds > 0) {
            seconds--;
            timerElement.innerText = seconds;
          } else {
            alert("Le temps est écoulé !");
            clearInterval(countdownInterval);
            document.getElementById("myForm").submit();
          }
}
    
const countdownInterval = setInterval(updateTimer, 1000);