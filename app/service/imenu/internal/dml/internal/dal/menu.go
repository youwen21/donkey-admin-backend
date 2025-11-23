package dal

import (
	"errors"
	"gofly/app/model"
	"gofly/app/service/imenu/menu_def"
	"gofly/apperror"
	"gofly/conf"

	"gorm.io/gorm"
)

/*  */

type menuDal struct{}

var (
	MenuDal = &menuDal{}
)

func (d *menuDal) GetSessionByModel(m *model.Menu) *gorm.DB {
	session := d.newSession()

	if m.Id != 0 {
		session.Where("id = ?", m.Id)
	}

	if m.SystemId != 0 {
		session.Where("system_id = ?", m.SystemId)
	}

	if m.ParentId != 0 {
		session.Where("parent_id = ?", m.ParentId)
	}

	if m.Name != "" {
		session.Where("name like ?", "%"+m.Name+"%")
	}
	if m.Level != 0 {
		session.Where("`level` = ?", m.Level)
	}
	if m.Status != 0 {
		session.Where("`status` = ?", m.Status)
	}

	return session
}

func (d *menuDal) GetSessionByForm(f *menu_def.MenuQueryForm) *gorm.DB {
	session := d.GetSessionByModel(&f.Menu)

	if len(f.IdList) > 0 {
		session.Where("id in (?)", f.IdList)
	}

	return session
}

func (d *menuDal) Count(f *menu_def.MenuQueryForm) (int64, error) {
	session := d.GetSessionByForm(f)

	var total int64

	if err := session.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *menuDal) Query(f *menu_def.MenuQueryForm) (*menu_def.MenuQueryRes, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var total int64
	var list []model.Menu

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&list).Error; err != nil {
		return nil, err
	}

	return &menu_def.MenuQueryRes{Total: total, List: list}, nil
}

func (d *menuDal) GetList(f *menu_def.MenuQueryForm) ([]model.Menu, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var results []model.Menu

	err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *menuDal) GetAll() ([]model.Menu, error) {
	var results []model.Menu

	session := d.newSession()
	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *menuDal) Get(pk int) (*model.Menu, error) {
	info := &model.Menu{}
	session := d.newSession()
	if err := session.Where("`id`= ?", pk).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *menuDal) GetBy(m *model.Menu) (*model.Menu, error) {
	session := d.GetSessionByModel(m)

	info := &model.Menu{}

	if err := session.First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *menuDal) GetLisByPkList(pkList []int) ([]model.Menu, error) {
	var results []model.Menu

	session := d.newSession()
	query := session.Where("`id` IN ?", pkList)
	err := query.Find(&results).Error
	return results, err
}

func (d *menuDal) GetMulti(pkList []int) (map[int]model.Menu, error) {
	if len(pkList) == 0 {
		return nil, apperror.PkListEmpty
	}

	var mMap = make(map[int]model.Menu)

	results, err := d.GetLisByPkList(pkList)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		mMap[v.Id] = v
	}
	return mMap, nil
}

func (d *menuDal) Insert(m *model.Menu) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *menuDal) BatchInsert(bm []*model.Menu, batchSize int) (int64, error) {
	session := d.newSession()
	err := session.CreateInBatches(bm, batchSize).Error
	return session.RowsAffected, err
}

func (d *menuDal) Update(m *model.Menu) (int64, error) {
	session := d.newSession()
	err := session.Updates(m).Error
	return session.RowsAffected, err
}

func (d *menuDal) UpdateBy(f *model.Menu, data map[string]any) (int64, error) {
	session := d.newSession()

	// TODO where clause

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

// SetInfo 允许 0 和 "" 值，优先使用 Update
func (d *menuDal) SetInfo(data map[string]any) (int64, error) {
	session := d.newSession()

	if _, ok := data["id"]; !ok {
		return 0, errors.New("lost pk")
	}

	// where clause
	session.Where("`id` = ?", data["id"]).Omit("id")

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

func (d *menuDal) Delete(pk int) error {
	session := d.newSession()
	err := session.Where("`id` = ?", pk).Delete(model.Menu{}).Error
	return err
}

func (d *menuDal) Exec(sql string, values ...interface{}) error {
	session := d.newSession()
	err := session.Exec(sql, values...).Error
	return err
}

func (d *menuDal) RawGet(sql string) (*model.Menu, error) {
	info := &model.Menu{}
	session := d.newSession()
	if err := session.Raw(sql).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *menuDal) RawFind(sql string) ([]model.Menu, error) {
	var results []model.Menu
	session := d.newSession()
	err := session.Raw(sql).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (d *menuDal) newEngine() *gorm.DB {
	return conf.Config.MysqlDefault.GetDb()
}

func (d *menuDal) newSession() *gorm.DB {
	return conf.Config.MysqlDefault.GetSession().Table("t_menu")
}
