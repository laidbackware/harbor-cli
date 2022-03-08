// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/preheat"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationPreheatCreatePolicyCmd returns a cmd to handle operation createPolicy
func makeOperationPreheatCreatePolicyCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "CreatePolicy",
		Short: `Create a preheat policy under a project`,
		RunE:  runOperationPreheatCreatePolicy,
	}

	if err := registerOperationPreheatCreatePolicyParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPreheatCreatePolicy uses cmd flags to call endpoint api
func runOperationPreheatCreatePolicy(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := preheat.NewCreatePolicyParams()
	if err, _ := retrieveOperationPreheatCreatePolicyXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPreheatCreatePolicyPolicyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPreheatCreatePolicyProjectNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationPreheatCreatePolicyResult(appCli.Preheat.CreatePolicy(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationPreheatCreatePolicyParamFlags registers all flags needed to fill params
func registerOperationPreheatCreatePolicyParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPreheatCreatePolicyXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPreheatCreatePolicyPolicyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPreheatCreatePolicyProjectNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPreheatCreatePolicyXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationPreheatCreatePolicyPolicyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var policyFlagName string
	if cmdPrefix == "" {
		policyFlagName = "policy"
	} else {
		policyFlagName = fmt.Sprintf("%v.policy", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(policyFlagName, "", "Optional json string for [policy]. The policy schema info")

	// add flags for body
	if err := registerModelPreheatPolicyFlags(0, "preheatPolicy", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationPreheatCreatePolicyProjectNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationPreheatCreatePolicyXRequestIDFlag(m *preheat.CreatePolicyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationPreheatCreatePolicyPolicyFlag(m *preheat.CreatePolicyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("policy") {
		// Read policy string from cmd and unmarshal
		policyValueStr, err := cmd.Flags().GetString("policy")
		if err != nil {
			return err, false
		}

		policyValue := models.PreheatPolicy{}
		if err := json.Unmarshal([]byte(policyValueStr), &policyValue); err != nil {
			return fmt.Errorf("cannot unmarshal policy string in models.PreheatPolicy: %v", err), false
		}
		m.Policy = &policyValue
	}
	policyValueModel := m.Policy
	if swag.IsZero(policyValueModel) {
		policyValueModel = &models.PreheatPolicy{}
	}
	err, added := retrieveModelPreheatPolicyFlags(0, policyValueModel, "preheatPolicy", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Policy = policyValueModel
	}
	if dryRun && debug {

		policyValueDebugBytes, err := json.Marshal(m.Policy)
		if err != nil {
			return err, false
		}
		logDebugf("Policy dry-run payload: %v", string(policyValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationPreheatCreatePolicyProjectNameFlag(m *preheat.CreatePolicyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationPreheatCreatePolicyResult parses request result and return the string content
func parseOperationPreheatCreatePolicyResult(resp0 *preheat.CreatePolicyCreated, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning createPolicyCreated is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*preheat.CreatePolicyBadRequest)
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
		resp2, ok := iResp2.(*preheat.CreatePolicyUnauthorized)
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
		resp3, ok := iResp3.(*preheat.CreatePolicyForbidden)
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
		resp4, ok := iResp4.(*preheat.CreatePolicyConflict)
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
		resp5, ok := iResp5.(*preheat.CreatePolicyInternalServerError)
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

	// warning: non schema response createPolicyCreated is not supported by go-swagger cli yet.

	return "", nil
}