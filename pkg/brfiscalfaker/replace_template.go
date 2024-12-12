package brfiscalfaker

import (
	"fmt"
	"github.com/mayckol/brfiscalfaker/pkg/br_documents"
	"log"
	"regexp"
	"strings"
)

type DependencyGraph map[string][]string

func topologicalSort(keys []string, dependencies DependencyGraph) ([]string, error) {
	sorted := []string{}
	visited := make(map[string]bool)
	temp := make(map[string]bool)

	var visit func(string) error
	visit = func(n string) error {
		if temp[n] {
			return fmt.Errorf("circular dependency detected at %s", n)
		}
		if !visited[n] {
			temp[n] = true
			for _, m := range dependencies[n] {
				if err := visit(m); err != nil {
					return err
				}
			}
			temp[n] = false
			visited[n] = true
			sorted = append(sorted, n)
		}
		return nil
	}

	for _, key := range keys {
		if err := visit(key); err != nil {
			return nil, err
		}
	}

	return sorted, nil
}

// ReplaceTemplate takes an XML template and replaces placeholders with mock values.
// It handles dependencies between placeholders and removes entire tags for blocked placeholders,
// including any surrounding whitespace and newline characters to prevent blank lines.
func ReplaceTemplate(template string, options ...Option) ([]byte, error) {
	cfg := &generationConfig{}
	for _, opt := range options {
		opt(cfg)
	}

	// Regular expression to find placeholders in the form {%key%}
	re := regexp.MustCompile(`\{\%(\w+)%\}`)

	// Find all matches of the placeholders
	matches := re.FindAllStringSubmatch(template, -1)

	// Collect unique keys
	keySet := make(map[string]struct{})
	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		key := match[1]
		keySet[key] = struct{}{}
	}

	// Create a slice of keys
	var keys []string
	for key := range keySet {
		keys = append(keys, key)
	}

	// Define dependencies
	dependencies := DependencyGraph{
		"accessKey": {"emitCNPJ"},
		// Add more dependencies as needed
	}

	// Perform topological sort
	sortedKeys, err := topologicalSort(keys, dependencies)
	if err != nil {
		return nil, fmt.Errorf("error sorting keys: %v", err)
	}

	// Map to store generated values for each unique key
	replacements := make(map[string]string)

	// Iterate through all sorted keys and generate mock values
	for _, key := range sortedKeys {
		replacements[key] = generateMockValue(key, replacements, cfg)
	}

	// Replace all placeholders in the template with generated values
	result := template
	for key, value := range replacements {
		placeholder := fmt.Sprintf("{%%%s%%}", key)
		result = strings.ReplaceAll(result, placeholder, value)
	}

	// Remove entire XML tags that correspond to blocked placeholders,
	// including any surrounding whitespace and newline characters
	for _, blockedKey := range cfg.blockedPlaceholders {
		// Define a regex pattern to match the entire tag containing the blocked placeholder
		// The (?s) flag enables dot-all mode, allowing .*? to match newline characters
		// \s* ensures that any leading or trailing whitespace (including newlines) is captured
		tagPattern := fmt.Sprintf(`(?s)\s*<%s\b[^>]*>.*?</%s>\s*`, regexp.QuoteMeta(blockedKey), regexp.QuoteMeta(blockedKey))
		tagRe := regexp.MustCompile(tagPattern)
		result = tagRe.ReplaceAllString(result, "")
	}

	result = strings.TrimSpace(result)

	return []byte(result), nil
}

