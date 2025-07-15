# Submore

**`submore`** is a command-line tool that brings order to your domain lists. It takes a list of subdomains, automatically discovers all their parent domains, and sorts the entire collection hierarchically. Use it to quickly organize complex domain lists or to scope your results to a specific target.

## ❯ Features ✨

* **Hierarchical Sorting:** Sorts domains by their level (e.g., `target.com` before `test.target.com`).
* **Parent Discovery:** Automatically finds and includes parent domains (e.g., finds `target.com` from `test.target.com`).
* **Scope Filtering:** Easily filter the list to show only subdomains of a specific domain.
* **Stdin and File Input:** Accepts input directly from a file (`-f`) or from other tools via `stdin`.

---

## ❯ Installation

You can install `submore` in two ways.

**1. Using `go install` (Recommended)**
```bash
go install -v github.com/ir4gh4v/submore@latest
```

2. Manual Build from Source
```bash
# Clone the repository
git clone https://github.com/ir4gh4v/submore.git
cd submore

# Build the binary
go build .

# Move the binary to your PATH (optional)
sudo cp submore /usr/local/bin/
````

## ❯ Usage
```bash
# Basic usage from a file
submore -f <file>

# Filter results for a specific domain
submore -f <file> -d <domain-scope>

# Use with stdin from another tool
cat <file> | submore -d <domain-scope>

Flags:

-f: Specifies the input file containing a list of domains.

-d: Specifies the domain to scope the results to.
```

## ❯ Examples
Let's use the following domains.txt file for all scenarios:
```bash
$ cat domains.txt
api.test.target.com
www.target.com
assets.other.com
test.target.com
```

1. Basic Sorting and Discovery
This example shows how submore automatically finds all parent domains (other.com, target.com) from the input list and then sorts the combined list hierarchically.

```bash
$ submore -f domains.txt
other.com
assets.other.com
target.com
test.target.com
api.test.target.com
www.target.com
```

2. Scoping to a Root Domain
Here, we only want domains that belong to target.com. The -d flag filters the results accordingly.

```bash
$ submore -f domains.txt -d target.com
target.com
test.target.com
api.test.target.com
www.target.com
```

3. Scoping to a Specific Subdomain
You can also scope the output to a deeper subdomain. The tool will show the specified subdomain and any of its children.

```bash
$ submore -f domains.txt -d test.target.com
test.target.com
api.test.target.com
```
