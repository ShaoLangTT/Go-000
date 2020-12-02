学习笔记

```batch
问题: 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
```
思路：
```batch
在 dao 层中 err 不会空时，我们要进行判断：
  1.当返回 sql.ErrNoRows 时（这种情况是没有数据，我认为并不是报错的），我们可以将其转换为内部自定义的错误，这样的好处是: 
    a.在service层可以兼容不同数据库查询为空情况 
    b.我们也可以快速定位
    c.针对这种情况也可以进行降级处理
  
  2. 当返回不为 sql.ErrNoRows 时，我们应该 Wrap 这个 error，封装其堆栈信息，便于后期定位
    
```

```batch
在service层时，我们要先进行错误判断，进行相关降级处理之后，就可以将错误直接抛到 controller层
```

```batch
controller层进行日志记录
```