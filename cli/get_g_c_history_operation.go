// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/gc"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationGcGetGCHistoryCmd returns a cmd to handle operation getGCHistory
func makeOperationGcGetGCHistoryCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getGCHistory",
		Short: `This endpoint let user get gc execution history.`,
		RunE:  runOperationGcGetGCHistory,
	}

	if err := registerOperationGcGetGCHistoryParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationGcGetGCHistory uses cmd flags to call endpoint api
func runOperationGcGetGCHistory(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := gc.NewGetGCHistoryParams()
	if err, _ := retrieveOperationGcGetGCHistoryXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationGcGetGCHistoryPageFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationGcGetGCHistoryPageSizeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationGcGetGCHistoryQFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationGcGetGCHistorySortFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationGcGetGCHistoryResult(appCli.Gc.GetGCHistory(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationGcGetGCHistoryParamFlags registers all flags needed to fill params
func registerOperationGcGetGCHistoryParamFlags(cmd *cobra.Command) error {
	if err := registerOperationGcGetGCHistoryXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationGcGetGCHistoryPageParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationGcGetGCHistoryPageSizeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationGcGetGCHistoryQParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationGcGetGCHistorySortParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationGcGetGCHistoryXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationGcGetGCHistoryPageParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationGcGetGCHistoryPageSizeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationGcGetGCHistoryQParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationGcGetGCHistorySortParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationGcGetGCHistoryXRequestIDFlag(m *gc.GetGCHistoryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationGcGetGCHistoryPageFlag(m *gc.GetGCHistoryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationGcGetGCHistoryPageSizeFlag(m *gc.GetGCHistoryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationGcGetGCHistoryQFlag(m *gc.GetGCHistoryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationGcGetGCHistorySortFlag(m *gc.GetGCHistoryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationGcGetGCHistoryResult parses request result and return the string content
func parseOperationGcGetGCHistoryResult(resp0 *gc.GetGCHistoryOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*gc.GetGCHistoryOK)
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
		resp1, ok := iResp1.(*gc.GetGCHistoryUnauthorized)
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
		resp2, ok := iResp2.(*gc.GetGCHistoryForbidden)
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
		resp3, ok := iResp3.(*gc.GetGCHistoryInternalServerError)
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
