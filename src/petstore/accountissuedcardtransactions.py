"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from .sdkconfiguration import SDKConfiguration
from petstore import utils
from petstore.models import errors, operations, shared
from typing import Optional

class AccountIssuedCardTransactions:
    sdk_configuration: SDKConfiguration

    def __init__(self, sdk_config: SDKConfiguration) -> None:
        self.sdk_configuration = sdk_config
        
    
    def list(self, request: operations.ListAccountIssuedCardTransactionsRequest) -> operations.ListAccountIssuedCardTransactionsResponse:
        r"""Get account transactions
        List issued card transactions associated with a Moov account
        """
        base_url = utils.template_url(*self.sdk_configuration.get_server_details())
        
        url = utils.generate_url(operations.ListAccountIssuedCardTransactionsRequest, base_url, '/issuing/{accountID}/transactions', request)
        headers = {}
        query_params = utils.get_query_params(operations.ListAccountIssuedCardTransactionsRequest, request)
        headers['Accept'] = 'application/json'
        headers['user-agent'] = f'speakeasy-sdk/{self.sdk_configuration.language} {self.sdk_configuration.sdk_version} {self.sdk_configuration.gen_version} {self.sdk_configuration.openapi_doc_version}'
        
        client = self.sdk_configuration.security_client
        
        http_res = client.request('GET', url, params=query_params, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.ListAccountIssuedCardTransactionsResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[list[shared.IssuedCardTransaction]])
                res.issued_card_transactions = out
            else:
                raise errors.SDKError(f'unknown content-type received: {content_type}', http_res.status_code, http_res.text, http_res)
        else:
            pass

        return res

    