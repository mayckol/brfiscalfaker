package brfiscalfaker

import (
	"fmt"
	"github.com/mayckol/brfiscalfaker/pkg/br_documents"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func init() {
	gofakeit.Seed(time.Now().UnixNano())
}

// xNome generates a mock name.
func xNome() string {
	return gofakeit.Company()
}

// xFant generates a mock trade name.
func xFant() string {
	return gofakeit.CompanySuffix()
}

// xLgr generates a mock street name.
func xLgr() string {
	return gofakeit.Street()
}

// nro generates a mock number.
func nro() string {
	return string(rune(gofakeit.Number(1, 9999)))
}

// xCpl generates a mock complement.
func xCpl() string {
	return gofakeit.Word()
}

// xBairro generates a mock neighborhood.
func xBairro() string {
	return gofakeit.City()
}

// cMun generates a mock municipality code.
func cMun() string {
	return strconv.Itoa(gofakeit.Number(1000000, 9999999))
}

// xMun generates a mock municipality name.
func xMun() string {
	return gofakeit.City()
}

// UF generates a mock state abbreviation.
func UF() string {
	return gofakeit.StateAbr()
}

// CEP generates a mock postal code.
func CEP() string {
	return gofakeit.Zip()
}

// cPais generates a mock country code.
func cPais() string {
	return "1058" // Brazil's country code
}

// xPais generates a mock country name.
func xPais() string {
	return "BRASIL"
}

// fone generates a mock phone number.
func fone() string {
	return gofakeit.Phone()
}

// IE generates a mock State Registration.
func IE() string {
	return string(rune(gofakeit.Number(8, 12)))
}

// CRT generates a mock Tax Regime code.
func CRT() string {
	return string(rune(gofakeit.Number(1, 3)))
}

// indIEDest generates a mock indicator.
func indIEDest() string {
	return string(rune(gofakeit.Number(0, 2)))
}

// email generates a mock email address.
func email() string {
	return gofakeit.Email()
}

// nItem generates a mock item number.
func nItem() string {
	return strconv.Itoa(gofakeit.Number(1, 999))
}

// cProd generates a mock product code.
func cProd() string {
	return fmt.Sprintf("%s.%s.%s", gofakeit.Numerify("##.##.#########"), gofakeit.Word(), gofakeit.Numerify("####"))
}

// cEAN generates a mock EAN code.
func cEAN() string {
	return gofakeit.RandomString([]string{"789", "790", "791", "792", "793", "794", "795", "796", "797", "798", "799"})
}

// xProd generates a mock product name.
func xProd() string {
	return gofakeit.ProductName()
}

// NCM generates a mock NCM code.
func NCM() string {
	return gofakeit.Numerify("########")
}

// CFOP generates a mock CFOP code.
func CFOP() string {
	return gofakeit.RandomString([]string{"abc", "def", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yz"})
}

// uCom generates a mock unit of measure.
func uCom() string {
	return "PC" // Example: PC, UN, KG, etc.
}

// qCom generates a mock quantity.
func qCom() string {
	return fmt.Sprintf("%.4f", gofakeit.Float64Range(0.1, 100.0))
}

// vUnCom generates a mock unit value.
func vUnCom() string {
	return fmt.Sprintf("%.10f", gofakeit.Float64Range(0.01, 1000.0))
}

// vProd generates a mock product value.
func vProd() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.01, 10000.0))
}

// cEANTrib generates a mock EAN code for taxation.
func cEANTrib() string {
	return gofakeit.RandomString([]string{"789", "790", "791", "792", "793", "794", "795", "796", "797", "798", "799"})
}

// uTrib generates a mock unit of taxation.
func uTrib() string {
	return "PC" // Example: PC, UN, KG, etc.
}

// qTrib generates a mock quantity for taxation.
func qTrib() string {
	return fmt.Sprintf("%.4f", gofakeit.Float64Range(0.1, 100.0))
}

// vUnTrib generates a mock unit value for taxation.
func vUnTrib() string {
	return fmt.Sprintf("%.4f", gofakeit.Float64Range(0.01, 1000.0))
}

// vDesc generates a mock discount value.
func vDesc() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 1000.0))
}

