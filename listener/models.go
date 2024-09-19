package listener

type WebhookListener func(message *WebhookMessage) error

type WebhookMessage struct {
	Wook        string             `json:"wook"`
	Status      string             `json:"status"`
	Type        string             `json:"type"`
	FromMe      bool               `json:"fromMe"`
	Id          string             `json:"id"`
	Session     string             `json:"session"`
	IsGroupMsg  bool               `json:"isGroupMsg"`
	Author      string             `json:"author"`
	Name        string             `json:"name"`
	To          string             `json:"to"`
	From        string             `json:"from"`
	Thumbnail   string             `json:"thumbnail"`
	Url         string             `json:"url"`
	QuotedMsg   string             `json:"quotedMsg"`
	QuotedMsgId string             `json:"quotedMsgId"`
	Datetime    string             `json:"datetime"`
	Data        WebhookMessageData `json:"data"`
}

type WebhookMessageData struct {
	Id                                    string                        `json:"id"`
	Viewed                                bool                          `json:"viewed"`
	Body                                  string                        `json:"body"`
	Type                                  string                        `json:"type"`
	Subtype                               string                        `json:"subtype"`
	T                                     int                           `json:"t"`
	NotifyName                            string                        `json:"notifyName"`
	From                                  string                        `json:"from"`
	To                                    string                        `json:"to"`
	Author                                string                        `json:"author"`
	Ack                                   int                           `json:"ack"`
	Invis                                 bool                          `json:"invis"`
	IsNewMsg                              bool                          `json:"isNewMsg"`
	Star                                  bool                          `json:"star"`
	KicNotified                           bool                          `json:"kicNotified"`
	RecvFresh                             bool                          `json:"recvFresh"`
	IsFromTemplate                        bool                          `json:"isFromTemplate"`
	MatchedText                           string                        `json:"matchedText"`
	Thumbnail                             string                        `json:"thumbnail"`
	RichPreviewType                       int                           `json:"richPreviewType"`
	Rcat                                  any                           `json:"rcat"`
	PollInvalidated                       bool                          `json:"pollInvalidated"`
	IsSentCagPollCreation                 bool                          `json:"isSentCagPollCreation"`
	LatestEditMsgKey                      any                           `json:"latestEditMsgKey"`
	LatestEditSenderTimestampMs           any                           `json:"latestEditSenderTimestampMs"`
	MentionedJidList                      []any                         `json:"mentionedJidList"`
	GroupMentions                         []any                         `json:"groupMentions"`
	IsEventCanceled                       bool                          `json:"isEventCanceled"`
	EventInvalidated                      bool                          `json:"eventInvalidated"`
	IsVcardOverMmsDocument                bool                          `json:"isVcardOverMmsDocument"`
	IsForwarded                           bool                          `json:"isForwarded"`
	ForwardingScore                       int                           `json:"forwardingScore"`
	Labels                                []string                      `json:"labels"`
	HasReaction                           bool                          `json:"hasReaction"`
	InviteGrpType                         string                        `json:"inviteGrpType"`
	CtwaContext                           WebhookMessageDataCTWAContext `json:"ctwaContext"`
	ProductHeaderImageRejected            bool                          `json:"productHeaderImageRejected"`
	LastPlaybackProgress                  int                           `json:"lastPlaybackProgress"`
	IsDynamicReplyButtonsMsg              bool                          `json:"isDynamicReplyButtonsMsg"`
	IsCarouselCard                        bool                          `json:"isCarouselCard"`
	ParentMsgId                           any                           `json:"parentMsgId"`
	IsMdHistoryMsg                        bool                          `json:"isMdHistoryMsg"`
	StickerSentTs                         int                           `json:"stickerSentTs"`
	IsAvatar                              bool                          `json:"isAvatar"`
	LastUpdateFromServerTs                int                           `json:"lastUpdateFromServerTs"`
	InvokedBotWid                         any                           `json:"invokedBotWid"`
	BizBotType                            any                           `json:"bizBotType"`
	BotResponseTargetId                   any                           `json:"botResponseTargetId"`
	BotPluginType                         any                           `json:"botPluginType"`
	BotPluginReferenceIndex               any                           `json:"botPluginReferenceIndex"`
	BotPluginSearchProvider               any                           `json:"botPluginSearchProvider"`
	BotPluginSearchUrl                    any                           `json:"botPluginSearchUrl"`
	BotPluginSearchQuery                  any                           `json:"botPluginSearchQuery"`
	BotPluginMaybeParent                  bool                          `json:"botPluginMaybeParent"`
	BotReelPluginThumbnailCdnUrl          any                           `json:"botReelPluginThumbnailCdnUrl"`
	BotMsgBodyType                        any                           `json:"botMsgBodyType"`
	RequiresDirectConnection              any                           `json:"requiresDirectConnection"`
	BizContentPlaceholderType             any                           `json:"bizContentPlaceholderType"`
	HostedBizEncStateMismatch             bool                          `json:"hostedBizEncStateMismatch"`
	SenderOrRecipientAccountTypeHosted    bool                          `json:"senderOrRecipientAccountTypeHosted"`
	PlaceholderCreatedWhenAccountIsHosted bool                          `json:"placeholderCreatedWhenAccountIsHosted"`
	ChatId                                string                        `json:"chatId"`
	FromMe                                bool                          `json:"fromMe"`
	Sender                                WebhookMessageDataSender      `json:"sender"`
	Timestamp                             int                           `json:"timestamp"`
	Content                               string                        `json:"content"`
	IsGroupMsg                            bool                          `json:"isGroupMsg"`
	MediaData                             any                           `json:"mediaData"`
}

