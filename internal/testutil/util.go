// Package with util functions used in the tests.
package testutil

import (
	"io/ioutil"
	"log"
	"os"
)

const (
	commandSection = `#==[ Command ]======================================#
# /bin/ps axwwo user,pid,ppid,%cpu,%mem,vsz,rss,stat,time,cmd
USER       PID  PPID %CPU %MEM    VSZ   RSS STAT     TIME CMD
root         1     0 45.1  0.0 250192 38856 Rs   8-03:50:41 /usr/lib/systemd/systemd --switched-root --system --deserialize 23
root         2     0  0.0  0.0      0     0 S    00:00:09 [kthreadd]
root         4     2  0.0  0.0      0     0 S<   00:00:00 [kworker/0:0H]
root         7     2  0.0  0.0      0     0 S<   00:00:00 [mm_percpu_wq]
root         8     2  0.0  0.0      0     0 S    00:03:45 [ksoftirqd/0]
root         9     2  0.1  0.0      0     0 S    00:49:50 [rcu_sched]
root        10     2  0.0  0.0      0     0 S    00:00:00 [rcu_bh]
root        11     2  0.0  0.0      0     0 S    00:00:32 [migration/0]
root        12     2  0.0  0.0      0     0 S    00:00:04 [watchdog/0]

`
	summarySection = `#==[ Summary ]======================================#
# AppArmor REJECT Messages
* apparmor.service - Load AppArmor profiles
   Loaded: loaded (/usr/lib/systemd/system/apparmor.service; enabled; vendor preset: enabled)
   Active: active (exited) since Mon 2020-05-25 15:10:11 CEST; 2 weeks 4 days ago
 Main PID: 1070 (code=exited, status=0/SUCCESS)
    Tasks: 0
   CGroup: /system.slice/apparmor.service

Warning: Journal has been rotated since unit was started. Log output is incomplete or unavailable.

0 Reject Messages, AppArmor Module: Loaded  

`
	TESTING_PATTERN    = "buddy_testing"
	BASIC_HEALTH_CHECK = "basic-health-check.txt"
)

func createBasicHealthCheckTestFile(tempFile *os.File, commandSectionCount int, summarySectionCount int, sections []string) {
	for i := 0; i < commandSectionCount; i++ {
		if _, err := tempFile.Write([]byte(commandSection)); err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < summarySectionCount; i++ {
		if _, err := tempFile.Write([]byte(summarySection)); err != nil {
			log.Fatal(err)
		}
	}

	for _, section := range sections {
		if _, err := tempFile.Write([]byte(section)); err != nil {
			log.Fatal(err)
		}
	}
}

func CreateBasicHealthCheckTestFile(commandSectionCount int, summarySectionCount int) string {
	tempDir := CreateTempDir()
	tempFile, err := os.OpenFile(tempDir+"/"+BASIC_HEALTH_CHECK, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer tempFile.Close()

	createBasicHealthCheckTestFile(tempFile, commandSectionCount, summarySectionCount, []string{})
	return tempFile.Name()
}

func CreateBasicHealthCheckTestFileWithSectionsAt(directory string, sections []string) string {
	tempFile, err := os.OpenFile(directory+"/"+BASIC_HEALTH_CHECK, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer tempFile.Close()

	createBasicHealthCheckTestFile(tempFile, 0, 0, sections)

	return tempFile.Name()
}

func CreateTempDir() string {
	dir, err := ioutil.TempDir("", TESTING_PATTERN)
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
