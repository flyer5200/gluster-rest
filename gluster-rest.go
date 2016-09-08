package main

import (
	"github.com/flyer5200/gorest"
	"net/http"
	"os/exec"
        "fmt"
)

type GlusterService struct {
        gorest.RestService `root:"/gluster/"`
	gluster            gorest.EndPoint `method:"GET" path:"/{...:string}" output:"string"`
}

func (serv GlusterService) Gluster(vars ...string) string {
	if len(vars) < 1 {
		return "incorrect url"
	}
        args := append(vars, "--xml")
        gCmd := exec.Command("gluster", args...)
        fmt.Print(vars, args, gCmd.Path, gCmd.Args)
	output, err := gCmd.CombinedOutput()
        if err != nil {
                return err.Error()
        }
	return string(output)
}

func main() {
	_, err := exec.LookPath("gluster")
	if err != nil {
		panic(err.Error())
	}
	gorest.RegisterService(new(GlusterService))
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":7331", nil)
}
