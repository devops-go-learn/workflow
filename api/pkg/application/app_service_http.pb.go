// Code generated by protoc-gen-go-http. DO NOT EDIT.

package application

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				Path:              "/workflow.application.Service/CreateApplication",
				FunctionName:      "CreateApplication",
				Method:            "POST",
				Resource:          "application",
				AuthEnable:        true,
				PermissionEnable:  true,
				AuditLog:          false,
				RequiredNamespace: false,
				Labels:            map[string]string{"action": "create"},
				Extension:         map[string]string{},
			},
			{
				Path:              "/workflow.application.Service/QueryApplication",
				FunctionName:      "QueryApplication",
				Method:            "GET",
				Resource:          "application",
				AuthEnable:        false,
				PermissionEnable:  false,
				AuditLog:          false,
				RequiredNamespace: false,
				Labels:            map[string]string{},
				Extension:         map[string]string{},
			},
		},
	}
	return set
}
