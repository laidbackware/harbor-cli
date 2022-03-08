// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for Search

// register flags to command
func registerModelSearchFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSearchChart(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSearchProject(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSearchRepository(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSearchChart(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: chart []*SearchResult array type is not supported by go-swagger cli yet

	return nil
}

func registerSearchProject(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: project []*Project array type is not supported by go-swagger cli yet

	return nil
}

func registerSearchRepository(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: repository []*SearchRepository array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSearchFlags(depth int, m *models.Search, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, chartAdded := retrieveSearchChartFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || chartAdded

	err, projectAdded := retrieveSearchProjectFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || projectAdded

	err, repositoryAdded := retrieveSearchRepositoryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || repositoryAdded

	return nil, retAdded
}

func retrieveSearchChartFlags(depth int, m *models.Search, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	chartFlagName := fmt.Sprintf("%v.chart", cmdPrefix)
	if cmd.Flags().Changed(chartFlagName) {
		// warning: chart array type []*SearchResult is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSearchProjectFlags(depth int, m *models.Search, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	projectFlagName := fmt.Sprintf("%v.project", cmdPrefix)
	if cmd.Flags().Changed(projectFlagName) {
		// warning: project array type []*Project is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSearchRepositoryFlags(depth int, m *models.Search, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	repositoryFlagName := fmt.Sprintf("%v.repository", cmdPrefix)
	if cmd.Flags().Changed(repositoryFlagName) {
		// warning: repository array type []*SearchRepository is not supported by go-swagger cli yet
	}

	return nil, retAdded
}