// function get(movieNum) {
//     let obj1 = {
//         "movieNum": movieNum,
//     };
//     let httpRequest = new XMLHttpRequest();
//     httpRequest.open("GET", '接口', true);
//     httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
//     httpRequest.send(JSON.stringify(obj1));
//     httpRequest.onreadystatechange = () => {
//         if (httpRequest.readyState == 4 && httpRequest.status == 200) {
//             let data = JSON.parse(httpRequest.responseText);

//         }
//     }

// }



let date = {
    "name": "movieInfo.Name",          //电影中文名字
    "otherName": 'movieInfo.OtherName',     //别名
    "score": 'movieInfo.Score',         //评分
    "year": 'movieInfo.Year',          //出版年份
    "time": 'movieInfo.Time + "分钟"',  //时长
    "area": 'movieInfo.Area',          //地区
    "director": 'movieInfo.Director',      //导演
    "starring": 'movieInfo.Starring',      //主演
    "CommentNum": 'movieInfo.CommentNum',    //影评数
    "Introduce": 'movieInfo.Introduce',     //简介
    "WantSee": 'movieInfo.WantSee',       //想看人数
    "Seen": 'movieInfo.Seen',          //看过人数
    "types": 'movieInfo.Types',         //类型
    // "img": 'movieInfo.Img',           //图片url
    "img": './111.jpg',
    "url": "http://101.201.234.29:8080/movie/" + 'movieNum',
}

let short =
{
    "data": [
        {
            "Topic": 'topic',
            "MovieNum": 1,
            "Username": 'username',
            "Txt": 'txt1',
            "Time": "2022-01-20 23:02:41",
            "LikeNum": 0,
            "Url": "http://101.201.234.29:8080/movieInfo/1"
        },
        {
            "Topic": 'topic',
            "MovieNum": 1,
            "Username": 'username',
            "Txt": 'txt2',
            "Time": "2022-01-20 23:02:41",
            "LikeNum": 0,
            "Url": "http://101.201.234.29:8080/movieInfo/1"
        }
    ]
}

let long = {

    "data": [
        {
            "Topic": 'topic',
            "MovieNum": 1,
            "Username": 'username',
            "Txt": 'txt1',
            "Time": "2022-01-20 23:02:41",
            "LikeNum": 0,
            "Url": "http://101.201.234.29:8080/movieInfo/1"
        },
        {
            "Topic": 'topic',
            "MovieNum": 1,
            "Username": 'username',
            "Txt": 'txt2',
            "Time": "2022-01-20 23:02:41",
            "LikeNum": 0,
            "Url": "http://101.201.234.29:8080/movieInfo/1"
        }
    ]
}

document.getElementById('movieName').innerHTML = date.name + "(" + date.year + ")";

document.getElementById('img').innerHTML = "<img src='./111.jpg'  width='135px' height='200px'>"

document.getElementById('dir').innerHTML = date.director;
document.getElementById('star').innerHTML = date.starring;
document.getElementById('type').innerHTML = date.types;
document.getElementById('dir').innerHTML = date.director;
document.getElementById('area').innerHTML = date.area;
document.getElementById('year').innerHTML = date.year;
document.getElementById('time').innerHTML = date.time;
document.getElementById('other').innerHTML = date.otherName;

document.getElementById('df').innerHTML = date.score
// 简介
document.getElementById('Introh').innerHTML = date.name + "的剧情简介......";
document.getElementById('Intro').innerHTML = date.Introduce;
// 短评
document.getElementById('shorth').innerHTML = date.name + "的短评.........(全部" + short.data.length + "条短评)" + '<button class="btn">我要写短评</button>';
document.getElementById('short1t').innerHTML = short.data[0].Username + "  " + short.data[0].Topic + "  " + short.data[0].LikeNum + "  " + short.data[0].Time;
document.getElementById('short1b').innerHTML = short.data[0].Txt;

document.getElementById('short2t').innerHTML = short.data[0].Username + "  " + short.data[0].Topic + "  " + short.data[0].LikeNum + "  " + short.data[0].Time;
document.getElementById('short2b').innerHTML = short.data[0].Txt;

document.getElementById('short3t').innerHTML = short.data[0].Username + "  " + short.data[0].Topic + "  " + short.data[0].LikeNum + "  " + short.data[0].Time;
document.getElementById('short3b').innerHTML = short.data[0].Txt;

document.getElementById('short4t').innerHTML = short.data[0].Username + "  " + short.data[0].Topic + "  " + short.data[0].LikeNum + "  " + short.data[0].Time;
document.getElementById('short4b').innerHTML = short.data[0].Txt;

//影评 
document.getElementById('longh').innerHTML = date.name + "的影评.........(全部" + long.data.length + "条影评)" + '<button class="btn">我要写影评</button>';
document.getElementById('long1t').innerHTML = long.data[0].Username + "  " + long.data[0].Topic + "  " + long.data[0].LikeNum + "  " + long.data[0].Time;
document.getElementById('long1b').innerHTML = long.data[0].Txt;

document.getElementById('long2t').innerHTML = long.data[0].Username + "  " + long.data[0].Topic + "  " + long.data[0].LikeNum + "  " + long.data[0].Time;
document.getElementById('long2b').innerHTML = long.data[0].Txt;

document.getElementById('long3t').innerHTML = long.data[0].Username + "  " + long.data[0].Topic + "  " + long.data[0].LikeNum + "  " + long.data[0].Time;
document.getElementById('long3b').innerHTML = long.data[0].Txt;

document.getElementById('long4t').innerHTML = long.data[0].Username + "  " + long.data[0].Topic + "  " + long.data[0].LikeNum + "  " + long.data[0].Time;
document.getElementById('long4b').innerHTML = long.data[0].Txt;

document.getElementById('Areah').innerHTML = date.name + "的讨论区.........(全部" + long.data.length + "条影评)" + '<button class="btn">我要写影评</button>';