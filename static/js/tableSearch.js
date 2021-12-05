function search(){
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