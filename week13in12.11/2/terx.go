package main

func main() {
	dir, err : os.Open(".")
	if err != nil {
		return
	}
	defer dir.close()

	fileIndos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		filename := fi.Name()
		fmt
	}
}