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
	Author      interface{}        `json:"author"`
	Name        string             `json:"name"`
	To          string             `json:"to"`
	From        string             `json:"from"`
	Mimetype    string             `json:"mimetype"`
	Thumbnail   string             `json:"thumbnail"`
	Url         string             `json:"url"`
	QuotedMsg   interface{}        `json:"quotedMsg"`
	QuotedMsgId string             `json:"quotedMsgId"`
	Datetime    string             `json:"datetime"`
	Base64      string             `json:"base64"`
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
	Author                                interface{}                   `json:"author"`
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
	Rcat                                  interface{}                   `json:"rcat"`
	PollInvalidated                       bool                          `json:"pollInvalidated"`
	IsSentCagPollCreation                 bool                          `json:"isSentCagPollCreation"`
	LatestEditMsgKey                      interface{}                   `json:"latestEditMsgKey"`
	LatestEditSenderTimestampMs           interface{}                   `json:"latestEditSenderTimestampMs"`
	MentionedJidList                      []interface{}                 `json:"mentionedJidList"`
	GroupMentions                         []interface{}                 `json:"groupMentions"`
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
	ParentMsgId                           interface{}                   `json:"parentMsgId"`
	IsMdHistoryMsg                        bool                          `json:"isMdHistoryMsg"`
	StickerSentTs                         int                           `json:"stickerSentTs"`
	IsAvatar                              bool                          `json:"isAvatar"`
	LastUpdateFromServerTs                int                           `json:"lastUpdateFromServerTs"`
	InvokedBotWid                         interface{}                   `json:"invokedBotWid"`
	BizBotType                            interface{}                   `json:"bizBotType"`
	BotResponseTargetId                   interface{}                   `json:"botResponseTargetId"`
	BotPluginType                         interface{}                   `json:"botPluginType"`
	BotPluginReferenceIndex               interface{}                   `json:"botPluginReferenceIndex"`
	BotPluginSearchProvider               interface{}                   `json:"botPluginSearchProvider"`
	BotPluginSearchUrl                    interface{}                   `json:"botPluginSearchUrl"`
	BotPluginSearchQuery                  interface{}                   `json:"botPluginSearchQuery"`
	BotPluginMaybeParent                  bool                          `json:"botPluginMaybeParent"`
	BotReelPluginThumbnailCdnUrl          interface{}                   `json:"botReelPluginThumbnailCdnUrl"`
	BotMsgBodyType                        interface{}                   `json:"botMsgBodyType"`
	RequiresDirectConnection              interface{}                   `json:"requiresDirectConnection"`
	BizContentPlaceholderType             interface{}                   `json:"bizContentPlaceholderType"`
	HostedBizEncStateMismatch             bool                          `json:"hostedBizEncStateMismatch"`
	SenderOrRecipientAccountTypeHosted    bool                          `json:"senderOrRecipientAccountTypeHosted"`
	PlaceholderCreatedWhenAccountIsHosted bool                          `json:"placeholderCreatedWhenAccountIsHosted"`
	ChatId                                string                        `json:"chatId"`
	FromMe                                bool                          `json:"fromMe"`
	Sender                                WebhookMessageDataSender      `json:"sender"`
	Timestamp                             int                           `json:"timestamp"`
	Content                               string                        `json:"content"`
	IsGroupMsg                            bool                          `json:"isGroupMsg"`
	MediaData                             WebhookMessageDataMedia       `json:"mediaData"`
	Mimetype                              string                        `json:"mimetype"`
	Duration                              string                        `json:"duration"`
	Filehash                              string                        `json:"filehash"`
	EncFilehash                           string                        `json:"encFilehash"`
	Size                                  int                           `json:"size"`
	MediaKey                              string                        `json:"mediaKey"`
	MediaKeyTimestamp                     int                           `json:"mediaKeyTimestamp"`
	DirectPath                            string                        `json:"directPath"`
	DeprecatedMms3Url                     string                        `json:"deprecatedMms3Url"`
	Waveform                              map[string]int                `json:"waveform"`
	IsViewOnce                            bool                          `json:"isViewOnce"`
}

type WebhookMessageDataSender struct {
	Id                       string             `json:"id"`
	Pushname                 string             `json:"pushname"`
	Type                     string             `json:"type"`
	VerifiedName             string             `json:"verifiedName"`
	IsBusiness               bool               `json:"isBusiness"`
	IsEnterprise             bool               `json:"isEnterprise"`
	IsSmb                    bool               `json:"isSmb"`
	VerifiedLevel            int                `json:"verifiedLevel"`
	PrivacyMode              interface{}        `json:"privacyMode"`
	Labels                   []string           `json:"labels"`
	IsContactSyncCompleted   int                `json:"isContactSyncCompleted"`
	TextStatusLastUpdateTime int                `json:"textStatusLastUpdateTime"`
	SyncToAddressbook        bool               `json:"syncToAddressbook"`
	FormattedName            string             `json:"formattedName"`
	IsMe                     bool               `json:"isMe"`
	IsMyContact              bool               `json:"isMyContact"`
	IsPSA                    bool               `json:"isPSA"`
	IsUser                   bool               `json:"isUser"`
	IsWAContact              bool               `json:"isWAContact"`
	ProfilePicThumbObj       ProfilePicThumbObj `json:"profilePicThumbObj"`
	Msgs                     interface{}        `json:"msgs"`
}

type ProfilePicThumbObj struct {
	Eurl    string `json:"eurl"`
	Id      string `json:"id"`
	Img     string `json:"img"`
	ImgFull string `json:"imgFull"`
	Tag     string `json:"tag"`
}

type WebhookMessageDataMedia struct {
	Type                   string `json:"type"`
	MediaStage             string `json:"mediaStage"`
	AnimationDuration      int    `json:"animationDuration"`
	AnimatedAsNewMsg       bool   `json:"animatedAsNewMsg"`
	IsViewOnce             bool   `json:"isViewOnce"`
	SwStreamingSupported   bool   `json:"_swStreamingSupported"`
	ListeningToSwSupport   bool   `json:"_listeningToSwSupport"`
	IsVcardOverMmsDocument bool   `json:"isVcardOverMmsDocument"`
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
