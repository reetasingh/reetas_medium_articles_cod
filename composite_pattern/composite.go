package main

import (
	"fmt"
	"os"
)

type IComponent interface {
	Show() error
	Delete() error
}

// File is a single unit
type File struct {
	name string
}

// Directory is a composite unit which consists of multiple Directories or Files as part of components
type Directory struct {
	components []IComponent
	name       string
}

// Create file with the given contents in data
func (f File) Create(data string) error {
	fileHandle, err := os.Create(f.name)
	if err != nil {
		return err
	}
	fileHandle.WriteString(data)
	return nil
}

// Delete file
func (f File) Delete() error {
	fmt.Printf("Deleting file %s\n", f.name)
	err := os.Remove(f.name)
	if err != nil {
		return err
	}
	return nil
}

// Show the contents of file
func (f File) Show() error {
	fmt.Printf("Showing contents of file %s\n", f.name)
	data, err := os.ReadFile(f.name)
	if err != nil {
		return err
	}
	// print the file contents
	fmt.Println(string(data))
	return nil
}

// Create directory
func (d Directory) Create() error {
	d.components = make([]IComponent, 0)
	return nil
}

// Delete directory
func (d Directory) Delete() error {
	fmt.Printf("Deleting directory %s\n", d.name)
	for _, component := range d.components {
		component.Delete()
	}
	return nil
}

// Show contents of a directory
func (d Directory) Show() error {
	fmt.Printf("Showing contents of directory %s\n", d.name)
	for _, component := range d.components {
		component.Show()
	}
	return nil
}

// Add a new component to Directory which is a composite type
func (d *Directory) Add(component IComponent) error {
	d.components = append(d.components, component)
	return nil
}

func main() {

	file_a := File{"a"}
	file_a.Create("this is file a")
	file_a.Show()

	file_b := File{"b"}
	file_b.Create("this is file b")
	file_b.Show()

	dir_m := Directory{name: "x"}
	dir_m.Create()
	dir_n := Directory{name: "y"}
	dir_n.Create()

	dir_m.Add(file_a)
	dir_m.Add(file_b)
	dir_m.Add(dir_n)

	dir_m.Show()
}
