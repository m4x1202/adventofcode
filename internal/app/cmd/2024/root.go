package cmd2024

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"

	"github.com/m4x1202/adventofcode/internal/app/cmd"
	"github.com/spf13/cobra"
)

const (
	AoCBaseURL = "https://adventofcode.com"
	Year       = 2024
)

func init() {
	cmd.RootCmd.AddCommand(cmd2024)

	cmd2024.AddCommand(downloadInputCmd)
}

var (
	cmd2024 = &cobra.Command{
		Use:   "2024",
		Short: "2024 puzzles",
	}
	downloadInputCmd = &cobra.Command{
		Use:       "downloadInput",
		Short:     "Downloads the input for a given day",
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		ValidArgs: []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25"},
		RunE: func(cmd *cobra.Command, args []string) error {
			rawDay := args[0]

			err := os.Mkdir(fmt.Sprintf("resources/%d/day%s", Year, rawDay), 0755)
			if err != nil && !os.IsExist(err) {
				return err
			}

			day := func(in string) uint8 {
				parsed, _ := strconv.ParseInt(in, 10, 0)
				return uint8(parsed)
			}(rawDay)
			inputDownloadUrl, err := url.Parse(fmt.Sprintf("%s/%d/day/%d/input", AoCBaseURL, Year, day))
			if err != nil {
				return err
			}
			jar, err := cookiejar.New(nil)
			if err != nil {
				return err
			}
			currentSession := &http.Cookie{
				Name:  "session",
				Value: "53616c7465645f5fe2cbd8ffc02fdfbdb8b61f9238e950954e5f5fe99c1bba39a0333aa0d462db581cbe882a3aabd04635995cc3d78b97bc7530675e882f2ba3", // Needs to be set whenever we want to download input files
			}
			jar.SetCookies(inputDownloadUrl, []*http.Cookie{currentSession})
			client := &http.Client{
				Jar: jar,
			}

			resp, err := client.Get(inputDownloadUrl.String())
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			inputFile, err := os.Create(fmt.Sprintf("resources/%d/day%s/input.txt", Year, rawDay))
			if err != nil {
				return err
			}
			defer inputFile.Close()

			_, err = io.Copy(inputFile, resp.Body)
			if err != nil {
				return err
			}
			return nil
		},
	}
)
