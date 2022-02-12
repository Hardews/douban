API

# Login

## `POST` `/api/login`

###`application/form-data`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| logName | 必选 | 用户名 |
| password | 必选 | 密码 |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |
| token        | 用户token    |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`输入的账号为空`|`logName`为空|
|`false`|`输入的密码为空`|`password`为空|
|`false`|`无此账号`|`logName`不存在|
|`false`|`密码错误`|`logName`与`password`不匹配|
|`ture`|`成功`|成功|

# Register

## `POST` `api/register`
###`application/form-data`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| signName | 必选 | 用户名 |
| signPassword | 必选 | 密码 |
| nickName | 非必选 | 昵称 |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`输入的账号为空`|`signName`为空|
|`false`|`输入的密码为空`|`signPassword`为空|
|`false`|`用户名已存在`|`signName`已存在|
|`false`|`用户名含有敏感词汇`|`nickName`含有敏感词汇|
|`false`|`昵称含有敏感词汇`|`nickName`含有敏感词汇|
|`false`|`密码长度不合法`|`signPassword`太长或太短|
|`ture`|`成功`|注册成功|

# User

## `POST` `api/user/change`
###`application/form-data`

####`修改密码`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| oldPassword | 必选 | 旧密码 |
| newPassword | 必选 | 新密码 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`输入的旧密码为空`|`oldPassword`为空|
|`false`|`输入的新密码为空`|`newPassword`为空|
|`false`|`旧密码不正确`|`oldPassword`已存在|
|`false`|`新密码长度不合法`|`newPassword`含有敏感词汇|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`修改成功`|修改成功|


## `POST` `api/user/introduce`
### `application/form-data`

#### `设置自我介绍`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| introduce | 必选 | 自我介绍 |
| token | 必选 | token |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`输入的自我介绍为空`|`introduce`为空|
|`false`|`输入的自我介绍含有敏感词`|`introduce`含有敏感词|
|`false`|`自我介绍长度不合法`|`introduce`不合法|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`设置成功`|设置成功|

## `POST` `api/user/setQuestion`
### `application/form-data`

#### `设置密保问题`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| question | 必选 | 密保问题 |
| answer | 必选 | 密保答案 |
| password | 必选 | 密码 |
| token | 必选 | token |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`密码为空`|`password`为空|
|`false`|`问题为空`|`question`为空|
|`false`|`答案为空`|`answer`为空|
|`false`|`密码错误`|`password`与用户名不匹配|
|`false`|`已设置密保`|`question`已设置|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`设置成功`|`设置成功`|

## `POST` `api/user/retrieve`
`application/form-data`

`通过密保找回密码`
`只有当question，answer，username都正确时，可发送新密码`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| answer | 必选 | 密保答案 |
| username | 必选 | 用户名 |
|newPassword|必选|新密码|

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`密码为空`|`newPassword`为空|
|`false`|`用户名为空`|`username`为空|
|`false`|`无此账号`|没有这个用户|
|`false`|`该账号无密保，可通过申诉找回`|该用户没设置`question`|
|`false`|`答案不正确`|`question`与`answer`不匹配|
|`false`|`答案为空`|`answer`为空|
|`ture`|`修改成功`|找回成功|

## `GET` `/api/:username/menu`
### `获取username为:username的用户基础信息`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :username | 必选 | 用户id |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|false|"无此用户"|`username`不存在|
|true|参见以下代码|成功|

```
{
        "username":  username,
        "nickName":  user.NickName || nil,
        "introduce": user.Introduce || nil,
}
```

## `GET` `api/:username/Comment`
### `获取username为:username的用户影评和短评`


| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :username | 必选 | 用户id |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|false|"无此用户"|`username`不存在|
|true|"该用户暂时无评论”|获取成功|
|true|参见以下代码|获取成功|

```
{
	"url": "http://101.201.234.29:8080/movie/movieNum",
	"username": username,
	"txt":      shortComments || longComments,
	"time":     Time,
}
```

## `GET` `api/:username/wantSee`
###`获取username为:username的用户的想看文件夹`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :username | 必选 | 用户id |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|false|"无此用户"|`username`不存在|
|true|"该用户暂时无评论”|获取成功|
|true|参见以下代码|获取成功|

```
{
	"username": username,
	"comment":  Comment || nil,
	"movieNum": MovieNum,
	"label":    Label || nil,
	"url":      "http://101.201.234.29:8080/movie/movieNum",
}
{
        "num": "n部想看",
}
```

## `GET` `/api/user/:username/Seen`
###`获取username为:username的用户的看过文件夹`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :username | 必选 | 用户id |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|false|"无此用户"|`username`不存在|
|true|"该用户暂时无评论”|获取成功|
|true|参见以下代码|获取成功|

