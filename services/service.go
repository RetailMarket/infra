package main

import (
	"path"
	"os/exec"
	"os"
	"fmt"
	"strconv"
	"io/ioutil"
)

const rootPath = ".."

func main() {
	args := os.Args[1:]
	repos := []string{"priceManager", "workflow", "priceSync", "workflowSync", "priceWeb"}
	switch args[0] {
	case "start":
		start(args[1], repos)
	case "stop":
		stop(args[1], repos)
	}
}

func startService(repoPath string) {
	pidPath := path.Join(repoPath, "server.pid")
	mainFile := path.Join(repoPath, "app/main/main.go")

	if _, err := os.Stat(path.Join(rootPath, mainFile)); os.IsNotExist(err) {
		command := exec.Command("go", "run", path.Join(repoPath, "app/main/main.go"))
		err := command.Start();
		if (err != nil) {
			fmt.Printf("Failed to process: %s %s", err, string(""))
		} else {
			file, err := os.Create(pidPath)
			if (err != nil) {
				fmt.Errorf("failed to creat pid file: %s", err)
			}
			file.Write([]byte(strconv.Itoa(command.Process.Pid)))
		}
	} else {
		fmt.Printf("main file doesn't exist: %v\n", mainFile)
	}
}

func start(service string, repos []string) {
	if (service == "all") {
		for _, repo := range repos {
			repoPath := path.Join(rootPath, repo)
			startService(repoPath)
		}
	} else {
		repoPath := path.Join(rootPath, service)
		startService(repoPath)
	}
}

func stop(service string, repos []string) {
	if (service == "all") {
		for _, repo := range repos {
			repoPath := path.Join(rootPath, repo)
			stopService(repoPath)
		}
	} else {
		repoPath := path.Join(rootPath, service)
		stopService(repoPath)
	}
}

func stopService(repoPath string) {
	pidPath := path.Join(repoPath, "server.pid")
	mainFile := path.Join(repoPath, "app/main/main.go")

	if _, err := os.Stat(path.Join(rootPath, mainFile)); os.IsNotExist(err) {
		pidBytes, err := ioutil.ReadFile(pidPath)
		if (err != nil) {
			fmt.Errorf("failed to creat pid file: %s", err)
		}
		pid, err := strconv.ParseInt(string(pidBytes), 10, 64)
		if (err != nil) {
			fmt.Errorf("wrong pid in file %s: %s", pidPath, err)
		}
		process := os.Process{Pid:int(pid)}
		err = process.Kill()
		if (err != nil) {
			fmt.Errorf("failed to kill pid %s: %s", string(pid), err)
		}
		err = os.Remove(pidPath)
		if (err != nil) {
			fmt.Errorf("failed to remove pid file %s: %s", pidPath, err)
		}
	} else {
		fmt.Printf("main file doesn't exist: %v\n", mainFile)
	}
}