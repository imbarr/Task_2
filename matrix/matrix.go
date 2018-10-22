package matrix

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"regexp"
)

type BinaryMatrix struct {
	height int
	width int
	values [][]bool
}

func getSplitFunc(delim string) bufio.SplitFunc, error {
	quoted = QuoteMeta(delim)
	regex, e := regexp.Compile("[" + delim + "]*([" + delim + "]+)([" + delim + "]*)")
	if(e != nul)
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}


		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return len(data), dropCR(data), nil
		}
		// Request more data.
		return 0, nil, nil
	}
}

func ReadMatrix(reader io.Reader) (BinaryMatrix, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(ScanC)


	var empty BinaryMatrix
	str, e := reader.(' ')
	if e != nil { return empty, e }
	height, e := strconv.Atoi(str)
	if e != nil { return empty, errors.New("failed to read height: " + e.Error()) }

	str, e = reader.ReadString('\n')
	if e != nil { return empty, e }
	width, e := strconv.Atoi(str)
	if e != nil { return empty, errors.New("failed to read width: " + e.Error()) }

	values := make([][]bool, height, width)
	for i := 0; i < height; i++ {
		for j := 0; j < width - 1; j++ {
			values[i][j], e = readValue(reader, ' ')
			if e != nil { return empty, e }
		}
		values[i][width - 1], e = readValue(reader, '\n')
		if e != nil { return empty, e }
	}
	return BinaryMatrix{height:height, width:width, values:values}, nil
}

func readValue(r bufio.Reader, delim byte) (bool, error) {
	str, e := r.ReadString(delim)
	if e != nil { return false, e }
	value, e := toBoolean(str)
	if e != nil { return false, e }
	return value, nil
}

func toBoolean(str string) (bool, error) {
	value, e :=  strconv.Atoi(str)
	if e != nil {
		return false, errors.New("failed converting value to bool: " + e.Error())
	}
	if value == 1 {
		return true, nil
	} else if value == 0 {
		return  false, nil
	}
	return false, errors.New("failed converting value to bool: value is not 0 or 1")
}