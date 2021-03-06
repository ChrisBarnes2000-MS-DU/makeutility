
package cmd

import (
    "testing"
)


func TestGetTimeStamp(t *testing.T) {
    if getTimeStamp(`2009-01-01T01:02:01.111+02:00`) != 1230764521111 {
        t.Error("Time Stamp For 2009-01-01T01:02:01.111+02:00 Is Expected To Be 1230764521111")
    }
}



func TestTableGetTimeStamp(t *testing.T) {
    var tests = []struct {
        input    string
        expected int
    }{
        {`2009-01-01T01:02:01.111+02:00`, 1230764521111},
        {`2021-03-31T01:02:01.111+02:00`, 1617145321111},
        {`2021-04-04T12:00:00.000-08:00`, 1617566400000}, // 12:00pm PST 
        {`2021-04-04T14:00:00.000-08:00`, 1617573600000},	// 2:00pm PST
        {`2021-04-04T17:00:00.000-08:00`, 1617584400000}, // 5:00pm PST
    }

    for _, test := range tests {
        if output := getTimeStamp(test.input); output != test.expected {
            t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
        }
    }
}

