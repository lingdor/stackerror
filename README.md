# Overview
[![Build Status](https://travis-ci.org/lingdor/stackerror.svg?branch=master)](https://travis-ci.org/lingdor/stackerror)
[![codecov](https://codecov.io/gh/lingdor/stackerror/branch/master/graph/badge.svg)](https://codecov.io/gh/lingdor/stackerror)
[![Go Report Card](https://goreportcard.com/badge/github.com/lingdor/stackError)](https://goreportcard.com/report/github.com/lingdor/stackError)

Go native error library is a great design, but in some scenarios, there are some flaws. In a complex business scenario, it is necessary to quickly locate the code exception location, but the native error only supports viewing the stack data when panic is used. When multiple returns or panic / recover are used, it is difficult to get the actual starting point of the exception. This makes code debugging a lot of difficulties!

# Get code

## method 1

```shell script
  go get github.com/lingdor/stackerror
```

##method 2

go.mod
```go
  require github.com/lingdor/stackerror v0.1.5
```
```shell script
  go mod download
```

#Usage

Get a stackError
```go
err:=stackerror.New("your message")
return err
```
Throw a stackerror
```go
stackerror.Panic("your message")
```
Grace method checking error
```go
func aa() error {
    return stackerror.New("err")
}

func main(){
	defer func(){
		if err:=recover();err!=nil {
			fmt.Println(err)
		}
	}()

	err:=aa()
	stackerror.CheckPanic(err)
}

```

Output: \
*stackerror.stackError : err\
  at main.aa( /Users/user/go/testApp/src/main/aa.go:8 )\
  at main.main( /Users/user/go/testApp/src/main/aa.go:18 )\
  at runtime.main( /usr/local/Cellar/go/1.13.4/libexec/src/runtime/proc.go:203 )\

# Thanks