// indTot generates a mock indicator for total.
func indTot() string {
	return gofakeit.RandomString([]string{"0", "1"})
}

// vTotTrib generates a mock total tributary value.
func vTotTrib() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 1000.0))
}

// orig generates a mock origin code.
func orig() string {
	return gofakeit.RandomString([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
}

// CSOSN generates a mock CSOSN code.
func CSOSN() string {
	return "102" // Example: 102, 300, etc.
}

// CST_PIS generates a mock CST code for PIS.
func CST_PIS() string {
	return "49" // Example: 49, 50, etc.
}

// vBC_PIS generates a mock PIS base value.
func vBC_PIS() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// pPIS generates a mock PIS percentage.
func pPIS() string {
	return fmt.Sprintf("%.4f", gofakeit.Float64Range(0.0, 100.0))
}

// vPIS generates a mock PIS value.
func vPIS() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// CST_COFINS generates a mock CST code for COFINS.
func CST_COFINS() string {
	return "49" // Example: 49, 50, etc.
}

// vBC_COFINS generates a mock COFINS base value.
func vBC_COFINS() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// pCOFINS generates a mock COFINS percentage.
func pCOFINS() string {
	return fmt.Sprintf("%.4f", gofakeit.Float64Range(0.0, 100.0))
}

// vCOFINS generates a mock COFINS value.
func vCOFINS() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// infAdProd generates a mock additional product information.
func infAdProd() string {
	return gofakeit.Sentence(10)
}

// modFrete generates a mock freight mode.
func modFrete() string {
	return gofakeit.RandomString([]string{"0", "1", "2", "3", "4", "9"})
}

// tPag generates a mock payment type.
func tPag() string {
	return "03" // Example: 01, 02, 03, etc.
}

// vPag generates a mock payment value.
func vPag() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// tpIntegra generates a mock integration type.
func tpIntegra() string {
	return "1" // Example: 0 or 1
}

// cardCNPJ generates a mock CNPJ for card.
func cardCNPJ() string {
	return br_documents.CNPJ()
}

// tBand generates a mock brand type.
func tBand() string {
	return "02" // Example: 01, 02, etc.
}

// cAut generates a mock authorization code.
func cAut() string {
	return gofakeit.Numerify("######")
}

// qrCode generates a mock QR Code URL.
func qrCode() string {
	return gofakeit.URL()
}

// urlChave generates a mock URL for chave.
func urlChave() string {
	return "www.nfce.fazenda.rj.gov.br/consulta"
}

// DigestValue generates a mock digest value.
func DigestValue() string {
	return gofakeit.RandomString([]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"})
}

// SignatureValue generates a mock signature value.
func SignatureValue() string {
	return gofakeit.UUID()
}

// X509Certificate generates a mock X509 certificate.
func X509Certificate() string {
	return gofakeit.UUID()
}

// tpAmbProt generates a mock environment type.
func tpAmbProt() string {
	return gofakeit.RandomString([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"})
}

// verAplic generates a mock application version.
func verAplic() string {
	return gofakeit.UUID()
}

// chNFe generates a mock NFe key.
func chNFe() string {
	return gofakeit.UUID()
}

// dhRecbto generates a mock receipt date.
func dhRecbto() string {
	return gofakeit.Date().String()
}

// nProt generates a mock protocol number.
func nProt() string {
	return gofakeit.UUID()
}

// digVal generates a mock digest value.
func digVal() string {
	return gofakeit.UUID()
}

// cStat generates a mock status code.
func cStat() string {
	return gofakeit.UUID()
}

// xMotivo generates a mock reason.
func xMotivo() string {
	return gofakeit.Sentence(10)
}

// vBC_total generates a mock total base value.
func vBC_total() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vICMS_total generates a mock total ICMS value.
func vICMS_total() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vICMSDeson generates a mock ICMS deson value.
func vICMSDeson() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vFCP generates a mock FCP value.
func vFCP() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vBCST generates a mock base value for ST.
func vBCST() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vST generates a mock ST value.
func vST() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vFCPST generates a mock FCP ST value.
func vFCPST() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vFCPSTRet generates a mock FCP ST Ret value.
func vFCPSTRet() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vProd_total generates a mock total product value.
func vProd_total() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vFrete generates a mock freight value.
func vFrete() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vSeg generates a mock insurance value.
func vSeg() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vDesc_total generates a mock total discount value.
func vDesc_total() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vII generates a mock II value.
func vII() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vIPI generates a mock IPI value.
func vIPI() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vIPIDevol generates a mock IPI devol value.
func vIPIDevol() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vPIS_total generates a mock total PIS value.
func vPIS_total() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vCOFINS_total generates a mock total COFINS value.
func vCOFINS_total() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vOutro generates a mock other value.
func vOutro() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vNF generates a mock NF value.
func vNF() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// vTotTrib_total generates a mock total tributary value.
func vTotTrib_total() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000.0))
}

// cUF generates a mock UF code.
func cUF() string {
	return gofakeit.RandomString([]string{"12", "27", "13", "16", "29", "23", "53", "32", "52", "21", "31", "50", "51", "15", "25", "26", "22", "41", "33", "24", "43", "11", "14", "42", "28", "35", "17", "29"})
}

// Number generates a mock number within a specified range.
func Number(min, max int) string {
	return strconv.Itoa(gofakeit.Number(min, max))
}

// NatOp generates a mock nature of operation.
func NatOp() string {
	return gofakeit.Sentence(3)
}

// nNF generates a mock NF number.
func nNF() string {
	return Number(1, 999999999)
}

// DhEmi generates a mock emission date.
func DhEmi() string {
	return gofakeit.Date().Format("2006-01-02T15:04:05-07:00")
}

// tpNF generates a mock NF type.
func tpNF() string {
	return gofakeit.RandomString([]string{"0", "1"})
}

// idDest generates a mock destination ID.
func idDest() string {
	return gofakeit.RandomString([]string{"1", "2", "3"})
}

// tpImp generates a mock print type.
func tpImp() string {
	return gofakeit.RandomString([]string{"0", "1", "2", "3", "4"})
}

// tpEmis generates a mock emission type.
func tpEmis() string {
	return gofakeit.RandomString([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"})
}

// cDV generates a mock check digit.
func cDV() string {
	return Number(0, 9)
}

// tpAmb generates a mock environment type.
func tpAmb() string {
	return gofakeit.RandomString([]string{"1", "2"})
}

// finNFe generates a mock NF purpose.
func finNFe() string {
	return gofakeit.RandomString([]string{"1", "2", "3", "4"})
}

// indFinal generates a mock final consumer indicator.
func indFinal() string {
	return gofakeit.RandomString([]string{"0", "1"})
}

// indPres generates a mock presence indicator.
func indPres() string {
	return gofakeit.RandomString([]string{"0", "1", "2", "3", "4", "9"})
}

// indIntermed generates a mock intermediary indicator.
func indIntermed() string {
	return gofakeit.RandomString([]string{"0", "1"})
}

// cNF generates a mock CFOP code.
func cNF() string {
	return fmt.Sprintf("%09d", gofakeit.Number(100000000, 999999999))
}

// natOp generates a mock nature of operation.
func natOp() string {
	return gofakeit.RandomString([]string{
		"Venda a vista",
		"Venda a prazo",
		"Devolucao de venda",
		"Compra para revenda",
	})
}

// indPag generates a mock payment indicator.
func indPag() string {
	return fmt.Sprintf("%d", gofakeit.Number(0, 1)) // 0 = Pagamento à vista, 1 = Pagamento a prazo
}

// serie generates a mock series number.
func serie() string {
	return fmt.Sprintf("%d", gofakeit.Number(1, 999))
}

// dEmi generates a mock emission date in YYYY-MM-DD format.
func dEmi() string {
	return gofakeit.Date().Format("2006-01-02")
}

// dSaiEnt generates a mock exit date in YYYY-MM-DD format.
func dSaiEnt() string {
	return dEmi() // For simplicity, using the same as emission date
}

// cMunFG generates a mock municipality code.
func cMunFG() string {
	return fmt.Sprintf("%07d", gofakeit.Number(1000000, 9999999))
}

// procEmi generates a mock process of emission.
func procEmi() string {
	return fmt.Sprintf("%d", gofakeit.Number(0, 6)) // 0 = Emissão de NF-e pelo contribuinte
}

// verProc generates a mock process version.
func verProc() string {
	return gofakeit.Word() // Example: "NF-eletronica.com"
}

// emitCNPJ generates a mock Brazilian CNPJ.
func emitCNPJ() string {
	return br_documents.CNPJ()
}

// emitXNome generates a mock company name.
func emitXNome() string {
	return gofakeit.Company()
}

// emitXFant generates a mock trade name.
func emitXFant() string {
	return gofakeit.CompanySuffix()
}

// enderEmitXLgr generates a mock street name for emitter's address.
func enderEmitXLgr() string {
	return gofakeit.Street()
}

// enderEmitNro generates a mock street number for emitter's address.
func enderEmitNro() string {
	return fmt.Sprintf("%d", gofakeit.Number(1, 9999))
}

// enderEmitXCpl generates a mock complement for emitter's address.
func enderEmitXCpl() string {
	return gofakeit.Word()
}

// enderEmitXBairro generates a mock neighborhood for emitter's address.
func enderEmitXBairro() string {
	return gofakeit.Address().State
}

// enderEmitCMun generates a mock municipality code for emitter's address.
func enderEmitCMun() string {
	return fmt.Sprintf("%07d", gofakeit.Number(1000000, 9999999))
}

// enderEmitXMun generates a mock municipality name for emitter's address.
func enderEmitXMun() string {
	return gofakeit.City()
}

// enderEmitUF generates a mock state abbreviation for emitter's address.
func enderEmitUF() string {
	return gofakeit.StateAbr()
}

// enderEmitCEP generates a mock postal code for emitter's address.
func enderEmitCEP() string {
	return gofakeit.Zip()
}

// enderEmitCPais generates a mock country code for emitter's address.
func enderEmitCPais() string {
	return "1058" // Brazil's country code
}

// enderEmitXPais generates a mock country name for emitter's address.
func enderEmitXPais() string {
	return "BRASIL"
}

// enderEmitFone generates a mock phone number for emitter's address.
func enderEmitFone() string {
	return gofakeit.PhoneFormatted()
}

// emitIE generates a mock State Registration for emitter.
func emitIE() string {
	return gofakeit.RandomString([]string{"ISENTO", "ISENTA"})
}

// destCNPJ generates a mock Brazilian CNPJ for destination.
func destCNPJ() string {
	return br_documents.CNPJ()
}

// destXNome generates a mock company name for destination.
func destXNome() string {
	return gofakeit.Company()
}

// enderDestXLgr generates a mock street name for destination's address.
func enderDestXLgr() string {
	return gofakeit.Street()
}

// enderDestNro generates a mock street number for destination's address.
func enderDestNro() string {
	return fmt.Sprintf("%d", gofakeit.Number(1, 9999))
}

// enderDestXCpl generates a mock complement for destination's address.
func enderDestXCpl() string {
	return gofakeit.Word()
}

// enderDestXBairro generates a mock neighborhood for destination's address.
func enderDestXBairro() string {
	return gofakeit.Address().State
}

// enderDestCMun generates a mock municipality code for destination's address.
func enderDestCMun() string {
	return fmt.Sprintf("%07d", gofakeit.Number(1000000, 9999999))
}

// enderDestXMun generates a mock municipality name for destination's address.
func enderDestXMun() string {
	return gofakeit.City()
}

// enderDestUF generates a mock state abbreviation for destination's address.
func enderDestUF() string {
	return gofakeit.StateAbr()
}

// enderDestCEP generates a mock postal code for destination's address.
func enderDestCEP() string {
	return gofakeit.Zip()
}

// enderDestCPais generates a mock country code for destination's address.
func enderDestCPais() string {
	return "1058" // Brazil's country code
}

// enderDestXPais generates a mock country name for destination's address.
func enderDestXPais() string {
	return "BRASIL"
}

// enderDestFone generates a mock phone number for destination's address.
func enderDestFone() string {
	return gofakeit.PhoneFormatted()
}

// destIE generates a mock State Registration for destination.
func destIE() string {
	return gofakeit.RandomString([]string{"ISENTO", "ISENTA"})
}

// retiradaCNPJ generates a mock CNPJ for retirada.
func retiradaCNPJ() string {
	return br_documents.CNPJ()
}

// retiradaXLgr generates a mock street name for retirada.
func retiradaXLgr() string {
	return gofakeit.Street()
}

// retiradaNro generates a mock street number for retirada.
func retiradaNro() string {
	return fmt.Sprintf("%d", gofakeit.Number(1, 99999))
}

// retiradaXCpl generates a mock complement for retirada.
func retiradaXCpl() string {
	return gofakeit.Word()
}

// retiradaXBairro generates a mock neighborhood for retirada.
func retiradaXBairro() string {
	return gofakeit.Address().State
}

// retiradaCMun generates a mock municipality code for retirada.
func retiradaCMun() string {
	return fmt.Sprintf("%07d", gofakeit.Number(1000000, 9999999))
}

// retiradaXMun generates a mock municipality name for retirada.
func retiradaXMun() string {
	return gofakeit.City()
}

// retiradaUF generates a mock state abbreviation for retirada.
func retiradaUF() string {
	return gofakeit.StateAbr()
}

// entregaCNPJ generates a mock CNPJ for entrega.
func entregaCNPJ() string {
	return br_documents.CNPJ()
}

// entregaXLgr generates a mock street name for entrega.
func entregaXLgr() string {
	return gofakeit.Street()
}

// entregaNro generates a mock street number for entrega.
func entregaNro() string {
	return fmt.Sprintf("%d", gofakeit.Number(1, 9999))
}

// entregaXCpl generates a mock complement for entrega.
func entregaXCpl() string {
	return gofakeit.Word()
}

// entregaXBairro generates a mock neighborhood for entrega.
func entregaXBairro() string {
	return gofakeit.Address().State
}

// entregaCMun generates a mock municipality code for entrega.
func entregaCMun() string {
	return fmt.Sprintf("%07d", gofakeit.Number(1000000, 9999999))
}

// entregaXMun generates a mock municipality name for entrega.
func entregaXMun() string {
	return gofakeit.City()
}

// entregaUF generates a mock state abbreviation for entrega.
func entregaUF() string {
	return gofakeit.StateAbr()
}

// detNItem generates a mock item number for det.
func detNItem() string {
	return fmt.Sprintf("%d", gofakeit.Number(1, 100))
}

// detProdCProd generates a mock product code for det.
func detProdCProd() string {
	return fmt.Sprintf("%05d", gofakeit.Number(1, 99999))
}

// detProdCEAN generates a mock EAN code for det. Can be empty.
func detProdCEAN() string {
	if gofakeit.Bool() {
		return gofakeit.RandomString([]string{"789", "790", "791", "792", "793", "794", "795", "796", "797", "798", "799"})
	}
	return ""
}

// detProdXProd generates a mock product name for det.
func detProdXProd() string {
	return gofakeit.ProductName()
}

// detProdCFOP generates a mock CFOP code for det.
func detProdCFOP() string {
	return fmt.Sprintf("%04d", gofakeit.Number(5101, 5999)) // Example range
}

// detProdUCom generates a mock unit of measure for det.
func detProdUCom() string {
	return gofakeit.RandomString([]string{"PC", "UN", "KG", "LT", "CX", "MT", "M2", "M3", "SC", "FD", "PCT", "GL", "FR", "BD", "PT", "KG", "M", "M2", "M3", "CM", "CM2", "CM3", "KG", "G", "MG", "L"})
}

// detProdQCom generates a mock quantity for det.
func detProdQCom() string {
	return fmt.Sprintf("%.4f", gofakeit.Float64Range(0.0001, 1000000.0))
}

// detProdVUnCom generates a mock unit value for det.
func detProdVUnCom() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.01, 1000.0))
}

// detProdVProd generates a mock product value for det.
func detProdVProd() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.01, 10000000.0))
}

