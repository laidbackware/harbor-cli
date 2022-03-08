// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for SystemInfo

// register flags to command
func registerModelSystemInfoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSystemInfoStorage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSystemInfoStorage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: storage []*Storage array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSystemInfoFlags(depth int, m *models.SystemInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, storageAdded := retrieveSystemInfoStorageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || storageAdded

	return nil, retAdded
}

func retrieveSystemInfoStorageFlags(depth int, m *models.SystemInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	storageFlagName := fmt.Sprintf("%v.storage", cmdPrefix)
	if cmd.Flags().Changed(storageFlagName) {
		// warning: storage array type []*Storage is not supported by go-swagger cli yet
	}

	return nil, retAdded
}