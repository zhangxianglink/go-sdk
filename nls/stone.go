package nls

/**
appkey: str,
token: 密钥,
file_link: https文件路径,
callback_url: 回调地址,
open_ws: 实时 true/离线模式 false,
task_id= 任务编号，不传输服务器随机生成,
enable_semantic_sentence_detection: 符号,
enable_inverse_text_normalization: 中文转阿拉伯数字,
enable_hot_rule: 热词转换,
enable_cn_rule: 热词转换前后对比,
enable_db_rule: 分贝,
enable_speed_rule: 语速,
以下不确定可以不设置，服务器会采用默认最优策略：
enable_default_noise_pass 单/左声道 滤波器,
enable_default_noise_command      动态压缩,
enable_default_noise_preemphasis  预加重,
enable_default_noise_equalizer    均衡,
enable_default_noise_speed        加速,
enable_right_noise_pass  右声道，同上,
enable_right_noise_command=None,
enable_right_noise_preemphasis=None,
enable_right_noise_equalizer=None,
enable_right_noise_speed=None
*/

type VadParam struct {
	AppKey                          string `json:"appkey"`
	Token                           string `json:"token"`
	FileLink                        string `json:"file_link"`
	CallbackUrl                     string `json:"callback_url"`
	OpenWs                          bool   `json:"open_ws"`
	EnableCallback                  bool   `json:"enable_callback"`
	EnableKafkaRule                 bool   `json:"enable_kafka_rule"`
	TaskId                          string `json:"task_id"`
	MaxSingleSegmentTime            int    `json:"max_single_segment_time"`
	NoiseType                       int    `json:"noise_type"`
	EnableSemanticSentenceDetection bool   `json:"enable_semantic_sentence_detection"`
	EnableInverseTextNormalization  bool   `json:"enable_inverse_text_normalization"`
	EnableHotRule                   bool   `json:"enable_hot_rule"`
	EnableCnRule                    bool   `json:"enable_cn_rule"`
	EnableDbRule                    bool   `json:"enable_db_rule"`
	EnableSpeedRule                 bool   `json:"enable_speed_rule"`
	EnableDefaultNoisePass          bool   `json:"enable_default_noise_pass"`
	EnableDefaultNoiseCommand       bool   `json:"enable_default_noise_command"`
	EnableDefaultNoisePreemphasis   bool   `json:"enable_default_noise_preemphasis"`
	EnableDefaultNoiseEqualizer     bool   `json:"enable_default_noise_equalizer"`
	EnableDefaultNoiseSpeed         bool   `json:"enable_default_noise_speed"`
	EnableRightNoisePass            bool   `json:"enable_right_noise_pass"`
	EnableRightNoiseCommand         bool   `json:"enable_right_noise_command"`
	EnableRightNoisePreemphasis     bool   `json:"enable_right_noise_preemphasis"`
	EnableRightNoiseEqualizer       bool   `json:"enable_right_noise_equalizer"`
	EnableRightNoiseSpeed           bool   `json:"enable_right_noise_speed"`
}

func NewVadParam(appKey string, token string, fileLink string, callbackUrl string, openWs bool) *VadParam {
	return &VadParam{
		AppKey:               appKey,
		Token:                token,
		FileLink:             fileLink,
		CallbackUrl:          callbackUrl,
		OpenWs:               openWs,
		EnableCallback:       true,
		EnableKafkaRule:      false,
		MaxSingleSegmentTime: 500,
		NoiseType:            3,
	}
}
