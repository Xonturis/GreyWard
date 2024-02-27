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

package vendors

import (
	"errors"
	"net/url"

	unfichier "github.com/xonturis/greyward/src/vendors/1fichier"
)

func GetVendor(link string) (Vendor, error) {
	url, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	host := url.Host
	url.Query().Del("af")

	switch host {
	case "1fichier.com":
		return unfichier.Vendor1Fichier{Link: url.String()}, nil
	default:
		return nil, errors.New("couldn't find corresponding vendor")
	}

}
