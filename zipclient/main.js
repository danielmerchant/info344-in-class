var submitButton = document.getElementById("submit");
var inputBox = document.getElementById("input-field");
var inputResult = document.getElementById("name-result");
var memoryResult = document.getElementById("memory-result");

submitButton.addEventListener("click", displayHello);
window.setInterval(memoryCheck, 1000);

function displayHello() {
    var name = inputBox.value;
    fetch("http://localhost:4000/hello?name=" + name).then(function(response) {
        return response.text();
    }).then(function(myText) {
        inputResult.innerHTML = myText;
    });
}

function memoryCheck() {
    memoryResult.innerHTML = "";
    fetch("http://localhost:4000/memory").then(function(response) {
        return response.json();
    }).then(function(myJson) {
        memoryResult.innerHTML = myJson["Alloc"];
    })
}