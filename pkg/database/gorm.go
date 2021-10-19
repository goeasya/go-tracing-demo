package database

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	mysql "go.elastic.co/apm/module/apmgormv2/driver/mysql"
	"gorm.io/gorm"
)

type DBManager struct {
	DB       *gorm.DB
	models   []ModelInterface
	initOnce sync.Once
}

func NewGormDB(dbConnInfo string) (*DBManager, error) {
	dsn := dbConnInfo + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DBManager{
		DB:       db,
		initOnce: sync.Once{},
	}, nil
}

type ModelInterface interface {
	TableName() string
}

func (m *DBManager) InitDataBase() {
	m.initModels()
	m.checkAndInitTable()
}

func (m *DBManager) initModels() {
	m.models = append(m.models, &EventModel{})
}

func (m *DBManager) checkAndInitTable() {
	m.initOnce.Do(func() {
		for _, mo := range m.models {
			if !m.DB.Migrator().HasTable(mo) {
				err := m.DB.Migrator().CreateTable(mo)
				if err != nil {
					log.Printf("auto create table %s error", mo.TableName())
				} else {
					log.Printf("auto create table %s success", mo.TableName())
				}
			}
		}
	})
}

func (m *DBManager) EventDao() EventInterface {
	return &Event{
		DB: m.DB,
	}
}

type EventModel struct {
	ID          uint32    `gorm:"column:id;primaryKey" json:"id"`
	EventID     string    `gorm:"column:event_id;uniqueIndex;size:32;not null" json:"eventId"`
	Message     string    `gorm:"column:message;size:128;" json:"message"`
	CreatedTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdatedTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (e *EventModel) TableName() string {
	return eventTableName
}

type Event struct {
	DB *gorm.DB
}

func (e *Event) NewEvent(ctx context.Context, message string) error {
	ev := new(EventModel)
	ev.Message = message
	ev.EventID = strings.Replace(uuid.Must(uuid.NewUUID()).String(), "-", "", -1)
	ev.CreatedTime = time.Now()
	ev.UpdatedTime = time.Now()
	return e.DB.WithContext(ctx).Create(&ev).Error
}

func (e *Event) GetByEventId(ctx context.Context, eventId string) (*EventModel, error) {
	var ev *EventModel
	err := e.DB.WithContext(ctx).Where("event_id = ?", eventId).Find(&ev).Error
	return ev, err
}

func (e *Event) List(ctx context.Context) ([]*EventModel, error) {
	var eventList []*EventModel
	err := e.DB.WithContext(ctx).Find(&eventList).Error
	return eventList, err
}

