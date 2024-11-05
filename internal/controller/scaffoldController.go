package controller

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/ylanzinhoy/internal/entity"
	typefield "github.com/ylanzinhoy/internal/typeField"
	txt "golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const entityTemplate = `package {{.Package}};



import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;

@Entity
public class {{.EntityName}} {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    {{range .Fields}}
    private {{.Type}} {{.Name}};
    {{end}}

    // Getters e setters
    {{range .Fields}}
    public {{.Type}} get{{.CapitalizedName}}() {
        return {{.Name}};
    }

    public void set{{.CapitalizedName}}({{.Type}} {{.Name}}) {
        this.{{.Name}} = {{.Name}};
    }
    {{end}}
}`

func ScaffoldController(args []string, packageFlag string) {

	entityName := args[0]
	fields := args[1:]

	// Processa os campos e extrai tipo e nome
	var parsedFields []entity.FieldEntity
	for _, field := range fields {
		parts := strings.Split(field, ":")
		if len(parts) != 2 {
			fmt.Printf("Formato de campo inválido: %s\n", field)
			return
		}
		parsedFields = append(parsedFields, entity.FieldEntity{
			Name:            parts[0],
			Type:            mapType(parts[1]),
			CapitalizedName: txt.Title(language.AmericanEnglish, txt.Compact).String(parts[0]),
		})
	}

	// Dados para preencher o template
	data := struct {
		Package    string
		EntityName string
		Fields     []entity.FieldEntity
	}{
		Package:    packageFlag, // Ajuste o nome do pacote conforme necessário
		EntityName: entityName,
		Fields:     parsedFields,
	}

	// Cria o arquivo Java
	fileName := fmt.Sprintf("%s.java", entityName)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Erro ao criar o arquivo: %v\n", err)
		return
	}
	defer file.Close()

	// Compila e executa o template
	tmpl, err := template.New("entity").Parse(entityTemplate)
	if err != nil {
		fmt.Printf("Erro ao parsear o template: %v\n", err)
		return
	}

	if err := tmpl.Execute(file, data); err != nil {
		fmt.Printf("Erro ao executar o template: %v\n", err)
		return
	}

	fmt.Printf("Arquivo %s gerado com sucesso!\n", fileName)

}

func mapType(fieldTypes string) string {
	switch fieldTypes {
	case typefield.String:
		return typefield.String
	case typefield.Integer:
		return typefield.Integer
	case typefield.Double:
		return typefield.Double
	case typefield.Float:
		return typefield.Float
	case typefield.Long:
		return typefield.Long
	default:
		return typefield.Object
	}

}
