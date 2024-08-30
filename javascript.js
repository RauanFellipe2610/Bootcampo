function loadHTML(id, url) {
    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState === 4 && this.status === 200) {
            document.getElementById(id).innerHTML = this.responseText;
        }
    };
    xhttp.open("GET", url, true);
    xhttp.send();
}

loadHTML("header", "header.html");
loadHTML("main", "main.html");
loadHTML("footer", "footer.html");
