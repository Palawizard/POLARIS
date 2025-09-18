package utils

// Player holds all runtime data for the current character.
type Player struct {
	Name                  string            // Display name
	Class                 string            // Chosen class ID
	EXP                   float64           // Current experience
	EXPToNextLevel        float64           // Threshold for next level
	Level                 int               // Current level
	Money                 float64           // Coins on hand
	MaxHealth             float64           // Max HP (includes gear bonuses)
	Health                float64           // Current HP
	Skills                map[string]int    // Known skills and counts (e.g., spellbooks/uses)
	Inventory             map[string]int    // Item ID -> quantity
	InventoryMax          int               // Inventory capacity
	InventoryUpgradesUsed int               // How many capacity upgrades were bought
	Equipment             map[string]int    // Owned equipment ID -> quantity
	Equipped              map[string]string // Slot -> equipment ID (e.g., "Head": "Leather Cap")
	Initiative            float64           // Used to decide turn order
	Mana                  float64           // Current MP
	MaxMana               float64           // Max MP
	ManaRegen             float64           // MP gained per turn
}