// detProdCEANTrib generates a mock EAN code for taxation in det. Can be empty.
func detProdCEANTrib() string {
	if gofakeit.Bool() {
		return gofakeit.RandomString([]string{"789", "790", "791", "792", "793", "794", "795", "796", "797", "798", "799"})
	}
	return ""
}

// detProdUTrib generates a mock unit of taxation for det.
func detProdUTrib() string {
	return gofakeit.UUID()
}

// detProdQTrib generates a mock quantity of taxation for det.
func detProdQTrib() string {
	return fmt.Sprintf("%.4f", gofakeit.Float64Range(0.0001, 12000000.0))
}

// detProdVUnTrib generates a mock unit value of taxation for det.
func detProdVUnTrib() string {
	return fmt.Sprintf("%.4f", gofakeit.Float64Range(0.01, 1000.0))
}

// impostoICMS00orig generates a mock origin code for ICMS00.
func impostoICMS00orig() string {
	return fmt.Sprintf("%d", gofakeit.Number(0, 3)) // 0 = Nacional, 1 = Estrangeira - Importação Direta, etc.
}

// impostoICMS00CST generates a mock CST code for ICMS00.
func impostoICMS00CST() string {
	return "00" // Example: 00 = Tributado integralmente
}

// impostoICMS00modBC generates a mock BC model for ICMS00.
func impostoICMS00modBC() string {
	return fmt.Sprintf("%d", gofakeit.Number(0, 3)) // 0 = Margem Valor Agregado, 1 = Pauta, etc.
}

