"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
from ..shared import cardverificationresult as shared_cardverificationresult
from dataclasses_json import Undefined, dataclass_json
from petstore import utils
from typing import Optional


@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class CardVerifications:
    r"""The results of submitting cardholder data to a card network for verification"""
    address_line1: Optional[shared_cardverificationresult.CardVerificationResult] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('addressLine1'), 'exclude': lambda f: f is None }})
    cvv: Optional[shared_cardverificationresult.CardVerificationResult] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('cvv'), 'exclude': lambda f: f is None }})
    postal_code: Optional[shared_cardverificationresult.CardVerificationResult] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('postalCode'), 'exclude': lambda f: f is None }})
    

