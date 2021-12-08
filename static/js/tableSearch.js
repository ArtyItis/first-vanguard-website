function searchByName() {
    let input = document.getElementById("search");
    let table = document.getElementById("table");
    let rows = table.getElementsByTagName("tr");

    let name;

    for (let i = 0; i < rows.length; i++) {
        name = rows[i].getElementsByTagName("td")[0];
        if (name) {
            txtValue = name.textContent || name.innerText;
            if (txtValue.indexOf(input.value) > -1) {
                rows[i].style.display = "";
            } else {
                rows[i].style.display = "none";
            }
        }
    }
}

function filterByCompany(input) {
    let table = document.getElementById("table");
    let rows = table.getElementsByTagName("tr");

    let company;
    for (let i = 0; i < rows.length; i++) {
        company = rows[i].getElementsByTagName("td")[2];
        if (company) {
            txtValue = company.textContent || company.innerText;
            if (txtValue == input) {
                rows[i].style.display = "";
            } else {
                rows[i].style.display = "none";
            }
        }
    }
}