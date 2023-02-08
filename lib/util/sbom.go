package util

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/anchore/syft/syft"
	"github.com/anchore/syft/syft/sbom"
)

func ConvertToSPDXJson(sbom *sbom.SBOM) {
	local, _ := ioutil.TempDir("", "go-vcs")

	bytes, err := syft.Encode(*sbom, syft.FormatByName("spdx-2-json"))
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create(local + "/spdx.json")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(string(bytes))

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(local + "/spdx.json" + " done")
}
