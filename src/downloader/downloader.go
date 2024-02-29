// GreyWard - An HTTP server to unblock links and launch media with VLC
// Copyright (C) 2024  Lylian Siffre

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package downloader

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
	"github.com/xonturis/greyward/src/vendors"
	"golang.org/x/term"
)

// This function copies the data from a io.Reader src to a io.Writer dst
// This function also prints the current progress of the copy depending on the passed contentLength
// argument. The progress is printed in the form of a progressbar.
func copyDataToFile(dst io.Writer, src io.Reader, contentLength int64) error {
	var written int64

	// Allocating copy buffer
	buf := make([]byte, 32*1024)
	loopI := 0
	width, _, err := term.GetSize(0)
	if err != nil {
		width = 100
		err = nil
	} else {
		width -= 7
		if width <= 0 {
			width = 100
		}
	}

	// Print an empty line for the progressbar
	fmt.Println()

	// Re implementation of io.Copy in order to have download progress bar in console
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = errors.New("invalid write result")
				}
			}
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
		// Update progress bar every XXX loop iterations
		if loopI > 2000 {
			// Compute progress in percents, then in "." relative to the console width
			done := (float64(written) / float64(contentLength))
			percent := int(done * float64(width))
			progressBar := strings.Repeat(".", percent)

			// Clear the progress bar line
			// https://stackoverflow.com/questions/75300588/how-to-clear-last-line-in-terminal-with-golang
			fmt.Printf("\033[1A\033[K")

			// Print progress bar
			fmt.Printf(fmt.Sprintf("[%%-%ds] %d%%%%\n", width, percent), progressBar)

			// Reset loopI counter
			loopI = 0
		}
		loopI++
	}

	return err
}

func download(name string, link string) error {
	out, err := os.Create(path.Join(viper.GetString("download.path"), name))
	if err != nil {
		fmt.Printf("An error occured when creating file: %s\n", err.Error())
		return err
	}
	defer out.Close()
	resp, err := http.Get(link)
	if err != nil {
		fmt.Printf("An error occured when downloading file: %s\n", err.Error())
		return err
	}
	defer resp.Body.Close()

	err = copyDataToFile(out, resp.Body, resp.ContentLength)

	if err != nil {
		fmt.Printf("An error occured when writing file: %s\n", err.Error())
		return err
	}

	fmt.Println("Download complete...")

	return nil
}

func DownloadFile(link string) error {
	link = strings.TrimSpace(link)
	vendor, err := vendors.GetVendor(link)
	if err != nil {
		return err
	}

	link, err = vendor.GetFileDDLLink()

	if err != nil {
		return err
	}

	name, err := vendor.GetFileName()

	if err != nil {
		return err
	}

	fmt.Printf("Downloading: %s\nLink: %s\n", name, link)

	go download(name, link)

	return nil
}
