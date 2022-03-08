// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/project"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationProjectGetProjectSummaryCmd returns a cmd to handle operation getProjectSummary
func makeOperationProjectGetProjectSummaryCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getProjectSummary",
		Short: `Get summary of the project.`,
		RunE:  runOperationProjectGetProjectSummary,
	}

	if err := registerOperationProjectGetProjectSummaryParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationProjectGetProjectSummary uses cmd flags to call endpoint api
func runOperationProjectGetProjectSummary(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := project.NewGetProjectSummaryParams()
	if err, _ := retrieveOperationProjectGetProjectSummaryXIsResourceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectGetProjectSummaryXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectGetProjectSummaryProjectNameOrIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationProjectGetProjectSummaryResult(appCli.Project.GetProjectSummary(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationProjectGetProjectSummaryParamFlags registers all flags needed to fill params
func registerOperationProjectGetProjectSummaryParamFlags(cmd *cobra.Command) error {
	if err := registerOperationProjectGetProjectSummaryXIsResourceNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectGetProjectSummaryXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectGetProjectSummaryProjectNameOrIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationProjectGetProjectSummaryXIsResourceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationProjectGetProjectSummaryXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationProjectGetProjectSummaryProjectNameOrIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationProjectGetProjectSummaryXIsResourceNameFlag(m *project.GetProjectSummaryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationProjectGetProjectSummaryXRequestIDFlag(m *project.GetProjectSummaryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationProjectGetProjectSummaryProjectNameOrIDFlag(m *project.GetProjectSummaryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationProjectGetProjectSummaryResult parses request result and return the string content
func parseOperationProjectGetProjectSummaryResult(resp0 *project.GetProjectSummaryOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*project.GetProjectSummaryOK)
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
		resp1, ok := iResp1.(*project.GetProjectSummaryBadRequest)
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
		resp2, ok := iResp2.(*project.GetProjectSummaryUnauthorized)
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
		resp3, ok := iResp3.(*project.GetProjectSummaryForbidden)
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
		resp4, ok := iResp4.(*project.GetProjectSummaryNotFound)
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
		resp5, ok := iResp5.(*project.GetProjectSummaryInternalServerError)
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
