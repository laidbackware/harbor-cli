// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for Scanner

// register flags to command
func registerModelScannerFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerScannerName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerVendor(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerScannerName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Name of the scanner`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerScannerVendor(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vendorDescription := `Name of the scanner provider`

	var vendorFlagName string
	if cmdPrefix == "" {
		vendorFlagName = "vendor"
	} else {
		vendorFlagName = fmt.Sprintf("%v.vendor", cmdPrefix)
	}

	var vendorFlagDefault string

	_ = cmd.PersistentFlags().String(vendorFlagName, vendorFlagDefault, vendorDescription)

	return nil
}

func registerScannerVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	versionDescription := `Version of the scanner adapter`

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "version"
	} else {
		versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
	}

	var versionFlagDefault string

	_ = cmd.PersistentFlags().String(versionFlagName, versionFlagDefault, versionDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelScannerFlags(depth int, m *models.Scanner, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, nameAdded := retrieveScannerNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, vendorAdded := retrieveScannerVendorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vendorAdded

	err, versionAdded := retrieveScannerVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	return nil, retAdded
}

func retrieveScannerNameFlags(depth int, m *models.Scanner, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerVendorFlags(depth int, m *models.Scanner, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vendorFlagName := fmt.Sprintf("%v.vendor", cmdPrefix)
	if cmd.Flags().Changed(vendorFlagName) {

		var vendorFlagName string
		if cmdPrefix == "" {
			vendorFlagName = "vendor"
		} else {
			vendorFlagName = fmt.Sprintf("%v.vendor", cmdPrefix)
		}

		vendorFlagValue, err := cmd.Flags().GetString(vendorFlagName)
		if err != nil {
			return err, false
		}
		m.Vendor = vendorFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerVersionFlags(depth int, m *models.Scanner, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	versionFlagName := fmt.Sprintf("%v.version", cmdPrefix)
	if cmd.Flags().Changed(versionFlagName) {

		var versionFlagName string
		if cmdPrefix == "" {
			versionFlagName = "version"
		} else {
			versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
		}

		versionFlagValue, err := cmd.Flags().GetString(versionFlagName)
		if err != nil {
			return err, false
		}
		m.Version = versionFlagValue

		retAdded = true
	}

	return nil, retAdded
}
