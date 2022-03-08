// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for UserCreationReq

// register flags to command
func registerModelUserCreationReqFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUserCreationReqComment(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserCreationReqEmail(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserCreationReqPassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserCreationReqRealname(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserCreationReqUsername(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUserCreationReqComment(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	commentDescription := ``

	var commentFlagName string
	if cmdPrefix == "" {
		commentFlagName = "comment"
	} else {
		commentFlagName = fmt.Sprintf("%v.comment", cmdPrefix)
	}

	var commentFlagDefault string

	_ = cmd.PersistentFlags().String(commentFlagName, commentFlagDefault, commentDescription)

	return nil
}

func registerUserCreationReqEmail(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	emailDescription := ``

	var emailFlagName string
	if cmdPrefix == "" {
		emailFlagName = "email"
	} else {
		emailFlagName = fmt.Sprintf("%v.email", cmdPrefix)
	}

	var emailFlagDefault string

	_ = cmd.PersistentFlags().String(emailFlagName, emailFlagDefault, emailDescription)

	return nil
}

func registerUserCreationReqPassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	passwordDescription := ``

	var passwordFlagName string
	if cmdPrefix == "" {
		passwordFlagName = "password"
	} else {
		passwordFlagName = fmt.Sprintf("%v.password", cmdPrefix)
	}

	var passwordFlagDefault string

	_ = cmd.PersistentFlags().String(passwordFlagName, passwordFlagDefault, passwordDescription)

	return nil
}

func registerUserCreationReqRealname(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	realnameDescription := ``

	var realnameFlagName string
	if cmdPrefix == "" {
		realnameFlagName = "realname"
	} else {
		realnameFlagName = fmt.Sprintf("%v.realname", cmdPrefix)
	}

	var realnameFlagDefault string

	_ = cmd.PersistentFlags().String(realnameFlagName, realnameFlagDefault, realnameDescription)

	return nil
}

func registerUserCreationReqUsername(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	usernameDescription := ``

	var usernameFlagName string
	if cmdPrefix == "" {
		usernameFlagName = "username"
	} else {
		usernameFlagName = fmt.Sprintf("%v.username", cmdPrefix)
	}

	var usernameFlagDefault string

	_ = cmd.PersistentFlags().String(usernameFlagName, usernameFlagDefault, usernameDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUserCreationReqFlags(depth int, m *models.UserCreationReq, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, commentAdded := retrieveUserCreationReqCommentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || commentAdded

	err, emailAdded := retrieveUserCreationReqEmailFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || emailAdded

	err, passwordAdded := retrieveUserCreationReqPasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || passwordAdded

	err, realnameAdded := retrieveUserCreationReqRealnameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || realnameAdded

	err, usernameAdded := retrieveUserCreationReqUsernameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || usernameAdded

	return nil, retAdded
}

func retrieveUserCreationReqCommentFlags(depth int, m *models.UserCreationReq, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	commentFlagName := fmt.Sprintf("%v.comment", cmdPrefix)
	if cmd.Flags().Changed(commentFlagName) {

		var commentFlagName string
		if cmdPrefix == "" {
			commentFlagName = "comment"
		} else {
			commentFlagName = fmt.Sprintf("%v.comment", cmdPrefix)
		}

		commentFlagValue, err := cmd.Flags().GetString(commentFlagName)
		if err != nil {
			return err, false
		}
		m.Comment = commentFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserCreationReqEmailFlags(depth int, m *models.UserCreationReq, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	emailFlagName := fmt.Sprintf("%v.email", cmdPrefix)
	if cmd.Flags().Changed(emailFlagName) {

		var emailFlagName string
		if cmdPrefix == "" {
			emailFlagName = "email"
		} else {
			emailFlagName = fmt.Sprintf("%v.email", cmdPrefix)
		}

		emailFlagValue, err := cmd.Flags().GetString(emailFlagName)
		if err != nil {
			return err, false
		}
		m.Email = emailFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserCreationReqPasswordFlags(depth int, m *models.UserCreationReq, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	passwordFlagName := fmt.Sprintf("%v.password", cmdPrefix)
	if cmd.Flags().Changed(passwordFlagName) {

		var passwordFlagName string
		if cmdPrefix == "" {
			passwordFlagName = "password"
		} else {
			passwordFlagName = fmt.Sprintf("%v.password", cmdPrefix)
		}

		passwordFlagValue, err := cmd.Flags().GetString(passwordFlagName)
		if err != nil {
			return err, false
		}
		m.Password = passwordFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserCreationReqRealnameFlags(depth int, m *models.UserCreationReq, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	realnameFlagName := fmt.Sprintf("%v.realname", cmdPrefix)
	if cmd.Flags().Changed(realnameFlagName) {

		var realnameFlagName string
		if cmdPrefix == "" {
			realnameFlagName = "realname"
		} else {
			realnameFlagName = fmt.Sprintf("%v.realname", cmdPrefix)
		}

		realnameFlagValue, err := cmd.Flags().GetString(realnameFlagName)
		if err != nil {
			return err, false
		}
		m.Realname = realnameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserCreationReqUsernameFlags(depth int, m *models.UserCreationReq, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	usernameFlagName := fmt.Sprintf("%v.username", cmdPrefix)
	if cmd.Flags().Changed(usernameFlagName) {

		var usernameFlagName string
		if cmdPrefix == "" {
			usernameFlagName = "username"
		} else {
			usernameFlagName = fmt.Sprintf("%v.username", cmdPrefix)
		}

		usernameFlagValue, err := cmd.Flags().GetString(usernameFlagName)
		if err != nil {
			return err, false
		}
		m.Username = usernameFlagValue

		retAdded = true
	}

	return nil, retAdded
}
