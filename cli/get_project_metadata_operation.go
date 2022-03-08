// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/project_metadata"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationProjectMetadataGetProjectMetadataCmd returns a cmd to handle operation getProjectMetadata
func makeOperationProjectMetadataGetProjectMetadataCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getProjectMetadata",
		Short: `Get the specific metadata of the specific project`,
		RunE:  runOperationProjectMetadataGetProjectMetadata,
	}

	if err := registerOperationProjectMetadataGetProjectMetadataParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationProjectMetadataGetProjectMetadata uses cmd flags to call endpoint api
func runOperationProjectMetadataGetProjectMetadata(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := project_metadata.NewGetProjectMetadataParams()
	if err, _ := retrieveOperationProjectMetadataGetProjectMetadataXIsResourceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectMetadataGetProjectMetadataXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectMetadataGetProjectMetadataMetaNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectMetadataGetProjectMetadataProjectNameOrIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationProjectMetadataGetProjectMetadataResult(appCli.ProjectMetadata.GetProjectMetadata(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationProjectMetadataGetProjectMetadataParamFlags registers all flags needed to fill params
func registerOperationProjectMetadataGetProjectMetadataParamFlags(cmd *cobra.Command) error {
	if err := registerOperationProjectMetadataGetProjectMetadataXIsResourceNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectMetadataGetProjectMetadataXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectMetadataGetProjectMetadataMetaNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectMetadataGetProjectMetadataProjectNameOrIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationProjectMetadataGetProjectMetadataXIsResourceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationProjectMetadataGetProjectMetadataXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationProjectMetadataGetProjectMetadataMetaNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	metaNameDescription := `Required. The name of metadata.`

	var metaNameFlagName string
	if cmdPrefix == "" {
		metaNameFlagName = "meta_name"
	} else {
		metaNameFlagName = fmt.Sprintf("%v.meta_name", cmdPrefix)
	}

	var metaNameFlagDefault string

	_ = cmd.PersistentFlags().String(metaNameFlagName, metaNameFlagDefault, metaNameDescription)

	return nil
}
func registerOperationProjectMetadataGetProjectMetadataProjectNameOrIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationProjectMetadataGetProjectMetadataXIsResourceNameFlag(m *project_metadata.GetProjectMetadataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationProjectMetadataGetProjectMetadataXRequestIDFlag(m *project_metadata.GetProjectMetadataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationProjectMetadataGetProjectMetadataMetaNameFlag(m *project_metadata.GetProjectMetadataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("meta_name") {

		var metaNameFlagName string
		if cmdPrefix == "" {
			metaNameFlagName = "meta_name"
		} else {
			metaNameFlagName = fmt.Sprintf("%v.meta_name", cmdPrefix)
		}

		metaNameFlagValue, err := cmd.Flags().GetString(metaNameFlagName)
		if err != nil {
			return err, false
		}
		m.MetaName = metaNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationProjectMetadataGetProjectMetadataProjectNameOrIDFlag(m *project_metadata.GetProjectMetadataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationProjectMetadataGetProjectMetadataResult parses request result and return the string content
func parseOperationProjectMetadataGetProjectMetadataResult(resp0 *project_metadata.GetProjectMetadataOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*project_metadata.GetProjectMetadataOK)
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
		resp1, ok := iResp1.(*project_metadata.GetProjectMetadataBadRequest)
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
		resp2, ok := iResp2.(*project_metadata.GetProjectMetadataUnauthorized)
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
		resp3, ok := iResp3.(*project_metadata.GetProjectMetadataForbidden)
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
		resp4, ok := iResp4.(*project_metadata.GetProjectMetadataNotFound)
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
		resp5, ok := iResp5.(*project_metadata.GetProjectMetadataInternalServerError)
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