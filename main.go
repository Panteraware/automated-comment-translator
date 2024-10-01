package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var translateLink string
	var translateApiKey string
	var file string
	var fromLanguage string
	var toLanguage string
	var saveLocation string

	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("Version=%s Commit=%s Build Time=%s\n", version, commit, date)
	}

	app := &cli.App{
		Name:    "Automated Comment Translator",
		Usage:   "Automatically translate comments in code files.",
		Version: version,
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Konotorii",
				Email: "github@konotorii.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "translateLink",
				Usage:       "Link to translation server",
				Destination: &translateLink,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "translateApiKey",
				Usage:       "API Key to translation server",
				Destination: &translateApiKey,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "file", // Not actual file, just file contents
				Usage:       "Path to file to be translated",
				Destination: &file,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "fromLanguage",
				Value:       "de",
				Usage:       "Language to translate from",
				Destination: &fromLanguage,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "toLanguage",
				Value:       "en",
				Usage:       "Language to translate to",
				Destination: &toLanguage,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "saveLocation",
				Usage:       "File location to save the final output at",
				Destination: &saveLocation,
				Required:    true,
			},
		},
		Action: func(c *cli.Context) error {
			replaced := file

			total := CountLines(file)

			println(fmt.Sprintf("Found %s comments to translate...", FormatNumber(total)))

			f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
			if err != nil {
				log.Fatalf("open file error: %v", err)
				return nil
			}
			defer f.Close()

			index := 0

			startTime := time.Now()

			sc := bufio.NewScanner(f)
			for sc.Scan() {
				text := sc.Text()

				finalLine := text

				characters := strings.Split(text, "")

				check := CheckingArray(characters)

				if check.Matches {
					findIndex := slices.Index(characters, ";")

					reconstruct := strings.Join(slices.Delete(characters, findIndex, len(characters)-1), "")

					spacesBefore := strings.Join(slices.Delete(characters, 0, findIndex), "")

					translated := Translate(c, translateLink, translateApiKey, Format(reconstruct), fromLanguage, toLanguage)

					finalLine = finalLine + "\n" + spacesBefore + *check.Value + translated

					index++

				}
			}
			if err := sc.Err(); err != nil {
				log.Fatalf("scan file error: %v", err)
				return nil
			}

			println(fmt.Sprintf("Translated %d comment lines in %s.", index, time.Since(startTime).String()))

			err = os.WriteFile(saveLocation, []byte(replaced), 0644)
			if err != nil {
				log.Fatal(err)
				return err
			}

			println(fmt.Sprintf("Saved translation to: %s", saveLocation))

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
