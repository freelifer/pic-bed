package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

    "gopkg.in/ini.v1"
	"github.com/gin-gonic/gin"
	"github.com/freelifer/pic-bed/api"
)

var cfg *ini.File
func settings(c *gin.Context) {
	
	c.JSON(200, gin.H{
		"message": "settings",
	})
}

func main() {
	// 本地配置初始化
	cfg = configInit()

	// gin初始化
	ginInit()
}


func configInit() *ini.File{
	homeDir, err := Home()
	if nil != err {
		panic(err)
	}

	fmt.Println(homeDir)
	folderPath := filepath.Join(homeDir, ".picbed")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.Mkdir(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	configpath := filepath.Join(folderPath, "conf.ini")
	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		os.Create(configpath)
	}

	cfg, err := ini.Load(configpath)
	if err != nil {
		panic(err)
	}
	return cfg
}

func ginInit() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiGroup:=r.Group("/api")

	api.Route(apiGroup, "settings", &api.SettingsController{})
	
	r.Run(":54321") // listen and serve on 0.0.0.0:8080
}

// Home returns the home directory for the executing user.
//
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
