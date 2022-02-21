let btn = document.getElementById("ssbtn")
let find = document.getElementById("find")


btn.onclick = function () {
    let httpRequest = new XMLHttpRequest();
    httpRequest.open("GET", 'http://49.235.99.195:8090/home/research/' + find.value, true);
    httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    httpRequest.send("");
    httpRequest.onreadystatechange = () => {
        if (httpRequest.readyState == 4 && httpRequest.status == 200) {
            var res = httpRequest.responseText;
            let dx = res.replace('"{', '{"')
            dx = dx.replace('}"', '}')
            // let dx = res.replace(/}}/g, "}},");
            // let item = JSON.parse("[" + dx.substring(0, dx.length - 1) + "]");
            var item = JSON.parse(dx)
            // data = data.concat(item);
            // window.onload = function () {
            //     window.location.href = "http://www.baidu.com";
            // }
            if (item.data == undefined) {
                alert("没有你想看的东西哦")
            }
            else {
                window.location.replace(item.data.Url);
            }


        }
    }
}