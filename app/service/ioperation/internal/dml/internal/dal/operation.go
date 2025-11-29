package dal

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/ioperation/operation_def"
	"donkey-admin/apperror"
	"donkey-admin/conf"
	"errors"

	"gorm.io/gorm"
)

/*  */

type operationDal struct{}

var (
	OperationDal = &operationDal{}
)

func (d *operationDal) GetSessionByModel(m *model.Operation) *gorm.DB {
	session := d.newSession()

	if m.Id != 0 {
		session.Where("`id` = ?", m.Id)
	}

	if m.SystemId != 0 {
		session.Where("`system_id` = ?", m.SystemId)
	}

	if m.MenuId != 0 {
		session.Where("`menu_id` = ?", m.MenuId)
	}

	if m.Code != "" {
		session.Where("`code` = ?", m.Code)
	}

	if m.Name != "" {
		session.Where("`name` = ?", m.Name)
	}

	if m.Status != 0 {
		session.Where("`status` = ?", m.Status)
	}

	return session
}

func (d *operationDal) GetSessionByForm(f *operation_def.OperationQueryForm) *gorm.DB {
	session := d.GetSessionByModel(&f.Operation)

	if len(f.IdList) > 0 {
		session.Where("id in (?)", f.IdList)
	}

	return session
}

func (d *operationDal) Count(f *operation_def.OperationQueryForm) (int64, error) {
	session := d.GetSessionByForm(f)

	var total int64

	if err := session.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *operationDal) Query(f *operation_def.OperationQueryForm) (*operation_def.OperationQueryRes, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var total int64
	var list []model.Operation

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&list).Error; err != nil {
		return nil, err
	}

	return &operation_def.OperationQueryRes{Total: total, List: list}, nil
}

func (d *operationDal) GetList(f *operation_def.OperationQueryForm) ([]model.Operation, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var results []model.Operation

	err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *operationDal) GetAll() ([]model.Operation, error) {
	var results []model.Operation

	session := d.newSession()
	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *operationDal) Get(pk int) (*model.Operation, error) {
	info := &model.Operation{}
	session := d.newSession()
	if err := session.Where("`id`= ?", pk).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *operationDal) GetBy(m *model.Operation) (*model.Operation, error) {
	session := d.GetSessionByModel(m)

	info := &model.Operation{}

	if err := session.First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *operationDal) GetLisByPkList(pkList []int) ([]model.Operation, error) {
	var results []model.Operation

	session := d.newSession()
	query := session.Where("`id` IN ?", pkList)
	err := query.Find(&results).Error
	return results, err
}

func (d *operationDal) GetMulti(pkList []int) (map[int]model.Operation, error) {
	if len(pkList) == 0 {
		return nil, apperror.PkListEmpty
	}

	var mMap = make(map[int]model.Operation)

	results, err := d.GetLisByPkList(pkList)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		mMap[v.Id] = v
	}
	return mMap, nil
}

func (d *operationDal) Insert(m *model.Operation) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *operationDal) BatchInsert(bm []*model.Operation, batchSize int) (int64, error) {
	session := d.newSession()
	err := session.CreateInBatches(bm, batchSize).Error
	return session.RowsAffected, err
}

func (d *operationDal) Update(m *model.Operation) (int64, error) {
	session := d.newSession()
	err := session.Updates(m).Error
	return session.RowsAffected, err
}

func (d *operationDal) UpdateBy(f *model.Operation, data map[string]any) (int64, error) {
	session := d.newSession()

	// TODO where clause

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

// SetInfo 允许 0 和 "" 值，优先使用 Update
func (d *operationDal) SetInfo(data map[string]any) (int64, error) {
	session := d.newSession()

	if _, ok := data["id"]; !ok {
		return 0, errors.New("lost pk")
	}

	// where clause
	session.Where("`id` = ?", data["id"]).Omit("id")

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

func (d *operationDal) Delete(pk int) error {
	session := d.newSession()
	err := session.Where("`id` = ?", pk).Delete(model.Operation{}).Error
	return err
}

func (d *operationDal) Exec(sql string, values ...interface{}) error {
	session := d.newSession()
	err := session.Exec(sql, values...).Error
	return err
}

func (d *operationDal) RawGet(sql string) (*model.Operation, error) {
	info := &model.Operation{}
	session := d.newSession()
	if err := session.Raw(sql).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *operationDal) RawFind(sql string) ([]model.Operation, error) {
	var results []model.Operation
	session := d.newSession()
	err := session.Raw(sql).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (d *operationDal) newEngine() *gorm.DB {
	return conf.Config.MysqlDefault.GetDb()
}

func (d *operationDal) newSession() *gorm.DB {
	return conf.Config.MysqlDefault.GetSession().Table("t_operation")
}
