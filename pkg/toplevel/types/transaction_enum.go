// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package types

// NewTransactionStatusPhase - transalate between NW change Phase and Transaction Phase
func NewTransactionStatusPhase(value int) TransactionStatusPhase {
	switch value {
	case 0:
		return TransactionStatusPhaseTRANSACTIONCHANGE
	default:
		return TransactionStatusPhaseTRANSACTIONROLLBACK
	}
}

// NewTransactionStatusState - transalate between NW change State and Transaction State
func NewTransactionStatusState(value int) TransactionStatusState {
	switch value {
	case 0:
		return TransactionStatusStateTRANSACTIONPENDING
	case 2:
		return TransactionStatusStateTRANSACTIONCOMPLETE
	default:
		return TransactionStatusStateTRANSACTIONFAILED
	}
}
