// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/member"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationMemberListProjectMembersCmd returns a cmd to handle operation listProjectMembers
func makeOperationMemberListProjectMembersCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "listProjectMembers",
		Short: `Get all project member information`,
		RunE:  runOperationMemberListProjectMembers,
	}

	if err := registerOperationMemberListProjectMembersParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationMemberListProjectMembers uses cmd flags to call endpoint api
func runOperationMemberListProjectMembers(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := member.NewListProjectMembersParams()
	if err, _ := retrieveOperationMemberListProjectMembersXIsResourceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationMemberListProjectMembersXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationMemberListProjectMembersEntitynameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationMemberListProjectMembersPageFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationMemberListProjectMembersPageSizeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationMemberListProjectMembersProjectNameOrIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationMemberListProjectMembersResult(appCli.Member.ListProjectMembers(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationMemberListProjectMembersParamFlags registers all flags needed to fill params
func registerOperationMemberListProjectMembersParamFlags(cmd *cobra.Command) error {
	if err := registerOperationMemberListProjectMembersXIsResourceNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationMemberListProjectMembersXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationMemberListProjectMembersEntitynameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationMemberListProjectMembersPageParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationMemberListProjectMembersPageSizeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationMemberListProjectMembersProjectNameOrIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationMemberListProjectMembersXIsResourceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationMemberListProjectMembersXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationMemberListProjectMembersEntitynameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	entitynameDescription := `The entity name to search.`

	var entitynameFlagName string
	if cmdPrefix == "" {
		entitynameFlagName = "entityname"
	} else {
		entitynameFlagName = fmt.Sprintf("%v.entityname", cmdPrefix)
	}

	var entitynameFlagDefault string

	_ = cmd.PersistentFlags().String(entitynameFlagName, entitynameFlagDefault, entitynameDescription)

	return nil
}
func registerOperationMemberListProjectMembersPageParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	pageDescription := `The page number`

	var pageFlagName string
	if cmdPrefix == "" {
		pageFlagName = "page"
	} else {
		pageFlagName = fmt.Sprintf("%v.page", cmdPrefix)
	}

	var pageFlagDefault int64 = 1

	_ = cmd.PersistentFlags().Int64(pageFlagName, pageFlagDefault, pageDescription)

	return nil
}
func registerOperationMemberListProjectMembersPageSizeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	pageSizeDescription := `The size of per page`

	var pageSizeFlagName string
	if cmdPrefix == "" {
		pageSizeFlagName = "page_size"
	} else {
		pageSizeFlagName = fmt.Sprintf("%v.page_size", cmdPrefix)
	}

	var pageSizeFlagDefault int64 = 10

	_ = cmd.PersistentFlags().Int64(pageSizeFlagName, pageSizeFlagDefault, pageSizeDescription)

	return nil
}
func registerOperationMemberListProjectMembersProjectNameOrIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationMemberListProjectMembersXIsResourceNameFlag(m *member.ListProjectMembersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationMemberListProjectMembersXRequestIDFlag(m *member.ListProjectMembersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationMemberListProjectMembersEntitynameFlag(m *member.ListProjectMembersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("entityname") {

		var entitynameFlagName string
		if cmdPrefix == "" {
			entitynameFlagName = "entityname"
		} else {
			entitynameFlagName = fmt.Sprintf("%v.entityname", cmdPrefix)
		}

		entitynameFlagValue, err := cmd.Flags().GetString(entitynameFlagName)
		if err != nil {
			return err, false
		}
		m.Entityname = &entitynameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationMemberListProjectMembersPageFlag(m *member.ListProjectMembersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("page") {

		var pageFlagName string
		if cmdPrefix == "" {
			pageFlagName = "page"
		} else {
			pageFlagName = fmt.Sprintf("%v.page", cmdPrefix)
		}

		pageFlagValue, err := cmd.Flags().GetInt64(pageFlagName)
		if err != nil {
			return err, false
		}
		m.Page = &pageFlagValue

	}
	return nil, retAdded
}
func retrieveOperationMemberListProjectMembersPageSizeFlag(m *member.ListProjectMembersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("page_size") {

		var pageSizeFlagName string
		if cmdPrefix == "" {
			pageSizeFlagName = "page_size"
		} else {
			pageSizeFlagName = fmt.Sprintf("%v.page_size", cmdPrefix)
		}

		pageSizeFlagValue, err := cmd.Flags().GetInt64(pageSizeFlagName)
		if err != nil {
			return err, false
		}
		m.PageSize = &pageSizeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationMemberListProjectMembersProjectNameOrIDFlag(m *member.ListProjectMembersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationMemberListProjectMembersResult parses request result and return the string content
func parseOperationMemberListProjectMembersResult(resp0 *member.ListProjectMembersOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*member.ListProjectMembersOK)
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
		resp1, ok := iResp1.(*member.ListProjectMembersBadRequest)
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
		resp2, ok := iResp2.(*member.ListProjectMembersUnauthorized)
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
		resp3, ok := iResp3.(*member.ListProjectMembersForbidden)
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
		resp4, ok := iResp4.(*member.ListProjectMembersNotFound)
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
		resp5, ok := iResp5.(*member.ListProjectMembersInternalServerError)
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