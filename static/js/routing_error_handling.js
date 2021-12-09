const url = new URL(window.location.href);
const error = url.searchParams.get("error");
const type = url.searchParams.get("type");
const success = url.searchParams.get("success");

switch (error) {
    case "authentication":
        if (type == "login") {
            loginError();
        }
        break;
    case "login":
        loginError();
        break;
    case "password":
        passwordError();
        break;
    case "user":

        break;
    default:
        break;
}

switch (success) {
    case "application":
        applicationSuccess();
        break;
    case "passwordChange":

        break;
}

function loginError() {
    let loginModal = document.getElementById("loginModal");
    let loginModalInstance = bootstrap.Modal.getOrCreateInstance(loginModal);
    let loginModalLabel = document.getElementById("loginModalLabel");
    loginModalLabel.classList.add("text-danger");
    switch (type) {
        case "username":
            loginModalLabel.innerHTML = "username ist falsch";
            break;
        case "password":
            loginModalLabel.innerHTML = "Passwort wurde falsch eingegeben";
            break;
    }
    loginModalInstance.show();
}

function passwordError() {
    let content = document.getElementById("content");
    let msg;
    let div = document.createElement("div");
    switch (type) {
        case "old":
            msg = "altes Passwort ist inkorrekt"
            break;
        case "new":
            msg = "neues Passwort wurde falsch wiederholt"
            break;
        default:
            break;
    }
    div.innerHTML = msg;
    div.classList.add("text-danger");
    content.insertBefore(div, content.firstChild);

}

function applicationSuccess() {
    let myToastEl = document.getElementById('applicationSuccessToast')
    let myToast = bootstrap.Toast.getOrCreateInstance(myToastEl)
    myToast.show();
}