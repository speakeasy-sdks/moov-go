"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from .sdkconfiguration import SDKConfiguration
from petstore import utils
from petstore.models import errors, operations, shared
from typing import Optional

class IssuedCard:
    sdk_configuration: SDKConfiguration

    def __init__(self, sdk_config: SDKConfiguration) -> None:
        self.sdk_configuration = sdk_config
        
    
    def get(self, request: operations.GetIssuedCardRequest) -> operations.GetIssuedCardResponse:
        r"""Get issued card
        Retrieve a single issued card associated with a Moov account
        <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/issued-cards.read` scope.
        """
        base_url = utils.template_url(*self.sdk_configuration.get_server_details())
        
        url = utils.generate_url(operations.GetIssuedCardRequest, base_url, '/issuing/{accountID}/issued-cards/{issuedCardID}', request)
        headers = {}
        headers['Accept'] = 'application/json'
        headers['user-agent'] = f'speakeasy-sdk/{self.sdk_configuration.language} {self.sdk_configuration.sdk_version} {self.sdk_configuration.gen_version} {self.sdk_configuration.openapi_doc_version}'
        
        client = self.sdk_configuration.security_client
        
        http_res = client.request('GET', url, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.GetIssuedCardResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.IssuedCard])
                res.issued_card = out
            else:
                raise errors.SDKError(f'unknown content-type received: {content_type}', http_res.status_code, http_res.text, http_res)
        else:
            pass

        return res

    
    def update(self, request: operations.UpdateIssuedCardRequest) -> operations.UpdateIssuedCardResponse:
        r"""Update issued card
        Update a Moov issued card
        <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/issued-cards.write` scope.
        """
        base_url = utils.template_url(*self.sdk_configuration.get_server_details())
        
        url = utils.generate_url(operations.UpdateIssuedCardRequest, base_url, '/issuing/{accountID}/issued-cards/{issuedCardID}', request)
        headers = {}
        req_content_type, data, form = utils.serialize_request_body(request, "update_issued_card", 'json')
        if req_content_type not in ('multipart/form-data', 'multipart/mixed'):
            headers['content-type'] = req_content_type
        if data is None and form is None:
            raise Exception('request body is required')
        headers['Accept'] = '*/*'
        headers['user-agent'] = f'speakeasy-sdk/{self.sdk_configuration.language} {self.sdk_configuration.sdk_version} {self.sdk_configuration.gen_version} {self.sdk_configuration.openapi_doc_version}'
        
        client = self.sdk_configuration.security_client
        
        http_res = client.request('PATCH', url, data=data, files=form, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.UpdateIssuedCardResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        

        return res

    