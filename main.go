package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	tokenDir      = ".terraform.d/tokens"
	mainTokenFile = ".terraform.d/credentials.tfrc.json"
)

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	cmd := os.Args[1]
	var credsName string
	if len(os.Args) > 2 {
		credsName = os.Args[2]
	}

	switch cmd {
	case "use":
		if err := useCredentials(credsName); err != nil {
			fmt.Println("Error:", err)
		}
	case "create":
		if err := createCredentials(credsName); err != nil {
			fmt.Println("Error:", err)
		}
	case "update":
		if err := updateCredentials(credsName); err != nil {
			fmt.Println("Error:", err)
		}
	case "delete":
		if err := deleteCredentials(credsName); err != nil {
			fmt.Println("Error:", err)
		}
	case "list":
		if err := listCredentials(); err != nil {
			fmt.Println("Error:", err)
		}
	case "-h", "--help":
		showHelp()
	default:
		fmt.Println("Error: Invalid command", cmd)
		showHelp()
	}
}

func useCredentials(name string) error {
	if name == "" {
		return errors.New("no credentials name provided")
	}

	tokenFile := filepath.Join(os.Getenv("HOME"), tokenDir, name+".tfrc.json")
	if _, err := os.Stat(tokenFile); os.IsNotExist(err) {
		return fmt.Errorf("token file for name '%s' not found", name)
	}

	mainTokenFilePath := filepath.Join(os.Getenv("HOME"), mainTokenFile)
	return copyFile(tokenFile, mainTokenFilePath)
}

func createCredentials(name string) error {
	if name == "" {
		return errors.New("no credentials name provided")
	}

	tokenFile := filepath.Join(os.Getenv("HOME"), tokenDir, name+".tfrc.json")
	if err := os.MkdirAll(filepath.Dir(tokenFile), 0700); err != nil {
		return err
	}

	if err := runTerraformLogin(); err != nil {
		return err
	}

	return moveFile(filepath.Join(os.Getenv("HOME"), mainTokenFile), tokenFile)
}

func updateCredentials(name string) error {
	if name == "" {
		return errors.New("no credentials name provided")
	}

	tokenFile := filepath.Join(os.Getenv("HOME"), tokenDir, name+".tfrc.json")
	if _, err := os.Stat(tokenFile); os.IsNotExist(err) {
		return fmt.Errorf("token file for name '%s' not found", name)
	}

	if err := runTerraformLogin(); err != nil {
		return err
	}

	return moveFile(filepath.Join(os.Getenv("HOME"), mainTokenFile), tokenFile)
}

func deleteCredentials(name string) error {
	if name == "" {
		return errors.New("no credentials name provided")
	}

	tokenFile := filepath.Join(os.Getenv("HOME"), tokenDir, name+".tfrc.json")
	if _, err := os.Stat(tokenFile); os.IsNotExist(err) {
		return fmt.Errorf("token file for name '%s' not found", name)
	}

	return os.Remove(tokenFile)
}

func listCredentials() error {
	files, err := filepath.Glob(filepath.Join(os.Getenv("HOME"), tokenDir, "*.tfrc.json"))
	if err != nil {
		return err
	}

	fmt.Printf("%-20s %-20s\n", "NAME", "DATE CREATED")
	fmt.Println(strings.Repeat("-", 40))

	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			return err
		}

		name := strings.TrimSuffix(filepath.Base(file), ".tfrc.json")
		dateCreated := info.ModTime().Format("2006-01-02 15:04:05")
		fmt.Printf("%-20s %-20s\n", name, dateCreated)
	}

	return nil
}

func showHelp() {
	fmt.Println("Usage: tcreds <command> [credentials_name]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  use <credentials_name>       Switch to the specified Terraform credentials.")
	fmt.Println("  create <credentials_name>    Create new Terraform credentials and store them under the provided name.")
	fmt.Println("  update <credentials_name>    Update the existing Terraform credentials for the specified name.")
	fmt.Println("  delete <credentials_name>    Delete the Terraform credentials associated with the specified name.")
	fmt.Println("  list                        List all stored Terraform credentials with their creation date.")
	fmt.Println("  -h, --help                  Display this help menu.")
	fmt.Println()
}

func copyFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	if err := os.WriteFile(dst, input, 0600); err != nil {
		return err
	}

	fmt.Println("Switched Terraform credentials to:", filepath.Base(dst))
	return nil
}

func moveFile(src, dst string) error {
	if err := copyFile(src, dst); err != nil {
		return err
	}

	return os.Remove(src)
}

func runTerraformLogin() error {
	cmd := exec.Command("terraform", "login")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
