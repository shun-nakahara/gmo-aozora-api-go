/*
 * GMO Aozora Net Bank Open API
 *
 * <p>オープンAPI仕様書（PDF版）は下記リンクをご参照ください</p> <div>   <div style='display:inline-block;'><a style='text-decoration:none; font-weight:bold; color:#00b8d4;' href='https://gmo-aozora.com/business/service/api-specification.html' target='_blank'>オープンAPI仕様書</a></div><div style='display:inline-block; margin-left:2px; left:2px; width:10px; height:10px; border-top:2px solid #00b8d4; border-right:2px solid #00b8d4; transparent;-webkit-transform:rotate(45deg); transform: rotate(45deg);'></div> </div> <h4 style='margin-top:30px; border-left: solid 4px #1B2F48; padding: 0.1em 0.5em; color:#1B2F48;'>共通仕様</h4> <div style='width:100%; margin:10px;'>   <p style='font-weight:bold; color:#616161;'>＜HTTPリクエストヘッダ＞</p>   <div style='display:table; margin-left:10px; background-color:#29659b;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff;'>項目</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; color:#fff;'>仕様</div>   </div>   <div style='display:table; margin-left:10px;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff; background-color:#29659b;'>プロトコル</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; background-color:#f8f8f8;'>HTTP1.1/HTTPS</div>   </div>   <div style='display:table; margin-left:10px;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff; background-color:#29659b;'>charset</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; background-color:#f8f8f8;'>UTF-8</div>   </div>   <div style='display:table; margin-left:10px;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff; background-color:#29659b;'>content-type</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; background-color:#f8f8f8;'>application/json</div>   </div>   <div style='display:table; margin-left:10px;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff; background-color:#29659b;'>domain_name</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; background-color:#f8f8f8;'>       本番環境：api.gmo-aozora.com</br>       開発環境：stg-api.gmo-aozora.com     </div>   </div>   <div style='display:table; margin-left:10px;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff; background-color:#29659b;'>メインURL</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; background-color:#f8f8f8;'>       https://{domain_name}/ganb/api/personal/{version}</br>       <span style='border-bottom:solid 1px;'>Version:1.x.x</span> の場合</br>       　https://api.gmo-aozora.com/ganb/api/personal/<span style='border-bottom:solid 1px;'>v1</span>     </div>   </div> </div> <div style='margin:20px 10px;'>   <p style='font-weight:bold; color:#616161;'>＜リクエスト共通仕様＞</p>   <p style='padding-left:20px; font-weight:bold; color:#616161;'>NULLデータの扱い</p>   <p style='padding-left:40px;'>パラメータの値が空の場合、またはパラメータ自体が設定されていない場合、どちらもNULLとして扱います</p> </div> <div style='margin:20px 10px;'>   <p style='font-weight:bold; color:#616161;'>＜レスポンス共通仕様＞</p>   <p style='padding-left:20px; font-weight:bold; color:#616161;'>NULLデータの扱い</p>   <ul>     <li>レスポンスデータ</li>       <ul>         <li style='list-style-type:none;'>レスポンスデータの値が空の場合または、レスポンスデータ自体が設定されない場合は「項目自体を設定しません」と記載</li>       </ul>     <li>配列</li>       <ul>         <li style='list-style-type:none;'>配列の要素の値が空の場合は「空のリスト」と記載</li>         <li style='list-style-type:none;'>配列自体が設定されない場合は「項目自体を設定しません」と記載</li>       </ul>   </ul> </div> <div style='margin:20px 10px;'>   <p style='font-weight:bold; color:#616161;'>＜更新系APIに関する注意事項＞</p>   <ul>     <li style='list-style-type:none;'>更新系処理がタイムアウトとなった場合、処理自体は実行されている可能性がありますので、</li>     <li style='list-style-type:none;'>再実行を行う必要がある場合は必ず照会系の処理で実行状況を確認してから再実行を行ってください</li>   </ul> </div> 
 *
 * API version: 1.1.12
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BulkTransferDetail struct {
	// 振込ステータス 半角数字 2:申請中、3:差戻、4:取下げ、5:期限切れ、8:承認取消/予約取消、 11:予約中、12:手続中、13:リトライ中、 20:手続済、30:不能・組戻あり、40:手続不成立 
	TransferStatus string `json:"transferStatus,omitempty"`
	// 振込ステータス名 全角文字 
	TransferStatusName string `json:"transferStatusName,omitempty"`
	// 種類 全角文字 総合振込　を表示 
	TransferTypeName string `json:"transferTypeName,omitempty"`
	// 会社コード(振込依頼人コード) 銀行側で番号を付与している場合のみ表示 該当する情報が無い場合は項目自体を設定しません 
	RemitterCode string `json:"remitterCode,omitempty"`
	// 振込無料回数利用可否 振込利用回数の利用可否（無料回数の利用可否の設定であり、実際の利用有無ではありません） 総合振込では無料回数は利用できないため、常にfalse 
	IsFeeFreeUse bool `json:"isFeeFreeUse,omitempty"`
	// ポイント利用可否 ポイント会社の利用可否 
	IsFeePointUse bool `json:"isFeePointUse,omitempty"`
	// ポイント会社名 全角文字 
	PointName string `json:"pointName,omitempty"`
	// 手数料後払区分 「true=後払い」以外の場合は項目自体を設定しません 
	FeeLaterPaymentFlg bool `json:"feeLaterPaymentFlg,omitempty"`
	// 合計手数料 半角数字 振り込み完了時以外は、予定の手数料 
	TotalFee string `json:"totalFee,omitempty"`
	// 合計出金金額 半角数字 手数料+振込金額　ただし、振込完了時以外は、予定の手数料 
	TotalDebitAmount string `json:"totalDebitAmount,omitempty"`
	// 振込申請情報 振込申請情報のリスト 
	TransferApplies []TransferApply `json:"transferApplies,omitempty"`
	// 振込受付情報 振込受付情報のリスト 該当する情報が無い場合は空のリストを返却 
	TransferAccepts []TransferAccept `json:"transferAccepts,omitempty"`
	// 総合振込レスポンス情報 総合振込レスポンス情報のリスト 該当する情報が無い場合は空のリストを返却 
	BulktransferResponses []BulkTransferResponse `json:"bulktransferResponses,omitempty"`
}
