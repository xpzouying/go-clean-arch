# Demo - Twitter


## 用户

### 获取用户

**请求**

```bash
curl --location --request GET 'http://127.0.0.1:8080/get-user' \
--header 'Content-Type: application/json' \
--header 'Content-Length: ' \
--data-raw '{"uid": 1}'
```


**响应**

```json
{
    "uid": 1,
    "name": "杨磊",
    "avatar": "http://dummyimage.com/100x100"
}

```


### 创建用户

**请求**

```bash
curl --location --request POST 'http://127.0.0.1:8080/create-user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "杨磊",
    "avatar": "http://dummyimage.com/100x100"
}'
```



**响应**

```json
{
    "uid": 1
}
```

