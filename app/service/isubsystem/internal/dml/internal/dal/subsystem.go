package dal

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/isubsystem/subsystem_def"
	"donkey-admin/apperror"
	"donkey-admin/conf"
	"errors"

	"gorm.io/gorm"
)

/*  */

type subsystemDal struct{}

var (
	SubsystemDal = &subsystemDal{}
)

func (d *subsystemDal) GetSessionByModel(m *model.Subsystem) *gorm.DB {
	session := d.newSession()

	if m.Id != 0 {
		session.Where("id = ?", m.Id)
	}

	return session
}

func (d *subsystemDal) GetSessionByForm(f *subsystem_def.SubsystemQueryForm) *gorm.DB {
	session := d.GetSessionByModel(&f.Subsystem)

	if len(f.IdList) > 0 {
		session.Where("id in (?)", f.IdList)
	}

	return session
}

func (d *subsystemDal) Count(f *subsystem_def.SubsystemQueryForm) (int64, error) {
	session := d.GetSessionByForm(f)

	var total int64

	if err := session.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *subsystemDal) Query(f *subsystem_def.SubsystemQueryForm) (*subsystem_def.SubsystemQueryRes, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var total int64
	var list []model.Subsystem

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&list).Error; err != nil {
		return nil, err
	}

	return &subsystem_def.SubsystemQueryRes{Total: total, List: list}, nil
}

func (d *subsystemDal) GetList(f *subsystem_def.SubsystemQueryForm) ([]model.Subsystem, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var results []model.Subsystem

	err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *subsystemDal) GetAll() ([]model.Subsystem, error) {
	var results []model.Subsystem

	session := d.newSession()
	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *subsystemDal) Get(pk int) (*model.Subsystem, error) {
	info := &model.Subsystem{}
	session := d.newSession()
	if err := session.Where("`id`= ?", pk).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *subsystemDal) GetBy(m *model.Subsystem) (*model.Subsystem, error) {
	session := d.GetSessionByModel(m)

	info := &model.Subsystem{}

	if err := session.First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *subsystemDal) GetLisByPkList(pkList []int) ([]model.Subsystem, error) {
	var results []model.Subsystem

	session := d.newSession()
	query := session.Where("`id` IN ?", pkList)
	err := query.Find(&results).Error
	return results, err
}

func (d *subsystemDal) GetMulti(pkList []int) (map[int]model.Subsystem, error) {
	if len(pkList) == 0 {
		return nil, apperror.PkListEmpty
	}

	var mMap = make(map[int]model.Subsystem)

	results, err := d.GetLisByPkList(pkList)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		mMap[v.Id] = v
	}
	return mMap, nil
}

func (d *subsystemDal) Insert(m *model.Subsystem) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *subsystemDal) BatchInsert(bm []*model.Subsystem, batchSize int) (int64, error) {
	session := d.newSession()
	err := session.CreateInBatches(bm, batchSize).Error
	return session.RowsAffected, err
}

func (d *subsystemDal) Update(m *model.Subsystem) (int64, error) {
	session := d.newSession()
	err := session.Updates(m).Error
	return session.RowsAffected, err
}

func (d *subsystemDal) UpdateBy(f *model.Subsystem, data map[string]any) (int64, error) {
	session := d.newSession()

	// TODO where clause

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

// SetInfo 允许 0 和 "" 值，优先使用 Update
func (d *subsystemDal) SetInfo(data map[string]any) (int64, error) {
	session := d.newSession()

	if _, ok := data["id"]; !ok {
		return 0, errors.New("lost pk")
	}

	// where clause
	session.Where("`id` = ?", data["id"]).Omit("id")

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

func (d *subsystemDal) Delete(pk int) error {
	session := d.newSession()
	err := session.Where("`id` = ?", pk).Delete(model.Subsystem{}).Error
	return err
}

func (d *subsystemDal) Exec(sql string, values ...interface{}) error {
	session := d.newSession()
	err := session.Exec(sql, values...).Error
	return err
}

func (d *subsystemDal) RawGet(sql string) (*model.Subsystem, error) {
	info := &model.Subsystem{}
	session := d.newSession()
	if err := session.Raw(sql).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *subsystemDal) RawFind(sql string) ([]model.Subsystem, error) {
	var results []model.Subsystem
	session := d.newSession()
	err := session.Raw(sql).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (d *subsystemDal) newEngine() *gorm.DB {
	return conf.Config.MysqlDefault.GetDb()
}

func (d *subsystemDal) newSession() *gorm.DB {
	return conf.Config.MysqlDefault.GetSession().Table("t_subsystem")
}
