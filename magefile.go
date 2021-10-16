//go:build mage
// +build mage

package main

import (
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/phillipsj/mage-helpers/dl"
)

var Default = Build

func Clean() error {
	return sh.Rm("bin")
}

func Build() error {
	mg.Deps(Clean)
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	return sh.Run("go", "build", "-o", "./bin/")
}

func Download() error {
	mg.Deps(Build)

	url := "https://gist.githubusercontent.com/phillipsj/07fed8ce06f932c19ab7613d8426d922/raw/13d3fc0ca54d136ad5744fd4448b65dbc87f32dc/random.txt"
	return dl.GetFile(url, filepath.Join("bin", "random.txt"))
}
