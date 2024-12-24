package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const configFileName = ".gatorconfig.json"

func Read() (Config, error){
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("cant get file path error: %v", err)
	}
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("cant open file error: %v", err)
	}
	defer file.Close()

	dat, err := io.ReadAll(file)
	if err != nil {
		return Config{}, fmt.Errorf("cant read file error: %v", err)
	}
	var jsonFile Config
	err = json.Unmarshal(dat, &jsonFile)
	if err != nil{
		return Config{}, fmt.Errorf("can't unmarshal data error: %v", err)
	}
	
	return jsonFile, nil
}

func getConfigFilePath()(string, error){
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting the home path: %v", err)
	}
	fullPath := homeDirectory + "/" +configFileName
	
	return fullPath, nil
}

func (c *Config) SetUser(userName string) error{
	c.Current_user_name = userName
	err := write(*c)
	if err != nil {
		return fmt.Errorf("cant set user error: %v", err)
	}
	//fmt.Println(c.Current_user_name)
	return nil
}

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("can't marshal config: %v", err)
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("cant get file path error: %v", err)
	}
	err = os.WriteFile(filePath, jsonData, 0755)
	if err != nil {
		return fmt.Errorf("cant write to file error: %v", err)
	}

	return nil
}

func (c *Config) SetUrl(db_url string) error {
	c.Db_url = db_url
	err := write(*c)
	if err != nil {
		return fmt.Errorf("can't set URL error: %v", err)
	}
	return nil
}
