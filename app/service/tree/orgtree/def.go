package orgtree

import "gofly/app/model"

type TreeOrg struct {
	model.Organization
	Children []*TreeOrg `json:"children"`
}
