fetch('navbar.html')
    .then(res => res.text())
    .then(text => {
        let oldelem = document.querySelector("script#replace_with_navbar");
        let newelem = document.createElement("div");
        newelem.innerHTML = text;
        oldelem.parentNode.replaceChild(newelem, oldelem);
    })

// Créer un élément script
var script2 = document.createElement("script");

// Spécifier le chemin vers script2.js
script2.src = "../static/js/navbar.js";

// Attacher un gestionnaire d'événement pour gérer le chargement du script
script2.onload = function () {
    console.log("script2.js a été chargé avec succès.");
};

// Ajouter le script au DOM
document.head.appendChild(script2);
const links = document.querySelectorAll('nav li')

icons.addEventListener("click", () => {
    nav.classList.toggle("active");
})

links.forEach((links) => {
    links.addEventListener("click", () => { nav.classList.remove('active'); })
})