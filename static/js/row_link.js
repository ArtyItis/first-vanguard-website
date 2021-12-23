function openApplication(id) {
    window.location.replace("/applications/"+id);
}

function openUser(id,company){
    window.location.replace("/companies/"+company+"/members/"+id);
}