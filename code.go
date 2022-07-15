package main
import (
    base64 "encoding/base64"
    "fmt"
)
func reverse(str string) (result string) {
    for _, v := range str {
        result = string(v) + result
    }
    return
}
func main(){
    var whatIsIt string
    secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
    sd, _ := base64.StdEncoding.DecodeString(secret)

    whatIsIt = reverse(string(sd))
    fmt.Println(string(whatIsIt))
}
