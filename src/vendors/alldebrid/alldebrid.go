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

package alldebrid

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

type VendorAlldebrid struct {
	Link string
}

type fileData struct {
	Link       string   `json:"link"`
	Filename   string   `json:"filename"`
	Host       string   `json:"host"`
	Streams    []string `json:"streams"`
	Paws       bool     `json:"paws"`
	Filesize   int      `json:"filesize"`
	Id         string   `json:"id"`
	HostDomain string   `json:"hostdomain"`
	Delayed    int      `json:"delayed"`
}

type responseDDLAlldebrid struct {
	Status string   `json:"status"`
	Data   fileData `json:"data"`
}

func getAlldebridData(v VendorAlldebrid) (*fileData, error) {
	resp, err := http.Get(
		fmt.Sprintf("https://api.alldebrid.com/v4/link/unlock?agent=greyward&apikey=%s&link=%s",
			viper.GetString("api-key.alldebrid"),
			v.Link))
	if err != nil {
		fmt.Println("Error " + err.Error())
		return nil, err
	}

	body, _ := io.ReadAll(resp.Body)

	var response responseDDLAlldebrid

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error " + err.Error())
		return nil, err
	}

	if response.Status != "success" {
		fmt.Println("Error " + response.Status)
		return nil, errors.New("alldebrid error: " + response.Status)
	}

	return &response.Data, nil
}

func (v VendorAlldebrid) GetFileDDLLink() (string, error) {
	response, err := getAlldebridData(v)

	return response.Link, err
}

func (v VendorAlldebrid) GetFileName() (string, error) {
	response, err := getAlldebridData(v)

	return response.Filename, err
}
