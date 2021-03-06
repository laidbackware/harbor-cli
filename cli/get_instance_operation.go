// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/preheat"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationPreheatGetInstanceCmd returns a cmd to handle operation getInstance
func makeOperationPreheatGetInstanceCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "GetInstance",
		Short: `Get a P2P provider instance`,
		RunE:  runOperationPreheatGetInstance,
	}

	if err := registerOperationPreheatGetInstanceParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPreheatGetInstance uses cmd flags to call endpoint api
func runOperationPreheatGetInstance(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := preheat.NewGetInstanceParams()
	if err, _ := retrieveOperationPreheatGetInstanceXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPreheatGetInstancePreheatInstanceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationPreheatGetInstanceResult(appCli.Preheat.GetInstance(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationPreheatGetInstanceParamFlags registers all flags needed to fill params
func registerOperationPreheatGetInstanceParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPreheatGetInstanceXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPreheatGetInstancePreheatInstanceNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPreheatGetInstanceXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationPreheatGetInstancePreheatInstanceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	preheatInstanceNameDescription := `Required. Instance Name`

	var preheatInstanceNameFlagName string
	if cmdPrefix == "" {
		preheatInstanceNameFlagName = "preheat_instance_name"
	} else {
		preheatInstanceNameFlagName = fmt.Sprintf("%v.preheat_instance_name", cmdPrefix)
	}

	var preheatInstanceNameFlagDefault string

	_ = cmd.PersistentFlags().String(preheatInstanceNameFlagName, preheatInstanceNameFlagDefault, preheatInstanceNameDescription)

	return nil
}

func retrieveOperationPreheatGetInstanceXRequestIDFlag(m *preheat.GetInstanceParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationPreheatGetInstancePreheatInstanceNameFlag(m *preheat.GetInstanceParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("preheat_instance_name") {

		var preheatInstanceNameFlagName string
		if cmdPrefix == "" {
			preheatInstanceNameFlagName = "preheat_instance_name"
		} else {
			preheatInstanceNameFlagName = fmt.Sprintf("%v.preheat_instance_name", cmdPrefix)
		}

		preheatInstanceNameFlagValue, err := cmd.Flags().GetString(preheatInstanceNameFlagName)
		if err != nil {
			return err, false
		}
		m.PreheatInstanceName = preheatInstanceNameFlagValue

	}
	return nil, retAdded
}

// parseOperationPreheatGetInstanceResult parses request result and return the string content
func parseOperationPreheatGetInstanceResult(resp0 *preheat.GetInstanceOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*preheat.GetInstanceOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*preheat.GetInstanceBadRequest)
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
		resp2, ok := iResp2.(*preheat.GetInstanceUnauthorized)
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
		resp3, ok := iResp3.(*preheat.GetInstanceForbidden)
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
		resp4, ok := iResp4.(*preheat.GetInstanceNotFound)
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
		resp5, ok := iResp5.(*preheat.GetInstanceInternalServerError)
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

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