// impostoICMS00vBC generates a mock base value for ICMS00.
func impostoICMS00vBC() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000000.0))
}

// impostoICMS00pICMS generates a mock ICMS percentage for ICMS00.
func impostoICMS00pICMS() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 100.0))
}

// impostoICMS00vICMS generates a mock ICMS value for ICMS00.
func impostoICMS00vICMS() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 1800000.0))
}

// impostoPISAliqCST generates a mock CST code for PISAliq.
func impostoPISAliqCST() string {
	return "01" // Example: 01 = Operação Tributável
}

// impostoPISAliqvBC generates a mock base value for PISAliq.
func impostoPISAliqvBC() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000000.0))
}

// impostoPISAliqpPIS generates a mock PIS percentage for PISAliq.
func impostoPISAliqpPIS() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 100.0))
}

// impostoPISAliqvPIS generates a mock PIS value for PISAliq.
func impostoPISAliqvPIS() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 130000.0))
}

// impostoCOFINSAliqCST generates a mock CST code for COFINSAliq.
func impostoCOFINSAliqCST() string {
	return "01" // Example: 01 = Operação Tributável
}

// impostoCOFINSAliqvBC generates a mock base value for COFINSAliq.
func impostoCOFINSAliqvBC() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 10000000.0))
}

// impostoCOFINSAliqpCOFINS generates a mock COFINS percentage for COFINSAliq.
func impostoCOFINSAliqpCOFINS() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 100.0))
}

