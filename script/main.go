package main

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/nxadm/tail"
	"github.com/smira/go-statsd"
)

var inputLogFile = "./input.log"
var outputLogFile = "/etc/fluent-bit/output.log"

func main() {
	// Track changes in input file
	t, err := tail.TailFile(inputLogFile, tail.Config{
		Follow:   true,
		ReOpen:   true,
		Location: &tail.SeekInfo{Offset: 0, Whence: io.SeekEnd}, // <- line changed
	})
	if err != nil {
		panic(err)
	}

	for line := range t.Lines {
		log := strings.Fields(line.Text)

		timestamp := log[0] + " " + log[1]
		serviceName := log[2]
		statusCode := log[3]
		responseTime, _ := strconv.Atoi(log[4][:len(log[4])-2])
		userId := log[5]
		transactionId := log[6]
		additionalInfo := strings.Join(log[7:], " ")

		// Send relevant metrics
		client := statsd.NewClient("localhost:8125")
		client.Incr("log_request_count", 1)
		client.PrecisionTiming("log_response_time", time.Duration(responseTime*1000000))
		defer client.Close()

		// Prepare logs in structured format (JSON)
		logs := map[string]interface{}{
			"timestamp":       timestamp,
			"service_name":    serviceName,
			"status_code":     statusCode,
			"response_time":   responseTime,
			"user_id":         userId,
			"transaction_id":  transactionId,
			"additional_info": additionalInfo,
			"message":         additionalInfo,
		}

		// Output logs in JSON format
		logsJSON, err := json.Marshal(logs)
		if err != nil {
			panic(err)
		}
		print(string(logsJSON))

		// Write to output file
		f, err := os.OpenFile(outputLogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(string(logsJSON) + "\n"); err != nil {
			panic(err)
		}

	}
}
