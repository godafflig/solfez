function updateTimer() {
          const timerElement = document.getElementById("timer");
          let seconds = parseInt(timerElement.innerText);
          
          if (seconds > 0) {
            seconds--;
            timerElement.innerText = seconds;
          } else {
            alert("Le temps est écoulé !");
            document.getElementById("myForm").submit();

            //window.location.reload();
          }
}
    
const countdownInterval = setInterval(updateTimer, 1000);
    
setTimeout(() => {
          clearInterval(countdownInterval);
}, 30000);