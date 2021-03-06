// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/retention"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationRetentionListRetentionTasksCmd returns a cmd to handle operation listRetentionTasks
func makeOperationRetentionListRetentionTasksCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "listRetentionTasks",
		Short: `Get Retention tasks, each repository as a task.`,
		RunE:  runOperationRetentionListRetentionTasks,
	}

	if err := registerOperationRetentionListRetentionTasksParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationRetentionListRetentionTasks uses cmd flags to call endpoint api
func runOperationRetentionListRetentionTasks(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := retention.NewListRetentionTasksParams()
	if err, _ := retrieveOperationRetentionListRetentionTasksXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationRetentionListRetentionTasksEidFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationRetentionListRetentionTasksIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationRetentionListRetentionTasksPageFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationRetentionListRetentionTasksPageSizeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationRetentionListRetentionTasksResult(appCli.Retention.ListRetentionTasks(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationRetentionListRetentionTasksParamFlags registers all flags needed to fill params
func registerOperationRetentionListRetentionTasksParamFlags(cmd *cobra.Command) error {
	if err := registerOperationRetentionListRetentionTasksXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationRetentionListRetentionTasksEidParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationRetentionListRetentionTasksIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationRetentionListRetentionTasksPageParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationRetentionListRetentionTasksPageSizeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationRetentionListRetentionTasksXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationRetentionListRetentionTasksEidParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	eidDescription := `Required. Retention execution ID.`

	var eidFlagName string
	if cmdPrefix == "" {
		eidFlagName = "eid"
	} else {
		eidFlagName = fmt.Sprintf("%v.eid", cmdPrefix)
	}

	var eidFlagDefault int64

	_ = cmd.PersistentFlags().Int64(eidFlagName, eidFlagDefault, eidDescription)

	return nil
}
func registerOperationRetentionListRetentionTasksIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. Retention ID.`

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
func registerOperationRetentionListRetentionTasksPageParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	pageDescription := `The page number.`

	var pageFlagName string
	if cmdPrefix == "" {
		pageFlagName = "page"
	} else {
		pageFlagName = fmt.Sprintf("%v.page", cmdPrefix)
	}

	var pageFlagDefault int64

	_ = cmd.PersistentFlags().Int64(pageFlagName, pageFlagDefault, pageDescription)

	return nil
}
func registerOperationRetentionListRetentionTasksPageSizeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	pageSizeDescription := `The size of per page.`

	var pageSizeFlagName string
	if cmdPrefix == "" {
		pageSizeFlagName = "page_size"
	} else {
		pageSizeFlagName = fmt.Sprintf("%v.page_size", cmdPrefix)
	}

	var pageSizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(pageSizeFlagName, pageSizeFlagDefault, pageSizeDescription)

	return nil
}

func retrieveOperationRetentionListRetentionTasksXRequestIDFlag(m *retention.ListRetentionTasksParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationRetentionListRetentionTasksEidFlag(m *retention.ListRetentionTasksParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("eid") {

		var eidFlagName string
		if cmdPrefix == "" {
			eidFlagName = "eid"
		} else {
			eidFlagName = fmt.Sprintf("%v.eid", cmdPrefix)
		}

		eidFlagValue, err := cmd.Flags().GetInt64(eidFlagName)
		if err != nil {
			return err, false
		}
		m.Eid = eidFlagValue

	}
	return nil, retAdded
}
func retrieveOperationRetentionListRetentionTasksIDFlag(m *retention.ListRetentionTasksParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationRetentionListRetentionTasksPageFlag(m *retention.ListRetentionTasksParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationRetentionListRetentionTasksPageSizeFlag(m *retention.ListRetentionTasksParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationRetentionListRetentionTasksResult parses request result and return the string content
func parseOperationRetentionListRetentionTasksResult(resp0 *retention.ListRetentionTasksOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*retention.ListRetentionTasksOK)
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
		resp1, ok := iResp1.(*retention.ListRetentionTasksUnauthorized)
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
		resp2, ok := iResp2.(*retention.ListRetentionTasksForbidden)
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
		resp3, ok := iResp3.(*retention.ListRetentionTasksInternalServerError)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
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
