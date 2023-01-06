/*
                       _
   ___   _   _   ___  | |_    ___    _ __ ___
  / __| | | | | / __| | __|  / _ \  | '_ ` _ \
 | (__  | |_| | \__ \ | |_  | (_) | | | | | | |
  \___|  \__,_| |___/  \__|  \___/  |_| |_| |_|

*/
// * Code generated by protoc-gen-custom. DO NOT EDIT.
// * versions:
// * - protoc-gen-custom (devel)
// * - protoc                   v3.14.0
// * source: enum_pb

package pb

func (x UserRole) FromString(str string) UserRole {
	switch str {
	case UserRole_USER_UNKNOWN.String():
		return UserRole_USER_UNKNOWN
	case UserRole_SUPER_ADMIN.String():
		return UserRole_SUPER_ADMIN
	case UserRole_ADMIN.String():
		return UserRole_ADMIN
	case UserRole_SELLER.String():
		return UserRole_SELLER
	case UserRole_MANAGER.String():
		return UserRole_MANAGER
	}
	return UserRole(0)
}

func (x ProjectStatus) FromString(str string) ProjectStatus {
	switch str {
	case ProjectStatus_ProjectStatus_NONE.String():
		return ProjectStatus_ProjectStatus_NONE
	case ProjectStatus_ProjectStatus_IN_PROGRESS.String():
		return ProjectStatus_ProjectStatus_IN_PROGRESS
	case ProjectStatus_ProjectStatus_PAUSE.String():
		return ProjectStatus_ProjectStatus_PAUSE
	case ProjectStatus_ProjectStatus_CLOSE.String():
		return ProjectStatus_ProjectStatus_CLOSE
	case ProjectStatus_ProjectStatus_DRAFT.String():
		return ProjectStatus_ProjectStatus_DRAFT
	}
	return ProjectStatus(0)
}
