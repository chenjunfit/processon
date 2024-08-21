// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"processon/internal/app/server/dao/internal"
)

// internalDesiredResultDao is internal type for wrapping internal DAO implements.
type internalDesiredResultDao = *internal.DesiredResultDao

// desiredResultDao is the data access object for table desired_result.
// You can define custom methods on it to extend its functionality as you wish.
type desiredResultDao struct {
	internalDesiredResultDao
}

var (
	// DesiredResult is globally public accessible object for table desired_result operations.
	DesiredResult = desiredResultDao{
		internal.NewDesiredResultDao(),
	}
)

// Fill with you ideas below.