package lcuapi

import (
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const MaxTickTimes = 10

func GetUxProcessByPsutil() (*process.Process, error) {
	//defer func() func() {
	//	start := time.Now()
	//	return func() {
	//		log.Printf("GetUxProcess cost %v", time.Since(start))
	//	}
	//}()()

	var (
		processChan = make(chan *process.Process)
		//tickChan    = time.Tick(500 * time.Millisecond)
		tickTimes = 0
	)
	defer close(processChan)

	//for {
	//	select {
	//	case <-tickChan:
	//if tickTimes >= MaxTickTimes {
	//	return nil, errors.New("LeagueClientUx.exe not found")
	//}
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	for _, p := range processes {
		//TODO wait util gopsutil implement Status()
		//
		//status, err := p.Status()
		//if err != nil {
		//	log.Printf("Error getting process status: %v", err)
		//	continue
		//}
		//if status == "Z" {
		//	continue
		//}

		name, err := p.Name()
		if err != nil {
			log.Printf("Error getting process name: %v", err)
			continue
		}

		if name == "LeagueClientUx.exe" || name == "LeagueClientUx" {
			return p, nil
		}
	}
	tickTimes++
	//		continue
	//	}
	//}
	return nil, nil
}

func GetRiotClientProcessByPsutil() (*process.Process, error) {
	//defer func() func() {
	//	start := time.Now()
	//	return func() {
	//		log.Printf("GetUxProcess cost %v", time.Since(start))
	//	}
	//}()()

	var (
		processChan = make(chan *process.Process)
		//tickChan    = time.Tick(500 * time.Millisecond)
		tickTimes = 0
	)
	defer close(processChan)

	//for {
	//	select {
	//	case <-tickChan:
	//if tickTimes >= MaxTickTimes {
	//	return nil, errors.New("LeagueClientUx.exe not found")
	//}
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	for _, p := range processes {
		//TODO wait util gopsutil implement Status()
		//
		//status, err := p.Status()
		//if err != nil {
		//	log.Printf("Error getting process status: %v", err)
		//	continue
		//}
		//if status == "Z" {
		//	continue
		//}

		name, err := p.Name()
		if err != nil {
			log.Printf("Error getting process name: %v", err)
			continue
		}

		if name == "RiotClientServices.exe" || name == "RiotClientServices" {
			return p, nil
		}
	}
	tickTimes++
	//		continue
	//	}
	//}
	return nil, nil

}

func GetUxProcessCommandlineMapByCmd() (mp map[string]string, err error) {
	args := []string{
		"/c",
		"wmic PROCESS WHERE name='LeagueClientUx.exe' GET commandline",
	}
	cmd := exec.Command("cmd", args...)
	res, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	// 通过判断是否有CommandLine来判断是否启动客户端
	var FoundClientExe = strings.HasPrefix(string(res), "CommandLine")

	mp = flagsToMap(string(res))
	if len(mp) == 0 {
		err = errors.New("need admin")
		if !FoundClientExe {
			err = errors.New("can't find client")
		}
		return
	}
	return
}

func getPIDByProcessName(processName string) (int, error) {
	cmd := exec.Command("tasklist", "/fo", "csv", "/nh")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		fields := strings.Split(line, ",")
		if len(fields) >= 2 {
			name := strings.Trim(fields[0], "\"")                   // 进程名，去除双引号
			pid := strings.TrimSpace(strings.Trim(fields[1], "\"")) // PID，去除空格与双引号
			if strings.EqualFold(processName, name) {
				return strconv.Atoi(pid)
			}
		}
	}

	return 0, fmt.Errorf("process not found: %s", processName)
}

func flagsToMap(res string) map[string]string {
	// 使用正则表达式来提取键值对
	re := regexp.MustCompile(`--([^= ]+)=([^ ]+)`)
	matches := re.FindAllStringSubmatch(res, -1)

	// 创建一个map来存储键值对
	configMap := make(map[string]string)

	// 将提取的键值对存储到map中
	for _, match := range matches {
		key := match[1]
		value := match[2]
		configMap[key] = value
	}

	//// 打印map中的键值对
	//for key, value := range configMap {
	//	fmt.Printf("%s: %s\n", key, value)
	//}
	return configMap
}
