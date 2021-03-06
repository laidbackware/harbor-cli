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

// makeOperationPreheatDeletePolicyCmd returns a cmd to handle operation deletePolicy
func makeOperationPreheatDeletePolicyCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "DeletePolicy",
		Short: `Delete a preheat policy`,
		RunE:  runOperationPreheatDeletePolicy,
	}

	if err := registerOperationPreheatDeletePolicyParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPreheatDeletePolicy uses cmd flags to call endpoint api
func runOperationPreheatDeletePolicy(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := preheat.NewDeletePolicyParams()
	if err, _ := retrieveOperationPreheatDeletePolicyXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPreheatDeletePolicyPreheatPolicyNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPreheatDeletePolicyProjectNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationPreheatDeletePolicyResult(appCli.Preheat.DeletePolicy(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationPreheatDeletePolicyParamFlags registers all flags needed to fill params
func registerOperationPreheatDeletePolicyParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPreheatDeletePolicyXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPreheatDeletePolicyPreheatPolicyNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPreheatDeletePolicyProjectNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPreheatDeletePolicyXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationPreheatDeletePolicyPreheatPolicyNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	preheatPolicyNameDescription := `Required. Preheat Policy Name`

	var preheatPolicyNameFlagName string
	if cmdPrefix == "" {
		preheatPolicyNameFlagName = "preheat_policy_name"
	} else {
		preheatPolicyNameFlagName = fmt.Sprintf("%v.preheat_policy_name", cmdPrefix)
	}

	var preheatPolicyNameFlagDefault string

	_ = cmd.PersistentFlags().String(preheatPolicyNameFlagName, preheatPolicyNameFlagDefault, preheatPolicyNameDescription)

	return nil
}
func registerOperationPreheatDeletePolicyProjectNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	projectNameDescription := `Required. The name of the project`

	var projectNameFlagName string
	if cmdPrefix == "" {
		projectNameFlagName = "project_name"
	} else {
		projectNameFlagName = fmt.Sprintf("%v.project_name", cmdPrefix)
	}

	var projectNameFlagDefault string

	_ = cmd.PersistentFlags().String(projectNameFlagName, projectNameFlagDefault, projectNameDescription)

	return nil
}

func retrieveOperationPreheatDeletePolicyXRequestIDFlag(m *preheat.DeletePolicyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationPreheatDeletePolicyPreheatPolicyNameFlag(m *preheat.DeletePolicyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("preheat_policy_name") {

		var preheatPolicyNameFlagName string
		if cmdPrefix == "" {
			preheatPolicyNameFlagName = "preheat_policy_name"
		} else {
			preheatPolicyNameFlagName = fmt.Sprintf("%v.preheat_policy_name", cmdPrefix)
		}

		preheatPolicyNameFlagValue, err := cmd.Flags().GetString(preheatPolicyNameFlagName)
		if err != nil {
			return err, false
		}
		m.PreheatPolicyName = preheatPolicyNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPreheatDeletePolicyProjectNameFlag(m *preheat.DeletePolicyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("project_name") {

		var projectNameFlagName string
		if cmdPrefix == "" {
			projectNameFlagName = "project_name"
		} else {
			projectNameFlagName = fmt.Sprintf("%v.project_name", cmdPrefix)
		}

		projectNameFlagValue, err := cmd.Flags().GetString(projectNameFlagName)
		if err != nil {
			return err, false
		}
		m.ProjectName = projectNameFlagValue

	}
	return nil, retAdded
}

// parseOperationPreheatDeletePolicyResult parses request result and return the string content
func parseOperationPreheatDeletePolicyResult(resp0 *preheat.DeletePolicyOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning deletePolicyOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*preheat.DeletePolicyBadRequest)
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
		resp2, ok := iResp2.(*preheat.DeletePolicyUnauthorized)
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
		resp3, ok := iResp3.(*preheat.DeletePolicyForbidden)
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
		resp4, ok := iResp4.(*preheat.DeletePolicyNotFound)
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
		resp5, ok := iResp5.(*preheat.DeletePolicyInternalServerError)
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

	// warning: non schema response deletePolicyOK is not supported by go-swagger cli yet.

	return "", nil
}
