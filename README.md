# Project Title

Posts discussion prompts from a Google Sheet to a discord channel. Intended to be run as a cron job

## Getting Started

### Dependencies

* Go
* [SheetDB](https://sheetdb.io)
* A [Google Sheet](https://sheets.google.com/) with the columns `id`, `prompt` and `is_posted`

### Usage
#### Configuration:
1. Copy the `.env.example` file into a `.env` file
2. Create a Webhook in Discord and copy the URL into your `.env` file
3. Create a Google Sheet with the required columns
4. Paste the Google Sheet into SheetDB
5. Copy the SheetDB API Endpoint for your Sheet into the `.env` file
6. Set the correct ID for your Discussion Ping role

#### Compilation:
```bash
$ go build
```

#### Running:
```bash
$ discussion-prompts post
```

## License

This project is licensed under the MIT License 

MIT License

Copyright (c) 2021 Ishan Das Sharma

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.


