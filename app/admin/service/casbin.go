package service

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	log "github.com/sirupsen/logrus"
	"time"
	"vab-admin/go/pkg/config"
	"vab-admin/go/pkg/db"
	"vab-admin/go/pkg/errors"
	model2 "vab-admin/go/pkg/model"
)

type Enforcer struct {
	*casbin.SyncedEnforcer
}

type CasbinAdapter struct{}

// LoadPolicy
// @date: 2022-02-01 18:50:10
func (c *CasbinAdapter) LoadPolicy(model model.Model) error {
	log.Debug("LoadPolicy")

	ctx := context.Background()
	err := c.loadRolePolicy(ctx, model)
	if err != nil {
		log.WithError(err).Error("Load casbin role policy")
		return err
	}

	err = c.loadUserPolicy(ctx, model)
	if err != nil {
		log.WithError(err).Error("Load casbin user policy")
		return err
	}

	return nil
}

func (c *CasbinAdapter) SavePolicy(_ model.Model) error { return nil }

// AddPolicy
// @date: 2022-02-01 18:50:03
func (c *CasbinAdapter) AddPolicy(_ string, _ string, _ []string) error { return nil }

// RemovePolicy
// @date: 2022-02-01 18:50:06
func (c *CasbinAdapter) RemovePolicy(sec string, _ string, rule []string) error { return nil }

// RemoveFilteredPolicy
// @date: 2022-02-01 18:50:05
func (c *CasbinAdapter) RemoveFilteredPolicy(sec string, _ string, fieldIndex int, fieldValues ...string) error {
	return nil
}

// Load user policy (g,user_id,role_id)
func (c *CasbinAdapter) loadUserPolicy(ctx context.Context, m model.Model) error {
	log.Debug("loadUserPolicy")

	var rows []*model2.AdminUserRole
	tx := db.Session(ctx).Model(&model2.AdminUserRole{}).Select([]string{"user_id", "role_id"}).Find(&rows)
	if err := tx.Error; err != nil {
		log.WithError(err).Error("获取用户角色失败")
		return err
	}

	for _, row := range rows {
		line := fmt.Sprintf("g,%d,%d", row.UserID, row.RoleID)
		_ = persist.LoadPolicyLine(line, m)
	}

	return nil
}

// loadRolePolicy Load role policy (p,role_id,path,method)
func (c *CasbinAdapter) loadRolePolicy(ctx context.Context, m model.Model) error {

	var groups []*model2.AdminRole

	db.Session(ctx).Model(&model2.AdminRole{}).Preload("Rules").Find(&groups)

	for _, group := range groups {
		for _, rule := range group.Rules {
			for _, action := range rule.Apis {

				line := fmt.Sprintf("p,%d,%s,%s", group.GetId(), action.GetPath(), action.GetMethod())
				_ = persist.LoadPolicyLine(line, m)
			}
		}
	}

	return nil
}

// NewCasbin
// @date: 2022-02-01 18:49:42
func NewCasbin(conf config.Config) (*Enforcer, error) {

	e, err := casbin.NewSyncedEnforcer(conf.Casbin.Model)
	if err != nil {
		return nil, err
	}

	e.EnableLog(conf.Casbin.Log)

	err = e.InitWithModelAndAdapter(e.GetModel(), &CasbinAdapter{})
	if err != nil {
		log.WithError(err).Fatal("创建casbin失败")
		return nil, err
	}

	e.EnableEnforce(true)

	e.StartAutoLoadPolicy(time.Hour * 24)

	return &Enforcer{e}, nil
}

// LoadPolicy
// @date: 2022-02-01 18:49:41
func (e *Enforcer) LoadPolicy() error {

	if err := e.Enforcer.LoadPolicy(); err != nil {
		return errors.ErrInternalServer
	}
	return nil
}
