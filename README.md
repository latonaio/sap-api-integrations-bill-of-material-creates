# sap-api-integrations-bill-of-material-creates
sap-api-integrations-bill-of-material-creates は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 部品表（BOM）データを登録するマイクロサービスです。    
sap-api-integrations-bill-of-material-creates には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-bill-of-material-creates は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_BILL_OF_MATERIAL_SRV_0002/overview  

## 動作環境  
sap-api-integrations-bill-of-material-creates は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-bill-of-material-creates は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-bill-of-material-creates が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_BILL_OF_MATERIAL_SRV_0002/overview
* APIサービス名(=baseURL): PI_BILL_OF_MATERIAL_SRV;v=0002

## 本レポジトリ に 含まれる API名
sap-api-integrations-bill-of-material-creates には、次の API をコールするためのリソースが含まれています。  

* MaterialBOM（部品表）
## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめ登録することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "SAPBillOfMaterialCreates",
	"accepter": ["Header"],
	"bill_of_material": "",
	"deleted": false
```
  
* 全データを登録する際のsample.jsonの記載例(2)  

全データを登録する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SAPXBillOfMaterialCreates",
	"accepter": ["All"],
	"bill_of_material": "",
	"deleted": false
```
## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncPostBillOfMaterial(
	headerItem *requests.HeaderItem,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "HeaderItem":
			func() {
				c.HeaderItem(headerItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 部品表 の ヘッダデータ が登録された結果の JSON の例です。  
以下の項目のうち、"BillOfMaterial" ～ "to_BillOfMaterialItem" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
XXXXXXXXXXXX
XXXXXXXXXXXX
```