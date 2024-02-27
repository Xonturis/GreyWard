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

package unfichier

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

type Vendor1Fichier struct {
	Link string
}

type responseDDL1Fichier struct {
	Url     string `json:"url"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type responseInfo1Fichier struct {
	Filename string `json:"filename"`
}

type request1Fichier struct {
	Url string `json:"url"`
}

func (v Vendor1Fichier) GetFileDDLLink() (string, error) {
	requestBody, err := json.Marshal(request1Fichier{Url: v.Link})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.1fichier.com/v1/download/get_token.cgi",
		bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+viper.GetString("api-key.1fichier"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var response responseDDL1Fichier

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	if response.Status == "KO" {
		return "", errors.New("1fichier error: " + response.Message)
	}

	return response.Url, nil
}

func (v Vendor1Fichier) GetFileName() (string, error) {
	requestBody, err := json.Marshal(request1Fichier{Url: v.Link})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.1fichier.com/v1/file/info.cgi",
		bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+viper.GetString("api-key.1fichier"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return "", errors.New("file not found")
	}

	body, _ := io.ReadAll(resp.Body)

	var response responseInfo1Fichier

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	return response.Filename, nil
}
