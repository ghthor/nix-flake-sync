LICENSE:
	nix-shell -p license-generator --run "license-generator --author 'Will Drengwitz <ghthor@gmail.com>' --year 2024 MIT"
	copywrite license
