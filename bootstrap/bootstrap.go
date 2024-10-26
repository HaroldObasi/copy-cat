package bootstrap

import (
	"fmt"
	"os"
	"os/exec"
)

type PackageJSON struct {
	Name            string            `json:"name"`
	Private         bool              `json:"private"`
	Version         string            `json:"version"`
	Type            string            `json:"type"`
	Scripts         map[string]string `json:"scripts"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func (pj PackageJSON) String() string {
	return fmt.Sprintf("Name: %v", pj.Name)
}

func BootStrapApp(appName string) {

	fmt.Println("Bootstrapping App with vite.")

	// create vite app
	viteAppScaffold := exec.Command(
		"npm",
		"create",
		"vite@latest",
		appName,
		"--",
		"--template",
		"vanilla",
	)

	_, err := viteAppScaffold.Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Installing dependencies.")
	// install tailwind, postcss, autoprefixer
	npmInstallTailwind := exec.Command(
		"npm",
		"install",
		"-D",
		"tailwindcss",
		"postcss",
		"autoprefixer",
	)

	npmInstallTailwind.Dir = appName

	_, err = npmInstallTailwind.Output()
	if err != nil {
		fmt.Println(err)
	}

	//install lucide
	npmInstallLucide := exec.Command(
		"npm",
		"install",
		"lucide",
	)

	npmInstallLucide.Dir = appName

	_, err = npmInstallLucide.Output()
	if err != nil {
		fmt.Println(err)
	}

	// write into tailwind.config.js
	os.WriteFile(fmt.Sprintf("./%s/tailwind.config.js", appName), []byte(TAILWIND_CONTENT), 0644)

	// write into postcss.config.js
	os.WriteFile(fmt.Sprintf("./%s/postcss.config.js", appName), []byte(POST_CSS_CONTENT), 0644)

	// write into style.css
	os.WriteFile(fmt.Sprintf("./%s/style.css", appName), []byte(STYLE_CSS_CONTENT), 0644)

	// write into main.js
	os.WriteFile(fmt.Sprintf("./%s/main.js", appName), []byte(MAIN_JS_CONTENT), 0644)

	fmt.Println("App bootstrapped successfully.")

}
