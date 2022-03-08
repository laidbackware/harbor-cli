// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for LdapFailedImportUser

// register flags to command
func registerModelLdapFailedImportUserFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerLdapFailedImportUserError(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLdapFailedImportUserUID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerLdapFailedImportUserError(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	errorDescription := `fail reason.`

	var errorFlagName string
	if cmdPrefix == "" {
		errorFlagName = "error"
	} else {
		errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
	}

	var errorFlagDefault string

	_ = cmd.PersistentFlags().String(errorFlagName, errorFlagDefault, errorDescription)

	return nil
}

func registerLdapFailedImportUserUID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	uidDescription := `the uid cant add to system.`

	var uidFlagName string
	if cmdPrefix == "" {
		uidFlagName = "uid"
	} else {
		uidFlagName = fmt.Sprintf("%v.uid", cmdPrefix)
	}

	var uidFlagDefault string

	_ = cmd.PersistentFlags().String(uidFlagName, uidFlagDefault, uidDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelLdapFailedImportUserFlags(depth int, m *models.LdapFailedImportUser, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, errorAdded := retrieveLdapFailedImportUserErrorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorAdded

	err, uidAdded := retrieveLdapFailedImportUserUIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || uidAdded

	return nil, retAdded
}

func retrieveLdapFailedImportUserErrorFlags(depth int, m *models.LdapFailedImportUser, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	errorFlagName := fmt.Sprintf("%v.error", cmdPrefix)
	if cmd.Flags().Changed(errorFlagName) {

		var errorFlagName string
		if cmdPrefix == "" {
			errorFlagName = "error"
		} else {
			errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
		}

		errorFlagValue, err := cmd.Flags().GetString(errorFlagName)
		if err != nil {
			return err, false
		}
		m.Error = errorFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveLdapFailedImportUserUIDFlags(depth int, m *models.LdapFailedImportUser, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	uidFlagName := fmt.Sprintf("%v.uid", cmdPrefix)
	if cmd.Flags().Changed(uidFlagName) {

		var uidFlagName string
		if cmdPrefix == "" {
			uidFlagName = "uid"
		} else {
			uidFlagName = fmt.Sprintf("%v.uid", cmdPrefix)
		}

		uidFlagValue, err := cmd.Flags().GetString(uidFlagName)
		if err != nil {
			return err, false
		}
		m.UID = uidFlagValue

		retAdded = true
	}

	return nil, retAdded
}
