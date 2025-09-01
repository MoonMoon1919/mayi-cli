package documents

import "github.com/MoonMoon1919/doyoucompute"

func ReadMe() (doyoucompute.Document, error) {
	document, err := doyoucompute.NewDocument("MAYI")
	if err != nil {
		return doyoucompute.Document{}, err
	}

	document.WriteIntro().
		Text("A powerful CLI tool for managing codeowners files with intelligent conflict detection, rule analysis, and automatic optimization.").
		Text("Powered by").
		Link("mayi", "https://github.com/MoonMoon1919/mayi")

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
