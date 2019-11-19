# 介绍
go原生error库是一个很棒的设计，但是在一些场景中，存在一些缺陷。在复杂的业务场景中，需要快速定位代码异常位置，但原生的error中只支持panic时查看堆栈数据，在多次return时，或panic/recover后，很难获得异常实际出发在哪。这让代码调试带来了很多困难！

# 获取代码
方式1
```shell script
  go get github.com/lingdor/stackerror
```
方式2 \
go.mod
```go
  require github.com/lingdor/stackerror v0.1.5
```
```shell script
  go mod download
```

#使用方法

创建一个stackError
```go
err:=stackerror.New("your message")
return err
```
抛出一个stackerror
```go
stackerror.Panic("your message")
```
优雅处理error
```go
func aa() error {
    return stackerror.New("err")
}
func main(){
    err:=aa()
    stackerror.CheckPanic(err)
}

```







#Thanks



