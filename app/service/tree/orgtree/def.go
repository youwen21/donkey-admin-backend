package orgtree

import "donkey-admin/app/model"

type TreeOrg struct {
	model.Organization
	Children []*TreeOrg `json:"children"`
}