```
{
	"username": username,
	"comment":  Comment || nil,
	"movieNum": MovieNum,
	"label":    Label || nil,
	"url":      "http://101.201.234.29:8080/movie/movieNum", //电影详情页url
}
{
        "num": "n部看过",
}
```
# `Home`
## `GET` `/api/home/research/:find`
### `搜索功能`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :find | 必选 | 搜索内容 |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|false|"抱歉，暂时没有你想要的电影”|搜索不到|
|true|参见以下代码|获取成功|

```
data:{
	"name":       movieInfo.Name,          //电影中文名字
	"otherName":  movieInfo.OtherName,     //别名
	"score":      movieInfo.Score,         //评分
	"year":       movieInfo.Year,          //出版年份
	"time":       movieInfo.Time + "分钟",  //时长
	"area":       movieInfo.Area,          //地区
	"director":   movieInfo.Director,      //导演
	"starring":   movieInfo.Starring,      //主演
	"CommentNum": movieInfo.CommentNum,    //影评数
	"Introduce":  movieInfo.Introduce,     //简介
	"WantSee":    movieInfo.WantSee,       //想看人数
	"Seen":       movieInfo.Seen,          //看过人数
	"types":      movieInfo.Types,         //类型
	"img":        movieInfo.Img,           //图片url
	"url":        "http://101.201.234.29:8080/movie/" + movieNum,
}
```
## `GET` `/api/home/recommend`
### `推荐功能（随机`


| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|true|参见以下代码|获取成功|

```
data:{
	"name":       movieInfo.Name,          //电影中文名字
	"otherName":  movieInfo.OtherName,     //别名
	"score":      movieInfo.Score,         //评分
	"year":       movieInfo.Year,          //出版年份
	"time":       movieInfo.Time + "分钟",  //时长
	"area":       movieInfo.Area,          //地区
	"director":   movieInfo.Director,      //导演
	"starring":   movieInfo.Starring,      //主演
	"CommentNum": movieInfo.CommentNum,    //影评数
	"Introduce":  movieInfo.Introduce,     //简介
	"WantSee":    movieInfo.WantSee,       //想看人数
	"Seen":       movieInfo.Seen,          //看过人数
	"types":      movieInfo.Types,         //类型
	"img":        movieInfo.Img,           //图片url
	"url":        "http://101.201.234.29:8080/movie/" + movieNum,
}
```

## `GET` `/api/home/:category`
###`分类功能`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :find | 必选 | 搜索内容 |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|false|"抱歉，暂时没有你想要的电影”|搜索不到|
|true|参见以下代码|获取成功|

```
data:{
	"name":       movieInfo.Name,          //电影中文名字
	"otherName":  movieInfo.OtherName,     //别名
	"score":      movieInfo.Score,         //评分
	"year":       movieInfo.Year,          //出版年份
	"time":       movieInfo.Time + "分钟",  //时长
	"area":       movieInfo.Area,          //地区
	"director":   movieInfo.Director,      //导演
	"starring":   movieInfo.Starring,      //主演
	"CommentNum": movieInfo.CommentNum,    //影评数
	"Introduce":  movieInfo.Introduce,     //简介
	"WantSee":    movieInfo.WantSee,       //想看人数
	"Seen":       movieInfo.Seen,          //看过人数
	"types":      movieInfo.Types,         //类型
	"img":        movieInfo.Img,           //图片url
	"url":        "http://101.201.234.29:8080/movie/" + movieNum,
}
```

