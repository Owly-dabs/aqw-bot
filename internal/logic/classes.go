package logic

import (
	"aqw/internal/window"
	"time"
)

type ActionStep struct {
	Key   string
	Delay time.Duration
}

var DefaultSteps = []ActionStep{
	{"5", 2 * time.Second},
	{"4", 2 * time.Second},
	{"3", 2 * time.Second},
	{"2", 2 * time.Second},
	{"2", 2 * time.Second},
	{"2", 2 * time.Second},
}

type FarmConfig struct {
	AttackDuration time.Duration
	RestDuration   time.Duration
	Cooldown       time.Duration
	AttackFunc     func(*window.Window) error
}

var DefaultCfg = FarmConfig{
	AttackDuration: 100 * time.Second,
	RestDuration:   30 * time.Second,
	Cooldown:       3 * time.Second,
	AttackFunc:     NinjaAttack,
}

var ContinuousCfg = FarmConfig{
	AttackDuration: 100 * time.Second,
	RestDuration:   0 * time.Second,
	Cooldown:       3 * time.Second,
	AttackFunc:     NinjaAttack,
}

// TODO: Add checking that total delay is 12 seconds

var WarriorSteps = []ActionStep{
	{"4", 2 * time.Second},
	{"5", 2 * time.Second},
	{"3", 2 * time.Second},
	{"2", 2 * time.Second},
	{"2", 2 * time.Second},
	{"2", 2 * time.Second},
}

var WarriorCfg = FarmConfig{
	AttackDuration: 100 * time.Second,
	RestDuration:   30 * time.Second,
	Cooldown:       3 * time.Second,
	AttackFunc:     WarriorAttack,
}

var HealerSteps = []ActionStep{
	{"5", 2 * time.Second},
	{"4", 2 * time.Second},
	{"3", 2 * time.Second},
	{"2", 2 * time.Second},
	{"2", 4 * time.Second},
}

var HealerCfg = FarmConfig{
	AttackDuration: 30 * time.Second,
	RestDuration:   0 * time.Second,
	Cooldown:       3 * time.Second,
	AttackFunc:     HealerAttack,
}

var DragonslayerSteps = []ActionStep{
	{"5", 2 * time.Second},
	{"4", 2 * time.Second},
	{"3", 2 * time.Second},
	{"2", 6 * time.Second},
}

var NinjaSteps = []ActionStep{
	{"5", 2 * time.Second},
	{"4", 2 * time.Second},
	{"3", 2 * time.Second},
	{"2", 2 * time.Second},
	{"2", 2 * time.Second},
	{"2", 2 * time.Second},
}
