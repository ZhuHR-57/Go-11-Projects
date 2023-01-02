> 1. 删除指定的slice元素

```go
slice = append(slice[:index],slice[index+1:]...)
```

> 2. json方法的选择
>
> + json.NewDecoder用于http连接与socket连接的读取与写入，或者文件读取；（性能高）
> + json.Unmarshal用于直接是byte的输入。

##### `1：json.Unmarshal进行解码`

```go
func HandleUse(w http.ResponseWriter, r *http.Request) {
    var u Use //此处的Use是一个结构体
    data, err := ioutil.ReadAll(r.Body)//此处的r是http请求得到的json格式数据-->然后转化为[]byte格式数据.
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    if err := json.Unmarshal(data, &u); err != nil { //经过这一步将json解码赋值给结构体，由json转化为结构体数据
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "姓名：%s，年龄：%d", u.Name, u.Age)

}
123456789101112131415
```

##### `2. json.NewDecoder解码`

```go
func HandleUse(w http.ResponseWriter, r *http.Request) {
    var u Use
    if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "姓名：%s，年龄：%d", u.Name, u.Age)

}
123456789
```

> 3. 字符串 < - > Int

```go
// 字符串 -> int
ID,err := strconv.ParseInt(bookId,0,0)

// int -> 字符串
book.ID = strconv.Itoa(rand.Intn(10000000))
```

