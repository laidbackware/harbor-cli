// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/spf13/cobra"
)

// Schema cli for ProjectMember

// register flags to command
func registerModelProjectMemberFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerProjectMemberMemberGroup(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerProjectMemberMemberUser(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerProjectMemberRoleID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerProjectMemberMemberGroup(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var memberGroupFlagName string
	if cmdPrefix == "" {
		memberGroupFlagName = "member_group"
	} else {
		memberGroupFlagName = fmt.Sprintf("%v.member_group", cmdPrefix)
	}

	if err := registerModelUserGroupFlags(depth+1, memberGroupFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerProjectMemberMemberUser(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var memberUserFlagName string
	if cmdPrefix == "" {
		memberUserFlagName = "member_user"
	} else {
		memberUserFlagName = fmt.Sprintf("%v.member_user", cmdPrefix)
	}

	if err := registerModelUserEntityFlags(depth+1, memberUserFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerProjectMemberRoleID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	roleIdDescription := `The role id 1 for projectAdmin, 2 for developer, 3 for guest, 4 for maintainer`

	var roleIdFlagName string
	if cmdPrefix == "" {
		roleIdFlagName = "role_id"
	} else {
		roleIdFlagName = fmt.Sprintf("%v.role_id", cmdPrefix)
	}

	var roleIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(roleIdFlagName, roleIdFlagDefault, roleIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelProjectMemberFlags(depth int, m *models.ProjectMember, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, memberGroupAdded := retrieveProjectMemberMemberGroupFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || memberGroupAdded

	err, memberUserAdded := retrieveProjectMemberMemberUserFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || memberUserAdded

	err, roleIdAdded := retrieveProjectMemberRoleIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || roleIdAdded

	return nil, retAdded
}

func retrieveProjectMemberMemberGroupFlags(depth int, m *models.ProjectMember, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	memberGroupFlagName := fmt.Sprintf("%v.member_group", cmdPrefix)
	if cmd.Flags().Changed(memberGroupFlagName) {
		// info: complex object member_group UserGroup is retrieved outside this Changed() block
	}
	memberGroupFlagValue := m.MemberGroup
	if swag.IsZero(memberGroupFlagValue) {
		memberGroupFlagValue = &models.UserGroup{}
	}

	err, memberGroupAdded := retrieveModelUserGroupFlags(depth+1, memberGroupFlagValue, memberGroupFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || memberGroupAdded
	if memberGroupAdded {
		m.MemberGroup = memberGroupFlagValue
	}

	return nil, retAdded
}

func retrieveProjectMemberMemberUserFlags(depth int, m *models.ProjectMember, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	memberUserFlagName := fmt.Sprintf("%v.member_user", cmdPrefix)
	if cmd.Flags().Changed(memberUserFlagName) {
		// info: complex object member_user UserEntity is retrieved outside this Changed() block
	}
	memberUserFlagValue := m.MemberUser
	if swag.IsZero(memberUserFlagValue) {
		memberUserFlagValue = &models.UserEntity{}
	}

	err, memberUserAdded := retrieveModelUserEntityFlags(depth+1, memberUserFlagValue, memberUserFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || memberUserAdded
	if memberUserAdded {
		m.MemberUser = memberUserFlagValue
	}

	return nil, retAdded
}

func retrieveProjectMemberRoleIDFlags(depth int, m *models.ProjectMember, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	roleIdFlagName := fmt.Sprintf("%v.role_id", cmdPrefix)
	if cmd.Flags().Changed(roleIdFlagName) {

		var roleIdFlagName string
		if cmdPrefix == "" {
			roleIdFlagName = "role_id"
		} else {
			roleIdFlagName = fmt.Sprintf("%v.role_id", cmdPrefix)
		}

		roleIdFlagValue, err := cmd.Flags().GetInt64(roleIdFlagName)
		if err != nil {
			return err, false
		}
		m.RoleID = roleIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