// impostoCOFINSAliqvCOFINS generates a mock COFINS value for COFINSAliq.
func impostoCOFINSAliqvCOFINS() string {
	return fmt.Sprintf("%.2f", gofakeit.Float64Range(0.0, 400000.0))
}

// totalICMSTotvBC generates a mock total base value for ICMSTot.
func totalICMSTotvBC() string {
	return "20000000.00" // Example fixed value
}

// totalICMSTotvICMS generates a mock total ICMS value for ICMSTot.
func totalICMSTotvICMS() string {
	return "18.00" // Example fixed value
}

// totalICMSTotvBCST generates a mock total BC ST value for ICMSTot.
func totalICMSTotvBCST() string {
	return "0"
}

// totalICMSTotvST generates a mock total ST value for ICMSTot.
func totalICMSTotvST() string {
	return "0"
}

// totalICMSTotvProd generates a mock total product value for ICMSTot.
func totalICMSTotvProd() string {
	return "20000000.00" // Example fixed value
}

// totalICMSTotvFrete generates a mock total freight value for ICMSTot.
func totalICMSTotvFrete() string {
	return "0"
}

// totalICMSTotvSeg generates a mock total insurance value for ICMSTot.
func totalICMSTotvSeg() string {
	return "0"
}

// totalICMSTotvDesc generates a mock total discount value for ICMSTot.
func totalICMSTotvDesc() string {
	return "0"
}

