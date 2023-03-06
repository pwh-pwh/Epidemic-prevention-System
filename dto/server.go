package dto

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Sys struct {
	ComputerName string `json:"computerName"`
	ComputerIp   string `json:"computerIp"`
	UserDir      string `json:"userDir"`
	OsName       string `json:"osName"`
	OsArch       string `json:"osArch"`
}

type Cpu struct {
	//Total  float64 `json:"total"`
	Wait   float64 `json:"wait"` //not use
	Used   float64 `json:"used"` //
	Sys    float64 `json:"sys"`  // user/sys =8.2/6.7 比例生成
	Free   float64 `json:"free"` //100-used sys
	CpuNum int     `json:"cpuNum"`
}

type Mem struct {
	Total float64 `json:"total"`
	Used  float64 `json:"used"`
	Free  float64 `json:"free"`
	Usage float64 `json:"usage"`
}

type DirInfo struct {
	DirName     string  `json:"dirName"`
	SysTypeName string  `json:"sysTypeName"`
	TypeName    string  `json:"typeName"`
	Total       string  `json:"total"`
	Free        string  `json:"free"`
	Used        string  `json:"used"`
	Usage       float64 `json:"usage"`
}

type Go struct {
	Version   string `json:"version"` //GOVERSION
	Home      string `json:"home"`    //GOROOT
	Name      string `json:"name"`    //go
	StartTime string `json:"startTime"`
	RunTime   string `json:"runTime"`
}
type Serve struct {
	Cpu      *Cpu       `json:"cpu"`
	Mem      *Mem       `json:"mem"`
	Go       *Go        `json:"jvm"`
	Sys      *Sys       `json:"sys"`
	SysFiles []*DirInfo `json:"sysFiles"`
}

func GetServer() *Serve {
	server := Serve{}
	server.Cpu = GetCpu()
	//fmt.Println(server.Cpu)
	server.Mem = GetMem()
	server.Sys = GetSys()
	server.SysFiles = GetDirInfo()
	server.Go = GetGoInfo()
	return &server
}

func GetDirInfo() []*DirInfo {
	partitions, _ := disk.Partitions(false)
	var dirs []*DirInfo
	for _, p := range partitions {
		d := new(DirInfo)
		//fmt.Println(p)
		d.DirName = p.Mountpoint
		d.SysTypeName = p.Fstype
		d.TypeName = p.Device
		usage, _ := disk.Usage(p.Mountpoint)
		gb := float64(usage.Total) / float64(1024*1024*1024)
		d.Total = fmt.Sprintf("%v GB", gb)
		free := float64(usage.Free) / float64(1024*1024*1024)
		d.Free = fmt.Sprintf("%v GB", free)
		u := float64(usage.Used) / float64(1024*1024*1024)
		d.Used = fmt.Sprintf("%v GB", u)
		if gb == 0 {
			d.Usage = 0
		} else {
			d.Usage = u / gb
		}
		dirs = append(dirs, d)
	}
	return dirs
}

func GetGoInfo() *Go {
	g := new(Go)
	stdout := exec.Command("go", "env")
	output, err := stdout.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(stdout.String())
	sc := bufio.NewScanner(bytes.NewReader(output))
	for sc.Scan() {
		text := sc.Text()
		//fmt.Println(text)
		if strings.Contains(text, "GOVERSION") {
			index := strings.Index(text, "=")
			g.Version = text[index+1 : len(text)]
		}
		if strings.Contains(text, "GOROOT") {
			index := strings.Index(text, "=")
			g.Home = text[index+1 : len(text)]
		}
	}
	g.Name = "go"
	seconds := time.Since(common.StartTime).Seconds()
	g.StartTime = common.StartTime.Format(common.TimeFormat)
	g.RunTime = fmt.Sprintf("%v 秒", seconds)
	return g
}

func GetMem() *Mem {
	m := new(Mem)
	memory, _ := mem.VirtualMemory()
	total := float64(memory.Total) / (1024 * 1024 * 1024)
	used := float64(memory.Used) / (1024 * 1024 * 1024)
	free := float64(memory.Free) / float64(1024*1024*1024)
	usage := float64(used) / float64(total)
	m.Total = float64(total)
	m.Used = float64(used)
	m.Free = float64(free)
	m.Usage = usage
	//fmt.Printf("%v %v %v %v", total, used, free, usage)
	return m
}

func GetCpu() *Cpu {
	c := new(Cpu)
	c.CpuNum = runtime.NumCPU()
	percent, _ := cpu.Percent(time.Second, false)
	if len(percent) == 0 {
		percent = append(percent, 15.53)
	}
	c.Free = 100 - percent[0]
	//  user/sys =8.2/6.7 比例生成
	c.Used = percent[0] * 0.5503
	c.Sys = percent[0] * 4.496
	return c
}

func GetSys() *Sys {
	goarch := runtime.GOARCH
	goos := runtime.GOOS
	osName, _ := os.Hostname()
	dir, _ := os.Getwd()
	ip := GetLocalIp()
	//fmt.Printf("arch:%v, goos:%v, hostname:%v\n", goarch, goos, osName)
	//fmt.Printf("homedir:%v, ip:%v\n", dir, ip)
	sys := Sys{
		OsArch:       goarch,
		ComputerName: osName,
		OsName:       goos,
		UserDir:      dir,
		ComputerIp:   ip,
	}
	//fmt.Println(sys)
	return &sys
}

func GetLocalIp() string {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}
