package dal

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/iorganization/organization_def"
	"donkey-admin/apperror"
	"donkey-admin/conf"
	"errors"
	"gorm.io/gorm"
)

/*  */

type organizationDal struct{}

var (
	OrganizationDal = &organizationDal{}
)

func (d *organizationDal) GetSessionByModel(m *model.Organization) *gorm.DB {
	session := d.newSession()

	if m.Id != 0 {
		session.Where("`id` = ?", m.Id)
	}

	if m.ParentId != 0 {
		session.Where("`parent_id` = ?", m.ParentId)
	}

	if m.Name != "" {
		session.Where("`name` = ?", m.Name)
	}

	if m.Level != 0 {
		session.Where("`level` = ?", m.Level)
	}

	if m.NodePath != "" {
		session.Where("`node_path` = ?", m.NodePath)
	}

	if m.Status != 0 {
		session.Where("`status` = ?", m.Status)
	}

	return session
}

func (d *organizationDal) GetSessionByForm(f *organization_def.OrganizationQueryForm) *gorm.DB {
	session := d.GetSessionByModel(&f.Organization)

	if len(f.IdList) > 0 {
		session.Where("id in (?)", f.IdList)
	}

	return session
}

func (d *organizationDal) Count(f *organization_def.OrganizationQueryForm) (int64, error) {
	session := d.GetSessionByForm(f)

	var total int64

	if err := session.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *organizationDal) Query(f *organization_def.OrganizationQueryForm) (*organization_def.OrganizationQueryRes, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var total int64
	var list []model.Organization

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&list).Error; err != nil {
		return nil, err
	}

	return &organization_def.OrganizationQueryRes{Total: total, List: list}, nil
}

func (d *organizationDal) GetList(f *organization_def.OrganizationQueryForm) ([]model.Organization, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var results []model.Organization

	err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *organizationDal) GetAll() ([]model.Organization, error) {
	var results []model.Organization

	session := d.newSession()
	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *organizationDal) Get(pk int) (*model.Organization, error) {
	info := &model.Organization{}
	session := d.newSession()
	if err := session.Where("`id`= ?", pk).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *organizationDal) GetBy(m *model.Organization) (*model.Organization, error) {
	session := d.GetSessionByModel(m)

	info := &model.Organization{}

	if err := session.First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *organizationDal) GetLisByPkList(pkList []int) ([]model.Organization, error) {
	var results []model.Organization

	session := d.newSession()
	query := session.Where("`id` IN ?", pkList)
	err := query.Find(&results).Error
	return results, err
}

func (d *organizationDal) GetMulti(pkList []int) (map[int]model.Organization, error) {
	if len(pkList) == 0 {
		return nil, apperror.PkListEmpty
	}

	var mMap = make(map[int]model.Organization)

	results, err := d.GetLisByPkList(pkList)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		mMap[v.Id] = v
	}
	return mMap, nil
}

func (d *organizationDal) Insert(m *model.Organization) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *organizationDal) BatchInsert(bm []*model.Organization, batchSize int) (int64, error) {
	session := d.newSession()
	err := session.CreateInBatches(bm, batchSize).Error
	return session.RowsAffected, err
}

func (d *organizationDal) Update(m *model.Organization) (int64, error) {
	session := d.newSession()
	err := session.Updates(m).Error
	return session.RowsAffected, err
}

func (d *organizationDal) UpdateBy(f *model.Organization, data map[string]any) (int64, error) {
	// where clause
	session := d.GetSessionByModel(f)

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

// SetInfo 允许 0 和 "" 值，优先使用 Update
func (d *organizationDal) SetInfo(data map[string]any) (int64, error) {
	session := d.newSession()

	if _, ok := data["id"]; !ok {
		return 0, errors.New("lost pk id")
	}

	// where clause
	session.Where("`id` = ?", data["id"]).Omit("id")

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

func (d *organizationDal) Delete(pk int) error {
	session := d.newSession()
	err := session.Where("`id` = ?", pk).Delete(model.Organization{}).Error
	return err
}

func (d *organizationDal) Exec(sql string, values ...interface{}) error {
	session := d.newSession()
	err := session.Exec(sql, values...).Error
	return err
}

func (d *organizationDal) RawGet(sql string) (*model.Organization, error) {
	info := &model.Organization{}
	session := d.newSession()
	if err := session.Raw(sql).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *organizationDal) RawFind(sql string) ([]model.Organization, error) {
	var results []model.Organization
	session := d.newSession()
	err := session.Raw(sql).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (d *organizationDal) newEngine() *gorm.DB {
	return conf.Config.MysqlDefault.GetDb()
}

func (d *organizationDal) newSession() *gorm.DB {
	return conf.Config.MysqlDefault.GetSession().Table("t_organization")
}
