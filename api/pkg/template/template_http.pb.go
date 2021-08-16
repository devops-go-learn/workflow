// Code generated by protoc-gen-go-http. DO NOT EDIT.

package template

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				Path:         "/workflow.template.Service/CreateTemplate",
				FunctionName: "CreateTemplate",
			},
			{
				Path:         "/workflow.template.Service/QueryTemplate",
				FunctionName: "QueryTemplate",
			},
			{
				Path:         "/workflow.template.Service/DescribeTemplate",
				FunctionName: "DescribeTemplate",
			},
			{
				Path:         "/workflow.template.Service/UpdateAction",
				FunctionName: "UpdateAction",
			},
			{
				Path:         "/workflow.template.Service/DeleteTemplate",
				FunctionName: "DeleteTemplate",
			},
		},
	}
	return set
}
