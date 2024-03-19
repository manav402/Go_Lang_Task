let elem = document.getElementById("pageno")
let no = parseInt(elem.innerText)
let next = no + 1
let prev = (no - 1) < 0 ? 0 : no - 1
document.getElementById("next").setAttribute("href", "/get/" + next)
document.getElementById("prev").setAttribute("href", "/get/" + prev)