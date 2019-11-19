# Overview
[![Build Status](https://travis-ci.org/lingdor/stackError.svg?branch=master)](https://travis-ci.org/lingdor/stackError)
[![codecov](https://codecov.io/gh/lingdor/stackError/branch/master/graph/badge.svg)](https://codecov.io/gh/lingdor/stackError)


Go native error library is a great design, but in some scenarios, there are some flaws. In a complex business scenario, it is necessary to quickly locate the code exception location, but the native error only supports viewing the stack data when panic is used. When multiple returns or panic / recover are used, it is difficult to get the actual starting point of the exception. This makes code debugging a lot of difficulties!

# Get code

## method 1

```shell script
  go get github.com/lingdor/stackError
```

##method 2

go.mod
```go
  require github.com/lingdor/stackError v0.1.5
```
```shell script
  go mod download
```

#Usage

Get a stackError
```go
err:=stackError.New("your message")
return err
```
Throw a stackerror
```go
stackError.Panic("your message")
```
Grace method checking error
```go
func aa() error {
    return stackError.New("err")
}

func main(){
	defer func(){
		if err:=recover();err!=nil {
			fmt.Println(err)
		}
	}()

	err:=aa()
	stackError.CheckPanic(err)
}

```

Output: \
*stackError.stackError : err
  at main.aa( /Users/bobby96333/go/testApp/src/main/aa.go:8 )
  at main.main( /Users/bobby96333/go/testApp/src/main/aa.go:18 )
  at runtime.main( /usr/local/Cellar/go/1.13.4/libexec/src/runtime/proc.go:203 )

# Thanks



