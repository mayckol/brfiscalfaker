<div align="center">
  <img src="logo.svg" alt="BrFiscalFaker Logo" width="400" height="400">
  <h1>BrFiscalFaker</h1>
</div>

## Overview

brfiscalfaker is a Go-based command-line tool designed to generate mock Brazilian fiscal invoices (NF-e, NFC-e, CFe, NFeDevolucao) for testing and development purposes. It allows users to create realistic invoice XML files with customizable data, facilitating the development of applications that interact with Brazilian fiscal systems.

## Features

- **Supports Multiple Invoice Types:** Generate NF-e, NFC-e, CFe, and NFeDevolucao invoices.
- **Customizable Data:** Provide custom CPF and CNPJ numbers.
- **Block Specific Tags:** Remove or block specific XML tags using the `--block-tags` flag.
- **Dependency Management:** Ensures dependent placeholders are processed in the correct order.
- **Cross-Platform:** Works seamlessly on various operating systems.
- **Comprehensive Logging:** Provides detailed logs for debugging and transparency.
- **Br documents:** This project includes utilities to generate random yet valid Brazilian fiscal identifiers such as Access Key (Chave de Acesso), CPF, and CNPJ. These are essential for creating mock data for testing purposes.
- **Unit Tested:** Robust unit tests to ensure reliability.

## Installation

Ensure you have Go installed on your system (version 1.16 or higher).

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/mayckol/brfiscalfaker.git
   cd brfiscalfaker
    ```
   
2. **Build the Application:**

   ```bash
   go build -o bfiscalfaker ./cmd/bfiscalfaker/main.go
   ```
   
## Usage

Run the application using the `go run command or the built binary.

### Command-Line Flags

- **`--cpf` (`optional`):** --cpf: (Optional) Provide a custom CPF number to include in the invoice.
- **`--cnpj` (`optional`):** --cnpj: (Optional) Provide a custom CNPJ number to include in the invoice.
- **`--block-tags` (`optional`):** --block-tags: (Optional) Block specific XML tags from being included in the invoice.
- **`--type` (`default NFCe`):** --type: (Optional) Specify the type of invoice to generate (NF-e, NFC-e, CFe, NFeDevolucao).

### Examples
* **Generate a Basic NFC-e Invoice:**

   ```bash
   go run cmd/bfiscalfaker/main.go
   ```
* **Generate an NF-e Invoice with Custom CPF and CNPJ:**

   ```bash
    go run cmd/bfiscalfaker/main.go --type NFe --cpf 12345678900 --cnpj 12345678901234
    ```
* **Generate a CFe Invoice with Blocked Tags:**

   ```bash
   go run cmd/bfiscalfaker/main.go --type CFe --block-tags "nItem, vProd"
   ```
* **Generate an NFCe Invoice with Custom Data and Blocked Tags:**

   ```bash
   go run cmd/bfiscalfaker/main.go --type NFCe --cpf 12345678900 --cnpj 12345678901234 --block-tags "nItem, vProd"
   ```

## Library Usage

### Download the Library

```bash
go get github.com/mayckol/brfiscalfaker
```
### Create a Template Generator
* **Random CNPJ, CPF, and Access Key Generation:**

   ```go
   package main
    import (
        "fmt"
        "github.com/mayckol/brfiscalfaker/pkg/br_documents"
    )
    func main() {
        cpf := br_documents.CPF()
        cnpj := br_documents.CNPJ()
        accessKey := br_documents.AccessKey()
        fmt.Println("CPF:", cpf)
        fmt.Println("CNPJ:", cnpj)
        fmt.Println("Access Key:", accessKey)
    }
    ```  
**Parameters Description:**

- **CPF()**
   - **Config**: Optional parameter to control output format.
      - `Masked`: `bool`. If `true`, the CPF is formatted as `XXX.XXX.XXX-XX`. Default is `false` (raw digits).

- **CNPJ()**
   - **Config**: Optional parameter to control output format.
      - `Masked`: `bool`. If `true`, the CNPJ is formatted as `XX.XXX.XXX/XXXX-XX`. Default is `false` (raw digits).

- **AccessKey()**
   - **Config**: Optional parameter to customize the generated key.
      - `Masked`: `bool`. If `true`, the Access Key is formatted with separators for readability. Default is `false` (raw digits).

```go
package main

import (
   "log"
   "github.com/mayckol/brfiscalfaker/pkg/nfs"
   "fmt"
)

func main() {
   // Define the desired template type
   templateType := nfs.NFCe

   // Create the generator
   generator, err := nfs.NewTemplateGenerator(templateType)
   if err != nil {
      log.Fatalf("Failed to create generator: %v", err)
   }
   options := []nfs.Option{
	   nfs.WithCPF("01234567890"),
	   nfs.WithCNPJ("12345678901234"),
   }

   // Generate the XML template
   xmlBytes, err := generator.Generate(options...)
   if err != nil {
      log.Fatalf("Failed to generate XML template: %v", err)
   }

   fmt.Println(string(xmlBytes))

}

```

## Docker

You can also run the application using Docker.
   ```shell
   docker pull mayckol/bfiscalfaker
   ```
   ```shell
   docker run mayckol/bfiscalfaker --type NFe --cpf 12345678900 --cnpj 12345678901234
   ```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to help improve this project.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.