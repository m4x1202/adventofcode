package cmd2021

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
	Year       = 2021
)

func init() {
	cmd.RootCmd.AddCommand(cmd2021)

	cmd2021.AddCommand(downloadInputCmd)
}

var (
	cmd2021 = &cobra.Command{
		Use:   "2021",
		Short: "2021 puzzles",
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
				Value: "53616c7465645f5f3515af344771bb615a081235d7b00a94e2da9b0d2c978120cc836b72cff58d294854200f244a78165ecc92c9b58a47fc5a993c7cb2b2a5db", // Needs to be set whenever we want to download input files
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
