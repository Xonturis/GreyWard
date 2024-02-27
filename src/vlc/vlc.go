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

package vlc

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/xonturis/greyward/src/vendors"
)

func startVLC(link string) {
	fmt.Println("Starting VLC...")
	command := exec.Command(VLC_COMMAND, append(VLC_COMMAND_ARGS, link)...)
	err := command.Run()

	if err != nil {
		fmt.Printf("An error occured: %s\n", err.Error())
	}
}

func Start(link string) error {
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

	fmt.Printf("Starting: %s\nLink: %s\n", name, link)

	go startVLC(link)

	return nil
}
