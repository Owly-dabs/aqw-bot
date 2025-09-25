package logic

import (
	"aqw/internal/window"
	"log"
	"time"
)

func ExecuteFarm(win *window.Window, cfg FarmConfig) error {
	for {
		// Attack phase
		log.Printf("🗡️ Starting attack phase (%s)", cfg.AttackDuration)
		if err := AttackRandomMonster(win); err != nil {
			return err
		}

		attackEnd := time.Now().Add(cfg.AttackDuration)
		for time.Now().Before(attackEnd) {
			if err := cfg.AttackFunc(win); err != nil {
				return err
			}
			time.Sleep(cfg.Cooldown)
		}

		// Rest phase
		log.Printf("💤 Starting rest phase (%s)", cfg.RestDuration)
		restEnd := time.Now().Add(cfg.RestDuration)
		for time.Now().Before(restEnd) {
			if err := Rest(win); err != nil {
				return err
			}
			time.Sleep(5 * time.Second) // configurable if needed
		}
	}
}

func FarmWarrior(win *window.Window) error {
	for {
		// Attack phase: 100 seconds
		log.Println("🗡️ Starting attack phase (100s)")
		attackEnd := time.Now().Add(100 * time.Second)
		for time.Now().Before(attackEnd) {
			if err := AttackRandomMonster(win); err != nil {
				return err
			}
			if err := WarriorAttack(win); err != nil {
				return err
			}
			time.Sleep(3 * time.Second) // Cooldown for the attack
			if err := WarriorAttack(win); err != nil {
				return err
			}
			time.Sleep(3 * time.Second) // Cooldown for the attack
		}

		// Rest phase: 30 seconds
		log.Println("💤 Starting rest phase (30s)")
		restEnd := time.Now().Add(30 * time.Second)
		for time.Now().Before(restEnd) {
			if err := Rest(win); err != nil {
				return err
			}
			time.Sleep(5 * time.Second) // Adjust based on how often to rest
		}
	}
}

func FarmHealer(win *window.Window) error {
	for {
		// Attack phase: 100 seconds
		log.Println("🗡️ Starting attack phase (100s)")
		attackEnd := time.Now().Add(100 * time.Second)
		if err := AttackRandomMonster(win); err != nil {
			return err
		}
		for time.Now().Before(attackEnd) {
			if err := HealerAttack(win); err != nil {
				return err
			}
			time.Sleep(3 * time.Second) // Cooldown for the attack
		}

		// // Rest phase: 30 seconds
		// log.Println("💤 Starting rest phase (45s)")
		// restEnd := time.Now().Add(45 * time.Second)
		// for time.Now().Before(restEnd) {
		// 	if err := Rest(win); err != nil {
		// 		return err
		// 	}
		// 	time.Sleep(5 * time.Second) // Adjust based on how often to rest
		// }
	}
}
