package documents

import "github.com/MoonMoon1919/doyoucompute"

func ReadMe() (doyoucompute.Document, error) {
	document, err := doyoucompute.NewDocument("MAYI-CLI")
	if err != nil {
		return doyoucompute.Document{}, err
	}

	document.WriteIntro().
		Text("A powerful CLI tool for managing codeowners files with intelligent conflict detection, rule analysis, and automatic optimization.").
		Text("Powered by").
		Link("mayi.", "https://github.com/MoonMoon1919/mayi")

	// Features
	featuresSection := document.CreateSection("Features")
	featureList := featuresSection.CreateList(doyoucompute.BULLET)
	featureList.Append("🖊️ Add, update, and remove rules")
	featureList.Append("✅ Validate rules and accuracy")
	featureList.Append("🔎 Check who can review files & directories")
	featureList.Append("📄 Generate files")

	// Installation
	quickStartSection := document.CreateSection("Quick Start")
	installation := quickStartSection.CreateSection("Installation")
	installation.WriteCodeBlock(
		"sh",
		[]string{"go install github.com/MoonMoon1919/mayi-cli@latest"},
		doyoucompute.Exec,
	)

	// TODO: Add commands to quickstart!
	basicCommands := quickStartSection.CreateSection("Basic commands")
	basicCommands.WriteCodeBlock("sh", []string{`mayi-cli create

# Add a rule
mayi-cli add rule --pattern 'docs/*' --owner '@MoonMoon1919' --owner '@toastsandwich123'

# Search for owners
mayi-cli get owners --pattern docs/

# Add another owner
mayi-cli add owner --pattern 'docs/*' --owner '@example'

# Analyze the file
mayi-cli analyze`}, doyoucompute.Exec)

	// Usage
	usageSection := document.CreateSection("Usage")
	creation := usageSection.CreateSection("Creating files")
	creation.WriteCodeBlock("sh", []string{`# Defaults to .github/CODEOWNERS
mayi-cli create

# Or override to your desired path
mayi-cli create --path CODEOWNERS`}, doyoucompute.Exec)

	adding := usageSection.CreateSection("Adding rules")
	adding.WriteCodeBlock("sh", []string{`# Add a basic rule to have one owner for docs
mayi-cli add rule --pattern 'docs/*' --owner '@MoonMoon1919'

# Except for samples - require no owner for those
mayi-cli add rule --pattern 'docs/internal/samples/*' --owner '' --action exclude`}, doyoucompute.Exec)

	addingOwners := usageSection.CreateSection("Adding rule owners")
	addingOwners.WriteCodeBlock("sh", []string{"mayi-cli add owner --pattern 'docs/*' --owner @example"}, doyoucompute.Exec)

	removing := usageSection.CreateSection("Removing rules")
	removing.WriteCodeBlock("sh", []string{"mayi-cli delete rule --pattern 'docs/internal/samples/*'"}, doyoucompute.Exec)

	removingOwners := usageSection.CreateSection("Removing rule owners")
	removingOwners.WriteCodeBlock("sh", []string{"mayi-cli delete owner --pattern 'docs/*' --owner @example"}, doyoucompute.Exec)

	getting := usageSection.CreateSection("Searching for rules")
	getting.WriteCodeBlock("sh", []string{"mayi-cli get owners --pattern docs/"}, doyoucompute.Exec)

	moving := usageSection.CreateSection("Moving rules")
	moving.WriteCodeBlock("sh", []string{`# Add a random rule
mayi-cli add rule --pattern '*.md' --owner @MoonMoon1919

# Then move it
mayi-cli move --source-pattern '*.md' --destination-pattern 'docs/*' --direction before`}, doyoucompute.Exec)

	analyze := usageSection.CreateSection("Analyzing")
	analyze.WriteCodeBlock("sh", []string{`# Analyze without fixing
mayi-cli analyze

# Or fix any errors the analyzer encounters
mayi-cli analyze --fix
	`}, doyoucompute.Exec)

	// Contrib
	contributing := document.CreateSection("Contributing")
	contributing.WriteIntro().
		Text("See").
		Link("CONTRIBUTING", "./CONTRIBUTING.md").
		Text("for details.")

	// License
	licenseSection := document.CreateSection("License")
	licenseSection.WriteIntro().
		Text("MIT License - see").
		Link("LICENSE", "./LICENSE").
		Text("for details.")

	// Disclaimer
	disclaimerSection := document.CreateSection("Disclaimers")
	disclaimerSection.WriteIntro().
		Text("This work does not represent the interests or technologies of any employer, past or present.").
		Text("It is a personal project only.")

	return document, nil
}