// totalICMSTotvII generates a mock total II value for ICMSTot.
func totalICMSTotvII() string {
	return "0"
}

// totalICMSTotvIPI generates a mock total IPI value for ICMSTot.
func totalICMSTotvIPI() string {
	return "0"
}

// totalICMSTotvPIS generates a mock total PIS value for ICMSTot.
func totalICMSTotvPIS() string {
	return "130000.00"
}

// totalICMSTotvCOFINS generates a mock total COFINS value for ICMSTot.
func totalICMSTotvCOFINS() string {
	return "400000.00"
}

// totalICMSTotvOutro generates a mock total outro value for ICMSTot.
func totalICMSTotvOutro() string {
	return "0"
}

// totalICMSTotvNF generates a mock total NF value for ICMSTot.
func totalICMSTotvNF() string {
	return "20000000.00"
}

// transpModFrete generates a mock freight mode.
func transpModFrete() string {
	return fmt.Sprintf("%d", gofakeit.Number(0, 9)) // 0 = Sem Frete, 1 = Por conta do Emitente, etc.
}

// transpTransportaCNPJ generates a mock CNPJ for transportadora.
func transpTransportaCNPJ() string {
	return br_documents.CNPJ()
}

// transpTransportaXNome generates a mock company name for transportadora.
func transpTransportaXNome() string {
	return gofakeit.Company()
}

