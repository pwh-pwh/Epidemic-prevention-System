package service

import (
	"context"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/mysql"
	"github.com/pwh-pwh/Epidemic-prevention-System/dao/query"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
)

func GetAllDept(flag string) ([]*models.SysDept, error) {
	sysDeptQ := query.Use(mysql.DB).SysDept
	deptList, err := sysDeptQ.WithContext(context.Background()).Where().Find()
	if err != nil {
		return nil, err
	}
	return toTree(deptList, 0), nil
}

func findChildren(parent *models.SysDept, treeList []*models.SysDept) *models.SysDept {
	for _, children := range treeList {
		if parent.DeptID == children.ParentID {
			parent.Children = append(parent.Children, findChildren(children, treeList))
		}
	}
	return parent
}
func toTree(treeList []*models.SysDept, pid int64) []*models.SysDept {
	var returnList []*models.SysDept
	for _, parent := range treeList {
		if pid == parent.ParentID {
			returnList = append(returnList, findChildren(parent, treeList))
		}
	}
	return returnList
}
