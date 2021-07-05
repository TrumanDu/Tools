package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()

	helper := &ReplaceHelper{
		Root:    flag.Arg(0),
		OldText: flag.Arg(1),
		NewText: flag.Arg(2),
	}

	err := filepath.Walk(helper.Root, helper.walkCallback)
	if err == nil {
		fmt.Println("done!")
	} else {
		fmt.Println("error:", err.Error())
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type ReplaceHelper struct {
	Root    string
	OldText string
	NewText string
}

func (helper *ReplaceHelper) walkCallback(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}
	if info.IsDir() {
		return nil
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(buf)
	newContent := strings.Replace(content, helper.OldText, helper.NewText, -1)

	err = ioutil.WriteFile(path, []byte(newContent), 0)
	if err != nil {
		return err
	}
	fmt.Println(path + " done.")
	return nil
}
