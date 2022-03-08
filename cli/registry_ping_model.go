// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for RegistryPing

// register flags to command
func registerModelRegistryPingFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRegistryPingAccessKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegistryPingAccessSecret(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegistryPingCredentialType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegistryPingID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegistryPingInsecure(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegistryPingType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegistryPingURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRegistryPingAccessKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	accessKeyDescription := `The registry access key.`

	var accessKeyFlagName string
	if cmdPrefix == "" {
		accessKeyFlagName = "access_key"
	} else {
		accessKeyFlagName = fmt.Sprintf("%v.access_key", cmdPrefix)
	}

	var accessKeyFlagDefault string

	_ = cmd.PersistentFlags().String(accessKeyFlagName, accessKeyFlagDefault, accessKeyDescription)

	return nil
}

func registerRegistryPingAccessSecret(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	accessSecretDescription := `The registry access secret.`

	var accessSecretFlagName string
	if cmdPrefix == "" {
		accessSecretFlagName = "access_secret"
	} else {
		accessSecretFlagName = fmt.Sprintf("%v.access_secret", cmdPrefix)
	}

	var accessSecretFlagDefault string

	_ = cmd.PersistentFlags().String(accessSecretFlagName, accessSecretFlagDefault, accessSecretDescription)

	return nil
}

func registerRegistryPingCredentialType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	credentialTypeDescription := `Credential type of the registry, e.g. basic.`

	var credentialTypeFlagName string
	if cmdPrefix == "" {
		credentialTypeFlagName = "credential_type"
	} else {
		credentialTypeFlagName = fmt.Sprintf("%v.credential_type", cmdPrefix)
	}

	var credentialTypeFlagDefault string

	_ = cmd.PersistentFlags().String(credentialTypeFlagName, credentialTypeFlagDefault, credentialTypeDescription)

	return nil
}

func registerRegistryPingID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `The registry ID.`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault int64

	_ = cmd.PersistentFlags().Int64(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerRegistryPingInsecure(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	insecureDescription := `Whether or not the certificate will be verified when Harbor tries to access the server.`

	var insecureFlagName string
	if cmdPrefix == "" {
		insecureFlagName = "insecure"
	} else {
		insecureFlagName = fmt.Sprintf("%v.insecure", cmdPrefix)
	}

	var insecureFlagDefault bool

	_ = cmd.PersistentFlags().Bool(insecureFlagName, insecureFlagDefault, insecureDescription)

	return nil
}

func registerRegistryPingType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Type of the registry, e.g. harbor.`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	return nil
}

func registerRegistryPingURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	urlDescription := `The registry URL.`

	var urlFlagName string
	if cmdPrefix == "" {
		urlFlagName = "url"
	} else {
		urlFlagName = fmt.Sprintf("%v.url", cmdPrefix)
	}

	var urlFlagDefault string

	_ = cmd.PersistentFlags().String(urlFlagName, urlFlagDefault, urlDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRegistryPingFlags(depth int, m *models.RegistryPing, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, accessKeyAdded := retrieveRegistryPingAccessKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || accessKeyAdded

	err, accessSecretAdded := retrieveRegistryPingAccessSecretFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || accessSecretAdded

	err, credentialTypeAdded := retrieveRegistryPingCredentialTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || credentialTypeAdded

	err, idAdded := retrieveRegistryPingIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, insecureAdded := retrieveRegistryPingInsecureFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || insecureAdded

	err, typeAdded := retrieveRegistryPingTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	err, urlAdded := retrieveRegistryPingURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || urlAdded

	return nil, retAdded
}

func retrieveRegistryPingAccessKeyFlags(depth int, m *models.RegistryPing, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	accessKeyFlagName := fmt.Sprintf("%v.access_key", cmdPrefix)
	if cmd.Flags().Changed(accessKeyFlagName) {

		var accessKeyFlagName string
		if cmdPrefix == "" {
			accessKeyFlagName = "access_key"
		} else {
			accessKeyFlagName = fmt.Sprintf("%v.access_key", cmdPrefix)
		}

		accessKeyFlagValue, err := cmd.Flags().GetString(accessKeyFlagName)
		if err != nil {
			return err, false
		}
		m.AccessKey = &accessKeyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRegistryPingAccessSecretFlags(depth int, m *models.RegistryPing, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	accessSecretFlagName := fmt.Sprintf("%v.access_secret", cmdPrefix)
	if cmd.Flags().Changed(accessSecretFlagName) {

		var accessSecretFlagName string
		if cmdPrefix == "" {
			accessSecretFlagName = "access_secret"
		} else {
			accessSecretFlagName = fmt.Sprintf("%v.access_secret", cmdPrefix)
		}

		accessSecretFlagValue, err := cmd.Flags().GetString(accessSecretFlagName)
		if err != nil {
			return err, false
		}
		m.AccessSecret = &accessSecretFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRegistryPingCredentialTypeFlags(depth int, m *models.RegistryPing, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	credentialTypeFlagName := fmt.Sprintf("%v.credential_type", cmdPrefix)
	if cmd.Flags().Changed(credentialTypeFlagName) {

		var credentialTypeFlagName string
		if cmdPrefix == "" {
			credentialTypeFlagName = "credential_type"
		} else {
			credentialTypeFlagName = fmt.Sprintf("%v.credential_type", cmdPrefix)
		}

		credentialTypeFlagValue, err := cmd.Flags().GetString(credentialTypeFlagName)
		if err != nil {
			return err, false
		}
		m.CredentialType = &credentialTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRegistryPingIDFlags(depth int, m *models.RegistryPing, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetInt64(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = &idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRegistryPingInsecureFlags(depth int, m *models.RegistryPing, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	insecureFlagName := fmt.Sprintf("%v.insecure", cmdPrefix)
	if cmd.Flags().Changed(insecureFlagName) {

		var insecureFlagName string
		if cmdPrefix == "" {
			insecureFlagName = "insecure"
		} else {
			insecureFlagName = fmt.Sprintf("%v.insecure", cmdPrefix)
		}

		insecureFlagValue, err := cmd.Flags().GetBool(insecureFlagName)
		if err != nil {
			return err, false
		}
		m.Insecure = &insecureFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRegistryPingTypeFlags(depth int, m *models.RegistryPing, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = &typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRegistryPingURLFlags(depth int, m *models.RegistryPing, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	urlFlagName := fmt.Sprintf("%v.url", cmdPrefix)
	if cmd.Flags().Changed(urlFlagName) {

		var urlFlagName string
		if cmdPrefix == "" {
			urlFlagName = "url"
		} else {
			urlFlagName = fmt.Sprintf("%v.url", cmdPrefix)
		}

		urlFlagValue, err := cmd.Flags().GetString(urlFlagName)
		if err != nil {
			return err, false
		}
		m.URL = &urlFlagValue

		retAdded = true
	}

	return nil, retAdded
}
