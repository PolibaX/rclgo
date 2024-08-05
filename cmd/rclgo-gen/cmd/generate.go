/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
    http://www.apache.org/licenses/LICENSE-2.0
*/

package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/PolibaX/rclgo/pkg/gogen"
	"golang.org/x/tools/go/packages"
)

const correctDistro = "jazzy"

func validateGenerateArgs(cmd *cobra.Command, _ []string) error {
	rootPaths := getRootPaths(cmd)
	if len(rootPaths) == 0 {
		if os.Getenv("AMENT_PREFIX_PATH") == "" {
			return fmt.Errorf("You haven't sourced your ROS2 environment! Cannot autodetect --root-path. Source your ROS2 or pass --root-path")
		}
		return fmt.Errorf("root-path is required")
	}

	distro := os.Getenv("ROS_DISTRO")
	if getBool(cmd, "ignore-ros-distro-mismatch") {
		if distro != correctDistro {
			gogen.PrintErrf("NOTE: Environment variable ROS_DISTRO is set to %q, generating files for %q\n", distro, correctDistro)
		}
	} else if distro != correctDistro {
		return fmt.Errorf("ROS_DISTRO should be set to %q", correctDistro)
	}

	destPath := getString(cmd, "dest-path")
	if destPath == "" {
		return fmt.Errorf("dest-path is required")
	}

	_, err := os.Stat(destPath)
	if errors.Is(err, os.ErrNotExist) {
		//#nosec G301 -- The generated directory doesn't contain secrets.
		err = os.MkdirAll(destPath, 0755)
	}
	if err != nil {
		return fmt.Errorf("dest-path error: %v", err)
	}

	return nil
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Go bindings for ROS2 interface definitions under <root-path>",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := getGogenConfig(cmd)
		if err != nil {
			return err
		}
		gen := gogen.New(config)
		if err := gen.GenerateGolangMessageTypes(); err != nil {
			return fmt.Errorf("failed to generate interface bindings: %w", err)
		}
		if err := gen.GenerateROS2AllMessagesImporter(); err != nil {
			return fmt.Errorf("failed to generate all importer: %w", err)
		}
		if err := gen.GenerateCGOFlags(); err != nil {
			return fmt.Errorf("failed to generate CGO flags: %w", err)
		}
		return nil
	},
	Args: validateGenerateArgs,
}

var generateRclgoCmd = &cobra.Command{
	Use:   "generate-rclgo",
	Short: "Generate Go code that forms a part of rclgo",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := getGogenConfig(cmd)
		if err != nil {
			return err
		}
		gen := gogen.New(config)
		if err := gen.GeneratePrimitives(); err != nil {
			return fmt.Errorf("failed to generate primitive types: %w", err)
		}
		if err := gen.GenerateRclgoFlags(); err != nil {
			return fmt.Errorf("failed to generate rclgo flags: %w", err)
		}
		if err := gen.GenerateTestGogenFlags(); err != nil {
			return fmt.Errorf("failed to generate rclgo flags: %w", err)
		}
		if err := gen.GenerateROS2ErrorTypes(); err != nil {
			return fmt.Errorf("failed to generate error types: %w", err)
		}
		return nil
	},
	Args: validateGenerateArgs,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	configureFlags(generateCmd, ".")

	rootCmd.AddCommand(generateRclgoCmd)
	configureFlags(generateRclgoCmd, gogen.RclgoRepoRootPath())
}

