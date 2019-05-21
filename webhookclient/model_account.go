/*
 * GMO Aozora Net Bank Open API
 *
 * <p>Ph2.5向けに作成したもの</p> 
 *
 * API version: 1.1.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 入金口座情報 
type Account struct {
	// 入金口座ID 半角数字 入金先の口座を識別するID 
	RaId string `json:"raId"`
	// 入金口座　支店コード 半角数字 
	RaBranchCode string `json:"raBranchCode"`
	// 半角文字 
	RaBranchNameKana string `json:"raBranchNameKana"`
	// 半角数字 
	RaAccountNumber string `json:"raAccountNumber"`
	// 全角文字 
	RaHolderName string `json:"raHolderName"`
	// 基準日 半角文字 応答日付、もしくは入金明細の基準日を示す。 YYYY-MM-DD形式 
	BaseDate string `json:"baseDate"`
	// 基準時刻 半角文字 応答時刻、もしくは入金明細の基準時刻を示す。  ISO8601 時差(offset)も表記 HH:MM:SS+09:00形式 
	BaseTime string `json:"baseTime"`
}
