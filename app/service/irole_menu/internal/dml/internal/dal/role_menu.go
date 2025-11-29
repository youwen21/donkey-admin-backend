package dal

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/irole_menu/role_menu_def"
	"donkey-admin/apperror"
	"donkey-admin/conf"

	"gorm.io/gorm"
)

/*  */

type roleMenuDal struct{}

var (
	RoleMenuDal = &roleMenuDal{}
)

func (d *roleMenuDal) GetSessionByModel(m *model.RoleMenu) *gorm.DB {
	session := d.newSession()

	if m.Id != 0 {
		session.Where("id = ?", m.Id)
	}

	if m.RoleId != 0 {
		session.Where("role_id = ?", m.RoleId)
	}

	if m.MenuId != 0 {
		session.Where("menu_id = ?", m.MenuId)
	}

	if m.SystemId != 0 {
		session.Where("system_id = ?", m.SystemId)
	}

	return session
}

func (d *roleMenuDal) GetSessionByForm(f *role_menu_def.RoleMenuQueryForm) *gorm.DB {
	session := d.GetSessionByModel(&f.RoleMenu)

	if len(f.IdList) > 0 {
		session.Where("id in (?)", f.IdList)
	}

	if len(f.RoleIdList) > 0 {
		session.Where("role_id in (?)", f.RoleIdList)
	}

	// if len(f.MenuIdList) > 0 {
	// 	session.Where("menu_id in (?)", f.MenuIdList)
	// }

	// if len(f.SystemIdList) > 0 {
	// 	session.Where("system_id in (?)", f.SystemIdList)
	// }

	return session
}

func (d *roleMenuDal) Count(f *role_menu_def.RoleMenuQueryForm) (int64, error) {
	session := d.GetSessionByForm(f)

	var total int64

	if err := session.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *roleMenuDal) Query(f *role_menu_def.RoleMenuQueryForm) (*role_menu_def.RoleMenuQueryRes, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var total int64
	var list []model.RoleMenu

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&list).Error; err != nil {
		return nil, err
	}

	return &role_menu_def.RoleMenuQueryRes{Total: total, List: list}, nil
}

func (d *roleMenuDal) GetList(f *role_menu_def.RoleMenuQueryForm) ([]model.RoleMenu, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var results []model.RoleMenu

	err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *roleMenuDal) GetAll() ([]model.RoleMenu, error) {
	var results []model.RoleMenu

	session := d.newSession()
	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *roleMenuDal) Get(pk int) (*model.RoleMenu, error) {
	info := &model.RoleMenu{}
	session := d.newSession()
	if err := session.Where("`id`= ?", pk).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *roleMenuDal) GetBy(m *model.RoleMenu) (*model.RoleMenu, error) {
	session := d.GetSessionByModel(m)

	info := &model.RoleMenu{}

	if err := session.First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *roleMenuDal) GetLisByPkList(pkList []int) ([]model.RoleMenu, error) {
	var results []model.RoleMenu

	session := d.newSession()
	query := session.Where("`id` IN ?", pkList)
	err := query.Find(&results).Error
	return results, err
}

func (d *roleMenuDal) GetMulti(pkList []int) (map[int]model.RoleMenu, error) {
	if len(pkList) == 0 {
		return nil, apperror.PkListEmpty
	}

	var mMap = make(map[int]model.RoleMenu)

	results, err := d.GetLisByPkList(pkList)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		mMap[v.Id] = v
	}
	return mMap, nil
}

func (d *roleMenuDal) Insert(m *model.RoleMenu) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *roleMenuDal) BatchInsert(bm []*model.RoleMenu, batchSize int) (int64, error) {
	session := d.newSession()
	err := session.CreateInBatches(bm, batchSize).Error
	return session.RowsAffected, err
}

func (d *roleMenuDal) Update(m *model.RoleMenu) (int64, error) {
	session := d.newSession()
	err := session.Updates(m).Error
	return session.RowsAffected, err
}

func (d *roleMenuDal) UpdateBy(f *model.RoleMenu, data map[string]any) (int64, error) {
	session := d.newSession()

	// TODO where clause

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

func (d *roleMenuDal) Delete(pk int) error {
	session := d.newSession()
	err := session.Where("`id` = ?", pk).Delete(model.RoleMenu{}).Error
	return err
}

func (d *roleMenuDal) Exec(sql string, values ...interface{}) error {
	session := d.newSession()
	err := session.Exec(sql, values...).Error
	return err
}

func (d *roleMenuDal) RawGet(sql string) (*model.RoleMenu, error) {
	info := &model.RoleMenu{}
	session := d.newSession()
	if err := session.Raw(sql).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *roleMenuDal) RawFind(sql string) ([]model.RoleMenu, error) {
	var results []model.RoleMenu
	session := d.newSession()
	err := session.Raw(sql).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (d *roleMenuDal) newEngine() *gorm.DB {
	return conf.Config.MysqlDefault.GetDb()
}

func (d *roleMenuDal) newSession() *gorm.DB {
	return conf.Config.MysqlDefault.GetSession().Table("t_role_menu")
}
