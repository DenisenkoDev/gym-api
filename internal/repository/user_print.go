package repository

import (
	"fmt"
	"gym-api/ent"
)

func PrintUserInfo(viz *ent.User, err error) {
	fmt.Println("=================================================")
	if err != nil {
		fmt.Println("❗❗❗  Ошибка:", err)
		return
	}
	fmt.Printf("👤 Посетитель: ID=%d | %s %s\n", viz.ID, viz.FirstName, viz.LastName)
	fmt.Println()

	for _, gym := range viz.Edges.VisitorGym {
		fmt.Println("🏋️ Зал:", gym.ID, "-", gym.Name)

		if len(gym.Edges.UserRoles) > 0 {
			fmt.Println("  👥 Роли пользователя:")
			for _, role := range gym.Edges.UserRoles {
				fmt.Printf("    - ID: %d | Роль: %s\n", role.ID, role.UserRole)
			}
		}
		fmt.Println()
	}

	for _, gym := range viz.Edges.ManagerGym {
		fmt.Println("🏋️ Зал:", gym.ID, "-", gym.Name)

		if len(gym.Edges.UserRoles) > 0 {
			fmt.Println("  👥 Роли пользователя:")
			for _, role := range gym.Edges.UserRoles {
				fmt.Printf("    - ID: %d | Роль: %s\n", role.ID, role.UserRole)
			}
		}
		fmt.Println()
	}

	for _, gym := range viz.Edges.OwnedGyms {
		fmt.Println("🏋️ Зал:", gym.ID, "-", gym.Name)

		if len(gym.Edges.UserRoles) > 0 {
			fmt.Println("  👥 Роли пользователя:")
			for _, role := range gym.Edges.UserRoles {
				fmt.Printf("    - ID: %d | Роль: %s\n", role.ID, role.UserRole)
			}
		}
		fmt.Println()
	}

	fmt.Println("=================================================")
}

func PrintVisitorInfo(viz *ent.User, err error) {
	fmt.Println("=================================================")
	if err != nil {
		fmt.Println("❗❗❗  Ошибка:", err)
		return
	}
	fmt.Printf("👤 Посетитель: ID=%d | %s %s\n", viz.ID, viz.FirstName, viz.LastName)
	fmt.Println()

	for _, gym := range viz.Edges.VisitorGym {
		fmt.Println("🏋️ Зал:", gym.ID, "-", gym.Name)

		if len(gym.Edges.UserRoles) > 0 {
			fmt.Println("  👥 Роли пользователя:")
			for _, role := range gym.Edges.UserRoles {
				fmt.Printf("    - ID: %d | Роль: %s\n", role.ID, role.UserRole)
			}
		}

		if len(gym.Edges.ManagerRoles) > 0 {
			fmt.Println("  🧑‍💼 Роли менеджера:")
			for _, role := range gym.Edges.ManagerRoles {
				fmt.Printf("    - ID: %d | Роль: %s\n", role.ID, role.ManagerRole)
			}
		}

		if len(gym.Edges.Abonements) > 0 {
			fmt.Println("  📄 Абонементы:")
			for _, ab := range gym.Edges.Abonements {
				fmt.Printf("    - ID: %d | Name: %s | Цена: %.2f\n", ab.ID, ab.Name, ab.Price)
				if ab.Edges.Coach != nil {
					fmt.Printf("      🧑‍🏫 Тренер: %s %s\n", ab.Edges.Coach.FirstName, ab.Edges.Coach.LastName)
				}
				if len(ab.Edges.Payments) > 0 {
					fmt.Println("      💰 Платежи:")
					for _, payment := range ab.Edges.Payments {
						fmt.Printf("        - ID: %d | Сумма: %.2f\n", payment.ID, payment.Amount)
					}
				}
			}
		}
		fmt.Println()
	}
	fmt.Println("=================================================")
}

