var url = new URL(window.location.href);
loginError();
passwordChangeError();
applicationSuccess();

function loginError() {
    var loginError = url.searchParams.get("loginError");
    if (loginError != null) {
        var loginModal = document.getElementById("loginModal");
        var modal = bootstrap.Modal.getOrCreateInstance(loginModal)
        var loginModalLabel = document.getElementById("loginModalLabel");
        if (loginError == "password") {
            loginModalLabel.classList.add("text-danger");
            loginModalLabel.innerHTML = "Passwort wurde falsch eingegeben";
        } else if (loginError == "username") {
            loginModalLabel.classList.add("text-danger");
            loginModalLabel.innerHTML = "username ist falsch";
        }
        modal.show();
    }
}

function passwordChangeError() {
    var passwordChangeError = url.searchParams.get("changePasswordError");
    if (passwordChangeError != null) {
        var content = document.getElementById("content");
        var msg;
        var div = document.createElement("div");
        if (passwordChangeError == "newPassword") {
            msg = "neues Passwort wurde falsch wiederholt"
        } else if (passwordChangeError == "oldPassword") {
            msg = "altes Passwort ist inkorrekt"
        }
        div.innerHTML = msg;
        div.classList.add("text-danger");
        content.insertBefore(div, content.firstChild);
    }
}

function applicationSuccess() {
    var applicationSuccess = url.searchParams.get("application");
    if (applicationSuccess == "success") {
        var myToastEl = document.getElementById('applicationSuccessToast')
        var myToast = bootstrap.Toast.getOrCreateInstance(myToastEl)
        myToast.show();
    }
}