package dal

import (
	"errors"
	"gofly/app/model"
	"gofly/app/service/irole/role_def"
	"gofly/apperror"
	"gofly/conf"

	"gorm.io/gorm"
)

/*  */

type roleDal struct{}

var (
	RoleDal = &roleDal{}
)

func (d *roleDal) GetSessionByModel(m *model.Role) *gorm.DB {
	session := d.newSession()

	if m.Id != 0 {
		session.Where("id = ?", m.Id)
	}

	return session
}

func (d *roleDal) GetSessionByForm(f *role_def.RoleQueryForm) *gorm.DB {
	session := d.GetSessionByModel(&f.Role)

	if len(f.IdList) > 0 {
		session.Where("id in (?)", f.IdList)
	}

	return session
}

func (d *roleDal) Count(f *role_def.RoleQueryForm) (int64, error) {
	session := d.GetSessionByForm(f)

	var total int64

	if err := session.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *roleDal) Query(f *role_def.RoleQueryForm) (*role_def.RoleQueryRes, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var total int64
	var list []model.Role

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&list).Error; err != nil {
		return nil, err
	}

	return &role_def.RoleQueryRes{Total: total, List: list}, nil
}

func (d *roleDal) GetList(f *role_def.RoleQueryForm) ([]model.Role, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var results []model.Role

	err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *roleDal) GetAll() ([]model.Role, error) {
	var results []model.Role

	session := d.newSession()
	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *roleDal) Get(pk int) (*model.Role, error) {
	info := &model.Role{}
	session := d.newSession()
	if err := session.Where("`id`= ?", pk).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *roleDal) GetBy(m *model.Role) (*model.Role, error) {
	session := d.GetSessionByModel(m)

	info := &model.Role{}

	if err := session.First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *roleDal) GetLisByPkList(pkList []int) ([]model.Role, error) {
	var results []model.Role

	session := d.newSession()
	query := session.Where("`id` IN ?", pkList)
	err := query.Find(&results).Error
	return results, err
}

func (d *roleDal) GetMulti(pkList []int) (map[int]model.Role, error) {
	if len(pkList) == 0 {
		return nil, apperror.PkListEmpty
	}

	var mMap = make(map[int]model.Role)

	results, err := d.GetLisByPkList(pkList)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		mMap[v.Id] = v
	}
	return mMap, nil
}

func (d *roleDal) Insert(m *model.Role) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *roleDal) BatchInsert(bm []*model.Role, batchSize int) (int64, error) {
	session := d.newSession()
	err := session.CreateInBatches(bm, batchSize).Error
	return session.RowsAffected, err
}

func (d *roleDal) Update(m *model.Role) (int64, error) {
	session := d.newSession()
	err := session.Updates(m).Error
	return session.RowsAffected, err
}

func (d *roleDal) UpdateBy(f *model.Role, data map[string]any) (int64, error) {
	session := d.newSession()

	// TODO where clause

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

// SetInfo 允许 0 和 "" 值，优先使用 Update
func (d *roleDal) SetInfo(data map[string]any) (int64, error) {
	session := d.newSession()

	if _, ok := data["id"]; !ok {
		return 0, errors.New("lost pk")
	}

	// where clause
	session.Where("`id` = ?", data["id"]).Omit("id")

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

func (d *roleDal) Delete(pk int) error {
	session := d.newSession()
	err := session.Where("`id` = ?", pk).Delete(model.Role{}).Error
	return err
}

func (d *roleDal) Exec(sql string, values ...interface{}) error {
	session := d.newSession()
	err := session.Exec(sql, values...).Error
	return err
}

func (d *roleDal) RawGet(sql string) (*model.Role, error) {
	info := &model.Role{}
	session := d.newSession()
	if err := session.Raw(sql).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *roleDal) RawFind(sql string) ([]model.Role, error) {
	var results []model.Role
	session := d.newSession()
	err := session.Raw(sql).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (d *roleDal) newEngine() *gorm.DB {
	return conf.Config.MysqlDefault.GetDb()
}

func (d *roleDal) newSession() *gorm.DB {
	return conf.Config.MysqlDefault.GetSession().Table("t_role")
}
