package main

import (
	"fmt"

	"os"
	"os/exec"

	"strings"
)

func executeClearIECache() bool {
	/*
		_command := exec.Command("Cmd", "Rundll32.exe", "inetCpl.cpl,ClearMyTracksByProcess 8")
		_command.Run()
	*/

	//out, err := exec.Command("cmd", "rundll32.exe", "inetCpl.cpl,ClearMyTracksByProcess 8").Output()\

	//each argument needs top be seperated as its own string (comma seperated)

	//Temporary Internet Files
	/*
		out, err := exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "8").Output()
		out, err := exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "2").Output()
		out, err := exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "1").Output()
		out, err := exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "16").Output()
		out, err := exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "32").Output()
		out, err := exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "1").Output()

		out, err := exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "255").Output()
	*/
	//clearing temp Files
	out, err := exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "8").Output()
	//clearing Cookies
	exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "2").Output()
	//clearing History
	exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "1").Output()
	//Form Data
	exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "16").Output()
	//Passwords
	exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "32").Output()
	//Delete All
	exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "1").Output()

	exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "255").Output()

	//clear java see if the machine even has javaws installed

	//_javapath = filepath.Join(os.Getenv("programdata"), "Oracle", "Java", "javapath");
	_pathVars := os.Getenv("path")
	//split the path vars
	var _pathSplit []string = strings.Split(_pathVars, ";")

	//iterate each pathname and see if a javapath Exist
	for _, path := range _pathSplit {
		if strings.Contains(path, "javapath") {
			//we found out that the current machine has javapath intalled
			//we can clear the java cache

			exec.Command("javaws", "-uninstall", "1").Output()
			break //exit the loop
		}
	}

	if err != nil {
		fmt.Println("ERROR: ", err)
		return false
	}
	fmt.Printf("Command output : %s", out)
	return true
}

func main() {
	executeClearIECache()
}
