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

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
	"github.com/xonturis/greyward/src/license"
	"github.com/xonturis/greyward/src/server"
)

func main() {
	fmt.Println(`GreyWard  Copyright (C) 2024  Lylian Siffre
This program comes with ABSOLUTELY NO WARRANTY; for details type 'warranty'.
This is free software, and you are welcome to redistribute it
under certain conditions; type 'license' for details.`)

	viper.SetConfigFile("./config.yml")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("Error while reading config: %s\n", err.Error())
		os.Exit(1)
	}

	go server.StartGreywardServer()

	go license.ListenForCommand()

	// Create a channel to receive OS signals
	signalChannel := make(chan os.Signal, 1)

	// Notify the signal channel for SIGINT and SIGTERM
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	<-signalChannel

	fmt.Println("Signal received. Exiting...")

	os.Exit(0)
}
