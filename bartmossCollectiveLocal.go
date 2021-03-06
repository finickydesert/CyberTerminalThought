package main

import (  
    "fmt"
    "io/ioutil"
    "errors"
    "runtime"
    "os/exec"
    
)

//import runtime library of go by running -- go get github.com/go-openapi/runtime


func main() {  
    data, err := ioutil.ReadFile("Bartmoss_Collective.txt")
    if err != nil {
		var strs string
		err := errors.New(err.Error())
		strs="http://stackoverflow.com/search?q=[go]"+err.Error()
        openbrowser(strs)
    fmt.Println("Contents of file:", string(data))
}
func openbrowser(url string) {
	var err error
	err = errors.New("an error message")
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Printf(err.Error())
	}
}