// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		PickUpFlagChanges()
		log.WithFields(log.Fields{"Log Level Flag": loglevel}).Debug("Entered Run cmd")
		fmt.Println("What XML file would you like to read from?")
		FileName := GetXMLFileName(".xml")
		log.WithFields(log.Fields{"FileName": FileName, "Location": "main func past GetXMLFileName"}).Debug("FileName retrieved from user")

		OpenXMLFile(FileName)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// OpenXMLFile Opens and XML File
func OpenXMLFile(FileName string) (xmlFile *os.File) {
	log.WithFields(log.Fields{"FileName": FileName, "Location": "Begin OpenXMLFile"}).Debug("Entering OpenXMLFile")
	xmlFile, err := os.Open(strings.TrimSpace(FileName)) // Open our xmlFile
	if err != nil {
		log.WithFields(log.Fields{"FileName": FileName, "Location": "open err check", "Error": err}).Error("Failed to open file")
		fmt.Println(color.HiRedString("%+v\n", err))
		return
	}
	fmt.Printf(color.GreenString("Successfully Opened %s\n", FileName))
	defer xmlFile.Close() // defer the closing of our xmlFile so that we can parse it later on
	return xmlFile
}

// IsPathExtensionNotValid just checks file extension type and if it is not valid then it rejects the file
func IsPathExtensionNotValid(FileName string, ValidExtension string) bool {
	log.WithFields(log.Fields{"FileName": FileName, "Location": "Inside IsPathExtensionNotValid func"}).Debug("Check FileName's validity")
	if filepath.Ext(strings.TrimSpace(FileName)) != ValidExtension {
		log.WithFields(log.Fields{"FileName": FileName, "Location": "Inside if statement IsPathExtensionNotValid func"}).Debug("FileName is not valid. Try again!")
		fmt.Println(color.HiRedString("Error! file must be xml file"))
		return true
	}
	log.WithFields(log.Fields{"FileName": FileName, "Location": "Past if statement IsPathExtensionNotValid func filename should be valid now"}).Debug("FileName is valid!")
	return false
}

// GetXMLFileName gets the name of the xml file we want to read from
func GetXMLFileName(ValidExtension string) string {
	var FileName string
	IsNotValid := true
	for IsNotValid {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		FileName, _ = reader.ReadString('\n')
		log.WithFields(log.Fields{"FileName": FileName, "Location": "Inside for loop GetXMLFileName func"}).Debug("FileName retrieved from user")
		IsNotValid = IsPathExtensionNotValid(FileName, ValidExtension)
	}
	log.WithFields(log.Fields{"FileName": FileName, "Location": "exiting GetXMLFileName func"}).Debug("FileName is valid; returning to main")
	return FileName
}

// PickUpFlagChanges picks up and applies any flags that have been passed via cli
func PickUpFlagChanges() {
	switch {
	case strings.Compare(loglevel, "debug") == 0:
		log.SetLevel(log.DebugLevel)
	case strings.Compare(loglevel, "info") == 0:
		log.SetLevel(log.InfoLevel)
	case strings.Compare(loglevel, "error") == 0:
		log.SetLevel(log.ErrorLevel)
	case strings.Compare(loglevel, "fatal") == 0:
		log.SetLevel(log.FatalLevel)
	case strings.Compare(loglevel, "panic") == 0:
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}

	switch {
	case strings.Compare(setOutput, "stdout") == 0:
		log.SetOutput(os.Stdout)
	}
}

// ReadFile reads an XML file passed to the function
func ReadFile(xmlFile *os.File) {
	byteValue, _ := ioutil.ReadAll(xmlFile)

}
