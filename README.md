# Pull Request Summary
Pull Request Summary is a simple script that retrieves the latest pull request activity for a given Github repository and prints out the results. This script can be extended to support multiple communication channels such as email, slack, webhooks, etc.

## Usage
```./pullreqsum --help```
```
Fetch the Github pull requests created for a given repository and email the results to one or more recipients

Usage:
  pullreqsum [flags]

Flags:
      --config string       config file (default is $HOME/.pullreqsum.yaml)
  -h, --help                help for pullreqsum
      --recipients string   List of email addresses to recieve the summary report
      --repository string   Github.com repository {{owner}}/{{repo}} (Example: golang/go)
```