// transpTransportaIE generates a mock State Registration for transportadora.
func transpTransportaIE() string {
	return gofakeit.RandomString([]string{"ISENTO", "ISENTA", "ISENTADO", "ISENTADO"})
}

// transpTransportaXEnder generates a mock address for transportadora.
func transpTransportaXEnder() string {
	return fmt.Sprintf("%s %d - %s - %s", gofakeit.Street(), gofakeit.Number(1, 9999), gofakeit.City(), gofakeit.StateAbr())
}

// transpTransportaXMun generates a mock municipality name for transportadora.
func transpTransportaXMun() string {
	return gofakeit.City()
}

// transpTransportaUF generates a mock state abbreviation for transportadora.
func transpTransportaUF() string {
	return gofakeit.StateAbr()
}

// transpVeicTranspPlaca generates a mock vehicle plate.
func transpVeicTranspPlaca() string {
	return gofakeit.Car().Type
}

// transpVeicTranspUF generates a mock state abbreviation for vehicle.
func transpVeicTranspUF() string {
	return gofakeit.StateAbr()
}

// transpVeicTranspRNTC generates a mock RNTC code for vehicle.
func transpVeicTranspRNTC() string {
	return fmt.Sprintf("%d", gofakeit.Number(100000000, 999999999))
}

// transpReboquePlaca generates a mock reboque (trailer) plate.
func transpReboquePlaca() string {
	return gofakeit.Car().Brand
}

// transpReboqueUF generates a mock state abbreviation for reboque.
func transpReboqueUF() string {
	return gofakeit.StateAbr()
}

// transpReboqueRNTC generates a mock RNTC code for reboque.
func transpReboqueRNTC() string {
	return fmt.Sprintf("%d", gofakeit.Number(100000000, 999999999))
}

// transpVolQVol generates a mock volume quantity.
func transpVolQVol() string {
	return "10000" // Example fixed value
}

// transpVolEsp generates a mock volume specification.
func transpVolEsp() string {
	return "CAIXA"
}

// transpVolMarca generates a mock volume brand.
func transpVolMarca() string {
	return "LINDOYA"
}

// transpVolNVol generates a mock volume number.
func transpVolNVol() string {
	return "500"
}

// transpVolPesoL generates a mock volume net weight.
func transpVolPesoL() string {
	return "1000000000.000" // Example fixed value
}

// transpVolPesoB generates a mock volume gross weight.
func transpVolPesoB() string {
	return "1200000000.000" // Example fixed value
}

// transpVolLacresNLacre generates a mock lacre number.
func transpVolLacresNLacre() string {
	return gofakeit.Numerify("XYZ########")
}

// infAdicInfAdFisco generates a mock additional fiscal information.
func infAdicInfAdFisco() string {
	return "Nota Fiscal de exemplo NF-eletronica.com"
}
