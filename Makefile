LICENSE:
	nix-shell -p license-generator --run "license-generator --author 'Will Drengwitz <ghthor@gmail.com>' --year 2024 MIT"
	copywrite license

gomod2nix.toml: go.mod go.sum
	nix run '.#gomod2nix'
