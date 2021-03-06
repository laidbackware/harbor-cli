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

// makeOperationRegistryCreateRegistryCmd returns a cmd to handle operation createRegistry
func makeOperationRegistryCreateRegistryCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "createRegistry",
		Short: `Create a registry`,
		RunE:  runOperationRegistryCreateRegistry,
	}

	if err := registerOperationRegistryCreateRegistryParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationRegistryCreateRegistry uses cmd flags to call endpoint api
func runOperationRegistryCreateRegistry(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := registry.NewCreateRegistryParams()
	if err, _ := retrieveOperationRegistryCreateRegistryXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationRegistryCreateRegistryRegistryFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationRegistryCreateRegistryResult(appCli.Registry.CreateRegistry(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationRegistryCreateRegistryParamFlags registers all flags needed to fill params
func registerOperationRegistryCreateRegistryParamFlags(cmd *cobra.Command) error {
	if err := registerOperationRegistryCreateRegistryXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationRegistryCreateRegistryRegistryParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationRegistryCreateRegistryXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationRegistryCreateRegistryRegistryParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var registryFlagName string
	if cmdPrefix == "" {
		registryFlagName = "registry"
	} else {
		registryFlagName = fmt.Sprintf("%v.registry", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(registryFlagName, "", "Optional json string for [registry]. The registry")

	// add flags for body
	if err := registerModelRegistryFlags(0, "registry", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationRegistryCreateRegistryXRequestIDFlag(m *registry.CreateRegistryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationRegistryCreateRegistryRegistryFlag(m *registry.CreateRegistryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("registry") {
		// Read registry string from cmd and unmarshal
		registryValueStr, err := cmd.Flags().GetString("registry")
		if err != nil {
			return err, false
		}

		registryValue := models.Registry{}
		if err := json.Unmarshal([]byte(registryValueStr), &registryValue); err != nil {
			return fmt.Errorf("cannot unmarshal registry string in models.Registry: %v", err), false
		}
		m.Registry = &registryValue
	}
	registryValueModel := m.Registry
	if swag.IsZero(registryValueModel) {
		registryValueModel = &models.Registry{}
	}
	err, added := retrieveModelRegistryFlags(0, registryValueModel, "registry", cmd)
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

// parseOperationRegistryCreateRegistryResult parses request result and return the string content
func parseOperationRegistryCreateRegistryResult(resp0 *registry.CreateRegistryCreated, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning createRegistryCreated is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*registry.CreateRegistryBadRequest)
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
		resp2, ok := iResp2.(*registry.CreateRegistryUnauthorized)
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
		resp3, ok := iResp3.(*registry.CreateRegistryForbidden)
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
		resp4, ok := iResp4.(*registry.CreateRegistryConflict)
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
		resp5, ok := iResp5.(*registry.CreateRegistryInternalServerError)
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

	// warning: non schema response createRegistryCreated is not supported by go-swagger cli yet.

	return "", nil
}
