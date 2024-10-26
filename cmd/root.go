package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/HaroldObasi/copy-cat/bootstrap"
	"github.com/HaroldObasi/copy-cat/deployments"
	sourcecontrol "github.com/HaroldObasi/copy-cat/source-control"
	"github.com/HaroldObasi/copy-cat/template"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "copy-cat",
	Short: "copy cat is a very fast static site generator",
	Long: `Generate and deploy static site portfolio
	
	Example: copy-cat -a my-app -g YOUR_GITHUB_TOKEN -v YOUR_VERCEL_TOKEN -u user-info-path.json`,
	Run: RunCommand,
	PreRunE: PreRunChecks,
}

func init() {
	rootCmd.Flags().StringP("appName", "a", "", "Name of the app")
	rootCmd.Flags().StringP("githubToken", "g", "", "Github token")
	rootCmd.Flags().StringP("vercelToken", "v", "", "Vercel token")
	rootCmd.Flags().StringP("userInfoPath", "u", "", "Path to user info")
}

func RunCommand(cmd *cobra.Command, args []string){
	appName, _ := cmd.Flags().GetString("appName")
	githubToken, _ := cmd.Flags().GetString("githubToken")
	vercelToken, _ := cmd.Flags().GetString("vercelToken")
	userInfoPath, _ := cmd.Flags().GetString("userInfoPath")

	absPath, err := filepath.Abs(userInfoPath)

	if err != nil {
		fmt.Println("User info path is invalid")
		os.Exit(1)
	}

	userInfoPath = absPath

	bootstrap.BootStrapApp(appName)
	template.FormatTemplate(appName, userInfoPath)
	apiUrl, htmlUrl := sourcecontrol.CreateRepo(appName, githubToken)
	sourcecontrol.UploadDir(appName, apiUrl, githubToken)
	deployments.DeployToVercel(htmlUrl, vercelToken, appName)

	fmt.Println("App name: ", appName)
	fmt.Println("Github token: ", githubToken)
	fmt.Println("Vercel token: ", vercelToken)
	fmt.Println("User Info path: ", userInfoPath)

}

func PreRunChecks(cmd *cobra.Command, args []string) error {
	appName, _ := cmd.Flags().GetString("appName")
	githubToken, _ := cmd.Flags().GetString("githubToken")
	vercelToken, _ := cmd.Flags().GetString("vercelToken")
	userInfoPath, _ := cmd.Flags().GetString("userInfoPath")

	if appName == "" {
		return fmt.Errorf("app name is required")
	}

	if githubToken == "" {
		return fmt.Errorf("github token is required")
	}

	if vercelToken == "" {
		return fmt.Errorf("vercel token is required")
	}

	if userInfoPath == "" {
		return fmt.Errorf("user info path is required")
	}

	_, err := filepath.Abs(userInfoPath)
	if err != nil {
		return fmt.Errorf("user info path is invalid")
	}

	_, err = os.Stat(userInfoPath)

	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("user info path does not exist")
		}
		return fmt.Errorf("error accessing file: %w", err)
	}

	return nil
}
  
func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
}