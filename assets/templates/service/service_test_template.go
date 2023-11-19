package service

import (
	"fmt"
	"text/template"
)

type ServiceTestTemplateFields struct {
	ServiceName  string
	ServicePath  string
	ServiceTitle string
}

func GenerateServiceTestTemplate() (*template.Template, error) {
	tmpl, err := template.New("service").Parse(`import { describe, expect, it, vi } from 'vitest'
	import { AxiosHttpClientAdapter } from '@/services/axios/AxiosHttpClientAdapter'
	import * as Errors from '@/services/axios/errors'
	import { {{ .ServiceName }} } from '{{ .ServicePath }}'
	
	const basePayloadMock = {
		key: 'value'
	}
	
	const makeSut = () => {
		const sut = {{ .ServiceName }}
	
		return {
			sut
		}
	}
	
	describe('{{ .ServiceTitle }}', () => {
		it('should call API with correct params', async () => {
			const requestSpy = vi.spyOn(AxiosHttpClientAdapter, 'request').mockResolvedValueOnce({
				statusCode: 200,
				body: basePayloadMock
			})

			const { sut } = makeSut()
	
			await sut()
	
			expect(requestSpy).toHaveBeenCalledWith({
				url: '',
				method: 'GET',
				body: basePayloadMock
			})
		})
	
		it.each([
			{
				statusCode: 400,
				expectedError: new Errors.InvalidApiRequestError().message
			},
			{
				statusCode: 403,
				expectedError: new Errors.ForbiddenError().message
			},
			{
				statusCode: 404,
				expectedError: new Errors.NotFoundError().message
			},
			{
				statusCode: 500,
				expectedError: new Errors.InternalServerError().message
			},
			{
				statusCode: 'unmappedStatusCode',
				expectedError: new Errors.UnexpectedError().message
			}
		])(
			'should throw when request fails with status code $statusCode',
			async ({ statusCode, expectedError }) => {
				vi.spyOn(AxiosHttpClientAdapter, 'request').mockResolvedValueOnce({
					statusCode
				})
				const { sut } = makeSut()
	
				const request = sut()
	
				expect(request).rejects.toBe(expectedError)
			}
		)
	})
`)

	if err != nil {
		return nil, fmt.Errorf("service test template: %s", err.Error())
	}

	return tmpl, nil
}
