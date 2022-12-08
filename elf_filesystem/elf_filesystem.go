package elf_filesystem

type File struct {
	Name string
	Size int
}

func NewFile(name string, size int) *File {
	return &File{
		Name: name,
		Size: size,
	}
}

type Dir struct {
	Name       string
	Parent     *Dir
	Children   []*Dir
	Files      []*File
	TotalSize  int  // total amount in directory plus its children
	DFSVisited bool // if we've visited the node while doing DFS
}

func NewDir(name string, parent *Dir) *Dir {
	return &Dir{
		Name:       name,
		Parent:     parent,
		Children:   []*Dir{},
		Files:      []*File{},
		TotalSize:  0,
		DFSVisited: false,
	}
}

func AddSize(d *Dir, size int) {

	d.TotalSize += size
	if d.Parent != nil {
		AddSize(d.Parent, size)
	}

}

func Cd(current *Dir, name string) *Dir {
	for _, dir := range current.Children {
		if dir.Name == name {
			return dir
		}
	}
	panic(name + " doesn't exist in " + current.Name)
}

func Up(current *Dir) *Dir {
	return current.Parent
}

func GetRoot(current *Dir) *Dir {
	if current.Parent == nil {
		return current
	}
	return GetRoot(current.Parent)
}
