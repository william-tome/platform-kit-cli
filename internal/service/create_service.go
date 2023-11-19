package service

import (
	"github.com/william.tome/pk-cli/cmd/config"
	"github.com/william.tome/pk-cli/tools"
)

func CreateService(name string) error {
	servicesDir := "src/services/"
	subDirName := tools.HandleStringSuffix(name, "-services")
	serviceSubDir := servicesDir + subDirName
	serviceName := tools.HandleStringSuffix(name, "-service")
	serviceMethodName := tools.ConvertKebabToCamel(serviceName)
	servicePath := serviceSubDir + "/" + serviceName + ".js"
	makeUrlFileName := "make-" + name + "-base-url"
	makeUrlPath := serviceSubDir + "/" + makeUrlFileName + ".js"
	makeMethodName := tools.ConvertKebabToCamel(makeUrlFileName)
	indexPath := serviceSubDir + "/index.js"
	testSubDir := "src/tests/services/" + tools.HandleStringSuffix(name, "-services")
	testFilePath := testSubDir + "/" + serviceName + ".test.js"
	testServiceTitle := tools.ConvertKebabToPascal(subDirName)

	data := config.ServiceProps{
		ServicesDir:       servicesDir,
		ServiceSubDir:     serviceSubDir,
		ServiceName:       serviceName,
		ServicePath:       servicePath,
		ServiceMethodName: serviceMethodName,
		MakeUrlFileName:   makeUrlFileName,
		MakeUrlMethodName: makeMethodName,
		MakeUrlPath:       makeUrlPath,
		IndexFilePath:     indexPath,
		TestSubDir:        testSubDir,
		TestFilePath:      testFilePath,
		TestServiceTitle:  testServiceTitle,
	}

	if err := createServiceDir(&data); err != nil {
		return err
	}

	if err := createServiceFile(&data); err != nil {
		return err
	}

	if err := createServiceTestDir(&data); err != nil {
		return err
	}

	if err := createServiceTestFile(&data); err != nil {
		return err
	}

	if err := createBaseUrlFile(&data); err != nil {
		return err
	}

	if err := createIndexFile(&data); err != nil {
		return err
	}

	return nil
}
