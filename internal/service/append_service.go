package service

import (
	"github.com/william.tome/pk-cli/cmd/config"
	"github.com/william.tome/pk-cli/tools"
)

func AppendService(name, serviceSubDirName string) error {
	servicesDir := "src/services/"
	subDirName := tools.HandleStringSuffix(serviceSubDirName, "-services")
	serviceSubDir := servicesDir + subDirName
	serviceName := tools.HandleStringSuffix(name, "-service")
	serviceMethodName := tools.ConvertKebabToCamel(serviceName)
	servicePath := serviceSubDir + "/" + serviceName + ".js"
	makeUrlFileName := "make-" + name + "-base-url"
	makeUrlPath := serviceSubDir + "/" + makeUrlFileName + ".js"
	makeMethodName := tools.ConvertKebabToCamel(makeUrlFileName)
	indexPath := serviceSubDir + "/index.js"
	testSubDir := "src/tests/services/" + serviceSubDirName
	testFilePath := testSubDir + "/" + serviceName + ".test.js"
	testServiceName := tools.ConvertKebabToPascal(subDirName)

	data := config.ServiceProps{
		ServicesDir:       servicesDir,
		ServiceSubDir:     serviceSubDir,
		ServiceName:       serviceName,
		ServiceMethodName: serviceMethodName,
		ServicePath:       servicePath,
		MakeUrlFileName:   makeUrlFileName,
		MakeUrlPath:       makeUrlPath,
		MakeUrlMethodName: makeMethodName,
		IndexFilePath:     indexPath,
		TestSubDir:        testSubDir,
		TestFilePath:      testFilePath,
		TestServiceTitle:  testServiceName,
	}

	if err := createServiceFile(&data); err != nil {
		return err
	}

	if err := createBaseUrlFile(&data); err != nil {
		return err
	}

	if err := createServiceTestFile(&data); err != nil {
		return err
	}

	if err := appendServiceToIndex(&data); err != nil {
		return err
	}

	return nil
}
