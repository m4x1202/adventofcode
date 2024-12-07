package main

import (
	"github.com/m4x1202/adventofcode/internal/app/cmd"
	_ "github.com/m4x1202/adventofcode/internal/app/cmd/2021"
	_ "github.com/m4x1202/adventofcode/internal/app/cmd/2022"
	_ "github.com/m4x1202/adventofcode/internal/app/cmd/2023"
	_ "github.com/m4x1202/adventofcode/internal/app/cmd/2024"
)

func main() {
	cmd.Execute()
}
