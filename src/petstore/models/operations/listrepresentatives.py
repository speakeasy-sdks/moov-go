"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
import requests as requests_http
from ..shared import representative as shared_representative
from typing import Optional



@dataclasses.dataclass
class ListRepresentativesRequest:
    account_id: str = dataclasses.field(metadata={'path_param': { 'field_name': 'accountID', 'style': 'simple', 'explode': False }})
    r"""ID of the account"""
    




@dataclasses.dataclass
class ListRepresentativesResponse:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    raw_response: Optional[requests_http.Response] = dataclasses.field(default=None)
    representatives: Optional[list[shared_representative.Representative]] = dataclasses.field(default=None)
    r"""Successfully retrieved representatives"""
    

