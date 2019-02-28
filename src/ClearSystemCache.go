package main
/*
This is a program that clears your IE cache as well as clears your USD(unified Service Desk ) Cache

*/
//https://sausheong.github.io/posts/how-to-build-a-simple-artificial-neural-network-with-go/
import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func executeKillUSD() bool {
	exec.Command("TASKKILL", "/F", "/IM", "UnifiedServiceDesk.exe").Output()
	return true
}

func executeClearUSDCache() bool {
	//clear the USD configuration

	_appdataDir := os.Getenv("appdata")
	_usdCacheDir := filepath.Join(_appdataDir, "Microsoft", "USD")
	dir, err := os.Open(_usdCacheDir)
	defer dir.Close()
	if err != nil {
		fmt.Println("ERROR OCCURED Opening USD Cache folder: ", err)
		return false
	}

	_fileToDel := "Default_USD.tokens.dat"
	delErr := os.Remove(filepath.Join(_usdCacheDir, _fileToDel))

	if delErr != nil {
		fmt.Println("USD cache already Cleared ")
		return false
	}
	//fileList, _ := dir.Readdirnames(-1)

	/*
		for _, file := range fileList {
			fmt.Println("Removing: ", file)
			_rmErr := os.Remove(filepath.Join(_usdCacheDir, file))

			if _rmErr != nil {
				fmt.Println("ERROR OCCURED DELETING: ", _rmErr)
				return false
			}
		}
	*/

	return true

}
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
	executeKillUSD()

	//clearing temp Files
	_, err := exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "8").Output()
	//exec.Command("rundll32.exe", "InetCpl.cpl,ClearMyTracksByProcess", "8").Output()
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
			//fmt.Println("Detected java installed.. clearing its cache..")
			//exec.Command("javaws", "-uninstall", "1").Output()
			break //exit the loop
		}
	}

	if err != nil {
		//fmt.Println("ERROR: ", err)
		return false
	}
	//fmt.Printf("Command output : %s", out)
	return true
}

func main() {
	fmt.Println("Welcome to AD Cache Clearing Tool, Made in GOLANG!")
	executeKillUSD()
	executeClearUSDCache()
	executeClearIECache()

	fmt.Println("[!] Successfully Finished Execution, you are all set...\nProgram will close in 5 seconds! whoohoo")

	duration := time.Duration(5) * time.Second // Pause for 10 seconds
	time.Sleep(duration)

}
