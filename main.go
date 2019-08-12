package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	not_last_element = "├───"
	last_element = "└───"
	tab_space = "	"
)

func dirTreeHelp (output *bytes.Buffer, path_to_repr string, tab_passed string, show_files bool) error {

	files, err := ioutil.ReadDir(path_to_repr)
	if err != nil {
		log.Fatal(err)
	}

	var final_element string
	if show_files {
		for _, file := range files {
			final_element = file.Name()
		}
	} else {
		for _, file := range files {
			if file.IsDir() {
				final_element = file.Name()
			}
		}
	}

	for _, file := range files {

		var tab_to_add string = ""
		var tab_to_pass string = ""

		if file.Name() == final_element {
			tab_to_add = last_element
			tab_to_pass = " " + tab_space
		} else {
			tab_to_add = not_last_element
			tab_to_pass = "│" + tab_space
		}

		tab_to_pass = tab_passed + tab_to_pass

		/*
			if:
				demonstrate files && it is a file
			else if:
				it is a directory
		*/
		if show_files && !file.IsDir() {
			var file_size string
			if file.Size() != 0 {
				file_size = fmt.Sprintf("%vb", file.Size())
			} else {
				file_size = "empty"
			}
			file_reprsentation := fmt.Sprintf("%v%v%v (%v)\n", tab_passed, tab_to_add, file.Name(), file_size)
			if _, err = output.WriteString(file_reprsentation); err != nil {
				log.Fatal(err)
			}
		} else if file.IsDir() {
			directory_reprsentation := fmt.Sprintf("%v%v%v\n", tab_passed, tab_to_add, file.Name())
			if _, err = output.WriteString(directory_reprsentation); err != nil {
				log.Fatal(err)
			}

			next_path := fmt.Sprintf("%v/%v", path_to_repr, file.Name())

			err := dirTreeHelp(output, next_path, tab_to_pass, show_files)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil
}

func dirTree (output *bytes.Buffer, path_to_repr string, show_files bool) error {

	path_to_repr, err := filepath.Abs(path_to_repr)
	if err != nil {
		log.Fatal(err)
	}

	return dirTreeHelp(output, path_to_repr, "", show_files)
}

/*
	This is an implementation of tree Unix (command), which represents displays of a directory in a tree like format.
*/

func main() {
	out := new(bytes.Buffer)
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}

	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(out.String())

}