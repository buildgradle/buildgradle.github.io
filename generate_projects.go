package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type The struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	Files       []File `json:"files"`
}

type File struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	projects := &The{
		make([]Project, 0),
	}
	addProjects(projects, "files")
	b, err := json.MarshalIndent(projects, "", "  ")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(b)
	ioutil.WriteFile("projects.json", b, 0644)
}

func addProjects(projects *The, dir string) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	projectRoot := false
	for _, info := range infos {
		if !info.IsDir() {
			projectRoot = true
		}
	}
	if projectRoot {
		project := Project{
			Name:  dir[len("files/"):],
			Files: make([]File, 0),
		}
		b, err := ioutil.ReadFile(dir + "/project.json")
		if err == nil {
			tmp := Project{}
			json.Unmarshal(b, &tmp)
			if tmp.Name != "" {
				project.Name = tmp.Name
			}
			project.Description = tmp.Description
			project.URL = tmp.URL
		}
		addFiles(&project, dir, dir)
		projects.Projects = append(projects.Projects, project)
	} else {
		for _, info := range infos {
			addProjects(projects, dir+"/"+info.Name())
		}
	}
}

func addFiles(project *Project, root string, dir string) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, info := range infos {
		if !info.IsDir() && info.Name() != "project.json" {
			path := dir + "/" + info.Name()
			file := File{
				Name: path[len(root+"/"):],
				URL:  "https://buildgradle.github.io/" + path,
			}
			project.Files = append(project.Files, file)
		}
	}
	for _, info := range infos {
		if info.IsDir() {
			addFiles(project, root, dir+"/"+info.Name())
		}
	}
}
