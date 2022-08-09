# ユーザーCRUD・ログインAPI

## ユーザー作成 API

### リクエスト

`
POST /api/register
`

| param     | type   | description     |
|-----------|--------|-----------------|
| userID    | string | UserID          |
| name      | string | 表示名             |
| birthdate | string | 誕生日(YYYY-MM-DD) |
| sex       | int    | 性別(男、女で1or2)    |
| height    | int    | 身長              |
| weight    | int    | 体重              |
| password  | string | パスワード           |

userIDは3~20文字以内
passwordは6文字以上

```json
{
    "UsersID": "Pi",
    "Name": "ASDF",
    "Birthdate": "2002-01-01",
    "Sex": 1,
    "Height": 169,
    "Weight": 55,
    "Password": "Raspberry"
}
```

### レスポンス

#### 成功時

| param            | type   | description     |
|------------------|--------|-----------------|
| Detail.userID    | string | userID          |
| Detail.name      | string | 名前              |
| Detail.birthdate | string | 誕生日(YYYY-MM-DD) |
| Detail.sex       | int    | 性別(男、女で1/2)     |
| Detail.height    | int    | 身長              |
| Detail.weight    | int    | 体重              |
| Detail.objective | string | 目標消費カロリー        |

```json
{
    "detail": {
        "usersID": "Pi",
        "name": "ASDF",
        "birthdate": "2002-01-01",
        "sex": 1,
        "height": 169,
        "weight": 55,
        "objective": 100
    }
}
```

#### 失敗時

##### 400 BadRequest

* リクエストのBodyの不備
* IDが重複している場合

## ログインAPI

### リクエスト

`
POST /api/login
`

| param    | type   | description |
|----------|--------|-------------|
| userID   | string | userID      |
| password | string | パスワード       |

```json
{
    "ID": "Pi",
    "Password": "Raspberry"
}
```

### レスポンス

#### 成功時

| param         | type    | description |
|---------------|---------|-------------|
| Detail.Result | Boolean | trueなら承認    |

```json
{
    "Detail": {
        "Result": true
    }
}
```

### 失敗時

#### 400 BadRequest

* リクエストBodyの不備


## ユーザーデータ取得

### リクエスト

`
GET /api/user/getUser?userid={userID}
`

### レスポンス

#### 成功時

| param            | type   | description |
|------------------|--------|-------------|
| Detail.ID        | string | ID          |
| Detail.Name      | string | 名前          |
| Detail.Birthdate | int    | 誕生日(UNIX)   |
| Detail.Sex       | int    | 性別(男、女で1/2) |
| Detail.Height    | int    | 身長          |
| Detail.Weight    | int    | 体重          |
| Detail.Objective | int    | 目標消費カロリー    |

```json
{
    "Detail": {
        "ID": "Pi",
        "Name": "ASDF",
        "Birthdate": "2002-01-01",
        "Sex": 1,
        "Height": 169,
        "Weight": 55,
        "Objective": 100
    }
}
```

### 失敗時

#### 400 BadRequest

* IDが存在しない時
