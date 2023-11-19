package service

import (
	"fmt"
	"text/template"
)

type ServiceTemplateFields struct {
	ServiceBaseUrl     string
	ServiceBaseUrlPath string
	ServiceName        string
}

func GenerateServiceTemplate() (*template.Template, error) {
	tmpl, err := template.New("service").Parse(`import { AxiosHttpClientAdapter } from '@/services/axios/AxiosHttpClientAdapter'
import * as Errors from '@/services/axios/errors'
import { {{ .ServiceBaseUrl }} } from '{{ .ServiceBaseUrlPath }}'

export const {{ .ServiceName }} = async (payload) => {
  let httpResponse = await AxiosHttpClientAdapter.request({
    url: {{ .ServiceBaseUrl }}(),
    method: 'GET',
    body: payload
  })
  return parseHttpResponse(httpResponse)
}

/**
 * @param {Object} httpResponse - The HTTP response object.
 * @param {Object} httpResponse.body - The response body.
 * @param {String} httpResponse.statusCode - The HTTP status code.
 * @returns {Object} The result message based on the status code.
 * @throws {Error} If there is an error with the response.
 */

const parseHttpResponse = (httpResponse) => {
  switch (httpResponse.statusCode) {
    case 200:
      return httpResponse.body
    case 400:
			throw new Errors.BadRequestError().message
		case 403:
			throw new Errors.ForbiddenError().message
    case 404:
      throw new Errors.NotFoundError().message
    case 500:
      throw new Errors.InternalServerError().message
    default:
      throw new Errors.UnexpectedError().message
  }
}
`)

	if err != nil {
		return nil, fmt.Errorf("service template: %s", err.Error())
	}

	return tmpl, nil
}
