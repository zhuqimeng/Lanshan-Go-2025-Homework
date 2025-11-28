# 朱启梦的第五次 go 组作业

### 项目说明

文件检索的小项目在 catch 文件夹里面。

不是自己写的，太多文件操作和函数调用我不会，让 deepseek 生成的。

但是对于这份代码的思路和内容都看懂了，如果自己写估计也没法做什么优化了。

我就摆在这里，以后想复习了就来看看。

### 使用方式：

cd 到 .exe 所在的文件夹，然后终端运行：

```bash
./catch /path/to/your/code "function name"
```

### 项目总结

`filepath` 库里的 `Walk` 函数可以递归遍历目录树中的所有文件和目录。

`filepath.Walk` 接受两个参数：一个是根目录路径，另一个是一个函数，该函数将在遍历每个文件或目录时被调用。

这个函数的签名如下：

`type WalkFunc func(path string, info os.FileInfo, err error) error`

