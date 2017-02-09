# Scarecrow

Scarecrow is a chatbot written in Go. It connects to Slack and XMPP and can be
chatted with in your terminal window, and it will probably be updated to
connect to more things in the future too.

It uses [RiveScript](http://www.rivescript.com/) as its brain back-end, it
remembers information about the people it chats with, keeps log files, etc.

# Features

* [Slack](https://www.slack.com/) integration
  * Users can chat with it over direct message and carry on a conversation
  * It can join public channels where it will sit in silence until a user talks
    directly to it, either by at-mentioning its username or starting a message
    with its name.
* XMPP integration
  * Hangouts bot via the XMPP gateway: use `talk.google.com` port `443` instead
    of the standard XMPP ports (`5222` or `5223`)
  * Known to work with an `ejabberd` server with valid CA certificate (you may
    need to set `"tls-disable": "true"` in the bot's config)
* User roles/permissions
  * Admin users can reload the RiveScript brain without rebooting the entire bot
  * Users across different platforms are uniquely identifiable, so you can add
    admin users without risking a different user with a matching name on a
    different platform having admin rights.
* Goroutines are spawned for each individual bot connection, so you can run
  multiple bots from one instance of the program.
* Chat with the bot on the console, too

# Install and Build

Build it with `make build`

Run it with `./scarecrow-cli`

Command line options are pretty basic: `--debug` for debug mode and
`--version` to get the version number.

# Documentation

[Documentation](docs/README.md) is available as Markdown files in the `docs/`
directory.

# License

```
Scarecrow - A RiveScript Chatbot written in Go
Copyright (C) 2015  Noah Petherbridge

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License along
with this program; if not, write to the Free Software Foundation, Inc.,
51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
```

# See Also

RiveScript's official homepage, <http://www.rivescript.com/>

The RiveScript Go module, <https://github.com/aichaos/rivescript-go>
