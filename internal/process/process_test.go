package process

import (
	"gitlab.suse.de/jguilhermevanz/buddy/internal/testutil"
	"testing"
)

const commandSection = `#==[ Command ]======================================#
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

func Test_GetProcessData(t *testing.T) {
	tempDir := testutil.CreateTempDir()
	file := testutil.CreateBasicHealthCheckTestFileWithSectionsAt(tempDir, []string{commandSection})
	t.Logf("Processing file: %s", file)

	processes := GetProcessData(tempDir)

	if len(processes) != 9 {
		t.Fatalf("Wrong number of processes. Expected: %d, got %d", 9, len(processes))
	}
	expectedProcesses := map[string]map[string]string{
		"1":  {"ppid": "0", "user": "root", "state": "Rs", "cmd": "/usr/lib/systemd/systemd --switched-root --system --deserialize 23"},
		"4":  {"ppid": "2", "user": "root", "state": "S<", "cmd": "[kworker/0:0H]"},
		"7":  {"ppid": "2", "user": "root", "state": "S<", "cmd": "[mm_percpu_wq]"},
		"8":  {"ppid": "2", "user": "root", "state": "S", "cmd": "[ksoftirqd/0]"},
		"9":  {"ppid": "2", "user": "root", "state": "S", "cmd": "[rcu_sched]"},
		"10": {"ppid": "2", "user": "root", "state": "S", "cmd": "[rcu_bh]"},
		"11": {"ppid": "2", "user": "root", "state": "S", "cmd": "[migration/0]"},
		"12": {"ppid": "2", "user": "root", "state": "S", "cmd": "[watchdog/0]"},
		"2":  {"ppid": "0", "user": "root", "state": "S", "cmd": "[kthreadd]"},
	}
	for _, proc := range processes {
		if _, ok := expectedProcesses[proc.Pid]; !ok {
			t.Fatalf("PID %s not expected", proc.Pid)
		}
		if ppid, _ := expectedProcesses[proc.Pid]["ppid"]; ppid != proc.Ppid {
			t.Fatalf("PID %s has different ppid. Expected %s, got %s", proc.Pid, ppid, proc.Ppid)
		}
		if user, _ := expectedProcesses[proc.Pid]["user"]; user != proc.User {
			t.Fatalf("PID %s has different user. Expected %s, got %s", proc.Pid, user, proc.User)
		}
		if state, _ := expectedProcesses[proc.Pid]["state"]; state != proc.State {
			t.Fatalf("PID %s has different state. Expected %s, got %s", proc.Pid, state, proc.State)
		}
		if cmd, _ := expectedProcesses[proc.Pid]["cmd"]; cmd != proc.Cmd {
			t.Fatalf("PID %s has different cmd. Expected %s, got %s", proc.Pid, cmd, proc.Cmd)
		}
	}
}
