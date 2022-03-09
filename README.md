# Harbor CLI
This is a primitive CLI to use with the Harbor container registry. All output is in json form, so can be passed through a tool such as [jq](https://stedolan.github.io/jq/) for better formatting. E.g `harbor projects listProjects | jq`

## Installation
`harbor-cli` is a go binary and can be downloaded from the [releases](https://github.com/laidbackware/harbor-cli/releases) page.</br>
On Linux/Mac the file just needs to be made executable.</br>
To make it available to all users `sudo mv harbor-cli-<os>-<version> /usr/local/bin/harbor-cli`

## Usage
```
Usage:
  harbor-cli [command]

Available Commands:
  artifact               
  auditlog               
  completion             Generate completion script
  configure              
  gc                     
  health                 
  help                   Help about any command
  icon                   
  immutable              
  label                  
  ldap                   
  member                 
  oidc                   
  ping                   
  preheat                
  project                
  project_metadata       
  quota                  
  registry               
  replication            
  repository             
  retention              
  robot                  
  robotv1                
  scan                   
  scan_all               
  scanner                
  search                 
  statistic              
  system_c_v_e_allowlist 
  systeminfo             
  user                   
  usergroup              
  webhook                
  webhookjob             

Flags:
      --config string     config file path
      --debug             output debug logs
      --dry-run           do not send the request to server
  -h, --help              help for main
      --hostname string   hostname of the service (default "localhost")
      --password string   password for basic auth
      --scheme string     Choose from: [http https] (default "http")
      --username string   username for basic auth

Use "harbro-cli [command] --help" for more information about a command.
```
## Known issues
`harbor-cli systeminfo getCert` is currently broken and does not return a response.

## Build
### Dependencies
- Golang
- go-swagger - https://goswagger.io/install.html

## Generation
On first run the directory my be initialized as a module `go mod init github.com/laidbackware/harbor-cli`
- wget https://raw.githubusercontent.com/goharbor/harbor/main/api/v2.0/swagger.yaml
- remove single quotes from single line stings `sed -i "/description/s/'//g" swagger.yaml`
- switch all double to single quotes `sed -i "s/\"/'/g" swagger.yaml`
- `swagger generate cli --target=. --spec=swagger.yaml --cli-app-name harbor-cli`
  
### Post Generation Fixes
- A number of files in /cli need `"github.com/laidbackware/harbor-cli/models"` added as an import
- Swap `m.Status = &statusFlagValues` to `m.Status = statusFlagValues` in cli/list_webhook_jobs_operation.go
- Import modules `go get -u -f ./...` and `go mod tidy`
- cli/get_cert_operation.go is broken and needed lines 42-50 commenting, plus `appCli, err` changing to `_, err`

## Build
Call the build script with the version number
```
./scripts/build.sh 0.0.1
```
