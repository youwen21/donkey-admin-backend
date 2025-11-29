package dal

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/iuser_role/user_role_def"
	"donkey-admin/apperror"
	"donkey-admin/conf"

	"gorm.io/gorm"
)

/*  */

type userRoleDal struct{}

var (
	UserRoleDal = &userRoleDal{}
)

func (d *userRoleDal) GetSessionByModel(m *model.UserRole) *gorm.DB {
	session := d.newSession()

	if m.Id != 0 {
		session.Where("id = ?", m.Id)
	}

	return session
}

func (d *userRoleDal) GetSessionByForm(f *user_role_def.UserRoleQueryForm) *gorm.DB {
	session := d.GetSessionByModel(&f.UserRole)

	if len(f.IdList) > 0 {
		session.Where("id in (?)", f.IdList)
	}

	return session
}

func (d *userRoleDal) Count(f *user_role_def.UserRoleQueryForm) (int64, error) {
	session := d.GetSessionByForm(f)

	var total int64

	if err := session.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *userRoleDal) Query(f *user_role_def.UserRoleQueryForm) (*user_role_def.UserRoleQueryRes, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var total int64
	var list []model.UserRole

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&list).Error; err != nil {
		return nil, err
	}

	return &user_role_def.UserRoleQueryRes{Total: total, List: list}, nil
}

func (d *userRoleDal) GetList(f *user_role_def.UserRoleQueryForm) ([]model.UserRole, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var results []model.UserRole

	err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *userRoleDal) GetAll() ([]model.UserRole, error) {
	var results []model.UserRole

	session := d.newSession()
	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *userRoleDal) Get(pk int) (*model.UserRole, error) {
	info := &model.UserRole{}
	session := d.newSession()
	if err := session.Where("`id`= ?", pk).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *userRoleDal) GetBy(m *model.UserRole) (*model.UserRole, error) {
	session := d.GetSessionByModel(m)

	info := &model.UserRole{}

	if err := session.First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *userRoleDal) GetLisByPkList(pkList []int) ([]model.UserRole, error) {
	var results []model.UserRole

	session := d.newSession()
	query := session.Where("`id` IN ?", pkList)
	err := query.Find(&results).Error
	return results, err
}

func (d *userRoleDal) GetMulti(pkList []int) (map[int]model.UserRole, error) {
	if len(pkList) == 0 {
		return nil, apperror.PkListEmpty
	}

	var mMap = make(map[int]model.UserRole)

	results, err := d.GetLisByPkList(pkList)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		mMap[v.Id] = v
	}
	return mMap, nil
}

func (d *userRoleDal) Insert(m *model.UserRole) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *userRoleDal) BatchInsert(bm []*model.UserRole, batchSize int) (int64, error) {
	session := d.newSession()
	err := session.CreateInBatches(bm, batchSize).Error
	return session.RowsAffected, err
}

func (d *userRoleDal) Update(m *model.UserRole) (int64, error) {
	session := d.newSession()
	err := session.Updates(m).Error
	return session.RowsAffected, err
}

func (d *userRoleDal) UpdateBy(f *model.UserRole, data map[string]any) (int64, error) {
	session := d.newSession()

	// TODO where clause

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

func (d *userRoleDal) Delete(pk int) error {
	session := d.newSession()
	err := session.Where("`id` = ?", pk).Delete(model.UserRole{}).Error
	return err
}

func (d *userRoleDal) Exec(sql string, values ...interface{}) error {
	session := d.newSession()
	err := session.Exec(sql, values...).Error
	return err
}

func (d *userRoleDal) RawGet(sql string) (*model.UserRole, error) {
	info := &model.UserRole{}
	session := d.newSession()
	if err := session.Raw(sql).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *userRoleDal) RawFind(sql string) ([]model.UserRole, error) {
	var results []model.UserRole
	session := d.newSession()
	err := session.Raw(sql).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (d *userRoleDal) newEngine() *gorm.DB {
	return conf.Config.MysqlDefault.GetDb()
}

func (d *userRoleDal) newSession() *gorm.DB {
	return conf.Config.MysqlDefault.GetSession().Table("t_user_role")
}
