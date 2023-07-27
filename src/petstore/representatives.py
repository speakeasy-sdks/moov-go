"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from .sdkconfiguration import SDKConfiguration
from petstore import utils
from petstore.models import errors, operations, shared
from typing import Optional

class Representatives:
    r"""We think of a representative as an individual who can take major actions on behalf of a business. A representative can be the business owner, or anyone who owns 25% or more of the business. Some examples of representatives are the CEO, CFO, COO, or president. We require all business accounts to have valid information for at least one representative. Moov accounts must have verified business representatives before a business account can send funds, collect money from other accounts, or store funds in a wallet. To learn more, read our guide on [business representatives](https://docs.moov.io/guides/accounts/business-representatives/)."""
    sdk_configuration: SDKConfiguration

    def __init__(self, sdk_config: SDKConfiguration) -> None:
        self.sdk_configuration = sdk_config
        
    
    def create(self, request: operations.CreateRepresentativeRequest) -> operations.CreateRepresentativeResponse:
        r"""Create representative
        Moov accounts associated with businesses require information regarding individuals who represent the business. You can provide this information by creating a representative. Each account is allowed a maximum of 7 representatives.<br><br> To create a representative, you must specify the `/accounts/{accountID}/representatives.write` scope.
        """
        base_url = utils.template_url(*self.sdk_configuration.get_server_details())
        
        url = utils.generate_url(operations.CreateRepresentativeRequest, base_url, '/accounts/{accountID}/representatives', request)
        headers = {}
        req_content_type, data, form = utils.serialize_request_body(request, "create_representative", 'json')
        if req_content_type not in ('multipart/form-data', 'multipart/mixed'):
            headers['content-type'] = req_content_type
        if data is None and form is None:
            raise Exception('request body is required')
        headers['Accept'] = 'application/json'
        headers['user-agent'] = f'speakeasy-sdk/{self.sdk_configuration.language} {self.sdk_configuration.sdk_version} {self.sdk_configuration.gen_version} {self.sdk_configuration.openapi_doc_version}'
        
        client = self.sdk_configuration.security_client
        
        http_res = client.request('POST', url, data=data, files=form, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.CreateRepresentativeResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.Representative])
                res.representative = out
            else:
                raise errors.SDKError(f'unknown content-type received: {content_type}', http_res.status_code, http_res.text, http_res)
        else:
            pass

        return res

    
    def delete(self, request: operations.DeleteRepresentativeRequest) -> operations.DeleteRepresentativeResponse:
        r"""Delete a representative
        Deletes a business representative associated with a Moov account. <br><br> To use this endpoint, you'll need to specify the `/accounts/{accountID}/representatives.write` scope.
        """
        base_url = utils.template_url(*self.sdk_configuration.get_server_details())
        
        url = utils.generate_url(operations.DeleteRepresentativeRequest, base_url, '/accounts/{accountID}/representatives/{representativeID}', request)
        headers = {}
        headers['Accept'] = '*/*'
        headers['user-agent'] = f'speakeasy-sdk/{self.sdk_configuration.language} {self.sdk_configuration.sdk_version} {self.sdk_configuration.gen_version} {self.sdk_configuration.openapi_doc_version}'
        
        client = self.sdk_configuration.security_client
        
        http_res = client.request('DELETE', url, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.DeleteRepresentativeResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        

        return res

    
    def get(self, request: operations.GetRepresentativeRequest) -> operations.GetRepresentativeResponse:
        r"""Get representative
        Retrieve a specific representative associated with a given Moov account. <br><br> To get a representative, you'll need to specify the `/accounts/{accountID}/representatives.read` scope.
        """
        base_url = utils.template_url(*self.sdk_configuration.get_server_details())
        
        url = utils.generate_url(operations.GetRepresentativeRequest, base_url, '/accounts/{accountID}/representatives/{representativeID}', request)
        headers = {}
        headers['Accept'] = 'application/json'
        headers['user-agent'] = f'speakeasy-sdk/{self.sdk_configuration.language} {self.sdk_configuration.sdk_version} {self.sdk_configuration.gen_version} {self.sdk_configuration.openapi_doc_version}'
        
        client = self.sdk_configuration.security_client
        
        http_res = client.request('GET', url, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.GetRepresentativeResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.Representative])
                res.representative = out
            else:
                raise errors.SDKError(f'unknown content-type received: {content_type}', http_res.status_code, http_res.text, http_res)
        else:
            pass

        return res

    
    def list(self, request: operations.ListRepresentativesRequest) -> operations.ListRepresentativesResponse:
        r"""List representatives
        A Moov account may have multiple representatives depending on the associated business's ownership and management structure. You can use this method to list all the representatives for a given Moov account. Note that Moov accounts associated with an individual do not have representatives. <br><br> To list representatives, you need to specify the `/accounts/{accountID}/representatives.read` scope.
        """
        base_url = utils.template_url(*self.sdk_configuration.get_server_details())
        
        url = utils.generate_url(operations.ListRepresentativesRequest, base_url, '/accounts/{accountID}/representatives', request)
        headers = {}
        headers['Accept'] = 'application/json'
        headers['user-agent'] = f'speakeasy-sdk/{self.sdk_configuration.language} {self.sdk_configuration.sdk_version} {self.sdk_configuration.gen_version} {self.sdk_configuration.openapi_doc_version}'
        
        client = self.sdk_configuration.security_client
        
        http_res = client.request('GET', url, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.ListRepresentativesResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[list[shared.Representative]])
                res.representatives = out
            else:
                raise errors.SDKError(f'unknown content-type received: {content_type}', http_res.status_code, http_res.text, http_res)
        else:
            pass

        return res

    
    def update(self, request: operations.PatchRepresentativeRequest) -> operations.PatchRepresentativeResponse:
        r"""Patch representative
        If a representative's information has changed you can patch the information associated with a specific representative ID.  
        To patch a representative, you'll need to specify the `/accounts/{accountID}/representatives.write` scope and provide the changed information.

        When **can** profile data be updated:  
          + For unverified representatives, all profile data can be edited.
          + During the verification process, missing or incomplete profile data can be edited.
          + Verified representatives can only add missing profile data.  

          When **can't** profile data be updated:  
          + Verified representatives cannot change any existing profile data.  

        If you need to update information in a locked state, please contact Moov support.
        """
        base_url = utils.template_url(*self.sdk_configuration.get_server_details())
        
        url = utils.generate_url(operations.PatchRepresentativeRequest, base_url, '/accounts/{accountID}/representatives/{representativeID}', request)
        headers = {}
        req_content_type, data, form = utils.serialize_request_body(request, "patch_representative_request", 'json')
        if req_content_type not in ('multipart/form-data', 'multipart/mixed'):
            headers['content-type'] = req_content_type
        if data is None and form is None:
            raise Exception('request body is required')
        headers['Accept'] = 'application/json'
        headers['user-agent'] = f'speakeasy-sdk/{self.sdk_configuration.language} {self.sdk_configuration.sdk_version} {self.sdk_configuration.gen_version} {self.sdk_configuration.openapi_doc_version}'
        
        client = self.sdk_configuration.security_client
        
        http_res = client.request('PATCH', url, data=data, files=form, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.PatchRepresentativeResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.Representative])
                res.representative = out
            else:
                raise errors.SDKError(f'unknown content-type received: {content_type}', http_res.status_code, http_res.text, http_res)
        else:
            pass

        return res

    