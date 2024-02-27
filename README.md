Overview
========

This software exposes an HTTP interface for unlocking and starting media through VLC.
At this time, and due to my personal requirements, only `1fichier.com` is handled. If you need GreyWard to handle a new website, feel free to ask through issues, you can also fork or, even better, issue a push request.

Build
=====

To build you must have [Go](https://go.dev/doc/install) installed.
Then simply clone the repo and run `go build`.

At this time, GreyWard works with VLC only, make sure VLC is installed and be found at `C:\Program Files\VideoLAN\VLC\vlc.exe`, otherwise this software won't work.

Configuration
=============

To unlock 1fichier.com files, you need to fill the config file with your API key.
Don't forget to rename `config.yml.example` to `config.yml`.

License
=======

This software is distributed under GPLv3. See `LICENSE` file.