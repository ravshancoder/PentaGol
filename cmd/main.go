package main

import (
	"fmt"

	"github.com/PentaGol/api"
	"github.com/PentaGol/config"
	"github.com/PentaGol/pkg/logger"
	"github.com/casbin/casbin/util"
	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	gormadapter "github.com/casbin/gorm-adapter/v2"
)

func main() {
	cfg := config.Load(".")
	log := logger.New(cfg.LogLevel, "pentagol")
	
	fmt.Println(cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database)
	log.Info("main: sqlxConfig",
		logger.String("host", cfg.Postgres.Host),
		logger.String("port", cfg.Postgres.Port),
		logger.String("database", cfg.Postgres.Database),
	)
	psqlString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	a, err := gormadapter.NewAdapter("postgres", psqlString, true)
	if err != nil {
		fmt.Println(err)
		log.Error("new  adapter error", logger.Error(err))
		return
	}

	casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, a)
	if err != nil {
		fmt.Println(err)
		log.Error("new enforcer error", logger.Error(err))
		return
	}
	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		fmt.Println(err)
		log.Error("casbin error load policy", logger.Error(err))
		return
	}

	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		CasbinEnforcer: casbinEnforcer,
	})

	if err := server.Run(cfg.HttpPort); err != nil {
		fmt.Println(err)
		log.Fatal("failed to run HTTP server: ", logger.Error(err))
		panic(err)
	}

}
