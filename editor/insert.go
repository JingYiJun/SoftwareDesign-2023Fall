package editor

import (
	"design/util"
	"errors"
)

func Insert(lineNum int, content string, fileContent *[]string) (int, error) {
	if !util.IsInit() {
		return 0, errors.New("insert: No file in workspace")
	}
	if lineNum == -1 {
		*fileContent = append(*fileContent, content)
		return len(*fileContent), nil
	}
	if lineNum > len(*fileContent)+1 || lineNum < 1 {
		return 0, errors.New("insert: line number out of range")
	}

	// n-1 node before newline
	// get the 0~n-2 lines
	// fileContent = slices.Insert(fileContent, lineNum, newNode.Node2String())
	*fileContent = append((*fileContent)[:lineNum-1], append([]string{content}, (*fileContent)[lineNum-1:]...)...)

	return lineNum, nil
}
