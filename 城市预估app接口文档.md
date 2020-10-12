社区部分

[TOC]



##### post：搜索问题   title  (string类型)
 
##### get  获得收藏的问题  

##### get：得到问题的接口

返回形式相同
 
```markdown
{
  "data": [
    {
      "q_id": 777,
      "title": "非艺术专业的同学在红岩视觉能待下去吗？",
      "description": "。。。",
      "place": "重庆",//上传的地点，支持用户进行地区搜索
      "disappear_at": "2019-11-12 01:09:00",
      "nickname": "用户名字",
      "photo_avator":"" //用户头像
      "photo_url": [],//问题图片url
      "created_at": "2019-11-03 18:07:42",
      "update_at": "2020-08-23 23:28:56",
      "answer_num": 0
    },
    {
      "q_id": 774,
      "title": "红岩",
      "description": "大二还能加吗",
      "place": "重庆",//上传的地点，支持用户进行地区搜索
      "disappear_at": "2019-11-03 18:18:00",
      "nickname": "用户名字",
      "photo_avator":"" //用户头像
      "photo_url": [],
      "created_at": "2019-11-03 17:18:52",
      "update_at": "2020-08-23 22:34:06",
      "answer_num": 0
    },
  ],
  "info": "success",
  "status": 200
}
```

##### post：提出问题  根据    根据用户id

```markdown
      "title": "非艺术专业的同学在红岩视觉能待下去吗？",
      "description": "。。。",
      "disappear_at": "2019-11-12 01:09:00",
      "nickname": "用户名字",
      "photo_url": [],//图片url
      "created_at": "2019-11-03 18:07:42",
      "update_at": "2020-08-23 23:28:56",
      
      返回形式：
      {
      "info": "success",
      "status": 200
      }
```

##### post：提出评论   根据  问题  q_id   和根据用户id

```markdown
      "description": "。。。",
      "place": "重庆",//上传的地点，支持用户进行地区搜索
      "disappear_at": "2019-11-12 01:09:00",
      "nickname": "用户名字",
      "photo_url": [],//图片url
      "created_at": "2019-11-03 18:07:42",
      "update_at": "2020-08-23 23:28:56",
      
       返回形式：
      {
      "info": "success",
      "status": 200
      }
```

##### get：根据问题 q_id 得到评论

```markdown
{
  "data": [
    {
      "a_id": 777,
      "description": "。。。",
      "disappear_at": "2019-11-12 01:09:00",
      "nickname": "用户名字",
      "photo_avator":"" //用户头像
      "photo_url": [],//问题图片url
      "created_at": "2019-11-03 18:07:42",
      "update_at": "2020-08-23 23:28:56",
    },
    {
      "a_id": 774,      
      "description": "大二还能加吗",
      "disappear_at": "2019-11-03 18:18:00",
      "nickname": "用户名字",
      "photo_url": [],
      "created_at": "2019-11-03 17:18:52",
      "update_at": "2020-08-23 22:34:06",
    },
  ],
  "info": "success",
  "status": 200
}
```

##### post：问题图片上传   根据  q_id



##### post：回答图片上传   根据  a_id

   



##### post：收藏问题  q_id   根据用户id

##### post：点赞该问题 q_id 和   根据用户id

##### post：取消该问题点赞 q_id  根据用户id

##### post :  点赞评论  a_id   根据用户id

##### post  :  取消点赞评论  a_id  根据用户id

##### post：取消问题收藏  q_id   根据用户id

这六个返回形式：

```markdown
{
   "info": "success",
   "status": 200
}
```

##### 回复评论 相关接口：

###### get  得到 回复评论  根据 a_id

```markdown
//这里不允许上传图片所以就没有了得到图片的url
{
  "data": [
    {
      "c_id": 777,
      "description": "。。。",
      "disappear_at": "2019-11-12 01:09:00",
      "nickname": "用户名字",
      "photo_avator":"" //用户头像
      "created_at": "2019-11-03 18:07:42",
      "update_at": "2020-08-23 23:28:56",
    },
    {
      "c_id": 774,      
      "description": "大二还能加吗",
      "disappear_at": "2019-11-03 18:18:00",
      "nickname": "用户名字",
      "created_at": "2019-11-03 17:18:52",
      "update_at": "2020-08-23 22:34:06",
    },
  ],
  "info": "success",
  "status": 200
}
```



###### post  发布回复评论  根据  a_id

```markdown
     
      "description": "大二还能加吗",
      "disappear_at": "2019-11-03 18:18:00",
      "nickname": "用户名字",
      "created_at": "2019-11-03 17:18:52",
      "update_at": "2020-08-23 22:34:06",
       返回形式：
      {
      "info": "success",
      "status": 200
      }
```

