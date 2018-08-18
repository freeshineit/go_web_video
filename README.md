# web_voide

## api设计

### user 用户

|   name             | url       |  medthod  |  status code  |
|--------------------|-----------|:-----------:|---------------|
|创建用户（注册）       | /user     |  post     | 201、400、500 |
|用户登录             | /user/:username | post | 200、400、500|
|获取用户基本信息       | /user/:username | get | 200、400、401、403、500|
|用户注销             | /user/:username | delete | 204、400、401、403、500|

### video 视频

|   name             | url                             |  medthod    |  status code           |
|--------------------|---------------------------------|:-----------:|------------------------|
|视频列表             | /user/:username/videos          |  get        | 200、400、500          |
|播放视频             | /user/:username/videos/:vid-id  |  get        | 200、400、500          |
|删除视频             | /user/:username/videos/vid-id   |  delete     | 204、400、401、403、500 |

### comment 评论

|   name       | url                                    |  medthod    |  status code           |
|--------------|----------------------------------------|:-----------:|------------------------|
|评论          | /videos/:vid-id/comments                |  get        | 200、400、500          |
|提交评论       | /videos/:vid-id/comments               |  post       | 201、400、500          |
|删除评论       | /videos/:vid-id/comments/:comment-id   |  delete     | 204、400、401、403、500 |





