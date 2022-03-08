// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for OIDCCliSecretReq

// register flags to command
func registerModelOIDCCliSecretReqFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerOIDCCliSecretReqSecret(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerOIDCCliSecretReqSecret(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	secretDescription := `The new secret`

	var secretFlagName string
	if cmdPrefix == "" {
		secretFlagName = "secret"
	} else {
		secretFlagName = fmt.Sprintf("%v.secret", cmdPrefix)
	}

	var secretFlagDefault string

	_ = cmd.PersistentFlags().String(secretFlagName, secretFlagDefault, secretDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelOIDCCliSecretReqFlags(depth int, m *models.OIDCCliSecretReq, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, secretAdded := retrieveOIDCCliSecretReqSecretFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || secretAdded

	return nil, retAdded
}

func retrieveOIDCCliSecretReqSecretFlags(depth int, m *models.OIDCCliSecretReq, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	secretFlagName := fmt.Sprintf("%v.secret", cmdPrefix)
	if cmd.Flags().Changed(secretFlagName) {

		var secretFlagName string
		if cmdPrefix == "" {
			secretFlagName = "secret"
		} else {
			secretFlagName = fmt.Sprintf("%v.secret", cmdPrefix)
		}

		secretFlagValue, err := cmd.Flags().GetString(secretFlagName)
		if err != nil {
			return err, false
		}
		m.Secret = secretFlagValue

		retAdded = true
	}

	return nil, retAdded
}