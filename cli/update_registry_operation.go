// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/registry"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationRegistryUpdateRegistryCmd returns a cmd to handle operation updateRegistry
func makeOperationRegistryUpdateRegistryCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "updateRegistry",
		Short: `Update the registry`,
		RunE:  runOperationRegistryUpdateRegistry,
	}

	if err := registerOperationRegistryUpdateRegistryParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationRegistryUpdateRegistry uses cmd flags to call endpoint api
func runOperationRegistryUpdateRegistry(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := registry.NewUpdateRegistryParams()
	if err, _ := retrieveOperationRegistryUpdateRegistryXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationRegistryUpdateRegistryIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationRegistryUpdateRegistryRegistryFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationRegistryUpdateRegistryResult(appCli.Registry.UpdateRegistry(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationRegistryUpdateRegistryParamFlags registers all flags needed to fill params
func registerOperationRegistryUpdateRegistryParamFlags(cmd *cobra.Command) error {
	if err := registerOperationRegistryUpdateRegistryXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationRegistryUpdateRegistryIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationRegistryUpdateRegistryRegistryParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationRegistryUpdateRegistryXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	xRequestIdDescription := `An unique ID for the request`

	var xRequestIdFlagName string
	if cmdPrefix == "" {
		xRequestIdFlagName = "X-Request-Id"
	} else {
		xRequestIdFlagName = fmt.Sprintf("%v.X-Request-Id", cmdPrefix)
	}

	var xRequestIdFlagDefault string

	_ = cmd.PersistentFlags().String(xRequestIdFlagName, xRequestIdFlagDefault, xRequestIdDescription)

	return nil
}
func registerOperationRegistryUpdateRegistryIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. The registry ID`

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
func registerOperationRegistryUpdateRegistryRegistryParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var registryFlagName string
	if cmdPrefix == "" {
		registryFlagName = "registry"
	} else {
		registryFlagName = fmt.Sprintf("%v.registry", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(registryFlagName, "", "Optional json string for [registry]. The registry")

	// add flags for body
	if err := registerModelRegistryUpdateFlags(0, "registryUpdate", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationRegistryUpdateRegistryXRequestIDFlag(m *registry.UpdateRegistryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("X-Request-Id") {

		var xRequestIdFlagName string
		if cmdPrefix == "" {
			xRequestIdFlagName = "X-Request-Id"
		} else {
			xRequestIdFlagName = fmt.Sprintf("%v.X-Request-Id", cmdPrefix)
		}

		xRequestIdFlagValue, err := cmd.Flags().GetString(xRequestIdFlagName)
		if err != nil {
			return err, false
		}
		m.XRequestID = &xRequestIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationRegistryUpdateRegistryIDFlag(m *registry.UpdateRegistryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

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
		m.ID = idFlagValue

	}
	return nil, retAdded
}
func retrieveOperationRegistryUpdateRegistryRegistryFlag(m *registry.UpdateRegistryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("registry") {
		// Read registry string from cmd and unmarshal
		registryValueStr, err := cmd.Flags().GetString("registry")
		if err != nil {
			return err, false
		}

		registryValue := models.RegistryUpdate{}
		if err := json.Unmarshal([]byte(registryValueStr), &registryValue); err != nil {
			return fmt.Errorf("cannot unmarshal registry string in models.RegistryUpdate: %v", err), false
		}
		m.Registry = &registryValue
	}
	registryValueModel := m.Registry
	if swag.IsZero(registryValueModel) {
		registryValueModel = &models.RegistryUpdate{}
	}
	err, added := retrieveModelRegistryUpdateFlags(0, registryValueModel, "registryUpdate", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Registry = registryValueModel
	}
	if dryRun && debug {

		registryValueDebugBytes, err := json.Marshal(m.Registry)
		if err != nil {
			return err, false
		}
		logDebugf("Registry dry-run payload: %v", string(registryValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationRegistryUpdateRegistryResult parses request result and return the string content
func parseOperationRegistryUpdateRegistryResult(resp0 *registry.UpdateRegistryOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning updateRegistryOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*registry.UpdateRegistryUnauthorized)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*registry.UpdateRegistryForbidden)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*registry.UpdateRegistryNotFound)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp4 interface{} = respErr
		resp4, ok := iResp4.(*registry.UpdateRegistryConflict)
		if ok {
			if !swag.IsZero(resp4) && !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp5 interface{} = respErr
		resp5, ok := iResp5.(*registry.UpdateRegistryInternalServerError)
		if ok {
			if !swag.IsZero(resp5) && !swag.IsZero(resp5.Payload) {
				msgStr, err := json.Marshal(resp5.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response updateRegistryOK is not supported by go-swagger cli yet.

	return "", nil
}
