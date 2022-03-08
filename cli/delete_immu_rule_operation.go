// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/immutable"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationImmutableDeleteImmuRuleCmd returns a cmd to handle operation deleteImmuRule
func makeOperationImmutableDeleteImmuRuleCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "DeleteImmuRule",
		Short: ``,
		RunE:  runOperationImmutableDeleteImmuRule,
	}

	if err := registerOperationImmutableDeleteImmuRuleParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationImmutableDeleteImmuRule uses cmd flags to call endpoint api
func runOperationImmutableDeleteImmuRule(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := immutable.NewDeleteImmuRuleParams()
	if err, _ := retrieveOperationImmutableDeleteImmuRuleXIsResourceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImmutableDeleteImmuRuleXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImmutableDeleteImmuRuleImmutableRuleIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImmutableDeleteImmuRuleProjectNameOrIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationImmutableDeleteImmuRuleResult(appCli.Immutable.DeleteImmuRule(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationImmutableDeleteImmuRuleParamFlags registers all flags needed to fill params
func registerOperationImmutableDeleteImmuRuleParamFlags(cmd *cobra.Command) error {
	if err := registerOperationImmutableDeleteImmuRuleXIsResourceNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImmutableDeleteImmuRuleXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImmutableDeleteImmuRuleImmutableRuleIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImmutableDeleteImmuRuleProjectNameOrIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationImmutableDeleteImmuRuleXIsResourceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	xIsResourceNameDescription := `The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.`

	var xIsResourceNameFlagName string
	if cmdPrefix == "" {
		xIsResourceNameFlagName = "X-Is-Resource-Name"
	} else {
		xIsResourceNameFlagName = fmt.Sprintf("%v.X-Is-Resource-Name", cmdPrefix)
	}

	var xIsResourceNameFlagDefault bool

	_ = cmd.PersistentFlags().Bool(xIsResourceNameFlagName, xIsResourceNameFlagDefault, xIsResourceNameDescription)

	return nil
}
func registerOperationImmutableDeleteImmuRuleXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationImmutableDeleteImmuRuleImmutableRuleIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	immutableRuleIdDescription := `Required. The ID of the immutable rule`

	var immutableRuleIdFlagName string
	if cmdPrefix == "" {
		immutableRuleIdFlagName = "immutable_rule_id"
	} else {
		immutableRuleIdFlagName = fmt.Sprintf("%v.immutable_rule_id", cmdPrefix)
	}

	var immutableRuleIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(immutableRuleIdFlagName, immutableRuleIdFlagDefault, immutableRuleIdDescription)

	return nil
}
func registerOperationImmutableDeleteImmuRuleProjectNameOrIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	projectNameOrIdDescription := `Required. The name or id of the project`

	var projectNameOrIdFlagName string
	if cmdPrefix == "" {
		projectNameOrIdFlagName = "project_name_or_id"
	} else {
		projectNameOrIdFlagName = fmt.Sprintf("%v.project_name_or_id", cmdPrefix)
	}

	var projectNameOrIdFlagDefault string

	_ = cmd.PersistentFlags().String(projectNameOrIdFlagName, projectNameOrIdFlagDefault, projectNameOrIdDescription)

	return nil
}

func retrieveOperationImmutableDeleteImmuRuleXIsResourceNameFlag(m *immutable.DeleteImmuRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("X-Is-Resource-Name") {

		var xIsResourceNameFlagName string
		if cmdPrefix == "" {
			xIsResourceNameFlagName = "X-Is-Resource-Name"
		} else {
			xIsResourceNameFlagName = fmt.Sprintf("%v.X-Is-Resource-Name", cmdPrefix)
		}

		xIsResourceNameFlagValue, err := cmd.Flags().GetBool(xIsResourceNameFlagName)
		if err != nil {
			return err, false
		}
		m.XIsResourceName = &xIsResourceNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImmutableDeleteImmuRuleXRequestIDFlag(m *immutable.DeleteImmuRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationImmutableDeleteImmuRuleImmutableRuleIDFlag(m *immutable.DeleteImmuRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("immutable_rule_id") {

		var immutableRuleIdFlagName string
		if cmdPrefix == "" {
			immutableRuleIdFlagName = "immutable_rule_id"
		} else {
			immutableRuleIdFlagName = fmt.Sprintf("%v.immutable_rule_id", cmdPrefix)
		}

		immutableRuleIdFlagValue, err := cmd.Flags().GetInt64(immutableRuleIdFlagName)
		if err != nil {
			return err, false
		}
		m.ImmutableRuleID = immutableRuleIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImmutableDeleteImmuRuleProjectNameOrIDFlag(m *immutable.DeleteImmuRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("project_name_or_id") {

		var projectNameOrIdFlagName string
		if cmdPrefix == "" {
			projectNameOrIdFlagName = "project_name_or_id"
		} else {
			projectNameOrIdFlagName = fmt.Sprintf("%v.project_name_or_id", cmdPrefix)
		}

		projectNameOrIdFlagValue, err := cmd.Flags().GetString(projectNameOrIdFlagName)
		if err != nil {
			return err, false
		}
		m.ProjectNameOrID = projectNameOrIdFlagValue

	}
	return nil, retAdded
}

// parseOperationImmutableDeleteImmuRuleResult parses request result and return the string content
func parseOperationImmutableDeleteImmuRuleResult(resp0 *immutable.DeleteImmuRuleOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning deleteImmuRuleOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*immutable.DeleteImmuRuleBadRequest)
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
		resp2, ok := iResp2.(*immutable.DeleteImmuRuleUnauthorized)
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
		resp3, ok := iResp3.(*immutable.DeleteImmuRuleForbidden)
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
		resp4, ok := iResp4.(*immutable.DeleteImmuRuleInternalServerError)
		if ok {
			if !swag.IsZero(resp4) && !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response deleteImmuRuleOK is not supported by go-swagger cli yet.

	return "", nil
}