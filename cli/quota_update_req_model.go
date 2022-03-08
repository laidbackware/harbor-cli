// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for QuotaUpdateReq

// register flags to command
func registerModelQuotaUpdateReqFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerQuotaUpdateReqHard(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerQuotaUpdateReqHard(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: hard ResourceList map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelQuotaUpdateReqFlags(depth int, m *models.QuotaUpdateReq, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, hardAdded := retrieveQuotaUpdateReqHardFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hardAdded

	return nil, retAdded
}

func retrieveQuotaUpdateReqHardFlags(depth int, m *models.QuotaUpdateReq, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	hardFlagName := fmt.Sprintf("%v.hard", cmdPrefix)
	if cmd.Flags().Changed(hardFlagName) {
		// warning: hard map type ResourceList is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
