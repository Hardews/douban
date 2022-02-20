function getuser(username) {
    let httpRequest = new XMLHttpRequest();
    httpRequest.open("GET", '接口' + "username=" + username, true);
    httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    httpRequest.send("");
    httpRequest.onreadystatechange = () => {
        if (httpRequest.readyState == 4 && httpRequest.status == 200) {
            let data = JSON.parse(httpRequest.responseText);
            {
                data
            }
        }
    }

}