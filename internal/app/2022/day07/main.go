package day07

import (
	"fmt"
	"io/fs"
	"strings"
	"time"

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

func part1Func(preparedInput *DirInfo) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	partLogger.Info().Msgf("%d", preparedInput.GetSize())

	puzzleAnswer = cast.ToUint64(0)
	return puzzleAnswer
}

func part2Func(preparedInput any) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	// Logic here
	puzzleAnswer = cast.ToUint64(0)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) *DirInfo {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	root = &DirInfo{
		Parent: nil,
		Name:   "/",
		Dirs:   map[string]*DirInfo{},
		Files:  map[string]FileInfo{},
	}

	// Initial Command
	ParseCommand(input[0][2:]).Execute(root, input[1:], "")

	return root
}

var (
	root *DirInfo
)

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

func (c Command) Execute(pwd *DirInfo, s []string, dirName string) {
	switch c {
	case CdRoot:
		pwd = root
	case CdParent:
		pwd = pwd.Parent
	case CdDir:
		pwd = pwd.Dirs[dirName]
	case Ls:
		listedDir := []string{}
		for !strings.HasPrefix(s[0], "$") {
			listedDir = append(listedDir, s[0])
			s = s[1:]
		}
		pwd.ExploreDir(listedDir)
	}
	nextCommand := ParseCommand(strings.TrimPrefix(s[0], "$ "))
	if nextCommand == CdDir {
		nextCommand.Execute(pwd, s[1:], strings.TrimPrefix(s[0], "$ cd "))
	} else {
		nextCommand.Execute(pwd, s[1:], "")
	}
}

func CustomFS(in []string) fs.FS {
	return NewCustomFS(in)
}

type customFS struct{}

func NewCustomFS(in []string) customFS {
	return customFS{}
}

var (
	_ fs.FS        = customFS{}
	_ fs.ReadDirFS = customFS{}
)

var (
	_ fs.ReadDirFile = directory{}
	_ fs.DirEntry    = directory{}
)

type directory struct {
	name     string
	children []fs.DirEntry
}

func (d directory) ReadDir(n int) ([]fs.DirEntry, error) {
	return d.children, nil
}

func (d directory) Name() string {
	return d.name
}

func (_ directory) IsDir() bool {
	return true
}

func (_ directory) Type() fs.FileMode {
	return fs.ModeDir
}

func (d directory) Info() (fs.FileInfo, error) {
	return d, nil
}

func (_ directory) ModTime() time.Time {
	return time.Time{}
}

func (_ directory) Mode() fs.FileMode {
	return fs.ModeDir
}

func (_ directory) Sys() any {
	return nil
}

func (_ directory) Size() int64 {
	return 0
}

func (_ directory) Close() error {
	return nil
}

func (_ directory) Read(_ []byte) (int, error) {
	return 0, nil
}

func (d directory) Stat() (fs.FileInfo, error) {
	return d, nil
}

func (dir DirInfo) GetSize() (totalSize uint64) {
	for _, f := range dir.Files {
		totalSize += f.Size()
	}
	for _, d := range dir.Dirs {
		totalSize += d.GetSize()
	}
	return
}

func ParseDir(parent *DirInfo, in string) *DirInfo {
	return &DirInfo{
		Parent: parent,
		Name:   strings.TrimPrefix(in, "dir "),
		Dirs:   map[string]*DirInfo{},
		Files:  map[string]fs.FileInfo{},
	}
}

func (d *DirInfo) ExploreDir(splitDir []string) {
	for _, e := range splitDir {
		if strings.HasPrefix(e, "dir") {
			dir := ParseDir(d, e)
			d.Dirs[dir.Name] = dir
		} else {
			file := ParseFile(e)
			d.Files[file.Name()] = file
		}
	}
}

var (
	_ fs.File     = file{}
	_ fs.DirEntry = file{}
)

type file struct {
	name string
	size int64
}

func ParseFile(in string) file {
	splitFile := strings.Split(in, " ")
	return file{
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

func (f file) Stat() (fs.FileInfo, error) {
	return f, nil
}

func (_ file) IsDir() bool {
	return false
}

func (_ file) ModTime() time.Time {
	return time.Time{}
}

func (_ file) Mode() fs.FileMode {
	return fs.ModePerm
}

func (_ file) Sys() any {
	return nil
}

func (_ file) Close() error {
	return nil
}

func (_ file) Read(_ []byte) (int, error) {
	return 0, nil
}

func (f file) Info() (fs.FileInfo, error) {
	return f, nil
}

func (_ file) Type() fs.FileMode {
	return fs.ModePerm
}
