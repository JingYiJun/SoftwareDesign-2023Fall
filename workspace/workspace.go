package workspace

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	. "design/interfaces"
)

type Workspace struct {
	FileName               string
	Dirty                  bool // if Dirty, save to file
	UndoableCommandHistory []UndoableCommand
	UndoableCommandPointer int
	FileContent            []string
}

var CurrentWorkspace Workspace
var allWorkspaces map[string]Workspace
var path string

func init() {
	allWorkspaces = make(map[string]Workspace)
	path = "./files/"
}

func updateWorkspace(curWorkspace Workspace) {
	_, ok := allWorkspaces[curWorkspace.FileName]
	if ok {
		allWorkspaces[curWorkspace.FileName] = curWorkspace
	}
	// if nil
	// else ?
}

func (curWorkspace *Workspace) Load(fileName string) (*Workspace, error) {
	if curWorkspace == nil {
		return nil, errors.New("load: curWorkspace is nil")
	}
	updateWorkspace(*curWorkspace)

	ws, ok := allWorkspaces[fileName]
	if ok {
		CurrentWorkspace = ws
		return &CurrentWorkspace, nil
	}

	ws = Workspace{
		FileName:               fileName,
		Dirty:                  false,
		UndoableCommandHistory: make([]UndoableCommand, 0),
		UndoableCommandPointer: 0,
		FileContent:            make([]string, 0),
	}

	filePath := path + fileName

	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	reader := bufio.NewReader(file)
	for {
		content, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				return nil, errors.New(err.Error())
			}

			if content == "" {
				break
			} else {
				ws.FileContent = append(ws.FileContent, content)
				break
			}
			// 注意此处还要处理最后一行
		}
		// if content == "" {
		// 	continue
		// }
		ws.FileContent = append(ws.FileContent, content)
	}
	CurrentWorkspace = ws
	return &CurrentWorkspace, nil
}

func (curWorkspace *Workspace) Open(fileName string) (*Workspace, error) {
	if curWorkspace == nil {
		return nil, errors.New("open: curWorkspace is nil")
	}
	updateWorkspace(*curWorkspace)
	_, ok := allWorkspaces[fileName]
	if ok {
		CurrentWorkspace = allWorkspaces[fileName]
		return &CurrentWorkspace, nil
	}
	return nil, errors.New("open: no such file")
}

func (curWorkspace *Workspace) Close(fileName string) error {
	if curWorkspace == nil {
		return errors.New("close: curWorkspace is nil")
	}
	if curWorkspace.Dirty {
		fmt.Println("Do you want to save the current workspace [Y\\N] ？")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			return errors.New(err.Error())
		}
		if input == "Y" || input == "y" {
			updateWorkspace(*curWorkspace)
		}
	}

	_, ok := allWorkspaces[fileName]
	if ok {
		delete(allWorkspaces, fileName)
		return nil
	}
	return errors.New("close: no such file")

}

func (curWorkspace *Workspace) ListWorkspace() error {
	for fileName := range allWorkspaces {
		fmt.Println(fileName)
	}
	return nil
}