#Movie
## `GET` `api/movieInfo/:movieNum`
###`通过:movieNum获取电影信息`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|false|"无该电影的信息”|获取不到|
|true|参见以下代码|获取成功|
```
data:{
        "name":       movieInfo.Name,          //电影中文名字
        "otherName":  movieInfo.OtherName,     //别名
        "score":      movieInfo.Score,         //评分
        "year":       movieInfo.Year,          //出版年份
        "time":       movieInfo.Time + "分钟",  //时长
        "area":       movieInfo.Area,          //地区
        "director":   movieInfo.Director,      //导演
        "starring":   movieInfo.Starring,      //主演
        "CommentNum": movieInfo.CommentNum,    //影评数
        "Introduce":  movieInfo.Introduce,     //简介
        "WantSee":    movieInfo.WantSee,       //想看人数
        "Seen":       movieInfo.Seen,          //看过人数
        "types":      movieInfo.Types,         //类型
        "img":        movieInfo.Img,           //图片url
        "url":        "http://101.201.234.29:8080/movie/" + movieNum,
}
```
## `GET` `api/movieInfo/:movieNum/longComment`
### `通过:movieNum获取电影的影评`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|false|"无影评”|该电影下无影评|
|true|参见以下代码|获取成功|
```
{
    "data": [
        {
            "Topic": topic,
            "MovieNum": 1,
            "Username": username,
            "Txt": txt1,
            "Time": "2022-01-20 23:02:41",
            "LikeNum": 0,
            "Url": "http://101.201.234.29:8080/movieInfo/1"
        },
        {
            "Topic": topic,
            "MovieNum": 1,
            "Username": username,
            "Txt": txt2,
            "Time": "2022-01-20 23:02:41",
            "LikeNum": 0,
            "Url": "http://101.201.234.29:8080/movieInfo/1"
        }
    ]
}
```
## `GET` `api/movieInfo/:movieNum/shortComment`
###`通过:movieNum获取电影的短评`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|false|"无短评”|该电影下无短评|
|true|参见以下代码|获取成功|
```
{
    "data": [
        {
            "Topic": topic,
            "MovieNum": 1,
            "Username": username,
            "Txt": txt1,
            "Time": "2022-01-20 23:02:41",
            "LikeNum": 0,
            "Url": "http://101.201.234.29:8080/movieInfo/1"
        },
        {
            "Topic": topic,
            "MovieNum": 1,
            "Username": username,
            "Txt": txt2,
            "Time": "2022-01-20 23:02:41",
            "LikeNum": 0,
            "Url": "http://101.201.234.29:8080/movieInfo/1"
        }
    ]
}
```
## `GET` `api/movieInfo/:movieNum/commentArea`
### `通过:movieNum获取电影的讨论区`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|false|"无讨论区”|该电影下无讨论区|
|true|"无评论"|讨论区话题下无评论|
|true|参见以下代码|获取成功|
```
{
    "commentNum": 3,
    "likeNum": 1,
    "time": "2022-02-10 14:07:26",
    "topic": "1",
    "username": "1225101128"
}{
    "comment": "test",
    "likeNum": 0,
    "time": "2022-02-08 00:57:41",
    "username": "1225101128"
}{
    "comment": "tes3",
    "likeNum": 2,
    "time": "2022-02-08 01:02:42",
    "username": "123456"
}{
    "commentNum": 0,
    "likeNum": 0,
    "time": "2022-02-12 02:58:30",
    "topic": "2",
    "username": "1225101127"
}{
    "data": "无评论"
}
```
## `POST` `api/movie/:movieNum/wantSee`
### `application/form-data`

####`用户设置想看`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
| comment | 必选 | 收藏的简短评论 |
| label | 必选 | 存储标签 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`成功`|设置成功|

## `POST` `api/movie/:movieNum/seen`
### `application/form-data`

#### `用户设置看过`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
| comment | 必选 | 收藏的简短评论 |
| label | 必选 | 存储标签 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`"成功"`|设置成功|

## `DELETE` `api/movie/:movieNum/wantSee`
### `application/form-data`

#### `用户删除想看内容`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
| label | 非必选 | 存储标签 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`删除成功`|成功|

## `DELETE` `api/movie/:movieNum/seen`
### `application/form-data`

#### `用户删除看过内容`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
| label | 非必选 | 存储标签 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`删除成功`|成功|

## `POST` `api/movie/:movieNum/longComment`
###`application/form-data`

####`用户给予影评`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
| topic | 必选 | 影评标题 |
|LongComment|必选|影评内容|
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`影评含有敏感词汇`|`LongComment`含有敏感词汇|
|`false`|`长度不合法`|`LongComment`过长或过短|
|`false`|`已有影评`|`LongComment`已存在|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`成功`|成功|

## `DELETE` `api/movie/:movieNum/longComment`
### `application/form-data`

#### `用户删除影评`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`影评不存在`|`LongComment`不存在|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`删除成功`|成功|

## `PUT` `api/movie/:movieNum/longComment`
### `application/form-data`

#### `用户更新短评`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
|comment|必选|新影评|
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`影评不存在`|`LongComment`不存在|
|`false`|`影评含有敏感词汇`|`LongComment`含有敏感词汇|
|`false`|`长度不合法`|`LongComment`过长或过短|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`更新成功`|成功|

## `POST` `api/movie/:movieNum/shortComment`
###`application/form-data`

#### `用户给予短评`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
| topic | 必选 | 影评标题 |
|ShortComment|必选|影评内容|
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`短评含有敏感词汇`|`ShortComment`含有敏感词汇|
|`false`|`长度不合法`|`ShortComment`过长或过短|
|`false`|`已有短评`|`ShortComment`已存在|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`"成功"`|成功|
## `DELETE` `api/movie/:movieNum/shortComment`
###`application/form-data`

####`用户删除短评`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`短评不存在`|`ShortComment`不存在|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`"删除成功"`|成功|

## `PUT` `api/movie/:movieNum/shortComment`
### `application/form-data`

