const links = document.querySelectorAll('nav li')

icons.addEventListener("click", () => {
    nav.classList.toggle("active");
})

links.forEach((links) => {
    links.addEventListener("click", () => { nav.classList.remove('active'); })
})