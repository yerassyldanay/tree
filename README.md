# Tree Unix (command)  in golang

This is an implementation of tree (command) Unix, which displays contents of a directory in a tree-like format.

### Input #1
```
go run main.go ./testdata -f
```

### Output #1

```
├───project
│	├───file.txt (19b)
│	└───gopher.png (70372b)
├───static
│	├───a_lorem
│	│	├───dolor.txt (empty)
│	│	├───gopher.png (70372b)
│	│	└───ipsum
│	│	 	└───gopher.png (70372b)
│	├───css
│	│	└───body.css (28b)
│	├───empty.txt (empty)
│	├───html
│	│	└───index.html (57b)
│	├───js
│	│	└───site.js (10b)
│	└───z_lorem
│	 	├───dolor.txt (empty)
│	 	├───gopher.png (70372b)
│	 	└───ipsum
│	 	 	└───gopher.png (70372b)
├───zline
│	├───empty.txt (empty)
│	└───lorem
│	 	├───dolor.txt (empty)
│	 	├───gopher.png (70372b)
│	 	└───ipsum
│	 	 	└───gopher.png (70372b)
└───zzfile.txt (empty)

```

### Input #2
```
go run main.go ./testdata
```

### Output #2
```
├───project
├───static
│	├───a_lorem
│	│	└───ipsum
│	├───css
│	├───html
│	├───js
│	└───z_lorem
│	 	└───ipsum
└───zline
 	└───lorem
 	 	└───ipsum
```
