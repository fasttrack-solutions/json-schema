package validator

const (
	GetUserDetails       = "get/user_details"
	GetUserBlocks        = "get/user_blocks"
	GetUserConsents      = "get/user_consents"
	GetBonusList         = "get/bonus_list"
	PostBonusCredit      = "post/bonus_credit"
	PostBonusCreditFunds = "post/bonus_credit_funds"
	PortUserConsents     = "post/user_consents"
)

const (
	NotificationTypeBonus              = "bonus"
	NotificationTypeCart               = "cart"
	NotificationTypeCasino             = "casino"
	NotificationTypeCustom             = "custom"
	NotificationTypeGame               = "game"
	NotificationTypeLoginV2            = "login_v2"
	NotificationTypeLotteryV2          = "lottery_v2"
	NotificationTypeLottery            = "lottery"
	NotificationTypePayment            = "payment"
	NotificationTypeSportsbook         = "sportsbook"
	NotificationTypeUserBalancesUpdate = "user_balances_update"
	NotificationTypeUserBlockV2        = "user_block_v2"
	NotificationTypeUserConsentsV2     = "user_consents_v2"
	NotificationTypeUserCreateV2       = "user_create_v2"
	NotificationTypeUserUpdateV2       = "user_update_v2"
)

func getNotificationTypes() []string {
	return []string{
		NotificationTypeBonus,
		NotificationTypeCart,
		NotificationTypeCasino,
		NotificationTypeCustom,
		NotificationTypeGame,
		NotificationTypeLoginV2,
		NotificationTypeLottery,
		NotificationTypeLotteryV2,
		NotificationTypePayment,
		NotificationTypeSportsbook,
		NotificationTypeUserBalancesUpdate,
		NotificationTypeUserBlockV2,
		NotificationTypeUserConsentsV2,
		NotificationTypeUserCreateV2,
		NotificationTypeUserUpdateV2,
	}
}

func getOperatorAPIEndpoints() []string {
	return []string{
		GetBonusList,
		GetUserBlocks,
		GetUserConsents,
		GetUserDetails,
		PortUserConsents,
		PostBonusCredit,
		PostBonusCreditFunds,
	}
}