func PrintOwnerInfo(owner *ent.User, err error) {
	fmt.Println("=================================================")
	if err != nil {
		fmt.Println("❗❗❗  Ошибка:", err)
		return
	}
	fmt.Printf("🏠 Владелец: ID=%d | %s %s\n", owner.ID, owner.FirstName, owner.LastName)
	fmt.Println()

	for _, gym := range owner.Edges.OwnedGyms {
		fmt.Println("🏋️ Зал:", gym.ID, "-", gym.Name)

		if len(gym.Edges.Managers) > 0 {
			fmt.Println("  🧑‍💼 Менеджеры:")
			for _, manager := range gym.Edges.Managers {
				fmt.Printf("    - ID: %d | %s %s\n", manager.ID, manager.FirstName, manager.LastName)
				for _, role := range manager.Edges.UserRoles {
					fmt.Printf("      👔 Роль: %s (ID: %d)\n", role.UserRole, role.ID)
				}
			}
		}

		if len(gym.Edges.Visitors) > 0 {
			fmt.Println("  👥 Посетители:")
			for _, visitor := range gym.Edges.Visitors {
				fmt.Printf("    - ID: %d | %s %s\n", visitor.ID, visitor.FirstName, visitor.LastName)
				for _, role := range visitor.Edges.UserRoles {
					fmt.Printf("      🪪 Роль: %s (ID: %d)\n", role.UserRole, role.ID)
				}
				if len(visitor.Edges.Abonements) > 0 {
					fmt.Println("      📄 Абонементы:")
					for _, ab := range visitor.Edges.Abonements {
						fmt.Printf("        - ID: %d | Цена: %.2f\n", ab.ID, ab.Price)
						if ab.Edges.Type != nil {
							fmt.Println("          📌 Тип:", ab.Edges.Type.Name)
						}
						if len(ab.Edges.Payments) > 0 {
							fmt.Println("          💰 Платежи:")
							for _, payment := range ab.Edges.Payments {
								fmt.Printf("            - ID: %d | Сумма: %.2f\n", payment.ID, payment.Amount)
							}
						}
					}
				}
			}
		}
		fmt.Println()
	}
	fmt.Println("=================================================")
}

func PrintManagerInfo(manager *ent.User, err error) {
	fmt.Println("=================================================")
	if err != nil {
		fmt.Println("❗❗❗  Ошибка:", err)
		return
	}
	fmt.Printf("🧑‍💼 Менеджер: ID=%d | %s %s\n", manager.ID, manager.FirstName, manager.LastName)
	fmt.Println()

	for _, gym := range manager.Edges.ManagerGym {
		fmt.Println("🏋️ Зал:", gym.ID, "-", gym.Name)

		if len(gym.Edges.Managers) > 0 {
			fmt.Println("  🧑‍💼 Менеджеры:")
			for _, m := range gym.Edges.Managers {
				fmt.Printf("    - ID: %d | %s %s\n", m.ID, m.FirstName, m.LastName)
				for _, role := range m.Edges.UserRoles {
					fmt.Printf("      👔 Роль: %s (ID: %d)\n", role.UserRole, role.ID)
				}
			}
		}

		if len(gym.Edges.Visitors) > 0 {
			fmt.Println("  👥 Посетители:")
			for _, visitor := range gym.Edges.Visitors {
				fmt.Printf("    - ID: %d | %s %s\n", visitor.ID, visitor.FirstName, visitor.LastName)
				for _, role := range visitor.Edges.UserRoles {
					fmt.Printf("      🪪 Роль: %s (ID: %d)\n", role.UserRole, role.ID)
				}
				if len(visitor.Edges.Abonements) > 0 {
					fmt.Println("      📄 Абонементы:")
					for _, ab := range visitor.Edges.Abonements {
						fmt.Printf("        - ID: %d | Цена: %.2f\n", ab.ID, ab.Price)
						if ab.Edges.Type != nil {
							fmt.Println("          📌 Тип:", ab.Edges.Type.Name)
						}
						if len(ab.Edges.Payments) > 0 {
							fmt.Println("          💰 Платежи:")
							for _, payment := range ab.Edges.Payments {
								fmt.Printf("            - ID: %d | Сумма: %.2f\n", payment.ID, payment.Amount)
							}
						}
					}
				}
			}
		}
		fmt.Println()
	}
	fmt.Println("=================================================")
}
