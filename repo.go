package alpm

import (
	"log"
	"os/exec"
	"path/filepath"
	"sync"
)

type Repository struct {
	database string
	path     string
	mutex    *sync.Mutex
}

func NewRepository(path string) *Repository {
	r := new(Repository)
	r.mutex = new(sync.Mutex)
	r.database = path
	r.path = filepath.Dir(path)
	return r
}

func (r *Repository) Database() string {
	return r.database
}

func (r *Repository) Directory() string {
	return r.path
}

func (r *Repository) PackagePath(file string) string {
	return filepath.Join(r.path, file)
}

func (r *Repository) Add(pkg string, sign bool, delta bool) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	args := []string{"-R", "-f", "-q"}
	if sign {
		args = append(args, "-s", "-v")
	}
	if delta {
		args = append(args, "-d")
	}
	args = append(args, r.database, pkg)

	cmd := exec.Command("repo-add", args...)

	if output, err := cmd.CombinedOutput(); err != nil {
		log.Println(string(output))
		log.Fatalln(err)
	} else {
		log.Println(string(output))
	}

	return nil
}