func configureFlags(cmd *cobra.Command, destPathDefault string) {
	cmd.PersistentFlags().StringArrayP("root-path", "r", []string{os.Getenv("AMENT_PREFIX_PATH")}, "Root lookup path for ROS2 .msg files. If ROS2 environment is sourced, is autodetected.")
	cmd.PersistentFlags().StringP("dest-path", "d", destPathDefault, "Destination directory for the Golang typed converted ROS2 messages. ROS2 Message structure is preserved as <ros2-package>/msg/<msg-name>")
	cmd.PersistentFlags().String("rclgo-import-path", gogen.DefaultConfig.RclgoImportPath, "Import path of rclgo library")
	cmd.PersistentFlags().String("message-module-prefix", gogen.DefaultConfig.MessageModulePrefix, "Import path prefix for generated message binding modules")
	cmd.PersistentFlags().StringArray("include-package", nil, "Include only packages matching a regex. Can be passed multiple times. If multiple include options are passed, the union of the matches is generated.")
	cmd.PersistentFlags().StringArray("include-package-deps", nil, "Include only packages which are dependencies of listed packages. Can be passed multiple times. If multiple include options are passed, the union of the matches is generated.")
	cmd.PersistentFlags().StringArray("include-go-package-deps", nil, "Include only packages which are dependencies of listed Go packages. Can be passed multiple times. If multiple include options are passed, the union of the matches is generated.")
	cmd.PersistentFlags().Bool("ignore-ros-distro-mismatch", false, "If true, ignores possible mismatches in sourced and supported ROS distro")
	cmd.PersistentFlags().String("license-header-path", "", "Path to a file containing a license header to be added to generated files. By default no license is added.")
	cmd.PersistentFlags().String("cgo-flags-path", "cgo-flags.env", `Path to file where CGO flags are written. If empty, no flags are written. If "-", flags are written to stdout.`)
	bindPFlags(cmd)
}

func getPrefix(cmd *cobra.Command) string {
	parts := []string{}
	for c := cmd; c != c.Root(); c = c.Parent() {
		parts = append(parts, c.Name())
	}
	for i := 0; i < len(parts)/2; i++ {
		parts[i], parts[len(parts)-i-1] = parts[len(parts)-i-1], parts[i]
	}
	prefix := strings.Join(parts, ".")
	if prefix != "" {
		prefix += "."
	}
	return prefix
}

func getString(cmd *cobra.Command, key string) string {
	return viper.GetString(getPrefix(cmd) + key)
}

func getBool(cmd *cobra.Command, key string) bool {
	return viper.GetBool(getPrefix(cmd) + key)
}

func bindPFlags(cmd *cobra.Command) {
	prefix := getPrefix(cmd)
	cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		viper.BindPFlag(prefix+f.Name, f)
	})
	cmd.LocalFlags().VisitAll(func(f *pflag.Flag) {
		viper.BindPFlag(prefix+f.Name, f)
	})
}

func getGogenConfig(cmd *cobra.Command) (*gogen.Config, error) {
	destPath := getString(cmd, "dest-path")
	modulePrefix := getString(cmd, "message-module-prefix")

	if modulePrefix == gogen.DefaultConfig.MessageModulePrefix {
		pkgs, err := packages.Load(&packages.Config{})
		if err == nil && len(pkgs) > 0 {
			modulePrefix = path.Join(pkgs[0].PkgPath, destPath)
		}
	}
	rules, err := getPackageRules(cmd)
	if err != nil {
		return nil, err
	}
	licenseHeader := getString(cmd, "license-header-path")
	if licenseHeader != "" {
		headerBytes, err := os.ReadFile(licenseHeader)
		if err != nil {
			return nil, err
		}
		licenseHeader = string(headerBytes)
	}
	return &gogen.Config{
		RclgoImportPath:     getString(cmd, "rclgo-import-path"),
		MessageModulePrefix: modulePrefix,
		RootPaths:           getRootPaths(cmd),
		DestPath:            destPath,
		CGOFlagsPath:        getString(cmd, "cgo-flags-path"),

		RegexIncludes:  rules,
		ROSPkgIncludes: viper.GetStringSlice(getPrefix(cmd) + "include-package-deps"),
		GoPkgIncludes:  viper.GetStringSlice(getPrefix(cmd) + "include-go-package-deps"),

		LicenseHeader: licenseHeader,
	}, nil
}

func getRootPaths(cmd *cobra.Command) []string {
	pathLists := viper.GetStringSlice(getPrefix(cmd) + "root-path")
	found := make(map[string]bool)
	var paths []string
	for _, pl := range pathLists {
		for _, p := range filepath.SplitList(pl) {
			if !found[p] {
				found[p] = true
				paths = append(paths, p)
			}
		}
	}
	return paths
}

func getPackageRules(cmd *cobra.Command) (_ gogen.RuleSet, err error) {
	includes := viper.GetStringSlice(getPrefix(cmd) + "include-package")
	rules := make(gogen.RuleSet, len(includes))
	for i, pattern := range includes {
		rules[i], err = gogen.NewRule(pattern)
		if err != nil {
			return nil, err
		}
	}
	return rules, nil
}
