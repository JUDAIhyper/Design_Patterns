package model

import (
	. "Composite/controller"
	"fmt"
)

type Folder struct {
	Componets []Component
	Name      string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, composite := range f.Componets {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c Component) {
	f.Componets = append(f.Componets, c)
}
