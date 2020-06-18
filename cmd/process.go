package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.suse.de/jguilhermevanz/buddy/internal/process"
	"log"
	"strings"
)

var (
	pid        string
	getProcess = &cobra.Command{
		Use:   "process",
		Short: "Get info about processes",
		Run: func(cmd *cobra.Command, args []string) {
			// get processes data
			processes := process.GetProcessData(getSupportutilsDir())

			if len(pid) > 0 {
				for _, proc := range processes {
					if proc.Pid == pid {
						_, err := fmt.Printf("PID: %s\tPPID: %s\tSTATE: %s\tCOMMAND: %s\n", proc.Pid, proc.Ppid, proc.State, proc.Cmd)
						if err != nil {
							log.Fatal(err)
						}
					}
				}
				return
			}

			zombiesCount := 0
			deadCount := 0

			zombiesProc := ""
			zombieWriter := bytes.NewBufferString(zombiesProc)
			deadsProc := ""
			deadWriter := bytes.NewBufferString(deadsProc)

			_, err := fmt.Fprintf(zombieWriter, "Zombies process:\n")
			if err != nil {
				log.Fatal(err)
			}
			_, err = fmt.Fprintf(deadWriter, "Uninterruptible process:\n")
			if err != nil {
				log.Fatal(err)
			}
			for _, proc := range processes {
				if strings.Contains(proc.State, "Z") {
					zombiesCount++
					_, err := fmt.Fprintf(zombieWriter, "PID: %s\tPPID: %s\tSTATE: %s\tCOMMAND: %s\n", proc.Pid, proc.Ppid, proc.State, proc.Cmd)
					if err != nil {
						log.Fatal(err)
					}
				}
				if strings.Contains(proc.State, "D") {
					deadCount++
					_, err := fmt.Fprintf(deadWriter, "PID: %s\tPPID: %s\tSTATE: %s\tCOMMAND: %s\n", proc.Pid, proc.Ppid, proc.State, proc.Cmd)
					if err != nil {
						log.Fatal(err)
					}
				}
			}

			fmt.Fprintf(zombieWriter, "Zombies total: %d\n", zombiesCount)
			fmt.Fprintf(deadWriter, "Uninterruptible total: %d\n", deadCount)
			fmt.Print(zombieWriter.String())
			fmt.Print(deadWriter.String())
		},
	}
)

func init() {
	getProcess.PersistentFlags().StringVarP(&pid, "pid", "p", "", "Filter by PID")

}
