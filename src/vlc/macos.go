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

//go:build darwin
// +build darwin

package vlc

var VLC_COMMAND = "open"
var VLC_COMMAND_ARGS = []string{"-a", "VLC", "--args", "--play-and-exit"}