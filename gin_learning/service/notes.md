## 常用的database.DBconnect方法解析

### (1)`.Find()`操作

* `.Find(&users)`
  > 会将所有的users查询出来，并且赋值给users

* `.Where("name = ?", "John").Find(&users)`

  > 会将name为John的users查询出来，并且赋值给users

* `.Where("name = ?", "John").Find(&users, "age > ?", 18)`

  > 会将name为John并且age大于18的users查询出来，并且赋值给users

### (2)`Create()`操作

* `.Create(&user)`
  > 会将user插入到数据库中

* `.Create(&users)`
  > 会将users插入到数据库中

### (3)`Delete()`操作

如果被删除的对象不存在，会使得返回的`user`对象的`ID`为0，`DeletedAt`为`time.Time`类型的0值

* `.Delete(&user)`
  > 会将user从数据库中删除

### (4)`Updates()`操作

如果被更新的对象不存在，会使得返回的`user`对象的`ID`为0，`UpdatedAt`为`time.Time`类型的0值

* `.Updates(&user)`

  > 会将user更新到数据库中

* `.Updates(map[string]interface{}{"name": "John", "age": 18})`

  > 会将name为John，age为18的user更新到数据库中