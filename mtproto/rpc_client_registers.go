/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'codegen_rpc_helper.py'
 *
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mtproto

import (
	"github.com/golang/glog"
	"reflect"
)

type newRPCReplyFunc func() interface{}

type RPCContextTuple struct {
	Method       string
	NewReplyFunc newRPCReplyFunc
}

var rpcContextRegisters = map[string]RPCContextTuple{
	"TLAuthLogOut":                         RPCContextTuple{"/mtproto.RPCAuth/auth_logOut", func() interface{} { return new(Bool) }},
	"TLAuthResetAuthorizations":            RPCContextTuple{"/mtproto.RPCAuth/auth_resetAuthorizations", func() interface{} { return new(Bool) }},
	"TLAuthSendInvites":                    RPCContextTuple{"/mtproto.RPCAuth/auth_sendInvites", func() interface{} { return new(Bool) }},
	"TLAuthBindTempAuthKey":                RPCContextTuple{"/mtproto.RPCAuth/auth_bindTempAuthKey", func() interface{} { return new(Bool) }},
	"TLAuthCancelCode":                     RPCContextTuple{"/mtproto.RPCAuth/auth_cancelCode", func() interface{} { return new(Bool) }},
	"TLAuthDropTempAuthKeys":               RPCContextTuple{"/mtproto.RPCAuth/auth_dropTempAuthKeys", func() interface{} { return new(Bool) }},
	"TLAccountRegisterDevice":              RPCContextTuple{"/mtproto.RPCAccount/account_registerDevice", func() interface{} { return new(Bool) }},
	"TLAccountRegisterDeviceLayer74":       RPCContextTuple{"/mtproto.RPCAccount/account_registerDeviceLayer74", func() interface{} { return new(Bool) }},
	"TLAccountUnregisterDevice":            RPCContextTuple{"/mtproto.RPCAccount/account_unregisterDevice", func() interface{} { return new(Bool) }},
	"TLAccountUnregisterDeviceLayer74":     RPCContextTuple{"/mtproto.RPCAccount/account_unregisterDeviceLayer74", func() interface{} { return new(Bool) }},
	"TLAccountUpdateNotifySettings":        RPCContextTuple{"/mtproto.RPCAccount/account_updateNotifySettings", func() interface{} { return new(Bool) }},
	"TLAccountResetNotifySettings":         RPCContextTuple{"/mtproto.RPCAccount/account_resetNotifySettings", func() interface{} { return new(Bool) }},
	"TLAccountUpdateStatus":                RPCContextTuple{"/mtproto.RPCAccount/account_updateStatus", func() interface{} { return new(Bool) }},
	"TLAccountReportPeer":                  RPCContextTuple{"/mtproto.RPCAccount/account_reportPeer", func() interface{} { return new(Bool) }},
	"TLAccountCheckUsername":               RPCContextTuple{"/mtproto.RPCAccount/account_checkUsername", func() interface{} { return new(Bool) }},
	"TLAccountDeleteAccount":               RPCContextTuple{"/mtproto.RPCAccount/account_deleteAccount", func() interface{} { return new(Bool) }},
	"TLAccountSetAccountTTL":               RPCContextTuple{"/mtproto.RPCAccount/account_setAccountTTL", func() interface{} { return new(Bool) }},
	"TLAccountUpdateDeviceLocked":          RPCContextTuple{"/mtproto.RPCAccount/account_updateDeviceLocked", func() interface{} { return new(Bool) }},
	"TLAccountResetAuthorization":          RPCContextTuple{"/mtproto.RPCAccount/account_resetAuthorization", func() interface{} { return new(Bool) }},
	"TLAccountUpdatePasswordSettings":      RPCContextTuple{"/mtproto.RPCAccount/account_updatePasswordSettings", func() interface{} { return new(Bool) }},
	"TLAccountConfirmPhone":                RPCContextTuple{"/mtproto.RPCAccount/account_confirmPhone", func() interface{} { return new(Bool) }},
	"TLContactsDeleteContacts":             RPCContextTuple{"/mtproto.RPCContacts/contacts_deleteContacts", func() interface{} { return new(Bool) }},
	"TLContactsBlock":                      RPCContextTuple{"/mtproto.RPCContacts/contacts_block", func() interface{} { return new(Bool) }},
	"TLContactsUnblock":                    RPCContextTuple{"/mtproto.RPCContacts/contacts_unblock", func() interface{} { return new(Bool) }},
	"TLContactsResetTopPeerRating":         RPCContextTuple{"/mtproto.RPCContacts/contacts_resetTopPeerRating", func() interface{} { return new(Bool) }},
	"TLContactsResetSaved":                 RPCContextTuple{"/mtproto.RPCContacts/contacts_resetSaved", func() interface{} { return new(Bool) }},
	"TLMessagesSetTyping":                  RPCContextTuple{"/mtproto.RPCMessages/messages_setTyping", func() interface{} { return new(Bool) }},
	"TLMessagesReportSpam":                 RPCContextTuple{"/mtproto.RPCMessages/messages_reportSpam", func() interface{} { return new(Bool) }},
	"TLMessagesHideReportSpam":             RPCContextTuple{"/mtproto.RPCMessages/messages_hideReportSpam", func() interface{} { return new(Bool) }},
	"TLMessagesDiscardEncryption":          RPCContextTuple{"/mtproto.RPCMessages/messages_discardEncryption", func() interface{} { return new(Bool) }},
	"TLMessagesSetEncryptedTyping":         RPCContextTuple{"/mtproto.RPCMessages/messages_setEncryptedTyping", func() interface{} { return new(Bool) }},
	"TLMessagesReadEncryptedHistory":       RPCContextTuple{"/mtproto.RPCMessages/messages_readEncryptedHistory", func() interface{} { return new(Bool) }},
	"TLMessagesReportEncryptedSpam":        RPCContextTuple{"/mtproto.RPCMessages/messages_reportEncryptedSpam", func() interface{} { return new(Bool) }},
	"TLMessagesUninstallStickerSet":        RPCContextTuple{"/mtproto.RPCMessages/messages_uninstallStickerSet", func() interface{} { return new(Bool) }},
	"TLMessagesEditChatAdmin":              RPCContextTuple{"/mtproto.RPCMessages/messages_editChatAdmin", func() interface{} { return new(Bool) }},
	"TLMessagesReorderStickerSets":         RPCContextTuple{"/mtproto.RPCMessages/messages_reorderStickerSets", func() interface{} { return new(Bool) }},
	"TLMessagesSaveGif":                    RPCContextTuple{"/mtproto.RPCMessages/messages_saveGif", func() interface{} { return new(Bool) }},
	"TLMessagesSetInlineBotResults":        RPCContextTuple{"/mtproto.RPCMessages/messages_setInlineBotResults", func() interface{} { return new(Bool) }},
	"TLMessagesEditInlineBotMessage":       RPCContextTuple{"/mtproto.RPCMessages/messages_editInlineBotMessage", func() interface{} { return new(Bool) }},
	"TLMessagesSetBotCallbackAnswer":       RPCContextTuple{"/mtproto.RPCMessages/messages_setBotCallbackAnswer", func() interface{} { return new(Bool) }},
	"TLMessagesSaveDraft":                  RPCContextTuple{"/mtproto.RPCMessages/messages_saveDraft", func() interface{} { return new(Bool) }},
	"TLMessagesReadFeaturedStickers":       RPCContextTuple{"/mtproto.RPCMessages/messages_readFeaturedStickers", func() interface{} { return new(Bool) }},
	"TLMessagesSaveRecentSticker":          RPCContextTuple{"/mtproto.RPCMessages/messages_saveRecentSticker", func() interface{} { return new(Bool) }},
	"TLMessagesClearRecentStickers":        RPCContextTuple{"/mtproto.RPCMessages/messages_clearRecentStickers", func() interface{} { return new(Bool) }},
	"TLMessagesSetInlineGameScore":         RPCContextTuple{"/mtproto.RPCMessages/messages_setInlineGameScore", func() interface{} { return new(Bool) }},
	"TLMessagesToggleDialogPin":            RPCContextTuple{"/mtproto.RPCMessages/messages_toggleDialogPin", func() interface{} { return new(Bool) }},
	"TLMessagesReorderPinnedDialogs":       RPCContextTuple{"/mtproto.RPCMessages/messages_reorderPinnedDialogs", func() interface{} { return new(Bool) }},
	"TLMessagesSetBotShippingResults":      RPCContextTuple{"/mtproto.RPCMessages/messages_setBotShippingResults", func() interface{} { return new(Bool) }},
	"TLMessagesSetBotPrecheckoutResults":   RPCContextTuple{"/mtproto.RPCMessages/messages_setBotPrecheckoutResults", func() interface{} { return new(Bool) }},
	"TLMessagesFaveSticker":                RPCContextTuple{"/mtproto.RPCMessages/messages_faveSticker", func() interface{} { return new(Bool) }},
	"TLUploadSaveFilePart":                 RPCContextTuple{"/mtproto.RPCUpload/upload_saveFilePart", func() interface{} { return new(Bool) }},
	"TLUploadSaveBigFilePart":              RPCContextTuple{"/mtproto.RPCUpload/upload_saveBigFilePart", func() interface{} { return new(Bool) }},
	"TLHelpSaveAppLog":                     RPCContextTuple{"/mtproto.RPCHelp/help_saveAppLog", func() interface{} { return new(Bool) }},
	"TLHelpSetBotUpdatesStatus":            RPCContextTuple{"/mtproto.RPCHelp/help_setBotUpdatesStatus", func() interface{} { return new(Bool) }},
	"TLChannelsReadHistory":                RPCContextTuple{"/mtproto.RPCChannels/channels_readHistory", func() interface{} { return new(Bool) }},
	"TLChannelsReportSpam":                 RPCContextTuple{"/mtproto.RPCChannels/channels_reportSpam", func() interface{} { return new(Bool) }},
	"TLChannelsEditAbout":                  RPCContextTuple{"/mtproto.RPCChannels/channels_editAbout", func() interface{} { return new(Bool) }},
	"TLChannelsCheckUsername":              RPCContextTuple{"/mtproto.RPCChannels/channels_checkUsername", func() interface{} { return new(Bool) }},
	"TLChannelsUpdateUsername":             RPCContextTuple{"/mtproto.RPCChannels/channels_updateUsername", func() interface{} { return new(Bool) }},
	"TLChannelsSetStickers":                RPCContextTuple{"/mtproto.RPCChannels/channels_setStickers", func() interface{} { return new(Bool) }},
	"TLChannelsReadMessageContents":        RPCContextTuple{"/mtproto.RPCChannels/channels_readMessageContents", func() interface{} { return new(Bool) }},
	"TLChannelsDeleteHistory":              RPCContextTuple{"/mtproto.RPCChannels/channels_deleteHistory", func() interface{} { return new(Bool) }},
	"TLBotsAnswerWebhookJSONQuery":         RPCContextTuple{"/mtproto.RPCBots/bots_answerWebhookJSONQuery", func() interface{} { return new(Bool) }},
	"TLPaymentsClearSavedInfo":             RPCContextTuple{"/mtproto.RPCPayments/payments_clearSavedInfo", func() interface{} { return new(Bool) }},
	"TLPhoneReceivedCall":                  RPCContextTuple{"/mtproto.RPCPhone/phone_receivedCall", func() interface{} { return new(Bool) }},
	"TLPhoneSaveCallDebug":                 RPCContextTuple{"/mtproto.RPCPhone/phone_saveCallDebug", func() interface{} { return new(Bool) }},
	"TLAuthCheckPhone":                     RPCContextTuple{"/mtproto.RPCAuth/auth_checkPhone", func() interface{} { return new(Auth_CheckedPhone) }},
	"TLAuthSendCode":                       RPCContextTuple{"/mtproto.RPCAuth/auth_sendCode", func() interface{} { return new(Auth_SentCode) }},
	"TLAuthSendCodeLayer51":                RPCContextTuple{"/mtproto.RPCAuth/auth_sendCodeLayer51", func() interface{} { return new(Auth_SentCode) }},
	"TLAuthResendCode":                     RPCContextTuple{"/mtproto.RPCAuth/auth_resendCode", func() interface{} { return new(Auth_SentCode) }},
	"TLAccountSendChangePhoneCode":         RPCContextTuple{"/mtproto.RPCAccount/account_sendChangePhoneCode", func() interface{} { return new(Auth_SentCode) }},
	"TLAccountSendConfirmPhoneCode":        RPCContextTuple{"/mtproto.RPCAccount/account_sendConfirmPhoneCode", func() interface{} { return new(Auth_SentCode) }},
	"TLAuthSignUp":                         RPCContextTuple{"/mtproto.RPCAuth/auth_signUp", func() interface{} { return new(Auth_Authorization) }},
	"TLAuthSignIn":                         RPCContextTuple{"/mtproto.RPCAuth/auth_signIn", func() interface{} { return new(Auth_Authorization) }},
	"TLAuthImportAuthorization":            RPCContextTuple{"/mtproto.RPCAuth/auth_importAuthorization", func() interface{} { return new(Auth_Authorization) }},
	"TLAuthImportBotAuthorization":         RPCContextTuple{"/mtproto.RPCAuth/auth_importBotAuthorization", func() interface{} { return new(Auth_Authorization) }},
	"TLAuthCheckPassword":                  RPCContextTuple{"/mtproto.RPCAuth/auth_checkPassword", func() interface{} { return new(Auth_Authorization) }},
	"TLAuthRecoverPassword":                RPCContextTuple{"/mtproto.RPCAuth/auth_recoverPassword", func() interface{} { return new(Auth_Authorization) }},
	"TLAuthExportAuthorization":            RPCContextTuple{"/mtproto.RPCAuth/auth_exportAuthorization", func() interface{} { return new(Auth_ExportedAuthorization) }},
	"TLAuthRequestPasswordRecovery":        RPCContextTuple{"/mtproto.RPCAuth/auth_requestPasswordRecovery", func() interface{} { return new(Auth_PasswordRecovery) }},
	"TLAccountGetNotifySettings":           RPCContextTuple{"/mtproto.RPCAccount/account_getNotifySettings", func() interface{} { return new(PeerNotifySettings) }},
	"TLAccountUpdateProfile":               RPCContextTuple{"/mtproto.RPCAccount/account_updateProfile", func() interface{} { return new(User) }},
	"TLAccountUpdateUsername":              RPCContextTuple{"/mtproto.RPCAccount/account_updateUsername", func() interface{} { return new(User) }},
	"TLAccountChangePhone":                 RPCContextTuple{"/mtproto.RPCAccount/account_changePhone", func() interface{} { return new(User) }},
	"TLContactsImportCard":                 RPCContextTuple{"/mtproto.RPCContacts/contacts_importCard", func() interface{} { return new(User) }},
	"TLAccountGetPrivacy":                  RPCContextTuple{"/mtproto.RPCAccount/account_getPrivacy", func() interface{} { return new(Account_PrivacyRules) }},
	"TLAccountSetPrivacy":                  RPCContextTuple{"/mtproto.RPCAccount/account_setPrivacy", func() interface{} { return new(Account_PrivacyRules) }},
	"TLAccountGetAccountTTL":               RPCContextTuple{"/mtproto.RPCAccount/account_getAccountTTL", func() interface{} { return new(AccountDaysTTL) }},
	"TLAccountGetAuthorizations":           RPCContextTuple{"/mtproto.RPCAccount/account_getAuthorizations", func() interface{} { return new(Account_Authorizations) }},
	"TLAccountGetPassword":                 RPCContextTuple{"/mtproto.RPCAccount/account_getPassword", func() interface{} { return new(Account_Password) }},
	"TLAccountGetPasswordSettings":         RPCContextTuple{"/mtproto.RPCAccount/account_getPasswordSettings", func() interface{} { return new(Account_PasswordSettings) }},
	"TLAccountGetTmpPassword":              RPCContextTuple{"/mtproto.RPCAccount/account_getTmpPassword", func() interface{} { return new(Account_TmpPassword) }},
	"TLUsersGetFullUser":                   RPCContextTuple{"/mtproto.RPCUsers/users_getFullUser", func() interface{} { return new(UserFull) }},
	"TLContactsGetContacts":                RPCContextTuple{"/mtproto.RPCContacts/contacts_getContacts", func() interface{} { return new(Contacts_Contacts) }},
	"TLContactsImportContacts":             RPCContextTuple{"/mtproto.RPCContacts/contacts_importContacts", func() interface{} { return new(Contacts_ImportedContacts) }},
	"TLContactsDeleteContact":              RPCContextTuple{"/mtproto.RPCContacts/contacts_deleteContact", func() interface{} { return new(Contacts_Link) }},
	"TLContactsGetBlocked":                 RPCContextTuple{"/mtproto.RPCContacts/contacts_getBlocked", func() interface{} { return new(Contacts_Blocked) }},
	"TLContactsSearch":                     RPCContextTuple{"/mtproto.RPCContacts/contacts_search", func() interface{} { return new(Contacts_Found) }},
	"TLContactsResolveUsername":            RPCContextTuple{"/mtproto.RPCContacts/contacts_resolveUsername", func() interface{} { return new(Contacts_ResolvedPeer) }},
	"TLContactsGetTopPeers":                RPCContextTuple{"/mtproto.RPCContacts/contacts_getTopPeers", func() interface{} { return new(Contacts_TopPeers) }},
	"TLMessagesGetMessages":                RPCContextTuple{"/mtproto.RPCMessages/messages_getMessages", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesGetHistory":                 RPCContextTuple{"/mtproto.RPCMessages/messages_getHistory", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesGetHistoryLayer71":          RPCContextTuple{"/mtproto.RPCMessages/messages_getHistoryLayer71", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesSearch":                     RPCContextTuple{"/mtproto.RPCMessages/messages_search", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesSearchLayer68":              RPCContextTuple{"/mtproto.RPCMessages/messages_searchLayer68", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesSearchGlobal":               RPCContextTuple{"/mtproto.RPCMessages/messages_searchGlobal", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesGetUnreadMentions":          RPCContextTuple{"/mtproto.RPCMessages/messages_getUnreadMentions", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesGetRecentLocations":         RPCContextTuple{"/mtproto.RPCMessages/messages_getRecentLocations", func() interface{} { return new(Messages_Messages) }},
	"TLChannelsGetMessages":                RPCContextTuple{"/mtproto.RPCChannels/channels_getMessages", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesGetDialogs":                 RPCContextTuple{"/mtproto.RPCMessages/messages_getDialogs", func() interface{} { return new(Messages_Dialogs) }},
	"TLMessagesReadHistory":                RPCContextTuple{"/mtproto.RPCMessages/messages_readHistory", func() interface{} { return new(Messages_AffectedMessages) }},
	"TLMessagesReadHistoryLayer2":          RPCContextTuple{"/mtproto.RPCMessages/messages_readHistoryLayer2", func() interface{} { return new(Messages_AffectedMessages) }},
	"TLMessagesDeleteMessages":             RPCContextTuple{"/mtproto.RPCMessages/messages_deleteMessages", func() interface{} { return new(Messages_AffectedMessages) }},
	"TLMessagesReadMessageContents":        RPCContextTuple{"/mtproto.RPCMessages/messages_readMessageContents", func() interface{} { return new(Messages_AffectedMessages) }},
	"TLChannelsDeleteMessages":             RPCContextTuple{"/mtproto.RPCChannels/channels_deleteMessages", func() interface{} { return new(Messages_AffectedMessages) }},
	"TLMessagesDeleteHistory":              RPCContextTuple{"/mtproto.RPCMessages/messages_deleteHistory", func() interface{} { return new(Messages_AffectedHistory) }},
	"TLMessagesReadMentions":               RPCContextTuple{"/mtproto.RPCMessages/messages_readMentions", func() interface{} { return new(Messages_AffectedHistory) }},
	"TLChannelsDeleteUserHistory":          RPCContextTuple{"/mtproto.RPCChannels/channels_deleteUserHistory", func() interface{} { return new(Messages_AffectedHistory) }},
	"TLMessagesSendMessage":                RPCContextTuple{"/mtproto.RPCMessages/messages_sendMessage", func() interface{} { return new(Updates) }},
	"TLMessagesSendMedia":                  RPCContextTuple{"/mtproto.RPCMessages/messages_sendMedia", func() interface{} { return new(Updates) }},
	"TLMessagesForwardMessages":            RPCContextTuple{"/mtproto.RPCMessages/messages_forwardMessages", func() interface{} { return new(Updates) }},
	"TLMessagesEditChatTitle":              RPCContextTuple{"/mtproto.RPCMessages/messages_editChatTitle", func() interface{} { return new(Updates) }},
	"TLMessagesEditChatPhoto":              RPCContextTuple{"/mtproto.RPCMessages/messages_editChatPhoto", func() interface{} { return new(Updates) }},
	"TLMessagesAddChatUser":                RPCContextTuple{"/mtproto.RPCMessages/messages_addChatUser", func() interface{} { return new(Updates) }},
	"TLMessagesDeleteChatUser":             RPCContextTuple{"/mtproto.RPCMessages/messages_deleteChatUser", func() interface{} { return new(Updates) }},
	"TLMessagesCreateChat":                 RPCContextTuple{"/mtproto.RPCMessages/messages_createChat", func() interface{} { return new(Updates) }},
	"TLMessagesForwardMessage":             RPCContextTuple{"/mtproto.RPCMessages/messages_forwardMessage", func() interface{} { return new(Updates) }},
	"TLMessagesImportChatInvite":           RPCContextTuple{"/mtproto.RPCMessages/messages_importChatInvite", func() interface{} { return new(Updates) }},
	"TLMessagesStartBot":                   RPCContextTuple{"/mtproto.RPCMessages/messages_startBot", func() interface{} { return new(Updates) }},
	"TLMessagesToggleChatAdmins":           RPCContextTuple{"/mtproto.RPCMessages/messages_toggleChatAdmins", func() interface{} { return new(Updates) }},
	"TLMessagesMigrateChat":                RPCContextTuple{"/mtproto.RPCMessages/messages_migrateChat", func() interface{} { return new(Updates) }},
	"TLMessagesSendInlineBotResult":        RPCContextTuple{"/mtproto.RPCMessages/messages_sendInlineBotResult", func() interface{} { return new(Updates) }},
	"TLMessagesEditMessage":                RPCContextTuple{"/mtproto.RPCMessages/messages_editMessage", func() interface{} { return new(Updates) }},
	"TLMessagesGetAllDrafts":               RPCContextTuple{"/mtproto.RPCMessages/messages_getAllDrafts", func() interface{} { return new(Updates) }},
	"TLMessagesSetGameScore":               RPCContextTuple{"/mtproto.RPCMessages/messages_setGameScore", func() interface{} { return new(Updates) }},
	"TLMessagesSendScreenshotNotification": RPCContextTuple{"/mtproto.RPCMessages/messages_sendScreenshotNotification", func() interface{} { return new(Updates) }},
	"TLMessagesSendMultiMedia":             RPCContextTuple{"/mtproto.RPCMessages/messages_sendMultiMedia", func() interface{} { return new(Updates) }},
	"TLHelpGetAppChangelog":                RPCContextTuple{"/mtproto.RPCHelp/help_getAppChangelog", func() interface{} { return new(Updates) }},
	"TLChannelsCreateChannel":              RPCContextTuple{"/mtproto.RPCChannels/channels_createChannel", func() interface{} { return new(Updates) }},
	"TLChannelsEditAdmin":                  RPCContextTuple{"/mtproto.RPCChannels/channels_editAdmin", func() interface{} { return new(Updates) }},
	"TLChannelsEditTitle":                  RPCContextTuple{"/mtproto.RPCChannels/channels_editTitle", func() interface{} { return new(Updates) }},
	"TLChannelsEditPhoto":                  RPCContextTuple{"/mtproto.RPCChannels/channels_editPhoto", func() interface{} { return new(Updates) }},
	"TLChannelsJoinChannel":                RPCContextTuple{"/mtproto.RPCChannels/channels_joinChannel", func() interface{} { return new(Updates) }},
	"TLChannelsLeaveChannel":               RPCContextTuple{"/mtproto.RPCChannels/channels_leaveChannel", func() interface{} { return new(Updates) }},
	"TLChannelsInviteToChannel":            RPCContextTuple{"/mtproto.RPCChannels/channels_inviteToChannel", func() interface{} { return new(Updates) }},
	"TLChannelsDeleteChannel":              RPCContextTuple{"/mtproto.RPCChannels/channels_deleteChannel", func() interface{} { return new(Updates) }},
	"TLChannelsToggleInvites":              RPCContextTuple{"/mtproto.RPCChannels/channels_toggleInvites", func() interface{} { return new(Updates) }},
	"TLChannelsToggleSignatures":           RPCContextTuple{"/mtproto.RPCChannels/channels_toggleSignatures", func() interface{} { return new(Updates) }},
	"TLChannelsUpdatePinnedMessage":        RPCContextTuple{"/mtproto.RPCChannels/channels_updatePinnedMessage", func() interface{} { return new(Updates) }},
	"TLChannelsEditBanned":                 RPCContextTuple{"/mtproto.RPCChannels/channels_editBanned", func() interface{} { return new(Updates) }},
	"TLChannelsTogglePreHistoryHidden":     RPCContextTuple{"/mtproto.RPCChannels/channels_togglePreHistoryHidden", func() interface{} { return new(Updates) }},
	"TLPhoneDiscardCall":                   RPCContextTuple{"/mtproto.RPCPhone/phone_discardCall", func() interface{} { return new(Updates) }},
	"TLPhoneSetCallRating":                 RPCContextTuple{"/mtproto.RPCPhone/phone_setCallRating", func() interface{} { return new(Updates) }},
	"TLMessagesGetPeerSettings":            RPCContextTuple{"/mtproto.RPCMessages/messages_getPeerSettings", func() interface{} { return new(PeerSettings) }},
	"TLMessagesGetChats":                   RPCContextTuple{"/mtproto.RPCMessages/messages_getChats", func() interface{} { return new(Messages_Chats) }},
	"TLMessagesGetCommonChats":             RPCContextTuple{"/mtproto.RPCMessages/messages_getCommonChats", func() interface{} { return new(Messages_Chats) }},
	"TLMessagesGetAllChats":                RPCContextTuple{"/mtproto.RPCMessages/messages_getAllChats", func() interface{} { return new(Messages_Chats) }},
	"TLChannelsGetChannels":                RPCContextTuple{"/mtproto.RPCChannels/channels_getChannels", func() interface{} { return new(Messages_Chats) }},
	"TLChannelsGetAdminedPublicChannels":   RPCContextTuple{"/mtproto.RPCChannels/channels_getAdminedPublicChannels", func() interface{} { return new(Messages_Chats) }},
	"TLMessagesGetFullChat":                RPCContextTuple{"/mtproto.RPCMessages/messages_getFullChat", func() interface{} { return new(Messages_ChatFull) }},
	"TLChannelsGetFullChannel":             RPCContextTuple{"/mtproto.RPCChannels/channels_getFullChannel", func() interface{} { return new(Messages_ChatFull) }},
	"TLMessagesGetDhConfig":                RPCContextTuple{"/mtproto.RPCMessages/messages_getDhConfig", func() interface{} { return new(Messages_DhConfig) }},
	"TLMessagesRequestEncryption":          RPCContextTuple{"/mtproto.RPCMessages/messages_requestEncryption", func() interface{} { return new(EncryptedChat) }},
	"TLMessagesAcceptEncryption":           RPCContextTuple{"/mtproto.RPCMessages/messages_acceptEncryption", func() interface{} { return new(EncryptedChat) }},
	"TLMessagesSendEncrypted":              RPCContextTuple{"/mtproto.RPCMessages/messages_sendEncrypted", func() interface{} { return new(Messages_SentEncryptedMessage) }},
	"TLMessagesSendEncryptedFile":          RPCContextTuple{"/mtproto.RPCMessages/messages_sendEncryptedFile", func() interface{} { return new(Messages_SentEncryptedMessage) }},
	"TLMessagesSendEncryptedService":       RPCContextTuple{"/mtproto.RPCMessages/messages_sendEncryptedService", func() interface{} { return new(Messages_SentEncryptedMessage) }},
	"TLMessagesGetAllStickers":             RPCContextTuple{"/mtproto.RPCMessages/messages_getAllStickers", func() interface{} { return new(Messages_AllStickers) }},
	"TLMessagesGetMaskStickers":            RPCContextTuple{"/mtproto.RPCMessages/messages_getMaskStickers", func() interface{} { return new(Messages_AllStickers) }},
	"TLMessagesGetWebPagePreview":          RPCContextTuple{"/mtproto.RPCMessages/messages_getWebPagePreview", func() interface{} { return new(MessageMedia) }},
	"TLMessagesUploadMedia":                RPCContextTuple{"/mtproto.RPCMessages/messages_uploadMedia", func() interface{} { return new(MessageMedia) }},
	"TLMessagesExportChatInvite":           RPCContextTuple{"/mtproto.RPCMessages/messages_exportChatInvite", func() interface{} { return new(ExportedChatInvite) }},
	"TLChannelsExportInvite":               RPCContextTuple{"/mtproto.RPCChannels/channels_exportInvite", func() interface{} { return new(ExportedChatInvite) }},
	"TLMessagesCheckChatInvite":            RPCContextTuple{"/mtproto.RPCMessages/messages_checkChatInvite", func() interface{} { return new(ChatInvite) }},
	"TLMessagesGetStickerSet":              RPCContextTuple{"/mtproto.RPCMessages/messages_getStickerSet", func() interface{} { return new(Messages_StickerSet) }},
	"TLStickersCreateStickerSet":           RPCContextTuple{"/mtproto.RPCStickers/stickers_createStickerSet", func() interface{} { return new(Messages_StickerSet) }},
	"TLStickersRemoveStickerFromSet":       RPCContextTuple{"/mtproto.RPCStickers/stickers_removeStickerFromSet", func() interface{} { return new(Messages_StickerSet) }},
	"TLStickersChangeStickerPosition":      RPCContextTuple{"/mtproto.RPCStickers/stickers_changeStickerPosition", func() interface{} { return new(Messages_StickerSet) }},
	"TLStickersAddStickerToSet":            RPCContextTuple{"/mtproto.RPCStickers/stickers_addStickerToSet", func() interface{} { return new(Messages_StickerSet) }},
	"TLMessagesInstallStickerSet":          RPCContextTuple{"/mtproto.RPCMessages/messages_installStickerSet", func() interface{} { return new(Messages_StickerSetInstallResult) }},
	"TLMessagesGetDocumentByHash":          RPCContextTuple{"/mtproto.RPCMessages/messages_getDocumentByHash", func() interface{} { return new(Document) }},
	"TLMessagesSearchGifs":                 RPCContextTuple{"/mtproto.RPCMessages/messages_searchGifs", func() interface{} { return new(Messages_FoundGifs) }},
	"TLMessagesGetSavedGifs":               RPCContextTuple{"/mtproto.RPCMessages/messages_getSavedGifs", func() interface{} { return new(Messages_SavedGifs) }},
	"TLMessagesGetInlineBotResults":        RPCContextTuple{"/mtproto.RPCMessages/messages_getInlineBotResults", func() interface{} { return new(Messages_BotResults) }},
	"TLMessagesGetMessageEditData":         RPCContextTuple{"/mtproto.RPCMessages/messages_getMessageEditData", func() interface{} { return new(Messages_MessageEditData) }},
	"TLMessagesGetBotCallbackAnswer":       RPCContextTuple{"/mtproto.RPCMessages/messages_getBotCallbackAnswer", func() interface{} { return new(Messages_BotCallbackAnswer) }},
	"TLMessagesGetPeerDialogs":             RPCContextTuple{"/mtproto.RPCMessages/messages_getPeerDialogs", func() interface{} { return new(Messages_PeerDialogs) }},
	"TLMessagesGetPinnedDialogs":           RPCContextTuple{"/mtproto.RPCMessages/messages_getPinnedDialogs", func() interface{} { return new(Messages_PeerDialogs) }},
	"TLMessagesGetFeaturedStickers":        RPCContextTuple{"/mtproto.RPCMessages/messages_getFeaturedStickers", func() interface{} { return new(Messages_FeaturedStickers) }},
	"TLMessagesGetRecentStickers":          RPCContextTuple{"/mtproto.RPCMessages/messages_getRecentStickers", func() interface{} { return new(Messages_RecentStickers) }},
	"TLMessagesGetArchivedStickers":        RPCContextTuple{"/mtproto.RPCMessages/messages_getArchivedStickers", func() interface{} { return new(Messages_ArchivedStickers) }},
	"TLMessagesGetGameHighScores":          RPCContextTuple{"/mtproto.RPCMessages/messages_getGameHighScores", func() interface{} { return new(Messages_HighScores) }},
	"TLMessagesGetInlineGameHighScores":    RPCContextTuple{"/mtproto.RPCMessages/messages_getInlineGameHighScores", func() interface{} { return new(Messages_HighScores) }},
	"TLMessagesGetWebPage":                 RPCContextTuple{"/mtproto.RPCMessages/messages_getWebPage", func() interface{} { return new(WebPage) }},
	"TLMessagesGetFavedStickers":           RPCContextTuple{"/mtproto.RPCMessages/messages_getFavedStickers", func() interface{} { return new(Messages_FavedStickers) }},
	"TLMessagesUploadEncryptedFile":        RPCContextTuple{"/mtproto.RPCMessages/messages_uploadEncryptedFile", func() interface{} { return new(EncryptedFile) }},
	"TLUpdatesGetState":                    RPCContextTuple{"/mtproto.RPCUpdates/updates_getState", func() interface{} { return new(Updates_State) }},
	"TLUpdatesGetDifference":               RPCContextTuple{"/mtproto.RPCUpdates/updates_getDifference", func() interface{} { return new(Updates_Difference) }},
	"TLUpdatesGetChannelDifference":        RPCContextTuple{"/mtproto.RPCUpdates/updates_getChannelDifference", func() interface{} { return new(Updates_ChannelDifference) }},
	"TLPhotosUpdateProfilePhoto":           RPCContextTuple{"/mtproto.RPCPhotos/photos_updateProfilePhoto", func() interface{} { return new(UserProfilePhoto) }},
	"TLPhotosUploadProfilePhoto":           RPCContextTuple{"/mtproto.RPCPhotos/photos_uploadProfilePhoto", func() interface{} { return new(Photos_Photo) }},
	"TLPhotosGetUserPhotos":                RPCContextTuple{"/mtproto.RPCPhotos/photos_getUserPhotos", func() interface{} { return new(Photos_Photos) }},
	"TLUploadGetFile":                      RPCContextTuple{"/mtproto.RPCUpload/upload_getFile", func() interface{} { return new(Upload_File) }},
	"TLUploadGetWebFile":                   RPCContextTuple{"/mtproto.RPCUpload/upload_getWebFile", func() interface{} { return new(Upload_WebFile) }},
	"TLUploadGetCdnFile":                   RPCContextTuple{"/mtproto.RPCUpload/upload_getCdnFile", func() interface{} { return new(Upload_CdnFile) }},
	"TLHelpGetConfig":                      RPCContextTuple{"/mtproto.RPCHelp/help_getConfig", func() interface{} { return new(Config) }},
	"TLHelpGetNearestDc":                   RPCContextTuple{"/mtproto.RPCHelp/help_getNearestDc", func() interface{} { return new(NearestDc) }},
	"TLHelpGetAppUpdate":                   RPCContextTuple{"/mtproto.RPCHelp/help_getAppUpdate", func() interface{} { return new(Help_AppUpdate) }},
	"TLHelpGetAppUpdateLayer46":            RPCContextTuple{"/mtproto.RPCHelp/help_getAppUpdateLayer46", func() interface{} { return new(Help_AppUpdate) }},
	"TLHelpGetInviteText":                  RPCContextTuple{"/mtproto.RPCHelp/help_getInviteText", func() interface{} { return new(Help_InviteText) }},
	"TLHelpGetInviteTextLayer46":           RPCContextTuple{"/mtproto.RPCHelp/help_getInviteTextLayer46", func() interface{} { return new(Help_InviteText) }},
	"TLHelpGetSupport":                     RPCContextTuple{"/mtproto.RPCHelp/help_getSupport", func() interface{} { return new(Help_Support) }},
	"TLHelpGetTermsOfService":              RPCContextTuple{"/mtproto.RPCHelp/help_getTermsOfService", func() interface{} { return new(Help_TermsOfService) }},
	"TLHelpGetCdnConfig":                   RPCContextTuple{"/mtproto.RPCHelp/help_getCdnConfig", func() interface{} { return new(CdnConfig) }},
	"TLHelpGetRecentMeUrls":                RPCContextTuple{"/mtproto.RPCHelp/help_getRecentMeUrls", func() interface{} { return new(Help_RecentMeUrls) }},
	"TLHelpGetScheme":                		RPCContextTuple{"/mtproto.RPCHelp/help_getScheme", func() interface{} { return new(Scheme) }},
	"TLChannelsGetParticipants":            RPCContextTuple{"/mtproto.RPCChannels/channels_getParticipants", func() interface{} { return new(Channels_ChannelParticipants) }},
	"TLChannelsGetParticipant":             RPCContextTuple{"/mtproto.RPCChannels/channels_getParticipant", func() interface{} { return new(Channels_ChannelParticipant) }},
	"TLChannelsExportMessageLink":          RPCContextTuple{"/mtproto.RPCChannels/channels_exportMessageLink", func() interface{} { return new(ExportedMessageLink) }},
	"TLChannelsExportMessageLinkLayer74":   RPCContextTuple{"/mtproto.RPCChannels/channels_exportMessageLinkLayer74", func() interface{} { return new(ExportedMessageLink) }},
	"TLChannelsGetAdminLog":                RPCContextTuple{"/mtproto.RPCChannels/channels_getAdminLog", func() interface{} { return new(Channels_AdminLogResults) }},
	"TLBotsSendCustomRequest":              RPCContextTuple{"/mtproto.RPCBots/bots_sendCustomRequest", func() interface{} { return new(DataJSON) }},
	"TLPhoneGetCallConfig":                 RPCContextTuple{"/mtproto.RPCPhone/phone_getCallConfig", func() interface{} { return new(DataJSON) }},
	"TLPaymentsGetPaymentForm":             RPCContextTuple{"/mtproto.RPCPayments/payments_getPaymentForm", func() interface{} { return new(Payments_PaymentForm) }},
	"TLPaymentsGetPaymentReceipt":          RPCContextTuple{"/mtproto.RPCPayments/payments_getPaymentReceipt", func() interface{} { return new(Payments_PaymentReceipt) }},
	"TLPaymentsValidateRequestedInfo":      RPCContextTuple{"/mtproto.RPCPayments/payments_validateRequestedInfo", func() interface{} { return new(Payments_ValidatedRequestedInfo) }},
	"TLPaymentsSendPaymentForm":            RPCContextTuple{"/mtproto.RPCPayments/payments_sendPaymentForm", func() interface{} { return new(Payments_PaymentResult) }},
	"TLPaymentsGetSavedInfo":               RPCContextTuple{"/mtproto.RPCPayments/payments_getSavedInfo", func() interface{} { return new(Payments_SavedInfo) }},
	"TLPhoneRequestCall":                   RPCContextTuple{"/mtproto.RPCPhone/phone_requestCall", func() interface{} { return new(Phone_PhoneCall) }},
	"TLPhoneAcceptCall":                    RPCContextTuple{"/mtproto.RPCPhone/phone_acceptCall", func() interface{} { return new(Phone_PhoneCall) }},
	"TLPhoneConfirmCall":                   RPCContextTuple{"/mtproto.RPCPhone/phone_confirmCall", func() interface{} { return new(Phone_PhoneCall) }},
	"TLLangpackGetLangPack":                RPCContextTuple{"/mtproto.RPCLangpack/langpack_getLangPack", func() interface{} { return new(LangPackDifference) }},
	"TLLangpackGetDifference":              RPCContextTuple{"/mtproto.RPCLangpack/langpack_getDifference", func() interface{} { return new(LangPackDifference) }},

	// TODO(@benqi): Not support Vector Reply
	"TLAccountGetWallPapers":        RPCContextTuple{"/mtproto.RPCAccount/account_getWallPapers", func() interface{} { return new(Vector_WallPaper) }},
	"TLUsersGetUsers":               RPCContextTuple{"/mtproto.RPCUsers/users_getUsers", func() interface{} { return new(Vector_User) }},
	"TLContactsGetStatuses":         RPCContextTuple{"/mtproto.RPCContacts/contacts_getStatuses", func() interface{} { return new(Vector_ContactStatus) }},
	"TLContactsExportCard":          RPCContextTuple{"/mtproto.RPCContacts/contacts_exportCard", func() interface{} { return new(VectorInt) }},
	"TLMessagesGetMessagesViews":    RPCContextTuple{"/mtproto.RPCMessages/messages_getMessagesViews", func() interface{} { return new(VectorInt) }},
	"TLMessagesReceivedMessages":    RPCContextTuple{"/mtproto.RPCMessages/messages_receivedMessages", func() interface{} { return new(Vector_ReceivedNotifyMessage) }},
	"TLMessagesReceivedQueue":       RPCContextTuple{"/mtproto.RPCMessages/messages_receivedQueue", func() interface{} { return new(VectorLong) }},
	"TLPhotosDeletePhotos":          RPCContextTuple{"/mtproto.RPCPhotos/photos_deletePhotos", func() interface{} { return new(VectorLong) }},
	"TLMessagesGetAttachedStickers": RPCContextTuple{"/mtproto.RPCMessages/messages_getAttachedStickers", func() interface{} { return new(Vector_StickerSetCovered) }},
	"TLUploadReuploadCdnFile":       RPCContextTuple{"/mtproto.RPCUpload/upload_reuploadCdnFile", func() interface{} { return new(Vector_CdnFileHash) }},
	"TLUploadGetCdnFileHashes":      RPCContextTuple{"/mtproto.RPCUpload/upload_getCdnFileHashes", func() interface{} { return new(Vector_CdnFileHash) }},
	"TLLangpackGetStrings":          RPCContextTuple{"/mtproto.RPCLangpack/langpack_getStrings", func() interface{} { return new(Vector_LangPackString) }},
	"TLLangpackGetLanguages":        RPCContextTuple{"/mtproto.RPCLangpack/langpack_getLanguages", func() interface{} { return new(Vector_LangPackLanguage) }},
}

func FindRPCContextTuple(t interface{}) *RPCContextTuple {
	rt := reflect.TypeOf(t)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	m, ok := rpcContextRegisters[rt.Name()]
	if !ok {
		glog.Error("Can't find name: ", rt.Name())
		return nil
	}
	return &m
}
