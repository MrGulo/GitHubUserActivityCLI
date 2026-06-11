GitHub User Activity CLI

A lightweight command-line interface (CLI) application built in pure Go. This tool fetches and displays the recent public activity of any GitHub user directly in your terminal.

Built as part of the backend developer roadmap to master REST API integration, JSON parsing, and CLI application development without relying on external dependencies.

🛠️ Tech Stack

Language: Go (Golang 1.26)

Libraries: Pure Go Standard Library (Zero third-party dependencies)

net/http for API requests.

encoding/json for data parsing.

os for command-line arguments.

Key Features

Real-time Data: Fetches the latest events from the official GitHub API (api.github.com).

Event Parsing: Intelligently parses various GitHub event types and translates them into human-readable terminal output (e.g., PushEvent, CreateEvent, WatchEvent).

Graceful Error Handling: Provides clear feedback for invalid usernames (404 Not Found) or network/API failures.

Zero Dependencies: Keeps the binary extremely small and secure by avoiding external frameworks.

Getting Started

Prerequisites

Go installed on your machine.

Installation & Build

Clone the repository:

git clone https://github.com/MrGulo/GitHubUserActivityCLI.git
cd GitHubUserActivityCLI


Build the executable:

go build -o github-activity main.go


💻 Usage

Run the built executable and pass the GitHub username as an argument:

# On Linux/macOS
./github-activity MrGulo

# On Windows
github-activity.exe MrGulo


Example Output

Pushed 5 commits to MrGulo/developer-roadmap
Created branch/repository in MrGulo/Real-Time_Bidding_API
Starred golang/go
Event type ForkEvent in user/some-repo


Developed as a learning project to practice CLI architecture and API communication in Go.
