package common

import (
	"gitlab.suse.de/jguilhermevanz/buddy/internal/testutil"
	"testing"
)

const TOTAL_SECTION = 5

func Test_getFileSections(t *testing.T) {
	resultChannel := make(chan string)
	file := testutil.CreateBasicHealthCheckTestFile(TOTAL_SECTION, TOTAL_SECTION)
	t.Logf("Processing file: %s", file)

	go getFileSections(file, resultChannel)

	count := 0
	for range resultChannel {
		t.Log("Section found!")
		count++
	}
	expected_sections_count := TOTAL_SECTION * 2
	if count != expected_sections_count {
		t.Fatalf("Section count is wrong. Expected %d, got %d", expected_sections_count, count)
	}
}

func Test_GetCommand(t *testing.T) {
	cmdChannel := make(chan string)
	file := testutil.CreateBasicHealthCheckTestFile(TOTAL_SECTION, 1)
	t.Logf("Processing file: %s", file)

	go GetCommand(file, cmdChannel)

	count := 0
	for range cmdChannel {
		t.Log("Command found!")
		count++
	}
	if count != TOTAL_SECTION {
		t.Fatalf("Section count is wrong. Expected %d, got %d", TOTAL_SECTION, count)
	}
}
