// on page load set hyper link for next and prev button
let elem = document.getElementById("pageno")
let no = parseInt(elem.innerText)
let next = no + 1
let prev = (no - 1) < 0 ? 0 : no - 1
document.getElementById("next").setAttribute("href", "/get/" + next)
document.getElementById("prev").setAttribute("href", "/get/" + prev)

// if api response is valid populate the tbody with the data
let tbody = document.querySelector("tbody")
function tableCreator(data) {
    let tags = ""
    data.forEach(d => {
        tags += `<tr><td>${d["no."]}</td><td>${d["train no."]}</td><td>${d["train name"]}</td><td>${d.starts}</td><td>${d.ends}</td></tr>`
    });
    tbody.innerHTML = tags
}

let oldData = tbody.innerHTML

// request function send and update table based on received json data
function request() {
    isRunning = false
    tbody.innerHTML = "" //cleaning table body
    let query = document.querySelector("#search").value
    console.log(query)
    fetch("http://localhost:8000/search", {
        method: "POST",
        body: JSON.stringify({
            "query": query
        }),
        headers: {
            "Content-Type": "application/json"
        }
    }).then((res) => res.json()).then((data) => {
        if (data) {
            tableCreator(data)
        } else {
            tbody.innerHTML = `<tr><td colspan=5>No Data Found...</td></tr>`
        }
    })
}

// on input from inputbox search in the database for appropriete data
let inputBox = document.getElementById("search")
let isRunning = false
let timeOutId
inputBox.addEventListener("input", (e) => {
    if (e.target.value.length == 0) {
        tbody.innerHTML = oldData
        clearTimeout(timeOutId)
        return
    }
    if (isRunning) {
        clearTimeout(timeOutId)
    }
    isRunning = true
    timeOutId = setTimeout(request, 500)
})

