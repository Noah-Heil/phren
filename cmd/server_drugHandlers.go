package cmd

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

var drugdb Drugbanktype

// DrugHandler returns (in memory)
func DrugHandler(c *gin.Context) {
	// c.Header("Content-Type", "application/json")
	drug := c.Param("drug")

	c.String(http.StatusOK, "Hello %s", drug)
}

// RootDrugHandler returns (in memory)
func RootDrugHandler(c *gin.Context) {
	// c.Header("Content-Type", "application/json")

	c.String(http.StatusOK, "Hello Drug API \n")

	// Open our xmlFile
	xmlFile, err := os.Open("db/testdb.xml")

	// if os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we unmarshal our byteArray which contains our
	// xmlFiles content
	xml.Unmarshal(byteValue, &drugdb)

	c.String(http.StatusOK, spew.Sdump(drugdb))

	// drugdb.MarshalXML()
}

// DrugClassificationsHandler returns (in memory)
func DrugClassificationsHandler(c *gin.Context) {
	// c.Header("Content-Type", "application/json")

	c.String(http.StatusOK, "Hello Drug Classifications API")
}
