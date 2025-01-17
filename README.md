![License](https://img.shields.io/badge/license-sushiware-red)
![Issues open](https://img.shields.io/github/issues/crashbrz/ghorgs)
![GitHub pull requests](https://img.shields.io/github/issues-pr-raw/crashbrz/ghorgs)
![GitHub closed issues](https://img.shields.io/github/issues-closed-raw/crashbrz/ghorgs)
![GitHub last commit](https://img.shields.io/github/last-commit/crashbrz/ghorgs)

# ghorgs - GitHub [Enterprise] Organization Fetcher

This script dumps and displays information about GitHub organizations using the GitHub [Enterprise] API. It supports features like minimal output, result limits, and saving the output to a file.

## Usage

Run the script using the `go run` command. The following flags are available:

### Flags

- **`-t`**: *(Required)* Your GitHub Personal Access Token.
- **`-u`**: *(Optional)* The base URL of the GitHub API. Defaults to `https://api.github.com`.
- **`-O`**: *(Required)* Enable the script to list GitHub Enterprise organizations.
- **`-m`**: *(Optional)* Output minimal information (organization login only).
- **`-o`**: *(Optional)* Specify a file to save the output.
- **`-l`**: *(Optional)* Limit the number of results fetched. If omitted or set to `0`, all organizations will be fetched.

### Examples

#### Fetch all organizations
```
go run ghorgs-list.go -t YOUR_GITHUB_TOKEN -O
```

#### Fetch a maximum of 10 organizations
```
go run ghorgs-list.go -t YOUR_GITHUB_TOKEN -O -l 10
```

#### Fetch organizations with minimal output and save to a file
```
go run ghorgs-list.go -t YOUR_GITHUB_TOKEN -O -m -o output.txt
```

#### Fetch organizations from a custom GitHub Enterprise instance
```
go run ghorgs-list.go -t YOUR_GITHUB_TOKEN -u https://your-github-instance.com -O
```

## Output

The script prints results to the console. If the `-o` flag is set, it also writes the results to the specified file.

### Full Output Example
```
==== Listing Organizations ====
- org1 (ID: 12345, URL: https://api.github.com/orgs/org1)
- org2 (ID: 67890, URL: https://api.github.com/orgs/org2)
```

### Minimal Output Example (`-m`)
```
org1
org2
```
