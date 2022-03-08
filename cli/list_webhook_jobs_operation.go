// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/webhookjob"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationWebhookjobListWebhookJobsCmd returns a cmd to handle operation listWebhookJobs
func makeOperationWebhookjobListWebhookJobsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "ListWebhookJobs",
		Short: `This endpoint returns webhook jobs of a project.
`,
		RunE: runOperationWebhookjobListWebhookJobs,
	}

	if err := registerOperationWebhookjobListWebhookJobsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationWebhookjobListWebhookJobs uses cmd flags to call endpoint api
func runOperationWebhookjobListWebhookJobs(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := webhookjob.NewListWebhookJobsParams()
	if err, _ := retrieveOperationWebhookjobListWebhookJobsXIsResourceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationWebhookjobListWebhookJobsXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationWebhookjobListWebhookJobsPageFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationWebhookjobListWebhookJobsPageSizeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationWebhookjobListWebhookJobsPolicyIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationWebhookjobListWebhookJobsProjectNameOrIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationWebhookjobListWebhookJobsQFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationWebhookjobListWebhookJobsSortFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationWebhookjobListWebhookJobsStatusFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationWebhookjobListWebhookJobsResult(appCli.Webhookjob.ListWebhookJobs(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationWebhookjobListWebhookJobsParamFlags registers all flags needed to fill params
func registerOperationWebhookjobListWebhookJobsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationWebhookjobListWebhookJobsXIsResourceNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationWebhookjobListWebhookJobsXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationWebhookjobListWebhookJobsPageParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationWebhookjobListWebhookJobsPageSizeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationWebhookjobListWebhookJobsPolicyIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationWebhookjobListWebhookJobsProjectNameOrIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationWebhookjobListWebhookJobsQParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationWebhookjobListWebhookJobsSortParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationWebhookjobListWebhookJobsStatusParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationWebhookjobListWebhookJobsXIsResourceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationWebhookjobListWebhookJobsXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationWebhookjobListWebhookJobsPageParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationWebhookjobListWebhookJobsPageSizeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationWebhookjobListWebhookJobsPolicyIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	policyIdDescription := `Required. The policy ID.`

	var policyIdFlagName string
	if cmdPrefix == "" {
		policyIdFlagName = "policy_id"
	} else {
		policyIdFlagName = fmt.Sprintf("%v.policy_id", cmdPrefix)
	}

	var policyIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(policyIdFlagName, policyIdFlagDefault, policyIdDescription)

	return nil
}
func registerOperationWebhookjobListWebhookJobsProjectNameOrIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationWebhookjobListWebhookJobsQParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationWebhookjobListWebhookJobsSortParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationWebhookjobListWebhookJobsStatusParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	statusDescription := `The status of webhook job.`

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "status"
	} else {
		statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
	}

	var statusFlagDefault []string

	_ = cmd.PersistentFlags().StringSlice(statusFlagName, statusFlagDefault, statusDescription)

	return nil
}

func retrieveOperationWebhookjobListWebhookJobsXIsResourceNameFlag(m *webhookjob.ListWebhookJobsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationWebhookjobListWebhookJobsXRequestIDFlag(m *webhookjob.ListWebhookJobsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationWebhookjobListWebhookJobsPageFlag(m *webhookjob.ListWebhookJobsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationWebhookjobListWebhookJobsPageSizeFlag(m *webhookjob.ListWebhookJobsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationWebhookjobListWebhookJobsPolicyIDFlag(m *webhookjob.ListWebhookJobsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("policy_id") {

		var policyIdFlagName string
		if cmdPrefix == "" {
			policyIdFlagName = "policy_id"
		} else {
			policyIdFlagName = fmt.Sprintf("%v.policy_id", cmdPrefix)
		}

		policyIdFlagValue, err := cmd.Flags().GetInt64(policyIdFlagName)
		if err != nil {
			return err, false
		}
		m.PolicyID = policyIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationWebhookjobListWebhookJobsProjectNameOrIDFlag(m *webhookjob.ListWebhookJobsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationWebhookjobListWebhookJobsQFlag(m *webhookjob.ListWebhookJobsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationWebhookjobListWebhookJobsSortFlag(m *webhookjob.ListWebhookJobsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationWebhookjobListWebhookJobsStatusFlag(m *webhookjob.ListWebhookJobsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("status") {

		var statusFlagName string
		if cmdPrefix == "" {
			statusFlagName = "status"
		} else {
			statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
		}

		statusFlagValues, err := cmd.Flags().GetStringSlice(statusFlagName)
		if err != nil {
			return err, false
		}
		m.Status = &statusFlagValues

	}
	return nil, retAdded
}

// parseOperationWebhookjobListWebhookJobsResult parses request result and return the string content
func parseOperationWebhookjobListWebhookJobsResult(resp0 *webhookjob.ListWebhookJobsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*webhookjob.ListWebhookJobsOK)
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
		resp1, ok := iResp1.(*webhookjob.ListWebhookJobsBadRequest)
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
		resp2, ok := iResp2.(*webhookjob.ListWebhookJobsUnauthorized)
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
		resp3, ok := iResp3.(*webhookjob.ListWebhookJobsForbidden)
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
		resp4, ok := iResp4.(*webhookjob.ListWebhookJobsInternalServerError)
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

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
