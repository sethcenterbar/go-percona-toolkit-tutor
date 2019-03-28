package utilities

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"

	"github.com/mgutz/ansi"
	"github.com/sethcenterbar/percona-toolkit-tutor/structs"
	yaml "gopkg.in/yaml.v2"
)

func OpenToolboxWeb() structs.Toolbox {
	// Establish file location
	myself, error := user.Current()
	if error != nil {
		panic(error)
	}
	homedir := myself.HomeDir
	fileLocation := homedir + "/.percona-toolkit-trainer.yaml"

	var toolbox structs.Toolbox

	// Attempt to open file
	_, err := os.Open(fileLocation)
	if err != nil {
		if os.IsNotExist(err) {
			println("You don't already have the file, downloading file from github gist..")
			url := "https://gist.githubusercontent.com/sethcenterbar/620e9cd0f7288f91862a9763863e16e3/raw/ccb61dc0cf700e543facf7fee501356f77ad4db3/toolkit.yaml"
			reader := strings.NewReader(`{"body":123}`)
			request, err := http.NewRequest("GET", url, reader)
			if err != nil {
				panic("Error!")
			}
			client := &http.Client{}
			resp, err := client.Do(request)
			response, err := ioutil.ReadAll(resp.Body)

			yaml.Unmarshal(response, &toolbox)

			// Write file so we don't have to pull it every run
			permissions := os.FileMode(0644)
			err = ioutil.WriteFile(fileLocation, response, permissions)
			if err != nil {
				panic("Couldn't write file!")
			}
		} else {
			panic("A serious error occured.")
		}
	} else {
		myfile, err := ioutil.ReadFile(fileLocation)
		if err != nil {
			panic("aaah!")
		}
		yaml.Unmarshal(myfile, &toolbox)
	}

	return toolbox
}

func GrabTool(t structs.Toolbox, tool string) (structs.Tool, error) {
	for _, x := range t.Tools {
		if x.Name == tool {
			thisTool := x
			return thisTool, nil
		}
	}
	return structs.Tool{}, errors.New(tool + " is not a valid tool in the toolkit.")
}

func ListTools(t structs.Toolbox) string {
	output := "\n" + ansi.Color("The currently supported tools are: ", "green") + "\n"
	for _, tool := range t.Tools {
		output += "  " + tool.Name + "\n"
	}
	return output
}

func ValidateTool(tb structs.Toolbox, toolname string) bool {
	for _, tool := range tb.Tools {
		if tool.Name == toolname {
			return true
		}
	}
	return false
}
