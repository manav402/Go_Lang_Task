let elem = document.getElementById("pageno")
let no = parseInt(elem.innerText)
let next = no + 1
let prev = (no - 1) < 0 ? 0 : no - 1
document.getElementById("next").setAttribute("href", "/get/" + next)
document.getElementById("prev").setAttribute("href", "/get/" + prev)

var table  = document.querySelector("tbody")
function tableCreator(data){
    console.log(data)
    // var tr = document.createElement("tr")
    var tags = ""
    data.forEach(d => {
        tags += `<tr><td>${d["no."]}</td><td>${d["train no."]}</td><td>${d["train name"]}</td><td>${d.starts}</td><td>${d.ends}</td></tr>`
    });
    table.innerHTML = tags
}

var inputBox = document.getElementById("search")
var isChanged = false
var tbody = document.querySelector("tbody")
inputBox.addEventListener("input", (d) => {
   tbody.innerHTML = ""
   fetch("http://localhost:8000/search",{
    method:"POST",
    body:JSON.stringify({
        "query":d.target.value
    }),
    headers:{
        "Content-Type":"application/json"
    }
   }).then((res)=>res.json()).then((data)=>{
    // console.log(data)
    if (data){
        tableCreator(data)
    }else{
        tbody.innerHTML = `<tr><td colspan=5>No Data Found...</td></tr>`
    }
   })
})