package dal

import (
	"errors"
	"gofly/app/model"
	"gofly/app/service/iuser_permission/user_permission_def"
	"gofly/apperror"
	"gofly/conf"

	"gorm.io/gorm"
)

/*  */

type userPermissionDal struct{}

var (
	UserPermissionDal = &userPermissionDal{}
)

func (d *userPermissionDal) GetSessionByModel(m *model.UserPermission) *gorm.DB {
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

func (d *userPermissionDal) GetSessionByForm(f *user_permission_def.UserPermissionQueryForm) *gorm.DB {
	session := d.GetSessionByModel(&f.UserPermission)

	if len(f.IdList) > 0 {
		session.Where("id in (?)", f.IdList)
	}

	return session
}

func (d *userPermissionDal) Count(f *user_permission_def.UserPermissionQueryForm) (int64, error) {
	session := d.GetSessionByForm(f)

	var total int64

	if err := session.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *userPermissionDal) Query(f *user_permission_def.UserPermissionQueryForm) (*user_permission_def.UserPermissionQueryRes, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var total int64
	var list []model.UserPermission

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&list).Error; err != nil {
		return nil, err
	}

	return &user_permission_def.UserPermissionQueryRes{Total: total, List: list}, nil
}

func (d *userPermissionDal) GetList(f *user_permission_def.UserPermissionQueryForm) ([]model.UserPermission, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var results []model.UserPermission

	err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *userPermissionDal) GetAll() ([]model.UserPermission, error) {
	var results []model.UserPermission

	session := d.newSession()
	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *userPermissionDal) Get(pk int) (*model.UserPermission, error) {
	info := &model.UserPermission{}
	session := d.newSession()
	if err := session.Where("`id`= ?", pk).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *userPermissionDal) GetBy(m *model.UserPermission) (*model.UserPermission, error) {
	session := d.GetSessionByModel(m)

	info := &model.UserPermission{}

	if err := session.First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *userPermissionDal) GetLisByPkList(pkList []int) ([]model.UserPermission, error) {
	var results []model.UserPermission

	session := d.newSession()
	query := session.Where("`id` IN ?", pkList)
	err := query.Find(&results).Error
	return results, err
}

func (d *userPermissionDal) GetMulti(pkList []int) (map[int]model.UserPermission, error) {
	if len(pkList) == 0 {
		return nil, apperror.PkListEmpty
	}

	var mMap = make(map[int]model.UserPermission)

	results, err := d.GetLisByPkList(pkList)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		mMap[v.Id] = v
	}
	return mMap, nil
}

func (d *userPermissionDal) Insert(m *model.UserPermission) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *userPermissionDal) BatchInsert(bm []*model.UserPermission, batchSize int) (int64, error) {
	session := d.newSession()
	err := session.CreateInBatches(bm, batchSize).Error
	return session.RowsAffected, err
}

func (d *userPermissionDal) Update(m *model.UserPermission) (int64, error) {
	session := d.newSession()
	err := session.Updates(m).Error
	return session.RowsAffected, err
}

func (d *userPermissionDal) UpdateBy(f *model.UserPermission, data map[string]any) (int64, error) {
	// where clause
	session := d.GetSessionByModel(f)

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

// SetInfo 允许 0 和 "" 值，优先使用 Update
func (d *userPermissionDal) SetInfo(data map[string]any) (int64, error) {
	session := d.newSession()

	if _, ok := data["id"]; !ok {
		return 0, errors.New("lost pk id")
	}

	// where clause
	session.Where("`id` = ?", data["id"]).Omit("id")

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

func (d *userPermissionDal) Delete(pk int) error {
	session := d.newSession()
	err := session.Where("`id` = ?", pk).Delete(model.UserPermission{}).Error
	return err
}

func (d *userPermissionDal) Exec(sql string, values ...interface{}) (int64, error) {
	session := d.newSession()
	session.Exec(sql, values...)
	return session.RowsAffected, session.Error
}

func (d *userPermissionDal) RawGet(sql string) (*model.UserPermission, error) {
	info := &model.UserPermission{}
	session := d.newSession()
	if err := session.Raw(sql).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *userPermissionDal) RawFind(sql string) ([]model.UserPermission, error) {
	var results []model.UserPermission
	session := d.newSession()
	err := session.Raw(sql).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (d *userPermissionDal) newEngine() *gorm.DB {
	return conf.Config.MysqlDefault.GetDb()
}

func (d *userPermissionDal) newSession() *gorm.DB {
	return conf.Config.MysqlDefault.GetSession().Table("t_user_permission")
}
