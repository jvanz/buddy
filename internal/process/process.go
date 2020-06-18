package process

import (
	"bufio"
	"gitlab.suse.de/jguilhermevanz/buddy/internal/common"
	"log"
	"regexp"
	"strings"
	"sync"
)

const (
	PROCESS_FILE = "basic-health-check.txt"
	PROC_COMMAND = "bin/ps axwwo user,pid,ppid,%cpu,%mem,vsz,rss,stat,time,cmd"
)

type Process struct {
	User  string
	Pid   string
	Ppid  string
	State string
	Cmd   string
}

// GetProcessData read the basic-health-check.txt inside the given supportutils
// directory and process data containing the running processes.
// Returns a slice of the running processes.
func GetProcessData(supportutils string) []*Process {
	cmdData, err := common.FindCommand(supportutils+"/"+PROCESS_FILE, PROC_COMMAND)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(cmdData))
	scanner.Scan() // ignore the first line, the header
	scanner.Scan() // ignore the first line, the header
	scanner.Scan() // ignore the first line, the header
	regex := regexp.MustCompile(`^(?P<user>\S+)\s+(?P<pid>\S+)\s+(?P<ppid>\S+)\s+(?P<cpu>\S+)\s+(?P<memory>\S+)\s+(?P<vsz>\S+)\s+(?P<rss>\S+)\s+(?P<state>\S+)\s+(?P<time>\S+)\s+(?P<cmd>.+)$`)

	processes := []*Process{}
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for scanner.Scan() {
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
		procline := scanner.Text()
		wg.Add(1)
		go func(procline string) {
			matches := regex.FindStringSubmatch(procline)
			process := &Process{
				matches[1],
				matches[2],
				matches[3],
				matches[8],
				matches[10],
			}
			mutex.Lock()
			processes = append(processes, process)
			mutex.Unlock()
			wg.Done()

		}(procline)
	}
	wg.Wait()
	return processes
}