// generateMockValue generates mock data based on the placeholder key.
// It uses provided CPF/CNPJ if available.
func generateMockValue(key string, replacements map[string]string, cfg *generationConfig) string {
	switch key {
	case "accessKey":
		emitCNPJ, exists := replacements["emitCNPJ"]
		if !exists || emitCNPJ == "" {
			log.Printf("Warning: emitCNPJ not set before accessKey generation.")
			emitCNPJ = br_documents.CNPJ()
		}
		accessKey := br_documents.AccessKey(br_documents.AccessKeyConfig{
			CNPJ: emitCNPJ,
		})
		return accessKey
	case "cUF":
		return cUF()
	case "cNF":
		return Number(10000000, 99999999)
	case "natOp":
		return NatOp()
	case "serie":
		return Number(1, 999)
	case "nNF":
		return nNF()
	case "dhEmi":
		return DhEmi()
	case "tpNF":
		return tpNF()
	case "idDest":
		return idDest()
	case "cMunFG":
		return cMun()
	case "tpImp":
		return tpImp()
	case "tpEmis":
		return tpEmis()
	case "cDV":
		return cDV()
	case "tpAmb":
		return tpAmb()
	case "finNFe":
		return finNFe()
	case "indFinal":
		return indFinal()
	case "indPres":
		return indPres()
	case "indIntermed":
		return indIntermed()
	case "procEmi":
		return procEmi()
	case "verProc":
		return verProc()
	case "emitCNPJ", "CNPJ", "destCNPJ", "transpTransportaCNPJ", "cardCNPJ", "retiradaCNPJ", "entregaCNPJ":
		if cfg.CNPJ != "" {
			return cfg.CNPJ
		}
		return br_documents.CNPJ()
	case "emitXNome":
		return xNome()
	case "xLgr":
		return xLgr()
	case "nro":
		return nro()
	case "xCpl":
		return xCpl()
	case "xBairro":
		return xBairro()
	case "cMun":
		return cMun()
	case "xMun":
		return xMun()
	case "UF":
		return UF()
	case "CEP":
		return CEP()
	case "cPais":
		return cPais()
	case "xPais":
		return xPais()
	case "fone":
		return fone()
	case "IE":
		return IE()
	case "CRT":
		return CRT()
	case "CPF":
		if cfg.CPF != "" {
			return cfg.CPF
		}
		return br_documents.CPF()
	case "destXNome":
		return xNome()
	case "xLgrDest":
		return xLgr()
	case "nroDest":
		return nro()
	case "xCplDest":
		return xCpl()
	case "xBairroDest":
		return xBairro()
	case "cMunDest":
		return cMun()
	case "xMunDest":
		return xMun()
	case "UFDest":
		return UF()
	case "CEPDest":
		return CEP()
	case "cPaisDest":
		return cPais()
	case "xPaisDest":
		return xPais()
	case "foneDest":
		return fone()
	case "indIEDest":
		return indIEDest()
	case "email":
		return email()
	case "nItem":
		return nItem()
	case "cProd":
		return cProd()
	case "cEAN":
		return cEAN()
	case "xProd":
		return xProd()
	case "NCM":
		return NCM()
	case "CFOP":
		return CFOP()
	case "uCom":
		return uCom()
	case "qCom":
		return qCom()
	case "vUnCom":
		return vUnCom()
	case "vProd":
		return vProd()
	case "cEANTrib":
		return cEANTrib()
	case "uTrib":
		return uTrib()
	case "qTrib":
		return qTrib()
	case "vUnTrib":
		return vUnTrib()
	case "vDesc":
		return vDesc()
	case "indTot":
		return indTot()
	case "vTotTrib":
		return vTotTrib()
	case "orig":
		return orig()
	case "CSOSN":
		return CSOSN()
	case "CST_PIS":
		return CST_PIS()
	case "vBC_PIS":
		return vBC_PIS()
	case "pPIS":
		return pPIS()
	case "vPIS":
		return vPIS()
	case "CST_COFINS":
		return CST_COFINS()
	case "vBC_COFINS":
		return vBC_COFINS()
	case "pCOFINS":
		return pCOFINS()
	case "vCOFINS":
		return vCOFINS()
	case "infAdProd":
		return infAdProd()
	case "vBC_total":
		return vBC_total()
	case "vICMS_total":
		return vICMS_total()
	case "vICMSDeson":
		return vICMSDeson()
	case "vFCP":
		return vFCP()
	case "vBCST":
		return vBCST()
	case "vST":
		return vST()
	case "vFCPST":
		return vFCPST()
	case "vFCPSTRet":
		return vFCPSTRet()
	case "vProd_total":
		return vProd_total()
	case "vFrete":
		return vFrete()
	case "vSeg":
		return vSeg()
	case "vDesc_total":
		return vDesc_total()
	case "vII":
		return vII()
	case "vIPI":
		return vIPI()
	case "vIPIDevol":
		return vIPIDevol()
	case "vPIS_total":
		return vPIS_total()
	case "vCOFINS_total":
		return vCOFINS_total()
	case "vOutro":
		return vOutro()
	case "vNF":
		return vNF()
	case "vTotTrib_total":
		return vTotTrib_total()
	case "modFrete":
		return modFrete()
	case "tPag":
		return tPag()
	case "vPag":
		return vPag()
	case "tpIntegra":
		return tpIntegra()
	case "tBand":
		return tBand()
	case "cAut":
		return cAut()
	case "qrCode":
		return qrCode()
	case "urlChave":
		return urlChave()
	case "DigestValue":
		return DigestValue()
	case "SignatureValue":
		return SignatureValue()
	case "X509Certificate":
		return X509Certificate()
	case "tpAmbProt":
		return tpAmbProt()
	case "verAplic":
		return verAplic()
	case "chNFe":
		return chNFe()
	case "dhRecbto":
		return dhRecbto()
	case "nProt":
		return nProt()
	case "digVal":
		return digVal()
	case "cStat":
		return cStat()
	case "xMotivo":
		return xMotivo()
	case "transpTransportaXNome":
		return transpTransportaXNome()
	case "transpTransportaIE":
		return transpTransportaIE()
	case "transpTransportaXEnder":
		return transpTransportaXEnder()
	case "transpTransportaXMun":
		return transpTransportaXMun()
	case "transpTransportaUF":
		return transpTransportaUF()
	case "transpVeicTranspPlaca":
		return transpVeicTranspPlaca()
	case "transpVeicTranspUF":
		return transpVeicTranspUF()
	case "transpVeicTranspRNTC":
		return transpVeicTranspRNTC()
	case "transpReboquePlaca":
		return transpReboquePlaca()
	case "transpReboqueUF":
		return transpReboqueUF()
	case "transpReboqueRNTC":
		return transpReboqueRNTC()
	case "transpVolQVol":
		return transpVolQVol()
	case "transpVolEsp":
		return transpVolEsp()
	case "transpVolMarca":
		return transpVolMarca()
	case "transpVolNVol":
		return transpVolNVol()
	case "transpVolPesoL":
		return transpVolPesoL()
	case "transpVolPesoB":
		return transpVolPesoB()
	case "transpVolLacresNLacre":
		return transpVolLacresNLacre()
	case "infAdicInfAdFisco":
		return infAdicInfAdFisco()
	case "impostoCOFINSAliqCST":
		return impostoCOFINSAliqCST()
	case "impostoCOFINSAliqvBC":
		return impostoCOFINSAliqvBC()
	case "impostoCOFINSAliqpCOFINS":
		return impostoCOFINSAliqpCOFINS()
	case "impostoCOFINSAliqvCOFINS":
		return impostoCOFINSAliqvCOFINS()
	case "totalICMSTotvBC":
		return totalICMSTotvBC()
	case "totalICMSTotvICMS":
		return totalICMSTotvICMS()
	case "totalICMSTotvBCST":
		return totalICMSTotvBCST()
	case "totalICMSTotvST":
		return totalICMSTotvST()
	case "totalICMSTotvProd":
		return totalICMSTotvProd()
	case "totalICMSTotvFrete":
		return totalICMSTotvFrete()
	case "totalICMSTotvSeg":
		return totalICMSTotvSeg()
	case "totalICMSTotvDesc":
		return totalICMSTotvDesc()
	case "totalICMSTotvII":
		return totalICMSTotvII()
	case "totalICMSTotvIPI":
		return totalICMSTotvIPI()
	case "totalICMSTotvPIS":
		return totalICMSTotvPIS()
	case "totalICMSTotvCOFINS":
		return totalICMSTotvCOFINS()
	case "totalICMSTotvOutro":
		return totalICMSTotvOutro()
	case "totalICMSTotvNF":
		return totalICMSTotvNF()
	case "transpModFrete":
		return transpModFrete()
	case "emitXFant":
		return emitXFant()
	case "enderEmitXLgr":
		return enderEmitXLgr()
	case "enderEmitNro":
		return enderEmitNro()
	case "enderEmitXCpl":
		return enderEmitXCpl()
	case "enderEmitXBairro":
		return enderEmitXBairro()
	case "enderEmitCMun":
		return enderEmitCMun()
	case "enderEmitXMun":
		return enderEmitXMun()
	case "enderEmitUF":
		return enderEmitUF()
	case "enderEmitCEP":
		return enderEmitCEP()
	case "enderEmitCPais":
		return enderEmitCPais()
	case "enderEmitXPais":
		return enderEmitXPais()
	case "enderEmitFone":
		return enderEmitFone()
	case "emitIE":
		return emitIE()
	case "enderDestXLgr":
		return enderDestXLgr()
	case "enderDestNro":
		return enderDestNro()
	case "enderDestXCpl":
		return enderDestXCpl()
	case "enderDestXBairro":
		return enderDestXBairro()
	case "enderDestCMun":
		return enderDestCMun()
	case "enderDestXMun":
		return enderDestXMun()
	case "enderDestUF":
		return enderDestUF()
	case "enderDestCEP":
		return enderDestCEP()
	case "enderDestCPais":
		return enderDestCPais()
	case "enderDestXPais":
		return enderDestXPais()
	case "enderDestFone":
		return enderDestFone()
	case "destIE":
		return destIE()
	case "retiradaXLgr":
		return retiradaXLgr()
	case "retiradaNro":
		return retiradaNro()
	case "retiradaXCpl":
		return retiradaXCpl()
	case "retiradaXBairro":
		return retiradaXBairro()
	case "retiradaCMun":
		return retiradaCMun()
	case "retiradaXMun":
		return retiradaXMun()
	case "retiradaUF":
		return retiradaUF()
	case "entregaXLgr":
		return entregaXLgr()
	case "entregaNro":
		return entregaNro()
	case "entregaXCpl":
		return entregaXCpl()
	case "entregaXBairro":
		return entregaXBairro()
	case "entregaCMun":
		return entregaCMun()
	case "entregaXMun":
		return entregaXMun()
	case "entregaUF":
		return entregaUF()
	case "detNItem":
		return detNItem()
	case "detProdCProd":
		return detProdCProd()
	case "detProdCEAN":
		return detProdCEAN()
	case "detProdXProd":
		return detProdXProd()
	case "detProdCFOP":
		return detProdCFOP()
	case "detProdUCom":
		return detProdUCom()
	case "detProdQCom":
		return detProdQCom()
	case "detProdVUnCom":
		return detProdVUnCom()
	case "detProdVProd":
		return detProdVProd()
	case "detProdCEANTrib":
		return detProdCEANTrib()
	case "detProdUTrib":
		return detProdUTrib()
	case "detProdQTrib":
		return detProdQTrib()
	case "detProdVUnTrib":
		return detProdVUnTrib()
	case "impostoICMS00orig":
		return impostoICMS00orig()
	case "impostoICMS00CST":
		return impostoICMS00CST()
	case "impostoICMS00modBC":
		return impostoICMS00modBC()
	case "impostoICMS00vBC":
		return impostoICMS00vBC()
	case "impostoICMS00pICMS":
		return impostoICMS00pICMS()
	case "impostoICMS00vICMS":
		return impostoICMS00vICMS()
	case "impostoPISAliqCST":
		return impostoPISAliqCST()
	case "impostoPISAliqvBC":
		return impostoPISAliqvBC()
	case "impostoPISAliqpPIS":
		return impostoPISAliqpPIS()
	case "impostoPISAliqvPIS":
		return impostoPISAliqvPIS()
	case "indPag":
		return indPag()
	default:
		return ""
	}
}
