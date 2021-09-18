package golp

import "path/filepath"

func Template(p string, f string)  (string){
	f, _ = filepath.Abs(p + f)
	return f
}