package workspace

import (
	"errors"
	"fmt"
)

func (curWorkspace *Workspace) Exit() error {
	var dirty bool
	for _, workspace := range allWorkspaces {
		if workspace.Dirty {
			dirty = true
			break
		}
	}
	if dirty {
		fmt.Println("Do you want to save the unsaved workspace [Y\\N] ？")
		var input string
		for {
			_, err := fmt.Scanln(&input)
			if err != nil {
				return errors.New(err.Error())
			}
			if input == "Y" || input == "y" {
				for _, workspace := range allWorkspaces {
					if workspace.Dirty {
						err = workspace.Save()
						if err != nil {
							return err
						}
					}
				}
				break
			} else if input == "N" || input == "n" {
				break
			} else {
				fmt.Println("Please input Y or N")
			}
		}
	}

	allWorkspaces = make(map[string]Workspace)
	//ToDo
	*curWorkspace = Workspace{}
	return nil
}
