package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type _OrganMgr struct {
	*_BaseMgr
}

// OrganMgr open func
func OrganMgr(db *gorm.DB) *_OrganMgr {
	if db == nil {
		panic(fmt.Errorf("OrganMgr need init by db"))
	}
	return &_OrganMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrganMgr) GetTableName() string {
	return "organ"
}

// Get 获取
func (obj *_OrganMgr) Get() (result Organ, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info []User //
			err = obj.DB.New().Table("user").Where("sex = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.UserList = info
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_OrganMgr) Gets() (results []*Organ, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info []User //
				err = obj.DB.New().Table("user").Where("sex = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserList = info
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_OrganMgr) WithID(ID int) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithUserID userId获取
func (obj *_OrganMgr) WithUserID(UserID int) Option {
	return optionFunc(func(o *options) { o.query["userId"] = UserID })
}

// WithType type获取
func (obj *_OrganMgr) WithType(Type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = Type })
}

// WithScore score获取
func (obj *_OrganMgr) WithScore(Score int) Option {
	return optionFunc(func(o *options) { o.query["score"] = Score })
}

// GetByOption 功能选项模式获取
func (obj *_OrganMgr) GetByOption(opts ...Option) (result Organ, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info []User //
			err = obj.DB.New().Table("user").Where("sex = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.UserList = info
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_OrganMgr) GetByOptions(opts ...Option) (results []*Organ, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info []User //
				err = obj.DB.New().Table("user").Where("sex = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserList = info
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_OrganMgr) GetFromID(ID int) (result Organ, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info []User //
			err = obj.DB.New().Table("user").Where("sex = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.UserList = info
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_OrganMgr) GetBatchFromID(IDs []int) (results []*Organ, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info []User //
				err = obj.DB.New().Table("user").Where("sex = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserList = info
			}
		}
	}
	return
}

// GetFromUserID 通过userId获取内容
func (obj *_OrganMgr) GetFromUserID(UserID int) (results []*Organ, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("userId = ?", UserID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info []User //
				err = obj.DB.New().Table("user").Where("sex = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserList = info
			}
		}
	}
	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_OrganMgr) GetBatchFromUserID(UserIDs []int) (results []*Organ, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("userId IN (?)", UserIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info []User //
				err = obj.DB.New().Table("user").Where("sex = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserList = info
			}
		}
	}
	return
}

// GetFromType 通过type获取内容
func (obj *_OrganMgr) GetFromType(Type int) (results []*Organ, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", Type).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info []User //
				err = obj.DB.New().Table("user").Where("sex = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserList = info
			}
		}
	}
	return
}

// GetBatchFromType 批量唯一主键查找
func (obj *_OrganMgr) GetBatchFromType(Types []int) (results []*Organ, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type IN (?)", Types).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info []User //
				err = obj.DB.New().Table("user").Where("sex = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserList = info
			}
		}
	}
	return
}

// GetFromScore 通过score获取内容
func (obj *_OrganMgr) GetFromScore(Score int) (results []*Organ, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("score = ?", Score).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info []User //
				err = obj.DB.New().Table("user").Where("sex = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserList = info
			}
		}
	}
	return
}

// GetBatchFromScore 批量唯一主键查找
func (obj *_OrganMgr) GetBatchFromScore(Scores []int) (results []*Organ, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("score IN (?)", Scores).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info []User //
				err = obj.DB.New().Table("user").Where("sex = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserList = info
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrganMgr) FetchByPrimaryKey(ID int) (result Organ, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info []User //
			err = obj.DB.New().Table("user").Where("sex = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.UserList = info
		}
	}

	return
}

// FetchByIndex  获取多个内容
func (obj *_OrganMgr) FetchByIndex(UserID int) (results []*Organ, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("userId = ?", UserID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info []User //
				err = obj.DB.New().Table("user").Where("sex = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserList = info
			}
		}
	}
	return
}