#### `用户更新短评`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
|comment|必选|新短评|
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`短评不存在`|`ShortComment`不存在|
|`false`|`短评含有敏感词汇`|`ShortComment`含有敏感词汇|
|`false`|`长度不合法`|`ShortComment`过长或过短|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`更新成功`|成功|

## `POST` `api/movie/:movieNum/commentArea`
### `application/form-data`

#### `用户发表讨论区话题`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
| topic | 必选 | 话题 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`"已有话题"`|`topic`已存在|
|`false`|`"话题为空"`|`topic`为空|
|`false`|`"话题含有敏感词汇"`|`topic`含有敏感词汇|
|`false`|`"话题长度不合法"`|`topic`长度不合法|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`成功`|设置成功|

## `DELETE` `api/movie/:movieNum/:areaNum`
### `application/form-data`

#### `用户删除讨论区话题`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
|  :areaNum| 必选 | 话题编号 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`"话题不存在"`|`topic`不存在|
|`false`|`"话题含有敏感词汇"`|`topic`含有敏感词汇|
|`false`|`"话题长度不合法"`|`topic`不合法|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`删除成功`|设置成功|

## `PUT` `api/movie/:movieNum/commentArea`
### `application/form-data`

#### `用户更新讨论区话题`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
| topic | 必选 | 话题 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`"话题为空"`|`topic`为空|
|`false`|`"话题含有敏感词汇"`|`topic`有敏感词汇|
|`false`|`"话题长度不合法"`|`topic`长度不合法|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`更新成功`|设置成功|

## `POST` `api/movie/:movieNum/:areaNum/like`
### `application/form-data`

#### `用户点赞讨论区话题`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
|  :areaNum| 必选 | 话题编号 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`您已经点过赞啦`|用户点过赞了|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`点赞成功!`|成功|

## `POST` `api/movie/:movieNum/:areaNum/like`
###`application/form-data`

#### `用户取消点赞讨论区话题`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
|  :areaNum| 必选 | 话题编号 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`""`|成功|

##  `POST` `api/movie/:movieNum/:areaNum/comment`

###`application/form-data`

####`用户发表评论（讨论区话题`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
|comment | 必选 | 评论 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`"评论为空"`|`comment`为空|
|`false`|`"评论含有敏感词汇"`|`comment`有敏感词汇|
|`false`|`"评论长度不合法"`|`comment`长度不合法|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`成功`|设置成功|

## `DELETE` `api/movie/:movieNum/:areaNum/comment`
### `application/form-data`

#### `用户删除评论（讨论区话题`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
|  :areaNum| 必选 | 话题编号 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`"话题不存在"`|`comment`不存在|
|`false`|`"话题为空"`|`comment`为空|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`删除成功`|设置成功|

## `PUT` `api/movie/:movieNum/:areaNum/comment`
### `application/form-data`

#### `用户更新评论（讨论区话题`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| :movieNum | 必选 | 电影编号 |
| comment | 必选 | 评论 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`"评论为空"`|`comment`为空|
|`false`|`"评论含有敏感词汇"`|`comment`有敏感词汇|
|`false`|`"评论长度不合法"`|`comment`长度不合法|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`更新成功`|设置成功|

## `POST` `api/movie/:movieNum/:areaNum/comment/like`
### `application/form-data`

#### `用户点赞评论（讨论区话题`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
|username|必选|被点赞人的用户名|
| :movieNum | 必选 | 电影编号 |
|  :areaNum| 必选 | 话题编号 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`您已经点过赞啦`|用户点过赞了|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`点赞成功!`|成功|

## `POST` `api/movie/:movieNum/:areaNum/comment/like`
### `application/form-data`

#### `用户取消点赞评论（讨论区话题`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
|username|必选|被点赞人的用户名|
| :movieNum | 必选 | 电影编号 |
|  :areaNum| 必选 | 话题编号 |
| token | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`""`|成功|

#Administrator
## `POST` `api/administrator/setNewMovie`
### `application/form-data`

#### `管理员添加电影信息`

| 请求参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| movieName | 必选 | 电影编号 |
| otherName | 必选 | 别名 |
| score | 必选 | 评分 |
| Starring | 必选 | 主演 |
| Area | 必选 | 地区 |
| Time | 必选 | 时长 |
| Director | 必选 | 导演 |
| Types | 必选 | 类型 |
| Introduce | 必选 | 简介 |
| Year | 必选 | 出版年份 |
| token | 必选 | token |


| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status    | data     | 说明   |
| ------------ | ------------ | ------------ |
|`false`|`"非管理员无权限操作"`|权限不足|
|`false`|`token为空`|`token`为空|
|`false`|`token不正确`|`token`不正确|
|`false`|`token格式不正确`|`token`格式不正确|
|`false`|`token已过期`|`token`已过期|
|`ture`|`movieNum`|设置成功,返回电影编号|