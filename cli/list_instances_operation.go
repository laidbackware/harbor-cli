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

// makeOperationPreheatListInstancesCmd returns a cmd to handle operation listInstances
func makeOperationPreheatListInstancesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ListInstances",
		Short: `List P2P provider instances`,
		RunE:  runOperationPreheatListInstances,
	}

	if err := registerOperationPreheatListInstancesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPreheatListInstances uses cmd flags to call endpoint api
func runOperationPreheatListInstances(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := preheat.NewListInstancesParams()
	if err, _ := retrieveOperationPreheatListInstancesXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPreheatListInstancesPageFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPreheatListInstancesPageSizeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPreheatListInstancesQFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPreheatListInstancesSortFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationPreheatListInstancesResult(appCli.Preheat.ListInstances(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationPreheatListInstancesParamFlags registers all flags needed to fill params
func registerOperationPreheatListInstancesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPreheatListInstancesXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPreheatListInstancesPageParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPreheatListInstancesPageSizeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPreheatListInstancesQParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPreheatListInstancesSortParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPreheatListInstancesXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationPreheatListInstancesPageParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationPreheatListInstancesPageSizeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationPreheatListInstancesQParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	qDescription := `Query string to query resources. Supported query patterns are 'exact match(k=v)', 'fuzzy match(k=~v)', 'range(k=[min~max])', 'list with union releationship(k={v1 v2 v3})' and 'list with intersetion relationship(k=(v1 v2 v3))'. The value of range and list can be string(enclosed by ' or ), integer or time(in format '2020-04-09 02:36:00'). All of these query patterns should be put in the query string 'q=xxx' and splitted by ','. e.g. q=k1=v1,k2=~v2,k3=[min~max]`

	var qFlagName string
	if cmdPrefix == "" {
		qFlagName = "q"
	} else {
		qFlagName = fmt.Sprintf("%v.q", cmdPrefix)
	}

	var qFlagDefault string

	_ = cmd.PersistentFlags().String(qFlagName, qFlagDefault, qDescription)

	return nil
}
func registerOperationPreheatListInstancesSortParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sortDescription := `Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with 'sort=field1,-field2'`

	var sortFlagName string
	if cmdPrefix == "" {
		sortFlagName = "sort"
	} else {
		sortFlagName = fmt.Sprintf("%v.sort", cmdPrefix)
	}

	var sortFlagDefault string

	_ = cmd.PersistentFlags().String(sortFlagName, sortFlagDefault, sortDescription)

	return nil
}

func retrieveOperationPreheatListInstancesXRequestIDFlag(m *preheat.ListInstancesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationPreheatListInstancesPageFlag(m *preheat.ListInstancesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationPreheatListInstancesPageSizeFlag(m *preheat.ListInstancesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationPreheatListInstancesQFlag(m *preheat.ListInstancesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("q") {

		var qFlagName string
		if cmdPrefix == "" {
			qFlagName = "q"
		} else {
			qFlagName = fmt.Sprintf("%v.q", cmdPrefix)
		}

		qFlagValue, err := cmd.Flags().GetString(qFlagName)
		if err != nil {
			return err, false
		}
		m.Q = &qFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPreheatListInstancesSortFlag(m *preheat.ListInstancesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("sort") {

		var sortFlagName string
		if cmdPrefix == "" {
			sortFlagName = "sort"
		} else {
			sortFlagName = fmt.Sprintf("%v.sort", cmdPrefix)
		}

		sortFlagValue, err := cmd.Flags().GetString(sortFlagName)
		if err != nil {
			return err, false
		}
		m.Sort = &sortFlagValue

	}
	return nil, retAdded
}

// parseOperationPreheatListInstancesResult parses request result and return the string content
func parseOperationPreheatListInstancesResult(resp0 *preheat.ListInstancesOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*preheat.ListInstancesOK)
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
		resp1, ok := iResp1.(*preheat.ListInstancesBadRequest)
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
		resp2, ok := iResp2.(*preheat.ListInstancesUnauthorized)
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
		resp3, ok := iResp3.(*preheat.ListInstancesForbidden)
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
		resp4, ok := iResp4.(*preheat.ListInstancesNotFound)
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
		resp5, ok := iResp5.(*preheat.ListInstancesInternalServerError)
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