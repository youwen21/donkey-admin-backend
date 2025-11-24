package dal

import (
	"errors"
	"gofly/app/model"
	"gofly/app/service/iuser_permissions/user_permissions_def"
	"gofly/apperror"
	"gofly/conf"

	"gorm.io/gorm"
)

/*  */

type userPermissionsDal struct{}

var (
	UserPermissionsDal = &userPermissionsDal{}
)

func (d *userPermissionsDal) GetSessionByModel(m *model.UserPermissions) *gorm.DB {
	session := d.newSession()

	if m.Id != 0 {
		session.Where("id = ?", m.Id)
	}

	if m.UserId != 0 {
		session.Where("user_id = ?", m.UserId)
	}

	if m.SystemId != 0 {
		session.Where("system_id = ?", m.SystemId)
	}

	if m.MenuId != 0 {
		session.Where("menu_id = ?", m.MenuId)
	}

	return session
}

func (d *userPermissionsDal) GetSessionByForm(f *user_permissions_def.UserPermissionsQueryForm) *gorm.DB {
	session := d.GetSessionByModel(&f.UserPermissions)

	if len(f.IdList) > 0 {
		session.Where("id in (?)", f.IdList)
	}

	return session
}

func (d *userPermissionsDal) Count(f *user_permissions_def.UserPermissionsQueryForm) (int64, error) {
	session := d.GetSessionByForm(f)

	var total int64

	if err := session.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *userPermissionsDal) Query(f *user_permissions_def.UserPermissionsQueryForm) (*user_permissions_def.UserPermissionsQueryRes, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var total int64
	var list []model.UserPermissions

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&list).Error; err != nil {
		return nil, err
	}

	return &user_permissions_def.UserPermissionsQueryRes{Total: total, List: list}, nil
}

func (d *userPermissionsDal) GetList(f *user_permissions_def.UserPermissionsQueryForm) ([]model.UserPermissions, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var results []model.UserPermissions

	err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *userPermissionsDal) GetAll() ([]model.UserPermissions, error) {
	var results []model.UserPermissions

	session := d.newSession()
	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *userPermissionsDal) Get(pk int) (*model.UserPermissions, error) {
	info := &model.UserPermissions{}
	session := d.newSession()
	if err := session.Where("`id`= ?", pk).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *userPermissionsDal) GetBy(m *model.UserPermissions) (*model.UserPermissions, error) {
	session := d.GetSessionByModel(m)

	info := &model.UserPermissions{}

	if err := session.First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *userPermissionsDal) GetLisByPkList(pkList []int) ([]model.UserPermissions, error) {
	var results []model.UserPermissions

	session := d.newSession()
	query := session.Where("`id` IN ?", pkList)
	err := query.Find(&results).Error
	return results, err
}

func (d *userPermissionsDal) GetMulti(pkList []int) (map[int]model.UserPermissions, error) {
	if len(pkList) == 0 {
		return nil, apperror.PkListEmpty
	}

	var mMap = make(map[int]model.UserPermissions)

	results, err := d.GetLisByPkList(pkList)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		mMap[v.Id] = v
	}
	return mMap, nil
}

func (d *userPermissionsDal) Insert(m *model.UserPermissions) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *userPermissionsDal) BatchInsert(bm []*model.UserPermissions, batchSize int) (int64, error) {
	session := d.newSession()
	err := session.CreateInBatches(bm, batchSize).Error
	return session.RowsAffected, err
}

func (d *userPermissionsDal) Update(m *model.UserPermissions) (int64, error) {
	session := d.newSession()
	err := session.Updates(m).Error
	return session.RowsAffected, err
}

func (d *userPermissionsDal) UpdateBy(f *model.UserPermissions, data map[string]any) (int64, error) {
	// where clause
	session := d.GetSessionByModel(f)

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

// SetInfo 允许 0 和 "" 值，优先使用 Update
func (d *userPermissionsDal) SetInfo(data map[string]any) (int64, error) {
	session := d.newSession()

	if _, ok := data["id"]; !ok {
		return 0, errors.New("lost pk id")
	}

	// where clause
	session.Where("`id` = ?", data["id"]).Omit("id")

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

func (d *userPermissionsDal) Delete(pk int) error {
	session := d.newSession()
	err := session.Where("`id` = ?", pk).Delete(model.UserPermissions{}).Error
	return err
}

func (d *userPermissionsDal) Exec(sql string, values ...interface{}) (int64, error) {
	session := d.newSession()
	session.Exec(sql, values...)
	return session.RowsAffected, session.Error
}

func (d *userPermissionsDal) RawGet(sql string) (*model.UserPermissions, error) {
	info := &model.UserPermissions{}
	session := d.newSession()
	if err := session.Raw(sql).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *userPermissionsDal) RawFind(sql string) ([]model.UserPermissions, error) {
	var results []model.UserPermissions
	session := d.newSession()
	err := session.Raw(sql).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (d *userPermissionsDal) newEngine() *gorm.DB {
	return conf.Config.MysqlDefault.GetDb()
}

func (d *userPermissionsDal) newSession() *gorm.DB {
	return conf.Config.MysqlDefault.GetSession().Table("t_user_permissions")
}
