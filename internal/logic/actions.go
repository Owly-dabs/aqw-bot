package logic

import (
	"aqw/internal/input"
	"aqw/internal/window"
	"log"
	"time"
)

func AttackRandomMonster(win *window.Window) error {
	log.Println("Attempting to attack a random monster...")
	if err := win.Activate(); err != nil {
		return err
	}
	if err := input.SendText(win.ID, "t"); err != nil {
		return err
	}
	return input.SendText(win.ID, "1")
}

func ExecuteAttack(win *window.Window, name string, steps []ActionStep) error {
	log.Printf("Executing %s attack...\n", name)
	if err := win.Activate(); err != nil {
		return err
	}

	for _, step := range steps {
		if err := input.SendText(win.ID, step.Key); err != nil {
			return err
		}
		time.Sleep(step.Delay)
	}
	return nil
}

func Rest(win *window.Window) error {
	log.Println("Resting...")
	if err := win.Activate(); err != nil {
		return err
	}
	return input.SendText(win.ID, "x")
}

func JoinLocation(win *window.Window, location string) error {
	if err := win.Activate(); err != nil {
		return err
	}
	if err := input.SendEnter(win.ID); err != nil {
		return err
	}
	if err := input.SendText(win.ID, "/join "+location); err != nil {
		return err
	}
	return input.SendEnter(win.ID)
}

func WarriorAttack(win *window.Window) error {
	return ExecuteAttack(win, "warrior", WarriorSteps)
}

func HealerAttack(win *window.Window) error {
	return ExecuteAttack(win, "healer", HealerSteps)
}

func DragonslayerAttack(win *window.Window) error {
	return ExecuteAttack(win, "dragonslayer", DragonslayerSteps)
}

func NinjaAttack(win *window.Window) error {
	return ExecuteAttack(win, "ninja", NinjaSteps)
}