type WebhookMessageDataSender struct {
	Id                       string `json:"id"`
	Pushname                 string `json:"pushname"`
	Type                     string `json:"type"`
	Labels                   []any  `json:"labels"`
	IsContactSyncCompleted   int    `json:"isContactSyncCompleted"`
	TextStatusLastUpdateTime int    `json:"textStatusLastUpdateTime"`
	SyncToAddressbook        bool   `json:"syncToAddressbook"`
	FormattedName            string `json:"formattedName"`
	IsMe                     bool   `json:"isMe"`
	IsMyContact              bool   `json:"isMyContact"`
	IsPSA                    bool   `json:"isPSA"`
	IsUser                   bool   `json:"isUser"`
	IsWAContact              bool   `json:"isWAContact"`
	ProfilePicThumbObj       struct {
	} `json:"profilePicThumbObj"`
	Msgs any `json:"msgs"`
}

type WebhookMessageDataCTWAContext struct {
	SourceUrl                 string `json:"sourceUrl"`
	Description               string `json:"description"`
	Title                     string `json:"title"`
	Thumbnail                 string `json:"thumbnail"`
	ThumbnailUrl              string `json:"thumbnailUrl"`
	MediaType                 int    `json:"mediaType"`
	MediaUrl                  string `json:"mediaUrl"`
	AdContextPreviewDismissed bool   `json:"adContextPreviewDismissed"`
}

//type WebhookMessage struct {
//	IsStatusReply  bool                 `json:"isStatusReply"`
//	ChatLid        string               `json:"chatLid"`
//	ConnectedPhone string               `json:"connectedPhone"`
//	WaitingMessage bool                 `json:"waitingMessage"`
//	IsEdit         bool                 `json:"isEdit"`
//	IsGroup        bool                 `json:"isGroup"`
//	IsNewsletter   bool                 `json:"isNewsletter"`
//	InstanceId     string               `json:"instanceId"`
//	MessageId      string               `json:"messageId"`
//	Phone          string               `json:"phone"`
//	FromMe         bool                 `json:"fromMe"`
//	Momment        int64                `json:"momment"`
//	Status         string               `json:"status"`
//	ChatName       string               `json:"chatName"`
//	SenderPhoto    any          `json:"senderPhoto"`
//	SenderName     string               `json:"senderName"`
//	Photo          string               `json:"photo"`
//	Broadcast      bool                 `json:"broadcast"`
//	ParticipantLid any          `json:"participantLid"`
//	Forwarded      bool                 `json:"forwarded"`
//	Type           string               `json:"type"`
//	FromApi        bool                 `json:"fromApi"`
//	Text           *WebhookMessageText  `json:"text"`
//	Audio          *WebhookMessageAudio `json:"audio"`
//}
//
//type WebhookMessageAudio struct {
//	Ptt      bool   `json:"ptt"`
//	Seconds  int    `json:"seconds"`
//	AudioUrl string `json:"audioUrl"`
//	MimeType string `json:"mimeType"`
//	ViewOnce bool   `json:"viewOnce"`
//}
//
//type WebhookMessageText struct {
//	Message string `json:"message"`
//}
