#### 1.类型转换的问题
在go的后台开发过程中，用到了GitHub上的mysql驱动，因为表中有字段，updatetime 类型为timestamp
而这个类型对应的golang中的类型是time.Time，当我这样使用的时候发现报错：
```
t time.Time
rows,_ :=conn.Query("select updatetime from t_table")
.....
rows.scan(&t)
```
> 错误：unsupported Scan, storing driver.Value type []uint8 into type *time.Time

>解决的方法是：parseTime=true这个加在mysql open的时候，例如：
```
fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",.....)
```

#### 2.时区的问题
同样的，我发现，上面的问题解决了，但是我这里是北京东八区的时间，而日志中打印的总是0时区，而这影响了Unix()函数的转换结果,因为加载和转换的时候,默认用了UTC
```
t.Unix() //这个是2019-06-17 20:36:48.000 +0000 UTC---->1560803808
```
> 解决方法是同样的在open的时候指定时区loc=Local,
````
%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local
```