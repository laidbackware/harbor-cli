# Harbor CLI

## Dependencies
- Golang
- go-swagger - https://goswagger.io/install.html

## Generation
On first run the directory my be initialized as a module `go mod init github.com/laidbackware/harbor-cli`
- wget https://raw.githubusercontent.com/goharbor/harbor/master/api/v2.0/swagger.yaml
- remove single quotes from single line stings `sed -i "/description/s/'//g" swagger.yaml`
- switch all double to single quotes
- 
- fix quotes in string `sed -i "s/\"stop\" now/'stop' now/" swagger.yaml` and `sed -i "s/\"targets\" and \"event_types\"/'targets' and 'event_types'/" swagger.yaml`

### Post Generation Fixes
- A number of files in /cli need `"github.com/go-openapi/strfmt"` added as an import
- 