// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for ScannerCapability

// register flags to command
func registerModelScannerCapabilityFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerScannerCapabilityConsumesMimeTypes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerCapabilityProducesMimeTypes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerScannerCapabilityConsumesMimeTypes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: consumes_mime_types []string array type is not supported by go-swagger cli yet

	return nil
}

func registerScannerCapabilityProducesMimeTypes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: produces_mime_types []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelScannerCapabilityFlags(depth int, m *models.ScannerCapability, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, consumesMimeTypesAdded := retrieveScannerCapabilityConsumesMimeTypesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || consumesMimeTypesAdded

	err, producesMimeTypesAdded := retrieveScannerCapabilityProducesMimeTypesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || producesMimeTypesAdded

	return nil, retAdded
}

func retrieveScannerCapabilityConsumesMimeTypesFlags(depth int, m *models.ScannerCapability, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	consumesMimeTypesFlagName := fmt.Sprintf("%v.consumes_mime_types", cmdPrefix)
	if cmd.Flags().Changed(consumesMimeTypesFlagName) {
		// warning: consumes_mime_types array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveScannerCapabilityProducesMimeTypesFlags(depth int, m *models.ScannerCapability, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	producesMimeTypesFlagName := fmt.Sprintf("%v.produces_mime_types", cmdPrefix)
	if cmd.Flags().Changed(producesMimeTypesFlagName) {
		// warning: produces_mime_types array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
