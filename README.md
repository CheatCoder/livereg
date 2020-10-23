# livereg

[![PkgGoDev](https://pkg.go.dev/badge/github.com/cheatcoder/livereg)](https://pkg.go.dev/github.com/cheatcoder/livereg)

Test the Regexp with live data and work with the found items (ESC key to leave the regexp input) after found the right regexp

```
 package main
 
 import(
    "fmt"
    "log"
    
    "github.com/CheatCoder/livereg"
 )
 
 func main(){
    text,err := livereg.NewStringReg("Write Here your Text or give it with a String Variable - foo bar foobar")
    if err != nil{
      log.Panic(err)
    }
    for _,val := range text {
      fmt.Println(string(text))
    }
 }
```
