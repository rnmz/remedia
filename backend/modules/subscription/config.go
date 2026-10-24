package subscription

type SubscriptionConfig struct {
	Permissions SubscriptionPermissions `json:"permissions"` // Permissions for subscriptions
}

type SubscriptionPermissions struct {
	ChangePlan  []string `json:"change_plan"`  // Permissions to change subscription plans
	ChangePrice []string `json:"change_price"` // Permissions to change subscription prices
}
