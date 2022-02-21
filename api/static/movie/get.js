// 获得基本信息
function get(movieNum) {


    let httpRequest = new XMLHttpRequest();
    httpRequest.open("GET", 'http://49.235.99.195:8090/movieInfo/' + movieNum, true);
    httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    httpRequest.send("");
    httpRequest.onreadystatechange = () => {
        if (httpRequest.readyState == 4 && httpRequest.status == 200) {
            var use = JSON.parse(httpRequest.responseText);

            document.getElementById('movieName').innerHTML = use.data.Name + "(" + use.data.Year + ")";

            document.getElementById('img').innerHTML = "<img src='" + use.data.Img + "'  width='135px' height='200px'>"

            document.getElementById('dir').innerHTML = "导演:" + use.data.Director;
            document.getElementById('star').innerHTML = "主演：" + use.data.Starring;
            document.getElementById('type').innerHTML = "类型：" + use.data.Types;
            // document.getElementById('dir').innerHTML = ""use.data.Director;
            document.getElementById('area').innerHTML = "地区：" + use.data.Area;
            document.getElementById('year').innerHTML = "上映时间：" + use.data.Year;
            document.getElementById('time').innerHTML = "时长：" + use.data.Time + "分钟";
            document.getElementById('other').innerHTML = "别名:" + use.data.OtherName;

            document.getElementById('df').innerHTML = use.data.Score

            document.getElementById('Introh').innerHTML = use.data.Name + "的剧情简介......";
            document.getElementById('Intro').innerHTML = use.data.Introduce;
        }
    }

}



// let date = {
//     "name": "movieInfo.Name",          //电影中文名字
//     "otherName": 'movieInfo.OtherName',     //别名
//     "score": 'movieInfo.Score',         //评分
//     "year": 'movieInfo.Year',          //出版年份
//     "time": 'movieInfo.Time + "分钟"',  //时长
//     "area": 'movieInfo.Area',          //地区
//     "director": 'movieInfo.Director',      //导演
//     "starring": 'movieInfo.Starring',      //主演
//     "CommentNum": 'movieInfo.CommentNum',    //影评数
//     "Introduce": 'movieInfo.Introduce',     //简介
//     "WantSee": 'movieInfo.WantSee',       //想看人数
//     "Seen": 'movieInfo.Seen',          //看过人数
//     "types": 'movieInfo.Types',         //类型
//     // "img": 'movieInfo.Img',           //图片url
//     "img": './111.jpg',
//     "url": "http://101.201.234.29:8080/movie/" + 'movieNum',
// }

function short(movieNum) {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", 'http://49.235.99.195:8090/movieInfo/' + movieNum + '/shortComment', true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.send("");
    xhr.onreadystatechange = () => {
        if (xhr.readyState == 4 && xhr.status == 200) {
            let short = JSON.parse(xhr.responseText);

            // 短评
            document.getElementById('shorth').innerHTML = "短评.........(全部" + short.data.length + "条短评)" + '<button class="btn">我要写短评</button>';
            for (let i = 0; i < 3; i++) {
                document.getElementById('short1t').innerHTML = short.data[i].Username + "  " + short.data[i].Topic + "  " + short.data[i].LikeNum + "  " + short.data[0].Time;
                document.getElementById('short1b').innerHTML = short.data[i].Txt;
            }

        }
    }
}

// let short =
// {
//     "data": [
//         {
//             "Topic": 'topic',
//             "MovieNum": 1,
//             "Username": 'username',
//             "Txt": 'txt1',
//             "Time": "2022-01-20 23:02:41",
//             "LikeNum": 0,
//             "Url": "http://101.201.234.29:8080/movieInfo/1"
//         },
//         {
//             "Topic": 'topic',
//             "MovieNum": 1,
//             "Username": 'username',
//             "Txt": 'txt2',
//             "Time": "2022-01-20 23:02:41",
//             "LikeNum": 0,
//             "Url": "http://101.201.234.29:8080/movieInfo/1"
//         }
//     ]
// }



function long(movieNum) {
    let http = new XMLHttpRequest();
    http.open("GET", 'http://49.235.99.195:8090/movieInfo/' + movieNum + '/longComment', true);
    http.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    http.send("");
    http.onreadystatechange = () => {
        if (http.readyState == 4 && http.status == 200) {
            let long = JSON.parse(http.responseText);

            //影评
            document.getElementById('longh').innerHTML = "影评.........(全部" + long.data.length + "条影评)" + '<button class="btn">我要写影评</button>';
            for (let i = 0; i < 3; i++) {
                document.getElementById('long1t').innerHTML = long.data[i].Username + "  " + long.data[i].Topic + "  " + long.data[i].LikeNum + "  " + long.data[i].Time;
                document.getElementById('long1b').innerHTML = long.data[i].Txt;
            }

            // document.getElementById('Areah').innerHTML = date.name + "的讨论区.........(全部" + long.data.length + "条影评)" + '<button class="btn">我要写影评</button>';
        }
    }
}

// }
// let long = {

//     "data": [
//         {
//             "Topic": 'topic',
//             "MovieNum": 1,
//             "Username": 'username',
//             "Txt": 'txt1',
//             "Time": "2022-01-20 23:02:41",
//             "LikeNum": 0,
//             "Url": "http://101.201.234.29:8080/movieInfo/1"
//         },
//         {
//             "Topic": 'topic',
//             "MovieNum": 1,
//             "Username": 'username',
//             "Txt": 'txt2',
//             "Time": "2022-01-20 23:02:41",
//             "LikeNum": 0,
//             "Url": "http://101.201.234.29:8080/movieInfo/1"
//         }
//     ]
// }


get(1);
short(1);
long(1);
