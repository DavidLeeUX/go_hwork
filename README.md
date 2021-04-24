# go_hwork

1.我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
- 需要wrap处理错误， 将各层错误统一交个调用方处理
- [代码](week02/main.go)