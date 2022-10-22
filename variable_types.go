package rlbuilder

type VariableType string

const (
	// Email
	VariableTypeEmail VariableType = "EMAIL_ADDRESS"

	// IP address
	VariableTypeIP VariableType = "IP_ADDRESS"

	// Phone number
	VariableTypePhone VariableType = "PHONE_NUMBER"

	// Browser and device
	VariableTypeUserAgent   VariableType = "USER_AGENT"
	VariableTypeFingerprint VariableType = "FINGERPRINT"

	// Payment method
	VariableTypePaymentType VariableType = "PAYMENT_TYPE"
	VariableTypeCardBin     VariableType = "CARD_BIN"
	VariableTypeAuthCode    VariableType = "AUTH_CODE"
	VariableTypeAVS         VariableType = "AVS"

	// Billing address
	VariableTypeBillingName         VariableType = "BILLING_NAME"
	VariableTypeBillingPhone        VariableType = "BILLING_PHONE"
	VariableTypeBillingAddressLine1 VariableType = "BILLING_ADDRESS_L1"
	VariableTypeBillingAddressLine2 VariableType = "BILLING_ADDRESS_L2"
	VariableTypeBillingCity         VariableType = "BILLING_CITY"
	VariableTypeBillingState        VariableType = "BILLING_STATE"
	VariableTypeBillingCountry      VariableType = "BILLING_COUNTRY"
	VariableTypeBillingZip          VariableType = "BILLING_ZIP"

	// Shipping address
	VariableTypeShippingName         VariableType = "SHIPPING_NAME"
	VariableTypeShippingPhone        VariableType = "SHIPPING_PHONE"
	VariableTypeShippingAddressLine1 VariableType = "SHIPPING_ADDRESS_L1"
	VariableTypeShippingAddressLine2 VariableType = "SHIPPING_ADDRESS_L2"
	VariableTypeShippingCity         VariableType = "SHIPPING_CITY"
	VariableTypeShippingState        VariableType = "SHIPPING_STATE"
	VariableTypeShippingCountry      VariableType = "SHIPPING_COUNTRY"
	VariableTypeShippingZip          VariableType = "SHIPPING_ZIP"

	// Order
	VariableTypeOrderID         VariableType = "ORDER_ID"
	VariableTypeProductCategory VariableType = "PRODUCT_CATEGORY"
	VariableTypeCurrencyCode    VariableType = "CURRENCY_CODE"
	VariableTypePrice           VariableType = "PRICE"

	// Custom
	VariableTypeNumeric      VariableType = "NUMERIC"
	VariableTypeCategorical  VariableType = "CATEGORICAL"
	VariableTypeFreeFormText VariableType = "FREE_FORM_TEXT"
)
