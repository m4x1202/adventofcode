package day07

import (
	"fmt"
	"strings"

	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

const (
	DAY = "07"
)

var (
	dayLogger = log.With().
			Str("day", DAY).
			Logger()
	partLogger zerolog.Logger
)

func ExecutePart(p uint8) {
	preparedInput := prepareInput(readPuzzleInput())
	switch p {
	case 1:
		part1Func(preparedInput)
	case 2:
		part2Func(preparedInput)
	default:
		panic("part does not exist")
	}
}

func part1Func(preparedInput FileInfo) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	totalSize, _ := preparedInput.(*directory).SizeU100k()

	fmt.Printf("sum of dir sizes: %d\n", totalSize)
	puzzleAnswer = cast.ToUint64(totalSize)
	return puzzleAnswer
}

func part2Func(preparedInput any) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	justRightSize, _ := preparedInput.(*directory).SizeJustRight()

	fmt.Printf("just right size dir: %d\n", justRightSize)
	puzzleAnswer = cast.ToUint64(justRightSize)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) FileInfo {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	root = &directory{
		parent:   nil,
		name:     "/",
		children: []FileInfo{},
	}

	// Initial Command
	ParseCommand(input[0][2:]).Execute(root, input[1:], "")

	return root
}

const (
	fsSpace int64 = 70000000
)

func RequiredSpaceForUpdate() int64 {
	return 30000000 - (fsSpace - root.Size())
}

var root FileInfo

type Command uint8

const (
	_ Command = iota
	CdDir
	CdParent
	CdRoot
	Ls
)

func ParseCommand(in string) Command {
	splitCommand := strings.Split(in, " ")
	switch splitCommand[0] {
	case "ls":
		return Ls
	case "cd":
		switch splitCommand[1] {
		case "/":
			return CdRoot
		case "..":
			return CdParent
		default:
			return CdDir
		}
	default:
		return 0
	}
}

func (c Command) Execute(pwd FileInfo, s []string, dirName string) {
	switch c {
	case CdRoot:
		pwd = root
	case CdParent:
		pwd = pwd.(*directory).parent
	case CdDir:
		for _, c := range pwd.(*directory).children {
			if c.Name() == dirName {
				pwd = c
			}
		}
	case Ls:
		listedDir := []string{}
		for len(s) > 0 && !strings.HasPrefix(s[0], "$") {
			listedDir = append(listedDir, s[0])
			s = s[1:]
		}
		dir := pwd.(*directory)
		dir.ExploreDir(listedDir)
	}
	if len(s) == 0 {
		return
	}
	nextCommand := ParseCommand(strings.TrimPrefix(s[0], "$ "))
	if nextCommand == CdDir {
		nextCommand.Execute(pwd, s[1:], strings.TrimPrefix(s[0], "$ cd "))
	} else {
		nextCommand.Execute(pwd, s[1:], "")
	}
}

type FileInfo interface {
	IsDir() bool
	Name() string
	Size() int64
}

type directory struct {
	parent   FileInfo
	name     string
	children []FileInfo
}

func (d directory) Name() string {
	return d.name
}

func (d directory) SizeU100k() (totalSize int64, dirSize int64) {
	for _, e := range d.children {
		if d, ok := e.(*directory); ok {
			ts, ds := d.SizeU100k()
			totalSize += ts
			dirSize += ds
		} else {
			dirSize += e.Size()
		}
	}
	if dirSize <= 100000 {
		totalSize += dirSize
	}
	return
}

func (d directory) SizeJustRight() (justRightSize int64, dirSize int64) {
	justRightSize = fsSpace
	for _, e := range d.children {
		if d, ok := e.(*directory); ok {
			jrs, ds := d.SizeJustRight()
			if jrs == RequiredSpaceForUpdate() {
				return justRightSize, 0
			} else if jrs < justRightSize {
				justRightSize = jrs
			}
			dirSize += ds
		} else {
			dirSize += e.Size()
		}
	}
	if dirSize > RequiredSpaceForUpdate() && dirSize < justRightSize {
		justRightSize = dirSize
	}
	return
}

func (d directory) Size() (totalSize int64) {
	for _, e := range d.children {
		totalSize += e.Size()
	}
	return
}

func (_ directory) IsDir() bool {
	return true
}

func ParseDir(parent FileInfo, in string) *directory {
	return &directory{
		parent:   parent,
		name:     strings.TrimPrefix(in, "dir "),
		children: []FileInfo{},
	}
}

func (d *directory) ExploreDir(splitDir []string) {
	for _, e := range splitDir {
		if strings.HasPrefix(e, "dir") {
			d.children = append(d.children, ParseDir(d, e))
		} else {
			d.children = append(d.children, ParseFile(e))
		}
	}
}

type file struct {
	name string
	size int64
}

func ParseFile(in string) *file {
	splitFile := strings.Split(in, " ")
	return &file{
		size: cast.ToInt64(splitFile[0]),
		name: splitFile[1],
	}
}

func (f file) Name() string {
	return f.name
}

func (f file) Size() int64 {
	return f.size
}

func (_ file) IsDir() bool {
	return false
}
