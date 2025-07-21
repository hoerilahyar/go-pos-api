package casbin

import (
	"fmt"
	"log"
	"time"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var Enforcer *casbin.SyncedEnforcer

func InitCasbin(db *gorm.DB) *casbin.SyncedEnforcer {
	// Membuat adapter dari GORM DB
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("Failed to create adapter: %v", err)
	}

	// Gunakan SyncedEnforcer, bukan Enforcer biasa
	enforcer, err := casbin.NewSyncedEnforcer("pkg/casbin/policies/rbac_model.conf", adapter)
	if err != nil {
		log.Fatalf("Failed to create synced enforcer: %v", err)
	}

	// Aktifkan auto-save policy ke DB
	// enforcer.EnableAutoSave(true)

	// Mulai auto-reload policy dari DB setiap 1 menit
	enforcer.StartAutoLoadPolicy(time.Second * 1)

	// Load policy dari DB pertama kali
	if err := enforcer.LoadPolicy(); err != nil {
		log.Fatalf("Failed to load policy: %v", err)
	}

	// if err := policies.SetupPolicies(enforcer); err != nil {
	// 	log.Fatalf("failed to setup policies: %v", err)
	// }

	// Simpan ke variabel global jika ingin digunakan di package lain
	Enforcer = enforcer

	fmt.Println("Casbin initialized")

	return enforcer
}
