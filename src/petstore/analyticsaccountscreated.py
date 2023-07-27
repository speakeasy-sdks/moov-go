"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from .sdkconfiguration import SDKConfiguration
from petstore import utils
from petstore.models import operations, shared

class AnalyticsAccountsCreated:
    sdk_configuration: SDKConfiguration

    def __init__(self, sdk_config: SDKConfiguration) -> None:
        self.sdk_configuration = sdk_config
        
    
    def post(self, request: shared.TimeRange) -> operations.PostAnalyticsAccountsCreatedResponse:
        r"""Count the number of profiles created by an individual or business"""
        base_url = utils.template_url(*self.sdk_configuration.get_server_details())
        
        url = base_url + '/analytics/accounts/profiles-created'
        headers = {}
        req_content_type, data, form = utils.serialize_request_body(request, "request", 'json')
        if req_content_type not in ('multipart/form-data', 'multipart/mixed'):
            headers['content-type'] = req_content_type
        if data is None and form is None:
            raise Exception('request body is required')
        headers['Accept'] = '*/*'
        headers['user-agent'] = f'speakeasy-sdk/{self.sdk_configuration.language} {self.sdk_configuration.sdk_version} {self.sdk_configuration.gen_version} {self.sdk_configuration.openapi_doc_version}'
        
        client = self.sdk_configuration.security_client
        
        http_res = client.request('POST', url, data=data, files=form, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.PostAnalyticsAccountsCreatedResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        

        return res

    