# Foolder

A lightweight, fast web directory brute forcer built in Go. Built for security testing and CTF usage.

> Disclaimer: This is a personal project built for security testing purposes only. Foolder is designed to help identify web directories that should not be publicly accessible to unauthorized users.

---

## Features

- Concurrent scanning with a configurable worker pool
- Detects `200`, `301`, and `403` status codes
- Color-coded output
- Custom wordlist support
- Configurable timeout and worker count

---

## Install

```bash
git clone https://github.com/kajrab/foolder
cd foolder
go mod tidy
go build -o foolder
```

---

## Usage

```bash
./foolder --url http://target.com
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--url` | required | Target URL |
| `--wordlist` | `wordlists/default.txt` | Path to wordlist |
| `--workers` | `50` | Number of concurrent workers |
| `--timeout` | `3` | HTTP timeout in seconds |
| `--output` | none | Save results to file |

### Examples

```bash
# Basic scan
./foolder --url http://target.com

# Custom wordlist and worker count
./foolder --url http://target.com --wordlist /path/to/wordlist.txt --workers 100

# Save results
./foolder --url http://target.com --output results.txt

# Full options
./foolder --url http://target.com --workers 100 --timeout 5 --output results.txt
```

---

## Output

```
[200] Access Granted:  http://target.com/admin
[301] Redirected:      http://target.com/dashboard
[403] Access Denied:   http://target.com/config
```

---

## Project Structure

```
foolder/
  main.go          # Entry point, flags, setup
  scanner.go       # Worker logic, HTTP requests
  output.go        # Banner, color output
  wordlists/
    default.txt    # Default wordlist
  go.mod
  go.sum
  README.md
  LICENSE
  .gitignore
```

---

## Wordlist

Default wordlist sourced from [SecLists](https://github.com/danielmiessler/SecLists) by Daniel Miessler.

For deeper scans use `directory-list-2.3-medium.txt` from the same repo.

---

## Legal

Only use Foolder against targets you own or have explicit permission to test. Unauthorized scanning is illegal. The author is not responsible for misuse.

---

## License

MIT
