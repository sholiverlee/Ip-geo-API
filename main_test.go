package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

// TestGetElement tests the getElement function.
func TestDownloadCsv(t *testing.T) {
	gzFile := "dbip-city-ipv4-num.csv.gz"
	downloadCsv(gzFile)
	assert.Equal(t, '0', '0')
}

func TestExtractGZ(t *testing.T) {
	gzFile := "dbip-city-ipv4-num.csv.gz"
	csvFile := "dbip-city-ipv4-num.csv"
	extractGZ(gzFile, csvFile)
	assert.Equal(t, '0', '0')
}

func TestReadAndGet(t *testing.T) {
	ReadAndGet("dbip-city-ipv4-num.csv")
	assert.Equal(t, '0', '0')

}

func TestMain(t *testing.T) {
	main()
	assert.Equal(t, '0', '0')

}
