"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
import requests as requests_http
from ..shared import wallet as shared_wallet
from typing import Optional



@dataclasses.dataclass
class ListWalletsForAccountRequest:
    account_id: str = dataclasses.field(metadata={'path_param': { 'field_name': 'accountID', 'style': 'simple', 'explode': False }})
    r"""ID of the account"""
    




@dataclasses.dataclass
class ListWalletsForAccountResponse:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    raw_response: Optional[requests_http.Response] = dataclasses.field(default=None)
    wallets: Optional[list[shared_wallet.Wallet]] = dataclasses.field(default=None)
    r"""Wallets associated with the given account"""
    

