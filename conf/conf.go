package conf
import ( "io/ioutil" "log" "gopkg.in/yaml.v2" )
type Conf struct 
{
	ChatLog       string `yaml:"chatlog"`
	ServerIP      string `yaml:"serverip"`
	Serverport    int    `yaml:"serverport"`
	ApiServerPort int    `yaml:"apiserverport"`
}
func (k *Conf) GetConf(fname string) *Conf 
{
	yamlFile, err := ioutil.ReadFile(fname)
	if err != nil { log.Printf("yamlFile.Get err   #%v ", err }
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil{ log.Fatalf("Unmarshal: %v", err }
	return k
}
