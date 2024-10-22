package csharp

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func downloadDotnetCoverage(ctx context.Context) error {
	// Consruct arguments for command to check if dotnet-coverage is already installed or not.
	check_args := []string{
		"dotnet",
		"tool",
		"list",
		"-g",
	}

	checkCmd := exec.CommandContext(ctx, check_args[0], check_args[1:]...)
	checkCmd.Stdout = os.Stdout
	checkCmd.Stderr = os.Stderr

	if err := checkCmd.Run(); err != nil {
		return fmt.Errorf("failed to check for existing dotnet-coverage: %w", err)
	} else {
		if strings.Contains(checkCmd.String(), "dotnet-coverage") {
			fmt.Println("dotnet-coverage is already installed")
			return nil
		}
	}

	fmt.Println("dotnet-coverage not found. Installing...")

	// Construct the command arguments to install dotnet-coverage
	installArgs := []string{
		"dotnet",
		"tool",
		"install",
		"--global",
		"dotnet-coverage",
	}

	installCmd := exec.CommandContext(ctx, installArgs[0], installArgs[1:]...)
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr

	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("failed to install dotnet-coverage. Ensure .NET SDK is installed and try again: %w", err)
	}

	return nil
}