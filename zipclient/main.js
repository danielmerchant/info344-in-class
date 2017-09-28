var submitButton = document.getElementById("submit");
var inputBox = document.getElementById("input-field");
var inputResult = document.getElementById("name-result");

submitButton.addEventListener("click", displayHello);

function displayHello() {
    var name = inputBox.value;
    fetch("http://localhost:4000/hello?name=" + name).then(function(response) {
        return response.text();
    }).then(function(myBlob) {
        inputResult.innerHTML = myBlob;
    });
}